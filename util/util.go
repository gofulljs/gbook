package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

// Active 使某个path生效，即把path放在env的PATH最前面
func Active(path string) {
	newEnvPath := strings.Join([]string{path, os.Getenv("PATH")}, string(filepath.ListSeparator))
	os.Setenv("PATH", newEnvPath)
}

/**
 * @description: 获取路径信息
 * @param {string} path
 * @return {*}
 */
func GetPathInfo(path string) (fs fs.FileInfo, exist bool, err error) {
	fs, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, xerrors.Errorf("%w, err")
	}

	return fs, true, nil
}

/**
 * @description: 获取文件是否存在
 * @param {string} filepath
 * @return {*}
 */
func GetFileExist(filepath string) (bool, error) {
	fs, exist, err := GetPathInfo(filepath)
	if err != nil {
		return false, xerrors.Errorf("%w, err")
	}
	if exist {
		if fs.IsDir() {
			return false, xerrors.Errorf("filepath %v is a directory", filepath)
		}
	}

	return exist, nil
}

/**
 * @description: 删除文件
 * @param {string} fileFullPath
 * @return {*}
 */
func DeleteFileIfExist(fileFullPath string) error {
	isExist, err := GetFileExist(fileFullPath)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if isExist {
		err = os.Remove(fileFullPath)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
	}

	return nil
}
