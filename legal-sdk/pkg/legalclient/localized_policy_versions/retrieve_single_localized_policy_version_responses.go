// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package localized_policy_versions

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

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// RetrieveSingleLocalizedPolicyVersionReader is a Reader for the RetrieveSingleLocalizedPolicyVersion structure.
type RetrieveSingleLocalizedPolicyVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveSingleLocalizedPolicyVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveSingleLocalizedPolicyVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveSingleLocalizedPolicyVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /agreement/admin/localized-policy-versions/{localizedPolicyVersionId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRetrieveSingleLocalizedPolicyVersionOK creates a RetrieveSingleLocalizedPolicyVersionOK with default headers values
func NewRetrieveSingleLocalizedPolicyVersionOK() *RetrieveSingleLocalizedPolicyVersionOK {
	return &RetrieveSingleLocalizedPolicyVersionOK{}
}

/*RetrieveSingleLocalizedPolicyVersionOK handles this case with default header values.

  successful operation
*/
type RetrieveSingleLocalizedPolicyVersionOK struct {
	Payload *legalclientmodels.RetrieveLocalizedPolicyVersionResponse
}

func (o *RetrieveSingleLocalizedPolicyVersionOK) Error() string {
	return fmt.Sprintf("[GET /agreement/admin/localized-policy-versions/{localizedPolicyVersionId}][%d] retrieveSingleLocalizedPolicyVersionOK  %+v", 200, o.ToJSONString())
}

func (o *RetrieveSingleLocalizedPolicyVersionOK) ToJSONString() string {
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

func (o *RetrieveSingleLocalizedPolicyVersionOK) GetPayload() *legalclientmodels.RetrieveLocalizedPolicyVersionResponse {
	return o.Payload
}

func (o *RetrieveSingleLocalizedPolicyVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(legalclientmodels.RetrieveLocalizedPolicyVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveSingleLocalizedPolicyVersionBadRequest creates a RetrieveSingleLocalizedPolicyVersionBadRequest with default headers values
func NewRetrieveSingleLocalizedPolicyVersionBadRequest() *RetrieveSingleLocalizedPolicyVersionBadRequest {
	return &RetrieveSingleLocalizedPolicyVersionBadRequest{}
}

/*RetrieveSingleLocalizedPolicyVersionBadRequest handles this case with default header values.

  <table><tr><td>NumericErrorCode</td><td>ErrorCode</td></tr><tr><td>2912</td><td>errors.net.accelbyte.platform.legal.invalid_policy_version</td></tr></table>
*/
type RetrieveSingleLocalizedPolicyVersionBadRequest struct {
	Payload *legalclientmodels.ErrorEntity
}

func (o *RetrieveSingleLocalizedPolicyVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /agreement/admin/localized-policy-versions/{localizedPolicyVersionId}][%d] retrieveSingleLocalizedPolicyVersionBadRequest  %+v", 400, o.ToJSONString())
}

func (o *RetrieveSingleLocalizedPolicyVersionBadRequest) ToJSONString() string {
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

func (o *RetrieveSingleLocalizedPolicyVersionBadRequest) GetPayload() *legalclientmodels.ErrorEntity {
	return o.Payload
}

func (o *RetrieveSingleLocalizedPolicyVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(legalclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
