package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/geomark27/dotfilesGo/internal/platform"
)

// Install escribe los snippets en cada editor y retorna cuántos se instalaron.
func Install(data []byte) int {
	paths := snippetPaths()
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

	return installed
}

// snippetPaths retorna un mapa de editor → ruta de destino según el OS.
func snippetPaths() map[string]string {
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

		if platform.IsWSL() {
			if appdata := platform.WSLWindowsAppData(); appdata != "" {
				paths["VSCode (Windows)"] = filepath.Join(appdata, "Code", "User", "snippets", "go.json")
				paths["Zed (Windows)"] = filepath.Join(appdata, "Zed", "snippets", "go.json")
			}
		}
	}

	return paths
}
