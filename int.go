package bigx

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Int represents a big int value with a number of convenience methods and conversions built in
type Int struct {
	value *big.Int
}

// Int64 constructs Int from int64
func Int64(v int64) *Int {
	return &Int{
		value: big.NewInt(v),
	}
}

// New constructs new Int from string
func New(s string) (*Int, bool) {
	if bi, ok := big.NewInt(0).SetString(s, 10); ok {
		return &Int{value: bi}, true
	}
	return nil, false
}

// Add adds value from `that`. nil is considered to be zero
func (i *Int) Add(v *Int) *Int {
	switch {
	case i == nil && v == nil:
		return nil
	case i == nil:
		return v
	case v == nil:
		return i
	default:
		return &Int{
			value: big.NewInt(0).Add(i.value, v.value),
		}
	}
}

// BigInt provides interop with *big.Int
func (i *Int) BigInt() *big.Int {
	return i.value
}

// BigFloat provides interop with *big.Float
func (i *Int) BigFloat() *big.Float {
	return big.NewFloat(0).SetInt(i.value)
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (i *Int) Cmp(v *Int) int {
	switch {
	case i == nil && v == nil:
		return 0
	case i == nil:
		return -1
	case v == nil:
		return 1
	default:
		return i.value.Cmp(v.value)
	}
}

// MarshalDynamoDBAttributeValue implements dynamodb.Marshaler
func (i *Int) MarshalDynamoDBAttributeValue(item *dynamodb.AttributeValue) error {
	if i == nil {
		item.NULL = aws.Bool(true)
		return nil
	}

	item.N = aws.String(i.value.String())

	return nil
}

// MarshalJSON implements json.Marshaler
func (i *Int) MarshalJSON() ([]byte, error) {
	if i == nil {
		return []byte("null"), nil
	}

	return json.Marshal(i.value.String())
}

// Mul multiplies the value by `that`.  Returns nil if either value or that is nil
func (i *Int) Mul(that *Int) *Int {
	if i == nil || that == nil {
		return nil
	}

	return &Int{
		value: big.NewInt(0).Mul(i.value, that.value),
	}
}

// Quo divides value by `that`.  Returns nil if either value or that is nil
func (i *Int) Quo(that *Int) *Int {
	if i == nil || that == nil {
		return nil
	}

	return &Int{
		value: big.NewInt(0).Quo(i.value, that.value),
	}
}

// String expresses value as string
func (i *Int) String() string {
	if i == nil {
		return ""
	}
	return i.value.String()
}

// Sub subtracts value from `that`. nil is considered to be zero
func (i *Int) Sub(v *Int) *Int {
	switch {
	case i == nil && v == nil:
		return nil
	case i == nil:
		return &Int{
			value: big.NewInt(0).Sub(big.NewInt(0), v.value),
		}
	case v == nil:
		return i
	default:
		return &Int{
			value: big.NewInt(0).Sub(i.value, v.value),
		}
	}
}

// Uint64 converts value to int64.  nil will return 0
func (i *Int) Uint64() uint64 {
	if i == nil {
		return 0
	}
	return i.value.Uint64()
}

// UnmarshalDynamoDBAttributeValue implements dynamodb.Unmarshaler
func (i *Int) UnmarshalDynamoDBAttributeValue(item *dynamodb.AttributeValue) error {
	switch {
	case aws.BoolValue(item.NULL):
		return nil

	case item.N != nil:
		s := aws.StringValue(item.N)
		v, ok := big.NewInt(0).SetString(s, 10)
		if !ok {
			return fmt.Errorf("failed to parse bigx.Int, %v", s)
		}
		*i = Int{value: v}
		return nil

	default:
		return fmt.Errorf("don't know how to unmarshal item into bigx.Int")
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (i *Int) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("failed to unmarshal bigx.Int: %w", err)
	}

	if s == "" {
		return nil
	}

	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return fmt.Errorf("failed to parse bigx, %v", s)
	}

	*i = Int{
		value: v,
	}

	return nil
}
