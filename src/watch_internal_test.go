package pike

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWatch_EmptyARN(t *testing.T) {
	err := Watch("", 10)
	assert.Error(t, err)
	assert.IsType(t, &arnEmptyError{}, err)
}

func TestWatch_InvalidWaitTime(t *testing.T) {
	err := Watch("arn:aws:iam::123456789012:policy/test-policy", 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "wait time must be positive")

	err = Watch("arn:aws:iam::123456789012:policy/test-policy", -5)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "wait time must be positive")
}

func TestSortActions_WithArrayActions(t *testing.T) {
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": ["s3:PutObject", "s3:GetObject", "s3:DeleteObject"],
				"Resource": "*"
			}
		]
	}`

	result, err := sortActions(policy)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(*result), &parsed)
	assert.NoError(t, err)

	statements := parsed["Statement"].([]interface{})
	statement := statements[0].(map[string]interface{})
	actions := statement["Action"].([]interface{})

	// Verify actions are sorted
	assert.Equal(t, "s3:DeleteObject", actions[0])
	assert.Equal(t, "s3:GetObject", actions[1])
	assert.Equal(t, "s3:PutObject", actions[2])
}

func TestSortActions_WithStringAction(t *testing.T) {
	policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": "s3:GetObject",
				"Resource": "*"
			}
		]
	}`

	result, err := sortActions(policy)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	var parsed map[string]interface{}
	err = json.Unmarshal([]byte(*result), &parsed)
	assert.NoError(t, err)
}

func TestSortActions_InvalidJSON(t *testing.T) {
	policy := `{"invalid": json}`

	result, err := sortActions(policy)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.IsType(t, &unmarshallJSONError{}, err)
}

func TestSortActions_InvalidStatement(t *testing.T) {
	policy := `{
		"Version": "2012-10-17",
		"Statement": "invalid"
	}`

	result, err := sortActions(policy)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.IsType(t, &castToListOfInterfaceError{}, err)
}

func TestSortInterfaceStrings_Success(t *testing.T) {
	actions := []interface{}{"s3:PutObject", "s3:GetObject", "s3:DeleteObject"}

	result := sortInterfaceStrings(actions)

	assert.NotNil(t, result)
	assert.Len(t, result, 3)
	assert.Equal(t, "s3:DeleteObject", result[0])
	assert.Equal(t, "s3:GetObject", result[1])
	assert.Equal(t, "s3:PutObject", result[2])
}

func TestSortInterfaceStrings_InvalidInput(t *testing.T) {
	actions := "not an array"

	result := sortInterfaceStrings(actions)

	assert.Nil(t, result)
}

func TestSortInterfaceStrings_NonStringElements(t *testing.T) {
	actions := []interface{}{"s3:Xavier", "s3:GetObject", 123, "s3:PutObject"}

	result := sortInterfaceStrings(actions)

	assert.NotNil(t, result)
	assert.Len(t, result, 4)
	// Non-string elements should be skipped, but array length preserved
	assert.Equal(t, "s3:GetObject", result[1])
	assert.Equal(t, "", result[0]) // default zero value for skipped element
	assert.Equal(t, "s3:PutObject", result[2])
}

func TestWaitExpiredError_Error(t *testing.T) {
	err := &waitExpiredError{}
	assert.Equal(t, "wait expired with no change", err.Error())
}

func TestUrlEscapeError_Error(t *testing.T) {
	originalErr := errors.New("test error")
	err := &urlEscapeError{err: originalErr}
	assert.Contains(t, err.Error(), "failed to unescape url")
	assert.Contains(t, err.Error(), "test error")
}

func TestCastToListOfInterfaceError_Error(t *testing.T) {
	err := &castToListOfInterfaceError{}
	assert.Equal(t, "failed to convert to list of interfaces", err.Error())
}

func TestGetVersionError_Error(t *testing.T) {
	originalErr := errors.New("test error")
	err := &getVersionError{err: originalErr}
	assert.Contains(t, err.Error(), "failed to get version")
	assert.Contains(t, err.Error(), "test error")
}

func TestWaitForPolicyChangeError_Error(t *testing.T) {
	originalErr := errors.New("test error")
	err := &waitForPolicyChangeError{err: originalErr}
	assert.Contains(t, err.Error(), "failed to wait for policy change")
	assert.Contains(t, err.Error(), "test error")
}
