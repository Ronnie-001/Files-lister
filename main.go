package main

import (
	"fmt"
	"os"
	"log"
	"strings"
)

var (
	dirToFind string
	viewHiddenFiles string
)

func main() {
	fmt.Println("Which dir do you want to list?")
	fmt.Scan(&dirToFind)
	
	showFiles := false

	fmt.Println("Would you also like to view hidden files?")	
	fmt.Scan(&viewHiddenFiles)		
	
	if viewHiddenFiles == "yes" {
		showFiles = !showFiles
	}

	files, err := os.ReadDir(dirToFind)
	_ = files
	if err != nil {
		log.Fatal(err)
	}
	
	for _, val := range files {
		if strings.HasPrefix(val.Name(), ".") && !showFiles {
			continue	
		} else {
			fmt.Println(val.Name())
		}
	}
}
