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

// NewGetLolClashV1InvitedRosterIdsParams creates a new GetLolClashV1InvitedRosterIdsParams object
// with the default values initialized.
func NewGetLolClashV1InvitedRosterIdsParams() *GetLolClashV1InvitedRosterIdsParams {

	return &GetLolClashV1InvitedRosterIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1InvitedRosterIdsParamsWithTimeout creates a new GetLolClashV1InvitedRosterIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1InvitedRosterIdsParamsWithTimeout(timeout time.Duration) *GetLolClashV1InvitedRosterIdsParams {

	return &GetLolClashV1InvitedRosterIdsParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1InvitedRosterIdsParamsWithContext creates a new GetLolClashV1InvitedRosterIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1InvitedRosterIdsParamsWithContext(ctx context.Context) *GetLolClashV1InvitedRosterIdsParams {

	return &GetLolClashV1InvitedRosterIdsParams{

		Context: ctx,
	}
}

// NewGetLolClashV1InvitedRosterIdsParamsWithHTTPClient creates a new GetLolClashV1InvitedRosterIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1InvitedRosterIdsParamsWithHTTPClient(client *http.Client) *GetLolClashV1InvitedRosterIdsParams {

	return &GetLolClashV1InvitedRosterIdsParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1InvitedRosterIdsParams contains all the parameters to send to the API endpoint
for the get lol clash v1 invited roster ids operation typically these are written to a http.Request
*/
type GetLolClashV1InvitedRosterIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) WithTimeout(timeout time.Duration) *GetLolClashV1InvitedRosterIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) WithContext(ctx context.Context) *GetLolClashV1InvitedRosterIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) WithHTTPClient(client *http.Client) *GetLolClashV1InvitedRosterIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 invited roster ids params
func (o *GetLolClashV1InvitedRosterIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1InvitedRosterIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
