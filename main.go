package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Transacao representa os dados recebidos pelo microserviço
type Transacao struct {
	ID        string  `json:"id"`
	Moeda     string  `json:"moeda"`
	ValorBRL  float64 `json:"valor_brl"`
	Status    string  `json:"status"`
	Risco     string  `json:"risco"`
	Timestamp string  `json:"timestamp"`
}

// handlerAnalisar processa requisições e aplica regras de antifraude
func handlerAnalisar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simula uma transação de alta movimentação para demonstração via navegador (GET)
	tx := Transacao{
		ID:        "TX-99841",
		Moeda:     "USD",
		ValorBRL:  45000.00,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Regra de negócio de alta velocidade: valores acima de R$ 10.000 exigem revisão
	if tx.ValorBRL > 10000.00 {
		tx.Status = "RETIDA_PARA_ANALISE"
		tx.Risco = "ALTO"
	} else {
		tx.Status = "APROVADA"
		tx.Risco = "BAIXO"
	}

	json.NewEncoder(w).Encode(tx)
}

func main() {
	http.HandleFunc("/api/validar", handlerAnalisar)

	fmt.Println("Microserviço Go Antifraude rodando em http://localhost:8080/api/validar")
	fmt.Println("Pressione Ctrl+C no terminal para encerrar.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}