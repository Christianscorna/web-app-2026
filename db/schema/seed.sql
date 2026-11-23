
-- ==========================================
-- DATOS DE LAS PREGUNTAS Y SUS OPCIONES
-- 10 preguntas por materia
-- ==========================================

WITH datos(nombre_materia, enunciado, opciones) AS (
    VALUES

    -- ======================================
    -- SISTEMAS OPERATIVOS
    -- ======================================

    (
        'Sistemas Operativos',
        '¿Cuál es la función principal de un sistema operativo?',
        '[
            {"texto":"Administrar los recursos del hardware y proporcionar servicios a los programas.","correcta":true},
            {"texto":"Diseñar las páginas web del usuario.","correcta":false},
            {"texto":"Reemplazar completamente al procesador.","correcta":false},
            {"texto":"Convertir todos los archivos en código fuente.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué diferencia existe entre un proceso y un hilo?',
        '[
            {"texto":"Un hilo siempre tiene más memoria que un proceso.","correcta":false},
            {"texto":"Un proceso tiene su propio espacio de direcciones, mientras que los hilos de un proceso comparten ese espacio.","correcta":true},
            {"texto":"Un proceso solamente puede ejecutar una instrucción.","correcta":false},
            {"texto":"No existe ninguna diferencia entre ambos conceptos.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué algoritmo de planificación asigna la CPU por turnos utilizando un quantum?',
        '[
            {"texto":"First Come First Served.","correcta":false},
            {"texto":"Shortest Job First.","correcta":false},
            {"texto":"Planificación por prioridad sin desalojo.","correcta":false},
            {"texto":"Round Robin.","correcta":true}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué sucede cuando un proceso realiza una operación de entrada/salida y debe esperar su finalización?',
        '[
            {"texto":"El proceso termina definitivamente.","correcta":false},
            {"texto":"El proceso pasa al estado bloqueado.","correcta":true},
            {"texto":"El proceso pasa automáticamente a ejecutarse.","correcta":false},
            {"texto":"El proceso elimina su espacio de memoria.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué es un deadlock o interbloqueo?',
        '[
            {"texto":"Una situación en la que dos o más procesos quedan esperando indefinidamente recursos retenidos por otros.","correcta":true},
            {"texto":"Un error producido exclusivamente por falta de espacio en disco.","correcta":false},
            {"texto":"Una técnica para acelerar la ejecución de procesos.","correcta":false},
            {"texto":"La finalización normal de todos los procesos.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué función cumple la memoria virtual?',
        '[
            {"texto":"Eliminar la necesidad de memoria RAM.","correcta":false},
            {"texto":"Almacenar permanentemente los archivos del usuario.","correcta":false},
            {"texto":"Permitir que los procesos utilicen un espacio de direcciones virtual mayor que la memoria física disponible.","correcta":true},
            {"texto":"Reemplazar la unidad de procesamiento.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué es una interrupción?',
        '[
            {"texto":"Una señal que solicita la atención del procesador para atender un evento.","correcta":true},
            {"texto":"Una instrucción que elimina todos los procesos.","correcta":false},
            {"texto":"Un archivo especial del sistema operativo.","correcta":false},
            {"texto":"Una operación exclusiva de los discos rígidos.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué técnica permite que varios procesos compartan la CPU alternando su ejecución?',
        '[
            {"texto":"Fragmentación externa.","correcta":false},
            {"texto":"Planificación de procesos.","correcta":true},
            {"texto":"Compresión de archivos.","correcta":false},
            {"texto":"Formateo de memoria.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Cuál es la función principal del planificador de corto plazo?',
        '[
            {"texto":"Elegir qué proceso de los que están listos obtiene la CPU.","correcta":true},
            {"texto":"Decidir qué archivos deben eliminarse.","correcta":false},
            {"texto":"Administrar exclusivamente el espacio en disco.","correcta":false},
            {"texto":"Crear las cuentas de los usuarios.","correcta":false}
        ]'::jsonb
    ),
    (
        'Sistemas Operativos',
        '¿Qué es la paginación en la administración de memoria?',
        '[
            {"texto":"Una técnica para ordenar archivos alfabéticamente.","correcta":false},
            {"texto":"Un mecanismo para aumentar la velocidad del procesador.","correcta":false},
            {"texto":"Una técnica que divide la memoria virtual y física en bloques de tamaño fijo.","correcta":true},
            {"texto":"Un método para ejecutar procesos sin memoria.","correcta":false}
        ]'::jsonb
    ),

    -- ======================================
    -- PROGRAMACIÓN WEB
    -- ======================================

    (
        'Programación Web',
        '¿Qué protocolo se utiliza habitualmente para la comunicación entre un navegador y un servidor web?',
        '[
            {"texto":"FTP.","correcta":false},
            {"texto":"SMTP.","correcta":false},
            {"texto":"HTTP.","correcta":true},
            {"texto":"SSH.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Cuál es la función principal de HTML?',
        '[
            {"texto":"Administrar las conexiones con PostgreSQL.","correcta":false},
            {"texto":"Definir la estructura y el contenido de una página web.","correcta":true},
            {"texto":"Ejecutar consultas SQL.","correcta":false},
            {"texto":"Controlar directamente los procesos del sistema operativo.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué método HTTP se utiliza habitualmente para solicitar información a un servidor?',
        '[
            {"texto":"POST.","correcta":false},
            {"texto":"DELETE.","correcta":false},
            {"texto":"PATCH.","correcta":false},
            {"texto":"GET.","correcta":true}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Para qué se utiliza principalmente CSS?',
        '[
            {"texto":"Para definir la presentación y los estilos de una página web.","correcta":true},
            {"texto":"Para almacenar datos en una base de datos.","correcta":false},
            {"texto":"Para administrar procesos del sistema operativo.","correcta":false},
            {"texto":"Para reemplazar el protocolo HTTP.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué significa que una aplicación web utilice una arquitectura cliente-servidor?',
        '[
            {"texto":"Que todos los usuarios ejecutan el servidor en sus navegadores.","correcta":false},
            {"texto":"Que el cliente solicita servicios o recursos y el servidor responde a esas solicitudes.","correcta":true},
            {"texto":"Que no existe comunicación por red.","correcta":false},
            {"texto":"Que el servidor solamente puede atender a un usuario.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué código de estado HTTP indica que una solicitud fue procesada correctamente?',
        '[
            {"texto":"404.","correcta":false},
            {"texto":"500.","correcta":false},
            {"texto":"200.","correcta":true},
            {"texto":"301.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué característica tiene el formato JSON?',
        '[
            {"texto":"Es un formato de intercambio de datos basado en texto.","correcta":true},
            {"texto":"Es un lenguaje exclusivo para diseñar estilos.","correcta":false},
            {"texto":"Es un motor de bases de datos relacionales.","correcta":false},
            {"texto":"Es un protocolo de transferencia de archivos.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Para qué se utiliza el método POST de HTTP?',
        '[
            {"texto":"Para solicitar exclusivamente archivos CSS.","correcta":false},
            {"texto":"Para enviar datos al servidor, por ejemplo al crear un recurso.","correcta":true},
            {"texto":"Para eliminar automáticamente una base de datos.","correcta":false},
            {"texto":"Para mostrar una página sin comunicarse con el servidor.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué función cumple una API en una aplicación web?',
        '[
            {"texto":"Definir una interfaz para que distintos componentes o sistemas puedan comunicarse.","correcta":true},
            {"texto":"Reemplazar el sistema operativo del servidor.","correcta":false},
            {"texto":"Diseñar exclusivamente los colores de una página.","correcta":false},
            {"texto":"Impedir toda comunicación entre cliente y servidor.","correcta":false}
        ]'::jsonb
    ),
    (
        'Programación Web',
        '¿Qué ventaja ofrece utilizar una base de datos en una aplicación web?',
        '[
            {"texto":"Evitar que los usuarios puedan enviar solicitudes.","correcta":false},
            {"texto":"Hacer que el navegador ejecute código SQL directamente.","correcta":false},
            {"texto":"Permitir almacenar, consultar y administrar información de manera persistente.","correcta":true},
            {"texto":"Eliminar la necesidad de un servidor.","correcta":false}
        ]'::jsonb
    ),

    -- ======================================
    -- ARQUITECTURA DE COMPUTADORAS I
    -- ======================================

    (
        'Arquitectura de Computadoras I',
        '¿Qué componente del procesador realiza operaciones aritméticas y lógicas?',
        '[
            {"texto":"El contador de programa (PC).","correcta":false},
            {"texto":"El registro de instrucciones (IR).","correcta":false},
            {"texto":"La Unidad Aritmético-Lógica (ALU).","correcta":true},
            {"texto":"El bus de direcciones.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Cuál es la función del contador de programa (PC)?',
        '[
            {"texto":"Almacenar el resultado de todas las operaciones aritméticas.","correcta":false},
            {"texto":"Almacenar la dirección de la próxima instrucción que se debe buscar.","correcta":true},
            {"texto":"Almacenar permanentemente los datos de la memoria principal.","correcta":false},
            {"texto":"Decodificar las instrucciones del procesador.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué característica distingue a la arquitectura Harvard de la arquitectura Von Neumann?',
        '[
            {"texto":"No utiliza memoria para almacenar instrucciones.","correcta":false},
            {"texto":"No posee unidad de control.","correcta":false},
            {"texto":"Utiliza un único registro para todas las operaciones.","correcta":false},
            {"texto":"Utiliza memorias separadas para instrucciones y datos.","correcta":true}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué función cumple el registro de instrucciones (IR)?',
        '[
            {"texto":"Almacenar la instrucción que está siendo ejecutada o decodificada.","correcta":true},
            {"texto":"Almacenar la dirección de la próxima instrucción.","correcta":false},
            {"texto":"Controlar la temperatura del procesador.","correcta":false},
            {"texto":"Almacenar exclusivamente resultados de operaciones de entrada/salida.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Cuál es la función de la unidad de control del procesador?',
        '[
            {"texto":"Realizar exclusivamente sumas y restas.","correcta":false},
            {"texto":"Coordinar y controlar la ejecución de las instrucciones.","correcta":true},
            {"texto":"Almacenar todos los programas de forma permanente.","correcta":false},
            {"texto":"Proporcionar energía eléctrica a la computadora.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué representa el ciclo de instrucción de un procesador?',
        '[
            {"texto":"El proceso de encender y apagar la computadora.","correcta":false},
            {"texto":"La secuencia de buscar, decodificar y ejecutar una instrucción.","correcta":true},
            {"texto":"La instalación de un sistema operativo.","correcta":false},
            {"texto":"La transferencia de archivos entre discos.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué característica tiene la memoria RAM?',
        '[
            {"texto":"Es una memoria volátil que pierde su contenido al cortar la alimentación.","correcta":true},
            {"texto":"Conserva todos sus datos sin energía eléctrica.","correcta":false},
            {"texto":"Solo puede almacenar instrucciones del procesador.","correcta":false},
            {"texto":"Es siempre más lenta que un disco rígido.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué función cumple el bus de direcciones?',
        '[
            {"texto":"Transportar exclusivamente resultados aritméticos.","correcta":false},
            {"texto":"Indicar la dirección de memoria o dispositivo al que se desea acceder.","correcta":true},
            {"texto":"Almacenar instrucciones de manera permanente.","correcta":false},
            {"texto":"Ejecutar las instrucciones del programa.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué unidad se utiliza habitualmente para medir la frecuencia de un procesador?',
        '[
            {"texto":"Bytes.","correcta":false},
            {"texto":"Voltios.","correcta":false},
            {"texto":"Hertz (Hz).","correcta":true},
            {"texto":"Amperes.","correcta":false}
        ]'::jsonb
    ),
    (
        'Arquitectura de Computadoras I',
        '¿Qué sucede durante la etapa de decodificación de una instrucción?',
        '[
            {"texto":"Se interpreta la instrucción para determinar qué operación debe realizarse.","correcta":true},
            {"texto":"Se elimina el programa de la memoria.","correcta":false},
            {"texto":"Se apaga el procesador.","correcta":false},
            {"texto":"Se almacenan todos los archivos del usuario.","correcta":false}
        ]'::jsonb
    )
),

preguntas_insertadas AS (
    INSERT INTO pregunta (enunciado, id_materia)
    SELECT d.enunciado, m.id_materia
    FROM datos d
    JOIN materia m ON m.nombre = d.nombre_materia
    RETURNING id_pregunta, enunciado
)

INSERT INTO opcion (texto, es_correcta, id_pregunta)
SELECT
    elemento.opcion->>'texto',
    (elemento.opcion->>'correcta')::boolean,
    p.id_pregunta
FROM preguntas_insertadas p
JOIN datos d ON d.enunciado = p.enunciado
CROSS JOIN LATERAL jsonb_array_elements(d.opciones)
    AS elemento(opcion);
