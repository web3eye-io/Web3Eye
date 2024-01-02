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
func GetTokenURIInfo(ctx context.Context, uri string) (*TokenURIInfo, bool, error) {
	into := &TokenMetadata{}

	err := DecodeMetadataFromURI(ctx, uri, into)
	if err != nil {
		return nil, false, err
	}

	tokenURIType := TokenURIType(uri)
	name, description := FindNameAndDescription(ctx, *into)
	iURL, vURL := FindImageAndAnimationURLs(ctx, *into, uri, AnimationKeywords, ImageKeywords, true)

	complete := true
	if name == "" ||
		description == "" ||
		iURL == "" {
		complete = false
	}
	return &TokenURIInfo{
		URI:         uri,
		URIType:     tokenURIType,
		Name:        name,
		Description: description,
		ImageURL:    iURL,
		VideoURL:    vURL,
	}, complete, nil
}
