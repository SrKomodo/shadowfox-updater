# Shadowfox Updater

This is a cross-platform installer/uninstaller/updater for [Shadowfox](https://github.com/overdodactyl/ShadowFox), a universal dark theme for Firefox.

## Installing

- For all platforms: go to the [latest release](https://github.com/SrKomodo/shadowfox-updater/releases/latest) and download the respective file for your OS
  - If you are in Linux or Mac, you will probably need to run `chmod +x [filename]` for the OS to register it as an executable
- On Arch Linux you can install the package `shadowfox-updater` from AUR
- On MacOS, you can install with either [Homebrew](https://brew.sh/) or [MacPorts.](https://www.macports.org/) 

  - Homebrew installation:

  ```
  $ brew install srkomodo/tap/shadowfox-updater
  $ shadowfox
  ```

  - MacPorts installation:

  ```
  $ sudo port install shadowfox-updater
  $ shadowfox-updater
  ```

## How to use

There are various ways to use Shadowfox Updater

### GUI Mode

If you run the file from the command line, it will show you a text-based UI. You can use `TAB` to move between the different options (and `SHIFT+TAB` on some terminals to move backwards), and you can use `ENTER` to toggle the checkboxes and press the buttons.

The "Profile to use" dropdown will show all available profiles in which to install to, and you can cycle through them with the arrow keys.

The "Auto-Generate UUIDs" checkbox, if toggled, will make the updater automatically populate the `internal_UUIDs.txt` file, which is used for styling of extensions. Generally you would toggle this unless you want to manage precisely which extensions get styled.

The "Set Firefox dark theme" checkbox, if toggled, will make the updater automatically enable Firefox's dark theme for it's UI and devtools. If you already have the dark theme enabled, you shouldn't toggle this one.

Then the "Install/Update Shadowfox", "Uninstall Shadowfox" and "Exit" buttons are pretty self explanatory.

#### Fallback

If the text-based UI fails to load (something that happens in some terminals), the program will load a more basic text-only prompt that has the same features as the usual GUI but without the fancy buttons and dropdowns.

### CLI Mode

If you run the file with one or more arguments, the updater will work as a command line tool, which can be useful for automated scripts and such. Instead of explaining how it works I'm just going to paste the result of the command `shadowfox-updater -h`.

```
Usage of shadowfox-updater:
  -generate-uuids
    	Wheter to automatically generate UUIDs or not
  -profile-index int
    	Index of profile to use
  -profile-name string
    	Name of profile to use, if not defined or not found will fallback to profile-index
  -set-dark-theme
    	Wheter to automatically set Firefox's dark theme
  -uninstall
    	Wheter to install or uninstall ShadowFox
```

## Common issues

### ShadowFox couldn't automatically find 'profiles.ini'

If this error shows up then your Firefox installation is probably located in a non-standard location. In this case, the solution would be to move the shadowfox executable to wherever `profiles.ini` is located.

1. Open Firefox and go to `about:profiles`
2. Click "Open root folder"
3. Go back a few folders until you see `profiles.ini`
4. Copy the updater executable to where `profiles.ini` is located
5. Run the updater again

### Couldn't read prefs.js: no such file or directory

This issue can happen if the profile you are trying to install to hasn't ever been opened. It can be easily fixed by just running Firefox with that profile and then running the updater again.

### panic: key-value delimiter not found

This issue usually happens because `profiles.ini` is encoded in some encoding different from UTF-8, this can be easily fixed by changing `profiles.ini`'s encoding to UTF-8 with your favorite text editor or command line tool of choice.
