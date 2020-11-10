package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ClientHttp(arrStr, api string) {
	fmt.Println("传入：" + arrStr)
	pool := x509.NewCertPool()
	crtPath := "server.crt"
	crtFile, err := ioutil.ReadFile(crtPath)
	if err != nil {
		fmt.Println("crt file read err:", err.Error())
		return
	}
	pool.AppendCertsFromPEM(crtFile)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{
		Transport: tr,
	}
	form := url.Values{}
	form.Set("newarr", arrStr)
	b := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", "https://localhost:8888"+api, b)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
