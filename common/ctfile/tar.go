package ctfile

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func GenTarGZ(tarFile string, files []string) error {
	// file write
	fileWriter, err := os.Create(tarFile)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	// gzip write
	gw := gzip.NewWriter(fileWriter)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// read file and write
	for _, fi := range files {
		fileState, err := os.Stat(fi)
		if err != nil {
			return err
		}

		// set info to header
		tarHeader := new(tar.Header)
		tarHeader.Name = fileState.Name()
		tarHeader.Size = fileState.Size()
		tarHeader.Mode = int64(fileState.Mode())
		tarHeader.ModTime = fileState.ModTime()

		// write header
		err = tw.WriteHeader(tarHeader)
		if err != nil {
			return err
		}

		file, err := os.Open(fi)
		if err != nil {
			return err
		}

		// write file
		_, err = io.Copy(tw, file)
		if err != nil {
			return err
		}
	}
	return nil
}
