package elemhttp

import (
  "fmt"
  "errors"
  "io/ioutil"
  "time"
  "net/http"
  "net/url"
)

func Post (urlString string, obj []byte, maxRetry int) ([]byte, error) {
  postObj := url.Values{"tmsp": {string(obj)}}

  client := &http.Client{
     Timeout: time.Duration(5 * time.Second),
  }

  resp, err := client.PostForm(urlString, postObj)
  if err != nil {
    if maxRetry <= 0 {
      fmt.Println(fmt.Sprintf("Error reaching smart contract engine: %v", err))
      return nil, errors.New(fmt.Sprintf("Error reaching smart contract engine: %v", err))
    }
    fmt.Println("Retrying...")
    maxRetry--
    response, error := Post(urlString, obj, maxRetry)
    return response, error
  }

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  if resp.StatusCode < 400 {
    return body, nil
  }

  return nil, errors.New(fmt.Sprintf("%v", resp.StatusCode))
}
