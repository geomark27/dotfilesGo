BINARY  := dotfilesgo
MODULE  := github.com/geomark27/dotfilesGo/cmd/dotfilesgo
BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)

# Versioning
LAST_TAG    := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
MAJOR       := $(shell echo $(LAST_TAG) | cut -d. -f1 | tr -d v)
MINOR       := $(shell echo $(LAST_TAG) | cut -d. -f2)
PATCH       := $(shell echo $(LAST_TAG) | cut -d. -f3)
NEXT_PATCH  := v$(MAJOR).$(MINOR).$(shell echo $$(($(PATCH)+1)))
NEXT_MINOR  := v$(MAJOR).$(shell echo $$(($(MINOR)+1))).0
NEXT_MAJOR  := v$(shell echo $$(($(MAJOR)+1))).0.0

.DEFAULT_GOAL := help

# ============================================
# DESARROLLO
# ============================================

.PHONY: run
run: ## Ejecuta el CLI sin compilar binario
	go run ./cmd/dotfilesgo

.PHONY: build
build: ## Compila el binario inyectando la versión del último tag
	go build -ldflags "-X main.version=$(LAST_TAG)" -o $(BINARY) ./cmd/dotfilesgo
	@echo "✅ Binario: ./$(BINARY) ($(LAST_TAG))"

.PHONY: test
test: ## Corre todos los tests
	go test ./...

.PHONY: vet
vet: ## Analiza el código con go vet
	go vet ./...

.PHONY: check
check: vet test ## Corre vet + tests (requerido antes de release)
	@echo "✅ Check completado."

.PHONY: clean
clean: ## Elimina el binario compilado
	rm -f $(BINARY)
	@echo "🗑️  Binario eliminado."

# ============================================
# COMANDOS GIT
# ============================================

.PHONY: push
push: ## Commitea y pushea  →  make push m='mensaje'
	@if [ -z "$(m)" ]; then \
		echo "❌ Error: Debes proporcionar un mensaje"; \
		echo "   Uso: make push m='tu mensaje de commit'"; \
		exit 1; \
	fi
	@echo "📦 Agregando archivos..."
	@git add .
	@echo "✏️  Commiteando: $(m)"
	@git commit -m "$(m)"
	@echo "🚀 Pusheando a origin/$(BRANCH)..."
	@git push origin $(BRANCH)
	@echo "✅ Push completado exitosamente!"

.PHONY: pull
pull: ## Descarga los últimos cambios desde origin
	@echo "⬇️  Pulling desde origin/$(BRANCH)..."
	@git fetch origin
	@git pull origin $(BRANCH)
	@echo "✅ Pull completado!"

.PHONY: status
status: ## Muestra estado del repo y el último tag publicado
	@echo "📊 Estado de Git (rama: $(BRANCH)) — último tag: $(LAST_TAG)"
	@echo ""
	@git status

.PHONY: sync
sync: ## Pull + commit + push en un solo paso  →  make sync m='mensaje'
	@if [ -z "$(m)" ]; then \
		echo "❌ Error: Debes proporcionar un mensaje"; \
		echo "   Uso: make sync m='tu mensaje de commit'"; \
		exit 1; \
	fi
	@echo "⬇️  Pulling cambios..."
	@git pull origin $(BRANCH)
	@echo "📦 Agregando archivos..."
	@git add .
	@echo "✏️  Commiteando: $(m)"
	@git commit -m "$(m)"
	@echo "🚀 Pusheando a origin/$(BRANCH)..."
	@git push origin $(BRANCH)
	@echo "✅ Sincronización completada!"

.PHONY: log
log: ## Muestra los commits desde el último tag
	@echo "📋 Commits desde $(LAST_TAG):"
	@git log $(LAST_TAG)..HEAD --oneline

# ============================================
# RELEASES (tags)
# ============================================

.PHONY: release-patch
release-patch: check ## Bug fix      $(LAST_TAG) → $(NEXT_PATCH)
	@echo "🔖 $(LAST_TAG) → $(NEXT_PATCH)"
	@git tag $(NEXT_PATCH)
	@git push origin $(NEXT_PATCH)
	@echo "✅ Release $(NEXT_PATCH) publicado."

.PHONY: release-minor
release-minor: check ## Nueva feature  $(LAST_TAG) → $(NEXT_MINOR)
	@echo "🔖 $(LAST_TAG) → $(NEXT_MINOR)"
	@git tag $(NEXT_MINOR)
	@git push origin $(NEXT_MINOR)
	@echo "✅ Release $(NEXT_MINOR) publicado."

.PHONY: release-major
release-major: check ## Breaking change  $(LAST_TAG) → $(NEXT_MAJOR)
	@echo "🔖 $(LAST_TAG) → $(NEXT_MAJOR)"
	@git tag $(NEXT_MAJOR)
	@git push origin $(NEXT_MAJOR)
	@echo "✅ Release $(NEXT_MAJOR) publicado."

# ============================================
# AYUDA
# ============================================

.PHONY: help
help: ## Muestra todos los comandos disponibles
	@echo ""
	@echo "╔══════════════════════════════════════════════════╗"
	@echo "║              dotfilesGo — Makefile               ║"
	@echo "╚══════════════════════════════════════════════════╝"
	@echo ""
	@echo "  Rama activa : $(BRANCH)"
	@echo "  Último tag  : $(LAST_TAG)"
	@echo ""
	@echo "┌─ DESARROLLO ─────────────────────────────────────"
	@echo "│  make run            Ejecuta el CLI sin compilar"
	@echo "│  make build          Compila el binario (versión=$(LAST_TAG))"
	@echo "│  make test           Corre todos los tests"
	@echo "│  make vet            Analiza el código con go vet"
	@echo "│  make check          vet + tests (requerido antes de release)"
	@echo "│  make clean          Elimina el binario compilado"
	@echo "│"
	@echo "├─ GIT ────────────────────────────────────────────"
	@echo "│  make push  m='msg'  Commitea y pushea a $(BRANCH)"
	@echo "│  make pull           Descarga los últimos cambios"
	@echo "│  make sync  m='msg'  Pull + commit + push en un paso"
	@echo "│  make status         Estado del repo + último tag"
	@echo "│  make log            Commits desde $(LAST_TAG)"
	@echo "│"
	@echo "├─ RELEASES ───────────────────────────────────────"
	@echo "│  make release-patch  Bug fix   → $(NEXT_PATCH)"
	@echo "│  make release-minor  Feature   → $(NEXT_MINOR)"
	@echo "│  make release-major  Breaking  → $(NEXT_MAJOR)"
	@echo "│"
	@echo "└──────────────────────────────────────────────────"
	@echo ""
	@echo "  Flujo recomendado:"
	@echo "    1. make sync m='feat: nueva feature'"
	@echo "    2. make release-minor"
	@echo ""
