package main

type Value struct {
	Name   string `json:name`
	Binary bool   `json:binary`
	Main   string `json:main`
	Code   string `json:code`
}

type Init struct {
	Value Value
}
