package file

import (
	"os"
	"path"
)

// 追加方式打开文件
func OpenAppendFile(fileName string, filePath string) *os.File {
	// 日志文件完整路径
	logFullPath := path.Join(filePath, fileName)

	// 返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(logFullPath)
	switch {
	case os.IsNotExist(err):
		// 创建目录，完整目录路径，权限os.ModePerm为0777
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			panic(err)
		}
	case os.IsPermission(err):
		panic(err)
	}

	// 以追加或创建或写方式打开文件，没有就创建
	file, err := os.OpenFile(logFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// 打开文件
func OpenFile(fileName string, filePath string) *os.File {
	// 日志文件完整路径
	logFullPath := path.Join(filePath, fileName)

	// 返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(logFullPath)
	switch {
	case os.IsNotExist(err):
		// 创建目录，完整目录路径，权限os.ModePerm为0777
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			panic(err)
		}
	case os.IsPermission(err):
		panic(err)
	}

	// 以追加或创建或写方式打开文件，没有就创建
	file, err := os.OpenFile(logFullPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}
