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

// NewPostLolEndOfGameV1StateDismissStatsParams creates a new PostLolEndOfGameV1StateDismissStatsParams object
// with the default values initialized.
func NewPostLolEndOfGameV1StateDismissStatsParams() *PostLolEndOfGameV1StateDismissStatsParams {

	return &PostLolEndOfGameV1StateDismissStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolEndOfGameV1StateDismissStatsParamsWithTimeout creates a new PostLolEndOfGameV1StateDismissStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolEndOfGameV1StateDismissStatsParamsWithTimeout(timeout time.Duration) *PostLolEndOfGameV1StateDismissStatsParams {

	return &PostLolEndOfGameV1StateDismissStatsParams{

		timeout: timeout,
	}
}

// NewPostLolEndOfGameV1StateDismissStatsParamsWithContext creates a new PostLolEndOfGameV1StateDismissStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolEndOfGameV1StateDismissStatsParamsWithContext(ctx context.Context) *PostLolEndOfGameV1StateDismissStatsParams {

	return &PostLolEndOfGameV1StateDismissStatsParams{

		Context: ctx,
	}
}

// NewPostLolEndOfGameV1StateDismissStatsParamsWithHTTPClient creates a new PostLolEndOfGameV1StateDismissStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolEndOfGameV1StateDismissStatsParamsWithHTTPClient(client *http.Client) *PostLolEndOfGameV1StateDismissStatsParams {

	return &PostLolEndOfGameV1StateDismissStatsParams{
		HTTPClient: client,
	}
}

/*PostLolEndOfGameV1StateDismissStatsParams contains all the parameters to send to the API endpoint
for the post lol end of game v1 state dismiss stats operation typically these are written to a http.Request
*/
type PostLolEndOfGameV1StateDismissStatsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) WithTimeout(timeout time.Duration) *PostLolEndOfGameV1StateDismissStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) WithContext(ctx context.Context) *PostLolEndOfGameV1StateDismissStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) WithHTTPClient(client *http.Client) *PostLolEndOfGameV1StateDismissStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol end of game v1 state dismiss stats params
func (o *PostLolEndOfGameV1StateDismissStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolEndOfGameV1StateDismissStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}