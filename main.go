package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

type FeedType struct {
	Author  interface{} `json:"author"`
	Entry   interface{} `json:"entry"`
	Updated LabelType   `json:"updated"`
	Rights  LabelType   `json:"rights"`
	Title   LabelType   `json:"title"`
	Link    interface{} `json:"link"`
	Icon    LabelType   `json:"icon"`
	ID      LabelType   `json:"id"`
}
type LabelType struct {
	Label string `json:"label"`
}
type Data struct {
	Feed FeedType `json:"feed"`
}

func main() {
	content, err := ioutil.ReadFile("./results.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var data Data
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// map
	// authors := reflect.ValueOf(data.Feed.Author)
	// inspection(authors)

	// slice
	entry := reflect.ValueOf(data.Feed.Entry)
	inspect_rec([]string{"entry"}, entry)
	// inspection(entry)

	// link := reflect.ValueOf(data.Feed.Link)
	// inspect_rec([]string{"link"}, link)

	// struct
	// updated := reflect.ValueOf(data.Feed.Entry)
	// inspection(updated)
}

func inspection(data reflect.Value) {
	fmt.Printf("Type %v, Kind %v, typeOf %v \n", data.Type(), data.Kind(), reflect.TypeOf(data))
	switch data.Kind() {
	case reflect.Map:
		for _, key := range data.MapKeys() {
			value := data.MapIndex(key)
			fmt.Printf("key %v, value %v \n", key, value)
		}
	default:
		for i := 0; i < data.Len(); i++ {
			fmt.Printf("key %v, value %v \n", i, data.Index(i))
		}
	}
}

func inspect_rec(parent []string, data reflect.Value) {
	switch data.Kind() {
	case reflect.Slice:
		for i := 0; i < data.Len(); i++ {
			value := data.Index(i)
			fmt.Printf("Data %v \n", i)
			inspect_rec(parent, reflect.ValueOf(value.Interface()))
		}
	case reflect.Map:
		for _, key := range data.MapKeys() {
			value := data.MapIndex(key)
			parent = append(parent, key.String())
			inspect_rec(parent, reflect.ValueOf(value.Interface()))

			// pop out the last visited node name
			parent = parent[:len(parent)-1]
		}
	default:
		fmt.Printf("\t %v: %v \n", strings.Join(parent, "_"), data.String())
	}
}
