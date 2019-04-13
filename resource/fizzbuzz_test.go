package resource

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// fizzBuzzTest is used to test the execution of the function fizzBuzz and the parsing of its output
type fizzBuzzTest struct {
	params         FizzBuzzParams
	expectedOutput []string
	expectedErr    error
}

var (
	fizzBuzzTests = []fizzBuzzTest{
		{ // Valid
			params: FizzBuzzParams{
				Int1:    3,
				Int2:    5,
				String1: "fizz",
				String2: "buzz",
				Limit:   15,
			},
			expectedOutput: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"},
			expectedErr:    nil,
		},
		{ // Invalid: Limit should be greater than 0
			params: FizzBuzzParams{
				Int1:    3,
				Int2:    5,
				String1: "fizz",
				String2: "buzz",
				Limit:   0,
			},
			expectedOutput: nil,
			expectedErr:    ErrInvalidParameters,
		},
		{ // Invalid: Int1 should be greater than 0
			params: FizzBuzzParams{
				Int1:    0,
				Int2:    1,
				String1: "fizz",
				String2: "buzz",
				Limit:   1,
			},
			expectedOutput: nil,
			expectedErr:    ErrInvalidParameters,
		},
		{ // Invalid: Int2 should be greater than 0
			params: FizzBuzzParams{
				Int1:    1,
				Int2:    0,
				String1: "fizz",
				String2: "buzz",
				Limit:   1,
			},
			expectedOutput: nil,
			expectedErr:    ErrInvalidParameters,
		},
	}
)

func TestFizzBuzz(t *testing.T) {
	wg := sync.WaitGroup{}

	for _, test := range fizzBuzzTests {
		pr, pw := io.Pipe()

		wg.Add(1)
		go func(pr *io.PipeReader, test fizzBuzzTest) { // Reading the pipe should be done asynchronously
			defer func() {
				pr.Close()
				wg.Done()
			}()

			var response []string

			decoder := json.NewDecoder(pr)
			err := decoder.Decode(&response)
			if test.expectedErr != nil { // fizzBuzz should returns an error and the JSON decoding too
				require.Errorf(t, err, "An error is expected during JSON decoding")
			} else {
				require.NoErrorf(t, err, "Unexpected error during JSON decoding")
				require.Equalf(t, test.expectedOutput, response, "Response not expected")
			}

		}(pr, test)

		err := fizzBuzz(test.params, pw)
		require.Equalf(t, test.expectedErr, err, "Unexpected value of error during fizzBuzz processing")

		pw.Close()
		wg.Wait()
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range fizzBuzzTests {
			fizzBuzz(test.params, ioutil.Discard)
		}
	}
}
