package rich

import (
	"archive/zip"
	"io"
	"os"
	"path"
)

// Zip 压缩目录或文件
func Zip(srcDirPath string, destFilePath string) error {
	fw, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer fw.Close()

	tw := zip.NewWriter(fw)
	defer tw.Close()

	f, err := os.Open(srcDirPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.IsDir() {
		if err := zipDir(srcDirPath, path.Base(srcDirPath), tw); err != nil {
			return err
		}
	} else {
		if err := zipFile(srcDirPath, fi.Name(), tw); err != nil {
			return err
		}
	}
	return nil
}

func zipDir(srcDirPath string, recPath string, tw *zip.Writer) error {
	// Open source diretory
	dir, err := os.Open(srcDirPath)
	if err != nil {
		return err
	}
	defer dir.Close()
	// Get file info slice
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		// Append path
		curPath := srcDirPath + "/" + fi.Name()
		if fi.IsDir() {
			if err := zipDir(curPath, recPath+"/"+fi.Name(), tw); err != nil {
				return err
			}
		} else {
			if err := zipFile(curPath, recPath+"/"+fi.Name(), tw); err != nil {
				return err
			}
		}
	}
	return nil
}

// Deal with files
func zipFile(srcFile string, recPath string, tw *zip.Writer) error {
	// File reader
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()
	// Write hander
	w, err2 := tw.Create(recPath)
	if err2 != nil {
		return err2
	}
	// Write file data
	_, err = io.Copy(w, fr)
	if err != nil {
		return err
	}
	return nil
}
