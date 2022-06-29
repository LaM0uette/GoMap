package pkg

import (
	"ConcatFiles/loger"
	"GoMap/rgb"
	"fmt"
	"os"
)

func GetCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		loger.Error("Error with current dir:", err)
		os.Exit(1)
	}
	return pwd
}

func GetUserInput(msg string) any {

	rgb.GreenB.Print(msg)

	var input string
	_, err := fmt.Scanf("%v", &input)
	if err != nil {
		loger.Crash("Crash lors de la recuperation de la saisie utilisateur : ", err)
		return nil
	}

	return input
}
