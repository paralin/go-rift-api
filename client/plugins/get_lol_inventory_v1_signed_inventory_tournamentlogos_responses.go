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

// GetLolInventoryV1SignedInventoryTournamentlogosReader is a Reader for the GetLolInventoryV1SignedInventoryTournamentlogos structure.
type GetLolInventoryV1SignedInventoryTournamentlogosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolInventoryV1SignedInventoryTournamentlogosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolInventoryV1SignedInventoryTournamentlogosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolInventoryV1SignedInventoryTournamentlogosOK creates a GetLolInventoryV1SignedInventoryTournamentlogosOK with default headers values
func NewGetLolInventoryV1SignedInventoryTournamentlogosOK() *GetLolInventoryV1SignedInventoryTournamentlogosOK {
	return &GetLolInventoryV1SignedInventoryTournamentlogosOK{}
}

/*GetLolInventoryV1SignedInventoryTournamentlogosOK handles this case with default header values.

Successful response
*/
type GetLolInventoryV1SignedInventoryTournamentlogosOK struct {
	Payload map[string]string
}

func (o *GetLolInventoryV1SignedInventoryTournamentlogosOK) Error() string {
	return fmt.Sprintf("[GET /lol-inventory/v1/signedInventory/tournamentlogos][%d] getLolInventoryV1SignedInventoryTournamentlogosOK  %+v", 200, o.Payload)
}

func (o *GetLolInventoryV1SignedInventoryTournamentlogosOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GetLolInventoryV1SignedInventoryTournamentlogosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
