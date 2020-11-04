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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// TenancyTenantsUpdateReader is a Reader for the TenancyTenantsUpdate structure.
type TenancyTenantsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsUpdateOK creates a TenancyTenantsUpdateOK with default headers values
func NewTenancyTenantsUpdateOK() *TenancyTenantsUpdateOK {
	return &TenancyTenantsUpdateOK{}
}

/*TenancyTenantsUpdateOK handles this case with default header values.

TenancyTenantsUpdateOK tenancy tenants update o k
*/
type TenancyTenantsUpdateOK struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsUpdateDefault creates a TenancyTenantsUpdateDefault with default headers values
func NewTenancyTenantsUpdateDefault(code int) *TenancyTenantsUpdateDefault {
	return &TenancyTenantsUpdateDefault{
		_statusCode: code,
	}
}

/*TenancyTenantsUpdateDefault handles this case with default header values.

TenancyTenantsUpdateDefault tenancy tenants update default
*/
type TenancyTenantsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenants update default response
func (o *TenancyTenantsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancy_tenants_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
