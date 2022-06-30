package pkg

import (
	"path"
	"strings"
)

func CleanDlgFiles() {

	DrawStart("Clean DLG Files")
	DrawSep("RUN")

	for folder := range DLGFolders {
		if strings.Contains(DLGFolders[folder], "_") {
			continue
		}

		//RemoveAllFiles(path.Join(GetCurrentDir(), DLGFolders[folder]))
		p := "C:\\Users\\XD5965\\OneDrive - EQUANS\\Bureau\\V12"
		RemoveAllFiles(path.Join(p, DLGFolders[folder]))
	}
}
