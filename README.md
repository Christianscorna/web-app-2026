# Proyecto Programación Web 2026

Integrantes:

- **Christian Scornaienchi**
- **Mateo Menchon**
- **Lucio Manuel Storni**

## Sobre el proyecto:
La idea es tener una aplicación web sencilla que permita al usuario jugar aprendiendo, eligiendo las materias que quiere estudiar y el juego le propondrá una trivia con preguntas de las asignaturas elegidas. 

Aprendimos a levantar nuestro propio servidor web en Go que sirve un documento simple HTML. Tambien aprendimos a escribir nuestras primeras líneas de un documento web incluyendo un título, un párrafo y una cabecera con metadata.

Durante la segunda parte de nuestra aplicación, creamos las tablas que van a dar persistencia a las entidades del proyecto y ejecutamos una serie de pruebas unitarias para verificar el correcto funcionamiento de las operaciones CRUD en todas las tablas. 

### Estructura del Proyecto

```
tp-especial/
|-- db/
|   |-- queries/
|   |   └── queries.sql
|   |-- schema/
|   |   └── schema.sql
|   └── sqlc/
|       |-- db.go
|       |-- models.go
|       |-- materia_test.go
|       |-- preguntas_test.go
|       └── opciones_test.go
|-- docker-compose.yml
|-- sqlc.yaml
|-- Makefile
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md
└── static/
    └── index.html
```

### Levantar el proyecto

Descarga este proyecto y abrelo con tu terminal preferida. Luego simplemente párate en la raíz del proyecto, es decir en la carpeta tp-especial y corre: `make test` Si todo ha salido bien, deberías ver mensajes en pantalla de los test pasando correctamente
