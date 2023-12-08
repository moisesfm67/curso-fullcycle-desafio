package entity

import (
	"sort"
	"strings"

	"github.com/moisesfm67/curso-fullcycle-desafio/go/internal/people/dto"
)

func OrderByNameAndAge(peoples []dto.People) []dto.People {
	sort.Slice(peoples, func(i, j int) bool {
		nomeI := strings.ToLower(peoples[i].Name)
		nomeJ := strings.ToLower(peoples[j].Name)

		if nomeI == nomeJ {
			return peoples[i].Age < peoples[j].Age
		}

		return nomeI < nomeJ
	})

	return peoples
}
