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

// NewDeleteLolGameClientChatV1IgnoredSummonersParams creates a new DeleteLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized.
func NewDeleteLolGameClientChatV1IgnoredSummonersParams() *DeleteLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &DeleteLolGameClientChatV1IgnoredSummonersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithTimeout creates a new DeleteLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithTimeout(timeout time.Duration) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &DeleteLolGameClientChatV1IgnoredSummonersParams{

		timeout: timeout,
	}
}

// NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithContext creates a new DeleteLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithContext(ctx context.Context) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &DeleteLolGameClientChatV1IgnoredSummonersParams{

		Context: ctx,
	}
}

// NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient creates a new DeleteLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient(client *http.Client) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &DeleteLolGameClientChatV1IgnoredSummonersParams{
		HTTPClient: client,
	}
}

/*DeleteLolGameClientChatV1IgnoredSummonersParams contains all the parameters to send to the API endpoint
for the delete lol game client chat v1 ignored summoners operation typically these are written to a http.Request
*/
type DeleteLolGameClientChatV1IgnoredSummonersParams struct {

	/*SummonerName*/
	SummonerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) WithTimeout(timeout time.Duration) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) WithContext(ctx context.Context) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) WithHTTPClient(client *http.Client) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerName adds the summonerName to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) WithSummonerName(summonerName string) *DeleteLolGameClientChatV1IgnoredSummonersParams {
	o.SetSummonerName(summonerName)
	return o
}

// SetSummonerName adds the summonerName to the delete lol game client chat v1 ignored summoners params
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) SetSummonerName(summonerName string) {
	o.SummonerName = summonerName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLolGameClientChatV1IgnoredSummonersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param summonerName
	qrSummonerName := o.SummonerName
	qSummonerName := qrSummonerName
	if qSummonerName != "" {
		if err := r.SetQueryParam("summonerName", qSummonerName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
