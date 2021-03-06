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

// NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParams creates a new GetLolCollectionsV1InventoriesBySummonerIDSpellsParams object
// with the default values initialized.
func NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParams() *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSpellsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithTimeout creates a new GetLolCollectionsV1InventoriesBySummonerIDSpellsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSpellsParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithContext creates a new GetLolCollectionsV1InventoriesBySummonerIDSpellsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSpellsParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithHTTPClient creates a new GetLolCollectionsV1InventoriesBySummonerIDSpellsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDSpellsParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDSpellsParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV1InventoriesBySummonerIDSpellsParams contains all the parameters to send to the API endpoint
for the get lol collections v1 inventories by summoner Id spells operation typically these are written to a http.Request
*/
type GetLolCollectionsV1InventoriesBySummonerIDSpellsParams struct {

	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) WithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerID adds the summonerID to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) WithSummonerID(summonerID int64) *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v1 inventories by summoner Id spells params
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV1InventoriesBySummonerIDSpellsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
