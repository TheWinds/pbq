package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

func init() {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal("can not get user home dir", err)
	}
	HomeDir := curUser.HomeDir
	ConfigFileDir = HomeDir + "/.pbq"
	ConfigFilePath = ConfigFileDir + "/config.json"
}

// DefaultLayout default file format layout
const DefaultLayout = "%YYYY%MM%DD%UNIX-%FILENAME"

// HomeDir user home dir
var (
	ConfigFileDir  = ""
	ConfigFilePath = ""
)

// Config user config
type Config struct {
	// user account to access qiniu cloud
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	// when we upload file,we will use this layout to format file name
	UploadNameLayout string `json:"upload_name_layout"`
	// which bucket to upload picture
	BucketName string
}

// SaveToFile save config to ~/.pbq/config.json
func (config *Config) SaveToFile() error {
	config.SecretKey = base64.StdEncoding.EncodeToString([]byte(config.SecretKey))
	configBytes, err := json.Marshal(&config)
	if err != nil {
		return err
	}
	return saveFile(ConfigFilePath, configBytes)
}

// ReadFromFile read config from ~/.pbq/config.json
func ReadFromFile() (*Config, error) {
	fileBytes, readErr := readFile(ConfigFilePath)
	if readErr != nil {
		return nil, readErr
	}
	newConfig := new(Config)
	marshalErr := json.Unmarshal(fileBytes, newConfig)
	if marshalErr != nil {
		return nil, marshalErr
	}
	// decode secret key
	secretKeyBytes, decodeErr := base64.StdEncoding.DecodeString(newConfig.SecretKey)
	if decodeErr != nil {
		return nil, decodeErr
	}
	newConfig.SecretKey = string(secretKeyBytes)
	return newConfig, nil
}

// FormatUploadFileName format upload fileName
func (config *Config) FormatUploadFileName(fileName string) string {
	mapTimeToken := make(map[string]string, 0)
	mapTimeToken["%YYYY"] = "2006"
	mapTimeToken["%MM"] = "01"
	mapTimeToken["%DD"] = "02"

	layout := config.UploadNameLayout
	if !strings.Contains(layout, "%FILENAME") {
		layout = DefaultLayout
	}
	for k, v := range mapTimeToken {
		layout = strings.Replace(layout, k, v, 1)
	}
	timeNow := time.Now()

	ret := timeNow.Format(layout)
	ret = strings.Replace(ret, "%FILENAME", fileName, 1)
	ret = strings.Replace(ret, "%UNIX", fmt.Sprintf("%d", timeNow.Unix()), 1)

	return ret
}

func saveFile(fileName string, datas []byte) error {

	if _, err := os.Stat(ConfigFileDir); err != nil {
		err = os.MkdirAll(ConfigFileDir, 0755)
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
