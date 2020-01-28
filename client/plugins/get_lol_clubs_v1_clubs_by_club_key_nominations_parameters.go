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

// NewGetLolClubsV1ClubsByClubKeyNominationsParams creates a new GetLolClubsV1ClubsByClubKeyNominationsParams object
// with the default values initialized.
func NewGetLolClubsV1ClubsByClubKeyNominationsParams() *GetLolClubsV1ClubsByClubKeyNominationsParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyNominationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithTimeout creates a new GetLolClubsV1ClubsByClubKeyNominationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithTimeout(timeout time.Duration) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyNominationsParams{

		timeout: timeout,
	}
}

// NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithContext creates a new GetLolClubsV1ClubsByClubKeyNominationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithContext(ctx context.Context) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyNominationsParams{

		Context: ctx,
	}
}

// NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithHTTPClient creates a new GetLolClubsV1ClubsByClubKeyNominationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClubsV1ClubsByClubKeyNominationsParamsWithHTTPClient(client *http.Client) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	var ()
	return &GetLolClubsV1ClubsByClubKeyNominationsParams{
		HTTPClient: client,
	}
}

/*GetLolClubsV1ClubsByClubKeyNominationsParams contains all the parameters to send to the API endpoint
for the get lol clubs v1 clubs by club key nominations operation typically these are written to a http.Request
*/
type GetLolClubsV1ClubsByClubKeyNominationsParams struct {

	/*ClubKey*/
	ClubKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) WithTimeout(timeout time.Duration) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) WithContext(ctx context.Context) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) WithHTTPClient(client *http.Client) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClubKey adds the clubKey to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) WithClubKey(clubKey string) *GetLolClubsV1ClubsByClubKeyNominationsParams {
	o.SetClubKey(clubKey)
	return o
}

// SetClubKey adds the clubKey to the get lol clubs v1 clubs by club key nominations params
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) SetClubKey(clubKey string) {
	o.ClubKey = clubKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClubsV1ClubsByClubKeyNominationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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