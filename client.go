package main

import (
  "fmt"
  "git.apache.org/thrift.git/lib/go/thrift"
)

func main(){
  transport, err := thrift.NewTSocket('localhost:9090')
  protocol := thrift.NewTBinaryProtocolTransport(transport)
  var pro TProtocolFactory
  transport = transportFactory.GetTransport(transport)
  
  defer transport.Close()
  if err := transport.Open(); err != nil {
  	return err
  }
  client := handleClient(users.NewUserServiceClientFactory(transport, pro))

  user := NewUser()
  user.FirstName = "Arun"
  user.LastName = "Arun"
  user.Email = "a@g.com"

  client.AddUser(user)

  users = client.GetUsers()
  for _, usr := range users{
  	fmt.Println(usr)
  }
}

