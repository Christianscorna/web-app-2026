# Proyecto Programación Web 2026

Integrantes:

- **Christian Scornaienchi**
- **Mateo Menchon**
- **Lucio Manuel Storni**

## Sobre el proyecto:
La idea del proyecto es tener una aplicación web sencilla para probar todas las herramientas que fuimos aprendiendo durante el curso de Programación Web: Crear un contendor, crear y ejecutar archivos go, conectar la interfaz con una base de datos sencilla, automatizar la generación de código go mediante comandos, etc.

Una vez dentro de la página principal, el sitio le da la bienvenida y tendrá dos opciones principales: La de jugar y la de ir a las preguntas. La idea es que se pueda hacer un CRUD completo sobre una entidad y además que haya algo "más para hacer" con estas entidades y ver como se relacionan/interactuan entre ellas. 

Si queres podes agregar una pregunta (siempre en cuando sea respetuosa), la pregunta se asocia automaticamente a una materia, que previamente fue seleccionada cuando buscaba las preguntas de su materia. Podes ver las preguntas, modificarlas por completo e incluso eliminarlas. Acá lo interesante es que la base es consistente con la interfaz y siempre que se elimine una pregunta, se eliminan tambien las opciones a ella relacionadas y no suceden cosas raras como por ejemplo que se borre la materia asociada. 

Luego puede jugar, filtra materias por nombre, elige hasta un máximo de 5 materias (o menos), gira la ruleta y, de la materia que haya salido seleccionada, sale una pregunta aleatoria. 

La ruleta es una prueba de concepto, simplemente es un contenedorcito que almacena las materias seleccionadas (y le permite quitar materias de ella en cualquier momento del juego). 

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

### Levantar el proyecto

Para ejecutar este proyecto es necesario tener instalado sqlc. Puedes descargarlo desde la terminal con los siguientes comandos. 

> go get gorm.io/gorm <br>
> go get gorm.io/driver/postgres

Tambien es necesario tener descargado templ. Puedes descargarlo con el siguiente comando:

> go get github.com/a-h/templ

Puedes acceder a ambas dependencias desde `ls ~/go/bin/` y allí deberían figurar tanto sqlc como templ. 

Descarga este proyecto y abrilo con una terminal bash. Luego simplemente párate en la raíz del proyecto, es decir en la carpeta tp-especial y corre: `make tp5-run` Si todo ha salido bien, deberías ver el proceso de levantado del servidor, base de datos, generación de códigos sqlc y templ y el servidor vivo escuchando en el puerto 8080. Al finalizar tus pruebas recuerda hacer `ctrl + c` para dar de baja el servidor y correr un `make down` para apagar el contenedor y volúmenes viejos. 