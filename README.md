# filepanic

Wrapper around some file library functions and other file utilities aiming to simplify error handling. Any error causes panic.

## Example
``` go
package main

import (
	. "fmt"

	fp "github.com/wegmarken2006/filepanic"
)

func main() {
	filename := "tmp01.txt"
	file := fp.Create(filename)
	file.Write([]byte("hello"))
	file.Close()

	buf := make([]byte, 10)
	file = fp.Open(filename)
	file.Read(buf)
	Println(string(buf))

	file.Seek(0, 0)
	text := file.ReadLines()
	Println(text)
}
```
