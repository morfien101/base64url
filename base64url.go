package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(b64.URLEncoding.EncodeToString([]byte(os.Args[1])))
	} else {
		fmt.Println("Usage: base64url StringToEncode")
	}
}
