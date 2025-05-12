package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "api-server/docs" // Importação dos docs gerados pelo Swag
)

// @title API Go com Swagger
// @version 1.0
// @description API simples em Go demonstrando o uso do Swagger
// @host localhost:8080
// @BasePath /
// @schemes http

// Resposta padrão da API
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// homeHandler godoc
// @Summary Endpoint principal da API
// @Description Retorna uma mensagem de boas-vindas
// @Tags home
// @Produce json
// @Success 200 {object} Response
// @Router / [get]
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := Response{
		Message: "Bem-vindo à API Go",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// healthHandler godoc
// @Summary Verifica a saúde da API
// @Description Retorna o status de saúde do serviço
// @Tags health
// @Produce json
// @Success 200 {object} Response
// @Router /health [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "Serviço disponível",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Definindo as rotas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	
	// Swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Porta onde o servidor irá escutar
	port := 8080
	fmt.Printf("Servidor iniciado na porta %d...\n", port)
	fmt.Println("Acesse a documentação Swagger em: http://localhost:8080/swagger/index.html")
	
	// Iniciando o servidor
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}