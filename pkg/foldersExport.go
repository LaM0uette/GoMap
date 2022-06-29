package pkg

import "fmt"

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	nro := GetUserInput("Entrez le NRO : ")
	fmt.Println(nro)

}
