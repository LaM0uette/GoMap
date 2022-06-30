package pkg

import "path"

func CleanDlgFiles() {

	DrawStart("Clean DLG Files")
	DrawSep("RUN")

	dir := GetCurrentDir()

	RemoveAllFiles(path.Join(dir, ""))
}
