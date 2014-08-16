gophertron
==========

This package is just the code for a short talk on Go's coverage testing & the [httptest](http://golang.org/pkg/net/http/httptest/) package.

Gophertron is a JSON API for spawning, mutating, zapping & kapowing gophers.

Pre-requisite: install Go. Now install & run gophertron:

	go get github.com/laher/gophertron
	gophertron

To test, cd into the gophertron/gophers directory, and run `go test .`. 

For coverage testing, run ./cover.sh, or the following (for Windows you'll need to change the location of the cover.out file):

	go test -coverprofile=/tmp/cover.out .
	go tool cover -func=/tmp/cover.out
	go tool cover -html=/tmp/cover.out

You can also [visit gocover.io to see the coverage report](http://gocover.io/github.com/laher/gophertron/gophers) ![coverage status](http://gocover.io/_badge/github.com/laher/gophertron/gophers?)

Note: There is one bug which has been intentionally left uncaptured by the tests. See `Gopher.Kapow()`
