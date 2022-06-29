package pkg

import (
	"GoMap/coSql"
	"fmt"
	"path"
	"strings"
)

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	refcode3 := strings.ToUpper(GetUserInput("Entrez le REFCODE3 : "))
	phase := strings.ToUpper(GetUserInput("Entrez la PHASE : "))
	nLivaison := GetUserInput("Entrez le n° de LIVRAISON : ")
	nVersion := GetUserInput("Entrez le n° de VERSION : ")
	rc := coSql.GetRefcodeData("BIMA")

	folderDLG := path.Join(
		FolderExportsGrace,
		fmt.Sprintf("RIP%v", rc.RefCode1),
		fmt.Sprintf("NRO%v_PM%v_%s", rc.NRO, rc.SRO, refcode3),
		fmt.Sprintf("%s_%v", phase, nLivaison),
		fmt.Sprintf("V%v", nVersion))

	CreateNewFolder(folderDLG)
	CreateNewFolder(path.Join(folderDLG, "__Historiques__"))
	CreateNewFolder(path.Join(folderDLG, "DJANGO"))
	CreateNewFolder(path.Join(folderDLG, "DLG"))
	CreateNewFolder(path.Join(folderDLG, "DLG", ""))
	CreateNewFolder(path.Join(folderDLG, "PDB"))
	CreateNewFolder(path.Join(folderDLG, "ROPT"))
}
