package main

import "socket5"

func main() {
	c := socket5.NewConfig()
	s := socket5.NewServer(*c, nil, nil, nil)
	s.ListenAndServe()
}
