package main

type command struct {
	id         ID
	receipient string
	sender     string
	body       []byte
}

type ID int

const (
  REG ID = iota
  JOIN
  LEAVE
  MSG
  CHNS
  USRS
)
