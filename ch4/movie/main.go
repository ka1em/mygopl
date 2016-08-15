package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mygopl/test_aes/aes2"
	"os"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color, omitempty"`
	Actor []string
}

func main() {

	movies := []Movie{
		{Title: "AAA", Year: 1992, Color: false,
			Actor: []string{"HHH", "YYY"}},
		{Title: "XXX", Year: 1000, Color: true,
			Actor: []string{"wang"}},
	}

	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("Json Marshl err: %s", err)
	}

	fmt.Println("JSON")
	fmt.Printf("%s\n", data)

	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64 encoding")
	fmt.Println(str)

	tmp_data_2, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Errorf("base64 decode error %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Decode")
	fmt.Printf("%s", tmp_data_2)

	key := []byte("abcdefghijklmnopqnrstuvwxyzabcde")
	tmp_data_3, err := aes2.AesEncrypt(data, key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "AesEncrypt error %s", err)
		os.Exit(3)
	}

	fmt.Println("\nAES encoding")
	fmt.Println(tmp_data_3)
	//fmt.Println(base64.StdEncoding.EncodeToString(tmp_data_3))

	fmt.Println("\nAES Deconding")
	tmp_data_4, err := aes2.AesDecrypt(tmp_data_3, key)
	if err != nil {
		fmt.Errorf("Error %s", err)
		os.Exit(4)
	}

	fmt.Println(string(tmp_data_4))
	os.Exit(0)
}
