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

// NewPostLolClashV1RosterByRosterIDDeclineParams creates a new PostLolClashV1RosterByRosterIDDeclineParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDDeclineParams() *PostLolClashV1RosterByRosterIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDDeclineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDDeclineParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDDeclineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDDeclineParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDDeclineParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDDeclineParamsWithContext creates a new PostLolClashV1RosterByRosterIDDeclineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDDeclineParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDDeclineParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDDeclineParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDDeclineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDDeclineParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDDeclineParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDDeclineParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id decline operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDDeclineParams struct {

	/*RosterID*/
	RosterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDDeclineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDDeclineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDDeclineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDDeclineParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id decline params
func (o *PostLolClashV1RosterByRosterIDDeclineParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDDeclineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rosterId
	if err := r.SetPathParam("rosterId", o.RosterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
