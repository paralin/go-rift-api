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

// NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams creates a new GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized.
func NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams() *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithTimeout creates a new GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithContext creates a new GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithContext(ctx context.Context) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithHTTPClient creates a new GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	var ()
	return &GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams contains all the parameters to send to the API endpoint
for the get lol collections v2 inventories by summoner Id summoner icons operation typically these are written to a http.Request
*/
type GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams struct {

	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) WithContext(ctx context.Context) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) WithSummonerID(summonerID int64) *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v2 inventories by summoner Id summoner icons params
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV2InventoriesBySummonerIDSummonerIconsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
