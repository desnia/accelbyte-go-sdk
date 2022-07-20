// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package bans_deprecated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// GetBansTypeReader is a Reader for the GetBansType structure.
type GetBansTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBansTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBansTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBansTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBansTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/bans returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetBansTypeOK creates a GetBansTypeOK with default headers values
func NewGetBansTypeOK() *GetBansTypeOK {
	return &GetBansTypeOK{}
}

/*GetBansTypeOK handles this case with default header values.

  OK
*/
type GetBansTypeOK struct {
	Payload *iamclientmodels.AccountcommonBans
}

func (o *GetBansTypeOK) Error() string {
	return fmt.Sprintf("[GET /iam/bans][%d] getBansTypeOK  %+v", 200, o.Payload)
}

func (o *GetBansTypeOK) GetPayload() *iamclientmodels.AccountcommonBans {
	return o.Payload
}

func (o *GetBansTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.AccountcommonBans)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBansTypeUnauthorized creates a GetBansTypeUnauthorized with default headers values
func NewGetBansTypeUnauthorized() *GetBansTypeUnauthorized {
	return &GetBansTypeUnauthorized{}
}

/*GetBansTypeUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetBansTypeUnauthorized struct {
}

func (o *GetBansTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/bans][%d] getBansTypeUnauthorized ", 401)
}

func (o *GetBansTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBansTypeForbidden creates a GetBansTypeForbidden with default headers values
func NewGetBansTypeForbidden() *GetBansTypeForbidden {
	return &GetBansTypeForbidden{}
}

/*GetBansTypeForbidden handles this case with default header values.

  Forbidden
*/
type GetBansTypeForbidden struct {
}

func (o *GetBansTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/bans][%d] getBansTypeForbidden ", 403)
}

func (o *GetBansTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}