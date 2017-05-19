package main

import (
	"fmt"
	"log"

	// Path to protobuff
	"github.com/alextanhongpin/go-protobuf-example/protobuf"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	test := &example.SearchRequest{
		Query:         "hello", //proto.String("hello"),
		PageNumber:    12,      //proto.Int32(17),
		ResultPerPage: 123,     //proto.Int32(17),
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newTest := &example.SearchRequest{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}

	// Now test and newTest contains the same data
	if test.GetQuery() != newTest.GetQuery() {
		log.Fatalf("data mismatch %q != %q", test.GetQuery(), newTest.GetQuery())

	}
	fmt.Printf("%q", test)
}
