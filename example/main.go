package main

import (
	"github.com/xaitx/socket5"
)

func main() {
	c := socket5.NewConfig()
	s := socket5.NewServer(*c, &socket5.PasswordAuth{Username: "admin", Password: "admin"}, nil, nil)
	s.ListenAndServe()

}
