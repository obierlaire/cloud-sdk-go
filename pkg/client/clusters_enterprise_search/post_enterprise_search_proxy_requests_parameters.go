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

package clusters_enterprise_search

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

// NewPostEnterpriseSearchProxyRequestsParams creates a new PostEnterpriseSearchProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEnterpriseSearchProxyRequestsParams() *PostEnterpriseSearchProxyRequestsParams {
	return &PostEnterpriseSearchProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEnterpriseSearchProxyRequestsParamsWithTimeout creates a new PostEnterpriseSearchProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewPostEnterpriseSearchProxyRequestsParamsWithTimeout(timeout time.Duration) *PostEnterpriseSearchProxyRequestsParams {
	return &PostEnterpriseSearchProxyRequestsParams{
		timeout: timeout,
	}
}

// NewPostEnterpriseSearchProxyRequestsParamsWithContext creates a new PostEnterpriseSearchProxyRequestsParams object
// with the ability to set a context for a request.
func NewPostEnterpriseSearchProxyRequestsParamsWithContext(ctx context.Context) *PostEnterpriseSearchProxyRequestsParams {
	return &PostEnterpriseSearchProxyRequestsParams{
		Context: ctx,
	}
}

// NewPostEnterpriseSearchProxyRequestsParamsWithHTTPClient creates a new PostEnterpriseSearchProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEnterpriseSearchProxyRequestsParamsWithHTTPClient(client *http.Client) *PostEnterpriseSearchProxyRequestsParams {
	return &PostEnterpriseSearchProxyRequestsParams{
		HTTPClient: client,
	}
}

/* PostEnterpriseSearchProxyRequestsParams contains all the parameters to send to the API endpoint
   for the post enterprise search proxy requests operation.

   Typically these are written to a http.Request.
*/
type PostEnterpriseSearchProxyRequestsParams struct {

	/* XManagementRequest.

	   When set to `true`, includes the X-Management-Request header value.
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload that is used to proxy the EnterpriseSearch deployment.
	*/
	Body string

	/* ClusterID.

	   The EnterpriseSearch deployment identifier
	*/
	ClusterID string

	/* EnterpriseSearchPath.

	   The URL part to proxy to the EnterpriseSearch cluster. Example: /api/ent/v1/internal/read_only_mode or /api/ent/v1/internal/health
	*/
	EnterpriseSearchPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEnterpriseSearchProxyRequestsParams) WithDefaults() *PostEnterpriseSearchProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post enterprise search proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEnterpriseSearchProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithTimeout(timeout time.Duration) *PostEnterpriseSearchProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithContext(ctx context.Context) *PostEnterpriseSearchProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithHTTPClient(client *http.Client) *PostEnterpriseSearchProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *PostEnterpriseSearchProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithBody(body string) *PostEnterpriseSearchProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithClusterID(clusterID string) *PostEnterpriseSearchProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEnterpriseSearchPath adds the enterpriseSearchPath to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) WithEnterpriseSearchPath(enterpriseSearchPath string) *PostEnterpriseSearchProxyRequestsParams {
	o.SetEnterpriseSearchPath(enterpriseSearchPath)
	return o
}

// SetEnterpriseSearchPath adds the enterpriseSearchPath to the post enterprise search proxy requests params
func (o *PostEnterpriseSearchProxyRequestsParams) SetEnterpriseSearchPath(enterpriseSearchPath string) {
	o.EnterpriseSearchPath = enterpriseSearchPath
}

// WriteToRequest writes these params to a swagger request
func (o *PostEnterpriseSearchProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param enterprise_search_path
	if err := r.SetPathParam("enterprise_search_path", o.EnterpriseSearchPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
