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

// NewPostLolGameflowV1BattleTrainingStartParams creates a new PostLolGameflowV1BattleTrainingStartParams object
// with the default values initialized.
func NewPostLolGameflowV1BattleTrainingStartParams() *PostLolGameflowV1BattleTrainingStartParams {

	return &PostLolGameflowV1BattleTrainingStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolGameflowV1BattleTrainingStartParamsWithTimeout creates a new PostLolGameflowV1BattleTrainingStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolGameflowV1BattleTrainingStartParamsWithTimeout(timeout time.Duration) *PostLolGameflowV1BattleTrainingStartParams {

	return &PostLolGameflowV1BattleTrainingStartParams{

		timeout: timeout,
	}
}

// NewPostLolGameflowV1BattleTrainingStartParamsWithContext creates a new PostLolGameflowV1BattleTrainingStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolGameflowV1BattleTrainingStartParamsWithContext(ctx context.Context) *PostLolGameflowV1BattleTrainingStartParams {

	return &PostLolGameflowV1BattleTrainingStartParams{

		Context: ctx,
	}
}

// NewPostLolGameflowV1BattleTrainingStartParamsWithHTTPClient creates a new PostLolGameflowV1BattleTrainingStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolGameflowV1BattleTrainingStartParamsWithHTTPClient(client *http.Client) *PostLolGameflowV1BattleTrainingStartParams {

	return &PostLolGameflowV1BattleTrainingStartParams{
		HTTPClient: client,
	}
}

/*PostLolGameflowV1BattleTrainingStartParams contains all the parameters to send to the API endpoint
for the post lol gameflow v1 battle training start operation typically these are written to a http.Request
*/
type PostLolGameflowV1BattleTrainingStartParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) WithTimeout(timeout time.Duration) *PostLolGameflowV1BattleTrainingStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) WithContext(ctx context.Context) *PostLolGameflowV1BattleTrainingStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) WithHTTPClient(client *http.Client) *PostLolGameflowV1BattleTrainingStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol gameflow v1 battle training start params
func (o *PostLolGameflowV1BattleTrainingStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolGameflowV1BattleTrainingStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
