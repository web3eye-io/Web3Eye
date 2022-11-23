//nolint
package main

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool"
	proto "github.com/NpoolPlatform/message/npool/nftmeta/v1/token"
	crud "github.com/web3eye-io/cyber-tracer/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/milvusdb"
)

func api() {
	fmt.Println(milvusdb.Init(context.Background()))

	mgr := milvusdb.NewNFTConllectionMGR()

	ret, err := mgr.Query(context.Background(), []int64{437431862635267233, 437431862635266693})
	fmt.Println("err")
	fmt.Println(err)
	vec1, vec2 := ret[437431862635267233], ret[437431862635266693]

	// vecs := [][milvusdb.VectorDim]float32{}
	// for i := 0; i < 100; i++ {
	// 	vec := [milvusdb.VectorDim]float32{}
	// 	for j := range vec {
	// 		vec[j] = rand.Float32()
	// 	}
	// 	vecs = append(vecs, vec)
	// }
	// fmt.Println(mgr.Create(context.Background(), vecs))
	// fmt.Println(mgr.Delete(context.Background(), 436519957287423533))
	// fmt.Println(mgr.Query(context.Background(), []int64{436519957287424159}))
	// fmt.Println(mgr.Query(context.Background(), []int64{436519957287424217}))
	fmt.Println(mgr.Search(context.Background(), [][milvusdb.VectorDim]float32{vec1, vec2}, 3))
}

func img2vector() {
	fmt.Println(imageconvert.ImgURLConvertVector("ipfs://QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE"))
}
func main() {
	// api()
	// img2vector()
	rescanImg()
}

func rescanImg() {
	conds := &proto.Conds{
		VectorState: &npool.StringVal{
			Op:    "eq",
			Value: proto.ConvertState_Failed.String(),
		},
	}
	_, total, err := crud.Rows(context.Background(), conds, 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	pageNum := 10
	pages := total/10 + 1
	for i := 0; i < pages; i++ {
		fmt.Println(i)
		rows, _, err := crud.Rows(context.Background(), conds, i*pageNum, pageNum)
		if err != nil {
			fmt.Println(err)
		}

		for _, info := range rows {
			imageconvert.DealVectorState(context.Background(), info.ID)
		}
	}

}
