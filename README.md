
Siga los Siguinetes pasos para leventar el proyecto :

Instalar go : 
go version go1.22.0 darwin/amd64
instalar postman

instalar postgres
instalar postman


Crear base de datos : 

CREATE DATABASE apigo;

  asegúrate de ejecutar los scripts SQL para crear las tablas users, drugs, y vaccinations en tu base de datos apigo


Script para la creacion de tablas :

-- Script para crear la tabla 'User'
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Script para crear la tabla 'Drug'
CREATE TABLE drugs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    approved BOOLEAN NOT NULL,
    min_dose INTEGER NOT NULL,
    max_dose INTEGER NOT NULL,
    available_at TIMESTAMP WITHOUT TIME ZONE
);

-- Script para crear la tabla 'Vaccination'
CREATE TABLE vaccinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    drug_id INTEGER NOT NULL,
    dose INTEGER NOT NULL,
    date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    FOREIGN KEY (drug_id) REFERENCES drugs(id)
);


Ejecuta el siguiente comando para descargar las dependencias especificadas en tu go.mod:
     
     go mod tidy

go clean -modcache     Limpiar cache

go build ./...         compilar proeycto

go run cmd/api/main.go      ejecutar proyecto  



 Docker (Opcional): Si prefieres contenerizar tu aplicación y la base de datos, asegúrate de tener Docker y Docker Compose instalados. Luego, en la raíz de tu proyecto, ejecuta:
    docker-compose up --build



    A continuacion dejo las APIS Desarrolladas basadas en la indicaciones del Reto Tecnico:

-------------------------------------------
POST http://localhost:8080/signup

Headers : 

Content-Type : application/json

json :

   {
     "name": "denzel",
     "email": "denzel@gmail.com",
     "password": "123456"
   }



-------------------------------------------
POST  http://localhost:8080/login


Headers : 

Content-Type : application/json


json: 

{
    "email": "denzel@gmail.com",
    "password": "123456"
}



-------------------------------------------

POST http://localhost:8080/drugs

Headers:

Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDMxNDYsInN1YiI6MTR9.uPq5qRyIP0_lYGFqnTWN1sd7YsRkwX3VXySYwGD89Nk    (cambiar por el token generado)

json : 

{
    "name": "Ibuprofeno",
    "approved": true,
    "min_dose": 200,
    "max_dose": 800,
    "available_at": "2023-04-01T15:00:00Z"
}



-------------------------------------------

PUT http://localhost:8080/drugs/9    el numero 9 reemplazar por el Id generado que vas a actualizar

Headers : 

Content-Type : application/json
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwMjUwNjAsInN1YiI6MTN9.VzoPSyljCXk5Vru7qQzMtX345Z0qcYNWTbGikMxbnUw


json : 

{
    "name": "Ibuprofeno Modificado dos",
    "approved": true,
    "min_dose": 400,
    "max_dose": 1200,
    "available_at": "2023-05-01T15:00:00Z"
} 



-------------------------------------------

GET http://localhost:8080/drugs

Headers : 

Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwMjUwNjAsInN1YiI6MTN9.VzoPSyljCXk5Vru7qQzMtX345Z0qcYNWTbGikMxbnUw



-------------------------------------------


DELETE http://localhost:8080/drugs/12    numero 12 cambiarlo por el Id a Eliminar

Headers : 

Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDMxNDYsInN1YiI6MTR9.uPq5qRyIP0_lYGFqnTWN1sd7YsRkwX3VXySYwGD89Nk


-------------------------------------------

POST http://localhost:8080/vaccination

Headers : 

Content-Type : application/json
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDMxNDYsInN1YiI6MTR9.uPq5qRyIP0_lYGFqnTWN1sd7YsRkwX3VXySYwGD89Nk


json :

{
    "name": "Vacunación de prueba",
    "drug_id": 6,
    "dose": 200,
    "date": "2023-04-02T15:00:00Z"
}


-------------------------------------------

PUT http://localhost:8080/vaccination/4 

Content-Type : application/json
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDMxNDYsInN1YiI6MTR9.uPq5qRyIP0_lYGFqnTWN1sd7YsRkwX3VXySYwGD89Nk


Headers: 

Content-Type : application/json
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDMxNDYsInN1YiI6MTR9.uPq5qRyIP0_lYGFqnTWN1sd7YsRkwX3VXySYwGD89Nk

jason : 

{
    "name": "Vacunación Actualizada",
    "drug_id": 2,
    "dose": 250,
    "date": "2023-06-01T15:00:00Z"
}


-------------------------------------------

GET http://localhost:8080/vaccination

Headers : 

Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDEyMzQsInN1YiI6MTN9.QTvH2EuiGm0zEVbMk13NdS4HnkOIf4HWmgUzf_N91uA


-------------------------------------------

DELETE http://localhost:8080/vaccination/4

Headers : 

Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEwNDEyMzQsInN1YiI6MTN9.QTvH2EuiGm0zEVbMk13NdS4HnkOIf4HWmgUzf_N91uA





