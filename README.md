gophertron
==========

This package is just the code for a short talk on Go's coverage testing & the [httptest](http://golang.org/pkg/net/http/httptest/) package. I'll probably keep adding features for future discussions.

Gophertron is a JSON API for spawning, mutating, zapping & kapowing gophers.

Pre-requisite: install Go. Now install & run gophertron:

	go get github.com/laher/gophertron
	gophertron

To test, cd into the gophertron/gophers directory, and run `go test .`. 

Coverage isn't the be-all and end-all: there's a bug which has been intentionally left uncaptured by the tests, and won't be captured by coverage testing. See `Gopher.Kapow()`. 


httptest
--------

httptest is a small package providing 2 utilities for testing HTTP functionality. 
 
 * ResponseRecorder helps you test your http handlers. Sample code is in gophers/gopher_api_test.go. [httptest.ResponseRecorder](http://golang.org/pkg/net/http/httptest/#ResponseRecorder) implements the [http.ResponseWriter](http://golang.org/pkg/net/http/#ResponseWriter) interface, so that you can pass it to any http handler.
 * The second utility is an actual server. The system chooses a port itself, and supplies information for connecting. I haven't included examples here.

Note that my use of ResponseRecorder is for Martini handlers rather than actual [http.HandlerFunc](http://golang.org/pkg/net/http/#HandlerFunc)s, but the idea is the same.

Code Coverage
-------------

For coverage testing, you'll need to install go's coverage tooling, part of the 'go.tools' repo:

	go get code.google.com/p/go.tools/cmd/cover

(If you don't have it, the `go` tool will tell you where to get it as above).

At its most basic level, you can just run `go test -cover .`. This gives you percentages.

Now you can either run ./cover.sh, or the following (for Windows you'll need to change the location of the cover.out file):

	go test -coverprofile=/tmp/cover.out .
	go tool cover -func=/tmp/cover.out
	go tool cover -html=/tmp/cover.out

You can also [visit gocover.io to see the coverage report](http://gocover.io/github.com/laher/gophertron/gophers) ![coverage status](http://gocover.io/_badge/github.com/laher/gophertron/gophers?)


Reading on Go coverage testing:

 * [Go 1.2 Announcement](http://golang.org/doc/go1.2#cover)
 * [The cover story](http://blog.golang.org/cover)


### Go's own code coverage

Here's the output of coverage testing from go's source tree (v1.3):

`
../go-1.3/src/pkg$ go test -cover ./...
ok  	archive/tar	0.020s	coverage: 84.8% of statements
ok  	archive/zip	13.773s	coverage: 85.6% of statements
ok  	bufio	0.103s	coverage: 89.2% of statements
?   	builtin	[no test files]
ok  	bytes	1.530s	coverage: 94.8% of statements
ok  	compress/bzip2	0.073s	coverage: 88.1% of statements
ok  	compress/flate	18.315s	coverage: 91.1% of statements
ok  	compress/gzip	0.035s	coverage: 84.5% of statements
ok  	compress/lzw	0.108s	coverage: 86.9% of statements
ok  	compress/zlib	1.840s	coverage: 81.0% of statements
ok  	container/heap	0.004s	coverage: 100.0% of statements
ok  	container/list	0.003s	coverage: 100.0% of statements
ok  	container/ring	0.027s	coverage: 100.0% of statements
?   	crypto	[no test files]
ok  	crypto/aes	0.034s	coverage: 90.6% of statements
ok  	crypto/cipher	0.005s	coverage: 86.5% of statements
ok  	crypto/des	0.013s	coverage: 95.3% of statements
ok  	crypto/dsa	63.168s	coverage: 87.7% of statements
ok  	crypto/ecdsa	7.265s	coverage: 88.7% of statements
ok  	crypto/elliptic	1.112s	coverage: 96.8% of statements
ok  	crypto/hmac	0.004s	coverage: 100.0% of statements
ok  	crypto/md5	0.009s	coverage: 95.8% of statements
ok  	crypto/rand	0.640s	coverage: 57.8% of statements
ok  	crypto/rc4	0.135s	coverage: 76.0% of statements
ok  	crypto/rsa	1.883s	coverage: 83.6% of statements
ok  	crypto/sha1	0.004s	coverage: 99.1% of statements
ok  	crypto/sha256	0.004s	coverage: 98.8% of statements
ok  	crypto/sha512	0.006s	coverage: 97.8% of statements
ok  	crypto/subtle	0.010s	coverage: 92.6% of statements
ok  	crypto/tls	1.167s	coverage: 81.8% of statements
ok  	crypto/x509	1.216s	coverage: 80.2% of statements
?   	crypto/x509/pkix	[no test files]
ok  	database/sql	0.491s	coverage: 84.4% of statements
ok  	database/sql/driver	0.006s	coverage: 41.5% of statements
ok  	debug/dwarf	0.026s	coverage: 70.0% of statements
ok  	debug/elf	0.033s	coverage: 63.1% of statements
ok  	debug/gosym	0.216s	coverage: 37.6% of statements
ok  	debug/macho	0.043s	coverage: 64.6% of statements
ok  	debug/pe	0.037s	coverage: 47.3% of statements
ok  	debug/plan9obj	0.005s	coverage: 28.0% of statements
?   	encoding	[no test files]
ok  	encoding/ascii85	0.004s	coverage: 93.5% of statements
ok  	encoding/asn1	0.008s	coverage: 81.4% of statements
ok  	encoding/base32	0.004s	coverage: 96.2% of statements
ok  	encoding/base64	0.011s	coverage: 96.2% of statements
ok  	encoding/binary	0.010s	coverage: 80.1% of statements
ok  	encoding/csv	0.005s	coverage: 93.0% of statements
ok  	encoding/gob	3.283s	coverage: 91.0% of statements
ok  	encoding/hex	0.004s	coverage: 94.0% of statements
ok  	encoding/json	2.960s	coverage: 89.6% of statements
ok  	encoding/pem	0.004s	coverage: 86.1% of statements
ok  	encoding/xml	0.016s	coverage: 86.2% of statements
ok  	errors	0.004s	coverage: 100.0% of statements
ok  	expvar	0.007s	coverage: 84.9% of statements
ok  	flag	0.004s	coverage: 85.8% of statements
ok  	fmt	0.041s	coverage: 91.9% of statements
ok  	go/ast	0.006s	coverage: 45.2% of statements
ok  	go/build	10.860s	coverage: 76.6% of statements
ok  	go/doc	0.138s	coverage: 81.7% of statements
ok  	go/format	0.008s	coverage: 82.1% of statements
ok  	go/parser	0.057s	coverage: 84.9% of statements
ok  	go/printer	0.374s	coverage: 92.7% of statements
ok  	go/scanner	0.010s	coverage: 94.7% of statements
ok  	go/token	0.036s	coverage: 75.1% of statements
?   	hash	[no test files]
ok  	hash/adler32	0.012s	coverage: 54.2% of statements
ok  	hash/crc32	0.003s	coverage: 73.0% of statements
ok  	hash/crc64	0.004s	coverage: 70.8% of statements
ok  	hash/fnv	0.003s	coverage: 92.9% of statements
ok  	html	0.005s	coverage: 93.4% of statements
ok  	html/template	0.050s	coverage: 92.2% of statements
ok  	image	0.457s	coverage: 66.6% of statements
ok  	image/color	0.010s	coverage: 18.9% of statements
?   	image/color/palette	[no test files]
ok  	image/draw	0.087s	coverage: 85.2% of statements
ok  	image/gif	0.071s	coverage: 76.6% of statements
ok  	image/jpeg	0.198s	coverage: 85.2% of statements
ok  	image/png	0.114s	coverage: 82.6% of statements
ok  	index/suffixarray	0.010s	coverage: 93.9% of statements
ok  	io	0.019s	coverage: 91.0% of statements
ok  	io/ioutil	0.007s	coverage: 67.4% of statements
ok  	log	0.008s	coverage: 68.3% of statements
ok  	log/syslog	2.029s	coverage: 65.9% of statements
ok  	math	0.005s	coverage: 76.8% of statements
ok  	math/big	206.916s	coverage: 93.8% of statements
ok  	math/cmplx	0.007s	coverage: 82.7% of statements
ok  	math/rand	0.634s	coverage: 71.1% of statements
ok  	mime	0.028s	coverage: 94.1% of statements
ok  	mime/multipart	0.395s	coverage: 89.6% of statements
SIGQUIT: quit
PC=0x431b01

goroutine 0 [idle]:
runtime.futex(0x7d4eb0, 0x0, 0x0, 0x0)

.... (output truncated) ....

* Test killed with quit: ran too long (10m0s).
FAIL	net	600.016s
ok  	net/http	43.059s	coverage: 86.4% of statements
ok  	net/http/cgi	0.365s	coverage: 69.5% of statements
ok  	net/http/cookiejar	0.007s	coverage: 95.3% of statements
ok  	net/http/fcgi	0.011s	coverage: 30.8% of statements
ok  	net/http/httptest	0.306s	coverage: 60.5% of statements
ok  	net/http/httputil	0.014s	coverage: 47.7% of statements
?   	net/http/pprof	[no test files]
ok  	net/mail	0.005s	coverage: 86.5% of statements
ok  	net/rpc	0.033s	coverage: 76.7% of statements
ok  	net/rpc/jsonrpc	0.007s	coverage: 90.7% of statements
ok  	net/smtp	0.030s	coverage: 82.7% of statements
ok  	net/textproto	0.005s	coverage: 70.0% of statements
ok  	net/url	0.005s	coverage: 89.3% of statements
ok  	os	0.876s	coverage: 67.9% of statements
SIGQUIT: quit
PC=0x42e6c1

goroutine 0 [idle]:
runtime.futex(0x8dc350, 0x0, 0x0, 0x0)
...
Test killed with quit: ran too long (10m0s).
FAIL	os/exec	600.188s
ok  	os/signal	3.327s	coverage: 87.5% of statements
ok  	os/user	0.005s	coverage: 76.9% of statements
ok  	path	0.004s	coverage: 96.4% of statements
ok  	path/filepath	0.163s	coverage: 93.4% of statements
ok  	reflect	1.250s	coverage: 85.1% of statements
ok  	regexp	45.264s	coverage: 91.2% of statements
ok  	regexp/syntax	0.585s	coverage: 85.4% of statements
ok  	runtime	411.436s	coverage: 68.5% of statements
?   	runtime/cgo	[no test files]
ok  	runtime/debug	0.140s	coverage: 85.2% of statements
ok  	runtime/pprof	48.347s	coverage: 30.3% of statements
?   	runtime/race	[no test files]
ok  	sort	1.418s	coverage: 98.4% of statements
ok  	strconv	3.764s	coverage: 96.5% of statements
ok  	strings	1.729s	coverage: 96.9% of statements
ok  	sync	0.753s	coverage: 74.4% of statements
ok  	sync/atomic	1.546s	coverage: 0.0% of statements
ok  	syscall	0.060s	coverage: 23.6% of statements
ok  	testing	2.105s	coverage: 45.4% of statements
?   	testing/iotest	[no test files]
ok  	testing/quick	0.031s	coverage: 77.4% of statements
ok  	text/scanner	0.005s	coverage: 95.7% of statements
ok  	text/tabwriter	0.005s	coverage: 89.9% of statements
ok  	text/template	0.147s	coverage: 84.3% of statements
ok  	text/template/parse	0.009s	coverage: 93.5% of statements
ok  	time	9.409s	coverage: 90.3% of statements
ok  	unicode	0.004s	coverage: 89.7% of statements
ok  	unicode/utf16	0.003s	coverage: 100.0% of statements
ok  	unicode/utf8	0.006s	coverage: 97.8% of statements
?   	unsafe	[no test files]
`
