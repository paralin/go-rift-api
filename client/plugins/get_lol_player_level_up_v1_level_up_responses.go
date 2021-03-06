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

// GetLolPlayerLevelUpV1LevelUpReader is a Reader for the GetLolPlayerLevelUpV1LevelUp structure.
type GetLolPlayerLevelUpV1LevelUpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPlayerLevelUpV1LevelUpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPlayerLevelUpV1LevelUpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPlayerLevelUpV1LevelUpOK creates a GetLolPlayerLevelUpV1LevelUpOK with default headers values
func NewGetLolPlayerLevelUpV1LevelUpOK() *GetLolPlayerLevelUpV1LevelUpOK {
	return &GetLolPlayerLevelUpV1LevelUpOK{}
}

/*GetLolPlayerLevelUpV1LevelUpOK handles this case with default header values.

Successful response
*/
type GetLolPlayerLevelUpV1LevelUpOK struct {
	Payload *models.LolPlayerLevelUpPlayerLevelUpEvent
}

func (o *GetLolPlayerLevelUpV1LevelUpOK) Error() string {
	return fmt.Sprintf("[GET /lol-player-level-up/v1/level-up][%d] getLolPlayerLevelUpV1LevelUpOK  %+v", 200, o.Payload)
}

func (o *GetLolPlayerLevelUpV1LevelUpOK) GetPayload() *models.LolPlayerLevelUpPlayerLevelUpEvent {
	return o.Payload
}

func (o *GetLolPlayerLevelUpV1LevelUpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPlayerLevelUpPlayerLevelUpEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
