#!/bin/bash

include .env
export

run:
	go run main.go

test:
	go test ./tests/... -v