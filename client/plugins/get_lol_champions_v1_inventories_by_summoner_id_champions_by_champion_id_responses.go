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

// GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDReader is a Reader for the GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionID structure.
type GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK creates a GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK with default headers values
func NewGetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK() *GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK {
	return &GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK{}
}

/*GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK handles this case with default header values.

Successful response
*/
type GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK struct {
	Payload *models.LolChampionsCollectionsChampion
}

func (o *GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-champions/v1/inventories/{summonerId}/champions/{championId}][%d] getLolChampionsV1InventoriesBySummonerIdChampionsByChampionIdOK  %+v", 200, o.Payload)
}

func (o *GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK) GetPayload() *models.LolChampionsCollectionsChampion {
	return o.Payload
}

func (o *GetLolChampionsV1InventoriesBySummonerIDChampionsByChampionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChampionsCollectionsChampion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}