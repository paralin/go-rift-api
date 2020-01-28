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

// GetLolMatchHistoryV1FriendMatchlistsByAccountIDReader is a Reader for the GetLolMatchHistoryV1FriendMatchlistsByAccountID structure.
type GetLolMatchHistoryV1FriendMatchlistsByAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolMatchHistoryV1FriendMatchlistsByAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolMatchHistoryV1FriendMatchlistsByAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolMatchHistoryV1FriendMatchlistsByAccountIDOK creates a GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK with default headers values
func NewGetLolMatchHistoryV1FriendMatchlistsByAccountIDOK() *GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK {
	return &GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK{}
}

/*GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK handles this case with default header values.

Successful response
*/
type GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK struct {
	Payload *models.LolMatchHistoryMatchHistoryList
}

func (o *GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-match-history/v1/friend-matchlists/{accountId}][%d] getLolMatchHistoryV1FriendMatchlistsByAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK) GetPayload() *models.LolMatchHistoryMatchHistoryList {
	return o.Payload
}

func (o *GetLolMatchHistoryV1FriendMatchlistsByAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolMatchHistoryMatchHistoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
