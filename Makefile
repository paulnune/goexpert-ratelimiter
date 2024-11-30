# Makefile simplificado

# Targets principais

all: up install-lynx sleep read-files

up:
	@echo "📦 Starting services with Docker Compose..."
	@docker compose up -d
	@echo "✅ Services are up."

down:
	@echo "🛑 Stopping services with Docker Compose..."
	@docker compose down
	@echo "✅ Services are down."

restart:
	@echo "🔄 Restarting services with Docker Compose..."
	@make down
	@make up
	@echo "✅ Services restarted."

help:
	@echo "📜 Makefile Help:"
	@echo "  up               -> Start services with Docker Compose"
	@echo "  down             -> Stop services with Docker Compose"
	@echo "  restart          -> Restart services with Docker Compose"
	@echo "  install-lynx     -> Install lynx on your system"
	@echo "  read-files       -> Read HTML files with lynx"
	@echo "  help             -> Display this help message"

install-lynx:
	@echo "🔧 Installing lynx..."
	@if [ -f /etc/redhat-release ]; then \
		echo "Detectado Fedora/RHEL/CentOS. Instalando com dnf..."; \
		sudo dnf install -y lynx; \
	elif [ -f /etc/debian_version ]; then \
		echo "Detectado Debian/Ubuntu. Instalando com apt..."; \
		sudo apt update && sudo apt install -y lynx; \
	else \
		echo "Distribuição não suportada."; \
		exit 1; \
	fi
	@echo "✅ Lynx instalado."

read-files:
	@echo "📖 Lendo arquivos HTML com lynx..."
	@lynx -dump stress/summary-ip.html
	@lynx -dump stress/summary-token.html
	@echo "✅ Leitura concluída."

sleep:
	@echo "⏳ Aguardando geração de arquivos HTML para leitura com o Lynx (90 segundos)..."
	@sleep 90
	@echo "⏳ Tempo de espera concluído."

.PHONY: all up down restart help install-lynx read-files sleep
