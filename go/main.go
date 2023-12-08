package main

import (
	"os"
	"path/filepath"

	"github.com/moisesfm67/curso-fullcycle-desafio/go/internal/people/entity"
)

func main() {
	//para verificar se foi passado o arquivo de origem
	if len(os.Args) < 2 {
		println("Arquivo não fornecido")
		os.Exit(1)
	}

	expectedFileName := "origin-file.csv"
	actualFileName := filepath.Base(os.Args[1])

	if actualFileName != expectedFileName {
		println("Nome do arquivo incorreto")
		println("Use: go run main.go ./origin-file.csv")
		os.Exit(1)
	}

	peoples, err := entity.ReadFileCSV(os.Args[1])
	if err != nil {
		panic(err)
	}

	entity.OrderByName(peoples)
	err = entity.WriteFileCSV("ordered_by_name.csv", peoples)
	if err != nil {
		panic(err)
	}

	entity.OrderByAge(peoples)
	err = entity.WriteFileCSV("ordered_by_age.csv", peoples)
	if err != nil {
		panic(err)
	}

	entity.OrderByNameAndAge(peoples)
	err = entity.WriteFileCSV("ordered_by_name_and_age.csv", peoples)
	if err != nil {
		panic(err)
	}

	println("Arquivos gerados com sucesso!")
}
