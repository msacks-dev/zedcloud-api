// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package datastore_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// QueryDatastoresReader is a Reader for the QueryDatastores structure.
type QueryDatastoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDatastoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDatastoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryDatastoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryDatastoresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryDatastoresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryDatastoresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryDatastoresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryDatastoresOK creates a QueryDatastoresOK with default headers values
func NewQueryDatastoresOK() *QueryDatastoresOK {
	return &QueryDatastoresOK{}
}

/* QueryDatastoresOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryDatastoresOK struct {
	Payload *swagger_models.Datastores
}

func (o *QueryDatastoresOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresOK  %+v", 200, o.Payload)
}
func (o *QueryDatastoresOK) GetPayload() *swagger_models.Datastores {
	return o.Payload
}

func (o *QueryDatastoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.Datastores)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatastoresBadRequest creates a QueryDatastoresBadRequest with default headers values
func NewQueryDatastoresBadRequest() *QueryDatastoresBadRequest {
	return &QueryDatastoresBadRequest{}
}

/* QueryDatastoresBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryDatastoresBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryDatastoresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresBadRequest  %+v", 400, o.Payload)
}
func (o *QueryDatastoresBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryDatastoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatastoresUnauthorized creates a QueryDatastoresUnauthorized with default headers values
func NewQueryDatastoresUnauthorized() *QueryDatastoresUnauthorized {
	return &QueryDatastoresUnauthorized{}
}

/* QueryDatastoresUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryDatastoresUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryDatastoresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryDatastoresUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryDatastoresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatastoresForbidden creates a QueryDatastoresForbidden with default headers values
func NewQueryDatastoresForbidden() *QueryDatastoresForbidden {
	return &QueryDatastoresForbidden{}
}

/* QueryDatastoresForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryDatastoresForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryDatastoresForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresForbidden  %+v", 403, o.Payload)
}
func (o *QueryDatastoresForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryDatastoresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatastoresInternalServerError creates a QueryDatastoresInternalServerError with default headers values
func NewQueryDatastoresInternalServerError() *QueryDatastoresInternalServerError {
	return &QueryDatastoresInternalServerError{}
}

/* QueryDatastoresInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryDatastoresInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryDatastoresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryDatastoresInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryDatastoresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatastoresGatewayTimeout creates a QueryDatastoresGatewayTimeout with default headers values
func NewQueryDatastoresGatewayTimeout() *QueryDatastoresGatewayTimeout {
	return &QueryDatastoresGatewayTimeout{}
}

/* QueryDatastoresGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryDatastoresGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryDatastoresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/datastores][%d] queryDatastoresGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryDatastoresGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryDatastoresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}