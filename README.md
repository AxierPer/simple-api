# Simple CRUD API en Go

¡Bienvenido a la API CRUD simple hecha en Go! Esta API permite realizar operaciones básicas (Crear, Leer, Actualizar, Eliminar) sobre un conjunto de usuarios. Está construida utilizando Go y SQLite, y ofrece endpoints RESTful para interactuar con los datos.

## 🛠 Requisitos Previos

Antes de comenzar, asegúrate de tener lo siguiente instalado en tu máquina:

- [Go](https://golang.org/doc/install) (Versión 1.18 o superior)
- [SQLite](https://www.sqlite.org/download.html)

## 🔧 Instalación

Sigue estos pasos para configurar y ejecutar la API:

1. **Clona el repositorio**

   ```bash
   git clone https://github.com/tu-usuario/simple-go-crud-api.git
   cd simple-go-crud-api
   ```
2. **Instala las dependencias**

    Asegúrate de tener Go configurado correctamente en tu entorno, y luego instala las dependencias necesarias:
    ```bash
    go mod tidy
    ```
3. **Compila y ejecuta la API**
    En el directorio raíz del proyecto, ejecuta el siguiente comando para iniciar el servidor:
    ```bash
    go run main.go
    ```
    Esto lanzará el servidor en `http://localhost:8080`.

# 🌐 Endpoints Disponibles
1. **Crear un Usuario**

    **POST** `/users`

    Crea un nuevo usuario.

    **Request Body:**
    ```json
    {
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com"
    }
    ```
    **Respuesta Exitosa:**
    ```json
    {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com"
    }
    ```
2. **Obtener Todos los Usuarios**

      **GET** `/users`

      Obtiene la lista de todos los usuarios.

      **Respuesta Exitosa:**
      ```json
      [
        {
            "id": 1,
            "first_name": "John",
            "last_name": "Doe",
            "email": "john.doe@example.com"
        },
        {
            "id": 2,
            "first_name": "Jane",
            "last_name": "Smith",
            "email": "jane.smith@example.com"
        }
    ]  
      ```
3. **Obtener un Usuario por ID**
    
    **GET** `/users/{id}`

    Obtiene la información de un usuario específico por su ID.

    **Ejemplo:**
    ```bash
    GET http://localhost:8080/users/1
    ```
    **Respuesta Exitosa:**
    ```json
    {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@example.com"
    }
    ```
4. **Actualizar un Usuario**

    **PUT** `/users/{id}`

    Actualiza los datos de un usuario específico.

    **Request Body:**
    ```json
        {
            "first_name": "John",
            "last_name": "Doe",
            "email": "new.email@example.com"
        }
    ```

    **Respuesta Exitosa:**

    ```json
    {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "email": "new.email@example.com"
    }
    ```
5. **Eliminar un Usuario**

    **DELETE** `/users/{id}`

    Elimina un usuario por su ID.

    **Ejemplo:**

    ```bash
    DELETE http://localhost:8080/users/1
    ```
    **Respuesta Exitosa:**

    ```json
    {}
    ```

# 🧑‍💻 Estructura de Carpetas
El proyecto está organizado de la siguiente manera:

```python
simple-go-crud-api/
├── main.go            # Archivo principal que contiene la lógica del servidor y las rutas.
├── db/                # Contiene la lógica de la base de datos.
│   └── db.go          # Funciones para conectarse y realizar operaciones en la base de datos.
├── handlers/          # Contiene la lógica de las rutas y manejo de solicitudes HTTP.
│   └── user_handlers.go # Funciones para manejar las operaciones CRUD de los usuarios.
├── types/             # Contiene los tipos de datos.
│   └── user.go        # El tipo `User` utilizado por la API.
├── go.mod             # Archivo de configuración del módulo Go.
└── go.sum             # Archivo de dependencias de Go.
```
#  🛠 ¿Cómo funciona?
1. **Conexión a la Base de Datos:** La aplicación usa SQLite como base de datos local. Se crea una base de datos llamada `users.db`, que almacena la información de los usuarios. La base de datos se inicializa automáticamente al ejecutar la API.

2. **Manejo de Rutas:** Usamos el paquete gorilla/mux para manejar las rutas. Cada endpoint está asociado con una función que realiza la acción correspondiente en la base de datos (crear, leer, actualizar o eliminar un usuario).

3. **Estructura del Proyecto:** El código está dividido en carpetas para mantener todo organizado. `handlers` maneja las solicitudes HTTP, `db` maneja las operaciones con la base de datos, y `types` contiene los modelos de datos.

# 🔄 Pruebas
Puedes probar la API utilizando herramientas como `Postman` o `curl` desde la línea de comandos. A continuación se muestra un ejemplo de cómo hacer una solicitud usando `curl`:

```bash
curl -X GET http://localhost:8080/users
```
# 💡 Ideas para Mejorar
- **Autenticación:** Agregar autenticación con JWT para proteger los endpoints.
- **Validación de Datos:** Implementar validaciones de entrada para asegurarse de que los datos del usuario sean correctos.
- **Manejo de Errores:** Mejorar el manejo de errores y proporcionar mensajes de error más detallados.
- **Pruebas Unitarias:** Agregar pruebas unitarias para asegurar que los endpoints funcionan correctamente.

# 📜 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.

> ¡Gracias por usar esta API! Si tienes alguna pregunta o sugerencia, no dudes en abrir un issue en el repositorio.
