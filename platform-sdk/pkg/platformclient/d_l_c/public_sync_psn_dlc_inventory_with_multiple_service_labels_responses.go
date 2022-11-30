// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package d_l_c

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

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// PublicSyncPsnDlcInventoryWithMultipleServiceLabelsReader is a Reader for the PublicSyncPsnDlcInventoryWithMultipleServiceLabels structure.
type PublicSyncPsnDlcInventoryWithMultipleServiceLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/public/namespaces/{namespace}/users/{userId}/dlc/psn/sync/multiServiceLabels returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent creates a PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent with default headers values
func NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent() *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent {
	return &PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent{}
}

/*PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent handles this case with default header values.

  Successful operation
*/
type PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent struct {
}

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/public/namespaces/{namespace}/users/{userId}/dlc/psn/sync/multiServiceLabels][%d] publicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent ", 204)
}

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest creates a PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest with default headers values
func NewPublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest() *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest {
	return &PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest{}
}

/*PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>39125</td><td>Invalid platform [{platformId}] user token</td></tr><tr><td>39126</td><td>User id [{}] in namespace [{}] doesn't link platform [{}]</td></tr><tr><td>39127</td><td>Invalid service label [{serviceLabel}]</td></tr>
*/
type PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/public/namespaces/{namespace}/users/{userId}/dlc/psn/sync/multiServiceLabels][%d] publicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest  %+v", 400, o.ToJSONString())
}

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest) ToJSONString() string {
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

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicSyncPsnDlcInventoryWithMultipleServiceLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}