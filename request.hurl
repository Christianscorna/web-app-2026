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

# PREGUNTAS 
# ahora habría que adaptarlo para que lo acepte hurl. 

# Primera consulta agrego pregunta a materia 
curl -X POST http://localhost:8080/preguntas \
    -H "Content-Type: application/json" \
    -d '{ 
        "enunciado" : "Qué es un volumen?", 
        "id_materia" : 6 
    }'

# Agrego otra pregunta a la misma materia
curl -X POST http://localhost:8080/preguntas \
    -H "Content-Type: application/json" \
    -d '{ "enunciado" : "Qué es Docker?", "id_materia" : 6 }'

# Agrego una pregunta a una materia diferente
curl -X POST http://localhost:8080/preguntas \
    -H "Content-Type: application/json" \
    -d '{ "enunciado" : "Qué forma tiene la ecuación dada por x² + y² = 1", "id_materia" : 2 }'

# Agregamos una pregunta mal a propósito
curl -X POST http://localhost:8080/preguntas
    -H "Content-Type: application/json" \
    -d '{ "enunciado" : "Mi unica función es ser eliminada", "id_materia" : 2 }'

# Borramos la pregunta mal hecha
curl -X DELETE http://localhost:8080/preguntas/6

# Ponemos una pregunta equivocada
curl -X POST http://localhost:8080/preguntas
    -H "Content-Type: application/json" \
    -d '{ "enunciado" : "Mi unica funcion aquí es ser actualizada", "id_materia" : 1 }' 

# Acá corregimos la pregunta que nos equivocamos
curl -X PUT http://localhost:8080/preguntas/7 \
    -H "Content-Type: application/json" \
    -d '{"enunciado": "Qué significa TLC?", "id_materia": 1}'

# Si borras una materia con preguntas tambien se eliminan las preguntas asociadas a la misma
curl -X DELETE http://localhost:8080/materias/1

# luego al intentar traer esas preguntas asociadas a esa materia no trae nada --> null
curl http://localhost:8080/preguntas?id_materia=1