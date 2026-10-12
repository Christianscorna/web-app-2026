# Makefile para gestionar la Base de Datos PostgreSQL con Docker Compose

.PHONY: up down restart logs ps shell help test hurl-test run

# Comando por defecto (ayuda)
help:
	@echo "Comandos disponibles:"
	@echo "  make up      - Levanta el contenedor de PostgreSQL en segundo plano."
	@echo "  make down    - Detiene y elimina el contenedor."
	@echo "  make restart - Reinicia el contenedor."
	@echo "  make logs    - Muestra los logs en tiempo real."
	@echo "  make ps      - Lista el estado de los contenedores."
	@echo "  make shell   - Entra al shell psql dentro del contenedor para hacer consultas."
	@echo "  make clean   - Detiene y elimina con vervosidad."
	@echo "  make test    - Corre el programa."

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose down && docker compose up -d

logs:
	docker compose logs -f database

ps:
	docker compose ps

shell:
	docker compose exec database psql -U postgres -d web_db

clean:
	docker compose down -v

wait-db:
	@echo "Esperando a que PostgreSQL este listo..."
	@until docker compose exec -T database pg_isready -U postgres -d web_db > /dev/null 2>&1; do \
		echo "  [Aún no lista, reintentando en 1s...]"; \
		sleep 1; \
	done
	@echo "La bbdd esta lista y aceptando conexiones"

test: restart wait-db
	@echo "\generando código sqlc"
	~/go/bin/sqlc generate

	@echo "\nCorriendo la suite de tests de prueba:"
	go test -v ./db/tests
	
	@echo "\nBorrando contenedores y volúmenes viejos..."
	docker compose down -v

run: restart wait-db
	@set -e; \
	server_bin=$$(mktemp); server_pid=; \
	cleanup() { \
		if [ -n "$$server_pid" ]; then kill "$$server_pid" 2>/dev/null || true; wait "$$server_pid" 2>/dev/null || true; fi; \
		rm -f "$$server_bin"; \
		docker compose down -v; \
	}; \
	trap cleanup EXIT; \
	go build -o "$$server_bin" .; \
	"$$server_bin" & server_pid=$$!; \
	attempt=0; \
	sleep 2; \
	until curl -fsS http://localhost:8080/materias >/dev/null; do \
		if ! kill -0 "$$server_pid" 2>/dev/null; then echo "El servidor terminó antes de estar disponible"; exit 1; fi; \
		attempt=$$((attempt + 1)); \
		if [ "$$attempt" -ge 30 ]; then echo "El servidor no respondió dentro del tiempo esperado"; exit 1; fi; \
		sleep 1; \
	done; \
	hurl --test request.hurl