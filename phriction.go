package gonduit

import (
	"github.com/hexerei/gonduit/requests"
	"github.com/hexerei/gonduit/responses"
)

// PhrictionInfo performs a call to phriction.info
func (c *Conn) PhrictionInfo(
	req requests.PhrictionInfoRequest,
) (*responses.PhrictionInfoResponse, error) {
	var res responses.PhrictionInfoResponse

	if err := c.Call("phriction.info", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// PhrictionDocumentSearch performs a call to phriction.document.search
func (c *Conn) PhrictionDocumentSearch(
	req requests.PhrictionDocumentSearchRequest,
) (*responses.PhrictionDocumentSearchResponse, error) {
	var res responses.PhrictionDocumentSearchResponse

	if err := c.Call("phriction.document.search", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
