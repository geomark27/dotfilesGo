package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/geomark27/dotfilesGo/internal/assets"
	"github.com/geomark27/dotfilesGo/internal/installer"
	"github.com/geomark27/dotfilesGo/internal/updater"
)

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v
		}
	}
	return "dev"
}

func main() {
	version := getVersion()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--update":
			updater.Run()
			return
		case "--version", "-v":
			fmt.Printf("dotfilesGo %s\n", version)
			return
		}
	}

	// Consultar la última versión en paralelo, sin bloquear la instalación
	updateCh := make(chan string, 1)
	go func() {
		latest, err := updater.LatestVersion()
		if err != nil || latest == version || version == "dev" {
			updateCh <- ""
			return
		}
		updateCh <- latest
	}()

	data, err := assets.FS.ReadFile("go.json")
	if err != nil {
		fmt.Println("error leyendo snippets:", err)
		os.Exit(1)
	}

	installed := installer.Install(data)
	fmt.Printf("\n%d editor(es) configurados.\n", installed)

	if latest := <-updateCh; latest != "" {
		fmt.Printf("\n⚡ Nueva versión disponible: %s → %s\n", version, latest)
		fmt.Printf("   Actualiza con: dotfilesGo --update\n")
	}
}
