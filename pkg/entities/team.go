package entities

import (
	"time"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
)

type (
	TeamBase struct {
		Name   string `json:"name"`
		Active bool   `json:"active"`
		Logo   string `json:"logo"`
	}

	TeamResponse struct {
		TeamBase
		CreatedAt time.Time `json:"created_at"`
	}

	TeamResponseWithPagination struct {
		Teams []TeamResponse `json:"teams"`
		Pagination
	}
)

func CreateTeamResponse(team models.Team) TeamResponse {
	return TeamResponse{
		CreatedAt: team.CreatedAt,
		TeamBase: TeamBase{
			Name:   team.Name,
			Active: team.Active,
			Logo:   team.Logo,
		},
	}
}

func CreateTeamResponseWithpagination(teams []models.Team, page int, perPage int, total int) TeamResponseWithPagination {
	var teamsResponse []TeamResponse

	for _, team := range teams {
		teamsResponse = append(teamsResponse, CreateTeamResponse(team))
	}

	return TeamResponseWithPagination{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
	}
}
