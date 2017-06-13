package main

import (
	"testing"

	"os"

	"fmt"

	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveConfig(t *testing.T) {

	Convey("test save config to confing file", t, func() {
		config := &Config{
			AccessKey:        "AccessKeyAccessKeyAccessKey",
			SecretKey:        "SecretKeySecretKeySecretKey",
			UploadNameLayout: "$U-$filename",
			BucketName:       "blog",
		}
		err := config.SaveToFile()
		So(err, ShouldBeNil)
	})
}

func TestReadConfig(t *testing.T) {
	Convey("test read config file", t, func() {
		config := new(Config)
		config, err := ReadFromFile()
		So(err, ShouldBeNil)
		So(config.AccessKey, ShouldEqual, "AccessKeyAccessKeyAccessKey")
		So(config.SecretKey, ShouldEqual, "SecretKeySecretKeySecretKey")
		So(config.UploadNameLayout, ShouldEqual, "$U-$filename")
		So(config.BucketName, ShouldEqual, "blog")
		//del file
		if _, err := os.Stat(ConfigFilePath); err == nil {
			err = os.Remove(ConfigFilePath)
			if err != nil {
				t.Fatal(err)
			}
		}
		config, err = ReadFromFile()
		So(err, ShouldNotBeNil)
		So(config, ShouldBeNil)
	})
}

func TestFormatUploadFileName(t *testing.T) {
	Convey("test format upload file name", t, func() {
		config := &Config{UploadNameLayout: ""}
		ret := config.FormatUploadFileName("a.png")
		fmt.Println(ret)
		config.UploadNameLayout = "%YYYY/%MM/%DD/%FILENAME"
		ret = config.FormatUploadFileName("a.png")
		fmt.Println(ret)
		So(ret, ShouldEqual, time.Now().Format("2006/01/02/a.png"))
	})
}
