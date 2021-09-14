// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package volume_instance_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/zedcloud-api/swagger_models"
)

// GetVolumeInstanceStatusByNameReader is a Reader for the GetVolumeInstanceStatusByName structure.
type GetVolumeInstanceStatusByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeInstanceStatusByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVolumeInstanceStatusByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVolumeInstanceStatusByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVolumeInstanceStatusByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVolumeInstanceStatusByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVolumeInstanceStatusByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVolumeInstanceStatusByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVolumeInstanceStatusByNameOK creates a GetVolumeInstanceStatusByNameOK with default headers values
func NewGetVolumeInstanceStatusByNameOK() *GetVolumeInstanceStatusByNameOK {
	return &GetVolumeInstanceStatusByNameOK{}
}

/* GetVolumeInstanceStatusByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetVolumeInstanceStatusByNameOK struct {
	Payload *swagger_models.VolInstStatusMsg
}

func (o *GetVolumeInstanceStatusByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameOK  %+v", 200, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameOK) GetPayload() *swagger_models.VolInstStatusMsg {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.VolInstStatusMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeInstanceStatusByNameUnauthorized creates a GetVolumeInstanceStatusByNameUnauthorized with default headers values
func NewGetVolumeInstanceStatusByNameUnauthorized() *GetVolumeInstanceStatusByNameUnauthorized {
	return &GetVolumeInstanceStatusByNameUnauthorized{}
}

/* GetVolumeInstanceStatusByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type GetVolumeInstanceStatusByNameUnauthorized struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetVolumeInstanceStatusByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameUnauthorized) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeInstanceStatusByNameForbidden creates a GetVolumeInstanceStatusByNameForbidden with default headers values
func NewGetVolumeInstanceStatusByNameForbidden() *GetVolumeInstanceStatusByNameForbidden {
	return &GetVolumeInstanceStatusByNameForbidden{}
}

/* GetVolumeInstanceStatusByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type GetVolumeInstanceStatusByNameForbidden struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetVolumeInstanceStatusByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameForbidden  %+v", 403, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameForbidden) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeInstanceStatusByNameNotFound creates a GetVolumeInstanceStatusByNameNotFound with default headers values
func NewGetVolumeInstanceStatusByNameNotFound() *GetVolumeInstanceStatusByNameNotFound {
	return &GetVolumeInstanceStatusByNameNotFound{}
}

/* GetVolumeInstanceStatusByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetVolumeInstanceStatusByNameNotFound struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetVolumeInstanceStatusByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameNotFound) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeInstanceStatusByNameInternalServerError creates a GetVolumeInstanceStatusByNameInternalServerError with default headers values
func NewGetVolumeInstanceStatusByNameInternalServerError() *GetVolumeInstanceStatusByNameInternalServerError {
	return &GetVolumeInstanceStatusByNameInternalServerError{}
}

/* GetVolumeInstanceStatusByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type GetVolumeInstanceStatusByNameInternalServerError struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetVolumeInstanceStatusByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameInternalServerError) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeInstanceStatusByNameGatewayTimeout creates a GetVolumeInstanceStatusByNameGatewayTimeout with default headers values
func NewGetVolumeInstanceStatusByNameGatewayTimeout() *GetVolumeInstanceStatusByNameGatewayTimeout {
	return &GetVolumeInstanceStatusByNameGatewayTimeout{}
}

/* GetVolumeInstanceStatusByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type GetVolumeInstanceStatusByNameGatewayTimeout struct {
	Payload *swagger_models.ZsrvResponse
}

func (o *GetVolumeInstanceStatusByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}/status][%d] getVolumeInstanceStatusByNameGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetVolumeInstanceStatusByNameGatewayTimeout) GetPayload() *swagger_models.ZsrvResponse {
	return o.Payload
}

func (o *GetVolumeInstanceStatusByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(swagger_models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}