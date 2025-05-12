1. docker build -t go-api-swagger .

2. docker run -d -p 8080:8080 go-api-swagger:latest

acesse http://localhost:8080 para a rota default da api, que apenas retorna um 200.

OU

acesse http://localhost:8080/swagger para checar sua documentação by Swagger.