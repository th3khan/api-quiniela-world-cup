package seeders

import (
	"log"

	"github.com/th3khan/api-quiniela-world-cup/app/services/phase"
)

func (s *Seed) CreatePhases() error {
	phases := []string{"Fase de grupos", "Octavos de final", "Cuartos de final", "Semifinal", "Final"}

	for _, r := range phases {
		if _, err := phase.CreatePhase(r); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
