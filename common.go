package main

import (
	"flag"
	"os"
	"path/filepath"
)

var configFile *string = flag.String("config-file", "", "Location of the config file")

func importConfig() *Config {
	if len(*configFile) == 0 {
		homeDir := os.Getenv("HOME")
		if len(homeDir) == 0 {
			// alert(term, "$HOME not set. Please either export $HOME or use the -config-file option.\n")
			// TODO throw error
		}
		persistentDir := filepath.Join(homeDir, "Persistent")
		if stat, err := os.Lstat(persistentDir); err == nil && stat.IsDir() {
			// Looks like Tails.
			homeDir = persistentDir
		}
		*configFile = filepath.Join(homeDir, ".xmpp-gui")
	}

	config, err := ParseConfig(*configFile)
	
	if err != nil {
		// alert(term, "Failed to parse config file: "+err.Error())
		//config = new(Config)
		// if !enroll(config, term) {
		// 	return
		//}
		config.filename = *configFile
		config.Save()
	}

	return config
}
