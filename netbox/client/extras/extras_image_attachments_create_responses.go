// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// ExtrasImageAttachmentsCreateReader is a Reader for the ExtrasImageAttachmentsCreate structure.
type ExtrasImageAttachmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasImageAttachmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsCreateCreated creates a ExtrasImageAttachmentsCreateCreated with default headers values
func NewExtrasImageAttachmentsCreateCreated() *ExtrasImageAttachmentsCreateCreated {
	return &ExtrasImageAttachmentsCreateCreated{}
}

/*ExtrasImageAttachmentsCreateCreated handles this case with default header values.

ExtrasImageAttachmentsCreateCreated extras image attachments create created
*/
type ExtrasImageAttachmentsCreateCreated struct {
	Payload *models.ImageAttachment
}

func (o *ExtrasImageAttachmentsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/image-attachments/][%d] extrasImageAttachmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasImageAttachmentsCreateCreated) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsCreateDefault creates a ExtrasImageAttachmentsCreateDefault with default headers values
func NewExtrasImageAttachmentsCreateDefault(code int) *ExtrasImageAttachmentsCreateDefault {
	return &ExtrasImageAttachmentsCreateDefault{
		_statusCode: code,
	}
}

/*ExtrasImageAttachmentsCreateDefault handles this case with default header values.

ExtrasImageAttachmentsCreateDefault extras image attachments create default
*/
type ExtrasImageAttachmentsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments create default response
func (o *ExtrasImageAttachmentsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/image-attachments/][%d] extras_image-attachments_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
