package tests

import (
	. "github.com/saichler/shared/go/share/strings"
	"reflect"
	"strings"
	"testing"
)

var Str = New("")

func init() {
	Str.TypesPrefix = true
}

func instanceOf(t *testing.T, str string, expectError bool) (interface{}, bool) {
	v, e := InstanceOf(str, globals.Registry())
	if e != nil && !expectError {
		log.Fail(t, "Did not expect an error for str:", str)
		return nil, false
	} else if e == nil && expectError {
		log.Fail(t, "Did expect an error for str:", str)
		return nil, false
	}
	return v, true
}

func checkString(s *String, ex string, t *testing.T) bool {
	if s.String() != ex {
		log.Fail(t, "Expected String to be '", ex, "' but got ", s.String())
		return false
	}
	return true
}

func checkToString(any interface{}, ex string, t *testing.T) bool {
	return checkToFromString(any, ex, "xyz", t)
}

func checkToFromString(any interface{}, ex, ex2 string, t *testing.T) bool {
	s := Str.StringOf(any)
	fs, err := InstanceOf(s, globals.Registry())
	if err != nil {
		log.Fail(t, "Did not expect an error: ", err.Error())
		return false
	}
	reflect.DeepEqual(any, fs)
	// Until struct is implemented, skip it
	if !reflect.DeepEqual(any, fs) && !strings.Contains(s, ",25") {
		log.Fail("error from string:", reflect.DeepEqual(any, fs), ":", any, ":", fs, s)
		return false
	}

	_ex := Kind2String(reflect.ValueOf(any)).Add(ex).String()
	_ex2 := Kind2String(reflect.ValueOf(any)).Add(ex2).String()
	if s != _ex && s != _ex2 && s != ex {
		log.Fail(t, "Expected String to be '", ex, "' but got '", s, "'")
		return false
	}
	return true
}

func TestString(t *testing.T) {
	s := New("test")
	if ok := checkString(s, "test", t); !ok {
		return
	}

	s.Add("test")
	if ok := checkString(s, "testtest", t); !ok {
		return
	}

	s.Join(New("test"))
	if ok := checkString(s, "testtesttest", t); !ok {
		return
	}
	if s.IsBlank() {
		log.Fail(t, "Expected s to NOT be blank")
		return
	}
	s = New("")
	if !s.IsBlank() {
		log.Fail(t, "Expected s to be blank")
		return
	}
}

func TestToString(t *testing.T) {
	if ok := checkToString("test", "test", t); !ok {
		return
	}
	if ok := checkToString(int32(4343), "4343", t); !ok {
		return
	}
	if ok := checkToString(uint32(4342), "4342", t); !ok {
		return
	}
	if ok := checkToString(float32(4342.5454), "4342.5454", t); !ok {
		return
	}
	if ok := checkToString(float64(4342.5454), "4342.5454", t); !ok {
		return
	}
	if ok := checkToString(true, "true", t); !ok {
		return
	}
	if ok := checkToString(true, "true", t); !ok {
		return
	}
	if ok := checkToString(nil, "", t); !ok {
		return
	}
	type test struct{}
	globals.Registry().Register(&test{})
	if ok := checkToString(&test{}, "{22,25}test", t); !ok {
		return
	}
	st := &test{}
	st = nil
	if ok := checkToString(st, "<Nil>", t); !ok {
		return
	}
	if ok := checkToString([]int{}, "[]", t); !ok {
		return
	}
	if ok := checkToString([]int{1, 2, 3}, "[1,2,3]", t); !ok {
		return
	}
	if ok := checkToString([]byte("ABC"), "ABC", t); !ok {
		return
	}
	if ok := checkToFromString(map[string]int{"a": 1, "b": 2}, "[a=1,b=2]", "[b=2,a=1]", t); !ok {
		return
	}

	k := reflect.New(reflect.ValueOf("").Type()).Interface()

	if ok := checkToString(k, "", t); !ok {
		return
	}
}

