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

// NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams creates a new GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams object
// with the default values initialized.
func NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams() *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithTimeout creates a new GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams{

		timeout: timeout,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithContext creates a new GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams{

		Context: ctx,
	}
}

// NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithHTTPClient creates a new GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParamsWithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	var ()
	return &GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams{
		HTTPClient: client,
	}
}

/*GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams contains all the parameters to send to the API endpoint
for the get lol collections v1 inventories by summoner Id champion mastery top operation typically these are written to a http.Request
*/
type GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams struct {

	/*Limit*/
	Limit int64
	/*SortRule*/
	SortRule *string
	/*SummonerID*/
	SummonerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithTimeout(timeout time.Duration) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithContext(ctx context.Context) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithHTTPClient(client *http.Client) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithLimit(limit int64) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithSortRule adds the sortRule to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithSortRule(sortRule *string) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetSortRule(sortRule)
	return o
}

// SetSortRule adds the sortRule to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetSortRule(sortRule *string) {
	o.SortRule = sortRule
}

// WithSummonerID adds the summonerID to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WithSummonerID(summonerID int64) *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams {
	o.SetSummonerID(summonerID)
	return o
}

// SetSummonerID adds the summonerId to the get lol collections v1 inventories by summoner Id champion mastery top params
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) SetSummonerID(summonerID int64) {
	o.SummonerID = summonerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolCollectionsV1InventoriesBySummonerIDChampionMasteryTopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)
	if qLimit != "" {
		if err := r.SetQueryParam("limit", qLimit); err != nil {
			return err
		}
	}

	if o.SortRule != nil {

		// query param sortRule
		var qrSortRule string
		if o.SortRule != nil {
			qrSortRule = *o.SortRule
		}
		qSortRule := qrSortRule
		if qSortRule != "" {
			if err := r.SetQueryParam("sortRule", qSortRule); err != nil {
				return err
			}
		}

	}

	// path param summonerId
	if err := r.SetPathParam("summonerId", swag.FormatInt64(o.SummonerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
