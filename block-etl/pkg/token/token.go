package token

import (
	"context"
	"fmt"

	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/indexer"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	tokenProto "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

type TokenURIInfo struct {
	URI         string
	URIType     URIType
	ImageURL    string
	VideoURL    string
	Name        string
	Description string
}

// TODO: support special nft project
func GetTokenURIInfo(ctx context.Context, uri string) (*TokenURIInfo, error) {
	into := &TokenMetadata{}

	err := DecodeMetadataFromURI(ctx, uri, into)
	if err != nil {
		return nil, err
	}

	tokenURIType := TokenURIType(uri)
	name, description := FindNameAndDescription(ctx, *into)
	iURL, vURL := FindImageAndAnimationURLs(ctx, *into, uri, AnimationKeywords, ImageKeywords, true)

	return &TokenURIInfo{
		URI:         uri,
		URIType:     tokenURIType,
		Name:        name,
		Description: description,
		ImageURL:    iURL,
		VideoURL:    vURL,
	}, nil
}

func CheckTokenReq(info *tokenProto.TokenReq) *tokenProto.TokenReq {
	if info == nil {
		return nil
	}

	info.URIState = basetype.TokenURIState_TokenURIFinish.Enum()
	info.VectorState = tokenProto.ConvertState_Waiting.Enum()
	if *info.URI == "" {
		info.URIState = basetype.TokenURIState_TokenURIError.Enum()
		info.VectorState = tokenProto.ConvertState_Failed.Enum()
	}
	if *info.ImageURL == "" || *info.Name == "" || *info.Description == "" {
		info.URIState = basetype.TokenURIState_TokenURIIncomplete.Enum()
	}
	if *info.ImageURL == "" {
		info.VectorState = tokenProto.ConvertState_Failed.Enum()
	}

	if len(*info.URI) > indexer.MaxContentLength {
		remark := fmt.Sprintf("%v,uri too long(length: %v),skip to store it", info.Remark, len(*info.URI))
		uri := *info.URI
		uri = uri[:indexer.OverLimitStoreLength]
		info.URI = &uri
		info.Remark = &remark
	}

	if len(*info.ImageURL) > indexer.MaxContentLength {
		remark := fmt.Sprintf("%v,imageURL too long(length: %v),skip to store it", info.Remark, len(*info.URI))
		imageURL := *info.ImageURL
		imageURL = imageURL[:indexer.OverLimitStoreLength]
		info.ImageURL = &imageURL
		info.Remark = &remark
	}

	if len(*info.Description) > indexer.MaxContentLength {
		remark := fmt.Sprintf("%v,description too long(length: %v),skip to store it", info.Remark, len(*info.URI))
		description := *info.Description
		description = description[:indexer.OverLimitStoreLength]
		info.Description = &description
		info.Remark = &remark
	}
	return info
}
