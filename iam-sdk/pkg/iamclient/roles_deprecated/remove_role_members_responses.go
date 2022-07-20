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

// RemoveRoleMembersReader is a Reader for the RemoveRoleMembers structure.
type RemoveRoleMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveRoleMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveRoleMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveRoleMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveRoleMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRemoveRoleMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRemoveRoleMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/roles/{roleId}/members returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRemoveRoleMembersNoContent creates a RemoveRoleMembersNoContent with default headers values
func NewRemoveRoleMembersNoContent() *RemoveRoleMembersNoContent {
	return &RemoveRoleMembersNoContent{}
}

/*RemoveRoleMembersNoContent handles this case with default header values.

  Operation succeeded
*/
type RemoveRoleMembersNoContent struct {
}

func (o *RemoveRoleMembersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/members][%d] removeRoleMembersNoContent ", 204)
}

func (o *RemoveRoleMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleMembersBadRequest creates a RemoveRoleMembersBadRequest with default headers values
func NewRemoveRoleMembersBadRequest() *RemoveRoleMembersBadRequest {
	return &RemoveRoleMembersBadRequest{}
}

/*RemoveRoleMembersBadRequest handles this case with default header values.

  Invalid request
*/
type RemoveRoleMembersBadRequest struct {
}

func (o *RemoveRoleMembersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/members][%d] removeRoleMembersBadRequest ", 400)
}

func (o *RemoveRoleMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleMembersUnauthorized creates a RemoveRoleMembersUnauthorized with default headers values
func NewRemoveRoleMembersUnauthorized() *RemoveRoleMembersUnauthorized {
	return &RemoveRoleMembersUnauthorized{}
}

/*RemoveRoleMembersUnauthorized handles this case with default header values.

  Unauthorized access
*/
type RemoveRoleMembersUnauthorized struct {
}

func (o *RemoveRoleMembersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/members][%d] removeRoleMembersUnauthorized ", 401)
}

func (o *RemoveRoleMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleMembersForbidden creates a RemoveRoleMembersForbidden with default headers values
func NewRemoveRoleMembersForbidden() *RemoveRoleMembersForbidden {
	return &RemoveRoleMembersForbidden{}
}

/*RemoveRoleMembersForbidden handles this case with default header values.

  Forbidden
*/
type RemoveRoleMembersForbidden struct {
}

func (o *RemoveRoleMembersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/members][%d] removeRoleMembersForbidden ", 403)
}

func (o *RemoveRoleMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleMembersNotFound creates a RemoveRoleMembersNotFound with default headers values
func NewRemoveRoleMembersNotFound() *RemoveRoleMembersNotFound {
	return &RemoveRoleMembersNotFound{}
}

/*RemoveRoleMembersNotFound handles this case with default header values.

  Data not found
*/
type RemoveRoleMembersNotFound struct {
}

func (o *RemoveRoleMembersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/roles/{roleId}/members][%d] removeRoleMembersNotFound ", 404)
}

func (o *RemoveRoleMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}