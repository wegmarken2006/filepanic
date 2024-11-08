package filepanic

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"runtime"
)

type File struct {
	File *os.File
}

func Open(filename string) File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return File{File: file}
}

func Create(filename string) File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return File{File: file}
}

func ReadFile(filename string) []byte {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func ReadCommentedCSV(filename string, delimiter rune, comment rune) [][]string {
	var out [][]string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = delimiter
	r.Comment = comment

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		out = append(out, record)
	}
	return out
}

func FilesInDir(dirpath string) []string {
	var out []string

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			out = append(out, file.Name())
		}
	}

	return out
}

func DirsInDir(dirpath string) []string {
	var out []string

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			out = append(out, file.Name())
		}
	}

	return out
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}

func Mkdir(name string) {
	var err error
	if runtime.GOOS == "linux" {
		err = os.Mkdir(name, 0700)
	} else {
		err = os.Mkdir(name, fs.ModeDir)
	}
	
	if err != nil {
		panic(err)
	}
}

func ReadCSV(filename string, delimiter rune) [][]string {
	var out [][]string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = delimiter
	r.LazyQuotes = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		out = append(out, record)
	}
	return out
}

func (fpfile *File) Write(bytes []byte) int {
	n, err := fpfile.File.Write(bytes)
	if err != nil {
		panic(err)
	}
	return n
}

func (fpfile *File) Read(bytes []byte) int {
	n, err := fpfile.File.Read(bytes)
	if (err != nil) && (err != io.EOF){
		panic(err)
	}
	return n
}

func (fpfile *File) ReadLines() []string {
	scanner := bufio.NewScanner(fpfile.File)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	return text
}

func (fpfile *File) Seek(offset int64, whence int) int64 {
	x, err := fpfile.File.Seek(offset, whence)
	if err != nil {
		panic(err)
	}
	return x
}

func (fpfile *File) Close() {
	err := fpfile.File.Close()
	if err != nil {
		panic(err)
	}
}

func (fpfile *File) Metadata() fs.FileInfo {
	metadata, err := fpfile.File.Stat()
	if err != nil {
		panic(err)
	}
	return metadata
}
