// Code generated by go-swagger; DO NOT EDIT.

package data_deletion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/gdpr-sdk/pkg/gdprclientmodels"
)

// AdminSubmitUserAccountDeletionRequestReader is a Reader for the AdminSubmitUserAccountDeletionRequest structure.
type AdminSubmitUserAccountDeletionRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminSubmitUserAccountDeletionRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAdminSubmitUserAccountDeletionRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminSubmitUserAccountDeletionRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminSubmitUserAccountDeletionRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminSubmitUserAccountDeletionRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewAdminSubmitUserAccountDeletionRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminSubmitUserAccountDeletionRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminSubmitUserAccountDeletionRequestCreated creates a AdminSubmitUserAccountDeletionRequestCreated with default headers values
func NewAdminSubmitUserAccountDeletionRequestCreated() *AdminSubmitUserAccountDeletionRequestCreated {
	return &AdminSubmitUserAccountDeletionRequestCreated{}
}

/*AdminSubmitUserAccountDeletionRequestCreated handles this case with default header values.

  Created
*/
type AdminSubmitUserAccountDeletionRequestCreated struct {
	Payload *gdprclientmodels.ModelsRequestDeleteResponse
}

func (o *AdminSubmitUserAccountDeletionRequestCreated) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestCreated  %+v", 201, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestCreated) GetPayload() *gdprclientmodels.ModelsRequestDeleteResponse {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ModelsRequestDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSubmitUserAccountDeletionRequestUnauthorized creates a AdminSubmitUserAccountDeletionRequestUnauthorized with default headers values
func NewAdminSubmitUserAccountDeletionRequestUnauthorized() *AdminSubmitUserAccountDeletionRequestUnauthorized {
	return &AdminSubmitUserAccountDeletionRequestUnauthorized{}
}

/*AdminSubmitUserAccountDeletionRequestUnauthorized handles this case with default header values.

  Unauthorized
*/
type AdminSubmitUserAccountDeletionRequestUnauthorized struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminSubmitUserAccountDeletionRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestUnauthorized) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSubmitUserAccountDeletionRequestForbidden creates a AdminSubmitUserAccountDeletionRequestForbidden with default headers values
func NewAdminSubmitUserAccountDeletionRequestForbidden() *AdminSubmitUserAccountDeletionRequestForbidden {
	return &AdminSubmitUserAccountDeletionRequestForbidden{}
}

/*AdminSubmitUserAccountDeletionRequestForbidden handles this case with default header values.

  Forbidden
*/
type AdminSubmitUserAccountDeletionRequestForbidden struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminSubmitUserAccountDeletionRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestForbidden  %+v", 403, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestForbidden) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSubmitUserAccountDeletionRequestNotFound creates a AdminSubmitUserAccountDeletionRequestNotFound with default headers values
func NewAdminSubmitUserAccountDeletionRequestNotFound() *AdminSubmitUserAccountDeletionRequestNotFound {
	return &AdminSubmitUserAccountDeletionRequestNotFound{}
}

/*AdminSubmitUserAccountDeletionRequestNotFound handles this case with default header values.

  Not Found
*/
type AdminSubmitUserAccountDeletionRequestNotFound struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminSubmitUserAccountDeletionRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestNotFound  %+v", 404, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestNotFound) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSubmitUserAccountDeletionRequestConflict creates a AdminSubmitUserAccountDeletionRequestConflict with default headers values
func NewAdminSubmitUserAccountDeletionRequestConflict() *AdminSubmitUserAccountDeletionRequestConflict {
	return &AdminSubmitUserAccountDeletionRequestConflict{}
}

/*AdminSubmitUserAccountDeletionRequestConflict handles this case with default header values.

  Conflict
*/
type AdminSubmitUserAccountDeletionRequestConflict struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminSubmitUserAccountDeletionRequestConflict) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestConflict  %+v", 409, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestConflict) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSubmitUserAccountDeletionRequestInternalServerError creates a AdminSubmitUserAccountDeletionRequestInternalServerError with default headers values
func NewAdminSubmitUserAccountDeletionRequestInternalServerError() *AdminSubmitUserAccountDeletionRequestInternalServerError {
	return &AdminSubmitUserAccountDeletionRequestInternalServerError{}
}

/*AdminSubmitUserAccountDeletionRequestInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminSubmitUserAccountDeletionRequestInternalServerError struct {
	Payload *gdprclientmodels.ResponseError
}

func (o *AdminSubmitUserAccountDeletionRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gdpr/admin/namespaces/{namespace}/users/{userId}/deletions][%d] adminSubmitUserAccountDeletionRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminSubmitUserAccountDeletionRequestInternalServerError) GetPayload() *gdprclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminSubmitUserAccountDeletionRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gdprclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}