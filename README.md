# Go Microservice: Motor Antifraude em Tempo Real

Microserviço de alta performance desenvolvido em Golang (Go) utilizando apenas a biblioteca padrão (`net/http` e `encoding/json`) para validação instantânea de transações financeiras e detecção de valores atípicos.

## Tecnologias e Diferenciais
* **Golang (Go):** Concorrência nativa, baixo consumo de memória e resposta em microssegundos.
* **API REST / JSON:** Endpoint `/api/validar` pronto para integração com sistemas legados ou frontends modernos.

## Como executar
1. Certifique-se de ter o Go instalado (`go version`).
2. No terminal, execute: `go run main.go`
3. Acesse no navegador ou via cURL: `http://localhost:8080/api/validar`