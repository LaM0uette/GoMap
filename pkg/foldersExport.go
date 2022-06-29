package pkg

import (
	"GoMap/coSql"
	"fmt"
	"path"
	"strconv"
	"strings"
)

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	refcode3 := strings.ToUpper(GetUserInput("Entrez le REFCODE3 : "))
	phase := strings.ToUpper(GetUserInput("Entrez la PHASE : "))

	if strings.Contains(phase, "REC") {
		strings.Replace(phase, "REC", "DOE", 1)
	}

	_nLivaison := GetUserInput("Entrez le n° de LIVRAISON : ")
	nLivaison, _ := strconv.Atoi(_nLivaison)
	nVersion := GetUserInput("Entrez le n° de VERSION : ")

	rc := coSql.GetRefcodeData(refcode3)

	folderDLG := path.Join(
		FolderExportsGrace,
		fmt.Sprintf("RIP%v", rc.RefCode1),
		fmt.Sprintf("NRO%v_PM%v_%s", rc.NRO, rc.SRO, refcode3),
		fmt.Sprintf("%s_%v", phase, nLivaison),
		fmt.Sprintf("V%v", nVersion))

	fmt.Println(folderDLG)

	CreateNewFolder(folderDLG)
	CreateNewFolder(path.Join(folderDLG, "__Historiques__"))
	CreateNewFolder(path.Join(folderDLG, "DJANGO"))
	CreateNewFolder(path.Join(folderDLG, "DLG"))
	CreateNewFolder(path.Join(folderDLG, "DLG", fmt.Sprintf("%s-DLG-%v-%s-%s-%s-V%v", phase[0:3], rc.RefCode1, rc.RefCode2, refcode3, fmt.Sprintf("%02d", nLivaison), nVersion)))
	CreateNewFolder(path.Join(folderDLG, "PDB"))
	CreateNewFolder(path.Join(folderDLG, "ROPT"))
}
