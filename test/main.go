package main

import (
	"net/http"
	"math/rand"
	"time"
	"fmt"
)

func sayBitcoin(w http.ResponseWriter, r *http.Request){
	time.Sleep(1 * time.Second)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	currencyVal := 6000.0 + r1.Float64() * 1000.0 
	jpVal := 111.50 * currencyVal
	eurVal :=  0.89 * currencyVal
	json := fmt.Sprintf(`{"USD":%.2f,"JPY":%.2f, "EUR":%.2f}`, currencyVal, jpVal, eurVal)
	w.Write([]byte(json))
}

func main() {
	http.HandleFunc("/",sayBitcoin)
	if err := http.ListenAndServe(":8081",nil); err != nil {
		panic(err)
	}	
}
