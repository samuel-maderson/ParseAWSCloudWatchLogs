package main

import (
	cbase64 "cloudwatch-to-socket/libs/base64"
	cgzip "cloudwatch-to-socket/libs/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JSONEventData struct {
	Data string `json:"data"`
}

type JSONEvent struct {
	Awslogs JSONEventData `json:"awslogs"`
}

/* TODO - parse jcontent variable(JSON) to display each of the fields */

func main() {

	filename := "event.json"
	var jdata JSONEvent
	byte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(byte, &jdata)
	datagzip := cbase64.B64decode(jdata.Awslogs.Data)
	jcontent := cgzip.GzipDecompress(datagzip)
	fmt.Println(jcontent)
}
