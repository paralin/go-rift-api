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

// NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams creates a new PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized.
func NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams() *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithTimeout creates a new PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithTimeout(timeout time.Duration) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams{

		timeout: timeout,
	}
}

// NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithContext creates a new PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithContext(ctx context.Context) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams{

		Context: ctx,
	}
}

// NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithHTTPClient creates a new PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParamsWithHTTPClient(client *http.Client) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	var ()
	return &PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams{
		HTTPClient: client,
	}
}

/*PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams contains all the parameters to send to the API endpoint
for the post lol lobby team builder champ select v1 session trades by Id request operation typically these are written to a http.Request
*/
type PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) WithTimeout(timeout time.Duration) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) WithContext(ctx context.Context) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) WithHTTPClient(client *http.Client) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) WithID(id int32) *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post lol lobby team builder champ select v1 session trades by Id request params
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}