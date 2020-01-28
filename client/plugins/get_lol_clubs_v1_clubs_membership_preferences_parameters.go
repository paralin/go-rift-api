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

// NewGetLolClubsV1ClubsMembershipPreferencesParams creates a new GetLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized.
func NewGetLolClubsV1ClubsMembershipPreferencesParams() *GetLolClubsV1ClubsMembershipPreferencesParams {

	return &GetLolClubsV1ClubsMembershipPreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClubsV1ClubsMembershipPreferencesParamsWithTimeout creates a new GetLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClubsV1ClubsMembershipPreferencesParamsWithTimeout(timeout time.Duration) *GetLolClubsV1ClubsMembershipPreferencesParams {

	return &GetLolClubsV1ClubsMembershipPreferencesParams{

		timeout: timeout,
	}
}

// NewGetLolClubsV1ClubsMembershipPreferencesParamsWithContext creates a new GetLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClubsV1ClubsMembershipPreferencesParamsWithContext(ctx context.Context) *GetLolClubsV1ClubsMembershipPreferencesParams {

	return &GetLolClubsV1ClubsMembershipPreferencesParams{

		Context: ctx,
	}
}

// NewGetLolClubsV1ClubsMembershipPreferencesParamsWithHTTPClient creates a new GetLolClubsV1ClubsMembershipPreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClubsV1ClubsMembershipPreferencesParamsWithHTTPClient(client *http.Client) *GetLolClubsV1ClubsMembershipPreferencesParams {

	return &GetLolClubsV1ClubsMembershipPreferencesParams{
		HTTPClient: client,
	}
}

/*GetLolClubsV1ClubsMembershipPreferencesParams contains all the parameters to send to the API endpoint
for the get lol clubs v1 clubs membership preferences operation typically these are written to a http.Request
*/
type GetLolClubsV1ClubsMembershipPreferencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) WithTimeout(timeout time.Duration) *GetLolClubsV1ClubsMembershipPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) WithContext(ctx context.Context) *GetLolClubsV1ClubsMembershipPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) WithHTTPClient(client *http.Client) *GetLolClubsV1ClubsMembershipPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clubs v1 clubs membership preferences params
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClubsV1ClubsMembershipPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
