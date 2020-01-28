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

// NewPostLolChampSelectLegacyV1BattleTrainingLaunchParams creates a new PostLolChampSelectLegacyV1BattleTrainingLaunchParams object
// with the default values initialized.
func NewPostLolChampSelectLegacyV1BattleTrainingLaunchParams() *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {

	return &PostLolChampSelectLegacyV1BattleTrainingLaunchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithTimeout creates a new PostLolChampSelectLegacyV1BattleTrainingLaunchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithTimeout(timeout time.Duration) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {

	return &PostLolChampSelectLegacyV1BattleTrainingLaunchParams{

		timeout: timeout,
	}
}

// NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithContext creates a new PostLolChampSelectLegacyV1BattleTrainingLaunchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithContext(ctx context.Context) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {

	return &PostLolChampSelectLegacyV1BattleTrainingLaunchParams{

		Context: ctx,
	}
}

// NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithHTTPClient creates a new PostLolChampSelectLegacyV1BattleTrainingLaunchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChampSelectLegacyV1BattleTrainingLaunchParamsWithHTTPClient(client *http.Client) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {

	return &PostLolChampSelectLegacyV1BattleTrainingLaunchParams{
		HTTPClient: client,
	}
}

/*PostLolChampSelectLegacyV1BattleTrainingLaunchParams contains all the parameters to send to the API endpoint
for the post lol champ select legacy v1 battle training launch operation typically these are written to a http.Request
*/
type PostLolChampSelectLegacyV1BattleTrainingLaunchParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) WithTimeout(timeout time.Duration) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) WithContext(ctx context.Context) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) WithHTTPClient(client *http.Client) *PostLolChampSelectLegacyV1BattleTrainingLaunchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol champ select legacy v1 battle training launch params
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}