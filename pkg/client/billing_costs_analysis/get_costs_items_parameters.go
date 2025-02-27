// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package billing_costs_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCostsItemsParams creates a new GetCostsItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostsItemsParams() *GetCostsItemsParams {
	return &GetCostsItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostsItemsParamsWithTimeout creates a new GetCostsItemsParams object
// with the ability to set a timeout on a request.
func NewGetCostsItemsParamsWithTimeout(timeout time.Duration) *GetCostsItemsParams {
	return &GetCostsItemsParams{
		timeout: timeout,
	}
}

// NewGetCostsItemsParamsWithContext creates a new GetCostsItemsParams object
// with the ability to set a context for a request.
func NewGetCostsItemsParamsWithContext(ctx context.Context) *GetCostsItemsParams {
	return &GetCostsItemsParams{
		Context: ctx,
	}
}

// NewGetCostsItemsParamsWithHTTPClient creates a new GetCostsItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostsItemsParamsWithHTTPClient(client *http.Client) *GetCostsItemsParams {
	return &GetCostsItemsParams{
		HTTPClient: client,
	}
}

/* GetCostsItemsParams contains all the parameters to send to the API endpoint
   for the get costs items operation.

   Typically these are written to a http.Request.
*/
type GetCostsItemsParams struct {

	/* From.

	   A datetime for the beginning of the desired range for which to fetch costs. Defaults to start of current month. Note: there is currently a three-month maximum date range.
	*/
	From *string

	/* OrganizationID.

	   Identifier for the organization
	*/
	OrganizationID string

	/* To.

	   A datetime for the end of the desired range for which to fetch costs. Defaults to the current date. Note: there is currently a three-month maximum date range.
	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get costs items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostsItemsParams) WithDefaults() *GetCostsItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get costs items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostsItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get costs items params
func (o *GetCostsItemsParams) WithTimeout(timeout time.Duration) *GetCostsItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get costs items params
func (o *GetCostsItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get costs items params
func (o *GetCostsItemsParams) WithContext(ctx context.Context) *GetCostsItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get costs items params
func (o *GetCostsItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get costs items params
func (o *GetCostsItemsParams) WithHTTPClient(client *http.Client) *GetCostsItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get costs items params
func (o *GetCostsItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get costs items params
func (o *GetCostsItemsParams) WithFrom(from *string) *GetCostsItemsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get costs items params
func (o *GetCostsItemsParams) SetFrom(from *string) {
	o.From = from
}

// WithOrganizationID adds the organizationID to the get costs items params
func (o *GetCostsItemsParams) WithOrganizationID(organizationID string) *GetCostsItemsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get costs items params
func (o *GetCostsItemsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithTo adds the to to the get costs items params
func (o *GetCostsItemsParams) WithTo(to *string) *GetCostsItemsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get costs items params
func (o *GetCostsItemsParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostsItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if o.To != nil {

		// query param to
		var qrTo string

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
