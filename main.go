package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type Manifest struct {
	Description string `json:"description"`
	Home        string `json:"homepage"`
	Version     string `json:"version"`
	ModTime     int64
	Bucket      string
	App         string
	Installed   bool
	Updated     bool
	Hold        bool
}

type InstallJson struct {
	Hold bool `json:"hold"`
}

func scoopUpdate() string {
	var buf bytes.Buffer
	cmd := exec.Command("scoop", "update")
	//f, err := os.Create("scoop.out")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	w := io.MultiWriter(&buf, os.Stdout) //, f)
	//cmd.Stdout = os.Stdout
	cmd.Stdout = w

	if err := cmd.Run(); err != nil {
		log.Fatal(err, " cmd.Run() failed.")
	}
	return buf.String()
}

func hasArg(arg string) bool {
	if len(os.Args) > 1 {
		for _, a := range os.Args {
			if a == arg {
				return true
			}
		}
	}
	return false
}

func main() {

	var err error
	scoopRoot := path.Join(os.Getenv("USERPROFILE"), "scoop")
	bucketDir := scoopRoot + "/buckets"

	fzfpath, err := exec.LookPath("fzf")
	if err != nil {
		fmt.Println("Could not find fzf.")
		os.Exit(1)
	}

	log.SetFlags(log.Lshortfile)
	var (
		//ColorRed    = "[91m"
		ColorGreen  = "[92m"
		ColorYellow = "[93m"
		ColorCyan   = "[96m"
	)

	var buf string
	if hasArg("u") || hasArg("update") {
		buf = scoopUpdate()
	} else if hasArg("-p") {
		fmt.Println("Reading previous output.")
		b, err := os.ReadFile("scoop.out")
		if err != nil {
			log.Fatal(err)
		}
		buf = string(b)
		fmt.Println(buf)
		pressEnterToContinue()
	}
	//log.Println("buf:", buf)

	var updated []string
	for _, line := range strings.Split(buf, "\n") {
		//log.Println(line)
		if strings.HasPrefix(line, " * ") {
			app := strings.Split(line, " ")[3]
			app = strings.TrimSuffix(app, ":")
			app = strings.Split(app, "@")[0]
			updated = append(updated, app)
		}
	}
	//log.Println("updated:", updated)
	if len(updated) != 0 {
		pressEnterToContinue()
	}

	fmt.Println("Building app list..")
	bucketNames, err := ioutil.ReadDir(bucketDir)
	if err != nil {
		log.Fatal(err, "could not read bucket dir:", bucketDir)
	}

	app_dirs, err := os.ReadDir(scoopRoot + "/apps")
	if err != nil {
		log.Fatal(err, "Could not read installed apps dir")
	}

	var installed []string
	for _, f := range app_dirs {
		installed = append(installed, f.Name())
	}
	// scan for held apps
	var held []string
	for _, appdir := range app_dirs {
		// log.Println("appdir:", appdir.Name())
		installFile := filepath.Join(scoopRoot, "apps", appdir.Name(), "current", "install.json")
		jsonStr, err := os.ReadFile(installFile)
		if err != nil {
			log.Println(err)
			log.Println("Could not read install file", installFile)
			log.Println("installFile:", installFile)
			continue
		}
		// log.Println("jsonStr:", string(jsonStr))
		var jsn InstallJson
		err = json.Unmarshal(jsonStr, &jsn)
		if err != nil {
			log.Println("Could not unmarshal install json")
		}
		// log.Println("hold:", jsn.Hold)
		if jsn.Hold {
			held = append(held, appdir.Name())
		}

	}

	var entries []Manifest

	var bucketLocation = make(map[string]string)

	for _, f := range bucketNames {
		bucketName := f.Name()
		bdir := bucketDir + "/" + bucketName + "/bucket"
		bucketLocation[bucketName] = bdir
		files, err := os.ReadDir(bdir)
		if err != nil {
			log.Println("Could not read bucket dir.", bdir, err)
		}
		var manifest Manifest
		for _, v := range files {
			path := bdir + "/" + v.Name()
			if strings.HasSuffix(v.Name(), ".json") {
				appName := strings.TrimSuffix(v.Name(), filepath.Ext(v.Name()))
				//log.Println("appName:", appName)
				contents, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err, "could not read file", path)
				}
				err = json.Unmarshal(contents, &manifest)
				if err != nil {
					log.Println("skipping", appName, "Could not unmarshal json string.\n", string(contents))
					continue
				}
				info, err := v.Info()
				if err != nil {
					log.Fatal(err, "could not get info")
				}
				manifest.ModTime = info.ModTime().Unix()
				manifest.Bucket = bucketName
				manifest.App = appName
				manifest.Installed = contains(installed, appName)
				manifest.Updated = contains(updated, appName)
				manifest.Hold = contains(held, appName)

				entries = append(entries, manifest)
			}
		}
	}

	// sort.Slice(entries, func(i, j int) bool {
	// 	return entries[i].ModTime > entries[j].ModTime
	// })
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].App < entries[j].App
	})

	var sb strings.Builder
	var color string
	for _, manifest := range entries {
		if manifest.Installed {
			sb.WriteString("i")
		} else {
			sb.WriteString(" ")
		}
		if manifest.Updated {
			sb.WriteString("u")
		} else {
			sb.WriteString(" ")
		}
		if manifest.Hold {
			sb.WriteString("h")
		} else {
			sb.WriteString(" ")
		}

		fmt.Fprintf(&sb, "| %8s", manifest.Bucket)
		color = ""
		if manifest.Installed {
			color = ColorGreen
		}
		if manifest.Updated {
			if manifest.Installed {
				color = ColorYellow
			} else {
				color = ColorCyan
			}
		}
		fmt.Fprintf(&sb, "| %s%-25s[0m", color, manifest.App)
		sb.WriteString(
			fmt.Sprintf("| %s |%s|%s\n",
				manifest.Description,
				manifest.Home,
				manifest.Version))
	}

	// Call fzf with sb as stdin

	cmd := exec.Command(fzfpath,
		"--ansi", "--reverse", `--delimiter=\|`,
		"--no-hscroll", "--no-sort", "--multi",
		"--header", "F1:Homepage | F2:Install | F3:Uninstall | F4:Update | F5:Hold | F6:Unhold | Esc:Exit",
		"--bind", "f1:execute(scoop home {3})",
		"--bind", "f2:execute(scoop install {+3})",
		"--bind", "f3:execute(scoop uninstall {+3})",
		"--bind", "f4:execute(scoop update {+3})",
		"--bind", "f5:execute(scoop hold {+3})",
		"--bind", "f6:execute(scoop unhold {+3})",
		"--bind", "ctrl-a:toggle-all",
		"--bind", "change:top",
		"--preview-window=top:4:wrap",
		"--preview=echo {3} {-1} {2} && echo {5} && echo {4}")
	//log.Println(cmd.String())
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err, " could not pipe stdin")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, sb.String())
	}()

	if err = cmd.Run(); err != nil {
		log.Printf("%v could not get combined output.\n", err)
	}
}

func pressEnterToContinue() {
	fmt.Print("\n> Press Enter to continue ")
	fmt.Scanln()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			//log.Println(a, "==", e)
			return true
		}
	}
	return false
}

func RemoveIndex(items []string, index int) []string {
	return append(items[:index], items[index+1:]...)
}

// StringsFromFile()
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// StringsToFile
func WriteLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
