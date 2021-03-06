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
)

// NewPutLolLobbyV2LobbyPartyTypeParams creates a new PutLolLobbyV2LobbyPartyTypeParams object
// with the default values initialized.
func NewPutLolLobbyV2LobbyPartyTypeParams() *PutLolLobbyV2LobbyPartyTypeParams {
	var ()
	return &PutLolLobbyV2LobbyPartyTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLolLobbyV2LobbyPartyTypeParamsWithTimeout creates a new PutLolLobbyV2LobbyPartyTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLolLobbyV2LobbyPartyTypeParamsWithTimeout(timeout time.Duration) *PutLolLobbyV2LobbyPartyTypeParams {
	var ()
	return &PutLolLobbyV2LobbyPartyTypeParams{

		timeout: timeout,
	}
}

// NewPutLolLobbyV2LobbyPartyTypeParamsWithContext creates a new PutLolLobbyV2LobbyPartyTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLolLobbyV2LobbyPartyTypeParamsWithContext(ctx context.Context) *PutLolLobbyV2LobbyPartyTypeParams {
	var ()
	return &PutLolLobbyV2LobbyPartyTypeParams{

		Context: ctx,
	}
}

// NewPutLolLobbyV2LobbyPartyTypeParamsWithHTTPClient creates a new PutLolLobbyV2LobbyPartyTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLolLobbyV2LobbyPartyTypeParamsWithHTTPClient(client *http.Client) *PutLolLobbyV2LobbyPartyTypeParams {
	var ()
	return &PutLolLobbyV2LobbyPartyTypeParams{
		HTTPClient: client,
	}
}

/*PutLolLobbyV2LobbyPartyTypeParams contains all the parameters to send to the API endpoint
for the put lol lobby v2 lobby party type operation typically these are written to a http.Request
*/
type PutLolLobbyV2LobbyPartyTypeParams struct {

	/*PartyType*/
	PartyType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) WithTimeout(timeout time.Duration) *PutLolLobbyV2LobbyPartyTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) WithContext(ctx context.Context) *PutLolLobbyV2LobbyPartyTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) WithHTTPClient(client *http.Client) *PutLolLobbyV2LobbyPartyTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPartyType adds the partyType to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) WithPartyType(partyType string) *PutLolLobbyV2LobbyPartyTypeParams {
	o.SetPartyType(partyType)
	return o
}

// SetPartyType adds the partyType to the put lol lobby v2 lobby party type params
func (o *PutLolLobbyV2LobbyPartyTypeParams) SetPartyType(partyType string) {
	o.PartyType = partyType
}

// WriteToRequest writes these params to a swagger request
func (o *PutLolLobbyV2LobbyPartyTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PartyType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
