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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// UsersGroupsPartialUpdateReader is a Reader for the UsersGroupsPartialUpdate structure.
type UsersGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsPartialUpdateOK creates a UsersGroupsPartialUpdateOK with default headers values
func NewUsersGroupsPartialUpdateOK() *UsersGroupsPartialUpdateOK {
	return &UsersGroupsPartialUpdateOK{}
}

/*UsersGroupsPartialUpdateOK handles this case with default header values.

UsersGroupsPartialUpdateOK users groups partial update o k
*/
type UsersGroupsPartialUpdateOK struct {
	Payload *models.Group
}

func (o *UsersGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/groups/{id}/][%d] usersGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsPartialUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsPartialUpdateDefault creates a UsersGroupsPartialUpdateDefault with default headers values
func NewUsersGroupsPartialUpdateDefault(code int) *UsersGroupsPartialUpdateDefault {
	return &UsersGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*UsersGroupsPartialUpdateDefault handles this case with default header values.

UsersGroupsPartialUpdateDefault users groups partial update default
*/
type UsersGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups partial update default response
func (o *UsersGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UsersGroupsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/groups/{id}/][%d] users_groups_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
