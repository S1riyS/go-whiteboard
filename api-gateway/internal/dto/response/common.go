package response

import "github.com/google/uuid"

type DeleteResponse struct {
	ID uuid.UUID `json:"id"`
}
