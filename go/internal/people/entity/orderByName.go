package entity

import (
	"sort"
	"strings"

	"github.com/moisesfm67/curso-fullcycle-desafio/go/internal/people/dto"
)

func OrderByName(peoples []dto.People) []dto.People {
	sort.Slice(peoples, func(i, j int) bool {
		return strings.ToLower(peoples[i].Name) < strings.ToLower(peoples[j].Name)
	})

	return peoples
}
