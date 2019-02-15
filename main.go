/*
	texpackhelper is designed to aid in the creation of neural network-enhanced
	video game texture packs. This is the main entry point for the program.

	Author: Brenden Davidson
	Date: 2019-02-14
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/iAmSomeone2/texpackhelper/imageproc"
	"github.com/iAmSomeone2/texpackhelper/list"
	// "github.com/iAmSomeone2/texpackhelper/preset"
)

func main() {
	// String constants for the command-line flags

	const fileFlagDesc string = "The destination of the output file."
	const recursiveFlagDesc string = "Set this flag to traverse all files in the directory."
	const extFlagDesc string = "Only lists the files with the specified file extensions.\nMultiple extensions can be separated with ','"

	// Get relevant env variables
	// homeDir := os.Getenv("HOME")
	workDir := os.Getenv("PWD")

	// All flag pointers must be dereferenced to be used
	directoryPtr := flag.String("dir", "./", "The directory to make the list for.")
	recursiveTruePtr := flag.Bool("recursive", false, recursiveFlagDesc)
	fileNamePtr := flag.String("out", "file_list.txt", fileFlagDesc)
	extPtr := flag.String("ext", ".png", extFlagDesc) // Read through this flag to get comma separated values

	flag.Parse()

	// If the directory arg is not set, join into the current path
	var fullPath string
	if *directoryPtr == "./" {
		fullPath = workDir
	} else if !path.IsAbs(*directoryPtr) {
		fullPath = path.Join(workDir, *directoryPtr)
	} else {
		fullPath = *directoryPtr
	}

	fileName := path.Join(fullPath, *fileNamePtr)

	fmt.Println("\nReading from: " + fullPath + "\n")

	fmt.Println("Working...")
	// Get all of the files in a slice
	foundFiles := list.TraverseFolder(fullPath, *recursiveTruePtr)
	// If the ext flag is set, filter out everything that doesn't match it
	if *extPtr != "" {
		// Run filter function
		exts := strings.Split(*extPtr, ",")
		fmt.Println("Extensions to look for:")
		fmt.Println(exts)
		foundFiles = list.FilterExt(exts, foundFiles)
	}

	err := list.WriteList(fileName, foundFiles)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAll found files are listed in: " + fileName + "\n")

	// TODO: Call waifu2x from here to simplify the process.
	waifu := imageproc.NewWaifu2x()
	waifu.SetListPath(fileName)
	waifuErr := waifu.Run()
	if waifuErr != nil {
		// log.Fatal(err)
	}

	// TODO: Create a preset system with dolphin-emu being the default.

	/*
		Now, a list of all of the files that were output from the enhancer
		needs to be created.
		For now, it's assumed to be the default dolphin-emu texture load folder.
		Don't forget to change the work directory as needed.
	*/

}
