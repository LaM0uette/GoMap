package pkg

import (
	"ConcatFiles/loger"
	"GoMap/rgb"
	"bufio"
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

	input, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		loger.Crash("Crash lors de la recuperation de la saisie utilisateur :", err)
		return nil
	}

	return input
}
