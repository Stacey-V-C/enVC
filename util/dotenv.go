package util

import (
	"io/fs"
	"os"
	"strings"
)

func LoadEnv() {
	fileSystem := os.DirFS("./")

	file, err := fs.ReadFile(fileSystem, ".env")
	if err != nil {
		print("Could not load .env file - using default values\n")
		return
	} else {
		lines := strings.Split(string(file), "\n")

		for _, line := range lines {
			if len(line) > 0 {
				keyValue := strings.Split(line, "=")
				os.Setenv(keyValue[0], keyValue[1])
			}
		}
	}

}
