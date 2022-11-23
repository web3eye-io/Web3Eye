package token

import (
	"context"
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
// TODO: support context
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
