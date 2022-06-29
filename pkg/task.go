package pkg

import (
	"ConcatFiles/loger"
	"GoMap/rgb"
	"fmt"
	"os"
)

//
//
// FUNCTIONS

func CreateNewFolder(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		loger.Error("Erreur durant la création du dossier", err)
	}
}

//
//
// GETTERS

func GetCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		loger.Error("Error with current dir:", err)
		os.Exit(1)
	}
	return pwd
}

func GetUserInput(msg string) string {

	rgb.GreenB.Print(msg)

	input := ""
	_, err := fmt.Scanln(&input)
	if err != nil {
		loger.Crash("Crash lors de la recuperation de la saisie utilisateur : ", err)
		return ""
	}

	return input
}
