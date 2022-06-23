package main

import (
	"io/ioutil"
	"log"
	"protobuf-lesson/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_DESIGNER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": {}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Data{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

}
