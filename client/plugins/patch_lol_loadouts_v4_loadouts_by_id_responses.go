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

// PatchLolLoadoutsV4LoadoutsByIDReader is a Reader for the PatchLolLoadoutsV4LoadoutsByID structure.
type PatchLolLoadoutsV4LoadoutsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolLoadoutsV4LoadoutsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLolLoadoutsV4LoadoutsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolLoadoutsV4LoadoutsByIDOK creates a PatchLolLoadoutsV4LoadoutsByIDOK with default headers values
func NewPatchLolLoadoutsV4LoadoutsByIDOK() *PatchLolLoadoutsV4LoadoutsByIDOK {
	return &PatchLolLoadoutsV4LoadoutsByIDOK{}
}

/*PatchLolLoadoutsV4LoadoutsByIDOK handles this case with default header values.

Successful response
*/
type PatchLolLoadoutsV4LoadoutsByIDOK struct {
	Payload *models.LolLoadoutsScopedLoadout
}

func (o *PatchLolLoadoutsV4LoadoutsByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /lol-loadouts/v4/loadouts/{id}][%d] patchLolLoadoutsV4LoadoutsByIdOK  %+v", 200, o.Payload)
}

func (o *PatchLolLoadoutsV4LoadoutsByIDOK) GetPayload() *models.LolLoadoutsScopedLoadout {
	return o.Payload
}

func (o *PatchLolLoadoutsV4LoadoutsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLoadoutsScopedLoadout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}