package zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Zip(name string) ([]byte, error) {
	var buf = bytes.NewBuffer(nil)
	w := zip.NewWriter(buf)

	abs, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	info, _ := os.Stat(name)
	if info == nil {
		return nil, errors.New("not found")
	}
	zipFile := func(namepath string, path string) error {
		zw, err := w.Create(namepath)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zw, file)
		if err != nil {
			return err
		}
		err = w.Flush()
		return err
	}
	if info.IsDir() {
		err = filepath.Walk(abs, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(abs, path)
			if err != nil {
				return err
			}
			err = zipFile(rel, path)
			if err != nil {
				return err
			}
			return nil
		})
	} else {
		fp := filepath.Base(name)
		err = zipFile(fp, name)
	}
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnZip 解压
func UnZip(zipFile string, dstDir string) error {
	f, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer f.Close()

	os.MkdirAll(dstDir, 0777)

	for _, entry := range f.File {
		fp := filepath.Join(dstDir, entry.Name)

		if entry.FileInfo().IsDir() {
			os.MkdirAll(fp, entry.Mode())
			continue
		}

		file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, entry.Mode())
		if err != nil {
			return err
		}

		rc, err := entry.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(file, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
