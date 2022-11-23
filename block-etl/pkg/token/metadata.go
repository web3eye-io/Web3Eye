package token

import (
	"context"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/web3eye-io/cyber-tracer/block-etl/pkg/netutils"
)

const (
	// URITypeIPFS represents an IPFS URI
	URITypeIPFS URIType = "ipfs"
	// URITypeArweave represents an Arweave URI
	URITypeArweave URIType = "arweave"
	// URITypeHTTP represents an HTTP URI
	URITypeHTTP URIType = "http"
	// URITypeIPFSAPI represents an IPFS API URI
	URITypeIPFSAPI URIType = "ipfs-api"
	// URITypeIPFSGateway represents an IPFS Gateway URI
	URITypeIPFSGateway URIType = "ipfs-gateway"
	// URITypeBase64JSON represents a base64 encoded JSON document
	URITypeBase64JSON URIType = "base64json"
	// URITypeJSON represents a JSON document
	URITypeJSON URIType = "json"
	// URITypeBase64SVG represents a base64 encoded SVG
	URITypeBase64SVG URIType = "base64svg"
	//URITypeBase64BMP represents a base64 encoded BMP
	URITypeBase64BMP URIType = "base64bmp"
	// URITypeSVG represents an SVG
	URITypeSVG URIType = "svg"
	// URITypeENS represents an ENS domain
	URITypeENS URIType = "ens"
	// URITypeUnknown represents an unknown URI type
	URITypeUnknown URIType = "unknown"
	// URITypeInvalid represents an invalid URI type
	URITypeInvalid URIType = "invalid"
	// URITypeNone represents no URI
	URITypeNone URIType = "none"
)

var mediaTypePriorities = []MediaType{MediaTypeHTML, MediaTypeAudio, MediaTypeAnimation, MediaTypeVideo, MediaTypeBase64BMP, MediaTypeGIF, MediaTypeSVG, MediaTypeImage, MediaTypeJSON, MediaTypeBase64Text, MediaTypeText, MediaTypeSyncing, MediaTypeUnknown, MediaTypeInvalid}

var (
	AnimationKeywords  = []string{"animation", "video"}
	ImageKeywords      = []string{"image"}
	DefaultSearchDepth = 5
)

type URIType string

// TokenMetadata represents the JSON metadata for a token
type TokenMetadata map[string]interface{}

// Scan implements the database/sql Scanner interface for the TokenMetadata type
func (m *TokenMetadata) Scan(src interface{}) error {
	if src == nil {
		*m = TokenMetadata{}
		return nil
	}
	return json.Unmarshal(src.([]uint8), m)
}

// Value implements the database/sql/driver Valuer interface for the TokenMetadata type
func (m TokenMetadata) Value() (driver.Value, error) {
	return m.MarshallJSON()
}

// MarshallJSON implements the json.Marshaller interface for the TokenMetadata type
func (m TokenMetadata) MarshallJSON() ([]byte, error) {
	val, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	cleaned := strings.ToValidUTF8(string(val), "")
	// Replace literal '\\u0000' with empty string (marshal to JSON escapes each backslash)
	cleaned = strings.ReplaceAll(cleaned, "\\\\u0000", "")
	// Replace unicode NULL char (u+0000) i.e. '\u0000' with empty string
	cleaned = strings.ReplaceAll(cleaned, "\\u0000", "")
	return []byte(cleaned), nil
}

