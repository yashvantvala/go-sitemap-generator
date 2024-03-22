package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder(dirName string) {
	currentDir, err := os.Getwd()
	folderPath := filepath.Join(currentDir, dirName)
	err = os.Mkdir(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	fmt.Println("Folder created successfully:", dirName)
}
