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

// NewPostLolGameClientChatV1IgnoredSummonersParams creates a new PostLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized.
func NewPostLolGameClientChatV1IgnoredSummonersParams() *PostLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &PostLolGameClientChatV1IgnoredSummonersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolGameClientChatV1IgnoredSummonersParamsWithTimeout creates a new PostLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolGameClientChatV1IgnoredSummonersParamsWithTimeout(timeout time.Duration) *PostLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &PostLolGameClientChatV1IgnoredSummonersParams{

		timeout: timeout,
	}
}

// NewPostLolGameClientChatV1IgnoredSummonersParamsWithContext creates a new PostLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolGameClientChatV1IgnoredSummonersParamsWithContext(ctx context.Context) *PostLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &PostLolGameClientChatV1IgnoredSummonersParams{

		Context: ctx,
	}
}

// NewPostLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient creates a new PostLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient(client *http.Client) *PostLolGameClientChatV1IgnoredSummonersParams {
	var ()
	return &PostLolGameClientChatV1IgnoredSummonersParams{
		HTTPClient: client,
	}
}

/*PostLolGameClientChatV1IgnoredSummonersParams contains all the parameters to send to the API endpoint
for the post lol game client chat v1 ignored summoners operation typically these are written to a http.Request
*/
type PostLolGameClientChatV1IgnoredSummonersParams struct {

	/*SummonerName*/
	SummonerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) WithTimeout(timeout time.Duration) *PostLolGameClientChatV1IgnoredSummonersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) WithContext(ctx context.Context) *PostLolGameClientChatV1IgnoredSummonersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) WithHTTPClient(client *http.Client) *PostLolGameClientChatV1IgnoredSummonersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerName adds the summonerName to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) WithSummonerName(summonerName string) *PostLolGameClientChatV1IgnoredSummonersParams {
	o.SetSummonerName(summonerName)
	return o
}

// SetSummonerName adds the summonerName to the post lol game client chat v1 ignored summoners params
func (o *PostLolGameClientChatV1IgnoredSummonersParams) SetSummonerName(summonerName string) {
	o.SummonerName = summonerName
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolGameClientChatV1IgnoredSummonersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
