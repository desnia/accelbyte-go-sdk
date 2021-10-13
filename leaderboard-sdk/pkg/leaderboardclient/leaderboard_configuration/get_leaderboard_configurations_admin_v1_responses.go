// Code generated by go-swagger; DO NOT EDIT.

package leaderboard_configuration

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

// GetLeaderboardConfigurationsAdminV1Reader is a Reader for the GetLeaderboardConfigurationsAdminV1 structure.
type GetLeaderboardConfigurationsAdminV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeaderboardConfigurationsAdminV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeaderboardConfigurationsAdminV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLeaderboardConfigurationsAdminV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLeaderboardConfigurationsAdminV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLeaderboardConfigurationsAdminV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetLeaderboardConfigurationsAdminV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetLeaderboardConfigurationsAdminV1OK creates a GetLeaderboardConfigurationsAdminV1OK with default headers values
func NewGetLeaderboardConfigurationsAdminV1OK() *GetLeaderboardConfigurationsAdminV1OK {
	return &GetLeaderboardConfigurationsAdminV1OK{}
}

/*GetLeaderboardConfigurationsAdminV1OK handles this case with default header values.

  OK
*/
type GetLeaderboardConfigurationsAdminV1OK struct {
	Payload *leaderboardclientmodels.ModelsGetAllLeaderboardConfigsResp
}

func (o *GetLeaderboardConfigurationsAdminV1OK) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards][%d] getLeaderboardConfigurationsAdminV1OK  %+v", 200, o.Payload)
}

func (o *GetLeaderboardConfigurationsAdminV1OK) GetPayload() *leaderboardclientmodels.ModelsGetAllLeaderboardConfigsResp {
	return o.Payload
}

func (o *GetLeaderboardConfigurationsAdminV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ModelsGetAllLeaderboardConfigsResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLeaderboardConfigurationsAdminV1BadRequest creates a GetLeaderboardConfigurationsAdminV1BadRequest with default headers values
func NewGetLeaderboardConfigurationsAdminV1BadRequest() *GetLeaderboardConfigurationsAdminV1BadRequest {
	return &GetLeaderboardConfigurationsAdminV1BadRequest{}
}

/*GetLeaderboardConfigurationsAdminV1BadRequest handles this case with default header values.

  Bad Request
*/
type GetLeaderboardConfigurationsAdminV1BadRequest struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetLeaderboardConfigurationsAdminV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards][%d] getLeaderboardConfigurationsAdminV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetLeaderboardConfigurationsAdminV1BadRequest) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetLeaderboardConfigurationsAdminV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLeaderboardConfigurationsAdminV1Unauthorized creates a GetLeaderboardConfigurationsAdminV1Unauthorized with default headers values
func NewGetLeaderboardConfigurationsAdminV1Unauthorized() *GetLeaderboardConfigurationsAdminV1Unauthorized {
	return &GetLeaderboardConfigurationsAdminV1Unauthorized{}
}

/*GetLeaderboardConfigurationsAdminV1Unauthorized handles this case with default header values.

  Unauthorized
*/
type GetLeaderboardConfigurationsAdminV1Unauthorized struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetLeaderboardConfigurationsAdminV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards][%d] getLeaderboardConfigurationsAdminV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetLeaderboardConfigurationsAdminV1Unauthorized) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetLeaderboardConfigurationsAdminV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLeaderboardConfigurationsAdminV1Forbidden creates a GetLeaderboardConfigurationsAdminV1Forbidden with default headers values
func NewGetLeaderboardConfigurationsAdminV1Forbidden() *GetLeaderboardConfigurationsAdminV1Forbidden {
	return &GetLeaderboardConfigurationsAdminV1Forbidden{}
}

/*GetLeaderboardConfigurationsAdminV1Forbidden handles this case with default header values.

  Forbidden
*/
type GetLeaderboardConfigurationsAdminV1Forbidden struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetLeaderboardConfigurationsAdminV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards][%d] getLeaderboardConfigurationsAdminV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetLeaderboardConfigurationsAdminV1Forbidden) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetLeaderboardConfigurationsAdminV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLeaderboardConfigurationsAdminV1InternalServerError creates a GetLeaderboardConfigurationsAdminV1InternalServerError with default headers values
func NewGetLeaderboardConfigurationsAdminV1InternalServerError() *GetLeaderboardConfigurationsAdminV1InternalServerError {
	return &GetLeaderboardConfigurationsAdminV1InternalServerError{}
}

/*GetLeaderboardConfigurationsAdminV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetLeaderboardConfigurationsAdminV1InternalServerError struct {
	Payload *leaderboardclientmodels.ResponseErrorResponse
}

func (o *GetLeaderboardConfigurationsAdminV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /leaderboard/v1/admin/namespaces/{namespace}/leaderboards][%d] getLeaderboardConfigurationsAdminV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetLeaderboardConfigurationsAdminV1InternalServerError) GetPayload() *leaderboardclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *GetLeaderboardConfigurationsAdminV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(leaderboardclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}