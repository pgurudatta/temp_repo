package main

import (
    "crypto/md5"
    "fmt"
)

func main() {
    data := "Hello, world!"

    // Using MD5 (which is considered broken for cryptographic purposes)
    hash := md5.Sum([]byte(data))
    
    fmt.Printf("MD5 hash of '%s': %x\n", data, hash)
}
