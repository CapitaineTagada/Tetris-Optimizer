package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"tetris-optimizer/src/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go path/to/<input file>")
		return
	}

	filePath := os.Args[1]

	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Convert the content into a [][]string format, if not the !IsValid(Tetromino) doesn't work
	lines := strings.Split(string(content), "\n")
	var Tetromino [][]string
	for _, line := range lines {
		if line != "" {
			Tetromino = append(Tetromino, strings.Split(line, ""))
		}
	}

	if !utils.IsValid(Tetromino) {
		fmt.Println("ERROR")
		return
	}

	// Afficher le contenu du fichier pour vérifier que la lecture a été réussie
	fmt.Println(string(content))

}
