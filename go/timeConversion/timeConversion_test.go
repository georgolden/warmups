package timeConversion

import (
	"testing"
)

type testCase struct {
	arg  string
	want string
}

func TestTimeConversion(t *testing.T) {
	testCases := []testCase{
		{"12:00:00AM", "00:00:00"},
		{"12:00:00PM", "12:00:00"},
		{"09:00:00PM", "21:00:00"},
		{"09:00:00AM", "09:00:00"},
	}

	for _, tc := range testCases {
		got := TimeConversion(tc.arg)
		if tc.want != got {
			t.Errorf("Expected %v, got %v", tc.want, got)
		}
	}

}
