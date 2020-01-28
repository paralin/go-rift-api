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

// GetLolTftV1TftBattlepassReader is a Reader for the GetLolTftV1TftBattlepass structure.
type GetLolTftV1TftBattlepassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolTftV1TftBattlepassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolTftV1TftBattlepassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolTftV1TftBattlepassOK creates a GetLolTftV1TftBattlepassOK with default headers values
func NewGetLolTftV1TftBattlepassOK() *GetLolTftV1TftBattlepassOK {
	return &GetLolTftV1TftBattlepassOK{}
}

/*GetLolTftV1TftBattlepassOK handles this case with default header values.

Successful response
*/
type GetLolTftV1TftBattlepassOK struct {
	Payload *models.LolMissionsTftBattlepass
}

func (o *GetLolTftV1TftBattlepassOK) Error() string {
	return fmt.Sprintf("[GET /lol-tft/v1/tft/battlepass][%d] getLolTftV1TftBattlepassOK  %+v", 200, o.Payload)
}

func (o *GetLolTftV1TftBattlepassOK) GetPayload() *models.LolMissionsTftBattlepass {
	return o.Payload
}

func (o *GetLolTftV1TftBattlepassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolMissionsTftBattlepass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
