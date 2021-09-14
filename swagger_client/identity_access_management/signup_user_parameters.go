// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

	"github.com/zededa/zedcloud-api/swagger_models"
)

// NewSignupUserParams creates a new SignupUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSignupUserParams() *SignupUserParams {
	return &SignupUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSignupUserParamsWithTimeout creates a new SignupUserParams object
// with the ability to set a timeout on a request.
func NewSignupUserParamsWithTimeout(timeout time.Duration) *SignupUserParams {
	return &SignupUserParams{
		timeout: timeout,
	}
}

// NewSignupUserParamsWithContext creates a new SignupUserParams object
// with the ability to set a context for a request.
func NewSignupUserParamsWithContext(ctx context.Context) *SignupUserParams {
	return &SignupUserParams{
		Context: ctx,
	}
}

// NewSignupUserParamsWithHTTPClient creates a new SignupUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewSignupUserParamsWithHTTPClient(client *http.Client) *SignupUserParams {
	return &SignupUserParams{
		HTTPClient: client,
	}
}

/* SignupUserParams contains all the parameters to send to the API endpoint
   for the signup user operation.

   Typically these are written to a http.Request.
*/
type SignupUserParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *swagger_models.AAARequestAdminUserSignup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the signup user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignupUserParams) WithDefaults() *SignupUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the signup user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignupUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the signup user params
func (o *SignupUserParams) WithTimeout(timeout time.Duration) *SignupUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the signup user params
func (o *SignupUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the signup user params
func (o *SignupUserParams) WithContext(ctx context.Context) *SignupUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the signup user params
func (o *SignupUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the signup user params
func (o *SignupUserParams) WithHTTPClient(client *http.Client) *SignupUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the signup user params
func (o *SignupUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the signup user params
func (o *SignupUserParams) WithXRequestID(xRequestID *string) *SignupUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the signup user params
func (o *SignupUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the signup user params
func (o *SignupUserParams) WithBody(body *swagger_models.AAARequestAdminUserSignup) *SignupUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the signup user params
func (o *SignupUserParams) SetBody(body *swagger_models.AAARequestAdminUserSignup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SignupUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}