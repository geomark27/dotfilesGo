package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//go:embed snippets/go.json
var snippets embed.FS

func main() {
	data, err := snippets.ReadFile("snippets/go.json")
	if err != nil {
		fmt.Println("error leyendo snippets:", err)
		os.Exit(1)
	}

	paths := getSnippetPaths()
	installed := 0

	for editor, path := range paths {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("✗ %s: no se pudo crear directorio: %v\n", editor, err)
			continue
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			fmt.Printf("✗ %s: no se pudo escribir archivo: %v\n", editor, err)
			continue
		}
		fmt.Printf("✓ %s → %s\n", editor, path)
		installed++
	}

	fmt.Printf("\n%d editor(es) configurados.\n", installed)
}

func getSnippetPaths() map[string]string {
	home, _ := os.UserHomeDir()
	paths := map[string]string{}

	switch runtime.GOOS {
	case "windows":
		appdata := os.Getenv("APPDATA")
		paths["Zed"] = filepath.Join(appdata, "Zed", "snippets", "go.json")
		paths["VSCode"] = filepath.Join(appdata, "Code", "User", "snippets", "go.json")

	case "darwin":
		paths["Zed"] = filepath.Join(home, ".config", "zed", "snippets", "go.json")
		paths["VSCode"] = filepath.Join(home, "Library", "Application Support", "Code", "User", "snippets", "go.json")

	default: // linux
		paths["Zed"] = filepath.Join(home, ".config", "zed", "snippets", "go.json")
		paths["VSCode"] = filepath.Join(home, ".config", "Code", "User", "snippets", "go.json")

		// En WSL, tambien instalar en las rutas de Windows
		if isWSL() {
			if appdata := wslWindowsAppData(); appdata != "" {
				paths["VSCode (Windows)"] = filepath.Join(appdata, "Code", "User", "snippets", "go.json")
				paths["Zed (Windows)"] = filepath.Join(appdata, "Zed", "snippets", "go.json")
			}
		}
	}

	return paths
}

// isWSL detecta si el binario corre dentro de Windows Subsystem for Linux.
func isWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(data))
	return strings.Contains(lower, "microsoft") || strings.Contains(lower, "wsl")
}

// wslWindowsAppData retorna la ruta equivalente a %APPDATA% accesible desde WSL.
func wslWindowsAppData() string {
	// Intentar via variable de entorno expuesta por WSL interop
	if appdata := os.Getenv("APPDATA"); appdata != "" {
		return appdata
	}

	// Fallback: buscar en /mnt/c/Users/<usuario>/AppData/Roaming
	entries, err := os.ReadDir("/mnt/c/Users")
	if err != nil {
		return ""
	}
	skip := map[string]bool{
		"Public": true, "Default": true, "Default User": true,
		"All Users": true, "desktop.ini": true,
	}
	for _, e := range entries {
		if !e.IsDir() || skip[e.Name()] {
			continue
		}
		appdata := filepath.Join("/mnt/c/Users", e.Name(), "AppData", "Roaming")
		if _, err := os.Stat(appdata); err == nil {
			return appdata
		}
	}
	return ""
}