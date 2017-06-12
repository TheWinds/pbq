package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func init() {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal("can not get user home dir", err)
	}
	HomeDir := curUser.HomeDir
	ConfingFileDir = HomeDir + "/.pbq"
	AccountConfFilePath = ConfingFileDir + "/account.json"
	LayoutConfFilePath = ConfingFileDir + "/layout.json"
}

// HomeDir user home dir
var (
	ConfingFileDir      = ""
	AccountConfFilePath = ""
	LayoutConfFilePath  = ""
)

// const (
// 	// Slat key
// 	Slat = ")@!%#$&*(+_"
// )

// Account user account to access qiniu cloud
type Account struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// SaveAccount save account to ~/.pbq/account.json
func SaveAccount(account Account) error {
	account.SecretKey = base64.StdEncoding.EncodeToString([]byte(account.SecretKey))
	accountBytes, err := json.Marshal(&account)
	if err != nil {
		return err
	}
	return saveFile(AccountConfFilePath, accountBytes)
}

// ReadAccount read account from ~/.pbq/account.json
func ReadAccount() (*Account, error) {
	fileBytes, readErr := readFile(AccountConfFilePath)
	if readErr != nil {
		return nil, readErr
	}
	account := new(Account)
	marshalErr := json.Unmarshal(fileBytes, account)
	if marshalErr != nil {
		return nil, marshalErr
	}
	// decode secret key
	secretKeyBytes, decodeErr := base64.StdEncoding.DecodeString(account.SecretKey)
	if decodeErr != nil {
		return nil, decodeErr
	}
	account.SecretKey = string(secretKeyBytes)
	return account, nil
}

//Layout when we upload file,we will use this layout to format file name
type Layout struct {
	LayoutStr string
}

// Format format file name
func (layout *Layout) Format(fileNmae string) string {
	return ""
}

// SaveLayout save file name layout to ~/.pbq/layout.json
func SaveLayout(layout Layout) error {
	layoutBytes, err := json.Marshal(&layout)
	if err != nil {
		return err
	}
	return saveFile(LayoutConfFilePath, layoutBytes)
}

// ReadLayout read file name layout from ~/.pbq/layout.json
func ReadLayout() (*Layout, error) {
	fileBytes, readErr := readFile(LayoutConfFilePath)
	if readErr != nil {
		return nil, readErr
	}
	layout := new(Layout)
	marshalErr := json.Unmarshal(fileBytes, layout)
	if marshalErr != nil {
		return nil, marshalErr
	}
	return layout, nil
}

func saveFile(fileName string, datas []byte) error {
	// fmt.Println("dir:", dir)
	if _, err := os.Stat(ConfingFileDir); err != nil {
		err = os.MkdirAll(ConfingFileDir, 0755)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(fileName, datas, 0644)
}

func readFile(fileName string) ([]byte, error) {
	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}
	fileBytes, readErr := ioutil.ReadFile(fileName)
	if readErr != nil {
		return nil, readErr
	}
	return fileBytes, nil
}
