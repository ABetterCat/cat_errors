package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// 包装，并且增加堆栈信息
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	data := []byte{}
	f.Read(data)
	return data, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, "config.xml"))
	return config, errors.WithMessage(err, "could nor read config")
}

// pkg/errors 里边的errors.new 是包括堆栈信息的
func main() {
	_, err := ReadConfig()
	if err != nil {
		// 根原因的类型和值
		fmt.Printf("original error: %T %v \n", errors.Cause(err), errors.Cause(err))
		//打印堆栈信息，注意+v 会打印多的信息
		fmt.Printf("stack trace:  \n %+v \n", err)
		os.Exit(1)
	}
}