func TestFromStringPtr(t *testing.T) {
	s, e := InstanceOf("{22,24}test", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	s1 := *s.(*string)
	if s1 != "test" {
		log.Fail(t, "Expected value to be test but got ", s1)
		return
	}
}

func TestFromStringInt(t *testing.T) {
	v, e := InstanceOf("{2}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(int)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int {
		log.Fail(t, "From string failed for int")
		return
	}
	v, e = InstanceOf("{2}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringInt8(t *testing.T) {
	v, e := InstanceOf("{3}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(int8)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int8 {
		log.Fail(t, "From string failed for int8")
		return
	}
	v, e = InstanceOf("{3}5b", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringInt16(t *testing.T) {
	v, e := InstanceOf("{4}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(int16)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int16 {
		log.Fail(t, "From string failed for int16")
		return
	}
	v, e = InstanceOf("{4}5c", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringInt32(t *testing.T) {
	v, e := InstanceOf("{5}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(int32)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int32 {
		log.Fail(t, "From string failed for int32")
		return
	}
	v, e = InstanceOf("{5}5a", globals.Registry())
	v, e = InstanceOf("{5}", globals.Registry())
	r = v.(int32)
	if r != 0 {
		log.Fail(t, "From string failed for int32 blank")
		return
	}
}

func TestFromStringInt64(t *testing.T) {
	v, e := InstanceOf("{6}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(int64)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int64 {
		log.Fail(t, "From string failed for int64")
		return
	}
	v, e = InstanceOf("{6}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringUInt(t *testing.T) {
	v, e := InstanceOf("{7}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(uint)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Uint {
		log.Fail(t, "From string failed for Uint")
		return
	}
	v, e = InstanceOf("{7}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringUInt8(t *testing.T) {
	v, e := InstanceOf("{8}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.([]uint8)[0]
	//53 is the byte value of character 5
	if r != 53 || reflect.ValueOf(r).Kind() != reflect.Uint8 {
		log.Fail(t, "From string failed for Uint8")
		return
	}
}

func TestFromStringUInt16(t *testing.T) {
	v, e := InstanceOf("{9}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(uint16)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Uint16 {
		log.Fail(t, "From string failed for Uint16")
		return
	}
	v, e = InstanceOf("{9}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringUInt32(t *testing.T) {
	v, e := InstanceOf("{10}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(uint32)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Uint32 {
		log.Fail(t, "From string failed for Uint32")
		return
	}
	v, e = InstanceOf("{10}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringUInt64(t *testing.T) {
	v, e := InstanceOf("{11}5", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(uint64)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Uint64 {
		log.Fail(t, "From string failed for Uint64")
		return
	}
	v, e = InstanceOf("{11}5a", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringFloat32(t *testing.T) {
	v, e := InstanceOf("{13}5.8", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(float32)
	if r != 5.8 || reflect.ValueOf(r).Kind() != reflect.Float32 {
		log.Fail(t, "From string failed for float32")
		return
	}
	v, e = InstanceOf("{13}5.8d", globals.Registry())
	if e == nil {
		log.Fail(t, "Did expect an error")
	}
}

func TestFromStringSlice(t *testing.T) {
	s, e := InstanceOf("{23,24}[a,b]", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	s1 := s.([]string)
	if s1[0] != "a" {
		log.Fail(t, "value for index 0 was not equale to a")
		return
	}
	if s1[1] != "b" {
		log.Fail(t, "value for index 1 was not equale to b")
		return
	}
}

func TestFromStringInterface(t *testing.T) {
	v, e := InstanceOf("{20,13}5.8", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	r := v.(float32)
	if r != 5.8 || reflect.ValueOf(r).Kind() != reflect.Float32 {
		log.Fail(t, "From string failed for float32 interface")
		return
	}
}

func TestFromStringMap(t *testing.T) {
	s, e := InstanceOf("{21,24,2}[a=1,b=2]", globals.Registry())
	if e != nil {
		log.Fail(t, "Did not expect an error")
	}
	s1 := s.(map[string]int)
	if s1["a"] != 1 {
		log.Fail(t, "value for key 'a' was not found or not equale to 1")
		return
	}
	if s1["b"] != 2 {
		log.Fail(t, "value for key 'b' was not found or not equale to 2")
		return
	}
}

func TestAppendSpace(t *testing.T) {
	s := &String{AddSpaceWhenAdding: true}
	s.Add("a")
	s.Add("b")
	if s.String() != "a b" {
		log.Fail(t, "Expected a space:'"+s.String()+"'")
		return
	}
	if s.Len() != 3 {
		log.Fail("Expected lenght of 3")
		return
	}
	if len(s.Bytes()) != 3 {
		log.Fail("Expected lenght of 3")
		return
	}
	b := []byte{'c'}
	s.AddBytes(b)
	if len(s.Bytes()) != 4 {
		log.Fail("Expected lenght of 3")
		return
	}
}
