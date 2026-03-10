package main

import (
	"fmt"
	"os"

	"github.com/geomark27/dotfilesGo/internal/assets"
	"github.com/geomark27/dotfilesGo/internal/installer"
	"github.com/geomark27/dotfilesGo/internal/updater"
)

// version se inyecta en build time: -ldflags "-X main.version=v0.2.0"
var version = "dev"

func main() {
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
