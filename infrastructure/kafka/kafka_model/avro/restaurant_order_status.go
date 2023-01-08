// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     restaurant_approval_request.avsc
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

type RestaurantOrderStatus int32

const (
	RestaurantOrderStatusPAID RestaurantOrderStatus = 0
)

func (e RestaurantOrderStatus) String() string {
	switch e {
	case RestaurantOrderStatusPAID:
		return "PAID"
	}
	return "unknown"
}

func writeRestaurantOrderStatus(r RestaurantOrderStatus, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewRestaurantOrderStatusValue(raw string) (r RestaurantOrderStatus, err error) {
	switch raw {
	case "PAID":
		return RestaurantOrderStatusPAID, nil
	}

	return -1, fmt.Errorf("invalid value for RestaurantOrderStatus: '%s'", raw)

}

func (b RestaurantOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *RestaurantOrderStatus) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewRestaurantOrderStatusValue(stringVal)
	*b = val
	return err
}

type RestaurantOrderStatusWrapper struct {
	Target *RestaurantOrderStatus
}

func (b RestaurantOrderStatusWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b RestaurantOrderStatusWrapper) SetInt(v int32) {
	*(b.Target) = RestaurantOrderStatus(v)
}

func (b RestaurantOrderStatusWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b RestaurantOrderStatusWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b RestaurantOrderStatusWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b RestaurantOrderStatusWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b RestaurantOrderStatusWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b RestaurantOrderStatusWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b RestaurantOrderStatusWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b RestaurantOrderStatusWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b RestaurantOrderStatusWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b RestaurantOrderStatusWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b RestaurantOrderStatusWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b RestaurantOrderStatusWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b RestaurantOrderStatusWrapper) Finalize() {}
