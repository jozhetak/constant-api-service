package serializers

import "github.com/ninjadotorg/constant-api-service/models"

type ReserseParties struct {
	ListReserseParties []ReserseParty `json:"ListReserseParties"`
}

func NewReserseParties(data []models.ReserseParty) *ReserseParties {
	result := &ReserseParties{
		ListReserseParties: []ReserseParty{},
	}
	for _, d := range data {
		result.ListReserseParties = append(result.ListReserseParties, *NewReserseParty(&d))
	}
	return result
}

type ReserseParty struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

func NewReserseParty(data *models.ReserseParty) *ReserseParty {
	result := &ReserseParty{
		ID:   data.ID,
		Name: data.Name,
	}
	return result
}
