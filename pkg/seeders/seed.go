package seeders

import "fmt"

type Seeder interface {
	All()
	CreateRoles() error
}

type Seed struct {
	Name string
	Run  func() error
}

func NewSeeder() *Seed {
	return new(Seed)
}

func (s *Seed) All() {
	fmt.Println("aqui")
	seeders := []Seed{
		Seed{
			Name: "CreateRoles",
			Run: func() error {
				return s.CreateRoles()
			},
		},
	}

	for _, seed := range seeders {
		seed.Run()
	}
}