func TokenURIType(uri string) URIType {
	asString := uri
	asString = strings.TrimSpace(uri)
	switch {
	case strings.HasPrefix(asString, "ipfs"), strings.HasPrefix(asString, "Qm"):
		return URITypeIPFS
	case strings.HasPrefix(asString, "ar://"), strings.HasPrefix(asString, "arweave://"):
		return URITypeArweave
	case strings.HasPrefix(asString, "data:application/json;base64,"):
		return URITypeBase64JSON
	case strings.HasPrefix(asString, "data:image/svg+xml;base64,"), strings.HasPrefix(asString, "data:image/svg xml;base64,"):
		return URITypeBase64SVG
	case strings.HasPrefix(asString, "data:image/bmp;base64,"):
		return URITypeBase64BMP
	case strings.Contains(asString, "ipfs.io/api"):
		return URITypeIPFSAPI
	case strings.Contains(asString, "/ipfs/"):
		return URITypeIPFSGateway
	case strings.HasPrefix(asString, "http"), strings.HasPrefix(asString, "https"):
		return URITypeHTTP
	case strings.HasPrefix(asString, "{"), strings.HasPrefix(asString, "["), strings.HasPrefix(asString, "data:application/json"), strings.HasPrefix(asString, "data:text/plain,{"):
		return URITypeJSON
	case strings.HasPrefix(asString, "<svg"), strings.HasPrefix(asString, "data:image/svg+xml;utf8,"), strings.HasPrefix(asString, "data:image/svg+xml,"), strings.HasPrefix(asString, "data:image/svg xml,"):
		return URITypeSVG
	case strings.HasSuffix(asString, ".ens"):
		return URITypeENS
	case asString == string(URITypeInvalid):
		return URITypeInvalid
	case asString == "":
		return URITypeNone
	default:
		return URITypeUnknown
	}
}

func FindNameAndDescription(ctx context.Context, metadata TokenMetadata) (string, string) {
	name, ok := GetValueFromMapUnsafe(metadata, "name", DefaultSearchDepth).(string)
	if !ok {
		name = ""
	}

	description, ok := GetValueFromMapUnsafe(metadata, "description", DefaultSearchDepth).(string)
	if !ok {
		description = ""
	}

	return name, description
}

func FindImageAndAnimationURLs(ctx context.Context, metadata TokenMetadata, turi string, animationKeywords, imageKeywords []string, predict bool) (imgURL string, vURL string) {
	into := &TokenMetadata{}
	DecodeMetadataFromURI(context.Background(), turi, into)
	for _, keyword := range animationKeywords {
		if it, ok := GetValueFromMapUnsafe(metadata, keyword, DefaultSearchDepth).(string); ok && it != "" {
			vURL = it
			break
		}
	}

	for _, keyword := range imageKeywords {
		if it, ok := GetValueFromMapUnsafe(metadata, keyword, DefaultSearchDepth).(string); ok && it != "" && it != vURL {
			imgURL = it
			break
		}
	}

	if imgURL == "" && vURL == "" {
		imgURL = turi
	}

	if predict {
		return predictTrueURLs(ctx, imgURL, vURL)
	}
	return imgURL, vURL
}

// GetValueFromMap is a function that returns the value at the first occurence of a given key in a map that potentially contains nested maps
func GetValueFromMap(m map[string]interface{}, key string, searchDepth int) interface{} {
	if searchDepth == 0 {
		return nil
	}
	if _, ok := m[key]; ok {
		return m[key]
	}
	for k, v := range m {
		if strings.EqualFold(k, key) {
			return v
		}

		if nest, ok := v.(map[string]interface{}); ok {
			if nestVal := GetValueFromMap(nest, key, searchDepth-1); nestVal != nil {
				return nestVal
			}
		}
		if array, ok := v.([]interface{}); ok {
			for _, arrayVal := range array {
				if nest, ok := arrayVal.(map[string]interface{}); ok {
					if nestVal := GetValueFromMap(nest, key, searchDepth-1); nestVal != nil {
						return nestVal
					}
				}
			}
		}
	}
	return nil
}

