// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// QueryGlobalHardwareBrandsReader is a Reader for the QueryGlobalHardwareBrands structure.
type QueryGlobalHardwareBrandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryGlobalHardwareBrandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryGlobalHardwareBrandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryGlobalHardwareBrandsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryGlobalHardwareBrandsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryGlobalHardwareBrandsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryGlobalHardwareBrandsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewQueryGlobalHardwareBrandsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryGlobalHardwareBrandsOK creates a QueryGlobalHardwareBrandsOK with default headers values
func NewQueryGlobalHardwareBrandsOK() *QueryGlobalHardwareBrandsOK {
	return &QueryGlobalHardwareBrandsOK{}
}

/* QueryGlobalHardwareBrandsOK describes a response with status code 200, with default header values.

A successful response.
*/
type QueryGlobalHardwareBrandsOK struct {
	Payload *swagger_models.SysBrands
}

func (o *QueryGlobalHardwareBrandsOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsOK  %+v", 200, o.Payload)
}
func (o *QueryGlobalHardwareBrandsOK) GetPayload() *swagger_models.SysBrands {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.SysBrands)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGlobalHardwareBrandsBadRequest creates a QueryGlobalHardwareBrandsBadRequest with default headers values
func NewQueryGlobalHardwareBrandsBadRequest() *QueryGlobalHardwareBrandsBadRequest {
	return &QueryGlobalHardwareBrandsBadRequest{}
}

/* QueryGlobalHardwareBrandsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type QueryGlobalHardwareBrandsBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryGlobalHardwareBrandsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsBadRequest  %+v", 400, o.Payload)
}
func (o *QueryGlobalHardwareBrandsBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGlobalHardwareBrandsUnauthorized creates a QueryGlobalHardwareBrandsUnauthorized with default headers values
func NewQueryGlobalHardwareBrandsUnauthorized() *QueryGlobalHardwareBrandsUnauthorized {
	return &QueryGlobalHardwareBrandsUnauthorized{}
}

/* QueryGlobalHardwareBrandsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type QueryGlobalHardwareBrandsUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryGlobalHardwareBrandsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryGlobalHardwareBrandsUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGlobalHardwareBrandsForbidden creates a QueryGlobalHardwareBrandsForbidden with default headers values
func NewQueryGlobalHardwareBrandsForbidden() *QueryGlobalHardwareBrandsForbidden {
	return &QueryGlobalHardwareBrandsForbidden{}
}

/* QueryGlobalHardwareBrandsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type QueryGlobalHardwareBrandsForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryGlobalHardwareBrandsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsForbidden  %+v", 403, o.Payload)
}
func (o *QueryGlobalHardwareBrandsForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGlobalHardwareBrandsInternalServerError creates a QueryGlobalHardwareBrandsInternalServerError with default headers values
func NewQueryGlobalHardwareBrandsInternalServerError() *QueryGlobalHardwareBrandsInternalServerError {
	return &QueryGlobalHardwareBrandsInternalServerError{}
}

/* QueryGlobalHardwareBrandsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type QueryGlobalHardwareBrandsInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryGlobalHardwareBrandsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryGlobalHardwareBrandsInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryGlobalHardwareBrandsGatewayTimeout creates a QueryGlobalHardwareBrandsGatewayTimeout with default headers values
func NewQueryGlobalHardwareBrandsGatewayTimeout() *QueryGlobalHardwareBrandsGatewayTimeout {
	return &QueryGlobalHardwareBrandsGatewayTimeout{}
}

/* QueryGlobalHardwareBrandsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type QueryGlobalHardwareBrandsGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *QueryGlobalHardwareBrandsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global][%d] queryGlobalHardwareBrandsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *QueryGlobalHardwareBrandsGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *QueryGlobalHardwareBrandsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}