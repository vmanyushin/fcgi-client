// title       : fcgi-client.go
// description : make a request for fcgi application
// author      : v.manushin@corp.101xp.com
// date        : 2017-10-17
// version     : 0.1
// usage       : fcgi-client -h 127.0.0.1 -p 9001 -url /php-fpm_status

package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/flashmob/go-fastcgi-client"
)

func main() {
	scriptNamePtr := flag.String("url", "/status", "status url pm.status_page")
	portNumberPtr := flag.Int("port", 9000, "port number")
	hostAddrPtr := flag.String("hostname", "127.0.0.1", "hostname")
	flag.Parse()

	v := url.Values{}
	reqParams := v.Encode()

	env := make(map[string]string)

	env["REQUEST_METHOD"] = "GET"
	env["SCRIPT_FILENAME"] = *scriptNamePtr
	env["SCRIPT_NAME"] = *scriptNamePtr
	env["SERVER_PROTOCOL"] = "HTTP/1.0"
	env["QUERY_STRING"] = reqParams

	fcgi, err := fcgiclient.New(*hostAddrPtr, *portNumberPtr)

	if err != nil {
		fmt.Println("connection refused, check options")
		os.Exit(1)
	}

	content, _, err := fcgi.Request(env, reqParams)

	if err != nil {
		fmt.Println("check connection settings")
		os.Exit(1)
	}

	fmt.Printf("content: %s", content)
}
