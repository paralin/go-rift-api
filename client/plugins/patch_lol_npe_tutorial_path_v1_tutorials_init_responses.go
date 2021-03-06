// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchLolNpeTutorialPathV1TutorialsInitReader is a Reader for the PatchLolNpeTutorialPathV1TutorialsInit structure.
type PatchLolNpeTutorialPathV1TutorialsInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolNpeTutorialPathV1TutorialsInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchLolNpeTutorialPathV1TutorialsInitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolNpeTutorialPathV1TutorialsInitNoContent creates a PatchLolNpeTutorialPathV1TutorialsInitNoContent with default headers values
func NewPatchLolNpeTutorialPathV1TutorialsInitNoContent() *PatchLolNpeTutorialPathV1TutorialsInitNoContent {
	return &PatchLolNpeTutorialPathV1TutorialsInitNoContent{}
}

/*PatchLolNpeTutorialPathV1TutorialsInitNoContent handles this case with default header values.

No content
*/
type PatchLolNpeTutorialPathV1TutorialsInitNoContent struct {
}

func (o *PatchLolNpeTutorialPathV1TutorialsInitNoContent) Error() string {
	return fmt.Sprintf("[PATCH /lol-npe-tutorial-path/v1/tutorials/init][%d] patchLolNpeTutorialPathV1TutorialsInitNoContent ", 204)
}

func (o *PatchLolNpeTutorialPathV1TutorialsInitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
