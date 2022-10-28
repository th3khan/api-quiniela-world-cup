package seeders

import "fmt"

type Seeder interface {
	All()
	CreateRoles() error
	CreateUsers() error
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
		Seed{
			Name: "CreateUsers",
			Run: func() error {
				return s.CreateUsers()
			},
		},
	}

	for _, seed := range seeders {
		seed.Run()
	}
}
