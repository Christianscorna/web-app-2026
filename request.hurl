# request.hurl

POST http://localhost:8080/materias
Content-Type: application/json
{
    "nombre": "Programación Web Hurl",
    "descripcion": {
        "String": "Materia creada por la prueba de integración",
        "Valid": true
    }
}
HTTP 201
[Captures]
materia_id: jsonpath "$.id_materia"

GET http://localhost:8080/materias
HTTP 200
[Asserts]
jsonpath "$[*].id_materia" contains {{materia_id}}

GET http://localhost:8080/materias/{{materia_id}}
HTTP 200
[Asserts]
jsonpath "$.nombre" == "Programación Web Hurl"

PUT http://localhost:8080/materias/{{materia_id}}
Content-Type: application/json
{
    "nombre": "Programación Web Hurl actualizada",
    "descripcion": {
        "String": "Materia actualizada por la prueba de integración",
        "Valid": true
    }
}
HTTP 204

GET http://localhost:8080/materias/{{materia_id}}
HTTP 200
[Asserts]
jsonpath "$.nombre" == "Programación Web Hurl actualizada"

DELETE http://localhost:8080/materias/{{materia_id}}
HTTP 204

GET http://localhost:8080/materias/{{materia_id}}
HTTP 404

# MATERIAS
# Obtengo las dos materias existentes y capturo sus IDs
GET http://localhost:8080/materias
HTTP 200
[Asserts]
jsonpath "$[0].nombre" == "Programación Web"
jsonpath "$[1].nombre" == "Sistemas Operativos"
[Captures]
materia_programacion: jsonpath "$[0].id_materia"
materia_sistemas: jsonpath "$[1].id_materia"

# PREGUNTAS
# Creo una pregunta para Programación Web
POST http://localhost:8080/preguntas
Content-Type: application/json
{
    "enunciado": "¿Qué es un endpoint?",
    "id_materia": {{materia_programacion}}
}
HTTP 201
[Captures]
pregunta_programacion: jsonpath "$.id_pregunta"

# Verifico que puedo obtener la pregunta de Programación Web
GET http://localhost:8080/preguntas/{{pregunta_programacion}}
HTTP 200
[Asserts]
jsonpath "$.enunciado" == "¿Qué es un endpoint?"

# Actualizo la pregunta de Programación Web
PUT http://localhost:8080/preguntas/{{pregunta_programacion}}
Content-Type: application/json
{
    "enunciado": "¿Qué es una API?",
    "id_materia": {{materia_programacion}}
}
HTTP 204

# Verifico que la pregunta de Programación Web fue modificada
GET http://localhost:8080/preguntas/{{pregunta_programacion}}
HTTP 200
[Asserts]
jsonpath "$.enunciado" == "¿Qué es una API?"

# Creo una pregunta para Sistemas Operativos
POST http://localhost:8080/preguntas
Content-Type: application/json
{
    "enunciado": "¿Qué es un proceso?",
    "id_materia": {{materia_sistemas}}
}
HTTP 201
[Captures]
pregunta_sistemas: jsonpath "$.id_pregunta"

# Verifico que puedo obtener la pregunta de Sistemas Operativos
GET http://localhost:8080/preguntas/{{pregunta_sistemas}}
HTTP 200
[Asserts]
jsonpath "$.enunciado" == "¿Qué es un proceso?"

# Actualizo la pregunta de Sistemas Operativos
PUT http://localhost:8080/preguntas/{{pregunta_sistemas}}
Content-Type: application/json
{
    "enunciado": "¿Qué es un hilo?",
    "id_materia": {{materia_sistemas}}
}
HTTP 204

# Verifico que la pregunta de Sistemas Operativos fue modificada
GET http://localhost:8080/preguntas/{{pregunta_sistemas}}
HTTP 200
[Asserts]
jsonpath "$.enunciado" == "¿Qué es un hilo?"

# OPCIONES
# Creo una opción para la pregunta de Programación Web
POST http://localhost:8080/opciones
Content-Type: application/json
{
    "texto": "Una dirección para acceder a una funcionalidad del servidor",
    "es_correcta": true,
    "id_pregunta": {{pregunta_programacion}}
}
HTTP 201
[Captures]
opcion_programacion: jsonpath "$.id_opcion"

# Verifico que puedo obtener la opción de Programación Web
GET http://localhost:8080/opciones/{{opcion_programacion}}
HTTP 200
[Asserts]
jsonpath "$.texto" == "Una dirección para acceder a una funcionalidad del servidor"

# Actualizo la opción de Programación Web
PUT http://localhost:8080/opciones/{{opcion_programacion}}
Content-Type: application/json
{
    "texto": "Una ruta que identifica una operación de la API",
    "es_correcta": true
}
HTTP 204

# Verifico que la opción de Programación Web fue modificada
GET http://localhost:8080/opciones/{{opcion_programacion}}
HTTP 200
[Asserts]
jsonpath "$.texto" == "Una ruta que identifica una operación de la API"

# Elimino la opción de Programación Web
DELETE http://localhost:8080/opciones/{{opcion_programacion}}
HTTP 204

# Verifico que la opción de Programación Web ya no existe
GET http://localhost:8080/opciones/{{opcion_programacion}}
HTTP 404

# Creo una opción para la pregunta de Sistemas Operativos
POST http://localhost:8080/opciones
Content-Type: application/json
{
    "texto": "Un programa en ejecución",
    "es_correcta": true,
    "id_pregunta": {{pregunta_sistemas}}
}
HTTP 201
[Captures]
opcion_sistemas: jsonpath "$.id_opcion"

# Verifico que puedo obtener la opción de Sistemas Operativos
GET http://localhost:8080/opciones/{{opcion_sistemas}}
HTTP 200
[Asserts]
jsonpath "$.texto" == "Un programa en ejecución"

# Actualizo la opción de Sistemas Operativos
PUT http://localhost:8080/opciones/{{opcion_sistemas}}
Content-Type: application/json
{
    "texto": "Una instancia de un programa que se está ejecutando",
    "es_correcta": true
}
HTTP 204

# Verifico que la opción de Sistemas Operativos fue modificada
GET http://localhost:8080/opciones/{{opcion_sistemas}}
HTTP 200
[Asserts]
jsonpath "$.texto" == "Una instancia de un programa que se está ejecutando"

# Elimino la opción de Sistemas Operativos
DELETE http://localhost:8080/opciones/{{opcion_sistemas}}
HTTP 204

# Verifico que la opción de Sistemas Operativos ya no existe
GET http://localhost:8080/opciones/{{opcion_sistemas}}
HTTP 404

# PREGUNTAS
# Elimino la pregunta de Programación Web
DELETE http://localhost:8080/preguntas/{{pregunta_programacion}}
HTTP 204

# Verifico que la pregunta de Programación Web ya no existe
GET http://localhost:8080/preguntas/{{pregunta_programacion}}
HTTP 404

# Elimino la pregunta de Sistemas Operativos
DELETE http://localhost:8080/preguntas/{{pregunta_sistemas}}
HTTP 204

# Verifico que la pregunta de Sistemas Operativos ya no existe
GET http://localhost:8080/preguntas/{{pregunta_sistemas}}
HTTP 404

