package domain

import "strings"

type Person struct {
	Name string `json:"name"`
}

func (p *Person) ValidateName() {
	p.Name = strings.ToUpper(p.Name)
}
