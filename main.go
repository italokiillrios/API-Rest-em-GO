package main

import (
	"fmt"
	"golang4/database"
	"golang4/models"
	"golang4/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

}
