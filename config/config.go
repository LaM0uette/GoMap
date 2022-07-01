package config

import (
	"log"
	"os/user"
	"path/filepath"
)

const (
	Name    = "GoMap"
	Author  = "LaM0uette"
	Version = "0.2"
)

var (
	DstPath   = filepath.Join(GetTempDir(), Name+"_Temp")
	LogsPath  = filepath.Join(DstPath, "logs")
	DumpsPath = filepath.Join(DstPath, "dumps")
)

func GetTempDir() string {
	temp, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(temp.HomeDir)
}
