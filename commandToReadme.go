package main

// # commentToReadme
// ---
// This is a program will create a **Readme Markdown** file from the tagged comments and code in that source file.
//
// by Dave Larsen KV0S
//
// code at github.com/kv0s/commentToReadme
//
// Licensed under the GPL3
//
// ---

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const progname string = "commentToReadme"
const version string = "0.0.1"

func main() {
	// fmt.Printf("%s Version %s\n", progname, version)

	f, err := os.OpenFile("./commandToReadme.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		if strings.Contains(line, "//") {
			if strings.HasPrefix(line, "//") {
				fmt.Printf("%s", strings.TrimPrefix(line, "//"))
			}
		}
	}

}
