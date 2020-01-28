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

// GetLolAcsV2RecentlyPlayedChampionsByAccountIDReader is a Reader for the GetLolAcsV2RecentlyPlayedChampionsByAccountID structure.
type GetLolAcsV2RecentlyPlayedChampionsByAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolAcsV2RecentlyPlayedChampionsByAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolAcsV2RecentlyPlayedChampionsByAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolAcsV2RecentlyPlayedChampionsByAccountIDOK creates a GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK with default headers values
func NewGetLolAcsV2RecentlyPlayedChampionsByAccountIDOK() *GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK {
	return &GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK{}
}

/*GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK handles this case with default header values.

Successful response
*/
type GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK struct {
	Payload *models.LolAcsAcsChampionGamesCollection
}

func (o *GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-acs/v2/recently-played-champions/{accountId}][%d] getLolAcsV2RecentlyPlayedChampionsByAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK) GetPayload() *models.LolAcsAcsChampionGamesCollection {
	return o.Payload
}

func (o *GetLolAcsV2RecentlyPlayedChampionsByAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolAcsAcsChampionGamesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}