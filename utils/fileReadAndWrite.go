package utils

import (
	"io/ioutil"
	"os"
)

//读取文件
func ReadFile(filepath string) (*string, error) {
	//打开文件
	fi, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	//读取内容
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	err = fi.Close()
	if err != nil {
		return nil, err
	}

	fileVale := string(fd)
	return &fileVale, nil
}


//写入文件
func WriteStringToFile(filepath, content string) error {
	//打开文件，没有则创建，有则append内容
	w1, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	_, err = w1.Write([]byte(content))
	if err != nil {
		return err
	}

	return w1.Close()
}


//写入文件 性能更好
func WriteBytesToFile(filepath string, content []byte) error {
	//打开文件，没有此文件则创建文件，将写入的内容append进去
	w1, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	_, err = w1.Write(content)
	if err != nil {
		return err
	}

	return w1.Close()
}