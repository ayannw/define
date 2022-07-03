package lib

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "log"
)

type Config struct {
  App_ID  string `json:"app_id"`
  App_key string `json:"app_key"`
  Cache   bool   `json:"cache"`
  Colors  bool   `json:"colors"`
}

func GetConf() (conf Config) {

  homedir, h_err := os.UserHomeDir()

  if h_err != nil {
    log.Fatal("Unable to detect homedir.")
  }

  file, f_err := os.Open(homedir + "/.config/.define.conf.json")

  if f_err != nil {
    log.Fatal(f_err)
  }

  if h_err != nil {
    log.Fatal("Unable to read .define.conf.json")
  }

  defer file.Close()

  bytes, b_err := ioutil.ReadAll(file)

  if b_err != nil {
    log.Fatal(b_err)
  }

  json.Unmarshal(bytes, &conf)

  return conf
}


