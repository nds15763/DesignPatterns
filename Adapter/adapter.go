package Adapter

import "fmt"

type IFile interface {
	Read()
}

type ExcelFile struct {
}

func (excel *ExcelFile) ExcelRead() {
	fmt.Println("读取excel")
}

type CsvFile struct {
}

func (csv *CsvFile) CsvRead() {
	fmt.Println("读取csv")
}

type File struct {
	ftype string
	*ExcelFile
	*CsvFile
}

func (file *File) Read() {
	file.ExcelRead()
}
