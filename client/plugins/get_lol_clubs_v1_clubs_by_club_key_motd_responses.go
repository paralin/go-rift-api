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

// GetLolClubsV1ClubsByClubKeyMotdReader is a Reader for the GetLolClubsV1ClubsByClubKeyMotd structure.
type GetLolClubsV1ClubsByClubKeyMotdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClubsV1ClubsByClubKeyMotdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClubsV1ClubsByClubKeyMotdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClubsV1ClubsByClubKeyMotdOK creates a GetLolClubsV1ClubsByClubKeyMotdOK with default headers values
func NewGetLolClubsV1ClubsByClubKeyMotdOK() *GetLolClubsV1ClubsByClubKeyMotdOK {
	return &GetLolClubsV1ClubsByClubKeyMotdOK{}
}

/*GetLolClubsV1ClubsByClubKeyMotdOK handles this case with default header values.

Successful response
*/
type GetLolClubsV1ClubsByClubKeyMotdOK struct {
	Payload string
}

func (o *GetLolClubsV1ClubsByClubKeyMotdOK) Error() string {
	return fmt.Sprintf("[GET /lol-clubs/v1/clubs/{clubKey}/motd][%d] getLolClubsV1ClubsByClubKeyMotdOK  %+v", 200, o.Payload)
}

func (o *GetLolClubsV1ClubsByClubKeyMotdOK) GetPayload() string {
	return o.Payload
}

func (o *GetLolClubsV1ClubsByClubKeyMotdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
