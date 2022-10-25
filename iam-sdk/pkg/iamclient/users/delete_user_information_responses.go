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

// DeleteUserInformationReader is a Reader for the DeleteUserInformation structure.
type DeleteUserInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserInformationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserInformationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteUserInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserInformationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/namespaces/{namespace}/users/{userId}/information returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteUserInformationNoContent creates a DeleteUserInformationNoContent with default headers values
func NewDeleteUserInformationNoContent() *DeleteUserInformationNoContent {
	return &DeleteUserInformationNoContent{}
}

/*DeleteUserInformationNoContent handles this case with default header values.

  Operation succeeded
*/
type DeleteUserInformationNoContent struct {
}

func (o *DeleteUserInformationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/users/{userId}/information][%d] deleteUserInformationNoContent ", 204)
}

func (o *DeleteUserInformationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteUserInformationUnauthorized creates a DeleteUserInformationUnauthorized with default headers values
func NewDeleteUserInformationUnauthorized() *DeleteUserInformationUnauthorized {
	return &DeleteUserInformationUnauthorized{}
}

/*DeleteUserInformationUnauthorized handles this case with default header values.

  Unauthorized access
*/
type DeleteUserInformationUnauthorized struct {
}

func (o *DeleteUserInformationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/users/{userId}/information][%d] deleteUserInformationUnauthorized ", 401)
}

func (o *DeleteUserInformationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteUserInformationForbidden creates a DeleteUserInformationForbidden with default headers values
func NewDeleteUserInformationForbidden() *DeleteUserInformationForbidden {
	return &DeleteUserInformationForbidden{}
}

/*DeleteUserInformationForbidden handles this case with default header values.

  Forbidden
*/
type DeleteUserInformationForbidden struct {
}

func (o *DeleteUserInformationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/users/{userId}/information][%d] deleteUserInformationForbidden ", 403)
}

func (o *DeleteUserInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteUserInformationNotFound creates a DeleteUserInformationNotFound with default headers values
func NewDeleteUserInformationNotFound() *DeleteUserInformationNotFound {
	return &DeleteUserInformationNotFound{}
}

/*DeleteUserInformationNotFound handles this case with default header values.

  Data not found
*/
type DeleteUserInformationNotFound struct {
}

func (o *DeleteUserInformationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/namespaces/{namespace}/users/{userId}/information][%d] deleteUserInformationNotFound ", 404)
}

func (o *DeleteUserInformationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
