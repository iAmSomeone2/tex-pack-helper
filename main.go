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
	"strings"

	"github.com/iAmSomeone2/texpackhelper/list"
)

func main() {
	// String constants for the command-line flags
	const fileFlagDesc string = "The destination of the output file."
	const recursiveFlagDesc string = "Set this flag to traverse all files in the directory."
	const extFlagDesc string = "Only lists the files with the specified file extensions.\nMultiple extensions can be separated with ','"

	// All flag pointers must be dereferenced to be used
	directoryPtr := flag.String("dir", "./", "The directory to make the list for.")
	recursiveTruePtr := flag.Bool("recursive", false, recursiveFlagDesc)
	fileNamePtr := flag.String("out", "./file_list.txt", fileFlagDesc)
	extPtr := flag.String("ext", "", extFlagDesc) // Read through this flag to get comma separated values

	flag.Parse()

	fmt.Println("\nReading from: " + *directoryPtr + "\n")

	fmt.Println("Working...")
	// Get all of the files in a slice
	foundFiles := list.TraverseFolder(*directoryPtr, *recursiveTruePtr)
	// If the ext flag is set, filer out everything that doesn't match it
	if *extPtr != "" {
		// Run filter function
		exts := strings.Split(*extPtr, ",")
		fmt.Println("Extensions to look for:")
		fmt.Println(exts)
		foundFiles = list.FilterExt(exts, foundFiles)
	}

	err := list.WriteList(*fileNamePtr, foundFiles)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAll found files are listed in: " + *fileNamePtr)
}
