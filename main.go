package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

	if runtime.GOOS == "windows" {
		appdata := os.Getenv("APPDATA")
		return map[string]string{
			"Zed":    filepath.Join(appdata, "Zed", "snippets", "go.json"),
			"VSCode": filepath.Join(appdata, "Code", "User", "snippets", "go.json"),
		}
	}

	// Linux / macOS
	return map[string]string{
		"Zed":    filepath.Join(home, ".config", "zed", "snippets", "go.json"),
		"VSCode": filepath.Join(home, ".config", "Code", "User", "snippets", "go.json"),
	}
}