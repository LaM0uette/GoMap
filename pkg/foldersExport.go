package pkg

import (
	"GoMap/coSql"
	"GoMap/loger"
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

	DrawSep("CREATION DES DOSSIERS")

	createDlgFolders(dlg)
	loger.Ok("Dossiers créés !")
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

func createDlgFolders(dlg dlgData) {
	fRIP := fmt.Sprintf("RIP%v", dlg.Refcode1)
	fZone := fmt.Sprintf("NRO%v_PM%v_%s", dlg.NRO, dlg.SRO, dlg.Refcode3)

	var fLivraison string
	if dlg.Phase == "EXE" {
		fLivraison = dlg.Phase
	} else {
		fLivraison = fmt.Sprintf("%s_%v", dlg.Phase, dlg.Livaison)
	}

	fVersion := fmt.Sprintf("V%v", dlg.Version)

	p := path.Join(FolderExportsGrace, fRIP, fZone, fLivraison, fVersion)
	folders := []string{"__Historiques__", "DJANGO", "DLG", "PDB", "ROPT"}

	for folder := range folders {
		CreateNewFolder(path.Join(p, folders[folder]))
	}

	CreateNewFolder(path.Join(p,
		"DLG",
		fmt.Sprintf("%s-DLG-%v-%s-%s-%s-V%v",
			dlg.Phase[0:3],
			dlg.Refcode1,
			dlg.Refcode2,
			dlg.Refcode3,
			fmt.Sprintf("%02d", dlg.Livaison),
			dlg.Version)))
}
