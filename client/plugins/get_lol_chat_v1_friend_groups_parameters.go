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

// NewGetLolChatV1FriendGroupsParams creates a new GetLolChatV1FriendGroupsParams object
// with the default values initialized.
func NewGetLolChatV1FriendGroupsParams() *GetLolChatV1FriendGroupsParams {

	return &GetLolChatV1FriendGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChatV1FriendGroupsParamsWithTimeout creates a new GetLolChatV1FriendGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChatV1FriendGroupsParamsWithTimeout(timeout time.Duration) *GetLolChatV1FriendGroupsParams {

	return &GetLolChatV1FriendGroupsParams{

		timeout: timeout,
	}
}

// NewGetLolChatV1FriendGroupsParamsWithContext creates a new GetLolChatV1FriendGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChatV1FriendGroupsParamsWithContext(ctx context.Context) *GetLolChatV1FriendGroupsParams {

	return &GetLolChatV1FriendGroupsParams{

		Context: ctx,
	}
}

// NewGetLolChatV1FriendGroupsParamsWithHTTPClient creates a new GetLolChatV1FriendGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChatV1FriendGroupsParamsWithHTTPClient(client *http.Client) *GetLolChatV1FriendGroupsParams {

	return &GetLolChatV1FriendGroupsParams{
		HTTPClient: client,
	}
}

/*GetLolChatV1FriendGroupsParams contains all the parameters to send to the API endpoint
for the get lol chat v1 friend groups operation typically these are written to a http.Request
*/
type GetLolChatV1FriendGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) WithTimeout(timeout time.Duration) *GetLolChatV1FriendGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) WithContext(ctx context.Context) *GetLolChatV1FriendGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) WithHTTPClient(client *http.Client) *GetLolChatV1FriendGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol chat v1 friend groups params
func (o *GetLolChatV1FriendGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChatV1FriendGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
