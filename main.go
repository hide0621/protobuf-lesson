package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
)

func main() {

	//Employee構造体のインスタンスを作る
	employee := &pb.Employee{
		Id:          1,
		Name:        "suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,                                            //列挙型で定義していた
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},                        //stringのrepeatedで定義していた
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}}, //値が構造体なのでポインタ型
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Data{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	/*
		//インスタンスができたので、上記Employee構造体をシリアライズして、バイナリーファイルとして書き出す
		binData, err := proto.Marshal(employee)
		if err != nil {
			log.Fatalln("Can't serialize", err)
		}

		//シリアライズされたのでファイルに書き出す
		if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
			log.Fatalln("Can't write", err)
		}

		//test.binからバイナリーデータを読み込み、変数inに代入
		in, err := ioutil.ReadFile("test.bin")
		if err != nil {
			log.Fatalln("Can't read file", err)
		}

		readEmployee := &pb.Employee{}

		//変数inの中にあるバイナリーデータがでシリアライズされ、上記で生成した空の構造体に入る
		err = proto.Unmarshal(in, readEmployee)
		if err != nil {
			log.Fatalln("Can't deserialize", err)
		}

		//ちゃんとデシリアライズされているか確認
		fmt.Println(readEmployee)
	*/

	//json文字列データに変換
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee) //引数データをjson文字列に変換する
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}

	/*
		//ちゃんとjson文字列に変換されているか確認
		fmt.Println(out)
	*/

	//json文字列データから構造体に戻す
	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil { //変数outはjson文字列データなので、それをアンマーシャルして構造体（readEmployee）に代入
		log.Fatalln("Can't unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
