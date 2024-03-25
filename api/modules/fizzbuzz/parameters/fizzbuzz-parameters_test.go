package fizzbuzzParameters

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParseParams(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)

	expected := &Params{
		Int1:  3,
		Int2:  5,
		Limit: 15,
		Str1:  "Fizz",
		Str2:  "Buzz",
	}

	result, err := ParseParams(c)

	if err != nil {
		t.Errorf("ParseParams() returned an unexpected got: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseParams() = %v, want %v", result, expected)
	}
}

func TestParseParams_InvalidInt1(t *testing.T) {
	t.Run("int1 is an invalid integer", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=abc&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("int1 must be a valid integer")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

	t.Run("int1 is less than 1", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=0&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("int1 must be greater than 0")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

	t.Run("int1 is greater than math.MaxInt", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=9223372036854775808&int2=5&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}

func TestParseParams_InvalidInt2(t *testing.T) {
	t.Run("int2 is an invalid integer", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=abc&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("int2 must be a valid integer")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

	t.Run("int2 is less than 1", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=0&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("int2 must be greater than 0")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

	t.Run("int2 is greater than math.MaxInt", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=9223372036854775808&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("int2 is equal to int1", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=3&limit=15&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("int2 must be different from Int1")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})
}

func TestParseParams_InvalidLimit(t *testing.T) {
	t.Run("limit is an invalid integer", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=5&limit=abc&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("limit must be a valid integer")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

	t.Run("limit is less than 1", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/?int1=3&int2=5&limit=0&str1=Fizz&str2=Buzz", nil)

		_, err := ParseParams(c)

		expectedErr := errors.New("limit must be greater than 0")

		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("ParseParams() returned an unexpected got: %v, want %v", err, expectedErr)
		}
	})

}
