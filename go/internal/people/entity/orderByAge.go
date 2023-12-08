package entity

import (
	"sort"

	"github.com/moisesfm67/curso-fullcycle-desafio/go/internal/people/dto"
)

func OrderByAge(peoples []dto.People) []dto.People {
	sort.Slice(peoples, func(i, j int) bool {
		return peoples[i].Age < peoples[j].Age
	})

	return peoples
}
