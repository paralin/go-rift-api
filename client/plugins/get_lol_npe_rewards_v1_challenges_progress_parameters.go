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

// NewGetLolNpeRewardsV1ChallengesProgressParams creates a new GetLolNpeRewardsV1ChallengesProgressParams object
// with the default values initialized.
func NewGetLolNpeRewardsV1ChallengesProgressParams() *GetLolNpeRewardsV1ChallengesProgressParams {

	return &GetLolNpeRewardsV1ChallengesProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolNpeRewardsV1ChallengesProgressParamsWithTimeout creates a new GetLolNpeRewardsV1ChallengesProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolNpeRewardsV1ChallengesProgressParamsWithTimeout(timeout time.Duration) *GetLolNpeRewardsV1ChallengesProgressParams {

	return &GetLolNpeRewardsV1ChallengesProgressParams{

		timeout: timeout,
	}
}

// NewGetLolNpeRewardsV1ChallengesProgressParamsWithContext creates a new GetLolNpeRewardsV1ChallengesProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolNpeRewardsV1ChallengesProgressParamsWithContext(ctx context.Context) *GetLolNpeRewardsV1ChallengesProgressParams {

	return &GetLolNpeRewardsV1ChallengesProgressParams{

		Context: ctx,
	}
}

// NewGetLolNpeRewardsV1ChallengesProgressParamsWithHTTPClient creates a new GetLolNpeRewardsV1ChallengesProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolNpeRewardsV1ChallengesProgressParamsWithHTTPClient(client *http.Client) *GetLolNpeRewardsV1ChallengesProgressParams {

	return &GetLolNpeRewardsV1ChallengesProgressParams{
		HTTPClient: client,
	}
}

/*GetLolNpeRewardsV1ChallengesProgressParams contains all the parameters to send to the API endpoint
for the get lol npe rewards v1 challenges progress operation typically these are written to a http.Request
*/
type GetLolNpeRewardsV1ChallengesProgressParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) WithTimeout(timeout time.Duration) *GetLolNpeRewardsV1ChallengesProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) WithContext(ctx context.Context) *GetLolNpeRewardsV1ChallengesProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) WithHTTPClient(client *http.Client) *GetLolNpeRewardsV1ChallengesProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol npe rewards v1 challenges progress params
func (o *GetLolNpeRewardsV1ChallengesProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolNpeRewardsV1ChallengesProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
