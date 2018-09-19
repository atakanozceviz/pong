package pong

import (
	"fmt"
	"os"
	"path"
	"sync"
)

func Run(settingsPath, command, paths string) error {
	settings, err := loadSettings(settingsPath)
	if err != nil {
		// fmt.Printf("Cannot load settings: %s\n", err)
		// os.Exit(1)
		return err
	}

	if _, ok := settings.Commands[command]; !ok {
		// fmt.Printf("Provided settings doesn't contain \"command\" named: %s\n", command)
		// os.Exit(1)
		return err
	}

	if _, ok := settings.Paths[paths]; !ok {
		// fmt.Printf("Provided settings doesn't contain \"paths\" named: %s\n", paths)
		// os.Exit(1)
		return err
	}

	wg := &sync.WaitGroup{}
	jobs := make(chan *Job)

	for i := 0; i < settings.Commands[command].Workers; i++ {
		wg.Add(1)
		go worker(jobs, wg)
	}

	rootPath, err := os.Getwd()
	if err != nil {
		// fmt.Printf("cannot get working directory: %s\n", err)
		// os.Exit(1)
		return err
	}

	for _, dir := range settings.Paths[paths] {
		jobs <- &Job{
			Path:    path.Join(rootPath, dir),
			OnErr:   settings.Commands[command].OnErr,
			Command: settings.Commands[command],
		}
	}
	close(jobs)

	wg.Wait()
	return nil
}

func ParseSettings(settingsPath string) (commands map[string]string, paths []string) {
	settings, err := loadSettings(settingsPath)
	if err != nil {
		fmt.Printf("Cannot load settings: %s\n", err)
		os.Exit(1)
	}

	commands = make(map[string]string)
	paths = make([]string, 0, 2)

	for command, cmd := range settings.Commands {
		commands[command] = cmd.Description
	}

	for path := range settings.Paths {
		paths = append(paths, path)
	}
	return
}
