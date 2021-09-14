// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package cloud_diagnostics

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

// NewCreateCloudPolicyDocumentParams creates a new CreateCloudPolicyDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCloudPolicyDocumentParams() *CreateCloudPolicyDocumentParams {
	return &CreateCloudPolicyDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudPolicyDocumentParamsWithTimeout creates a new CreateCloudPolicyDocumentParams object
// with the ability to set a timeout on a request.
func NewCreateCloudPolicyDocumentParamsWithTimeout(timeout time.Duration) *CreateCloudPolicyDocumentParams {
	return &CreateCloudPolicyDocumentParams{
		timeout: timeout,
	}
}

// NewCreateCloudPolicyDocumentParamsWithContext creates a new CreateCloudPolicyDocumentParams object
// with the ability to set a context for a request.
func NewCreateCloudPolicyDocumentParamsWithContext(ctx context.Context) *CreateCloudPolicyDocumentParams {
	return &CreateCloudPolicyDocumentParams{
		Context: ctx,
	}
}

// NewCreateCloudPolicyDocumentParamsWithHTTPClient creates a new CreateCloudPolicyDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCloudPolicyDocumentParamsWithHTTPClient(client *http.Client) *CreateCloudPolicyDocumentParams {
	return &CreateCloudPolicyDocumentParams{
		HTTPClient: client,
	}
}

/* CreateCloudPolicyDocumentParams contains all the parameters to send to the API endpoint
   for the create cloud policy document operation.

   Typically these are written to a http.Request.
*/
type CreateCloudPolicyDocumentParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudPolicyDocumentParams) WithDefaults() *CreateCloudPolicyDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cloud policy document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCloudPolicyDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) WithTimeout(timeout time.Duration) *CreateCloudPolicyDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) WithContext(ctx context.Context) *CreateCloudPolicyDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) WithHTTPClient(client *http.Client) *CreateCloudPolicyDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) WithXRequestID(xRequestID *string) *CreateCloudPolicyDocumentParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create cloud policy document params
func (o *CreateCloudPolicyDocumentParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudPolicyDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}