package main

import (
	"flag"
	"fmt"
	"image/gif"
	"image/png"
	"os"
	"path/filepath"
)

var input string

func init() {
	flag.StringVar(&input, "f", "", "Input gif file")
	flag.Parse()

	if input == "" {
		fmt.Println("Please specify the gif file")
		os.Exit(1)
	}
}

func getDirAndName(path string) (dir string, name string) {
	path, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Open gif file fail: %v", err)
		return
	}

	dir, name = filepath.Split(path)

	extension := filepath.Ext(name)
	name = name[0 : len(name)-len(extension)]

	return dir, name
}

func getOutputFileName(dir string, name string, index int) string {
	return filepath.Join(dir, fmt.Sprintf("%s_%d.png", name, index))
}

func main() {
	f, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Open gif file fail")
		return
	}
	defer f.Close()

	//Decode gif file
	g, err := gif.DecodeAll(f)
	if err != nil {
		fmt.Printf("Open gif file: %v", err)
		return
	}

	dir, name := getDirAndName(input)

	for index, val := range g.Image {
		out, err := os.OpenFile(getOutputFileName(dir, name, index), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Create png file: %v", err)
			return
		}

		err = png.Encode(out, val)
		if err != nil {
			fmt.Printf("Write png file: %v", err)
		}

		out.Close()
	}
}
