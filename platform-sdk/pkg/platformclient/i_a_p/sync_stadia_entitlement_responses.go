// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncStadiaEntitlementReader is a Reader for the SyncStadiaEntitlement structure.
type SyncStadiaEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncStadiaEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSyncStadiaEntitlementNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /platform/public/namespaces/{namespace}/users/{userId}/iap/stadia/sync returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncStadiaEntitlementNoContent creates a SyncStadiaEntitlementNoContent with default headers values
func NewSyncStadiaEntitlementNoContent() *SyncStadiaEntitlementNoContent {
	return &SyncStadiaEntitlementNoContent{}
}

/*SyncStadiaEntitlementNoContent handles this case with default header values.

  Successful operation
*/
type SyncStadiaEntitlementNoContent struct {
}

func (o *SyncStadiaEntitlementNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/public/namespaces/{namespace}/users/{userId}/iap/stadia/sync][%d] syncStadiaEntitlementNoContent ", 204)
}

func (o *SyncStadiaEntitlementNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
