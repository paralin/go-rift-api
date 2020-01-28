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

// NewPostLolNpeRewardsV1ChallengesOptParams creates a new PostLolNpeRewardsV1ChallengesOptParams object
// with the default values initialized.
func NewPostLolNpeRewardsV1ChallengesOptParams() *PostLolNpeRewardsV1ChallengesOptParams {

	return &PostLolNpeRewardsV1ChallengesOptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolNpeRewardsV1ChallengesOptParamsWithTimeout creates a new PostLolNpeRewardsV1ChallengesOptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolNpeRewardsV1ChallengesOptParamsWithTimeout(timeout time.Duration) *PostLolNpeRewardsV1ChallengesOptParams {

	return &PostLolNpeRewardsV1ChallengesOptParams{

		timeout: timeout,
	}
}

// NewPostLolNpeRewardsV1ChallengesOptParamsWithContext creates a new PostLolNpeRewardsV1ChallengesOptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolNpeRewardsV1ChallengesOptParamsWithContext(ctx context.Context) *PostLolNpeRewardsV1ChallengesOptParams {

	return &PostLolNpeRewardsV1ChallengesOptParams{

		Context: ctx,
	}
}

// NewPostLolNpeRewardsV1ChallengesOptParamsWithHTTPClient creates a new PostLolNpeRewardsV1ChallengesOptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolNpeRewardsV1ChallengesOptParamsWithHTTPClient(client *http.Client) *PostLolNpeRewardsV1ChallengesOptParams {

	return &PostLolNpeRewardsV1ChallengesOptParams{
		HTTPClient: client,
	}
}

/*PostLolNpeRewardsV1ChallengesOptParams contains all the parameters to send to the API endpoint
for the post lol npe rewards v1 challenges opt operation typically these are written to a http.Request
*/
type PostLolNpeRewardsV1ChallengesOptParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) WithTimeout(timeout time.Duration) *PostLolNpeRewardsV1ChallengesOptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) WithContext(ctx context.Context) *PostLolNpeRewardsV1ChallengesOptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) WithHTTPClient(client *http.Client) *PostLolNpeRewardsV1ChallengesOptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol npe rewards v1 challenges opt params
func (o *PostLolNpeRewardsV1ChallengesOptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolNpeRewardsV1ChallengesOptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
