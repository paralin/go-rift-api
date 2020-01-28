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

// GetLolRankedV1RankedStatsReader is a Reader for the GetLolRankedV1RankedStats structure.
type GetLolRankedV1RankedStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRankedV1RankedStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRankedV1RankedStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRankedV1RankedStatsOK creates a GetLolRankedV1RankedStatsOK with default headers values
func NewGetLolRankedV1RankedStatsOK() *GetLolRankedV1RankedStatsOK {
	return &GetLolRankedV1RankedStatsOK{}
}

/*GetLolRankedV1RankedStatsOK handles this case with default header values.

Successful response
*/
type GetLolRankedV1RankedStatsOK struct {
	Payload map[string]models.LolRankedRankedStats
}

func (o *GetLolRankedV1RankedStatsOK) Error() string {
	return fmt.Sprintf("[GET /lol-ranked/v1/ranked-stats][%d] getLolRankedV1RankedStatsOK  %+v", 200, o.Payload)
}

func (o *GetLolRankedV1RankedStatsOK) GetPayload() map[string]models.LolRankedRankedStats {
	return o.Payload
}

func (o *GetLolRankedV1RankedStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}