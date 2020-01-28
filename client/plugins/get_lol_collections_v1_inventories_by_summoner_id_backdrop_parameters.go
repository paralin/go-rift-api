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

// NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParams creates a new GetLolCollectionsV1InventoriesBySummonerIDBackdropParams object
// with the default values initialized.
func NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParams() *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDBackdropParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithTimeout creates a new GetLolCollectionsV1InventoriesBySummonerIDBackdropParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDBackdropParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithContext creates a new GetLolCollectionsV1InventoriesBySummonerIDBackdropParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDBackdropParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithHTTPClient creates a new GetLolCollectionsV1InventoriesBySummonerIDBackdropParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDBackdropParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDBackdropParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV1InventoriesBySummonerIDBackdropParams contains all the parameters to send to the API endpoint
for the get lol collections v1 inventories by summoner Id backdrop operation typically these are written to a http.Request
*/
type GetLolCollectionsV1InventoriesBySummonerIDBackdropParams struct {

	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) WithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) WithSummonerID(summonerID int64) *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v1 inventories by summoner Id backdrop params
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV1InventoriesBySummonerIDBackdropParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}