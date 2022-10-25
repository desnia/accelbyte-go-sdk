// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteClientPermissionReader is a Reader for the DeleteClientPermission structure.
type DeleteClientPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClientPermissionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClientPermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClientPermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteClientPermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteClientPermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteClientPermissionNoContent creates a DeleteClientPermissionNoContent with default headers values
func NewDeleteClientPermissionNoContent() *DeleteClientPermissionNoContent {
	return &DeleteClientPermissionNoContent{}
}

/*DeleteClientPermissionNoContent handles this case with default header values.

  Operation succeeded
*/
type DeleteClientPermissionNoContent struct {
}

func (o *DeleteClientPermissionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action}][%d] deleteClientPermissionNoContent ", 204)
}

func (o *DeleteClientPermissionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteClientPermissionBadRequest creates a DeleteClientPermissionBadRequest with default headers values
func NewDeleteClientPermissionBadRequest() *DeleteClientPermissionBadRequest {
	return &DeleteClientPermissionBadRequest{}
}

/*DeleteClientPermissionBadRequest handles this case with default header values.

  Invalid request
*/
type DeleteClientPermissionBadRequest struct {
}

func (o *DeleteClientPermissionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action}][%d] deleteClientPermissionBadRequest ", 400)
}

func (o *DeleteClientPermissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteClientPermissionUnauthorized creates a DeleteClientPermissionUnauthorized with default headers values
func NewDeleteClientPermissionUnauthorized() *DeleteClientPermissionUnauthorized {
	return &DeleteClientPermissionUnauthorized{}
}

/*DeleteClientPermissionUnauthorized handles this case with default header values.

  Unauthorized access
*/
type DeleteClientPermissionUnauthorized struct {
}

func (o *DeleteClientPermissionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action}][%d] deleteClientPermissionUnauthorized ", 401)
}

func (o *DeleteClientPermissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteClientPermissionForbidden creates a DeleteClientPermissionForbidden with default headers values
func NewDeleteClientPermissionForbidden() *DeleteClientPermissionForbidden {
	return &DeleteClientPermissionForbidden{}
}

/*DeleteClientPermissionForbidden handles this case with default header values.

  Forbidden
*/
type DeleteClientPermissionForbidden struct {
}

func (o *DeleteClientPermissionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action}][%d] deleteClientPermissionForbidden ", 403)
}

func (o *DeleteClientPermissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteClientPermissionNotFound creates a DeleteClientPermissionNotFound with default headers values
func NewDeleteClientPermissionNotFound() *DeleteClientPermissionNotFound {
	return &DeleteClientPermissionNotFound{}
}

/*DeleteClientPermissionNotFound handles this case with default header values.

  Data not found
*/
type DeleteClientPermissionNotFound struct {
}

func (o *DeleteClientPermissionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/clients/{clientId}/clientpermissions/{resource}/{action}][%d] deleteClientPermissionNotFound ", 404)
}

func (o *DeleteClientPermissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
