package entity

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/moisesfm67/curso-fullcycle-desafio/go/internal/people/dto"
)

func ReadFileCSV(fileName string) ([]dto.People, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	readerCSV := csv.NewReader(file)
	readerCSV.Comma = ','

	lines, err := readerCSV.ReadAll()
	if err != nil {
		return nil, err
	}

	var peoples []dto.People
	for _, line := range lines[1:] {
		age, _ := strconv.Atoi(line[1])
		score, _ := strconv.Atoi(line[2])

		peoples = append(peoples, dto.People{
			Name:  line[0],
			Age:   age,
			Score: score,
		})
	}

	return peoples, nil
}

func WriteFileCSV(fileName string, peoples []dto.People) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writerCSV := csv.NewWriter(file)
	defer writerCSV.Flush()

	writerCSV.Write([]string{"Name", "Age", "Score"})

	for _, people := range peoples {
		line := []string{people.Name, strconv.Itoa(people.Age), strconv.Itoa(people.Score)}
		writerCSV.Write(line)
	}

	return nil
}
