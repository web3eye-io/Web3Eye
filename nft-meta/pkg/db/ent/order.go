// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
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
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber uint64 `json:"block_number,omitempty"`
	// TxIndex holds the value of the "tx_index" field.
	TxIndex uint32 `json:"tx_index,omitempty"`
	// LogIndex holds the value of the "log_index" field.
	LogIndex uint32 `json:"log_index,omitempty"`
	// Recipient holds the value of the "recipient" field.
	Recipient string `json:"recipient,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldCreatedAt, order.FieldUpdatedAt, order.FieldDeletedAt, order.FieldBlockNumber, order.FieldTxIndex, order.FieldLogIndex:
			values[i] = new(sql.NullInt64)
		case order.FieldChainType, order.FieldChainID, order.FieldTxHash, order.FieldRecipient, order.FieldRemark:
			values[i] = new(sql.NullString)
		case order.FieldEntID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = uint32(value.Int64)
		case order.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				o.EntID = *value
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = uint32(value.Int64)
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = uint32(value.Int64)
			}
		case order.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = uint32(value.Int64)
			}
		case order.FieldChainType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_type", values[i])
			} else if value.Valid {
				o.ChainType = value.String
			}
		case order.FieldChainID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				o.ChainID = value.String
			}
		case order.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				o.TxHash = value.String
			}
		case order.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value.Valid {
				o.BlockNumber = uint64(value.Int64)
			}
		case order.FieldTxIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tx_index", values[i])
			} else if value.Valid {
				o.TxIndex = uint32(value.Int64)
			}
		case order.FieldLogIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field log_index", values[i])
			} else if value.Valid {
				o.LogIndex = uint32(value.Int64)
			}
		case order.FieldRecipient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient", values[i])
			} else if value.Valid {
				o.Recipient = value.String
			}
		case order.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				o.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", o.EntID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", o.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", o.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", o.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("chain_type=")
	builder.WriteString(o.ChainType)
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(o.ChainID)
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(o.TxHash)
	builder.WriteString(", ")
	builder.WriteString("block_number=")
	builder.WriteString(fmt.Sprintf("%v", o.BlockNumber))
	builder.WriteString(", ")
	builder.WriteString("tx_index=")
	builder.WriteString(fmt.Sprintf("%v", o.TxIndex))
	builder.WriteString(", ")
	builder.WriteString("log_index=")
	builder.WriteString(fmt.Sprintf("%v", o.LogIndex))
	builder.WriteString(", ")
	builder.WriteString("recipient=")
	builder.WriteString(o.Recipient)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(o.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
