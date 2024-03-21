package utils

import "fmt"

func Error(message string) {
	fmt.Println(Red + message + Reset)
}

func Info(message string) {
	fmt.Println(Blue + message + Reset)
}

func Success(message string) {
	fmt.Println(Green + message + Reset)
}

func Warning(message string) {
	fmt.Println(Yellow + message + Reset)
}
