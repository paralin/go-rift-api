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

// NewGetLolClubsV1ClubsByClubKeyMembersParams creates a new GetLolClubsV1ClubsByClubKeyMembersParams object
// with the default values initialized.
func NewGetLolClubsV1ClubsByClubKeyMembersParams() *GetLolClubsV1ClubsByClubKeyMembersParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClubsV1ClubsByClubKeyMembersParamsWithTimeout creates a new GetLolClubsV1ClubsByClubKeyMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClubsV1ClubsByClubKeyMembersParamsWithTimeout(timeout time.Duration) *GetLolClubsV1ClubsByClubKeyMembersParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyMembersParams{

		timeout: timeout,
	}
}

// NewGetLolClubsV1ClubsByClubKeyMembersParamsWithContext creates a new GetLolClubsV1ClubsByClubKeyMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClubsV1ClubsByClubKeyMembersParamsWithContext(ctx context.Context) *GetLolClubsV1ClubsByClubKeyMembersParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyMembersParams{

		Context: ctx,
	}
}

// NewGetLolClubsV1ClubsByClubKeyMembersParamsWithHTTPClient creates a new GetLolClubsV1ClubsByClubKeyMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClubsV1ClubsByClubKeyMembersParamsWithHTTPClient(client *http.Client) *GetLolClubsV1ClubsByClubKeyMembersParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyMembersParams{
		HTTPClient: client,
	}
}

/*GetLolClubsV1ClubsByClubKeyMembersParams contains all the parameters to send to the API endpoint
for the get lol clubs v1 clubs by club key members operation typically these are written to a http.Request
*/
type GetLolClubsV1ClubsByClubKeyMembersParams struct {

	/*ClubKey*/
	ClubKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) WithTimeout(timeout time.Duration) *GetLolClubsV1ClubsByClubKeyMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) WithContext(ctx context.Context) *GetLolClubsV1ClubsByClubKeyMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) WithHTTPClient(client *http.Client) *GetLolClubsV1ClubsByClubKeyMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) WithClubKey(clubKey string) *GetLolClubsV1ClubsByClubKeyMembersParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the get lol clubs v1 clubs by club key members params
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClubsV1ClubsByClubKeyMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clubKey
	if err := r.SetPathParam("clubKey", o.ClubKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
