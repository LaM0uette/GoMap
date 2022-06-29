package pkg

import (
	"GoMap/coSql"
	"fmt"
	"path"
	"strconv"
	"strings"
)

type dlgData struct {
	NRO      int
	SRO      int
	Refcode1 int
	Refcode2 string
	Refcode3 string
	Phase    string
	Livaison int
	Version  int
}

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	dlg := getDlgData()

	folderDLG := path.Join(
		FolderExportsGrace,
		fmt.Sprintf("RIP%v", dlg.Refcode1),
		fmt.Sprintf("NRO%v_PM%v_%s", dlg.NRO, dlg.SRO, dlg.Refcode3),
		fmt.Sprintf("%s_%v", dlg.Phase, dlg.Livaison),
		fmt.Sprintf("V%v", dlg.Version))

	CreateNewFolder(folderDLG)
	CreateNewFolder(path.Join(folderDLG, "__Historiques__"))
	CreateNewFolder(path.Join(folderDLG, "DJANGO"))
	CreateNewFolder(path.Join(folderDLG, "DLG"))
	CreateNewFolder(path.Join(folderDLG, "DLG", fmt.Sprintf("%s-DLG-%v-%s-%s-%s-V%v", dlg.Phase[0:3], dlg.Refcode1, dlg.Refcode2, dlg.Refcode3, fmt.Sprintf("%02d", dlg.Livaison), dlg.Version)))
	CreateNewFolder(path.Join(folderDLG, "PDB"))
	CreateNewFolder(path.Join(folderDLG, "ROPT"))
}

func getDlgData() dlgData {
	refcode3 := strings.ToUpper(GetUserInput("Entrez le REFCODE3 : "))
	phase := strings.ToUpper(GetUserInput("Entrez la PHASE : "))
	livaison, _ := strconv.Atoi(GetUserInput("Entrez le n° de LIVRAISON : "))
	_version, _ := strconv.Atoi(GetUserInput("Entrez le n° de VERSION : "))

	rc := coSql.GetRefcodeData(refcode3)

	return dlgData{
		NRO:      rc.NRO,
		SRO:      rc.SRO,
		Refcode1: rc.RefCode1,
		Refcode2: rc.RefCode2,
		Refcode3: refcode3,
		Phase:    phase,
		Livaison: livaison,
		Version:  _version,
	}
}
