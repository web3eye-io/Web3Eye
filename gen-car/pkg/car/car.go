package car

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/common/ctfile"
	"github.com/web3eye-io/Web3Eye/common/unixfs"
)

const (
	DefaultCarVersion = 2
)

// CreateCar creates a car
// method from "github.com/ipld/go-car/cmd/car"
// the method is overwrite,because the original method is in main package
func CreateCar(ctx context.Context, carFilePath string, filesPath []string, version int) (rootCID string, err error) {
	dir, _ := filepath.Split(carFilePath)
	tarFilePath := fmt.Sprintf("%v/%v.tar.gz", dir, uuid.NewString())
	err = ctfile.GenTarGZ(tarFilePath, filesPath)
	if err != nil {
		return "", err
	}

	root, err := unixfs.CreateFilestore(context.Background(), tarFilePath, carFilePath)

	err = os.Remove(tarFilePath)
	if err != nil {
		return "", err
	}

	return root.String(), err
}

func Commp() {

}
