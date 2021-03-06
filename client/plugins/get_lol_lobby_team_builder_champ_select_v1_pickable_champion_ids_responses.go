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

// GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsReader is a Reader for the GetLolLobbyTeamBuilderChampSelectV1PickableChampionIds structure.
type GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK creates a GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK with default headers values
func NewGetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK() *GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK {
	return &GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK{}
}

/*GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK handles this case with default header values.

Successful response
*/
type GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK struct {
	Payload []int32
}

func (o *GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby-team-builder/champ-select/v1/pickable-champion-ids][%d] getLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetLolLobbyTeamBuilderChampSelectV1PickableChampionIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
