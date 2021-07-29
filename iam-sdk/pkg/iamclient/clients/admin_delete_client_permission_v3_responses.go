// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// AdminDeleteClientPermissionV3Reader is a Reader for the AdminDeleteClientPermissionV3 structure.
type AdminDeleteClientPermissionV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteClientPermissionV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeleteClientPermissionV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminDeleteClientPermissionV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteClientPermissionV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminDeleteClientPermissionV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminDeleteClientPermissionV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeleteClientPermissionV3NoContent creates a AdminDeleteClientPermissionV3NoContent with default headers values
func NewAdminDeleteClientPermissionV3NoContent() *AdminDeleteClientPermissionV3NoContent {
	return &AdminDeleteClientPermissionV3NoContent{}
}

/*AdminDeleteClientPermissionV3NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminDeleteClientPermissionV3NoContent struct {
}

func (o *AdminDeleteClientPermissionV3NoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action}][%d] adminDeleteClientPermissionV3NoContent ", 204)
}

func (o *AdminDeleteClientPermissionV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeleteClientPermissionV3BadRequest creates a AdminDeleteClientPermissionV3BadRequest with default headers values
func NewAdminDeleteClientPermissionV3BadRequest() *AdminDeleteClientPermissionV3BadRequest {
	return &AdminDeleteClientPermissionV3BadRequest{}
}

/*AdminDeleteClientPermissionV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type AdminDeleteClientPermissionV3BadRequest struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminDeleteClientPermissionV3BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action}][%d] adminDeleteClientPermissionV3BadRequest  %+v", 400, o.Payload)
}

func (o *AdminDeleteClientPermissionV3BadRequest) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminDeleteClientPermissionV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteClientPermissionV3Unauthorized creates a AdminDeleteClientPermissionV3Unauthorized with default headers values
func NewAdminDeleteClientPermissionV3Unauthorized() *AdminDeleteClientPermissionV3Unauthorized {
	return &AdminDeleteClientPermissionV3Unauthorized{}
}

/*AdminDeleteClientPermissionV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type AdminDeleteClientPermissionV3Unauthorized struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminDeleteClientPermissionV3Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action}][%d] adminDeleteClientPermissionV3Unauthorized  %+v", 401, o.Payload)
}

func (o *AdminDeleteClientPermissionV3Unauthorized) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminDeleteClientPermissionV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteClientPermissionV3Forbidden creates a AdminDeleteClientPermissionV3Forbidden with default headers values
func NewAdminDeleteClientPermissionV3Forbidden() *AdminDeleteClientPermissionV3Forbidden {
	return &AdminDeleteClientPermissionV3Forbidden{}
}

/*AdminDeleteClientPermissionV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr></table>
*/
type AdminDeleteClientPermissionV3Forbidden struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminDeleteClientPermissionV3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action}][%d] adminDeleteClientPermissionV3Forbidden  %+v", 403, o.Payload)
}

func (o *AdminDeleteClientPermissionV3Forbidden) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminDeleteClientPermissionV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteClientPermissionV3NotFound creates a AdminDeleteClientPermissionV3NotFound with default headers values
func NewAdminDeleteClientPermissionV3NotFound() *AdminDeleteClientPermissionV3NotFound {
	return &AdminDeleteClientPermissionV3NotFound{}
}

/*AdminDeleteClientPermissionV3NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10365</td><td>client not found</td></tr></table>
*/
type AdminDeleteClientPermissionV3NotFound struct {
	Payload *iamclientmodels.RestapiErrorResponse
}

func (o *AdminDeleteClientPermissionV3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/v3/admin/namespaces/{namespace}/clients/{clientId}/permissions/{resource}/{action}][%d] adminDeleteClientPermissionV3NotFound  %+v", 404, o.Payload)
}

func (o *AdminDeleteClientPermissionV3NotFound) GetPayload() *iamclientmodels.RestapiErrorResponse {
	return o.Payload
}

func (o *AdminDeleteClientPermissionV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestapiErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}