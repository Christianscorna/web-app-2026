CREATE TABLE materia (
    id_materia SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255)
);

CREATE TABLE pregunta (
    id_pregunta SERIAL PRIMARY KEY,
    enunciado VARCHAR(500) NOT NULL,
    id_materia INTEGER NOT NULL,

    FOREIGN KEY (id_materia)
    REFERENCES materia(id_materia)

    ON DELETE CASCADE
);

CREATE TABLE opcion (
    id_opcion SERIAL PRIMARY KEY,
    texto VARCHAR(255) NOT NULL,
    es_correcta BOOLEAN NOT NULL,
    id_pregunta INTEGER NOT NULL,

    FOREIGN KEY (id_pregunta)
    REFERENCES pregunta(id_pregunta)

    ON DELETE CASCADE
);


INSERT INTO materia (nombre, descripcion)
VALUES 
    ('Sistemas Operativos', 'Materia de 3er año de IS, dictada por Alejandro Zunino y Matias Hirch'),
    ('Programación Web', 'Materia de 3er año de IS, segundo cuatrimestre, dictada por Alejandro Zunino y Alfredo Teyseyre');