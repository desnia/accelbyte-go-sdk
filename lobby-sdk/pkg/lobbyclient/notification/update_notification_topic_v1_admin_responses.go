// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// UpdateNotificationTopicV1AdminReader is a Reader for the UpdateNotificationTopicV1Admin structure.
type UpdateNotificationTopicV1AdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationTopicV1AdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateNotificationTopicV1AdminNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNotificationTopicV1AdminBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateNotificationTopicV1AdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateNotificationTopicV1AdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateNotificationTopicV1AdminNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateNotificationTopicV1AdminInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateNotificationTopicV1AdminNoContent creates a UpdateNotificationTopicV1AdminNoContent with default headers values
func NewUpdateNotificationTopicV1AdminNoContent() *UpdateNotificationTopicV1AdminNoContent {
	return &UpdateNotificationTopicV1AdminNoContent{}
}

/*UpdateNotificationTopicV1AdminNoContent handles this case with default header values.

  No Content
*/
type UpdateNotificationTopicV1AdminNoContent struct {
}

func (o *UpdateNotificationTopicV1AdminNoContent) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminNoContent ", 204)
}

func (o *UpdateNotificationTopicV1AdminNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNotificationTopicV1AdminBadRequest creates a UpdateNotificationTopicV1AdminBadRequest with default headers values
func NewUpdateNotificationTopicV1AdminBadRequest() *UpdateNotificationTopicV1AdminBadRequest {
	return &UpdateNotificationTopicV1AdminBadRequest{}
}

/*UpdateNotificationTopicV1AdminBadRequest handles this case with default header values.

  Bad Request
*/
type UpdateNotificationTopicV1AdminBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UpdateNotificationTopicV1AdminBadRequest) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNotificationTopicV1AdminBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UpdateNotificationTopicV1AdminBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationTopicV1AdminUnauthorized creates a UpdateNotificationTopicV1AdminUnauthorized with default headers values
func NewUpdateNotificationTopicV1AdminUnauthorized() *UpdateNotificationTopicV1AdminUnauthorized {
	return &UpdateNotificationTopicV1AdminUnauthorized{}
}

/*UpdateNotificationTopicV1AdminUnauthorized handles this case with default header values.

  Unauthorized
*/
type UpdateNotificationTopicV1AdminUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UpdateNotificationTopicV1AdminUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateNotificationTopicV1AdminUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UpdateNotificationTopicV1AdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationTopicV1AdminForbidden creates a UpdateNotificationTopicV1AdminForbidden with default headers values
func NewUpdateNotificationTopicV1AdminForbidden() *UpdateNotificationTopicV1AdminForbidden {
	return &UpdateNotificationTopicV1AdminForbidden{}
}

/*UpdateNotificationTopicV1AdminForbidden handles this case with default header values.

  Forbidden
*/
type UpdateNotificationTopicV1AdminForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UpdateNotificationTopicV1AdminForbidden) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminForbidden  %+v", 403, o.Payload)
}

func (o *UpdateNotificationTopicV1AdminForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UpdateNotificationTopicV1AdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationTopicV1AdminNotFound creates a UpdateNotificationTopicV1AdminNotFound with default headers values
func NewUpdateNotificationTopicV1AdminNotFound() *UpdateNotificationTopicV1AdminNotFound {
	return &UpdateNotificationTopicV1AdminNotFound{}
}

/*UpdateNotificationTopicV1AdminNotFound handles this case with default header values.

  Not Found
*/
type UpdateNotificationTopicV1AdminNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UpdateNotificationTopicV1AdminNotFound) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNotificationTopicV1AdminNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UpdateNotificationTopicV1AdminNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationTopicV1AdminInternalServerError creates a UpdateNotificationTopicV1AdminInternalServerError with default headers values
func NewUpdateNotificationTopicV1AdminInternalServerError() *UpdateNotificationTopicV1AdminInternalServerError {
	return &UpdateNotificationTopicV1AdminInternalServerError{}
}

/*UpdateNotificationTopicV1AdminInternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateNotificationTopicV1AdminInternalServerError struct {
	Payload *lobbyclientmodels.RestapiErrorResponseV1
}

func (o *UpdateNotificationTopicV1AdminInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /lobby/v1/admin/notification/namespaces/{namespace}/topics/{topicName}][%d] updateNotificationTopicV1AdminInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNotificationTopicV1AdminInternalServerError) GetPayload() *lobbyclientmodels.RestapiErrorResponseV1 {
	return o.Payload
}

func (o *UpdateNotificationTopicV1AdminInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
