package pkg

import (
	"path"
	"strings"
)

func CleanDlgFiles() {

	DrawStart("Clean DLG Files")
	DrawSep("RUN")

	p := "C:\\Users\\XD5965\\OneDrive - EQUANS\\Bureau\\V12"

	dlg := getDLGName(p)

	for folder := range DLGFolders {
		if strings.Contains(DLGFolders[folder], "_") {
			continue
		}

		//RemoveAllDlgFiles(path.Join(GetCurrentDir(), DLGFolders[folder]))
		RemoveAllDlgFiles(path.Join(p, DLGFolders[folder]))
	}

	if len(dlg) > 0 {
		CreateNewFolder(path.Join(p, "DLG", dlg))
	}
}
