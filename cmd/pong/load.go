package pong

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadSettings(path string) (Settings, error) {
	settings := Settings{}

	f, err := os.Open(path)
	if err != nil {
		return settings, fmt.Errorf("cannot open file: %s", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return settings, fmt.Errorf("cannot read file: %s", err)
	}

	if err := json.Unmarshal(b, &settings); err != nil {
		return settings, fmt.Errorf("cannot unmarshal: %s", err)
	}

	// for _, paths := range settings.Paths {
	// 	if err := checkPaths(paths); err != nil {
	// 		return settings, err
	// 	}
	// }

	for k, v := range settings.Commands {
		if v.OnErr != "exit" && v.OnErr != "continue" {
			return settings, fmt.Errorf("%s \"onerror\" value must be \"exit\" or \"continue\"", k)
		}
	}

	return settings, nil
}

func checkPaths(paths ...[]string) error {
	rootFolder, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current working directory: %s\n", err)
	}
	for _, path := range paths {
		for _, projectPath := range path {
			projectAbsPath := rootFolder + "/" + projectPath
			if fi, err := os.Stat(projectAbsPath); os.IsNotExist(err) || !fi.IsDir() {
				if err == nil {
					return fmt.Errorf("%s is not a directory", fi.Name())
				}
				return fmt.Errorf("cannot find directory: %s", err)
			}
		}
	}
	return nil
}
