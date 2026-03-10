package updater

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	Module     = "github.com/geomark27/dotfilesGo"
	GithubRepo = "geomark27/dotfilesGo"
)

// Run ejecuta go install para actualizar el binario a la última versión.
func Run() {
	fmt.Println("Actualizando dotfilesGo a la última versión...")
	cmd := exec.Command("go", "install", Module+"@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("✗ Error al actualizar: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ dotfilesGo actualizado correctamente.")
}

// LatestVersion consulta la GitHub releases API y retorna el tag de la última release.
func LatestVersion() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/" + GithubRepo + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}
