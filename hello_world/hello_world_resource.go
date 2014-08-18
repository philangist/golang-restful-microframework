package main

import (
  "net/url"
  "philangist.github.com/restful-micro-framework/sleepy"
  )

type HelloResource struct {
  sleepy.PostNotSupported
  sleepy.PutNotSupported
  sleepy.DeleteNotSupported
}

func (HelloResource) Get(values url.Values) (int, interface{}) {
  data := map[string]string{"hello": "world"}
  return 200, data
}

func main(){
  helloResource := new(HelloResource)

  var api = new(sleepy.API)
  api.AddResource(helloResource, "/hello")
  api.Start(3000)
}
