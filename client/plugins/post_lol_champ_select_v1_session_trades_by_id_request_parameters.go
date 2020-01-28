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

// NewPostLolChampSelectV1SessionTradesByIDRequestParams creates a new PostLolChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized.
func NewPostLolChampSelectV1SessionTradesByIDRequestParams() *PostLolChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithTimeout creates a new PostLolChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithTimeout(timeout time.Duration) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDRequestParams{

		timeout: timeout,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithContext creates a new PostLolChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithContext(ctx context.Context) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDRequestParams{

		Context: ctx,
	}
}

// NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithHTTPClient creates a new PostLolChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolChampSelectV1SessionTradesByIDRequestParamsWithHTTPClient(client *http.Client) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolChampSelectV1SessionTradesByIDRequestParams{
		HTTPClient: client,
	}
}

/*PostLolChampSelectV1SessionTradesByIDRequestParams contains all the parameters to send to the API endpoint
for the post lol champ select v1 session trades by Id request operation typically these are written to a http.Request
*/
type PostLolChampSelectV1SessionTradesByIDRequestParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) WithTimeout(timeout time.Duration) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) WithContext(ctx context.Context) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) WithHTTPClient(client *http.Client) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) WithID(id int64) *PostLolChampSelectV1SessionTradesByIDRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol champ select v1 session trades by Id request params
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolChampSelectV1SessionTradesByIDRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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