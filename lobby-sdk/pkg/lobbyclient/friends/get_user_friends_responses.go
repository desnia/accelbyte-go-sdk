// Code generated by go-swagger; DO NOT EDIT.

package friends

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// GetUserFriendsReader is a Reader for the GetUserFriends structure.
type GetUserFriendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserFriendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserFriendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserFriendsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserFriendsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserFriendsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserFriendsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUserFriendsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /friends/namespaces/{namespace}/me returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserFriendsOK creates a GetUserFriendsOK with default headers values
func NewGetUserFriendsOK() *GetUserFriendsOK {
	return &GetUserFriendsOK{}
}

/*GetUserFriendsOK handles this case with default header values.

  OK
*/
type GetUserFriendsOK struct {
	Payload []*lobbyclientmodels.ModelGetUserFriendsResponse
}

func (o *GetUserFriendsOK) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsOK  %+v", 200, o.Payload)
}

func (o *GetUserFriendsOK) GetPayload() []*lobbyclientmodels.ModelGetUserFriendsResponse {
	return o.Payload
}

func (o *GetUserFriendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFriendsBadRequest creates a GetUserFriendsBadRequest with default headers values
func NewGetUserFriendsBadRequest() *GetUserFriendsBadRequest {
	return &GetUserFriendsBadRequest{}
}

/*GetUserFriendsBadRequest handles this case with default header values.

  Bad Request
*/
type GetUserFriendsBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *GetUserFriendsBadRequest) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserFriendsBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *GetUserFriendsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFriendsUnauthorized creates a GetUserFriendsUnauthorized with default headers values
func NewGetUserFriendsUnauthorized() *GetUserFriendsUnauthorized {
	return &GetUserFriendsUnauthorized{}
}

/*GetUserFriendsUnauthorized handles this case with default header values.

  Unauthorized
*/
type GetUserFriendsUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *GetUserFriendsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserFriendsUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *GetUserFriendsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFriendsForbidden creates a GetUserFriendsForbidden with default headers values
func NewGetUserFriendsForbidden() *GetUserFriendsForbidden {
	return &GetUserFriendsForbidden{}
}

/*GetUserFriendsForbidden handles this case with default header values.

  Forbidden
*/
type GetUserFriendsForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *GetUserFriendsForbidden) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserFriendsForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *GetUserFriendsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFriendsNotFound creates a GetUserFriendsNotFound with default headers values
func NewGetUserFriendsNotFound() *GetUserFriendsNotFound {
	return &GetUserFriendsNotFound{}
}

/*GetUserFriendsNotFound handles this case with default header values.

  Not Found
*/
type GetUserFriendsNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *GetUserFriendsNotFound) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserFriendsNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *GetUserFriendsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserFriendsInternalServerError creates a GetUserFriendsInternalServerError with default headers values
func NewGetUserFriendsInternalServerError() *GetUserFriendsInternalServerError {
	return &GetUserFriendsInternalServerError{}
}

/*GetUserFriendsInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetUserFriendsInternalServerError struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *GetUserFriendsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /friends/namespaces/{namespace}/me][%d] getUserFriendsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserFriendsInternalServerError) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *GetUserFriendsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}