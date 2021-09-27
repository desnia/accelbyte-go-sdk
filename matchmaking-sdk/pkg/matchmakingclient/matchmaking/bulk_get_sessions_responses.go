// Code generated by go-swagger; DO NOT EDIT.

package matchmaking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/matchmaking-sdk/pkg/matchmakingclientmodels"
)

// BulkGetSessionsReader is a Reader for the BulkGetSessions structure.
type BulkGetSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkGetSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkGetSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkGetSessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBulkGetSessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBulkGetSessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewBulkGetSessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewBulkGetSessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkGetSessionsOK creates a BulkGetSessionsOK with default headers values
func NewBulkGetSessionsOK() *BulkGetSessionsOK {
	return &BulkGetSessionsOK{}
}

/*BulkGetSessionsOK handles this case with default header values.

  Operation succeeded
*/
type BulkGetSessionsOK struct {
	Payload []*matchmakingclientmodels.ModelsMatchmakingResult
}

func (o *BulkGetSessionsOK) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsOK  %+v", 200, o.Payload)
}

func (o *BulkGetSessionsOK) GetPayload() []*matchmakingclientmodels.ModelsMatchmakingResult {
	return o.Payload
}

func (o *BulkGetSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetSessionsBadRequest creates a BulkGetSessionsBadRequest with default headers values
func NewBulkGetSessionsBadRequest() *BulkGetSessionsBadRequest {
	return &BulkGetSessionsBadRequest{}
}

/*BulkGetSessionsBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type BulkGetSessionsBadRequest struct {
	Payload *matchmakingclientmodels.ResponseErrorV1
}

func (o *BulkGetSessionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsBadRequest  %+v", 400, o.Payload)
}

func (o *BulkGetSessionsBadRequest) GetPayload() *matchmakingclientmodels.ResponseErrorV1 {
	return o.Payload
}

func (o *BulkGetSessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseErrorV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetSessionsUnauthorized creates a BulkGetSessionsUnauthorized with default headers values
func NewBulkGetSessionsUnauthorized() *BulkGetSessionsUnauthorized {
	return &BulkGetSessionsUnauthorized{}
}

/*BulkGetSessionsUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type BulkGetSessionsUnauthorized struct {
	Payload *matchmakingclientmodels.ResponseErrorV1
}

func (o *BulkGetSessionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *BulkGetSessionsUnauthorized) GetPayload() *matchmakingclientmodels.ResponseErrorV1 {
	return o.Payload
}

func (o *BulkGetSessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseErrorV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetSessionsForbidden creates a BulkGetSessionsForbidden with default headers values
func NewBulkGetSessionsForbidden() *BulkGetSessionsForbidden {
	return &BulkGetSessionsForbidden{}
}

/*BulkGetSessionsForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr><tr><td>20014</td><td>invalid audience</td></tr><tr><td>20015</td><td>insufficient scope</td></tr></table>
*/
type BulkGetSessionsForbidden struct {
	Payload *matchmakingclientmodels.ResponseErrorV1
}

func (o *BulkGetSessionsForbidden) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsForbidden  %+v", 403, o.Payload)
}

func (o *BulkGetSessionsForbidden) GetPayload() *matchmakingclientmodels.ResponseErrorV1 {
	return o.Payload
}

func (o *BulkGetSessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseErrorV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetSessionsNotFound creates a BulkGetSessionsNotFound with default headers values
func NewBulkGetSessionsNotFound() *BulkGetSessionsNotFound {
	return &BulkGetSessionsNotFound{}
}

/*BulkGetSessionsNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>510110</td><td>channel not found</td></tr></table>
*/
type BulkGetSessionsNotFound struct {
	Payload *matchmakingclientmodels.ResponseErrorV1
}

func (o *BulkGetSessionsNotFound) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsNotFound  %+v", 404, o.Payload)
}

func (o *BulkGetSessionsNotFound) GetPayload() *matchmakingclientmodels.ResponseErrorV1 {
	return o.Payload
}

func (o *BulkGetSessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseErrorV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkGetSessionsInternalServerError creates a BulkGetSessionsInternalServerError with default headers values
func NewBulkGetSessionsInternalServerError() *BulkGetSessionsInternalServerError {
	return &BulkGetSessionsInternalServerError{}
}

/*BulkGetSessionsInternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type BulkGetSessionsInternalServerError struct {
	Payload *matchmakingclientmodels.ResponseError
}

func (o *BulkGetSessionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /matchmaking/v1/admin/namespaces/{namespace}/channels/all/sessions/bulk][%d] bulkGetSessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *BulkGetSessionsInternalServerError) GetPayload() *matchmakingclientmodels.ResponseError {
	return o.Payload
}

func (o *BulkGetSessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
