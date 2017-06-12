package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	nonFlagNum := flag.NArg()
	fmt.Println(nonFlagNum)
	fmt.Println(flag.Arg(2))
	//pbq account ak sk
	//pbq layout $filename/$Y-$M-$D||$U-$filename
	//pbq fileNmae
}

// UploadFile upload file to server
func UploadFile(fileNmae, accessKey, secretKey string) {

}
