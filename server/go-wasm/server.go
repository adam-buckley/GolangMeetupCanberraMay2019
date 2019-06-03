package main

import _ "net/http/pprof"  
import (
	"log"
	"net/http"
)

func main() {
// go func() {  
	log.Println(http.ListenAndServe("localhost:6060", nil))  
// }()
}
