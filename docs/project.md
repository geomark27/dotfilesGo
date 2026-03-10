# dotfilesGo — Proyecto

## ¿Qué es?

CLI escrito en Go que instala snippets de código en los editores del usuario con un solo comando.
Los snippets están embebidos dentro del binario usando `//go:embed`, por lo que no requiere archivos externos ni conexión a internet después de instalarse.

```bash
go install github.com/geomark27/dotfilesGo/cmd/dotfilesgo@latest
dotfilesgo
```

---

## ¿Qué hace actualmente?

- [x] Instala snippets de Go en Zed y VSCode
- [x] Detecta el sistema operativo automáticamente (Linux, macOS, Windows)
- [x] Detecta si corre dentro de WSL e instala en rutas de Windows también
- [x] Detecta dinámicamente el usuario de Windows desde WSL (sin hardcodear rutas)
- [x] Crea los directorios necesarios si no existen
- [x] Notifica si hay una versión nueva disponible al instalar
- [x] Soporta `--version`, `--update`

---

## Roadmap

### v0.2 — Auto-update ✅
- [x] Agregar variable de versión en el binario (inyectable via `-ldflags "-X main.version=v0.2.0"`)
- [x] Al ejecutar `dotfilesGo`, consultar la última versión disponible en GitHub
- [x] Si hay una versión más nueva, notificar al usuario
- [x] Opción `dotfilesGo --update` que corre `go install` automáticamente
- [x] Flag `--version` / `-v` para ver la versión instalada

### v0.3 — Más editores
- [ ] Soporte para Neovim (LuaSnip)
- [ ] Soporte para IntelliJ / GoLand

### v0.4 — CLI mejorado
- [ ] Flag `--dry-run` para ver qué se instalaría sin instalar nada
- [ ] Flag `--only zed` o `--only vscode` para instalar en un editor específico
- [ ] Flag `--version` para ver la versión instalada *(movido a v0.2)*
- [ ] Output con colores

### v1.0 — Publicación
- [ ] Publicar en el Marketplace de VSCode como extensión
- [ ] Agregar GitHub Actions para releases automáticas con tags
- [ ] Documentación completa en README

---

## Estructura del proyecto

```
dotfilesGo/
├── cmd/
│   └── dotfilesgo/
│       └── main.go          ← entry point + flags
├── internal/
│   ├── assets/
│   │   ├── assets.go        ← embed de go.json
│   │   └── go.json          ← snippets
│   ├── installer/
│   │   └── installer.go     ← Install() + snippetPaths()
│   ├── platform/
│   │   └── platform.go      ← IsWSL() + WSLWindowsAppData()
│   └── updater/
│       └── updater.go       ← Run() + LatestVersion()
├── docs/
│   └── project.md
├── go.mod
└── README.md
```

---

## Cómo contribuir

1. Fork del repo
2. Agrega tus snippets en `internal/assets/go.json`
3. `go run ./cmd/dotfilesgo` para probar
4. Pull request con descripción del cambio