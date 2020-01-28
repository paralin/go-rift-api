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

// GetLolChampSelectLegacyV1BannableChampionIdsReader is a Reader for the GetLolChampSelectLegacyV1BannableChampionIds structure.
type GetLolChampSelectLegacyV1BannableChampionIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampSelectLegacyV1BannableChampionIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampSelectLegacyV1BannableChampionIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampSelectLegacyV1BannableChampionIdsOK creates a GetLolChampSelectLegacyV1BannableChampionIdsOK with default headers values
func NewGetLolChampSelectLegacyV1BannableChampionIdsOK() *GetLolChampSelectLegacyV1BannableChampionIdsOK {
	return &GetLolChampSelectLegacyV1BannableChampionIdsOK{}
}

/*GetLolChampSelectLegacyV1BannableChampionIdsOK handles this case with default header values.

Successful response
*/
type GetLolChampSelectLegacyV1BannableChampionIdsOK struct {
	Payload []int32
}

func (o *GetLolChampSelectLegacyV1BannableChampionIdsOK) Error() string {
	return fmt.Sprintf("[GET /lol-champ-select-legacy/v1/bannable-champion-ids][%d] getLolChampSelectLegacyV1BannableChampionIdsOK  %+v", 200, o.Payload)
}

func (o *GetLolChampSelectLegacyV1BannableChampionIdsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetLolChampSelectLegacyV1BannableChampionIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}