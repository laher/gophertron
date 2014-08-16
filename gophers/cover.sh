#!/bin/bash

go test -coverprofile=/tmp/cover.out .
go tool cover -func=/tmp/cover.out
go tool cover -html=/tmp/cover.out
