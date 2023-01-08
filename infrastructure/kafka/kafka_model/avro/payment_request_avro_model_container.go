// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     payment_request.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/container"
	"github.com/actgardner/gogen-avro/v10/vm"
)

func NewPaymentRequestAvroModelWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewPaymentRequestAvroModel()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type PaymentRequestAvroModelReader struct {
	r io.Reader
	p *vm.Program
}

func NewPaymentRequestAvroModelReader(r io.Reader) (*PaymentRequestAvroModelReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewPaymentRequestAvroModel()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &PaymentRequestAvroModelReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r PaymentRequestAvroModelReader) Read() (PaymentRequestAvroModel, error) {
	t := NewPaymentRequestAvroModel()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}