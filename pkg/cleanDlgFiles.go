package pkg

import (
	"fmt"
)

func CleanDlgFiles() {

	DrawStart("Clean DLG Files")
	DrawSep("RUN")

	files := []string{"pkg"}
	output := "done.zip"

	if err := ZipFiles(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File:", output)
}
