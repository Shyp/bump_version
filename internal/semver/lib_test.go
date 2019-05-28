package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeVersion(t *testing.T) {
	testCases := []struct {
		in    string
		vtype VersionType
		out   string
	}{
		{"0.4", Major, "1.0"},
		{"0.4.0", Major, "1.0.0"},
		{"1.0", Major, "2.0"},
		{"1", Major, "2"},
		{"1.0.1", Minor, "1.1.0"},
	}
	for _, tt := range testCases {
		v, err := changeVersion(tt.vtype, tt.in)
		if err != nil {
			t.Fatal(err)
		}
		if v.String() != tt.out {
			t.Errorf("changeVersion(%s, %s): got %s, want %s", tt.vtype, tt.in, v.String(), tt.out)
		}
	}
}

func TestExtract1(t *testing.T) {

	path1 := "../../test/1/main.go"
	v1 := "1.2.3"

	str, err := Extract(path1)
	assert.Nil(t, err, "err")
	assert.Equal(t, v1, str, "v1")
}

func TestExtract2(t *testing.T) {

	path1 := "../../test/2/main.go"
	v1 := "0.3.2"

	str, err := Extract(path1)
	assert.Nil(t, err, "err")
	assert.Equal(t, v1, str, "v1")
}

func TestExtractNonexistent(t *testing.T) {

	path1 := "../../test/no/such/path/main.go"
	v1 := ""

	str, err := Extract(path1)
	assert.NotNil(t, err, "err")
	assert.Equal(t, v1, str, "v1")
}
