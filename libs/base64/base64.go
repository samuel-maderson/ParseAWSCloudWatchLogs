package libs

import (
	"encoding/base64"
)

func B64decode(data string) string {
	rawDecodedText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	return string(rawDecodedText)
}
