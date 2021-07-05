package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"regexp"
)

func main() {
	url := "https://template-homedecor.onshopbase.com/collections/new-arrivals?sortby=name%3Aasc"
	resp, _ := http.Get(url)
	b, _ := ioutil.ReadAll(resp.Body)
	// regexpn := regexp.MustCompile(`<a.*?href="/collections/(.*?)</a>`)
	regexpn := regexp.MustCompile(`<a\shref="/collections(.*)</a>`)
	moviename := regexpn.FindAllStringSubmatch(string(b), -1)
	fmt.Println(moviename)
}