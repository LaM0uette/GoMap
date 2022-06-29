package pkg

import "fmt"

func FoldersExport() {

	DrawStart("Folders Export")
	DrawSep("INITIALISATION")

	nro := GetUserInput("Entrez le NRO : ")
	sro := GetUserInput("Entrez le SRO : ")
	refcode3 := GetUserInput("Entrez le REFCODE3 : ")
	phase := GetUserInput("Entrez la PHASE : ")
	nLivaison := GetUserInput("Entrez le n° de LIVRAISON : ")
	nVersion := GetUserInput("Entrez le n° de VERSION : ")

	fmt.Println(nro, sro, refcode3, phase, nLivaison, nVersion)

}
