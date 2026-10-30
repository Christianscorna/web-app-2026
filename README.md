# Proyecto Programación Web 2026

Integrantes:

- **Christian Scornaienchi**
- **Mateo Menchon**
- **Lucio Manuel Storni**

## Sobre el proyecto:
La idea es tener una aplicación web sencilla que permita al usuario jugar aprendiendo, eligiendo las materias que quiere estudiar y el juego le propondrá una trivia con preguntas de las asignaturas elegidas. Además es posible agregar preguntas desde interfaz si el usuario lo desea. 

Entidades:
#### Materia 
 - Atributos: idMateria, nombre, descripcion
 - Descripción: Es un elemento seleccionable de entre una lista de materias con las cuales el usuario puede jugar. 

#### Pregunta
 - Atributos: idPregunta, consigna, idMateria 
 - Descripción: Pregunta asociada a una materia, es lo que se va a presentar al usuario para responder. 

#### Opcion
 - Atributos: idOpción, texto, es_valida, idPregunta
 - Descripción: Opcion asociada a una materia. Se establece si es válida o no para comparar con la respuesta del jugador. 

#### Tp 1
Aprendimos a levantar nuestro propio servidor web en Go que sirve un documento simple HTML. Tambien aprendimos a escribir nuestras primeras líneas de un documento web incluyendo un título, un párrafo y una cabecera con metadata.

#### Tp 2
Durante la segunda parte de nuestra aplicación, creamos las tablas que van a dar persistencia a las entidades del proyecto y ejecutamos una serie de pruebas unitarias para verificar el correcto funcionamiento de las operaciones CRUD en todas las tablas. 

#### Tp 3
Durante esta tercera parte del proyecto hemos aprendido a controlar los endpoints de nuestra aplicación para que responda a cada una de las distintas solicitudes HTTP. Aprendimos sobre códigos de estado y a escribir consutlas cURL y a automatizarlas pruebas mediante un archivo de pruebas hURL. 

#### Tp 4
Esta parte del trabajo ha sido omitida por la cátedra para la entrega por cuestiones de tiempo. 

#### Tp 5 
En esta parte del proyecto hemos trabajado con Renderizado del Lado del Servidor con Templ, una tecnología que tal como lo hicimos con sqlc generate en su momento, nos permite generar código go automáticamente a raiz de los templates de código html. 

### Sobre la arquitectura del proyecto
Este proyecto utiliza una combinación de arquitectura cliente-servidor y organización por capas. El navegador funciona como cliente y se comunica mediante HTTP con el servidor desarrollado en go. La app se organiza en distintas responsabilidades: la capa de acceso a datos, capa de handlers y la capa de presentación. 

Este modo de organización permite organizar el código, facilitar su mantenimiento, entendimiento y distinción entre las diferentes partes de la lógica ya sea que este relacionada con la persistencia, procesamiento de solicitudes HTTP o atender a la interfaz de usuario. 

### Estructura del Proyecto

```
tp-especial/
|-- db/
|   |-- queries/
|   |   └── queries.sql
|   |-- schema/
|   |   └── schema.sql
|   └── tests/
|   |   |-- materia_test.go
|   |   |-- preguntas_test.go
|   |   └── opciones_test.go
|   └── coneccionbd.go
|-- handlers/
|   |-- materia_handler.go
|   |-- opcion_handler.go
|   └── pregunta_handler.go
|-- docker-compose.yml
|-- sqlc.yaml
|-- Makefile
|-- go.mod
|-- go.sum
|-- main.go
|-- request.hurl
|-- README.md
└── static/
    |-- crear_pregunta.html
    └── index.html
```

### Levantar el proyecto

Para ejecutar este proyecto es necesario tener instalado sqlc. Puedes descargarlo desde la terminal con los siguientes comandos. 

> go get gorm.io/gorm <br>
> go get gorm.io/driver/postgres

Descarga este proyecto y abrilo con una terminal bash. Luego simplemente párate en la raíz del proyecto, es decir en la carpeta tp-especial y corre: `make run` Si todo ha salido bien, deberías ver la salida de hurl exitosa de todos los comandos. 