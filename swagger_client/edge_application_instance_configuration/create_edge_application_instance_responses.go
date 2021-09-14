// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_application_instance_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateEdgeApplicationInstanceReader is a Reader for the CreateEdgeApplicationInstance structure.
type CreateEdgeApplicationInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEdgeApplicationInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEdgeApplicationInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEdgeApplicationInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEdgeApplicationInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEdgeApplicationInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEdgeApplicationInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEdgeApplicationInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateEdgeApplicationInstanceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEdgeApplicationInstanceOK creates a CreateEdgeApplicationInstanceOK with default headers values
func NewCreateEdgeApplicationInstanceOK() *CreateEdgeApplicationInstanceOK {
	return &CreateEdgeApplicationInstanceOK{}
}

/* CreateEdgeApplicationInstanceOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEdgeApplicationInstanceOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceOK) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceOK  %+v", 200, o.Payload)
}
func (o *CreateEdgeApplicationInstanceOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceBadRequest creates a CreateEdgeApplicationInstanceBadRequest with default headers values
func NewCreateEdgeApplicationInstanceBadRequest() *CreateEdgeApplicationInstanceBadRequest {
	return &CreateEdgeApplicationInstanceBadRequest{}
}

/* CreateEdgeApplicationInstanceBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateEdgeApplicationInstanceBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEdgeApplicationInstanceBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceUnauthorized creates a CreateEdgeApplicationInstanceUnauthorized with default headers values
func NewCreateEdgeApplicationInstanceUnauthorized() *CreateEdgeApplicationInstanceUnauthorized {
	return &CreateEdgeApplicationInstanceUnauthorized{}
}

/* CreateEdgeApplicationInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateEdgeApplicationInstanceUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEdgeApplicationInstanceUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceForbidden creates a CreateEdgeApplicationInstanceForbidden with default headers values
func NewCreateEdgeApplicationInstanceForbidden() *CreateEdgeApplicationInstanceForbidden {
	return &CreateEdgeApplicationInstanceForbidden{}
}

/* CreateEdgeApplicationInstanceForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type CreateEdgeApplicationInstanceForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceForbidden  %+v", 403, o.Payload)
}
func (o *CreateEdgeApplicationInstanceForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceConflict creates a CreateEdgeApplicationInstanceConflict with default headers values
func NewCreateEdgeApplicationInstanceConflict() *CreateEdgeApplicationInstanceConflict {
	return &CreateEdgeApplicationInstanceConflict{}
}

/* CreateEdgeApplicationInstanceConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge application instance record will conflict with an already existing edge application iinstance record.
*/
type CreateEdgeApplicationInstanceConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceConflict  %+v", 409, o.Payload)
}
func (o *CreateEdgeApplicationInstanceConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceInternalServerError creates a CreateEdgeApplicationInstanceInternalServerError with default headers values
func NewCreateEdgeApplicationInstanceInternalServerError() *CreateEdgeApplicationInstanceInternalServerError {
	return &CreateEdgeApplicationInstanceInternalServerError{}
}

/* CreateEdgeApplicationInstanceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateEdgeApplicationInstanceInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateEdgeApplicationInstanceInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeApplicationInstanceGatewayTimeout creates a CreateEdgeApplicationInstanceGatewayTimeout with default headers values
func NewCreateEdgeApplicationInstanceGatewayTimeout() *CreateEdgeApplicationInstanceGatewayTimeout {
	return &CreateEdgeApplicationInstanceGatewayTimeout{}
}

/* CreateEdgeApplicationInstanceGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateEdgeApplicationInstanceGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeApplicationInstanceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/apps/instances][%d] createEdgeApplicationInstanceGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateEdgeApplicationInstanceGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeApplicationInstanceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}