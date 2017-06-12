package main

import (
	"testing"

	"os"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveAccount(t *testing.T) {

	Convey("test save account to confing file", t, func() {
		err := SaveAccount(Account{
			AccessKey: "AccessKeyAccessKeyAccessKey",
			SecretKey: "SecretKeySecretKeySecretKey",
		})
		So(err, ShouldBeNil)
	})
}

func TestReadAccount(t *testing.T) {
	Convey("test read account confing file", t, func() {
		account, err := ReadAccount()
		So(err, ShouldBeNil)
		So(account.AccessKey, ShouldEqual, "AccessKeyAccessKeyAccessKey")
		So(account.SecretKey, ShouldEqual, "SecretKeySecretKeySecretKey")
		//del file
		if _, err := os.Stat(AccountConfFilePath); err == nil {
			err = os.Remove(AccountConfFilePath)
			if err != nil {
				t.Fatal(err)
			}
		}
		account, err = ReadAccount()
		So(err, ShouldNotBeNil)
		So(account, ShouldBeNil)
	})
}
