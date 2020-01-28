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

// GetLolClubsPublicV1ClubsPublicReader is a Reader for the GetLolClubsPublicV1ClubsPublic structure.
type GetLolClubsPublicV1ClubsPublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClubsPublicV1ClubsPublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClubsPublicV1ClubsPublicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClubsPublicV1ClubsPublicOK creates a GetLolClubsPublicV1ClubsPublicOK with default headers values
func NewGetLolClubsPublicV1ClubsPublicOK() *GetLolClubsPublicV1ClubsPublicOK {
	return &GetLolClubsPublicV1ClubsPublicOK{}
}

/*GetLolClubsPublicV1ClubsPublicOK handles this case with default header values.

Successful response
*/
type GetLolClubsPublicV1ClubsPublicOK struct {
	Payload []*models.LolClubsPublicClubsPublicData
}

func (o *GetLolClubsPublicV1ClubsPublicOK) Error() string {
	return fmt.Sprintf("[GET /lol-clubs-public/v1/clubs/public][%d] getLolClubsPublicV1ClubsPublicOK  %+v", 200, o.Payload)
}

func (o *GetLolClubsPublicV1ClubsPublicOK) GetPayload() []*models.LolClubsPublicClubsPublicData {
	return o.Payload
}

func (o *GetLolClubsPublicV1ClubsPublicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
