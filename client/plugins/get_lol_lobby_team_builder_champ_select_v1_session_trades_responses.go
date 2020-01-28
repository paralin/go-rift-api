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

// GetLolLobbyTeamBuilderChampSelectV1SessionTradesReader is a Reader for the GetLolLobbyTeamBuilderChampSelectV1SessionTrades structure.
type GetLolLobbyTeamBuilderChampSelectV1SessionTradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionTradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyTeamBuilderChampSelectV1SessionTradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SessionTradesOK creates a GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK with default headers values
func NewGetLolLobbyTeamBuilderChampSelectV1SessionTradesOK() *GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK {
	return &GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK{}
}

/*GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK handles this case with default header values.

Successful response
*/
type GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK struct {
	Payload []*models.LolLobbyTeamBuilderChampSelectTradeContract
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby-team-builder/champ-select/v1/session/trades][%d] getLolLobbyTeamBuilderChampSelectV1SessionTradesOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK) GetPayload() []*models.LolLobbyTeamBuilderChampSelectTradeContract {
	return o.Payload
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionTradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
