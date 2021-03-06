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

// NewPostLolClashV1RosterByRosterIDSubstituteDeclineParams creates a new PostLolClashV1RosterByRosterIDSubstituteDeclineParams object
// with the default values initialized.
func NewPostLolClashV1RosterByRosterIDSubstituteDeclineParams() *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteDeclineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithTimeout creates a new PostLolClashV1RosterByRosterIDSubstituteDeclineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteDeclineParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithContext creates a new PostLolClashV1RosterByRosterIDSubstituteDeclineParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteDeclineParams{

		Context: ctx,
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithHTTPClient creates a new PostLolClashV1RosterByRosterIDSubstituteDeclineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1RosterByRosterIDSubstituteDeclineParamsWithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	var ()
	return &PostLolClashV1RosterByRosterIDSubstituteDeclineParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1RosterByRosterIDSubstituteDeclineParams contains all the parameters to send to the API endpoint
for the post lol clash v1 roster by roster Id substitute decline operation typically these are written to a http.Request
*/
type PostLolClashV1RosterByRosterIDSubstituteDeclineParams struct {

	/*RosterID*/
	RosterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) WithTimeout(timeout time.Duration) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) WithContext(ctx context.Context) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) WithHTTPClient(client *http.Client) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRosterID adds the rosterID to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) WithRosterID(rosterID string) *PostLolClashV1RosterByRosterIDSubstituteDeclineParams {
	o.SetRosterID(rosterID)
	return o
}

// SetRosterID adds the rosterId to the post lol clash v1 roster by roster Id substitute decline params
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) SetRosterID(rosterID string) {
	o.RosterID = rosterID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
