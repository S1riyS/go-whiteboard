package response

import "github.com/S1riyS/go-whiteboard/api-gateway/internal/dto/model"

type DrawResponse struct {
	ID   string            `json:"id"`
	Type model.ElementType `json:"type"`
	Data any               `json:"data"`
}
