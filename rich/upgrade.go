package rich

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

// LastURL 最后一个下载地址
func LastURL(owner, project, tagName string) (string, string, error) {
	repoService := github.NewClient(nil).Repositories
	ctx := context.Background()
	rel, _, err := repoService.GetLatestRelease(ctx, owner, project)
	if err != nil {
		return "", "", err
	}
	if tagName == *rel.TagName {
		return "", tagName, errors.New("版本重复:" + tagName)
	}
	os := runtime.GOOS
	arac := runtime.GOARCH
	for _, a := range rel.Assets {
		url := a.GetBrowserDownloadURL()
		if strings.Contains(url, os) && strings.Contains(url, arac) {
			return url, *rel.TagName, nil
		}
	}
	return "", *rel.TagName, errors.New("缺少下载资源 操作系统:" + os + " 架构:" + arac)
}

// Download 下载并解压缩
func Download(url, target string) error {
	client := http.Client{Timeout: 900 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reader, err := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}
