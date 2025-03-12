package convert

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestToInt(t *testing.T) {
	val, err := ToInt("123")
	if err != nil || val != 123 {
		t.Errorf("ToInt failed: expected 123, got %d, err: %v", val, err)
	}
}

func TestToInt_Invalid(t *testing.T) {
	_, err := ToInt("abc")
	if err == nil {
		t.Error("ToInt should fail on invalid input, but got no error")
	}
}

func TestFromInt(t *testing.T) {
	result := FromInt(456)
	if result != "456" {
		t.Errorf("FromInt failed: expected \"456\", got %s", result)
	}
}

func TestToInt64(t *testing.T) {
	val, err := ToInt64("9876543210")
	if err != nil || val != 9876543210 {
		t.Errorf("ToInt64 failed: expected 9876543210, got %d, err: %v", val, err)
	}
}

func TestToInt64_Invalid(t *testing.T) {
	_, err := ToInt64("xyz")
	if err == nil {
		t.Error("ToInt64 should fail on invalid input, but got no error")
	}
}

func TestFromInt64(t *testing.T) {
	result := FromInt64(9876543210)
	if result != "9876543210" {
		t.Errorf("FromInt64 failed: expected \"9876543210\", got %s", result)
	}
}

func TestToFloat64(t *testing.T) {
	val, err := ToFloat64("3.1415")
	if err != nil || val != 3.1415 {
		t.Errorf("ToFloat64 failed: expected 3.1415, got %f, err: %v", val, err)
	}
}

func TestToFloat64_Invalid(t *testing.T) {
	_, err := ToFloat64("not-a-float")
	if err == nil {
		t.Error("ToFloat64 should fail on invalid input, but got no error")
	}
}

func TestFromFloat64(t *testing.T) {
	result := FromFloat64(2.71828)
	if result != "2.71828" {
		t.Errorf("FromFloat64 failed: expected \"2.71828\", got %s", result)
	}
}

func TestToUint64(t *testing.T) {
	val, err := ToUint64("1234567890")
	if err != nil || val != 1234567890 {
		t.Errorf("ToUint64 failed: expected 1234567890, got %d, err: %v", val, err)
	}
}

func TestToUint64_Invalid(t *testing.T) {
	_, err := ToUint64("-123")
	if err == nil {
		t.Error("ToUint64 should fail on negative input, but got no error")
	}
}

func TestFromUint64(t *testing.T) {
	result := FromUint64(1234567890)
	if result != "1234567890" {
		t.Errorf("FromUint64 failed: expected \"1234567890\", got %s", result)
	}
}

func TestToBool(t *testing.T) {
	tests := map[string]bool{
		"true":  true,
		"false": false,
		"1":     true,
		"0":     false,
	}
	for input, expected := range tests {
		result, err := ToBool(input)
		if err != nil || result != expected {
			t.Errorf("ToBool failed: input %s, expected %v, got %v, err: %v", input, expected, result, err)
		}
	}
}

func TestToBool_Invalid(t *testing.T) {
	_, err := ToBool("maybe")
	if err == nil {
		t.Error("ToBool should fail on invalid input, but got no error")
	}
}

func TestFromBool(t *testing.T) {
	if FromBool(true) != "true" {
		t.Errorf("FromBool(true) failed: expected \"true\"")
	}
	if FromBool(false) != "false" {
		t.Errorf("FromBool(false) failed: expected \"false\"")
	}
}

func TestToBase64(t *testing.T) {
	input := []byte("hello world")
	encoded := ToBase64(input)
	expected := "aGVsbG8gd29ybGQ="
	if encoded != expected {
		t.Errorf("ToBase64 failed: expected %s, got %s", expected, encoded)
	}
}

func TestFromBase64(t *testing.T) {
	input := "aGVsbG8gd29ybGQ="
	decoded, err := FromBase64(input)
	if err != nil {
		t.Errorf("FromBase64 returned error: %v", err)
	}
	expected := "hello world"
	if string(decoded) != expected {
		t.Errorf("FromBase64 failed: expected %s, got %s", expected, string(decoded))
	}
}

func TestFromBase64_Invalid(t *testing.T) {
	_, err := FromBase64("!!invalid!!")
	if err == nil {
		t.Error("FromBase64 should fail on invalid input, but got no error")
	}
}

func TestToHex(t *testing.T) {
	input := []byte("abc123")
	result := ToHex(input)
	expected := "616263313233"
	if result != expected {
		t.Errorf("ToHex failed: expected %s, got %s", expected, result)
	}
}

func TestFromHex(t *testing.T) {
	input := "616263313233"
	result, err := FromHex(input)
	if err != nil {
		t.Errorf("FromHex returned error: %v", err)
	}
	expected := []byte("abc123")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromHex failed: expected %s, got %s", expected, result)
	}
}

func TestFromHex_Invalid(t *testing.T) {
	_, err := FromHex("ZZZ123")
	if err == nil {
		t.Error("FromHex should fail on invalid input, but got no error")
	}
}

func TestToJSON(t *testing.T) {
	data := map[string]int{"one": 1, "two": 2}
	result, err := ToJSON(data)
	if err != nil {
		t.Errorf("ToJSON returned error: %v", err)
	}
	expected1 := `{"one":1,"two":2}`
	expected2 := `{"two":2,"one":1}` // порядок может отличаться
	if result != expected1 && result != expected2 {
		t.Errorf("ToJSON failed: got %s", result)
	}
}

func TestFromJSON(t *testing.T) {
	jsonStr := `{"one":1,"two":2}`
	var m map[string]int
	err := FromJSON(jsonStr, &m)
	if err != nil {
		t.Errorf("FromJSON returned error: %v", err)
	}
	expected := map[string]int{"one": 1, "two": 2}
	if !reflect.DeepEqual(m, expected) {
		t.Errorf("FromJSON failed: expected %v, got %v", expected, m)
	}
}

func TestFromJSON_Invalid(t *testing.T) {
	var m map[string]int
	err := FromJSON("{bad-json}", &m)
	if err == nil {
		t.Error("FromJSON should fail on invalid input, but got no error")
	}
}

func TestToXML(t *testing.T) {
	type Item struct {
		XMLName xml.Name `xml:"item"`
		ID      int      `xml:"id"`
		Name    string   `xml:"name"`
	}
	item := Item{ID: 42, Name: "Test"}
	xmlStr, err := ToXML(item)
	if err != nil {
		t.Errorf("ToXML returned error: %v", err)
	}
	if !(len(xmlStr) > 0 && (xmlStr[:5] == "<item" || xmlStr[:6] == "<?xml")) {
		t.Errorf("ToXML result seems invalid: %s", xmlStr)
	}
}

func TestFromXML(t *testing.T) {
	type Item struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	}
	xmlStr := `<Item><id>42</id><name>Test</name></Item>`
	var item Item
	err := FromXML(xmlStr, &item)
	if err != nil {
		t.Errorf("FromXML returned error: %v", err)
	}
	if item.ID != 42 || item.Name != "Test" {
		t.Errorf("FromXML failed: expected ID=42, Name=Test, got %+v", item)
	}
}

func TestFromXML_Invalid(t *testing.T) {
	var v any
	err := FromXML("<bad-xml>", &v)
	if err == nil {
		t.Error("FromXML should fail on invalid input, but got no error")
	}
}
