package pkg

import (
	"GoMap/coSql"
	"fmt"
	"strings"
)

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	//nro := GetUserInput("Entrez le NRO : ")
	//sro := GetUserInput("Entrez le SRO : ")
	refcode3 := strings.ToUpper(GetUserInput("Entrez le REFCODE3 : "))
	phase := strings.ToUpper(GetUserInput("Entrez la PHASE : "))
	nLivaison := GetUserInput("Entrez le n° de LIVRAISON : ")
	nVersion := GetUserInput("Entrez le n° de VERSION : ")

	fmt.Println(refcode3, phase, nLivaison, nVersion)

	coSql.GetRefcodeData("BIMA")

	//folderDlg := fmt.Sprintf("")
	//CreateNewFolder(path.Join(FolderExportsGrace, folderDlg))

}
