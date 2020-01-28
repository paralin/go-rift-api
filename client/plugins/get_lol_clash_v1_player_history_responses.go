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

// GetLolClashV1PlayerHistoryReader is a Reader for the GetLolClashV1PlayerHistory structure.
type GetLolClashV1PlayerHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1PlayerHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1PlayerHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1PlayerHistoryOK creates a GetLolClashV1PlayerHistoryOK with default headers values
func NewGetLolClashV1PlayerHistoryOK() *GetLolClashV1PlayerHistoryOK {
	return &GetLolClashV1PlayerHistoryOK{}
}

/*GetLolClashV1PlayerHistoryOK handles this case with default header values.

Successful response
*/
type GetLolClashV1PlayerHistoryOK struct {
	Payload []*models.LolClashRosterStats
}

func (o *GetLolClashV1PlayerHistoryOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/player/history][%d] getLolClashV1PlayerHistoryOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1PlayerHistoryOK) GetPayload() []*models.LolClashRosterStats {
	return o.Payload
}

func (o *GetLolClashV1PlayerHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
