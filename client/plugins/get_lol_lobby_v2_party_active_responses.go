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

// GetLolLobbyV2PartyActiveReader is a Reader for the GetLolLobbyV2PartyActive structure.
type GetLolLobbyV2PartyActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV2PartyActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV2PartyActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV2PartyActiveOK creates a GetLolLobbyV2PartyActiveOK with default headers values
func NewGetLolLobbyV2PartyActiveOK() *GetLolLobbyV2PartyActiveOK {
	return &GetLolLobbyV2PartyActiveOK{}
}

/*GetLolLobbyV2PartyActiveOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV2PartyActiveOK struct {
	Payload bool
}

func (o *GetLolLobbyV2PartyActiveOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v2/party-active][%d] getLolLobbyV2PartyActiveOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV2PartyActiveOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolLobbyV2PartyActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
