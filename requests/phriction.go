package requests

import "github.com/hexerei/gonduit/constants"

type PhrictionInfoRequest struct {
	Slug string `json:"slug"`
	Request
}

type PhrictionDocumentSearchRequest struct {
	QueryKey    string                                 `json:"queryKey"`
	Constraints map[string][]string                    `json:"constraints"`
	Attachments map[string]bool                        `json:"attachments"`
	Order       constants.PhrictionDocumentSearchOrder `json:"order"`
	Before      string                                 `json:"before"`
	After       string                                 `json:"after"`
	Limit       uint64                                 `json:"limit"`
	Request                                            // Includes __conduit__ field needed for authentication.
}
