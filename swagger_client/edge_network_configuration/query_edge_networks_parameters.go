// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_network_configuration

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
	"github.com/go-openapi/swag"
)

// NewQueryEdgeNetworksParams creates a new QueryEdgeNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryEdgeNetworksParams() *QueryEdgeNetworksParams {
	return &QueryEdgeNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryEdgeNetworksParamsWithTimeout creates a new QueryEdgeNetworksParams object
// with the ability to set a timeout on a request.
func NewQueryEdgeNetworksParamsWithTimeout(timeout time.Duration) *QueryEdgeNetworksParams {
	return &QueryEdgeNetworksParams{
		timeout: timeout,
	}
}

// NewQueryEdgeNetworksParamsWithContext creates a new QueryEdgeNetworksParams object
// with the ability to set a context for a request.
func NewQueryEdgeNetworksParamsWithContext(ctx context.Context) *QueryEdgeNetworksParams {
	return &QueryEdgeNetworksParams{
		Context: ctx,
	}
}

// NewQueryEdgeNetworksParamsWithHTTPClient creates a new QueryEdgeNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryEdgeNetworksParamsWithHTTPClient(client *http.Client) *QueryEdgeNetworksParams {
	return &QueryEdgeNetworksParams{
		HTTPClient: client,
	}
}

