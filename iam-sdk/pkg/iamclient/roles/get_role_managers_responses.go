// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// GetRoleManagersReader is a Reader for the GetRoleManagers structure.
type GetRoleManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleManagersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoleManagersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoleManagersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoleManagersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoleManagersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/roles/{roleId}/managers returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetRoleManagersOK creates a GetRoleManagersOK with default headers values
func NewGetRoleManagersOK() *GetRoleManagersOK {
	return &GetRoleManagersOK{}
}

/*GetRoleManagersOK handles this case with default header values.

  OK
*/
type GetRoleManagersOK struct {
	Payload *iamclientmodels.ModelRoleManagersResponse
}

func (o *GetRoleManagersOK) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/managers][%d] getRoleManagersOK  %+v", 200, o.ToJSONString())
}

func (o *GetRoleManagersOK) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetRoleManagersOK) GetPayload() *iamclientmodels.ModelRoleManagersResponse {
	return o.Payload
}

func (o *GetRoleManagersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(iamclientmodels.ModelRoleManagersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleManagersBadRequest creates a GetRoleManagersBadRequest with default headers values
func NewGetRoleManagersBadRequest() *GetRoleManagersBadRequest {
	return &GetRoleManagersBadRequest{}
}

/*GetRoleManagersBadRequest handles this case with default header values.

  Invalid request
*/
type GetRoleManagersBadRequest struct {
}

func (o *GetRoleManagersBadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/managers][%d] getRoleManagersBadRequest ", 400)
}

func (o *GetRoleManagersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewGetRoleManagersUnauthorized creates a GetRoleManagersUnauthorized with default headers values
func NewGetRoleManagersUnauthorized() *GetRoleManagersUnauthorized {
	return &GetRoleManagersUnauthorized{}
}

/*GetRoleManagersUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetRoleManagersUnauthorized struct {
}

func (o *GetRoleManagersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/managers][%d] getRoleManagersUnauthorized ", 401)
}

func (o *GetRoleManagersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewGetRoleManagersForbidden creates a GetRoleManagersForbidden with default headers values
func NewGetRoleManagersForbidden() *GetRoleManagersForbidden {
	return &GetRoleManagersForbidden{}
}

/*GetRoleManagersForbidden handles this case with default header values.

  Forbidden
*/
type GetRoleManagersForbidden struct {
}

func (o *GetRoleManagersForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/managers][%d] getRoleManagersForbidden ", 403)
}

func (o *GetRoleManagersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewGetRoleManagersNotFound creates a GetRoleManagersNotFound with default headers values
func NewGetRoleManagersNotFound() *GetRoleManagersNotFound {
	return &GetRoleManagersNotFound{}
}

/*GetRoleManagersNotFound handles this case with default header values.

  Data not found
*/
type GetRoleManagersNotFound struct {
}

func (o *GetRoleManagersNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/roles/{roleId}/managers][%d] getRoleManagersNotFound ", 404)
}

func (o *GetRoleManagersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
