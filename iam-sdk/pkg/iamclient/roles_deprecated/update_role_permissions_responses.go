// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package roles_deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRolePermissionsReader is a Reader for the UpdateRolePermissions structure.
type UpdateRolePermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRolePermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRolePermissionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRolePermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRolePermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRolePermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateRolePermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/roles/{roleId}/permissions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateRolePermissionsNoContent creates a UpdateRolePermissionsNoContent with default headers values
func NewUpdateRolePermissionsNoContent() *UpdateRolePermissionsNoContent {
	return &UpdateRolePermissionsNoContent{}
}

/*UpdateRolePermissionsNoContent handles this case with default header values.

  Operation succeeded
*/
type UpdateRolePermissionsNoContent struct {
}

func (o *UpdateRolePermissionsNoContent) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/permissions][%d] updateRolePermissionsNoContent ", 204)
}

func (o *UpdateRolePermissionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRolePermissionsBadRequest creates a UpdateRolePermissionsBadRequest with default headers values
func NewUpdateRolePermissionsBadRequest() *UpdateRolePermissionsBadRequest {
	return &UpdateRolePermissionsBadRequest{}
}

/*UpdateRolePermissionsBadRequest handles this case with default header values.

  Invalid request
*/
type UpdateRolePermissionsBadRequest struct {
}

func (o *UpdateRolePermissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/permissions][%d] updateRolePermissionsBadRequest ", 400)
}

func (o *UpdateRolePermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRolePermissionsUnauthorized creates a UpdateRolePermissionsUnauthorized with default headers values
func NewUpdateRolePermissionsUnauthorized() *UpdateRolePermissionsUnauthorized {
	return &UpdateRolePermissionsUnauthorized{}
}

/*UpdateRolePermissionsUnauthorized handles this case with default header values.

  Unauthorized access
*/
type UpdateRolePermissionsUnauthorized struct {
}

func (o *UpdateRolePermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/permissions][%d] updateRolePermissionsUnauthorized ", 401)
}

func (o *UpdateRolePermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRolePermissionsForbidden creates a UpdateRolePermissionsForbidden with default headers values
func NewUpdateRolePermissionsForbidden() *UpdateRolePermissionsForbidden {
	return &UpdateRolePermissionsForbidden{}
}

/*UpdateRolePermissionsForbidden handles this case with default header values.

  Forbidden
*/
type UpdateRolePermissionsForbidden struct {
}

func (o *UpdateRolePermissionsForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/permissions][%d] updateRolePermissionsForbidden ", 403)
}

func (o *UpdateRolePermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRolePermissionsNotFound creates a UpdateRolePermissionsNotFound with default headers values
func NewUpdateRolePermissionsNotFound() *UpdateRolePermissionsNotFound {
	return &UpdateRolePermissionsNotFound{}
}

/*UpdateRolePermissionsNotFound handles this case with default header values.

  Data not found
*/
type UpdateRolePermissionsNotFound struct {
}

func (o *UpdateRolePermissionsNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/roles/{roleId}/permissions][%d] updateRolePermissionsNotFound ", 404)
}

func (o *UpdateRolePermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}