// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
)

// Token is the model entity for the Token schema.
type Token struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// ChainType holds the value of the "chain_type" field.
	ChainType string `json:"chain_type,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID string `json:"chain_id,omitempty"`
	// Contract holds the value of the "contract" field.
	Contract string `json:"contract,omitempty"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type,omitempty"`
	// TokenID holds the value of the "token_id" field.
	TokenID string `json:"token_id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// URI holds the value of the "uri" field.
	URI string `json:"uri,omitempty"`
	// URIType holds the value of the "uri_type" field.
	URIType string `json:"uri_type,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// VideoURL holds the value of the "video_url" field.
	VideoURL string `json:"video_url,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// VectorID holds the value of the "vector_id" field.
	VectorID int64 `json:"vector_id,omitempty"`
	// VectorState holds the value of the "vector_state" field.
	VectorState string `json:"vector_state,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// IpfsImageURL holds the value of the "ipfs_image_url" field.
	IpfsImageURL string `json:"ipfs_image_url,omitempty"`
	// FileCid holds the value of the "file_cid" field.
	FileCid string `json:"file_cid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Token) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case token.FieldCreatedAt, token.FieldUpdatedAt, token.FieldDeletedAt, token.FieldVectorID:
			values[i] = new(sql.NullInt64)
		case token.FieldChainType, token.FieldChainID, token.FieldContract, token.FieldTokenType, token.FieldTokenID, token.FieldOwner, token.FieldURI, token.FieldURIType, token.FieldImageURL, token.FieldVideoURL, token.FieldDescription, token.FieldName, token.FieldVectorState, token.FieldRemark, token.FieldIpfsImageURL, token.FieldFileCid:
			values[i] = new(sql.NullString)
		case token.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Token", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Token fields.
func (t *Token) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case token.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case token.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = uint32(value.Int64)
			}
		case token.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = uint32(value.Int64)
			}
		case token.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = uint32(value.Int64)
			}
		case token.FieldChainType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_type", values[i])
			} else if value.Valid {
				t.ChainType = value.String
			}
		case token.FieldChainID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				t.ChainID = value.String
			}
		case token.FieldContract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract", values[i])
			} else if value.Valid {
				t.Contract = value.String
			}
		case token.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				t.TokenType = value.String
			}
		case token.FieldTokenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_id", values[i])
			} else if value.Valid {
				t.TokenID = value.String
			}
		case token.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				t.Owner = value.String
			}
		case token.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				t.URI = value.String
			}
		case token.FieldURIType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri_type", values[i])
			} else if value.Valid {
				t.URIType = value.String
			}
		case token.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				t.ImageURL = value.String
			}
		case token.FieldVideoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_url", values[i])
			} else if value.Valid {
				t.VideoURL = value.String
			}
		case token.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case token.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case token.FieldVectorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vector_id", values[i])
			} else if value.Valid {
				t.VectorID = value.Int64
			}
		case token.FieldVectorState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vector_state", values[i])
			} else if value.Valid {
				t.VectorState = value.String
			}
		case token.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				t.Remark = value.String
			}
		case token.FieldIpfsImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ipfs_image_url", values[i])
			} else if value.Valid {
				t.IpfsImageURL = value.String
			}
		case token.FieldFileCid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_cid", values[i])
			} else if value.Valid {
				t.FileCid = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Token.
// Note that you need to call Token.Unwrap() before calling this method if this Token
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Token) Update() *TokenUpdateOne {
	return (&TokenClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Token entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Token) Unwrap() *Token {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Token is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Token) String() string {
	var builder strings.Builder
	builder.WriteString("Token(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", t.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("chain_type=")
	builder.WriteString(t.ChainType)
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(t.ChainID)
	builder.WriteString(", ")
	builder.WriteString("contract=")
	builder.WriteString(t.Contract)
	builder.WriteString(", ")
	builder.WriteString("token_type=")
	builder.WriteString(t.TokenType)
	builder.WriteString(", ")
	builder.WriteString("token_id=")
	builder.WriteString(t.TokenID)
	builder.WriteString(", ")
	builder.WriteString("owner=")
	builder.WriteString(t.Owner)
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(t.URI)
	builder.WriteString(", ")
	builder.WriteString("uri_type=")
	builder.WriteString(t.URIType)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(t.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("video_url=")
	builder.WriteString(t.VideoURL)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("vector_id=")
	builder.WriteString(fmt.Sprintf("%v", t.VectorID))
	builder.WriteString(", ")
	builder.WriteString("vector_state=")
	builder.WriteString(t.VectorState)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(t.Remark)
	builder.WriteString(", ")
	builder.WriteString("ipfs_image_url=")
	builder.WriteString(t.IpfsImageURL)
	builder.WriteString(", ")
	builder.WriteString("file_cid=")
	builder.WriteString(t.FileCid)
	builder.WriteByte(')')
	return builder.String()
}

// Tokens is a parsable slice of Token.
type Tokens []*Token

func (t Tokens) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
