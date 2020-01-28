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

// NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams creates a new GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized.
func NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams() *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithTimeout creates a new GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithContext creates a new GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithHTTPClient creates a new GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams contains all the parameters to send to the API endpoint
for the get lol collections v1 inventories by summoner Id summoner icons operation typically these are written to a http.Request
*/
type GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams struct {

	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) WithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) WithSummonerID(summonerID int64) *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v1 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV1InventoriesBySummonerIDSummonerIconsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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