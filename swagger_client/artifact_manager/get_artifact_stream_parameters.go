// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package artifact_manager

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

// NewGetArtifactStreamParams creates a new GetArtifactStreamParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArtifactStreamParams() *GetArtifactStreamParams {
	return &GetArtifactStreamParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactStreamParamsWithTimeout creates a new GetArtifactStreamParams object
// with the ability to set a timeout on a request.
func NewGetArtifactStreamParamsWithTimeout(timeout time.Duration) *GetArtifactStreamParams {
	return &GetArtifactStreamParams{
		timeout: timeout,
	}
}

// NewGetArtifactStreamParamsWithContext creates a new GetArtifactStreamParams object
// with the ability to set a context for a request.
func NewGetArtifactStreamParamsWithContext(ctx context.Context) *GetArtifactStreamParams {
	return &GetArtifactStreamParams{
		Context: ctx,
	}
}

// NewGetArtifactStreamParamsWithHTTPClient creates a new GetArtifactStreamParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArtifactStreamParamsWithHTTPClient(client *http.Client) *GetArtifactStreamParams {
	return &GetArtifactStreamParams{
		HTTPClient: client,
	}
}

/* GetArtifactStreamParams contains all the parameters to send to the API endpoint
   for the get artifact stream operation.

   Typically these are written to a http.Request.
*/
type GetArtifactStreamParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for an artifact
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get artifact stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactStreamParams) WithDefaults() *GetArtifactStreamParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get artifact stream params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArtifactStreamParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get artifact stream params
func (o *GetArtifactStreamParams) WithTimeout(timeout time.Duration) *GetArtifactStreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact stream params
func (o *GetArtifactStreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact stream params
func (o *GetArtifactStreamParams) WithContext(ctx context.Context) *GetArtifactStreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact stream params
func (o *GetArtifactStreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact stream params
func (o *GetArtifactStreamParams) WithHTTPClient(client *http.Client) *GetArtifactStreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact stream params
func (o *GetArtifactStreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get artifact stream params
func (o *GetArtifactStreamParams) WithXRequestID(xRequestID *string) *GetArtifactStreamParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get artifact stream params
func (o *GetArtifactStreamParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get artifact stream params
func (o *GetArtifactStreamParams) WithID(id string) *GetArtifactStreamParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get artifact stream params
func (o *GetArtifactStreamParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactStreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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