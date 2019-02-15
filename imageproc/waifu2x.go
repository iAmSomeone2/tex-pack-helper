/*
	Written by Brenden Davidson
	2019-02-14
*/

package imageproc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// Waifu2x is a struct used for interacting with waifu2x functionality
type Waifu2x struct {
	installDir string
	listPath   string
	outputDir  string
	format     string
}

// NewWaifu2x creates and returns a new Waifu2x object.
// The function assumes that waifu2x is installed at $HOME/waifu2x,
// and it sets fileList to "$PWD/image_list.txt" and outputDir to
// "$PWD/output"
func NewWaifu2x() *Waifu2x {
	homeDir := os.Getenv("HOME")
	workDir := os.Getenv("PWD")

	return &Waifu2x{
		installDir: path.Join(homeDir, "waifu2x"),
		listPath:   path.Join(workDir, "image_list.txt"),
		outputDir:  path.Join(workDir, "output/"),
		format:     ".png",
	}
}

// SetInstallDir sets the directory in which to look for waifu2x.
// A fully-resolved path should be used.
func (waifu *Waifu2x) SetInstallDir(dir string) {
	waifu.installDir = path.Dir(dir)
}

// SetListPath sets the path to where the image list is located at.
// A fully-resolved path should be used.
func (waifu *Waifu2x) SetListPath(list string) {
	waifu.listPath = path.Clean(list)
}

// SetOutputDir sets the directory in which waifu2x should place its
// output files.
// A fully-resolved path should be used.
func (waifu *Waifu2x) SetOutputDir(dir string) {
	waifu.outputDir = path.Dir(dir)
}

// SetFormat sets the image output format for waifu2x to export to.
// Options are: .png or .jpg
func (waifu *Waifu2x) SetFormat(format string) {
	waifu.format = format
}

// Run finshes setting up waifu2x, then executes it.
// The set up process involves building the command, checking if the
// output folder exists, creating it if it doesn't, and running the program.
// A PathError is returned if the output folder can't be created.
func (waifu *Waifu2x) Run() error {
	const cmdFile string = "./runwaifu.sh"

	// Set up command to run
	thBinary, lookErr := exec.LookPath("th")
	if lookErr != nil {
		return lookErr
	}

	cmd := thBinary + " " + path.Join(waifu.installDir, "waifu2x.lua") + " -m scale"
	cmd += " -l " + waifu.listPath
	cmd += " -o " + path.Join(waifu.outputDir, "%s"+waifu.format)

	// fmt.Printf("\nCommand: %s\n\n", cmd)

	// Check if output folder exists, and create it if it doesn't
	if _, err := os.Stat(waifu.outputDir); err != nil {
		// Output folder does not exist, so it's created.
		createErr := os.MkdirAll(waifu.outputDir, os.ModePerm)
		if createErr != nil {
			// Folder couldn't be created.
			return createErr
		}
	}

	err := writeCmdFile(cmd, cmdFile)
	if err != nil {
		log.Println(err)
		fmt.Printf("\nDirectly calling waifu2x isn't working right now.\n")
		fmt.Printf("Run the following in another terminal, then return here where it finishes:\n\n%s", cmd)
		fmt.Println("\nPress ENTER in this terminal to continue when waifu finishes.")
		fmt.Scanln()
	} else {
		fmt.Println("Running waifu2x on the listed files...")
		command := exec.Command("bash", cmdFile)
		err := command.Run()
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

// writeCmdFile writes a shell script to execute waifu2x in a separate process.
// This may serve as a workaround to the trouble with running the command directly.
func writeCmdFile(cmd string, fileName string) error {
	var cmdFile *os.File
	defer cmdFile.Close()
	if _, err := os.Stat(fileName); err == nil {
		// Open file write-only
		cmdFile, err = os.OpenFile(fileName, os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		// Create the file
		cmdFile, err = os.Create(fileName)
		if err != nil {
			return err
		}
	}

	cmdFile.WriteString("#!/bin/bash\n\n" + cmd)

	return nil
}
