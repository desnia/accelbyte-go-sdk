// Code generated by go-swagger; DO NOT EDIT.

package users

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

// PublicUserVerificationV3Reader is a Reader for the PublicUserVerificationV3 structure.
type PublicUserVerificationV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicUserVerificationV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicUserVerificationV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicUserVerificationV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPublicUserVerificationV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicUserVerificationV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewPublicUserVerificationV3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicUserVerificationV3NoContent creates a PublicUserVerificationV3NoContent with default headers values
func NewPublicUserVerificationV3NoContent() *PublicUserVerificationV3NoContent {
	return &PublicUserVerificationV3NoContent{}
}

/*PublicUserVerificationV3NoContent handles this case with default header values.

  Operation succeeded
*/
type PublicUserVerificationV3NoContent struct {
}

func (o *PublicUserVerificationV3NoContent) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify][%d] publicUserVerificationV3NoContent ", 204)
}

func (o *PublicUserVerificationV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPublicUserVerificationV3BadRequest creates a PublicUserVerificationV3BadRequest with default headers values
func NewPublicUserVerificationV3BadRequest() *PublicUserVerificationV3BadRequest {
	return &PublicUserVerificationV3BadRequest{}
}

/*PublicUserVerificationV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PublicUserVerificationV3BadRequest struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicUserVerificationV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify][%d] publicUserVerificationV3BadRequest  %+v", 400, o.Payload)
}

func (o *PublicUserVerificationV3BadRequest) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicUserVerificationV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUserVerificationV3Unauthorized creates a PublicUserVerificationV3Unauthorized with default headers values
func NewPublicUserVerificationV3Unauthorized() *PublicUserVerificationV3Unauthorized {
	return &PublicUserVerificationV3Unauthorized{}
}

/*PublicUserVerificationV3Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr><tr><td>20022</td><td>token is not user token</td></tr></table>
*/
type PublicUserVerificationV3Unauthorized struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicUserVerificationV3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify][%d] publicUserVerificationV3Unauthorized  %+v", 401, o.Payload)
}

func (o *PublicUserVerificationV3Unauthorized) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicUserVerificationV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUserVerificationV3Forbidden creates a PublicUserVerificationV3Forbidden with default headers values
func NewPublicUserVerificationV3Forbidden() *PublicUserVerificationV3Forbidden {
	return &PublicUserVerificationV3Forbidden{}
}

/*PublicUserVerificationV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10152</td><td>verification code not found</td></tr><tr><td>10137</td><td>code is expired</td></tr><tr><td>10136</td><td>code is either been used or not valid anymore</td></tr><tr><td>10138</td><td>code not match</td></tr><tr><td>10149</td><td>verification contact type doesn't match</td></tr><tr><td>10148</td><td>verification code context doesn't match the required context</td></tr></table>
*/
type PublicUserVerificationV3Forbidden struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicUserVerificationV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify][%d] publicUserVerificationV3Forbidden  %+v", 403, o.Payload)
}

func (o *PublicUserVerificationV3Forbidden) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicUserVerificationV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicUserVerificationV3Conflict creates a PublicUserVerificationV3Conflict with default headers values
func NewPublicUserVerificationV3Conflict() *PublicUserVerificationV3Conflict {
	return &PublicUserVerificationV3Conflict{}
}

/*PublicUserVerificationV3Conflict handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10141</td><td>email verified</td></tr></table>
*/
type PublicUserVerificationV3Conflict struct {
	Payload *iamclientmodels.RestErrorResponse
}

func (o *PublicUserVerificationV3Conflict) Error() string {
	return fmt.Sprintf("[POST /iam/v3/public/namespaces/{namespace}/users/me/code/verify][%d] publicUserVerificationV3Conflict  %+v", 409, o.Payload)
}

func (o *PublicUserVerificationV3Conflict) GetPayload() *iamclientmodels.RestErrorResponse {
	return o.Payload
}

func (o *PublicUserVerificationV3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.RestErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}