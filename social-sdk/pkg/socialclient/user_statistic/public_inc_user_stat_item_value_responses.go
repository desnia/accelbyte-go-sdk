// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// PublicIncUserStatItemValueReader is a Reader for the PublicIncUserStatItemValue structure.
type PublicIncUserStatItemValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicIncUserStatItemValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicIncUserStatItemValueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicIncUserStatItemValueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicIncUserStatItemValueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewPublicIncUserStatItemValueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicIncUserStatItemValueOK creates a PublicIncUserStatItemValueOK with default headers values
func NewPublicIncUserStatItemValueOK() *PublicIncUserStatItemValueOK {
	return &PublicIncUserStatItemValueOK{}
}

/*PublicIncUserStatItemValueOK handles this case with default header values.

  successful operation
*/
type PublicIncUserStatItemValueOK struct {
	Payload *socialclientmodels.StatItemIncResult
}

func (o *PublicIncUserStatItemValueOK) Error() string {
	return fmt.Sprintf("[PATCH /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value][%d] publicIncUserStatItemValueOK  %+v", 200, o.ToJSONString())
}

func (o *PublicIncUserStatItemValueOK) ToJSONString() string {
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

func (o *PublicIncUserStatItemValueOK) GetPayload() *socialclientmodels.StatItemIncResult {
	return o.Payload
}

func (o *PublicIncUserStatItemValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(socialclientmodels.StatItemIncResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicIncUserStatItemValueBadRequest creates a PublicIncUserStatItemValueBadRequest with default headers values
func NewPublicIncUserStatItemValueBadRequest() *PublicIncUserStatItemValueBadRequest {
	return &PublicIncUserStatItemValueBadRequest{}
}

/*PublicIncUserStatItemValueBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12221</td><td>Invalid stat operator, expect [{expected}] but actual [{actual}]</td></tr></table>
*/
type PublicIncUserStatItemValueBadRequest struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *PublicIncUserStatItemValueBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value][%d] publicIncUserStatItemValueBadRequest  %+v", 400, o.ToJSONString())
}

func (o *PublicIncUserStatItemValueBadRequest) ToJSONString() string {
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

func (o *PublicIncUserStatItemValueBadRequest) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicIncUserStatItemValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicIncUserStatItemValueNotFound creates a PublicIncUserStatItemValueNotFound with default headers values
func NewPublicIncUserStatItemValueNotFound() *PublicIncUserStatItemValueNotFound {
	return &PublicIncUserStatItemValueNotFound{}
}

/*PublicIncUserStatItemValueNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12241</td><td>Stat [{statCode}] cannot be found in namespace [{namespace}]</td></tr><tr><td>12242</td><td>Stat item of [{statCode}] of user [{profileId}] cannot be found in namespace [{namespace}]</td></tr></table>
*/
type PublicIncUserStatItemValueNotFound struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *PublicIncUserStatItemValueNotFound) Error() string {
	return fmt.Sprintf("[PATCH /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value][%d] publicIncUserStatItemValueNotFound  %+v", 404, o.ToJSONString())
}

func (o *PublicIncUserStatItemValueNotFound) ToJSONString() string {
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

func (o *PublicIncUserStatItemValueNotFound) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicIncUserStatItemValueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicIncUserStatItemValueConflict creates a PublicIncUserStatItemValueConflict with default headers values
func NewPublicIncUserStatItemValueConflict() *PublicIncUserStatItemValueConflict {
	return &PublicIncUserStatItemValueConflict{}
}

/*PublicIncUserStatItemValueConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12273</td><td>Stat [{statCode}] is not decreasable</td></tr><tr><td>12275</td><td>[{action}] value: [{value}] of stat [{statCode}]  is out of range while minimum [{minimum}] and maximum [{maximum}] in namespace [{namespace}]</td></tr></table>
*/
type PublicIncUserStatItemValueConflict struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *PublicIncUserStatItemValueConflict) Error() string {
	return fmt.Sprintf("[PATCH /social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value][%d] publicIncUserStatItemValueConflict  %+v", 409, o.ToJSONString())
}

func (o *PublicIncUserStatItemValueConflict) ToJSONString() string {
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

func (o *PublicIncUserStatItemValueConflict) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicIncUserStatItemValueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
