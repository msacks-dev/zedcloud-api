// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package resource_group_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetResourceGroupStatusByIDReader is a Reader for the GetResourceGroupStatusByID structure.
type GetResourceGroupStatusByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceGroupStatusByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceGroupStatusByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceGroupStatusByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceGroupStatusByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceGroupStatusByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceGroupStatusByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResourceGroupStatusByIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceGroupStatusByIDOK creates a GetResourceGroupStatusByIDOK with default headers values
func NewGetResourceGroupStatusByIDOK() *GetResourceGroupStatusByIDOK {
	return &GetResourceGroupStatusByIDOK{}
}

/* GetResourceGroupStatusByIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetResourceGroupStatusByIDOK struct {
	Payload *swagger_models.TagStatusMsg
}

func (o *GetResourceGroupStatusByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdOK  %+v", 200, o.Payload)
}
func (o *GetResourceGroupStatusByIDOK) GetPayload() *swagger_models.TagStatusMsg {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.TagStatusMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByIDUnauthorized creates a GetResourceGroupStatusByIDUnauthorized with default headers values
func NewGetResourceGroupStatusByIDUnauthorized() *GetResourceGroupStatusByIDUnauthorized {
	return &GetResourceGroupStatusByIDUnauthorized{}
}

/* GetResourceGroupStatusByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetResourceGroupStatusByIDUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdUnauthorized  %+v", 401, o.Payload)
}
func (o *GetResourceGroupStatusByIDUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByIDForbidden creates a GetResourceGroupStatusByIDForbidden with default headers values
func NewGetResourceGroupStatusByIDForbidden() *GetResourceGroupStatusByIDForbidden {
	return &GetResourceGroupStatusByIDForbidden{}
}

/* GetResourceGroupStatusByIDForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetResourceGroupStatusByIDForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdForbidden  %+v", 403, o.Payload)
}
func (o *GetResourceGroupStatusByIDForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByIDNotFound creates a GetResourceGroupStatusByIDNotFound with default headers values
func NewGetResourceGroupStatusByIDNotFound() *GetResourceGroupStatusByIDNotFound {
	return &GetResourceGroupStatusByIDNotFound{}
}

/* GetResourceGroupStatusByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetResourceGroupStatusByIDNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceGroupStatusByIDNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByIDInternalServerError creates a GetResourceGroupStatusByIDInternalServerError with default headers values
func NewGetResourceGroupStatusByIDInternalServerError() *GetResourceGroupStatusByIDInternalServerError {
	return &GetResourceGroupStatusByIDInternalServerError{}
}

/* GetResourceGroupStatusByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetResourceGroupStatusByIDInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetResourceGroupStatusByIDInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceGroupStatusByIDGatewayTimeout creates a GetResourceGroupStatusByIDGatewayTimeout with default headers values
func NewGetResourceGroupStatusByIDGatewayTimeout() *GetResourceGroupStatusByIDGatewayTimeout {
	return &GetResourceGroupStatusByIDGatewayTimeout{}
}

/* GetResourceGroupStatusByIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetResourceGroupStatusByIDGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetResourceGroupStatusByIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/projects/id/{id}/status][%d] getResourceGroupStatusByIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetResourceGroupStatusByIDGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetResourceGroupStatusByIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}