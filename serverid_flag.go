package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ServerIDList []int

func (s *ServerIDList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *ServerIDList) Set(arg string) error {
	for _, str := range strings.Split(arg, ",") {
		id, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*s = append(*s, id)
	}
	return nil
}