/* QueryEdgeNetworksParams contains all the parameters to send to the API endpoint
   for the query edge networks operation.

   Typically these are written to a http.Request.
*/
type QueryEdgeNetworksParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// FilterNamePattern.
	FilterNamePattern *string

	// FilterProjectName.
	FilterProjectName *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy *string

	/* NextPageNum.

	   Page Number

	   Format: int64
	*/
	NextPageNum *int64

	/* NextPageSize.

	   Defines the page size

	   Format: int64
	*/
	NextPageSize *int64

	/* NextPageToken.

	   Page Token
	*/
	NextPageToken *string

	/* NextTotalPages.

	   Total number of pages to be fetched.

	   Format: int64
	*/
	NextTotalPages *int64

	// Summary.
	//
	// Format: boolean
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query edge networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeNetworksParams) WithDefaults() *QueryEdgeNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query edge networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryEdgeNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query edge networks params
func (o *QueryEdgeNetworksParams) WithTimeout(timeout time.Duration) *QueryEdgeNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query edge networks params
func (o *QueryEdgeNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query edge networks params
func (o *QueryEdgeNetworksParams) WithContext(ctx context.Context) *QueryEdgeNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query edge networks params
func (o *QueryEdgeNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query edge networks params
func (o *QueryEdgeNetworksParams) WithHTTPClient(client *http.Client) *QueryEdgeNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query edge networks params
func (o *QueryEdgeNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the query edge networks params
func (o *QueryEdgeNetworksParams) WithXRequestID(xRequestID *string) *QueryEdgeNetworksParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the query edge networks params
func (o *QueryEdgeNetworksParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithFilterNamePattern adds the filterNamePattern to the query edge networks params
func (o *QueryEdgeNetworksParams) WithFilterNamePattern(filterNamePattern *string) *QueryEdgeNetworksParams {
	o.SetFilterNamePattern(filterNamePattern)
	return o
}

// SetFilterNamePattern adds the filterNamePattern to the query edge networks params
func (o *QueryEdgeNetworksParams) SetFilterNamePattern(filterNamePattern *string) {
	o.FilterNamePattern = filterNamePattern
}

// WithFilterProjectName adds the filterProjectName to the query edge networks params
func (o *QueryEdgeNetworksParams) WithFilterProjectName(filterProjectName *string) *QueryEdgeNetworksParams {
	o.SetFilterProjectName(filterProjectName)
	return o
}

// SetFilterProjectName adds the filterProjectName to the query edge networks params
func (o *QueryEdgeNetworksParams) SetFilterProjectName(filterProjectName *string) {
	o.FilterProjectName = filterProjectName
}

// WithNextOrderBy adds the nextOrderBy to the query edge networks params
func (o *QueryEdgeNetworksParams) WithNextOrderBy(nextOrderBy *string) *QueryEdgeNetworksParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the query edge networks params
func (o *QueryEdgeNetworksParams) SetNextOrderBy(nextOrderBy *string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the query edge networks params
func (o *QueryEdgeNetworksParams) WithNextPageNum(nextPageNum *int64) *QueryEdgeNetworksParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the query edge networks params
func (o *QueryEdgeNetworksParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the query edge networks params
func (o *QueryEdgeNetworksParams) WithNextPageSize(nextPageSize *int64) *QueryEdgeNetworksParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the query edge networks params
func (o *QueryEdgeNetworksParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the query edge networks params
func (o *QueryEdgeNetworksParams) WithNextPageToken(nextPageToken *string) *QueryEdgeNetworksParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the query edge networks params
func (o *QueryEdgeNetworksParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the query edge networks params
func (o *QueryEdgeNetworksParams) WithNextTotalPages(nextTotalPages *int64) *QueryEdgeNetworksParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the query edge networks params
func (o *QueryEdgeNetworksParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithSummary adds the summary to the query edge networks params
func (o *QueryEdgeNetworksParams) WithSummary(summary *bool) *QueryEdgeNetworksParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the query edge networks params
func (o *QueryEdgeNetworksParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *QueryEdgeNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterNamePattern != nil {

		// query param filter.namePattern
		var qrFilterNamePattern string

		if o.FilterNamePattern != nil {
			qrFilterNamePattern = *o.FilterNamePattern
		}
		qFilterNamePattern := qrFilterNamePattern
		if qFilterNamePattern != "" {

			if err := r.SetQueryParam("filter.namePattern", qFilterNamePattern); err != nil {
				return err
			}
		}
	}

	if o.FilterProjectName != nil {

		// query param filter.projectName
		var qrFilterProjectName string

		if o.FilterProjectName != nil {
			qrFilterProjectName = *o.FilterProjectName
		}
		qFilterProjectName := qrFilterProjectName
		if qFilterProjectName != "" {

			if err := r.SetQueryParam("filter.projectName", qFilterProjectName); err != nil {
				return err
			}
		}
	}

	if o.NextOrderBy != nil {

		// query param next.orderBy
		var qrNextOrderBy string

		if o.NextOrderBy != nil {
			qrNextOrderBy = *o.NextOrderBy
		}
		qNextOrderBy := qrNextOrderBy
		if qNextOrderBy != "" {

			if err := r.SetQueryParam("next.orderBy", qNextOrderBy); err != nil {
				return err
			}
		}
	}

	if o.NextPageNum != nil {

		// query param next.pageNum
		var qrNextPageNum int64

		if o.NextPageNum != nil {
			qrNextPageNum = *o.NextPageNum
		}
		qNextPageNum := swag.FormatInt64(qrNextPageNum)
		if qNextPageNum != "" {

			if err := r.SetQueryParam("next.pageNum", qNextPageNum); err != nil {
				return err
			}
		}
	}

	if o.NextPageSize != nil {

		// query param next.pageSize
		var qrNextPageSize int64

		if o.NextPageSize != nil {
			qrNextPageSize = *o.NextPageSize
		}
		qNextPageSize := swag.FormatInt64(qrNextPageSize)
		if qNextPageSize != "" {

			if err := r.SetQueryParam("next.pageSize", qNextPageSize); err != nil {
				return err
			}
		}
	}

	if o.NextPageToken != nil {

		// query param next.pageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("next.pageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.NextTotalPages != nil {

		// query param next.totalPages
		var qrNextTotalPages int64

		if o.NextTotalPages != nil {
			qrNextTotalPages = *o.NextTotalPages
		}
		qNextTotalPages := swag.FormatInt64(qrNextTotalPages)
		if qNextTotalPages != "" {

			if err := r.SetQueryParam("next.totalPages", qNextTotalPages); err != nil {
				return err
			}
		}
	}

	if o.Summary != nil {

		// query param summary
		var qrSummary bool

		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {

			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}