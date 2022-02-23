package oss

import (
	"fmt"
	"testing"
)

func TestUploadContent(t *testing.T) {
	file := File{}
	fmt.Println(file.UploadWithContent("testFile", []byte("testFile--First"), nil))
}

func TestUploadWithURL(t *testing.T) {
	file := File{}
	fmt.Println(file.UploadWithURL("testFile", "http://www.juimg.com/tuku/yulantu/140112/328648-14011213253758.jpg", nil))
}

func TestUploadWithPath(t *testing.T) {
	file := File{}
	fmt.Println(file.UploadWithPath("testFile", "./file_test.go", nil))
}
