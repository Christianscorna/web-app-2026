# Makefile para gestionar la Base de Datos PostgreSQL con Docker Compose

.PHONY: up down restart logs ps shell help

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
	@echo "  make test     -Corre el programa."

# Levanta el contenedor en segundo plano (detached)
up:
	docker compose up -d

# Detiene y elimina el contenedor y sus recursos
down:
	docker compose down

# Detiene y levanta de nuevo el contenedor
restart:
	docker compose down && docker compose up -d

# Ver los logs del contenedor de Postgres
logs:
	docker compose logs -f database

# Ver el estado de los contenedores
ps:
	docker compose ps

# Conectarse a la base de datos a través de psql dentro del contenedor
shell:
	docker compose exec database psql -U postgres -d web_db

clean:
	docker compose down -v

test: restart
	@echo "\nLevantando el programa, espere por favorcito..."
	sleep 10
	@echo "\nCorriendo la suite de tests de prueba:"
	go test -v ./db/sqlc
	@echo "\nBorrando contenedores y volúmenes viejos..."
	docker compose down -v
