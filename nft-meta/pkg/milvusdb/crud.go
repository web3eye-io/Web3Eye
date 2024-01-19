package milvusdb

// TODO: support batch vectors

import (
	"context"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

type nftConllectionMGR struct {
	*entity.Schema
}

var c *nftConllectionMGR

func NewNFTConllectionMGR() *nftConllectionMGR {
	if c != nil {
		return c
	}
	c = &nftConllectionMGR{NFTSchema}
	return c
}

func (c *nftConllectionMGR) Create(ctx context.Context, nftVectors [][VectorDim]float32) ([]int64, error) {
	emptyRet := []int64{}
	cli, err := Client(ctx)
	if err != nil {
		return emptyRet, err
	}

	_nftVectors := make([][]float32, len(nftVectors))
	for i := range _nftVectors {
		_nftVectors[i] = make([]float32, len(nftVectors[0]))
		copy(_nftVectors[i], nftVectors[i][:])
	}

	vector := entity.NewColumnFloatVector(FieldsVector, VectorDim, _nftVectors)
	iDs, err := cli.Insert(ctx, c.CollectionName, DefaultPartition, vector)
	if err != nil {
		return emptyRet, err
	}

	if err := cli.Flush(ctx, c.CollectionName, false); err != nil {
		return emptyRet, err
	}

	idRows := iDs.(*entity.ColumnInt64)
	ret := idRows.Data()

	return ret, nil
}

func (c *nftConllectionMGR) Delete(ctx context.Context, iDs []int64) error {
	cli, err := Client(ctx)
	if err != nil {
		return err
	}

	pks := entity.NewColumnInt64(FieldsID, iDs)
	return cli.DeleteByPks(ctx, c.CollectionName, DefaultPartition, pks)
}

// return map[id]vector
func (c *nftConllectionMGR) Query(ctx context.Context, iDs []int64) (map[int64][VectorDim]float32, error) {
	ret := make(map[int64][VectorDim]float32)
	cli, err := Client(ctx)
	if err != nil {
		return ret, err
	}

	pks := entity.NewColumnInt64(FieldsID, iDs)
	rets, err := cli.QueryByPks(ctx, c.CollectionName, []string{DefaultPartition}, pks, []string{FieldsVector})
	if err != nil {
		return ret, err
	}

	var idColumn *entity.ColumnInt64
	var vecColumn *entity.ColumnFloatVector
	for _, col := range rets {
		if col.Name() == FieldsID {
			idColumn = col.(*entity.ColumnInt64)
		}
		if col.Name() == FieldsVector {
			vecColumn = col.(*entity.ColumnFloatVector)
		}
	}
	for i := 0; i < rets[0].Len(); i++ {
		val, err := idColumn.ValueByIdx(i)
		if err != nil {
			continue
		}
		_vec := vecColumn.Data()[i]
		vec := [VectorDim]float32{}
		copy(vec[:], _vec)
		ret[val] = vec
	}

	return ret, nil
}

// return []map[id]score
func (c *nftConllectionMGR) Search(ctx context.Context, nftVectors [][VectorDim]float32, limit int) ([]map[int64]float32, error) {
	ret := make([]map[int64]float32, 0)

	cli, err := Client(ctx)
	if err != nil {
		return ret, err
	}

	DISKANN_list := 100
	sParam, err := entity.NewIndexDISKANNSearchParam(DISKANN_list)
	if err != nil {
		return ret, err
	}

	vec := []entity.Vector{}
	for i := range nftVectors {
		vec = append(vec, entity.FloatVector(nftVectors[i][:]))
	}

	sRet, err := cli.Search(ctx, c.CollectionName, []string{DefaultPartition}, "", []string{FieldsID}, vec,
		FieldsVector, entity.L2, limit, sParam)
	if err != nil {
		return ret, err
	}

	for retI, sColumn := range sRet {
		ret = append(ret, make(map[int64]float32, sColumn.IDs.Len()))
		iDs := sColumn.IDs.(*entity.ColumnInt64).Data()

		for i := 0; i < sColumn.ResultCount; i++ {
			ret[retI][iDs[i]] = sColumn.Scores[i]
		}
	}
	return ret, nil
}

func (c *nftConllectionMGR) TotalNum(ctx context.Context) (int64, error) {
	cli, err := Client(ctx)
	if err != nil {
		return 0, err
	}

	segs, err := cli.GetPersistentSegmentInfo(ctx, c.CollectionName)
	if err != nil {
		return 0, err
	}

	if len(segs) == 0 {
		return 0, nil
	}

	num := segs[0].NumRows
	for _, v := range segs {
		if v.NumRows < num {
			num = v.NumRows
		}
	}

	return num, nil
}
