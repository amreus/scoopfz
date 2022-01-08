# scoopfz
fuzzy search your local scoop apps and descriptions.

Use Function keys to install & uninstall, update, and hold and unhold selected apps.


## Usage

    > scoopfz
    
`scoopfz` requires `fzf` (scoop install fzf) https://github.com/junegunn/fzf

Installed apps are colored green.

To view installed apps, search for `^i`

If you run `scoopfz u` scoopfz will run `scoop update` before listing the results. Updated apps are shown in cyan. Show only updated apps by using the search term of `^u`.

You can search for both insalled and updated apps using `^iu`

Keys:

```
F1  - `scoop home <selected>`
F2  - `scoop install <selected>` - use the tab key to select multiple apps
F3  - `scoop uninstll <selected>`
F4  - `scoop update <selected>`
F5  - `scoop hold <selected>`
F6  - `scoop unhold <selected>`
Esc - quit
```
