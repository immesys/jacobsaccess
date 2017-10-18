package main

import (
	"fmt"
	"os"
	"time"

	"github.com/immesys/ragent/ragentlib"
)

// ragent.cal-sdb.org:28590 MT3dKUYB8cnIfsbnPrrgy8Cb_8whVKM-Gtg2qd79Xco=
const ourEntity = ("\x32\x1f\x8f\x00\x47\x22\xb9\x74\x40\xc6\x6a\x09\x2a\xa9\x9c\xa7\x35\xe4\xe1" +
	"\x2d\xfc\xaf\x39\x32\x4d\x9c\xe8\xda\xe6\x8e\xe3\xee\xff\x81\x8c\xec\x81\xd3" +
	"\x71\x09\xc0\xd7\x28\x3a\xe4\x5e\x20\x61\xf5\x65\xd7\x45\x73\x1c\xa7\x43\x36" +
	"\x86\x7d\x17\x11\x42\x7a\x0b\xdb\x02\x08\x07\x94\x3a\x2f\x7c\xc1\xee\x14\x03" +
	"\x08\x1d\x92\x98\xf7\xbb\x23\x4f\x19\x06\x0c\x4a\x61\x63\x6f\x62\x73\x41\x63" +
	"\x63\x65\x73\x73\x00\x46\x69\x8c\xf8\xda\xdf\x07\xf2\xcb\x2b\xf9\xca\xcc\xdd" +
	"\x9c\xf1\x29\x1e\x61\x28\x57\x0c\x9a\xbc\x35\x13\xae\xc6\x72\x3e\xa8\xd4\x4a" +
	"\x7a\x29\x3a\xdc\xe8\xe7\x71\xeb\x79\xfb\xef\x71\x79\x98\xe2\xac\x70\x01\xab" +
	"\xd1\x84\x35\xb3\xe1\x65\xe5\x91\x56\xab\x1d\x08")

func main() {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Printf("failed to connect to remote server (check your internet): %v", r)
				os.Exit(1)
			}
		}()
		ragentlib.DoClientER([]byte(ourEntity), "ragent.cal-sdb.org:28590", "MT3dKUYB8cnIfsbnPrrgy8Cb_8whVKM-Gtg2qd79Xco=", "127.0.0.1:28588")
	}()
	time.Sleep(200 * time.Millisecond)
	tap("hamilton/sensors/jacobs/*")
}