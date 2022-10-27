package seeders

import (
	"fmt"
	"log"

	"github.com/th3khan/api-quiniela-world-cup/app/services/role"
)

func (s *Seed) CreateRoles() error {
	roles := []string{"Super Admin", "User"}

	fmt.Println(roles)

	for _, r := range roles {
		if err, _ := role.CreateRole(r); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
