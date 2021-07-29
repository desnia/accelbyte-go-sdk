// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminVerifyAccountV3Reader is a Reader for the AdminVerifyAccountV3 structure.
type AdminVerifyAccountV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminVerifyAccountV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminVerifyAccountV3NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminVerifyAccountV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminVerifyAccountV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminVerifyAccountV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminVerifyAccountV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminVerifyAccountV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminVerifyAccountV3NoContent creates a AdminVerifyAccountV3NoContent with default headers values
func NewAdminVerifyAccountV3NoContent() *AdminVerifyAccountV3NoContent {
	return &AdminVerifyAccountV3NoContent{}
}

/*AdminVerifyAccountV3NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminVerifyAccountV3NoContent struct {
}

func (o *AdminVerifyAccountV3NoContent) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3NoContent ", 204)
}

func (o *AdminVerifyAccountV3NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminVerifyAccountV3BadRequest creates a AdminVerifyAccountV3BadRequest with default headers values
func NewAdminVerifyAccountV3BadRequest() *AdminVerifyAccountV3BadRequest {
	return &AdminVerifyAccountV3BadRequest{}
}

/*AdminVerifyAccountV3BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type AdminVerifyAccountV3BadRequest struct {
}

func (o *AdminVerifyAccountV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3BadRequest ", 400)
}

func (o *AdminVerifyAccountV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminVerifyAccountV3Unauthorized creates a AdminVerifyAccountV3Unauthorized with default headers values
func NewAdminVerifyAccountV3Unauthorized() *AdminVerifyAccountV3Unauthorized {
	return &AdminVerifyAccountV3Unauthorized{}
}

/*AdminVerifyAccountV3Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminVerifyAccountV3Unauthorized struct {
}

func (o *AdminVerifyAccountV3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3Unauthorized ", 401)
}

func (o *AdminVerifyAccountV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminVerifyAccountV3Forbidden creates a AdminVerifyAccountV3Forbidden with default headers values
func NewAdminVerifyAccountV3Forbidden() *AdminVerifyAccountV3Forbidden {
	return &AdminVerifyAccountV3Forbidden{}
}

/*AdminVerifyAccountV3Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10152</td><td>verification code not found</td></tr><tr><td>10137</td><td>code is expired</td></tr><tr><td>10136</td><td>code is either been used or not valid anymore</td></tr><tr><td>10138</td><td>code not match</td></tr><tr><td>10149</td><td>verification contact type doesn't match</td></tr><tr><td>10148</td><td>verification code context doesn't match the required context</td></tr><tr><td>10162</td><td>invalid verification</td></tr></table>
*/
type AdminVerifyAccountV3Forbidden struct {
}

func (o *AdminVerifyAccountV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3Forbidden ", 403)
}

func (o *AdminVerifyAccountV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminVerifyAccountV3NotFound creates a AdminVerifyAccountV3NotFound with default headers values
func NewAdminVerifyAccountV3NotFound() *AdminVerifyAccountV3NotFound {
	return &AdminVerifyAccountV3NotFound{}
}

/*AdminVerifyAccountV3NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>10139</td><td>platform account not found</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type AdminVerifyAccountV3NotFound struct {
}

func (o *AdminVerifyAccountV3NotFound) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3NotFound ", 404)
}

func (o *AdminVerifyAccountV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminVerifyAccountV3InternalServerError creates a AdminVerifyAccountV3InternalServerError with default headers values
func NewAdminVerifyAccountV3InternalServerError() *AdminVerifyAccountV3InternalServerError {
	return &AdminVerifyAccountV3InternalServerError{}
}

/*AdminVerifyAccountV3InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type AdminVerifyAccountV3InternalServerError struct {
}

func (o *AdminVerifyAccountV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /iam/v3/admin/namespaces/{namespace}/users/{userId}/code/verify][%d] adminVerifyAccountV3InternalServerError ", 500)
}

func (o *AdminVerifyAccountV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}