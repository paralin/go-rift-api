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

// GetLolChampSelectV1ChampSfxNotificationsReader is a Reader for the GetLolChampSelectV1ChampSfxNotifications structure.
type GetLolChampSelectV1ChampSfxNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampSelectV1ChampSfxNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampSelectV1ChampSfxNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampSelectV1ChampSfxNotificationsOK creates a GetLolChampSelectV1ChampSfxNotificationsOK with default headers values
func NewGetLolChampSelectV1ChampSfxNotificationsOK() *GetLolChampSelectV1ChampSfxNotificationsOK {
	return &GetLolChampSelectV1ChampSfxNotificationsOK{}
}

/*GetLolChampSelectV1ChampSfxNotificationsOK handles this case with default header values.

Successful response
*/
type GetLolChampSelectV1ChampSfxNotificationsOK struct {
	Payload []*models.LolChampSelectChampionSfxNotification
}

func (o *GetLolChampSelectV1ChampSfxNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /lol-champ-select/v1/champ-sfx-notifications][%d] getLolChampSelectV1ChampSfxNotificationsOK  %+v", 200, o.Payload)
}

func (o *GetLolChampSelectV1ChampSfxNotificationsOK) GetPayload() []*models.LolChampSelectChampionSfxNotification {
	return o.Payload
}

func (o *GetLolChampSelectV1ChampSfxNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}