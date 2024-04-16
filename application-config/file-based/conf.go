package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}

func main() {
	configFileName, found := os.LookupEnv("MYVAR")
	if !found {
		fmt.Println("here")
		configFileName = "local"
	}
	fmt.Println("configuration file", configFileName)
	confFile, err := os.Open("./" + configFileName + ".json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	conf, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}
	myConf := Configuration{}
	err = json.Unmarshal(conf, &myConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(myConf)
}
