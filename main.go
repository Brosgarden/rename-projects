package main

import (
	"strings"
	"io/ioutil"

	"log"
	"os"
	"path/filepath"
	"fmt"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	renameMap := make(map[string]string)

	fmt.Printf("Finding child directories for %v\n", workDir)
	files, err := ioutil.ReadDir(workDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			newName := strings.ToLower(strings.Replace(f.Name(), ".", "-", -1))
			fmt.Printf("Renaming %v to %v\n", f.Name(), newName)
			os.Rename(f.Name(), newName)
			renameMap["':"+newName+"'"] = "':" + f.Name() + "'"
		}
	}

	fmt.Println("Finding Gradle Files")
	gradleFiles := make([]string, 0, 10)
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == "build.gradle" {
			fmt.Println(path)
			gradleFiles = append(gradleFiles, path)
		}
		return nil
	})

	fmt.Println("Replacing Gradle File Contents")
	for _, gradleFile := range gradleFiles {
		read, err := ioutil.ReadFile(gradleFile)
		if err != nil {
			log.Printf("Unable to handle %v\n", gradleFile)
		}
		newContent := string(read)
		for new, old := range renameMap {
			fmt.Printf("Replacing %v with %v in file: %v\n", old, new, gradleFile)
			newContent = strings.Replace(newContent, old, new, -1)
		}
		err = ioutil.WriteFile(gradleFile, []byte(newContent), 0)
		if err != nil {
			log.Printf("Unable to handle %v\n", gradleFile)
		}
	}
}
