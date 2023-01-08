// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     payment_response.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PaymentStatus int32

const (
	PaymentStatusCOMPLETED PaymentStatus = 0
	PaymentStatusCANCELLED PaymentStatus = 1
	PaymentStatusFAILED    PaymentStatus = 2
)

func (e PaymentStatus) String() string {
	switch e {
	case PaymentStatusCOMPLETED:
		return "COMPLETED"
	case PaymentStatusCANCELLED:
		return "CANCELLED"
	case PaymentStatusFAILED:
		return "FAILED"
	}
	return "unknown"
}

func writePaymentStatus(r PaymentStatus, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewPaymentStatusValue(raw string) (r PaymentStatus, err error) {
	switch raw {
	case "COMPLETED":
		return PaymentStatusCOMPLETED, nil
	case "CANCELLED":
		return PaymentStatusCANCELLED, nil
	case "FAILED":
		return PaymentStatusFAILED, nil
	}

	return -1, fmt.Errorf("invalid value for PaymentStatus: '%s'", raw)

}

func (b PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *PaymentStatus) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewPaymentStatusValue(stringVal)
	*b = val
	return err
}

type PaymentStatusWrapper struct {
	Target *PaymentStatus
}

func (b PaymentStatusWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b PaymentStatusWrapper) SetInt(v int32) {
	*(b.Target) = PaymentStatus(v)
}

func (b PaymentStatusWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b PaymentStatusWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b PaymentStatusWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b PaymentStatusWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b PaymentStatusWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b PaymentStatusWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b PaymentStatusWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b PaymentStatusWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b PaymentStatusWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b PaymentStatusWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b PaymentStatusWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b PaymentStatusWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b PaymentStatusWrapper) Finalize() {}