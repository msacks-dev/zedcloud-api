// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// CreateEdgeNodeReader is a Reader for the CreateEdgeNode structure.
type CreateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEdgeNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCreateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEdgeNodeOK creates a CreateEdgeNodeOK with default headers values
func NewCreateEdgeNodeOK() *CreateEdgeNodeOK {
	return &CreateEdgeNodeOK{}
}

/* CreateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateEdgeNodeOK struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeOK  %+v", 200, o.Payload)
}
func (o *CreateEdgeNodeOK) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeBadRequest creates a CreateEdgeNodeBadRequest with default headers values
func NewCreateEdgeNodeBadRequest() *CreateEdgeNodeBadRequest {
	return &CreateEdgeNodeBadRequest{}
}

/* CreateEdgeNodeBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type CreateEdgeNodeBadRequest struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEdgeNodeBadRequest) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeUnauthorized creates a CreateEdgeNodeUnauthorized with default headers values
func NewCreateEdgeNodeUnauthorized() *CreateEdgeNodeUnauthorized {
	return &CreateEdgeNodeUnauthorized{}
}

/* CreateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type CreateEdgeNodeUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateEdgeNodeUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeForbidden creates a CreateEdgeNodeForbidden with default headers values
func NewCreateEdgeNodeForbidden() *CreateEdgeNodeForbidden {
	return &CreateEdgeNodeForbidden{}
}

/* CreateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type CreateEdgeNodeForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeForbidden  %+v", 403, o.Payload)
}
func (o *CreateEdgeNodeForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeConflict creates a CreateEdgeNodeConflict with default headers values
func NewCreateEdgeNodeConflict() *CreateEdgeNodeConflict {
	return &CreateEdgeNodeConflict{}
}

/* CreateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge node record will conflict with an already existing edge node record.
*/
type CreateEdgeNodeConflict struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeConflict  %+v", 409, o.Payload)
}
func (o *CreateEdgeNodeConflict) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeInternalServerError creates a CreateEdgeNodeInternalServerError with default headers values
func NewCreateEdgeNodeInternalServerError() *CreateEdgeNodeInternalServerError {
	return &CreateEdgeNodeInternalServerError{}
}

/* CreateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type CreateEdgeNodeInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateEdgeNodeInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEdgeNodeGatewayTimeout creates a CreateEdgeNodeGatewayTimeout with default headers values
func NewCreateEdgeNodeGatewayTimeout() *CreateEdgeNodeGatewayTimeout {
	return &CreateEdgeNodeGatewayTimeout{}
}

/* CreateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type CreateEdgeNodeGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *CreateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/devices][%d] createEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}
func (o *CreateEdgeNodeGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *CreateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}