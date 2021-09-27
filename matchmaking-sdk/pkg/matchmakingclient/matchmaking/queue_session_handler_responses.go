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

// QueueSessionHandlerReader is a Reader for the QueueSessionHandler structure.
type QueueSessionHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueueSessionHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewQueueSessionHandlerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueueSessionHandlerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewQueueSessionHandlerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueueSessionHandlerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueueSessionHandlerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /matchmaking/namespaces/{namespace}/sessions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueueSessionHandlerNoContent creates a QueueSessionHandlerNoContent with default headers values
func NewQueueSessionHandlerNoContent() *QueueSessionHandlerNoContent {
	return &QueueSessionHandlerNoContent{}
}

/*QueueSessionHandlerNoContent handles this case with default header values.

  No Content
*/
type QueueSessionHandlerNoContent struct {
}

func (o *QueueSessionHandlerNoContent) Error() string {
	return fmt.Sprintf("[POST /matchmaking/namespaces/{namespace}/sessions][%d] queueSessionHandlerNoContent ", 204)
}

func (o *QueueSessionHandlerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueueSessionHandlerBadRequest creates a QueueSessionHandlerBadRequest with default headers values
func NewQueueSessionHandlerBadRequest() *QueueSessionHandlerBadRequest {
	return &QueueSessionHandlerBadRequest{}
}

/*QueueSessionHandlerBadRequest handles this case with default header values.

  Bad Request
*/
type QueueSessionHandlerBadRequest struct {
	Payload *matchmakingclientmodels.ResponseError
}

func (o *QueueSessionHandlerBadRequest) Error() string {
	return fmt.Sprintf("[POST /matchmaking/namespaces/{namespace}/sessions][%d] queueSessionHandlerBadRequest  %+v", 400, o.Payload)
}

func (o *QueueSessionHandlerBadRequest) GetPayload() *matchmakingclientmodels.ResponseError {
	return o.Payload
}

func (o *QueueSessionHandlerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueueSessionHandlerUnauthorized creates a QueueSessionHandlerUnauthorized with default headers values
func NewQueueSessionHandlerUnauthorized() *QueueSessionHandlerUnauthorized {
	return &QueueSessionHandlerUnauthorized{}
}

/*QueueSessionHandlerUnauthorized handles this case with default header values.

  Unauthorized
*/
type QueueSessionHandlerUnauthorized struct {
	Payload *matchmakingclientmodels.ResponseError
}

func (o *QueueSessionHandlerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /matchmaking/namespaces/{namespace}/sessions][%d] queueSessionHandlerUnauthorized  %+v", 401, o.Payload)
}

func (o *QueueSessionHandlerUnauthorized) GetPayload() *matchmakingclientmodels.ResponseError {
	return o.Payload
}

func (o *QueueSessionHandlerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueueSessionHandlerForbidden creates a QueueSessionHandlerForbidden with default headers values
func NewQueueSessionHandlerForbidden() *QueueSessionHandlerForbidden {
	return &QueueSessionHandlerForbidden{}
}

/*QueueSessionHandlerForbidden handles this case with default header values.

  Forbidden
*/
type QueueSessionHandlerForbidden struct {
	Payload *matchmakingclientmodels.ResponseError
}

func (o *QueueSessionHandlerForbidden) Error() string {
	return fmt.Sprintf("[POST /matchmaking/namespaces/{namespace}/sessions][%d] queueSessionHandlerForbidden  %+v", 403, o.Payload)
}

func (o *QueueSessionHandlerForbidden) GetPayload() *matchmakingclientmodels.ResponseError {
	return o.Payload
}

func (o *QueueSessionHandlerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueueSessionHandlerInternalServerError creates a QueueSessionHandlerInternalServerError with default headers values
func NewQueueSessionHandlerInternalServerError() *QueueSessionHandlerInternalServerError {
	return &QueueSessionHandlerInternalServerError{}
}

/*QueueSessionHandlerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type QueueSessionHandlerInternalServerError struct {
	Payload *matchmakingclientmodels.ResponseError
}

func (o *QueueSessionHandlerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /matchmaking/namespaces/{namespace}/sessions][%d] queueSessionHandlerInternalServerError  %+v", 500, o.Payload)
}

func (o *QueueSessionHandlerInternalServerError) GetPayload() *matchmakingclientmodels.ResponseError {
	return o.Payload
}

func (o *QueueSessionHandlerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(matchmakingclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
