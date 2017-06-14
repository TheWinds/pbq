package main

import (
	"fmt"
	"os"
	"path"

	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/api.v6/rs"

	"github.com/atotto/clipboard"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	rpc "qiniupkg.com/x/rpc.v7"
)

// PutRet upload response
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

// UploadFile upload file to server
func UploadFile(fileNmae string) {
	dir, err := os.Getwd()
	if err != nil {
		showTip("Failed to get file path❗ ️")
		os.Exit(0)
	}
	config, err := ReadFromFile()
	if err != nil {
		showTip("Failed to read configuration ❗️ ")
		showTip("If you don't have a configuration account, you should configure your account first ❗️ ")
		os.Exit(0)
	}
	conf.ACCESS_KEY = config.AccessKey
	conf.SECRET_KEY = config.SecretKey

	filePath := path.Join(dir, fileNmae)

	// create a client
	client := kodo.New(0, nil)
	// upload policy
	key := config.FormatUploadFileName(fileNmae)
	policy := &kodo.PutPolicy{
		Scope:   config.BucketName + ":" + key,
		Expires: 3600,
	}
	// make upload token
	token := client.MakeUptoken(policy)

	// build a uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	// call PutFile
	res := uploader.PutFile(nil, &ret, token, key, filePath, nil)

	if res != nil {
		showTip("Failed to upload file ❗️ ")
		showTip("Error:" + res.Error())
		return
	}

	mac := &digest.Mac{
		SecretKey: []byte(config.SecretKey),
		AccessKey: config.AccessKey,
	}
	domains, err := GetDomainsOfBucket(mac, config.BucketName)
	if res != nil {
		showTip("Failed to get domain ❗️ ")
		return
	}
	if len(domains) > 0 {
		domain := domains[0]
		fmt.Println()
		fmt.Println(fmt.Sprintf("URL: http://%s/%s", domain, key))
		fmt.Println(fmt.Sprintf("MarkDown: [](http://%s/%s)", domain, key))
		clipboard.WriteAll(fmt.Sprintf("[](http://%s/%s)", domain, key))
		showTip("MarkDown code has been copied to the clipboard 🍺 ")

	} else {
		showTip("Please make sure there is at least one domain ⚠️ ")
	}

}

// GetDomainsOfBucket get bucket domains
func GetDomainsOfBucket(mac *digest.Mac, bucket string) (domains []string, err error) {
	domains = make([]string, 0)
	client := rs.New(mac)
	getDomainsUrl := fmt.Sprintf("%s/v6/domain/list", "http://api.qiniu.com")
	postData := map[string][]string{
		"tbl": []string{bucket},
	}
	callErr := client.Conn.CallWithForm(nil, &domains, getDomainsUrl, postData)
	if callErr != nil {
		if v, ok := callErr.(*rpc.ErrorInfo); ok {
			err = fmt.Errorf("code: %d, %s, xreqid: %s", v.Code, v.Err, v.Reqid)
		} else {
			err = callErr
		}
	}
	return
}
