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

// NewGetLolLobbyV2PartyActiveParams creates a new GetLolLobbyV2PartyActiveParams object
// with the default values initialized.
func NewGetLolLobbyV2PartyActiveParams() *GetLolLobbyV2PartyActiveParams {

	return &GetLolLobbyV2PartyActiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLobbyV2PartyActiveParamsWithTimeout creates a new GetLolLobbyV2PartyActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLobbyV2PartyActiveParamsWithTimeout(timeout time.Duration) *GetLolLobbyV2PartyActiveParams {

	return &GetLolLobbyV2PartyActiveParams{

		timeout: timeout,
	}
}

// NewGetLolLobbyV2PartyActiveParamsWithContext creates a new GetLolLobbyV2PartyActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLobbyV2PartyActiveParamsWithContext(ctx context.Context) *GetLolLobbyV2PartyActiveParams {

	return &GetLolLobbyV2PartyActiveParams{

		Context: ctx,
	}
}

// NewGetLolLobbyV2PartyActiveParamsWithHTTPClient creates a new GetLolLobbyV2PartyActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLobbyV2PartyActiveParamsWithHTTPClient(client *http.Client) *GetLolLobbyV2PartyActiveParams {

	return &GetLolLobbyV2PartyActiveParams{
		HTTPClient: client,
	}
}

/*GetLolLobbyV2PartyActiveParams contains all the parameters to send to the API endpoint
for the get lol lobby v2 party active operation typically these are written to a http.Request
*/
type GetLolLobbyV2PartyActiveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) WithTimeout(timeout time.Duration) *GetLolLobbyV2PartyActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) WithContext(ctx context.Context) *GetLolLobbyV2PartyActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) WithHTTPClient(client *http.Client) *GetLolLobbyV2PartyActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol lobby v2 party active params
func (o *GetLolLobbyV2PartyActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLobbyV2PartyActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
