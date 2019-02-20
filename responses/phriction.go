package responses

import "github.com/hexerei/gonduit/entities"

type PhrictionInfoResponse entities.PhrictionInfoDocument

type PhrictionDocumentSearchResponse struct {
	Data   []entities.PhrictionDocument     `json:"data"`
	Maps   map[string]string                `json:"maps"`
	Query  map[string]string                `json:"query"`
	Cursor entities.PhrictionDocumentCursor `json:"cursor"`
	Status string                           `json:"status"`
}
