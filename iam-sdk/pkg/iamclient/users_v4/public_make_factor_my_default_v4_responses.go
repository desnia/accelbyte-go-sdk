// Code generated by go-swagger; DO NOT EDIT.

package users_v4

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

// PublicMakeFactorMyDefaultV4Reader is a Reader for the PublicMakeFactorMyDefaultV4 structure.
type PublicMakeFactorMyDefaultV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicMakeFactorMyDefaultV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicMakeFactorMyDefaultV4NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicMakeFactorMyDefaultV4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicMakeFactorMyDefaultV4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicMakeFactorMyDefaultV4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicMakeFactorMyDefaultV4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPublicMakeFactorMyDefaultV4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicMakeFactorMyDefaultV4NoContent creates a PublicMakeFactorMyDefaultV4NoContent with default headers values
func NewPublicMakeFactorMyDefaultV4NoContent() *PublicMakeFactorMyDefaultV4NoContent {
	return &PublicMakeFactorMyDefaultV4NoContent{}
}

/*PublicMakeFactorMyDefaultV4NoContent handles this case with default header values.

  Default method changed
*/
type PublicMakeFactorMyDefaultV4NoContent struct {
}

func (o *PublicMakeFactorMyDefaultV4NoContent) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4NoContent ", 204)
}

func (o *PublicMakeFactorMyDefaultV4NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicMakeFactorMyDefaultV4BadRequest creates a PublicMakeFactorMyDefaultV4BadRequest with default headers values
func NewPublicMakeFactorMyDefaultV4BadRequest() *PublicMakeFactorMyDefaultV4BadRequest {
	return &PublicMakeFactorMyDefaultV4BadRequest{}
}

/*PublicMakeFactorMyDefaultV4BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10191</td><td>email address not verified</td></tr><tr><td>10189</td><td>invalid factor</td></tr><tr><td>10192</td><td>factor not enabled</td></tr><tr><td>10171</td><td>email address not found</td></tr></table>
*/
type PublicMakeFactorMyDefaultV4BadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicMakeFactorMyDefaultV4BadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4BadRequest  %+v", 400, o.Payload)
}

func (o *PublicMakeFactorMyDefaultV4BadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicMakeFactorMyDefaultV4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicMakeFactorMyDefaultV4Unauthorized creates a PublicMakeFactorMyDefaultV4Unauthorized with default headers values
func NewPublicMakeFactorMyDefaultV4Unauthorized() *PublicMakeFactorMyDefaultV4Unauthorized {
	return &PublicMakeFactorMyDefaultV4Unauthorized{}
}

/*PublicMakeFactorMyDefaultV4Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type PublicMakeFactorMyDefaultV4Unauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicMakeFactorMyDefaultV4Unauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4Unauthorized  %+v", 401, o.Payload)
}

func (o *PublicMakeFactorMyDefaultV4Unauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicMakeFactorMyDefaultV4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicMakeFactorMyDefaultV4Forbidden creates a PublicMakeFactorMyDefaultV4Forbidden with default headers values
func NewPublicMakeFactorMyDefaultV4Forbidden() *PublicMakeFactorMyDefaultV4Forbidden {
	return &PublicMakeFactorMyDefaultV4Forbidden{}
}

/*PublicMakeFactorMyDefaultV4Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20003</td><td>forbidden access</td></tr></table>
*/
type PublicMakeFactorMyDefaultV4Forbidden struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicMakeFactorMyDefaultV4Forbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4Forbidden  %+v", 403, o.Payload)
}

func (o *PublicMakeFactorMyDefaultV4Forbidden) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicMakeFactorMyDefaultV4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicMakeFactorMyDefaultV4NotFound creates a PublicMakeFactorMyDefaultV4NotFound with default headers values
func NewPublicMakeFactorMyDefaultV4NotFound() *PublicMakeFactorMyDefaultV4NotFound {
	return &PublicMakeFactorMyDefaultV4NotFound{}
}

/*PublicMakeFactorMyDefaultV4NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10139</td><td>platform account not found</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type PublicMakeFactorMyDefaultV4NotFound struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicMakeFactorMyDefaultV4NotFound) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4NotFound  %+v", 404, o.Payload)
}

func (o *PublicMakeFactorMyDefaultV4NotFound) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicMakeFactorMyDefaultV4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicMakeFactorMyDefaultV4InternalServerError creates a PublicMakeFactorMyDefaultV4InternalServerError with default headers values
func NewPublicMakeFactorMyDefaultV4InternalServerError() *PublicMakeFactorMyDefaultV4InternalServerError {
	return &PublicMakeFactorMyDefaultV4InternalServerError{}
}

/*PublicMakeFactorMyDefaultV4InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type PublicMakeFactorMyDefaultV4InternalServerError struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicMakeFactorMyDefaultV4InternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/v4/public/namespaces/{namespace}/users/me/mfa/factor][%d] publicMakeFactorMyDefaultV4InternalServerError  %+v", 500, o.Payload)
}

func (o *PublicMakeFactorMyDefaultV4InternalServerError) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicMakeFactorMyDefaultV4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
