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

// NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams creates a new PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams object
// with the default values initialized.
func NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams() *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	var ()
	return &PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithTimeout creates a new PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithTimeout(timeout time.Duration) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	var ()
	return &PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams{

		timeout: timeout,
	}
}

// NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithContext creates a new PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithContext(ctx context.Context) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	var ()
	return &PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams{

		Context: ctx,
	}
}

// NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithHTTPClient creates a new PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParamsWithHTTPClient(client *http.Client) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	var ()
	return &PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams{
		HTTPClient: client,
	}
}

/*PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams contains all the parameters to send to the API endpoint
for the patch lol lobby team builder champ select v1 session my selection operation typically these are written to a http.Request
*/
type PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams struct {

	/*Selection*/
	Selection *models.LolLobbyTeamBuilderChampSelectMySelection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) WithTimeout(timeout time.Duration) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) WithContext(ctx context.Context) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) WithHTTPClient(client *http.Client) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelection adds the selection to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) WithSelection(selection *models.LolLobbyTeamBuilderChampSelectMySelection) *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams {
	o.SetSelection(selection)
	return o
}

// SetSelection adds the selection to the patch lol lobby team builder champ select v1 session my selection params
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) SetSelection(selection *models.LolLobbyTeamBuilderChampSelectMySelection) {
	o.Selection = selection
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLolLobbyTeamBuilderChampSelectV1SessionMySelectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
