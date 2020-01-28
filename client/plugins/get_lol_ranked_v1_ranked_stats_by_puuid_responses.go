// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// GetLolRankedV1RankedStatsByPuuidReader is a Reader for the GetLolRankedV1RankedStatsByPuuid structure.
type GetLolRankedV1RankedStatsByPuuidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRankedV1RankedStatsByPuuidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRankedV1RankedStatsByPuuidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRankedV1RankedStatsByPuuidOK creates a GetLolRankedV1RankedStatsByPuuidOK with default headers values
func NewGetLolRankedV1RankedStatsByPuuidOK() *GetLolRankedV1RankedStatsByPuuidOK {
	return &GetLolRankedV1RankedStatsByPuuidOK{}
}

/*GetLolRankedV1RankedStatsByPuuidOK handles this case with default header values.

Successful response
*/
type GetLolRankedV1RankedStatsByPuuidOK struct {
	Payload *models.LolRankedRankedStats
}

func (o *GetLolRankedV1RankedStatsByPuuidOK) Error() string {
	return fmt.Sprintf("[GET /lol-ranked/v1/ranked-stats/{puuid}][%d] getLolRankedV1RankedStatsByPuuidOK  %+v", 200, o.Payload)
}

func (o *GetLolRankedV1RankedStatsByPuuidOK) GetPayload() *models.LolRankedRankedStats {
	return o.Payload
}

func (o *GetLolRankedV1RankedStatsByPuuidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRankedRankedStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
