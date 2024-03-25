package fizzbuzz

import (
	"reflect"
	"testing"

	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
)

func TestFizzbuzz(t *testing.T) {
	t.Run("Fizzbuzz with default parameters", func(t *testing.T) {
		params := &fizzbuzzParameters.Params{
			Int1:  3,
			Int2:  5,
			Limit: 15,
			Str1:  "Fizz",
			Str2:  "Buzz",
		}

		expected := []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
		}

		result := Fizzbuzz(params)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Fizzbuzz() = %v, want %v", result, expected)
		}
	})

	t.Run("Fizzbuzz with custom parameters", func(t *testing.T) {
		params := &fizzbuzzParameters.Params{
			Int1:  2,
			Int2:  3,
			Limit: 10,
			Str1:  "Foo",
			Str2:  "Bar",
		}

		expected := []string{
			"1", "Foo", "Bar", "Foo", "5", "FooBar", "7", "Foo", "Bar", "Foo",
		}

		result := Fizzbuzz(params)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Fizzbuzz() = %v, want %v", result, expected)
		}
	})

	t.Run("Fizzbuzz with limit < int1", func(t *testing.T) {
		params := &fizzbuzzParameters.Params{
			Int1:  5,
			Int2:  3,
			Limit: 3,
			Str1:  "Fizz",
			Str2:  "Buzz",
		}

		expected := []string{
			"1", "2", "Buzz",
		}

		result := Fizzbuzz(params)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Fizzbuzz() = %v, want %v", result, expected)
		}
	})

	t.Run("Fizzbuzz with limit < int2", func(t *testing.T) {
		params := &fizzbuzzParameters.Params{
			Int1:  3,
			Int2:  5,
			Limit: 3,
			Str1:  "Fizz",
			Str2:  "Buzz",
		}

		expected := []string{
			"1", "2", "Fizz",
		}

		result := Fizzbuzz(params)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Fizzbuzz() = %v, want %v", result, expected)
		}
	})
}
