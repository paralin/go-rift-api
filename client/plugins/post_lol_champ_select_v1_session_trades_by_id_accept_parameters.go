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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostLolChampSelectV1SessionTradesByIDAcceptParams creates a new PostLolChampSelectV1SessionTradesByIDAcceptParams object
// with the default values initialized.
func NewPostLolChampSelectV1SessionTradesByIDAcceptParams() *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDAcceptParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithTimeout creates a new PostLolChampSelectV1SessionTradesByIDAcceptParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithTimeout(timeout time.Duration) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDAcceptParams{

		timeout: timeout,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithContext creates a new PostLolChampSelectV1SessionTradesByIDAcceptParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithContext(ctx context.Context) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDAcceptParams{

		Context: ctx,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithHTTPClient creates a new PostLolChampSelectV1SessionTradesByIDAcceptParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChampSelectV1SessionTradesByIDAcceptParamsWithHTTPClient(client *http.Client) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDAcceptParams{
		HTTPClient: client,
	}
}

/*PostLolChampSelectV1SessionTradesByIDAcceptParams contains all the parameters to send to the API endpoint
for the post lol champ select v1 session trades by Id accept operation typically these are written to a http.Request
*/
type PostLolChampSelectV1SessionTradesByIDAcceptParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) WithTimeout(timeout time.Duration) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) WithContext(ctx context.Context) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) WithHTTPClient(client *http.Client) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) WithID(id int64) *PostLolChampSelectV1SessionTradesByIDAcceptParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol champ select v1 session trades by Id accept params
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChampSelectV1SessionTradesByIDAcceptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
