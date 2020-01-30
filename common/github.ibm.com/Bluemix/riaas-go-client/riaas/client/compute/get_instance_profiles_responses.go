// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetInstanceProfilesReader is a Reader for the GetInstanceProfiles structure.
type GetInstanceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstanceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstanceProfilesOK creates a GetInstanceProfilesOK with default headers values
func NewGetInstanceProfilesOK() *GetInstanceProfilesOK {
	return &GetInstanceProfilesOK{}
}

/*GetInstanceProfilesOK handles this case with default header values.

instance profiles retrieved successfully
*/
type GetInstanceProfilesOK struct {
	Payload *GetInstanceProfilesOKBody
}

func (o *GetInstanceProfilesOK) Error() string {
	return fmt.Sprintf("[GET /instance/profiles][%d] getInstanceProfilesOK  %+v", 200, o.Payload)
}

func (o *GetInstanceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetInstanceProfilesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstanceProfilesOKBody instanceProfileCollection
swagger:model GetInstanceProfilesOKBody
*/
type GetInstanceProfilesOKBody struct {

	// first
	First *GetInstanceProfilesOKBodyFirst `json:"first,omitempty"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// next
	Next *models.Next `json:"next,omitempty"`

	// Collection of instance profiels
	Profiles []*models.InstanceProfile `json:"profiles,omitempty"`

	// The total number of resources across all pages
	// Minimum: 0
	TotalCount *int64 `json:"total_count,omitempty"`
}

// Validate validates this get instance profiles o k body
func (o *GetInstanceProfilesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstanceProfilesOKBody) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(o.First) { // not required
		return nil
	}

	if o.First != nil {
		if err := o.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstanceProfilesOK" + "." + "first")
			}
			return err
		}
	}

	return nil
}

func (o *GetInstanceProfilesOKBody) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(o.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("getInstanceProfilesOK"+"."+"limit", "body", int64(o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("getInstanceProfilesOK"+"."+"limit", "body", int64(o.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (o *GetInstanceProfilesOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if o.Next != nil {
		if err := o.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstanceProfilesOK" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (o *GetInstanceProfilesOKBody) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(o.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(o.Profiles); i++ {
		if swag.IsZero(o.Profiles[i]) { // not required
			continue
		}

		if o.Profiles[i] != nil {
			if err := o.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getInstanceProfilesOK" + "." + "profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetInstanceProfilesOKBody) validateTotalCount(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("getInstanceProfilesOK"+"."+"total_count", "body", int64(*o.TotalCount), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstanceProfilesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstanceProfilesOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstanceProfilesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstanceProfilesOKBodyFirst get instance profiles o k body first
swagger:model GetInstanceProfilesOKBodyFirst
*/
type GetInstanceProfilesOKBodyFirst struct {

	// href
	Href interface{} `json:"href,omitempty"`
}

// Validate validates this get instance profiles o k body first
func (o *GetInstanceProfilesOKBodyFirst) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstanceProfilesOKBodyFirst) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstanceProfilesOKBodyFirst) UnmarshalBinary(b []byte) error {
	var res GetInstanceProfilesOKBodyFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}