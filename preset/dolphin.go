package preset

import (
	"os"
	"path"
)

// Dolphin contains the info needed to find the texture files
// produced and used by a normal dolphin-emu install.
type Dolphin struct {
	dumpPath string
	loadPath string
}

// TODO: Add support for Windows and Mac installs

// InitDolphinPreset sets up the dophinPreset struct with the proper
// environment info and returns the struct.
func InitDolphinPreset() Dolphin {
	homeDir := os.Getenv("HOME")

	const defaultPath string = "/.local/share/dolphin-emu/"
	dump := path.Join(homeDir, defaultPath, "Dump/Textures")
	load := path.Join(homeDir, defaultPath, "Load/Textures")

	return Dolphin{
		dumpPath: dump,
		loadPath: load,
	}
}
