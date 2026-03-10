# dotfilesGo

> Instala snippets de Go en **Zed** y **VSCode** con un solo comando.

```bash
go install github.com/geomark27/dotfilesGo/cmd/dotfilesgo@latest && dotfilesgo
```

---

## Instalación

```bash
# 1. Instalar
go install github.com/geomark27/dotfilesGo/cmd/dotfilesgo@latest

# 2. Ejecutar
dotfilesgo
```

```
✓ Zed    → ~/.config/zed/snippets/go.json
✓ VSCode → ~/.config/Code/User/snippets/go.json
2 editor(es) configurados.
```

> En **Windows** y **WSL** detecta automáticamente las rutas de `APPDATA`.

---

## Comandos

| Comando | Descripción |
|---------|-------------|
| `dotfilesgo` | Instala los snippets en todos los editores detectados |
| `dotfilesgo --version` | Muestra la versión instalada |
| `dotfilesgo --update` | Actualiza a la última versión disponible |

Al instalar, el binario consulta automáticamente si hay una versión nueva y notifica al usuario:

```
⚡ Nueva versión disponible: v0.2.0 → v0.3.0
   Actualiza con: dotfilesgo --update
```

---

## Snippets incluidos

### 🔁 Control de flujo

| Prefix | Expansión |
|--------|-----------|
| `forr` | `for k, v := range collection { }` |
| `fori` | `for i := 0; i < n; i++ { }` |
| `sw` | `switch variable { case ...: default: }` |
| `sel` | `select { case v := <-ch: }` |

### 🧱 Tipos y estructuras

| Prefix | Expansión |
|--------|-----------|
| `st` | `type Name struct { }` |
| `stc` | `type Name struct { }` + `func NewName(...) *Name` |
| `iface` | `type Name interface { }` |

### ⚙️ Funciones

| Prefix | Expansión |
|--------|-----------|
| `fn` | `func Name(args) ReturnType { }` |
| `meth` | `func (r *Receiver) Name(args) ReturnType { }` |
| `main` | `package main` + `func main()` |
| `init` | `func init() { }` |

### 🚨 Manejo de errores

| Prefix | Expansión |
|--------|-----------|
| `ife` | `if err != nil { return nil, err }` |
| `errw` | `if err != nil { return nil, fmt.Errorf("context: %w", err) }` |
| `errt` | `errors.New("message")` |

### ⚡ Concurrencia

| Prefix | Expansión |
|--------|-----------|
| `gor` | `go func() { }()` |
| `ch` | `make(chan Type)` |
| `mu` | `sync.Mutex` + `Lock/Unlock` |
| `wg` | `sync.WaitGroup` completo |

### 🧪 Testing

| Prefix | Expansión |
|--------|-----------|
| `test` | `func TestName(t *testing.T) { }` |
| `bench` | `func BenchmarkName(b *testing.B) { }` |
| `tcase` | tabla de test cases completa |

### 🌐 HTTP

| Prefix | Expansión |
|--------|-----------|
| `hh` | `func NameHandler(w http.ResponseWriter, r *http.Request)` |
| `mid` | middleware con `next http.Handler` |

### 🛠️ Utilidades

| Prefix | Expansión |
|--------|-----------|
| `log` | `log.Printf("message: %v\n", value)` |
| `def` | `defer func()` |
| `msl` | `make([]Type, 0, cap)` |
| `mmap` | `make(map[KeyType]ValueType)` |

---

## Compatibilidad

| OS | Zed | VSCode |
|----|-----|--------|
| Linux | ✅ | ✅ |
| macOS | ✅ | ✅ |
| Windows | ✅ | ✅ |

---

## Cómo funciona

El binario usa `//go:embed` para empaquetar el JSON dentro del ejecutable — no necesita archivos externos ni conexión a internet después de instalarse. Detecta automáticamente el OS (Linux, macOS, Windows) y también si corre dentro de WSL.

---

## Autor

**Marcos Ramos** · [github.com/geomark27](https://github.com/geomark27)
