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
