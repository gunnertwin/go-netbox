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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// SecretsSecretRolesReadReader is a Reader for the SecretsSecretRolesRead structure.
type SecretsSecretRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsSecretRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretsSecretRolesReadOK creates a SecretsSecretRolesReadOK with default headers values
func NewSecretsSecretRolesReadOK() *SecretsSecretRolesReadOK {
	return &SecretsSecretRolesReadOK{}
}

/*SecretsSecretRolesReadOK handles this case with default header values.

SecretsSecretRolesReadOK secrets secret roles read o k
*/
type SecretsSecretRolesReadOK struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /secrets/secret-roles/{id}/][%d] secretsSecretRolesReadOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretRolesReadOK) GetPayload() *models.SecretRole {
	return o.Payload
}

func (o *SecretsSecretRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
