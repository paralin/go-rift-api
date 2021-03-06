// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLolEndOfGameV1ReportedPlayersReader is a Reader for the GetLolEndOfGameV1ReportedPlayers structure.
type GetLolEndOfGameV1ReportedPlayersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolEndOfGameV1ReportedPlayersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolEndOfGameV1ReportedPlayersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolEndOfGameV1ReportedPlayersOK creates a GetLolEndOfGameV1ReportedPlayersOK with default headers values
func NewGetLolEndOfGameV1ReportedPlayersOK() *GetLolEndOfGameV1ReportedPlayersOK {
	return &GetLolEndOfGameV1ReportedPlayersOK{}
}

/*GetLolEndOfGameV1ReportedPlayersOK handles this case with default header values.

Successful response
*/
type GetLolEndOfGameV1ReportedPlayersOK struct {
	Payload []int64
}

func (o *GetLolEndOfGameV1ReportedPlayersOK) Error() string {
	return fmt.Sprintf("[GET /lol-end-of-game/v1/reported-players][%d] getLolEndOfGameV1ReportedPlayersOK  %+v", 200, o.Payload)
}

func (o *GetLolEndOfGameV1ReportedPlayersOK) GetPayload() []int64 {
	return o.Payload
}

func (o *GetLolEndOfGameV1ReportedPlayersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
