package bigx

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tj/assert"
	"testing"
)

func TestInt_JSON(t *testing.T) {
	testCases := map[string]struct {
		Input *Int
		Want  string
	}{
		"nop": {
			Input: nil,
			Want:  `null`,
		},
		"1": {
			Input: Int64(1),
			Want:  `"1"`,
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			data, err := json.Marshal(tc.Input)
			assert.Nil(t, err)
			assert.Equal(t, tc.Want, string(data))
		})
	}
}

func TestInt_DynamoDB(t *testing.T) {
	testCases := map[string]struct {
		Value *Int
	}{
		"nop": {
			Value: nil,
		},
		"1": {
			Value: Int64(1),
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			item, err := dynamodbattribute.Marshal(tc.Value)
			assert.Nil(t, err)

			var got *Int
			err = dynamodbattribute.Unmarshal(item, &got)
			assert.Nil(t, err)

			assert.Equal(t, 0, tc.Value.Cmp(got))
		})
	}
}

func TestInt_Add(t *testing.T) {
	testCases := map[string]struct {
		A    *Int
		B    *Int
		Want *Int
	}{
		"both nil": {},
		"nil a": {
			A:    nil,
			B:    Int64(1),
			Want: Int64(1),
		},
		"nil b": {
			A:    Int64(1),
			B:    nil,
			Want: Int64(1),
		},
		"ok": {
			A:    Int64(1),
			B:    Int64(2),
			Want: Int64(3),
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			got := tc.A.Add(tc.B)
			assert.Equal(t, 0, tc.Want.Cmp(got))
		})
	}
}

func TestInt_Sub(t *testing.T) {
	testCases := map[string]struct {
		A    *Int
		B    *Int
		Want *Int
	}{
		"both nil": {},
		"nil a": {
			A:    nil,
			B:    Int64(1),
			Want: Int64(-1),
		},
		"nil b": {
			A:    Int64(1),
			B:    nil,
			Want: Int64(1),
		},
		"ok": {
			A:    Int64(5),
			B:    Int64(2),
			Want: Int64(3),
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			got := tc.A.Sub(tc.B)
			assert.Equal(t, 0, tc.Want.Cmp(got))
		})
	}
}

func TestInt_Mul(t *testing.T) {
	testCases := map[string]struct {
		A    *Int
		B    *Int
		Want *Int
	}{
		"both nil": {},
		"nil a": {
			A:    nil,
			B:    Int64(1),
			Want: nil,
		},
		"nil b": {
			A:    Int64(1),
			B:    nil,
			Want: nil,
		},
		"ok": {
			A:    Int64(5),
			B:    Int64(2),
			Want: Int64(10),
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			got := tc.A.Mul(tc.B)
			assert.Equal(t, 0, tc.Want.Cmp(got))
		})
	}
}

func TestInt_Quo(t *testing.T) {
	testCases := map[string]struct {
		A    *Int
		B    *Int
		Want *Int
	}{
		"both nil": {},
		"nil a": {
			A:    nil,
			B:    Int64(1),
			Want: nil,
		},
		"nil b": {
			A:    Int64(1),
			B:    nil,
			Want: nil,
		},
		"ok": {
			A:    Int64(10),
			B:    Int64(2),
			Want: Int64(5),
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			got := tc.A.Quo(tc.B)
			assert.Equal(t, 0, tc.Want.Cmp(got))
		})
	}
}

func TestNew(t *testing.T) {
	v, ok := New("123")
	assert.True(t, ok)
	assert.Equal(t, uint64(123), v.Uint64())
	assert.Equal(t, "123", v.String())
	assert.Equal(t, "123", v.BigInt().String())
	assert.Equal(t, "123", v.BigFloat().String())
}

func TestInt_UnmarshalJSON(t *testing.T) {
	want := Int64(123)
	data, err := json.Marshal(want)
	assert.Nil(t, err)

	var got *Int
	err = json.Unmarshal(data, &got)
	assert.Nil(t, err)

	assert.Equal(t, want, got)
}
