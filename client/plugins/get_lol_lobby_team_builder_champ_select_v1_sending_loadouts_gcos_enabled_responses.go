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

// GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledReader is a Reader for the GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabled structure.
type GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK creates a GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK with default headers values
func NewGetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK() *GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK {
	return &GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK{}
}

/*GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK handles this case with default header values.

Successful response
*/
type GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK struct {
	Payload bool
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby-team-builder/champ-select/v1/sending-loadouts-gcos-enabled][%d] getLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SendingLoadoutsGcosEnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
