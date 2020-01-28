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

// NewPostLolGameflowV1SessionTournamentEndedParams creates a new PostLolGameflowV1SessionTournamentEndedParams object
// with the default values initialized.
func NewPostLolGameflowV1SessionTournamentEndedParams() *PostLolGameflowV1SessionTournamentEndedParams {

	return &PostLolGameflowV1SessionTournamentEndedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolGameflowV1SessionTournamentEndedParamsWithTimeout creates a new PostLolGameflowV1SessionTournamentEndedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolGameflowV1SessionTournamentEndedParamsWithTimeout(timeout time.Duration) *PostLolGameflowV1SessionTournamentEndedParams {

	return &PostLolGameflowV1SessionTournamentEndedParams{

		timeout: timeout,
	}
}

// NewPostLolGameflowV1SessionTournamentEndedParamsWithContext creates a new PostLolGameflowV1SessionTournamentEndedParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolGameflowV1SessionTournamentEndedParamsWithContext(ctx context.Context) *PostLolGameflowV1SessionTournamentEndedParams {

	return &PostLolGameflowV1SessionTournamentEndedParams{

		Context: ctx,
	}
}

// NewPostLolGameflowV1SessionTournamentEndedParamsWithHTTPClient creates a new PostLolGameflowV1SessionTournamentEndedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolGameflowV1SessionTournamentEndedParamsWithHTTPClient(client *http.Client) *PostLolGameflowV1SessionTournamentEndedParams {

	return &PostLolGameflowV1SessionTournamentEndedParams{
		HTTPClient: client,
	}
}

/*PostLolGameflowV1SessionTournamentEndedParams contains all the parameters to send to the API endpoint
for the post lol gameflow v1 session tournament ended operation typically these are written to a http.Request
*/
type PostLolGameflowV1SessionTournamentEndedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) WithTimeout(timeout time.Duration) *PostLolGameflowV1SessionTournamentEndedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) WithContext(ctx context.Context) *PostLolGameflowV1SessionTournamentEndedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) WithHTTPClient(client *http.Client) *PostLolGameflowV1SessionTournamentEndedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol gameflow v1 session tournament ended params
func (o *PostLolGameflowV1SessionTournamentEndedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolGameflowV1SessionTournamentEndedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
