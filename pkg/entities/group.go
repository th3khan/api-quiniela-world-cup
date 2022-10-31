package entities

import "github.com/th3khan/api-quiniela-world-cup/app/models"

type (
	GroupBase struct {
		Name   string `json:"name" validate:"required"`
		Active bool   `json:"active" validate:"required"`
	}

	GroupRequest struct {
		GroupBase
	}

	GroupResponse struct {
		ID uint `json:"id"`
		GroupBase
	}

	GroupResponseWithPagination struct {
		Groups []GroupResponse `json:"groups"`
		Pagination
	}
)

func CreateGroupResponseWithPagination(page int, perPage int, total int, groups []models.Group) GroupResponseWithPagination {
	var groupsResponse []GroupResponse
	for _, group := range groups {
		groupsResponse = append(groupsResponse, CreateGroupresponse(group))
	}

	return GroupResponseWithPagination{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
		Groups: groupsResponse,
	}
}

func CreateGroupresponse(group models.Group) GroupResponse {
	return GroupResponse{
		ID: group.ID,
		GroupBase: GroupBase{
			Name:   group.Name,
			Active: group.Active,
		},
	}
}
