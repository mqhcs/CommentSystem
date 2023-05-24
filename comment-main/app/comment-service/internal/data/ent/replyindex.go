// Code generated by ent, DO NOT EDIT.

package ent

import (
	"comment-main/app/comment-service/internal/data/ent/replyindex"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ReplyIndex is the model entity for the ReplyIndex schema.
type ReplyIndex struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Oid holds the value of the "oid" field.
	Oid int64 `json:"oid,omitempty"`
	// Type holds the value of the "type" field.
	Type int8 `json:"type,omitempty"`
	// Mid holds the value of the "mid" field.
	Mid int64 `json:"mid,omitempty"`
	// Root holds the value of the "root" field.
	Root int64 `json:"root,omitempty"`
	// Parent holds the value of the "parent" field.
	Parent int64 `json:"parent,omitempty"`
	// Count holds the value of the "count" field.
	Count int32 `json:"count,omitempty"`
	// Rcount holds the value of the "rcount" field.
	Rcount int32 `json:"rcount,omitempty"`
	// Acount holds the value of the "acount" field.
	Acount int32 `json:"acount,omitempty"`
	// Likes holds the value of the "likes" field.
	Likes int32 `json:"likes,omitempty"`
	// Hates holds the value of the "hates" field.
	Hates int32 `json:"hates,omitempty"`
	// ReportCount holds the value of the "report_count" field.
	ReportCount int32 `json:"report_count,omitempty"`
	// State holds the value of the "state" field.
	State int8 `json:"state,omitempty"`
	// Ctime holds the value of the "ctime" field.
	Ctime time.Time `json:"ctime,omitempty"`
	// Mtime holds the value of the "mtime" field.
	Mtime        time.Time `json:"mtime,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReplyIndex) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case replyindex.FieldID, replyindex.FieldOid, replyindex.FieldType, replyindex.FieldMid, replyindex.FieldRoot, replyindex.FieldParent, replyindex.FieldCount, replyindex.FieldRcount, replyindex.FieldAcount, replyindex.FieldLikes, replyindex.FieldHates, replyindex.FieldReportCount, replyindex.FieldState:
			values[i] = new(sql.NullInt64)
		case replyindex.FieldCtime, replyindex.FieldMtime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReplyIndex fields.
func (ri *ReplyIndex) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case replyindex.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ri.ID = int64(value.Int64)
		case replyindex.FieldOid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				ri.Oid = value.Int64
			}
		case replyindex.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ri.Type = int8(value.Int64)
			}
		case replyindex.FieldMid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mid", values[i])
			} else if value.Valid {
				ri.Mid = value.Int64
			}
		case replyindex.FieldRoot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field root", values[i])
			} else if value.Valid {
				ri.Root = value.Int64
			}
		case replyindex.FieldParent:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent", values[i])
			} else if value.Valid {
				ri.Parent = value.Int64
			}
		case replyindex.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				ri.Count = int32(value.Int64)
			}
		case replyindex.FieldRcount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rcount", values[i])
			} else if value.Valid {
				ri.Rcount = int32(value.Int64)
			}
		case replyindex.FieldAcount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field acount", values[i])
			} else if value.Valid {
				ri.Acount = int32(value.Int64)
			}
		case replyindex.FieldLikes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likes", values[i])
			} else if value.Valid {
				ri.Likes = int32(value.Int64)
			}
		case replyindex.FieldHates:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hates", values[i])
			} else if value.Valid {
				ri.Hates = int32(value.Int64)
			}
		case replyindex.FieldReportCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field report_count", values[i])
			} else if value.Valid {
				ri.ReportCount = int32(value.Int64)
			}
		case replyindex.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ri.State = int8(value.Int64)
			}
		case replyindex.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				ri.Ctime = value.Time
			}
		case replyindex.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				ri.Mtime = value.Time
			}
		default:
			ri.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ReplyIndex.
// This includes values selected through modifiers, order, etc.
func (ri *ReplyIndex) Value(name string) (ent.Value, error) {
	return ri.selectValues.Get(name)
}

// Update returns a builder for updating this ReplyIndex.
// Note that you need to call ReplyIndex.Unwrap() before calling this method if this ReplyIndex
// was returned from a transaction, and the transaction was committed or rolled back.
func (ri *ReplyIndex) Update() *ReplyIndexUpdateOne {
	return NewReplyIndexClient(ri.config).UpdateOne(ri)
}

// Unwrap unwraps the ReplyIndex entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ri *ReplyIndex) Unwrap() *ReplyIndex {
	_tx, ok := ri.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReplyIndex is not a transactional entity")
	}
	ri.config.driver = _tx.drv
	return ri
}

// String implements the fmt.Stringer.
func (ri *ReplyIndex) String() string {
	var builder strings.Builder
	builder.WriteString("ReplyIndex(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ri.ID))
	builder.WriteString("oid=")
	builder.WriteString(fmt.Sprintf("%v", ri.Oid))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ri.Type))
	builder.WriteString(", ")
	builder.WriteString("mid=")
	builder.WriteString(fmt.Sprintf("%v", ri.Mid))
	builder.WriteString(", ")
	builder.WriteString("root=")
	builder.WriteString(fmt.Sprintf("%v", ri.Root))
	builder.WriteString(", ")
	builder.WriteString("parent=")
	builder.WriteString(fmt.Sprintf("%v", ri.Parent))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", ri.Count))
	builder.WriteString(", ")
	builder.WriteString("rcount=")
	builder.WriteString(fmt.Sprintf("%v", ri.Rcount))
	builder.WriteString(", ")
	builder.WriteString("acount=")
	builder.WriteString(fmt.Sprintf("%v", ri.Acount))
	builder.WriteString(", ")
	builder.WriteString("likes=")
	builder.WriteString(fmt.Sprintf("%v", ri.Likes))
	builder.WriteString(", ")
	builder.WriteString("hates=")
	builder.WriteString(fmt.Sprintf("%v", ri.Hates))
	builder.WriteString(", ")
	builder.WriteString("report_count=")
	builder.WriteString(fmt.Sprintf("%v", ri.ReportCount))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ri.State))
	builder.WriteString(", ")
	builder.WriteString("ctime=")
	builder.WriteString(ri.Ctime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(ri.Mtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ReplyIndexes is a parsable slice of ReplyIndex.
type ReplyIndexes []*ReplyIndex
