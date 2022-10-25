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

// EnableUserReader is a Reader for the EnableUser structure.
type EnableUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnableUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewEnableUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewEnableUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /iam/namespaces/{namespace}/users/{userId}/enable returns an error %d: %s", response.Code(), string(data))
	}
}

// NewEnableUserNoContent creates a EnableUserNoContent with default headers values
func NewEnableUserNoContent() *EnableUserNoContent {
	return &EnableUserNoContent{}
}

/*EnableUserNoContent handles this case with default header values.

  Operation succeeded
*/
type EnableUserNoContent struct {
}

func (o *EnableUserNoContent) Error() string {
	return fmt.Sprintf("[PUT /iam/namespaces/{namespace}/users/{userId}/enable][%d] enableUserNoContent ", 204)
}

func (o *EnableUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewEnableUserUnauthorized creates a EnableUserUnauthorized with default headers values
func NewEnableUserUnauthorized() *EnableUserUnauthorized {
	return &EnableUserUnauthorized{}
}

/*EnableUserUnauthorized handles this case with default header values.

  Unauthorized access
*/
type EnableUserUnauthorized struct {
}

func (o *EnableUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /iam/namespaces/{namespace}/users/{userId}/enable][%d] enableUserUnauthorized ", 401)
}

func (o *EnableUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewEnableUserForbidden creates a EnableUserForbidden with default headers values
func NewEnableUserForbidden() *EnableUserForbidden {
	return &EnableUserForbidden{}
}

/*EnableUserForbidden handles this case with default header values.

  Forbidden
*/
type EnableUserForbidden struct {
}

func (o *EnableUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /iam/namespaces/{namespace}/users/{userId}/enable][%d] enableUserForbidden ", 403)
}

func (o *EnableUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewEnableUserNotFound creates a EnableUserNotFound with default headers values
func NewEnableUserNotFound() *EnableUserNotFound {
	return &EnableUserNotFound{}
}

/*EnableUserNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type EnableUserNotFound struct {
}

func (o *EnableUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /iam/namespaces/{namespace}/users/{userId}/enable][%d] enableUserNotFound ", 404)
}

func (o *EnableUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewEnableUserInternalServerError creates a EnableUserInternalServerError with default headers values
func NewEnableUserInternalServerError() *EnableUserInternalServerError {
	return &EnableUserInternalServerError{}
}

/*EnableUserInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type EnableUserInternalServerError struct {
}

func (o *EnableUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /iam/namespaces/{namespace}/users/{userId}/enable][%d] enableUserInternalServerError ", 500)
}

func (o *EnableUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
