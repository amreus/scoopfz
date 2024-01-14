# scoopfz

fuzzy search your local scoop apps and descriptions.

Use Function keys to install & uninstall, update, and hold and unhold selected apps.

Each time scoopfz is run, it rebuilds a list of apps from your bucket list and
start `fzf` using the list as input.


## Install (via scoop)

    scoop install https://raw.githubusercontent.com/amreus/scoopfz/main/scoopfz.json


## Usage

    scoopfz
    
`scoopfz` requires `fzf` (scoop install fzf) https://github.com/junegunn/fzf

Installed apps are colored green.

To view installed apps, search for `^i` (use ctrl-i as a shortcut to installed apps.)

To filter on a specific bucket, use a query something like `'b:bucket appname`

If you run `scoopfz u` scoopfz will run `scoop update` before listing the results. Updated apps are shown in cyan. Show only updated apps by using the search term of `^u`.

You can search for both insalled and updated apps using `^iu`

Keys:

`fzf` default key bindings, plus:

```
F1  - `scoop home <selected>`
F2  - `scoop install <selected>` - use the tab key to select multiple apps
F3  - `scoop uninstll <selected>`
F4  - `scoop update <selected>`
F5  - `scoop hold <selected>`
F6  - `scoop unhold <selected>`

ctrl-i - shortcut to filter installed apps.

Esc - quit

```

![2022-01-08_155134](https://user-images.githubusercontent.com/38442825/148660771-98c7fff0-62cc-4195-8a1d-ee76d385a09a.png)
