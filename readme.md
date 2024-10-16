# XML para JSON Converter

![GitHub issues](https://img.shields.io/github/issues/rodlessa/xml_json_go) ![GitHub pull requests](https://img.shields.io/github/issues-pr/rodlessa/xml_json_go) ![GitHub stars](https://img.shields.io/github/stars/rodlessa/xml_json_go) ![GitHub forks](https://img.shields.io/github/forks/rodlessa/xml_json_go)

## Descrição

Este projeto é um conversor de XML para JSON, desenvolvido com um backend em Go e um frontend em Vue.js. A aplicação permite que os usuários insiram XML e converta-o em JSON de forma simples e rápida.

## Funcionalidades

- Conversão automática de XML para JSON.
- Interface simples e responsiva utilizando Vue.js.
- Backend robusto utilizando Gin para gerenciar requisições e conversões.

## Estrutura do Projeto
├── backend/ # Código-fonte do backend em Go  
└── main.go 
└── frontend/ # Código-fonte do frontend em Vue.js 
└── src/

## Tecnologias Utilizadas

- **Backend**: Go (Golang), Gin, goxml2json
- **Frontend**: Vue.js, Axios, Tailwind CSS

## Instalação

### Pré-requisitos

- [Go](https://golang.org/dl/) (versão 1.16 ou superior)
- [Node.js](https://nodejs.org/) (versão 14 ou superior)
- [npm](https://www.npmjs.com/get-npm) (geralmente instalado junto com o Node.js)

### Backend

1. Navegue até a pasta `backend`:

   ```bash
   cd backend
2. Instalar dependencias:
 ```go
 go mod tidy
```
3. Execute o servidor:
```go 
go run main.go 
```

### Front-end
Navegue até a pasta frontend:

````bash

cd frontend
````
Instale as dependências:
````
bash

npm install
````
Execute a aplicação:

````bash

npm run serve
````

Uso:

Acessar em http://localhost:8080