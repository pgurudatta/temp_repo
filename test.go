/*
Sample code for vulnerable type: Use of Password Hash With Insufficient Computational Effort
CWE : CWE-916
Description : Use of Password Hash With Insufficient Computational Effort
*/
package main
import (
	"crypto/md5"
	"fmt"
)
func main() {
	data := "SensitiveData123"
	hash := md5.Sum([]byte(data)) //source and sink
	hashString := fmt.Sprintf("%x", hash)
	fmt.Printf("MD5 Hash: %s\n", hashString)
}
