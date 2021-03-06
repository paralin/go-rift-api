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

// GetLolStoreV1GiftablefriendsReader is a Reader for the GetLolStoreV1Giftablefriends structure.
type GetLolStoreV1GiftablefriendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolStoreV1GiftablefriendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolStoreV1GiftablefriendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolStoreV1GiftablefriendsOK creates a GetLolStoreV1GiftablefriendsOK with default headers values
func NewGetLolStoreV1GiftablefriendsOK() *GetLolStoreV1GiftablefriendsOK {
	return &GetLolStoreV1GiftablefriendsOK{}
}

/*GetLolStoreV1GiftablefriendsOK handles this case with default header values.

Successful response
*/
type GetLolStoreV1GiftablefriendsOK struct {
	Payload []*models.LolStoreGiftingFriend
}

func (o *GetLolStoreV1GiftablefriendsOK) Error() string {
	return fmt.Sprintf("[GET /lol-store/v1/giftablefriends][%d] getLolStoreV1GiftablefriendsOK  %+v", 200, o.Payload)
}

func (o *GetLolStoreV1GiftablefriendsOK) GetPayload() []*models.LolStoreGiftingFriend {
	return o.Payload
}

func (o *GetLolStoreV1GiftablefriendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
