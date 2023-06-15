package car

import (
	"bufio"
	"context"
	"io"
	"os"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/filecoin-project/go-commp-utils/writer"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"golang.org/x/xerrors"
)

type CommPRet struct {
	Root        cid.Cid
	Size        uint64
	PayloadSize int64
	PieceSize   uint64
}

func ClientCalcCommP(ctx context.Context, inpath string) (*CommPRet, error) {
	rdr, err := os.Open(inpath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rdr.Close()
		logger.Sugar().Error(err)
	}()

	// check that the data is a car file; if it's not, retrieval won't work
	_, err = car.ReadHeader(bufio.NewReader(rdr))
	if err != nil {
		return nil, xerrors.Errorf("not a car file: %w", err)
	}

	if _, err := rdr.Seek(0, io.SeekStart); err != nil {
		return nil, xerrors.Errorf("seek to start: %w", err)
	}

	w := &writer.Writer{}
	_, err = io.CopyBuffer(w, rdr, make([]byte, writer.CommPBuf))
	if err != nil {
		return nil, xerrors.Errorf("copy into commp writer: %w", err)
	}

	commp, err := w.Sum()
	if err != nil {
		return nil, xerrors.Errorf("computing commP failed: %w", err)
	}

	return &CommPRet{
		Root:        commp.PieceCID,
		Size:        uint64(commp.PieceSize.Unpadded()),
		PayloadSize: commp.PayloadSize,
		PieceSize:   uint64(commp.PieceSize),
	}, nil
}
