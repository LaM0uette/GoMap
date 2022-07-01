package pkg

import (
	"GoMap/loger"
	"fmt"
	"path"
	"strings"
)

func CleanDlgFiles() {

	DrawStart("Clean DLG Files")
	DrawSep("RUN")

	DrawParam("CREATION DES VARIABLES...")

	dlgPath := GetCurrentDir()
	dlg := getDLGName(dlgPath)

	for folder := range DLGFolders {
		if strings.Contains(DLGFolders[folder], "_") {
			continue
		}

		//RemoveAllDlgFiles(path.Join(GetCurrentDir(), DLGFolders[folder]))
		RemoveAllDlgFiles(path.Join(dlgPath, DLGFolders[folder]))
		loger.Ok(fmt.Sprintf("Dosier %s vidé !", DLGFolders[folder]))
	}

	if len(dlg) > 0 {
		CreateNewFolder(path.Join(dlgPath, "DLG", dlg))
	}

	DrawSep("RESULTAT")
	loger.Ok("Nettoyage terminé !")
}
