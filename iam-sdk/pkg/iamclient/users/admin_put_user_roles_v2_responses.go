// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminPutUserRolesV2Reader is a Reader for the AdminPutUserRolesV2 structure.
type AdminPutUserRolesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminPutUserRolesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminPutUserRolesV2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminPutUserRolesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminPutUserRolesV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminPutUserRolesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminPutUserRolesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminPutUserRolesV2NoContent creates a AdminPutUserRolesV2NoContent with default headers values
func NewAdminPutUserRolesV2NoContent() *AdminPutUserRolesV2NoContent {
	return &AdminPutUserRolesV2NoContent{}
}

/*AdminPutUserRolesV2NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminPutUserRolesV2NoContent struct {
}

func (o *AdminPutUserRolesV2NoContent) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminPutUserRolesV2NoContent ", 204)
}

func (o *AdminPutUserRolesV2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewAdminPutUserRolesV2BadRequest creates a AdminPutUserRolesV2BadRequest with default headers values
func NewAdminPutUserRolesV2BadRequest() *AdminPutUserRolesV2BadRequest {
	return &AdminPutUserRolesV2BadRequest{}
}

/*AdminPutUserRolesV2BadRequest handles this case with default header values.

  Invalid request
*/
type AdminPutUserRolesV2BadRequest struct {
}

func (o *AdminPutUserRolesV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminPutUserRolesV2BadRequest ", 400)
}

func (o *AdminPutUserRolesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewAdminPutUserRolesV2Unauthorized creates a AdminPutUserRolesV2Unauthorized with default headers values
func NewAdminPutUserRolesV2Unauthorized() *AdminPutUserRolesV2Unauthorized {
	return &AdminPutUserRolesV2Unauthorized{}
}

/*AdminPutUserRolesV2Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminPutUserRolesV2Unauthorized struct {
}

func (o *AdminPutUserRolesV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminPutUserRolesV2Unauthorized ", 401)
}

func (o *AdminPutUserRolesV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewAdminPutUserRolesV2Forbidden creates a AdminPutUserRolesV2Forbidden with default headers values
func NewAdminPutUserRolesV2Forbidden() *AdminPutUserRolesV2Forbidden {
	return &AdminPutUserRolesV2Forbidden{}
}

/*AdminPutUserRolesV2Forbidden handles this case with default header values.

  Forbidden
*/
type AdminPutUserRolesV2Forbidden struct {
}

func (o *AdminPutUserRolesV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminPutUserRolesV2Forbidden ", 403)
}

func (o *AdminPutUserRolesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewAdminPutUserRolesV2NotFound creates a AdminPutUserRolesV2NotFound with default headers values
func NewAdminPutUserRolesV2NotFound() *AdminPutUserRolesV2NotFound {
	return &AdminPutUserRolesV2NotFound{}
}

/*AdminPutUserRolesV2NotFound handles this case with default header values.

  Data not found
*/
type AdminPutUserRolesV2NotFound struct {
}

func (o *AdminPutUserRolesV2NotFound) Error() string {
	return fmt.Sprintf("[PUT /iam/v2/admin/namespaces/{namespace}/users/{userId}/roles][%d] adminPutUserRolesV2NotFound ", 404)
}

func (o *AdminPutUserRolesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
