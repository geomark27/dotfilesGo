package platform

import (
	"os"
	"path/filepath"
	"strings"
)

// IsWSL detecta si el binario corre dentro de Windows Subsystem for Linux.
func IsWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(data))
	return strings.Contains(lower, "microsoft") || strings.Contains(lower, "wsl")
}

// WSLWindowsAppData retorna la ruta equivalente a %APPDATA% accesible desde WSL.
func WSLWindowsAppData() string {
	if appdata := os.Getenv("APPDATA"); appdata != "" {
		return appdata
	}

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
