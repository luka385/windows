package application

import (
	"unicode"

	"github.com/luka385/hexa-test4/domain"
)

type verySer struct{}

func (vs *verySer) VerificarPerson(p *domain.Person) domain.Person {

	if veryMayus(p.Name) {
		return domain.Person{
			ID:    p.ID,
			Name:  converMayus(p.Name),
			Email: p.Email,
		}
	}
	return *p
}

func veryMayus(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func converMayus(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
