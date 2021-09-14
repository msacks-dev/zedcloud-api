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

// NewGetEdgeNodeAttestationParams creates a new GetEdgeNodeAttestationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEdgeNodeAttestationParams() *GetEdgeNodeAttestationParams {
	return &GetEdgeNodeAttestationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEdgeNodeAttestationParamsWithTimeout creates a new GetEdgeNodeAttestationParams object
// with the ability to set a timeout on a request.
func NewGetEdgeNodeAttestationParamsWithTimeout(timeout time.Duration) *GetEdgeNodeAttestationParams {
	return &GetEdgeNodeAttestationParams{
		timeout: timeout,
	}
}

// NewGetEdgeNodeAttestationParamsWithContext creates a new GetEdgeNodeAttestationParams object
// with the ability to set a context for a request.
func NewGetEdgeNodeAttestationParamsWithContext(ctx context.Context) *GetEdgeNodeAttestationParams {
	return &GetEdgeNodeAttestationParams{
		Context: ctx,
	}
}

// NewGetEdgeNodeAttestationParamsWithHTTPClient creates a new GetEdgeNodeAttestationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEdgeNodeAttestationParamsWithHTTPClient(client *http.Client) *GetEdgeNodeAttestationParams {
	return &GetEdgeNodeAttestationParams{
		HTTPClient: client,
	}
}

/* GetEdgeNodeAttestationParams contains all the parameters to send to the API endpoint
   for the get edge node attestation operation.

   Typically these are written to a http.Request.
*/
type GetEdgeNodeAttestationParams struct {

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

// WithDefaults hydrates default values in the get edge node attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeAttestationParams) WithDefaults() *GetEdgeNodeAttestationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get edge node attestation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEdgeNodeAttestationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) WithTimeout(timeout time.Duration) *GetEdgeNodeAttestationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) WithContext(ctx context.Context) *GetEdgeNodeAttestationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) WithHTTPClient(client *http.Client) *GetEdgeNodeAttestationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) WithXRequestID(xRequestID *string) *GetEdgeNodeAttestationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) WithID(id string) *GetEdgeNodeAttestationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get edge node attestation params
func (o *GetEdgeNodeAttestationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEdgeNodeAttestationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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