package pkg

import (
	"ConcatFiles/loger"
	"GoMap/rgb"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//
//
// FUNCTIONS

func CreateNewFolder(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		loger.Error("Erreur durant la création du dossier", err)
	}
}

func RemoveAllDlgFiles(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		loger.Crash("Erreur durant l'ouverture du dossier", err)
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		loger.Crash("Erreur durant la lecture des fichiers", err)
	}

	for _, name := range names {
		if strings.Contains(strings.ToLower(name), "liste") || strings.Contains(strings.ToLower(name), "pbo") {
			return
		}

		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			loger.Crash("Erreur durant la suppression des fichiers", err)
		}
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
