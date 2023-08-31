// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderpair"
)

// OrderPair is the model entity for the OrderPair schema.
type OrderPair struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// Recipient holds the value of the "recipient" field.
	Recipient string `json:"recipient,omitempty"`
	// TargetID holds the value of the "target_id" field.
	TargetID string `json:"target_id,omitempty"`
	// BarterID holds the value of the "barter_id" field.
	BarterID string `json:"barter_id,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderPair) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderpair.FieldCreatedAt, orderpair.FieldUpdatedAt, orderpair.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case orderpair.FieldTxHash, orderpair.FieldRecipient, orderpair.FieldTargetID, orderpair.FieldBarterID, orderpair.FieldRemark:
			values[i] = new(sql.NullString)
		case orderpair.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderPair", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderPair fields.
func (op *OrderPair) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderpair.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				op.ID = *value
			}
		case orderpair.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = uint32(value.Int64)
			}
		case orderpair.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = uint32(value.Int64)
			}
		case orderpair.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				op.DeletedAt = uint32(value.Int64)
			}
		case orderpair.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				op.TxHash = value.String
			}
		case orderpair.FieldRecipient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recipient", values[i])
			} else if value.Valid {
				op.Recipient = value.String
			}
		case orderpair.FieldTargetID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field target_id", values[i])
			} else if value.Valid {
				op.TargetID = value.String
			}
		case orderpair.FieldBarterID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field barter_id", values[i])
			} else if value.Valid {
				op.BarterID = value.String
			}
		case orderpair.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				op.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OrderPair.
// Note that you need to call OrderPair.Unwrap() before calling this method if this OrderPair
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrderPair) Update() *OrderPairUpdateOne {
	return (&OrderPairClient{config: op.config}).UpdateOne(op)
}

// Unwrap unwraps the OrderPair entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrderPair) Unwrap() *OrderPair {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderPair is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrderPair) String() string {
	var builder strings.Builder
	builder.WriteString("OrderPair(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", op.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", op.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", op.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(op.TxHash)
	builder.WriteString(", ")
	builder.WriteString("recipient=")
	builder.WriteString(op.Recipient)
	builder.WriteString(", ")
	builder.WriteString("target_id=")
	builder.WriteString(op.TargetID)
	builder.WriteString(", ")
	builder.WriteString("barter_id=")
	builder.WriteString(op.BarterID)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(op.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// OrderPairs is a parsable slice of OrderPair.
type OrderPairs []*OrderPair

func (op OrderPairs) config(cfg config) {
	for _i := range op {
		op[_i].config = cfg
	}
}