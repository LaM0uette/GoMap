//go:generate goversioninfo -icon=GoMap.ico -manifest=GoMap.exe.manifest
package main

import (
	"flag"
	"fmt"
)

func main() {

	flagMode := flag.String("m", "fe", "Création des dossiers pour un exports")
	flag.Parse()

	fmt.Println("test")

}