// GetValueFromMapUnsafe is a function that returns the value at the first occurence of a given key in a map that potentially contains nested maps
// This function is unsafe because it will also return if the specified key is a substring of any key found in the map
func GetValueFromMapUnsafe(m map[string]interface{}, key string, searchDepth int) interface{} {
	if searchDepth == 0 {
		return nil
	}
	if _, ok := m[key]; ok {
		return m[key]
	}
	for k, v := range m {

		if strings.Contains(strings.ToLower(k), strings.ToLower(key)) {
			return v
		}

		if nest, ok := v.(map[string]interface{}); ok {
			if nestVal := GetValueFromMap(nest, key, searchDepth-1); nestVal != nil {
				return nestVal
			}
		}
		if array, ok := v.([]interface{}); ok {
			for _, arrayVal := range array {
				if nest, ok := arrayVal.(map[string]interface{}); ok {
					if nestVal := GetValueFromMap(nest, key, searchDepth-1); nestVal != nil {
						return nestVal
					}
				}
			}
		}
	}
	return nil
}

// GetMetadataFromURI parses and returns the NFT metadata for a given token URI
func GetMetadataFromURI(ctx context.Context, turi string) (TokenMetadata, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()
	var meta TokenMetadata
	err := DecodeMetadataFromURI(ctx, turi, &meta)
	if err != nil {
		return nil, err
	}

	return meta, nil

}

// DecodeMetadataFromURI calls URI and decodes the data into a metadata map
func DecodeMetadataFromURI(ctx context.Context, turi string, into *TokenMetadata) error {

	asString := turi

	switch TokenURIType(turi) {
	case URITypeBase64JSON:
		// decode the base64 encoded json
		b64data := asString[strings.IndexByte(asString, ',')+1:]
		decoded, err := base64.StdEncoding.DecodeString(string(b64data))
		if err != nil {
			return fmt.Errorf("error decoding base64 data: %s \n\n%s", err, b64data)
		}

		return json.Unmarshal(RemoveBOM(decoded), into)
	case URITypeBase64SVG:
		b64data := asString[strings.IndexByte(asString, ',')+1:]
		decoded, err := base64.StdEncoding.DecodeString(string(b64data))
		if err != nil {
			return fmt.Errorf("error decoding base64 data: %s \n\n%s", err, b64data)
		}
		into = &TokenMetadata{"image": string(decoded)}
		return nil
	case URITypeIPFS, URITypeIPFSGateway:
		bs, err := netutils.GetIPFSData(ctx, netutils.GetURIPath(asString, false))
		if err != nil {
			return err
		}
		return json.Unmarshal(bs, into)
	case URITypeArweave:
		path := strings.ReplaceAll(asString, "arweave://", "")
		path = strings.ReplaceAll(path, "ar://", "")
		// TODO: support ipfs node
		// result, err := GetArweaveData(arweaveClient, path)
		// if err != nil {
		// }
		result, err := netutils.GetArweaveDataHTTP(ctx, path)
		if err != nil {
			return err
		}
		return json.Unmarshal(result, into)
	case URITypeHTTP:
		req, err := http.NewRequestWithContext(ctx, "GET", asString, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %s", err)
		}
		resp, err := netutils.DefaultHTTPClient.Do(req)
		if err != nil {
			return fmt.Errorf("error getting data from http: %s", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode > 399 || resp.StatusCode < 200 {
			return netutils.ErrHTTP{Status: resp.StatusCode, URL: asString}
		}
		return json.NewDecoder(resp.Body).Decode(into)
	// TODO: I don`t what it means
	// case URITypeIPFSAPI:
	// 	parsedURL, err := url.Parse(asString)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	query := parsedURL.Query().Get("arg")
	// 	it, err := ipfsClient.Cat(query)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer it.Close()
	// 	return json.NewDecoder(it).Decode(into)
	case URITypeJSON, URITypeSVG:
		idx := strings.IndexByte(asString, '{')
		if idx == -1 {
			return json.Unmarshal(RemoveBOM([]byte(asString)), into)
		}
		return json.Unmarshal(RemoveBOM([]byte(asString[idx:])), into)

	default:
		return fmt.Errorf("unknown token URI type: %s", TokenURIType(turi))
	}

}

// RemoveBOM removes the byte order mark from a byte array
func RemoveBOM(bs []byte) []byte {
	if len(bs) > 3 && bs[0] == 0xEF && bs[1] == 0xBB && bs[2] == 0xBF {
		return bs[3:]
	}
	return bs
}
