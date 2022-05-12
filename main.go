package main

import (
	"fmt"

	"github.com/jheissonn/go-rest-api/database"
	"github.com/jheissonn/go-rest-api/models"
	"github.com/jheissonn/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go porta padrão: 8000")
	routes.HandleResquest()
}
