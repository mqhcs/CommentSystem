// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/replyarea"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ReplyArea is the model entity for the ReplyArea schema.
type ReplyArea struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Oid holds the value of the "oid" field.
	Oid int64 `json:"oid,omitempty"`
	// Type holds the value of the "type" field.
	Type int8 `json:"type,omitempty"`
	// Mid holds the value of the "mid" field.
	Mid int64 `json:"mid,omitempty"`
	// Count holds the value of the "count" field.
	Count int32 `json:"count,omitempty"`
	// RootCount holds the value of the "root_count" field.
	RootCount int32 `json:"root_count,omitempty"`
	// Acount holds the value of the "acount" field.
	Acount int32 `json:"acount,omitempty"`
	// State holds the value of the "state" field.
	State int8 `json:"state,omitempty"`
	// Ctime holds the value of the "ctime" field.
	Ctime time.Time `json:"ctime,omitempty"`
	// Mtime holds the value of the "mtime" field.
	Mtime time.Time `json:"mtime,omitempty"`
	// Attrs holds the value of the "attrs" field.
	Attrs int32 `json:"attrs,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta         string `json:"meta,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReplyArea) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case replyarea.FieldID, replyarea.FieldOid, replyarea.FieldType, replyarea.FieldMid, replyarea.FieldCount, replyarea.FieldRootCount, replyarea.FieldAcount, replyarea.FieldState, replyarea.FieldAttrs:
			values[i] = new(sql.NullInt64)
		case replyarea.FieldMeta:
			values[i] = new(sql.NullString)
		case replyarea.FieldCtime, replyarea.FieldMtime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReplyArea fields.
func (ra *ReplyArea) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case replyarea.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ra.ID = int64(value.Int64)
		case replyarea.FieldOid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				ra.Oid = value.Int64
			}
		case replyarea.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ra.Type = int8(value.Int64)
			}
		case replyarea.FieldMid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mid", values[i])
			} else if value.Valid {
				ra.Mid = value.Int64
			}
		case replyarea.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				ra.Count = int32(value.Int64)
			}
		case replyarea.FieldRootCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field root_count", values[i])
			} else if value.Valid {
				ra.RootCount = int32(value.Int64)
			}
		case replyarea.FieldAcount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field acount", values[i])
			} else if value.Valid {
				ra.Acount = int32(value.Int64)
			}
		case replyarea.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ra.State = int8(value.Int64)
			}
		case replyarea.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				ra.Ctime = value.Time
			}
		case replyarea.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				ra.Mtime = value.Time
			}
		case replyarea.FieldAttrs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attrs", values[i])
			} else if value.Valid {
				ra.Attrs = int32(value.Int64)
			}
		case replyarea.FieldMeta:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value.Valid {
				ra.Meta = value.String
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ReplyArea.
// This includes values selected through modifiers, order, etc.
func (ra *ReplyArea) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// Update returns a builder for updating this ReplyArea.
// Note that you need to call ReplyArea.Unwrap() before calling this method if this ReplyArea
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *ReplyArea) Update() *ReplyAreaUpdateOne {
	return NewReplyAreaClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the ReplyArea entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *ReplyArea) Unwrap() *ReplyArea {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReplyArea is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *ReplyArea) String() string {
	var builder strings.Builder
	builder.WriteString("ReplyArea(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ra.ID))
	builder.WriteString("oid=")
	builder.WriteString(fmt.Sprintf("%v", ra.Oid))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ra.Type))
	builder.WriteString(", ")
	builder.WriteString("mid=")
	builder.WriteString(fmt.Sprintf("%v", ra.Mid))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", ra.Count))
	builder.WriteString(", ")
	builder.WriteString("root_count=")
	builder.WriteString(fmt.Sprintf("%v", ra.RootCount))
	builder.WriteString(", ")
	builder.WriteString("acount=")
	builder.WriteString(fmt.Sprintf("%v", ra.Acount))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ra.State))
	builder.WriteString(", ")
	builder.WriteString("ctime=")
	builder.WriteString(ra.Ctime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(ra.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("attrs=")
	builder.WriteString(fmt.Sprintf("%v", ra.Attrs))
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(ra.Meta)
	builder.WriteByte(')')
	return builder.String()
}

// ReplyAreas is a parsable slice of ReplyArea.
type ReplyAreas []*ReplyArea
