package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	argNum := flag.NArg()
	firstArg := flag.Arg(0)
	switch argNum {
	case 0:
		showTip("Please enter a sub command  😅 ")
		os.Exit(0)
	case 1:
		fileName := flag.Arg(0)
		UploadFile(fileName)
	case 2:
		if firstArg != "layout" {
			showTip("Sub command '" + firstArg + "' does not support  😳 ")
			os.Exit(0)
		}
		// set layout
		config, err := ReadFromFile()
		if err != nil {
			showTip("Failed to read configuration ❗️ ")
			showTip("If you don't have a configuration account, you should configure your account first ❗️ ")
			os.Exit(0)
		}
		config.UploadNameLayout = flag.Arg(1)
		fmt.Println(config.UploadNameLayout)
		err = config.SaveToFile()
		if err != nil {
			showTip("Failed to save layout configuration ❗️ ")
			os.Exit(0)
		}
		showTip("Save configuration successfully ✅ ")
	case 4:
		if firstArg != "account" {
			showTip("Sub command '" + firstArg + "' does not support  😳 ")
			os.Exit(0)
		}

		args := flag.Args()
		accessKey := args[1]
		secretKey := args[2]
		bucketName := args[3]
		config := &Config{
			AccessKey:        accessKey,
			SecretKey:        secretKey,
			BucketName:       bucketName,
			UploadNameLayout: DefaultLayout,
		}
		err := config.SaveToFile()
		if err != nil {
			showTip("Failed to save configuration ❗️ ")
			os.Exit(0)
		}
		showTip("Save configuration successfully ✅ ")
	}
	// fmt.Println(flag.Arg(2))
	//pbq account -ak=accessKey -sk=secretKey -bn=bucketName
	//pbq layout $filename/$Y-$M-$D||$U-$filename
	//pbq fileNmae
	// if(flag.ar)
}

func showTip(tip string) {
	fmt.Println("[PBQ]:", tip)
}
