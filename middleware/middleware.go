package middleware

import (
  "fmt"
  "encoding/binary"
  "github.com/sahilkathpal/blockchain_engine/lib"
  . "github.com/tendermint/go-common"
	"github.com/tendermint/abci/types"
)

type MiddlewareApplication struct {
  hashCount int
  txCount int
  url string
}

func NewMiddlewareApplication (url string) *MiddlewareApplication {
  return &MiddlewareApplication {url: url}
}

func (app *MiddlewareApplication) Info(types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{Data: Fmt("hashes:%v, txs:%v", app.hashCount, app.txCount)}
}

func (app *MiddlewareApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {
  body, err := elemhttp.Post(app.url+"/append", tx, 3)
  if err != nil {
    fmt.Printf("DeliverTx error from Iris: %v", err)
    return types.ResponseCheckTx{Data: err, Code: 1}
  }

  app.txCount += 1

    return types.ResponseDeliverTx{Data: body, Code: 0}
}

func (app *MiddlewareApplication) CheckTx(tx []byte) types.ResponseCheckTx {
  body, err := elemhttp.Post(app.url+"/check", tx, 3)
  if err != nil {
    fmt.Printf("CheckTx error from Iris: %v", err)
    return types.ResponseCheckTx{Data: err, Code: 1}
  }

  return types.ResponseCheckTx{Data: body}
}


func (app *MiddlewareApplication) Commit() types.ResponseCommit {
	app.hashCount += 1

	if app.txCount == 0 {
		return types.ResponseCommit{}
	} else {
		hash := make([]byte, 8)
		binary.BigEndian.PutUint64(hash, uint64(app.txCount))
		return types.ResponseCommit{Data: hash}
	}
}

func (app *MiddlewareApplication) Query(query types.RequestQuery) types.ResponseQuery {
	return types.ResponseQuery {
    Log: Fmt("Query is not supported"),
  }
}

func (app *MiddlewareApplication) SetOption(key string, value string) (log string) {
	return ""
}

func (app *MiddlewareApplication) BeginBlock(params types.RequestBeginBlock) {
  fmt.Printf("Begin Block Baby!")
  return
}

func (app *MiddlewareApplication) EndBlock(height uint64) (resEndBlock types.ResponseEndBlock) {
	return
}

func (app *MiddlewareApplication) InitChain(validators types.RequestInitChain) {
  fmt.Println("Finally in InitChain")
}

}
