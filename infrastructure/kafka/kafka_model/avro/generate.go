package avro

//go:generate $GOPATH/bin/gogen-avro -containers . payment_request.avsc
//go:generate $GOPATH/bin/gogen-avro -containers . payment_response.avsc
//go:generate $GOPATH/bin/gogen-avro -containers . restaurant_approval_request.avsc
//go:generate $GOPATH/bin/gogen-avro -containers . restaurant_approval_response.avsc
