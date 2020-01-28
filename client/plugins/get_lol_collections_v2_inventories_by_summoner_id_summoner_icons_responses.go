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

// GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsReader is a Reader for the GetLolCollectionsV2InventoriesBySummonerIDSummonerIcons structure.
type GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK creates a GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK with default headers values
func NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK() *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK {
	return &GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK{}
}

/*GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK handles this case with default header values.

Successful response
*/
type GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK struct {
	Payload *models.LolCollectionsCollectionsSummonerIcons
}

func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK) Error() string {
	return fmt.Sprintf("[GET /lol-collections/v2/inventories/{summonerId}/summoner-icons][%d] getLolCollectionsV2InventoriesBySummonerIdSummonerIconsOK  %+v", 200, o.Payload)
}

func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK) GetPayload() *models.LolCollectionsCollectionsSummonerIcons {
	return o.Payload
}

func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCollectionsCollectionsSummonerIcons)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
