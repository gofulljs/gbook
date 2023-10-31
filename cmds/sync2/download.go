package sync2

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofulljs/gbook/global"
	"golang.org/x/xerrors"
)

// Download 3.2.3
func downloadGitbook(rootPath, sourceURI string) error {
	resp, err := http.Get(sourceURI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return tarGzUnzip(resp.Body, rootPath)
}

func MustDownloadGitbook(rootPath, proxy1, proxy2, sourceURI string) error {

	domain := "https://github.com/"
	subURI := strings.TrimPrefix(sourceURI, domain)

	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	if proxy1 != "" {
		proxy1 = strings.TrimSuffix(proxy1, "/")
		source, err := url.JoinPath(proxy1+"/"+domain, subURI)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		return downloadGitbook(rootPath, source)
	}

	if proxy2 != "" {
		source, err := url.JoinPath(proxy2, subURI)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		return downloadGitbook(rootPath, source)
	}

	// proxy1 download
	for i := 0; i < len(global.Proxy1s); i++ {
		proxy := strings.TrimSuffix(global.Proxy1s[i], "/")
		source := proxy + "/" + sourceURI

		err = downloadGitbook(rootPath, source)
		if err != nil {
			fmt.Printf("source: %v failed, err: %v\n", source, err)
			continue
		}
		fmt.Printf("use source:%v success\n", source)
		return nil
	}

	return downloadGitbook(rootPath, sourceURI)
}

// tarGzUnzip Decompress the tar.gz file
func tarGzUnzip(r io.Reader, dest string) error {
	// 创建一个 gzip 读取器
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	defer gzipReader.Close()

	// 创建一个 tar 读取器
	tarReader := tar.NewReader(gzipReader)

	// 读取并解压 tar 文件中的内容
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// 已经读取完所有文件
			break
		}
		if err != nil {
			return xerrors.Errorf("%w", err)
		}

		if header.Name == "pax_global_header" {
			continue
		}

		header.Name = strings.TrimPrefix(header.Name, "gitbook-")
		fileOrPathName := filepath.Join(dest, header.Name)

		if header.FileInfo().IsDir() {
			err = os.Mkdir(fileOrPathName, 0755)
			if err != nil {
				return xerrors.Errorf("%w", err)
			}
			continue
		}

		// 创建目标文件
		targetFile, err := os.Create(fileOrPathName)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
		defer targetFile.Close()

		// 将文件内容从 tar 文件中拷贝到目标文件
		if _, err := io.Copy(targetFile, tarReader); err != nil {
			return xerrors.Errorf("%w", err)
		}

		fmt.Println("解压文件:", targetFile.Name())
	}

	fmt.Println("解压完成")

	return nil
}
