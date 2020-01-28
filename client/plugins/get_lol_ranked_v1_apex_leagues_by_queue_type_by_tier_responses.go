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

// GetLolRankedV1ApexLeaguesByQueueTypeByTierReader is a Reader for the GetLolRankedV1ApexLeaguesByQueueTypeByTier structure.
type GetLolRankedV1ApexLeaguesByQueueTypeByTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRankedV1ApexLeaguesByQueueTypeByTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRankedV1ApexLeaguesByQueueTypeByTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRankedV1ApexLeaguesByQueueTypeByTierOK creates a GetLolRankedV1ApexLeaguesByQueueTypeByTierOK with default headers values
func NewGetLolRankedV1ApexLeaguesByQueueTypeByTierOK() *GetLolRankedV1ApexLeaguesByQueueTypeByTierOK {
	return &GetLolRankedV1ApexLeaguesByQueueTypeByTierOK{}
}

/*GetLolRankedV1ApexLeaguesByQueueTypeByTierOK handles this case with default header values.

Successful response
*/
type GetLolRankedV1ApexLeaguesByQueueTypeByTierOK struct {
	Payload *models.LolRankedLeagueLadderInfo
}

func (o *GetLolRankedV1ApexLeaguesByQueueTypeByTierOK) Error() string {
	return fmt.Sprintf("[GET /lol-ranked/v1/apex-leagues/{queueType}/{tier}][%d] getLolRankedV1ApexLeaguesByQueueTypeByTierOK  %+v", 200, o.Payload)
}

func (o *GetLolRankedV1ApexLeaguesByQueueTypeByTierOK) GetPayload() *models.LolRankedLeagueLadderInfo {
	return o.Payload
}

func (o *GetLolRankedV1ApexLeaguesByQueueTypeByTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRankedLeagueLadderInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}