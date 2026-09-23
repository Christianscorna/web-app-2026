# request.hurl

# obtener todas las materias
GET http://localhost:8080/materias
HTTP 200

# obtener la materia de id 2
GET  http://localhost:8080/materias/2
HTTP 200

# DELETE http://localhost:8080/materias/3
# HTTP 204

POST http://localhost:8080/materias
Content-Type : application/json
{
    "nombre" : "Programación Web",
    "descripcion" : {
        "String" : "Materia correspondiente a 3er año de Ing Sistemas, dicatada x Ale y Alfre",
        "Valid" : true
    }
}
HTTP 201