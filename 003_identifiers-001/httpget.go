package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {

	/*blank identifier allows us to tell the compiler that
	  we're not using something. In this case "err" which we're telling the compiler to
	  throw away the error code.*/
	res, _ := http.Get("https://git.thelema.academy")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
