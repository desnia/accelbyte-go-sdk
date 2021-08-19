// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// UpdateRootRegionOverrideReader is a Reader for the UpdateRootRegionOverride structure.
type UpdateRootRegionOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRootRegionOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRootRegionOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRootRegionOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRootRegionOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateRootRegionOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateRootRegionOverrideInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateRootRegionOverrideOK creates a UpdateRootRegionOverrideOK with default headers values
func NewUpdateRootRegionOverrideOK() *UpdateRootRegionOverrideOK {
	return &UpdateRootRegionOverrideOK{}
}

/*UpdateRootRegionOverrideOK handles this case with default header values.

  deployment region override updated
*/
type UpdateRootRegionOverrideOK struct {
	Payload *dsmcclientmodels.ModelsDeploymentWithOverride
}

func (o *UpdateRootRegionOverrideOK) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] updateRootRegionOverrideOK  %+v", 200, o.Payload)
}

func (o *UpdateRootRegionOverrideOK) GetPayload() *dsmcclientmodels.ModelsDeploymentWithOverride {
	return o.Payload
}

func (o *UpdateRootRegionOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsDeploymentWithOverride)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRootRegionOverrideBadRequest creates a UpdateRootRegionOverrideBadRequest with default headers values
func NewUpdateRootRegionOverrideBadRequest() *UpdateRootRegionOverrideBadRequest {
	return &UpdateRootRegionOverrideBadRequest{}
}

/*UpdateRootRegionOverrideBadRequest handles this case with default header values.

  malformed request
*/
type UpdateRootRegionOverrideBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateRootRegionOverrideBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] updateRootRegionOverrideBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRootRegionOverrideBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateRootRegionOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRootRegionOverrideUnauthorized creates a UpdateRootRegionOverrideUnauthorized with default headers values
func NewUpdateRootRegionOverrideUnauthorized() *UpdateRootRegionOverrideUnauthorized {
	return &UpdateRootRegionOverrideUnauthorized{}
}

/*UpdateRootRegionOverrideUnauthorized handles this case with default header values.

  Unauthorized
*/
type UpdateRootRegionOverrideUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateRootRegionOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] updateRootRegionOverrideUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRootRegionOverrideUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateRootRegionOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRootRegionOverrideNotFound creates a UpdateRootRegionOverrideNotFound with default headers values
func NewUpdateRootRegionOverrideNotFound() *UpdateRootRegionOverrideNotFound {
	return &UpdateRootRegionOverrideNotFound{}
}

/*UpdateRootRegionOverrideNotFound handles this case with default header values.

  deployment not found
*/
type UpdateRootRegionOverrideNotFound struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateRootRegionOverrideNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] updateRootRegionOverrideNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRootRegionOverrideNotFound) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateRootRegionOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRootRegionOverrideInternalServerError creates a UpdateRootRegionOverrideInternalServerError with default headers values
func NewUpdateRootRegionOverrideInternalServerError() *UpdateRootRegionOverrideInternalServerError {
	return &UpdateRootRegionOverrideInternalServerError{}
}

/*UpdateRootRegionOverrideInternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateRootRegionOverrideInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *UpdateRootRegionOverrideInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /dsmcontroller/admin/namespaces/{namespace}/configs/deployments/{deployment}/overrides/regions/{region}][%d] updateRootRegionOverrideInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRootRegionOverrideInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *UpdateRootRegionOverrideInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}