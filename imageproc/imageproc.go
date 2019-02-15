/*
	Package imageproc processes the images produced by the neural network.

	The first stage is to move the image types into their respective directories.
	These directories will be "masks" and "solid_color" for the time being.
	The remaining images stay in the main output directory.

	The second stage is to read through the "masks" directory.
	Images here will usually be a white mask on a black background. This is due to
	an error in the neural net where it removes transparency. Every image in this
	directory will be loaded in, the black portion of the image will be made transparent,
	and the resulting image will be saved and retain the transparency.

	The final stage is to read through the "solid_color" directory. This will be
	the most difficult stage because it involves comparing the output image with
	the original pixel-by-pixel. The goal in this step is to map the original
	image to a higher resolution.
*/

package imageproc

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"log"
	"os"
	"path"
)

// Begin starts the process of dealing with the output images.
// It should not be run in a goroutine because it manages all of the
// goroutines for processing the images.
func Begin(imageList []string) {
	workDir := path.Dir(imageList[0])
	maskDir := path.Join(workDir, "masks")
	solidDir := path.Join(workDir, "solid_color")

	// Make the "masks" and "solid_color" folders if they don't already exist.
	if _, err := os.Stat(maskDir); err == nil {
		// Folder found. Nothing needs to be done.
		log.Printf("Output dir: '%s' already exists.\n", maskDir)
	} else {
		// Folder not found. It must be created.
		createErr := os.MkdirAll(maskDir, os.ModePerm)
		if createErr != nil {
			log.Fatal(createErr)
		}
	}
	if _, err := os.Stat(solidDir); err == nil {
		// Folder found. Nothing needs to be done.
		log.Printf("Output dir: '%s' already exists.\n", solidDir)
	} else {
		// Folder not found. It must be created.
		createErr := os.MkdirAll(solidDir, os.ModePerm)
		if createErr != nil {
			log.Fatal(createErr)
		}
	}

	/*
		Start by creating a goroutine of classify() for each image in
		the list.
	*/
	for _, imageName := range imageList {
		go classify(imageName)
	}

	// Next, create file lists for the new directories

	// Run fixMask in a series of goroutines

}

// classify analyzes the image passed to it and determines which
// folder it should be placed in: "./masks", "./solid_color", or "./"
func classify(imageName string) {
	imagick.Initialize()
	defer imagick.Terminate() // Clean up imagick at the end of the function

	wand := imagick.NewMagickWand()
	defer wand.Destroy()

	// Read in image
	err := wand.ReadImage(imageName)
	if err != nil {
		log.Printf("ERROR %s not found!", imageName)
		return
	}

	// Get the number of unique colors in the image
	colorNum := wand.GetImageColors()

	if colorNum == 1 {
		// Image goes into the solid colors folder
		newLocation := path.Join("./solid_color/", imageName)
		err := os.Rename(imageName, newLocation)
		if err != nil {
			log.Fatal(err)
		}
	} else if colorNum == 2 {
		// Image goes into the mask folder
		newLocation := path.Join("./masks/", imageName)
		err := os.Rename(imageName, newLocation)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// fixMask converts black values in an image to full transparency and
// saves the new file over the old one.
func fixMask(maskName string) {

}
