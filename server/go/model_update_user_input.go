/*
 * HelloApiSchema
 *
 * Practice api schema
 *
 * API version: 1.0.0
 * Contact: doriven@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UpdateUserInput struct {

	Id int64 `json:"id"`

	EmailAddress string `json:"email_address,omitempty"`

	LastName string `json:"last_name,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	Birthday string `json:"birthday,omitempty"`

	Address string `json:"address,omitempty"`
}

// AssertUpdateUserInputRequired checks if the required fields are not zero-ed
func AssertUpdateUserInputRequired(obj UpdateUserInput) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUpdateUserInputRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UpdateUserInput (e.g. [][]UpdateUserInput), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUpdateUserInputRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUpdateUserInput, ok := obj.(UpdateUserInput)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUpdateUserInputRequired(aUpdateUserInput)
	})
}
