package main

import (
	"fmt"
	"initial/wallet"
	"log"
)

func initial() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	w := wallet.NewWallet()
	//fmt.Println(w.PublicKey())
	//fmt.Println(w.PrivateKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println("==========")
	fmt.Println(w.PublicKeyStr())
}
