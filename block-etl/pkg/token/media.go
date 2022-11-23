package token

import (
	"context"
	"strings"

	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/netutils"
)

const (
	// MediaTypeVideo represents a video
	MediaTypeVideo MediaType = "video"
	// MediaTypeImage represents an image
	MediaTypeImage MediaType = "image"
	// MediaTypeGIF represents a gif
	MediaTypeGIF MediaType = "gif"
	// MediaTypeSVG represents an SVG
	MediaTypeSVG MediaType = "svg"
	// MediaTypeBase64BMP represents a base64 encoded bmp file
	MediaTypeBase64BMP MediaType = "base64bmp"
	// MediaTypeText represents plain text
	MediaTypeText MediaType = "text"
	// MediaTypeHTML represents html
	MediaTypeHTML MediaType = "html"
	// MediaTypeBase64Text represents a base64 encoded plain text
	MediaTypeBase64Text MediaType = "base64text"
	// MediaTypeAudio represents audio
	MediaTypeAudio MediaType = "audio"
	// MediaTypeJSON represents json metadata
	MediaTypeJSON MediaType = "json"
	// MediaTypeAnimation represents an animation (.glb)
	MediaTypeAnimation MediaType = "animation"
	// MediaTypeInvalid represents an invalid media type such as when a token's external metadata's API is broken or no longer exists
	MediaTypeInvalid MediaType = "invalid"
	// MediaTypeUnknown represents an unknown media type
	MediaTypeUnknown MediaType = "unknown"
	// MediaTypeSyncing represents a syncing media
	MediaTypeSyncing MediaType = "syncing"
)

// MediaType represents the type of media that a token
type MediaType string

type mediaWithContentType struct {
	mediaType   MediaType
	contentType string
}

var postfixesToMediaTypes = map[string]mediaWithContentType{
	"jpg":  {MediaTypeImage, "image/jpeg"},
	"jpeg": {MediaTypeImage, "image/jpeg"},
	"png":  {MediaTypeImage, "image/png"},
	"webp": {MediaTypeImage, "image/webp"},
	"gif":  {MediaTypeGIF, "image/gif"},
	"mp4":  {MediaTypeVideo, "video/mp4"},
	"webm": {MediaTypeVideo, "video/webm"},
	"glb":  {MediaTypeAnimation, "model/gltf-binary"},
	"gltf": {MediaTypeAnimation, "model/gltf+json"},
	"svg":  {MediaTypeImage, "image/svg+xml"},
}

func predictTrueURLs(ctx context.Context, curImg, curV string) (string, string) {
	imgMediaType, _, _, err := PredictMediaType(ctx, curImg)
	if err != nil {
		return curImg, curV
	}
	vMediaType, _, _, err := PredictMediaType(ctx, curV)
	if err != nil {
		return curImg, curV
	}

	if IsAnimationLike(imgMediaType) && !IsAnimationLike(vMediaType) {
		return curV, curImg
	}

	if !IsValid(imgMediaType) || !IsValid(vMediaType) {
		return curImg, curV
	}

	if IsMorePriorityThan(imgMediaType, vMediaType) {
		return curV, curImg
	}

	return curImg, curV
}

// IsValid returns true if the media type is not unknown, syncing, or invalid
func IsValid(m MediaType) bool {
	return m != MediaTypeUnknown && m != MediaTypeInvalid && m != MediaTypeSyncing
}

// IsImageLike returns true if the media type is a type that is expected to be like an image and not live render
func IsImageLike(m MediaType) bool {
	return m == MediaTypeImage || m == MediaTypeGIF || m == MediaTypeBase64BMP || m == MediaTypeSVG
}

// IsAnimationLike returns true if the media type is a type that is expected to be like an animation and live render
func IsAnimationLike(m MediaType) bool {
	return m == MediaTypeVideo || m == MediaTypeHTML || m == MediaTypeAudio || m == MediaTypeAnimation
}

// IsMorePriorityThan returns true if the media type is more important than the other media type
func IsMorePriorityThan(m, other MediaType) bool {
	for _, t := range mediaTypePriorities {
		if t == m {
			return true
		}
		if t == other {
			return false
		}
	}
	return true
}

// PredictMediaType guesses the media type of the given URL.
func PredictMediaType(pCtx context.Context, url string) (MediaType, string, int64, error) {
	spl := strings.Split(url, ".")
	if len(spl) > 1 {
		ext := spl[len(spl)-1]
		ext = strings.Split(ext, "?")[0]
		if t, ok := postfixesToMediaTypes[ext]; ok {
			return t.mediaType, t.contentType, 0, nil
		}
	}
	asURI := url
	uriType := TokenURIType(asURI)
	switch uriType {
	case URITypeBase64JSON, URITypeJSON:
		return MediaTypeJSON, "application/json", int64(len(asURI)), nil
	case URITypeBase64SVG, URITypeSVG:
		return MediaTypeSVG, "image/svg", int64(len(asURI)), nil
	case URITypeBase64BMP:
		return MediaTypeBase64BMP, "image/bmp", int64(len(asURI)), nil
	case URITypeHTTP, URITypeIPFSAPI, URITypeIPFSGateway:
		contentType, contentLength, err := netutils.GetHTTPHeaders(pCtx, url)
		if err != nil {
			return MediaTypeUnknown, "", 0, err
		}
		return MediaFromContentType(contentType), contentType, contentLength, nil
	case URITypeIPFS:
		contentType, contentLength, err := netutils.GetIPFSHeaders(pCtx, strings.TrimPrefix(asURI, "ipfs://"))
		if err != nil {
			return MediaTypeUnknown, "", 0, err
		}
		return MediaFromContentType(contentType), contentType, contentLength, nil
	}
	return MediaTypeUnknown, "", 0, nil
}

func MediaFromContentType(contentType string) MediaType {
	contentType = strings.TrimSpace(contentType)
	whereCharset := strings.IndexByte(contentType, ';')
	if whereCharset != -1 {
		contentType = contentType[:whereCharset]
	}
	spl := strings.Split(contentType, "/")

	switch spl[0] {
	case "image":
		switch spl[1] {
		case "svg":
			return MediaTypeSVG
		case "gif":
			return MediaTypeGIF
		default:
			return MediaTypeImage
		}
	case "video":
		return MediaTypeVideo
	case "audio":
		return MediaTypeAudio
	case "text":
		switch spl[1] {
		case "html":
			return MediaTypeHTML
		default:
			return MediaTypeText
		}
	default:
		return MediaTypeUnknown
	}
}
