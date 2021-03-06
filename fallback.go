package main

import (
	"fmt"
	"strconv"

	"github.com/skratchdot/open-golang/open"
)

func createFallbackUI() {
	fmt.Println(header)

	shouldUpdate, newTag, err := checkForUpdate()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		panic(err)
	}

	if shouldUpdate {
		var choice string
		fmt.Printf("There is a new shadowfox-updater version available (%s -> %s)\nDo you want to update? [y/n]", tag, newTag)
		fmt.Scanln(&choice)
		wantToUpdate := (choice == "y" || choice == "Y")
		if wantToUpdate {
			open.Start("https://github.com/SrKomodo/shadowfox-updater/#installing")
		}
	}

	var choice string
	paths, names, err := getProfilePaths()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		panic(err)
	}

	if paths == nil {
		fmt.Println("ShadowFox couldn't automatically find 'profiles.ini'. Please follow these steps:")
		fmt.Println("  1. Close the program")
		fmt.Println("  2. Move the program to the folder 'profiles.ini' is located")
		fmt.Println("  3. Run the program")
		fmt.Scanln()
		return
	}

	fmt.Println("Available profiles:")
	for i, name := range names {
		fmt.Printf("  %d: %s\n", i, name)
	}

	fmt.Printf("\nWhich one would you like to use? [%d-%d] ", 0, len(names)-1)

	var profile string
	for {
		fmt.Scanln(&choice)
		i, err := strconv.Atoi(choice)
		if err != nil || i < 0 || i > len(paths) {
			fmt.Print("Please input a valid number ")
		} else {
			profile = paths[i]
			break
		}
	}

	fmt.Print("\nDo you want to (1) install or (2) uninstall ShadowFox? [1/2] ")
	fmt.Scanln(&choice)

	if choice == "2" {
		uninstall(profile)
		fmt.Print("\nShadowFox was successfully uninstalled! (Press 'enter' to exit)")
		fmt.Scanln()
		return
	}

	fmt.Print("\nWould you like to auto-generate UUIDs? [y/n] ")
	fmt.Scanln(&choice)
	uuids := (choice == "y" || choice == "Y")

	fmt.Print("\nWould you like to automatically set the Firefox dark theme? [y/n] ")
	fmt.Scanln(&choice)
	theme := (choice == "y" || choice == "Y")

	err = install(profile, uuids, theme)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		panic(err)
	}

	fmt.Print("\nShadowFox was successfully installed! (Press 'enter' to exit)")
	fmt.Scanln()
	return
}
