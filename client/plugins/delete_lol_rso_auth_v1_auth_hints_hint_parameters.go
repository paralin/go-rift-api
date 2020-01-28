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

// NewDeleteLolRsoAuthV1AuthHintsHintParams creates a new DeleteLolRsoAuthV1AuthHintsHintParams object
// with the default values initialized.
func NewDeleteLolRsoAuthV1AuthHintsHintParams() *DeleteLolRsoAuthV1AuthHintsHintParams {

	return &DeleteLolRsoAuthV1AuthHintsHintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolRsoAuthV1AuthHintsHintParamsWithTimeout creates a new DeleteLolRsoAuthV1AuthHintsHintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolRsoAuthV1AuthHintsHintParamsWithTimeout(timeout time.Duration) *DeleteLolRsoAuthV1AuthHintsHintParams {

	return &DeleteLolRsoAuthV1AuthHintsHintParams{

		timeout: timeout,
	}
}

// NewDeleteLolRsoAuthV1AuthHintsHintParamsWithContext creates a new DeleteLolRsoAuthV1AuthHintsHintParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolRsoAuthV1AuthHintsHintParamsWithContext(ctx context.Context) *DeleteLolRsoAuthV1AuthHintsHintParams {

	return &DeleteLolRsoAuthV1AuthHintsHintParams{

		Context: ctx,
	}
}

// NewDeleteLolRsoAuthV1AuthHintsHintParamsWithHTTPClient creates a new DeleteLolRsoAuthV1AuthHintsHintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolRsoAuthV1AuthHintsHintParamsWithHTTPClient(client *http.Client) *DeleteLolRsoAuthV1AuthHintsHintParams {

	return &DeleteLolRsoAuthV1AuthHintsHintParams{
		HTTPClient: client,
	}
}

/*DeleteLolRsoAuthV1AuthHintsHintParams contains all the parameters to send to the API endpoint
for the delete lol rso auth v1 auth hints hint operation typically these are written to a http.Request
*/
type DeleteLolRsoAuthV1AuthHintsHintParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) WithTimeout(timeout time.Duration) *DeleteLolRsoAuthV1AuthHintsHintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) WithContext(ctx context.Context) *DeleteLolRsoAuthV1AuthHintsHintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) WithHTTPClient(client *http.Client) *DeleteLolRsoAuthV1AuthHintsHintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol rso auth v1 auth hints hint params
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolRsoAuthV1AuthHintsHintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}