// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// NewPatchLolChampSelectLegacyV1SessionMySelectionParams creates a new PatchLolChampSelectLegacyV1SessionMySelectionParams object
// with the default values initialized.
func NewPatchLolChampSelectLegacyV1SessionMySelectionParams() *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	var ()
	return &PatchLolChampSelectLegacyV1SessionMySelectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithTimeout creates a new PatchLolChampSelectLegacyV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithTimeout(timeout time.Duration) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	var ()
	return &PatchLolChampSelectLegacyV1SessionMySelectionParams{

		timeout: timeout,
	}
}

// NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithContext creates a new PatchLolChampSelectLegacyV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithContext(ctx context.Context) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	var ()
	return &PatchLolChampSelectLegacyV1SessionMySelectionParams{

		Context: ctx,
	}
}

// NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithHTTPClient creates a new PatchLolChampSelectLegacyV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLolChampSelectLegacyV1SessionMySelectionParamsWithHTTPClient(client *http.Client) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	var ()
	return &PatchLolChampSelectLegacyV1SessionMySelectionParams{
		HTTPClient: client,
	}
}

/*PatchLolChampSelectLegacyV1SessionMySelectionParams contains all the parameters to send to the API endpoint
for the patch lol champ select legacy v1 session my selection operation typically these are written to a http.Request
*/
type PatchLolChampSelectLegacyV1SessionMySelectionParams struct {

	/*Selection*/
	Selection *models.LolChampSelectLegacyChampSelectMySelection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) WithTimeout(timeout time.Duration) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) WithContext(ctx context.Context) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) WithHTTPClient(client *http.Client) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelection adds the selection to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) WithSelection(selection *models.LolChampSelectLegacyChampSelectMySelection) *PatchLolChampSelectLegacyV1SessionMySelectionParams {
	o.SetSelection(selection)
	return o
}

// SetSelection adds the selection to the patch lol champ select legacy v1 session my selection params
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) SetSelection(selection *models.LolChampSelectLegacyChampSelectMySelection) {
	o.Selection = selection
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLolChampSelectLegacyV1SessionMySelectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Selection != nil {
		if err := r.SetBodyParam(o.Selection); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}