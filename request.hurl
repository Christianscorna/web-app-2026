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