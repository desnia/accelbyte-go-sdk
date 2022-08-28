// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveRoleAdminReader is a Reader for the RemoveRoleAdmin structure.
type RemoveRoleAdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveRoleAdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveRoleAdminNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveRoleAdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveRoleAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRemoveRoleAdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRemoveRoleAdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/roles/{roleId}/admin returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRemoveRoleAdminNoContent creates a RemoveRoleAdminNoContent with default headers values
func NewRemoveRoleAdminNoContent() *RemoveRoleAdminNoContent {
	return &RemoveRoleAdminNoContent{}
}

/*RemoveRoleAdminNoContent handles this case with default header values.

  Operation succeeded
*/
type RemoveRoleAdminNoContent struct {
}

func (o *RemoveRoleAdminNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/admin][%d] removeRoleAdminNoContent ", 204)
}

func (o *RemoveRoleAdminNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleAdminBadRequest creates a RemoveRoleAdminBadRequest with default headers values
func NewRemoveRoleAdminBadRequest() *RemoveRoleAdminBadRequest {
	return &RemoveRoleAdminBadRequest{}
}

/*RemoveRoleAdminBadRequest handles this case with default header values.

  Invalid request
*/
type RemoveRoleAdminBadRequest struct {
}

func (o *RemoveRoleAdminBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/admin][%d] removeRoleAdminBadRequest ", 400)
}

func (o *RemoveRoleAdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleAdminUnauthorized creates a RemoveRoleAdminUnauthorized with default headers values
func NewRemoveRoleAdminUnauthorized() *RemoveRoleAdminUnauthorized {
	return &RemoveRoleAdminUnauthorized{}
}

/*RemoveRoleAdminUnauthorized handles this case with default header values.

  Unauthorized access
*/
type RemoveRoleAdminUnauthorized struct {
}

func (o *RemoveRoleAdminUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/admin][%d] removeRoleAdminUnauthorized ", 401)
}

func (o *RemoveRoleAdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleAdminForbidden creates a RemoveRoleAdminForbidden with default headers values
func NewRemoveRoleAdminForbidden() *RemoveRoleAdminForbidden {
	return &RemoveRoleAdminForbidden{}
}

/*RemoveRoleAdminForbidden handles this case with default header values.

  Forbidden
*/
type RemoveRoleAdminForbidden struct {
}

func (o *RemoveRoleAdminForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/admin][%d] removeRoleAdminForbidden ", 403)
}

func (o *RemoveRoleAdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleAdminNotFound creates a RemoveRoleAdminNotFound with default headers values
func NewRemoveRoleAdminNotFound() *RemoveRoleAdminNotFound {
	return &RemoveRoleAdminNotFound{}
}

/*RemoveRoleAdminNotFound handles this case with default header values.

  Data not found
*/
type RemoveRoleAdminNotFound struct {
}

func (o *RemoveRoleAdminNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/admin][%d] removeRoleAdminNotFound ", 404)
}

func (o *RemoveRoleAdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}