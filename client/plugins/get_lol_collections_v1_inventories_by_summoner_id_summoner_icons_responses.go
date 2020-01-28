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

// GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsReader is a Reader for the GetLolCollectionsV1InventoriesBySummonerIDSummonerIcons structure.
type GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK creates a GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK with default headers values
func NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK() *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK {
	return &GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK{}
}

/*GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK handles this case with default header values.

Successful response
*/
type GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK struct {
	Payload *models.LolCollectionsCollectionsSummonerIcons
}

func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK) Error() string {
	return fmt.Sprintf("[GET /lol-collections/v1/inventories/{summonerId}/summoner-icons][%d] getLolCollectionsV1InventoriesBySummonerIdSummonerIconsOK  %+v", 200, o.Payload)
}

func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK) GetPayload() *models.LolCollectionsCollectionsSummonerIcons {
	return o.Payload
}

func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCollectionsCollectionsSummonerIcons)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
