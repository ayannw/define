package lib

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func FetchDefinition(word string) (resp OxfordResponse) {

  conf := GetConf()
  url := "https://od-api.oxforddictionaries.com/api/v2/entries/en-gb/" + word

  client := http.Client{ Timeout: time.Second * 5, }
  req, req_err := http.NewRequest(http.MethodGet, url, nil)

  if req_err != nil { log.Fatal(req_err) }

  req.Header = http.Header{
    "User-Agent": { "define_app" },
    "app_id": { conf.App_ID },
    "app_key": { conf.App_key },
  }

  res, res_err := client.Do(req)

  if res_err != nil { log.Fatal(res_err) }
 
  if res.Body != nil { defer res.Body.Close() }

  body, read_err := ioutil.ReadAll(res.Body)

  if read_err != nil { log.Fatal(read_err) }

  json_err := json.Unmarshal(body, &resp)

  if json_err != nil { log.Fatal(json_err) }

  return
}
