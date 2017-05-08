package main

import (
  "encoding/json"
  "log"
  "encoding/base64"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
  simplejson "github.com/bitly/go-simplejson"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    data, _ := evt.MarshalJSON()
    js, _ := simplejson.NewJson(data)
    decoded := js.Get("Records").GetIndex(0).Get("kinesis").Get("data").MustString()
    payload, _ := base64.StdEncoding.DecodeString(decoded)
    log.Printf("Decoded payload : %s", string(payload))
    return string(payload), nil
}
