// Code generated by go-swagger; DO NOT EDIT.

package event_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/eventlog-sdk/pkg/eventlogclientmodels"
)

// QueryEventStreamHandlerReader is a Reader for the QueryEventStreamHandler structure.
type QueryEventStreamHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryEventStreamHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryEventStreamHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryEventStreamHandlerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewQueryEventStreamHandlerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryEventStreamHandlerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewQueryEventStreamHandlerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueryEventStreamHandlerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 501:
		result := NewQueryEventStreamHandlerNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /event/v2/admin/namespaces/{namespace}/query returns an error %d: %s", response.Code(), string(data))
	}
}

// NewQueryEventStreamHandlerOK creates a QueryEventStreamHandlerOK with default headers values
func NewQueryEventStreamHandlerOK() *QueryEventStreamHandlerOK {
	return &QueryEventStreamHandlerOK{}
}

/*QueryEventStreamHandlerOK handles this case with default header values.

  OK
*/
type QueryEventStreamHandlerOK struct {
	Payload *eventlogclientmodels.ModelsEventResponseV2
}

func (o *QueryEventStreamHandlerOK) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerOK  %+v", 200, o.Payload)
}

func (o *QueryEventStreamHandlerOK) GetPayload() *eventlogclientmodels.ModelsEventResponseV2 {
	return o.Payload
}

func (o *QueryEventStreamHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(eventlogclientmodels.ModelsEventResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryEventStreamHandlerBadRequest creates a QueryEventStreamHandlerBadRequest with default headers values
func NewQueryEventStreamHandlerBadRequest() *QueryEventStreamHandlerBadRequest {
	return &QueryEventStreamHandlerBadRequest{}
}

/*QueryEventStreamHandlerBadRequest handles this case with default header values.

  Bad Request
*/
type QueryEventStreamHandlerBadRequest struct {
}

func (o *QueryEventStreamHandlerBadRequest) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerBadRequest ", 400)
}

func (o *QueryEventStreamHandlerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryEventStreamHandlerUnauthorized creates a QueryEventStreamHandlerUnauthorized with default headers values
func NewQueryEventStreamHandlerUnauthorized() *QueryEventStreamHandlerUnauthorized {
	return &QueryEventStreamHandlerUnauthorized{}
}

/*QueryEventStreamHandlerUnauthorized handles this case with default header values.

  Unauthorized
*/
type QueryEventStreamHandlerUnauthorized struct {
}

func (o *QueryEventStreamHandlerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerUnauthorized ", 401)
}

func (o *QueryEventStreamHandlerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryEventStreamHandlerForbidden creates a QueryEventStreamHandlerForbidden with default headers values
func NewQueryEventStreamHandlerForbidden() *QueryEventStreamHandlerForbidden {
	return &QueryEventStreamHandlerForbidden{}
}

/*QueryEventStreamHandlerForbidden handles this case with default header values.

  Forbidden
*/
type QueryEventStreamHandlerForbidden struct {
}

func (o *QueryEventStreamHandlerForbidden) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerForbidden ", 403)
}

func (o *QueryEventStreamHandlerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryEventStreamHandlerNotFound creates a QueryEventStreamHandlerNotFound with default headers values
func NewQueryEventStreamHandlerNotFound() *QueryEventStreamHandlerNotFound {
	return &QueryEventStreamHandlerNotFound{}
}

/*QueryEventStreamHandlerNotFound handles this case with default header values.

  Not Found
*/
type QueryEventStreamHandlerNotFound struct {
}

func (o *QueryEventStreamHandlerNotFound) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerNotFound ", 404)
}

func (o *QueryEventStreamHandlerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryEventStreamHandlerInternalServerError creates a QueryEventStreamHandlerInternalServerError with default headers values
func NewQueryEventStreamHandlerInternalServerError() *QueryEventStreamHandlerInternalServerError {
	return &QueryEventStreamHandlerInternalServerError{}
}

/*QueryEventStreamHandlerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type QueryEventStreamHandlerInternalServerError struct {
}

func (o *QueryEventStreamHandlerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerInternalServerError ", 500)
}

func (o *QueryEventStreamHandlerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryEventStreamHandlerNotImplemented creates a QueryEventStreamHandlerNotImplemented with default headers values
func NewQueryEventStreamHandlerNotImplemented() *QueryEventStreamHandlerNotImplemented {
	return &QueryEventStreamHandlerNotImplemented{}
}

/*QueryEventStreamHandlerNotImplemented handles this case with default header values.

  Internal Server Error
*/
type QueryEventStreamHandlerNotImplemented struct {
}

func (o *QueryEventStreamHandlerNotImplemented) Error() string {
	return fmt.Sprintf("[POST /event/v2/admin/namespaces/{namespace}/query][%d] queryEventStreamHandlerNotImplemented ", 501)
}

func (o *QueryEventStreamHandlerNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}