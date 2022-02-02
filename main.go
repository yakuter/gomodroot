package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const (
	pathArg       = "path"
	goModFileName = "go.mod"
)

var path string

func main() {
	flag.StringVar(&path, pathArg, "", "path to look for go mod root dir")
	flag.Parse()

	if path == "" {
		log.Println("path argument is empty")
		flag.PrintDefaults()
		os.Exit(1)
	}

	rootPath, err := goRootPath(path)
	if err != nil {
		log.Fatalf("Failed to locate go mod root path for path %s error: %v", path, err)
	}

	log.Printf("Go mod root path for %s is: %s", path, rootPath)

}

// goRootPath returns root dir of the project where go.mod file is exist
func goRootPath(path string) (string, error) {
	path = filepath.Clean(path)

	for {
		fi, err := os.Stat(filepath.Join(path, goModFileName))
		if err == nil && !fi.IsDir() {
			return path, nil
		}

		d := filepath.Dir(path)
		if d == path {
			break
		}

		path = d
	}

	return "", nil
}
