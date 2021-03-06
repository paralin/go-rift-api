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

// NewGetLolInventoryV1SignedInventoryParams creates a new GetLolInventoryV1SignedInventoryParams object
// with the default values initialized.
func NewGetLolInventoryV1SignedInventoryParams() *GetLolInventoryV1SignedInventoryParams {
	var ()
	return &GetLolInventoryV1SignedInventoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolInventoryV1SignedInventoryParamsWithTimeout creates a new GetLolInventoryV1SignedInventoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolInventoryV1SignedInventoryParamsWithTimeout(timeout time.Duration) *GetLolInventoryV1SignedInventoryParams {
	var ()
	return &GetLolInventoryV1SignedInventoryParams{

		timeout: timeout,
	}
}

// NewGetLolInventoryV1SignedInventoryParamsWithContext creates a new GetLolInventoryV1SignedInventoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolInventoryV1SignedInventoryParamsWithContext(ctx context.Context) *GetLolInventoryV1SignedInventoryParams {
	var ()
	return &GetLolInventoryV1SignedInventoryParams{

		Context: ctx,
	}
}

// NewGetLolInventoryV1SignedInventoryParamsWithHTTPClient creates a new GetLolInventoryV1SignedInventoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolInventoryV1SignedInventoryParamsWithHTTPClient(client *http.Client) *GetLolInventoryV1SignedInventoryParams {
	var ()
	return &GetLolInventoryV1SignedInventoryParams{
		HTTPClient: client,
	}
}

/*GetLolInventoryV1SignedInventoryParams contains all the parameters to send to the API endpoint
for the get lol inventory v1 signed inventory operation typically these are written to a http.Request
*/
type GetLolInventoryV1SignedInventoryParams struct {

	/*InventoryTypes*/
	InventoryTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) WithTimeout(timeout time.Duration) *GetLolInventoryV1SignedInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) WithContext(ctx context.Context) *GetLolInventoryV1SignedInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) WithHTTPClient(client *http.Client) *GetLolInventoryV1SignedInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryTypes adds the inventoryTypes to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) WithInventoryTypes(inventoryTypes []string) *GetLolInventoryV1SignedInventoryParams {
	o.SetInventoryTypes(inventoryTypes)
	return o
}

// SetInventoryTypes adds the inventoryTypes to the get lol inventory v1 signed inventory params
func (o *GetLolInventoryV1SignedInventoryParams) SetInventoryTypes(inventoryTypes []string) {
	o.InventoryTypes = inventoryTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolInventoryV1SignedInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesInventoryTypes := o.InventoryTypes

	joinedInventoryTypes := swag.JoinByFormat(valuesInventoryTypes, "")
	// query array param inventoryTypes
	if err := r.SetQueryParam("inventoryTypes", joinedInventoryTypes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
