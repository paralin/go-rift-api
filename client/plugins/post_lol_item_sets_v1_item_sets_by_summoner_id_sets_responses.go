// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolItemSetsV1ItemSetsBySummonerIDSetsReader is a Reader for the PostLolItemSetsV1ItemSetsBySummonerIDSets structure.
type PostLolItemSetsV1ItemSetsBySummonerIDSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolItemSetsV1ItemSetsBySummonerIDSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent creates a PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent with default headers values
func NewPostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent() *PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent {
	return &PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent{}
}

/*PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent handles this case with default header values.

No content
*/
type PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent struct {
}

func (o *PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-item-sets/v1/item-sets/{summonerId}/sets][%d] postLolItemSetsV1ItemSetsBySummonerIdSetsNoContent ", 204)
}

func (o *PostLolItemSetsV1ItemSetsBySummonerIDSetsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}