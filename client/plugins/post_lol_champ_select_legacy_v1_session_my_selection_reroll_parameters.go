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

// NewPostLolChampSelectLegacyV1SessionMySelectionRerollParams creates a new PostLolChampSelectLegacyV1SessionMySelectionRerollParams object
// with the default values initialized.
func NewPostLolChampSelectLegacyV1SessionMySelectionRerollParams() *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {

	return &PostLolChampSelectLegacyV1SessionMySelectionRerollParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithTimeout creates a new PostLolChampSelectLegacyV1SessionMySelectionRerollParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithTimeout(timeout time.Duration) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {

	return &PostLolChampSelectLegacyV1SessionMySelectionRerollParams{

		timeout: timeout,
	}
}

// NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithContext creates a new PostLolChampSelectLegacyV1SessionMySelectionRerollParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithContext(ctx context.Context) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {

	return &PostLolChampSelectLegacyV1SessionMySelectionRerollParams{

		Context: ctx,
	}
}

// NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithHTTPClient creates a new PostLolChampSelectLegacyV1SessionMySelectionRerollParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChampSelectLegacyV1SessionMySelectionRerollParamsWithHTTPClient(client *http.Client) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {

	return &PostLolChampSelectLegacyV1SessionMySelectionRerollParams{
		HTTPClient: client,
	}
}

/*PostLolChampSelectLegacyV1SessionMySelectionRerollParams contains all the parameters to send to the API endpoint
for the post lol champ select legacy v1 session my selection reroll operation typically these are written to a http.Request
*/
type PostLolChampSelectLegacyV1SessionMySelectionRerollParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) WithTimeout(timeout time.Duration) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) WithContext(ctx context.Context) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) WithHTTPClient(client *http.Client) *PostLolChampSelectLegacyV1SessionMySelectionRerollParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol champ select legacy v1 session my selection reroll params
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChampSelectLegacyV1SessionMySelectionRerollParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
