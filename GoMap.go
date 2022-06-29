//go:generate goversioninfo -icon=GoMap.ico -manifest=GoMap.exe.manifest
package main

import (
	"GoMap/loger"
	"GoMap/pkg"
	"GoMap/rgb"
	"bufio"
	"flag"
	"os"
)

func main() {

	flagMode := flag.String("m", "fe", "Creation des dossiers pour un exports")
	flag.Parse()

	switch *flagMode {
	case "fe":
		pkg.FoldersExport()
	}

	pkg.DrawSep(" FIN ")
	pkg.DrawEnd()

	rgb.GreenB.Print("Appuyer sur Entrée pour quitter...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		loger.Crash("Crash :", err)
	}
}
