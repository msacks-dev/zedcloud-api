// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

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

// NewDeActivateEdgeNodeParams creates a new DeActivateEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeActivateEdgeNodeParams() *DeActivateEdgeNodeParams {
	return &DeActivateEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeActivateEdgeNodeParamsWithTimeout creates a new DeActivateEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewDeActivateEdgeNodeParamsWithTimeout(timeout time.Duration) *DeActivateEdgeNodeParams {
	return &DeActivateEdgeNodeParams{
		timeout: timeout,
	}
}

// NewDeActivateEdgeNodeParamsWithContext creates a new DeActivateEdgeNodeParams object
// with the ability to set a context for a request.
func NewDeActivateEdgeNodeParamsWithContext(ctx context.Context) *DeActivateEdgeNodeParams {
	return &DeActivateEdgeNodeParams{
		Context: ctx,
	}
}

// NewDeActivateEdgeNodeParamsWithHTTPClient creates a new DeActivateEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeActivateEdgeNodeParamsWithHTTPClient(client *http.Client) *DeActivateEdgeNodeParams {
	return &DeActivateEdgeNodeParams{
		HTTPClient: client,
	}
}

/* DeActivateEdgeNodeParams contains all the parameters to send to the API endpoint
   for the de activate edge node operation.

   Typically these are written to a http.Request.
*/
type DeActivateEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the de activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeActivateEdgeNodeParams) WithDefaults() *DeActivateEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the de activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeActivateEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the de activate edge node params
func (o *DeActivateEdgeNodeParams) WithTimeout(timeout time.Duration) *DeActivateEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the de activate edge node params
func (o *DeActivateEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the de activate edge node params
func (o *DeActivateEdgeNodeParams) WithContext(ctx context.Context) *DeActivateEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the de activate edge node params
func (o *DeActivateEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the de activate edge node params
func (o *DeActivateEdgeNodeParams) WithHTTPClient(client *http.Client) *DeActivateEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the de activate edge node params
func (o *DeActivateEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the de activate edge node params
func (o *DeActivateEdgeNodeParams) WithXRequestID(xRequestID *string) *DeActivateEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the de activate edge node params
func (o *DeActivateEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the de activate edge node params
func (o *DeActivateEdgeNodeParams) WithID(id string) *DeActivateEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the de activate edge node params
func (o *DeActivateEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeActivateEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}