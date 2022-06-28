//go:generate goversioninfo -icon=GoMap.ico -manifest=GoMap.exe.manifest
package main

import (
	"GoMap/pkg"
	"flag"
)

func main() {

	flagMode := flag.String("m", "fe", "Création des dossiers pour un exports")
	flag.Parse()

	txtMode := ""
	switch *flagMode {
	case "fe":
		txtMode = "Folders Export"
	}

	pkg.DrawStart(txtMode)
	pkg.DrawSep("BUILD")

}
