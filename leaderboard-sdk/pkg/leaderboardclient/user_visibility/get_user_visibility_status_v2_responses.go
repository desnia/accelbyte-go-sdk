// Code generated by go-swagger; DO NOT EDIT.

package user_visibility

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/leaderboard-sdk/pkg/leaderboardclientmodels"
)

// GetUserVisibilityStatusV2Reader is a Reader for the GetUserVisibilityStatusV2 structure.
type GetUserVisibilityStatusV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserVisibilityStatusV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserVisibilityStatusV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserVisibilityStatusV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserVisibilityStatusV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserVisibilityStatusV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserVisibilityStatusV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUserVisibilityStatusV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserVisibilityStatusV2OK creates a GetUserVisibilityStatusV2OK with default headers values
func NewGetUserVisibilityStatusV2OK() *GetUserVisibilityStatusV2OK {
	return &GetUserVisibilityStatusV2OK{}
}

/*GetUserVisibilityStatusV2OK handles this case with default header values.

  OK
*/
type GetUserVisibilityStatusV2OK struct {
	Payload *leaderboardclientmodels.ModelsGetUserVisibilityResponse
}

func (o *GetUserVisibilityStatusV2OK) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2OK  %+v", 200, o.Payload)
}

func (o *GetUserVisibilityStatusV2OK) GetPayload() *leaderboardclientmodels.ModelsGetUserVisibilityResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ModelsGetUserVisibilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVisibilityStatusV2BadRequest creates a GetUserVisibilityStatusV2BadRequest with default headers values
func NewGetUserVisibilityStatusV2BadRequest() *GetUserVisibilityStatusV2BadRequest {
	return &GetUserVisibilityStatusV2BadRequest{}
}

/*GetUserVisibilityStatusV2BadRequest handles this case with default header values.

  Bad Request
*/
type GetUserVisibilityStatusV2BadRequest struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetUserVisibilityStatusV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetUserVisibilityStatusV2BadRequest) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVisibilityStatusV2Unauthorized creates a GetUserVisibilityStatusV2Unauthorized with default headers values
func NewGetUserVisibilityStatusV2Unauthorized() *GetUserVisibilityStatusV2Unauthorized {
	return &GetUserVisibilityStatusV2Unauthorized{}
}

/*GetUserVisibilityStatusV2Unauthorized handles this case with default header values.

  Unauthorized
*/
type GetUserVisibilityStatusV2Unauthorized struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetUserVisibilityStatusV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2Unauthorized  %+v", 401, o.Payload)
}

func (o *GetUserVisibilityStatusV2Unauthorized) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVisibilityStatusV2Forbidden creates a GetUserVisibilityStatusV2Forbidden with default headers values
func NewGetUserVisibilityStatusV2Forbidden() *GetUserVisibilityStatusV2Forbidden {
	return &GetUserVisibilityStatusV2Forbidden{}
}

/*GetUserVisibilityStatusV2Forbidden handles this case with default header values.

  Forbidden
*/
type GetUserVisibilityStatusV2Forbidden struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetUserVisibilityStatusV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetUserVisibilityStatusV2Forbidden) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVisibilityStatusV2NotFound creates a GetUserVisibilityStatusV2NotFound with default headers values
func NewGetUserVisibilityStatusV2NotFound() *GetUserVisibilityStatusV2NotFound {
	return &GetUserVisibilityStatusV2NotFound{}
}

/*GetUserVisibilityStatusV2NotFound handles this case with default header values.

  Not Found
*/
type GetUserVisibilityStatusV2NotFound struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetUserVisibilityStatusV2NotFound) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2NotFound  %+v", 404, o.Payload)
}

func (o *GetUserVisibilityStatusV2NotFound) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVisibilityStatusV2InternalServerError creates a GetUserVisibilityStatusV2InternalServerError with default headers values
func NewGetUserVisibilityStatusV2InternalServerError() *GetUserVisibilityStatusV2InternalServerError {
	return &GetUserVisibilityStatusV2InternalServerError{}
}

/*GetUserVisibilityStatusV2InternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetUserVisibilityStatusV2InternalServerError struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetUserVisibilityStatusV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v2/admin/namespaces/{namespace}/leaderboards/{leaderboardCode}/users/{userId}/visibility][%d] getUserVisibilityStatusV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserVisibilityStatusV2InternalServerError) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetUserVisibilityStatusV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}