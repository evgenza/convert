package convert

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"strconv"
)

// ToInt converts the given string to an integer.
// Returns an error if the string is not a valid integer.
func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// FromInt converts an integer to its string representation.
func FromInt(i int) string {
	return strconv.Itoa(i)
}

// ToInt64 converts the given string to an int64.
// Returns an error if the string is not a valid int64.
func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// FromInt64 converts an int64 to its string representation.
func FromInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToFloat64 converts the given string to a float64.
// Returns an error if the string is not a valid float64.
func ToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// FromFloat64 converts a float64 to its string representation
// without using exponential notation.
func FromFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// ToUint64 converts the given string to a uint64.
// Returns an error if the string is not a valid unsigned integer.
func ToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// FromUint64 converts a uint64 to its string representation.
func FromUint64(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// ToBool converts the given string to a boolean.
// Accepts "1", "t", "T", "true", "TRUE", "True" → true;
// and "0", "f", "F", "false", "FALSE", "False" → false.
// Returns an error for invalid input.
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// FromBool converts a boolean value to a string.
// true → "true", false → "false".
func FromBool(b bool) string {
	return strconv.FormatBool(b)
}

// ToBase64 encodes a byte slice to a standard Base64 string.
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromBase64 decodes a Base64-encoded string to a byte slice.
// Returns an error if the input string is not valid Base64.
func FromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// ToHex encodes a byte slice to a hexadecimal string.
func ToHex(data []byte) string {
	return hex.EncodeToString(data)
}

// FromHex decodes a hexadecimal string to a byte slice.
// Returns an error if the input is not valid hexadecimal.
func FromHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

// ToJSON serializes a value to a JSON string.
// Returns an error if marshalling fails.
func ToJSON(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// FromJSON deserializes a JSON string into the provided pointer v.
// v must be a pointer to the target variable.
func FromJSON(s string, v any) error {
	return json.Unmarshal([]byte(s), v)
}

// ToXML serializes a value to an XML string.
// Returns an error if marshalling fails.
func ToXML(v any) (string, error) {
	b, err := xml.Marshal(v)
	return string(b), err
}

// FromXML deserializes an XML string into the provided pointer v.
// v must be a pointer to the target variable.
func FromXML(s string, v any) error {
	return xml.Unmarshal([]byte(s), v)
}