package elemhttp

import (
  "fmt"
  "bytes"
  "errors"
  "io/ioutil"
  "time"
  "net/http"
)

func Post (url, header string, obj []byte, maxRetry int) ([]byte, error) {
  postObj := bytes.NewBuffer(obj)

  client := &http.Client{
     Timeout: time.Duration(5 * time.Second),
  }

  resp, err := client.Post(url, header, postObj)
  if err != nil {
    if maxRetry == 0 {
      return nil, errors.New(fmt.Sprintf("Error reaching smart contract engine: %v", err))
    }
    fmt.Println("Retrying...")
    maxRetry--
    response, error := Post(url, header, obj, maxRetry)
    return response, error
  }

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  if resp.StatusCode < 400 {
    return body, nil
  }

  return nil, errors.New(fmt.Sprintf("%v", resp.StatusCode))
}
