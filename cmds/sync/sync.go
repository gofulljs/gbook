package sync

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

var Run = &cli.Command{
	Name:  "sync",
	Usage: "sync gitbook 3.2.3",
	Action: func(cctx *cli.Context) error {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil
		}
		rootPath := filepath.Join(homeDir, ".gitbook", "versions")
		if checkGitbook323IsExist(rootPath) {
			return nil
		}

		err = os.MkdirAll(rootPath, 0755)
		if err != nil {
			return nil
		}

		return nil
	},
}

// checkGitbook323IsExist check gitbook is in $HOME dir exist
func checkGitbook323IsExist(rootPath string) bool {
	bookJsFullPath := filepath.Join(rootPath, "bin", "gitbook.js")
	_, err := os.Stat(bookJsFullPath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// Download 3.2.3
func downloadGitbook323(rootPath string) error {
	http.Get("")
}

func tarGzUnzip(zipFile, dest string) error {
	fr, err := os.Open(zipFile)
	if err != nil {
		log.Error("err", err)
		return err
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Error("err", err)
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("err", err)
			return err
		}
		fw, err := os.OpenFile(dest+h.Name, os.O_CREATE|os.O_WRONLY, 0666 /*os.FileMode(h.Mode)*/)
		if err != nil {
			log.Error("err", err)
			return err
		}
		defer fw.Close()
		_, err = io.Copy(fw, tr)
		if err != nil {
			log.Error("err", err)
			return err
		}
	}
	log.Info("unzip tar.gz ok")
	return nil
}
