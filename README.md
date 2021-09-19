# filepanic

Wrapper around some file library functions and other file utilities aiming to simplify error handling. Any error causes panic.

## Example
``` go
//go.mode file
module testfp

require github.com/wegmarken2006/filepanic v1.0.2
//uncomment the line below when locally developing filepanic module
//replace github.com/wegmarken2006/filepanic v1.0.2 => ./filepanic

go 1.17
```

``` go
//testfp.go file
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

	fp.Mkdir("tmp")
	text = fp.FilesInDir("tmp")
	Println(text)
	text = fp.DirsInDir("./")
	Println(text)
}
```
