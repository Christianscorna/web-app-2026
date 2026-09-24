# Makefile para gestionar la Base de Datos PostgreSQL con Docker Compose

.PHONY: up down restart logs ps shell help test hurl-test

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
	go run .