package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// ReplyIndex holds the schema definition for the ReplyIndex entity.
type ReplyIndex struct {
	ent.Schema
}

// Fields of the ReplyIndex.
func (ReplyIndex) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("oid"),
		field.Int8("type"),
		field.Int64("mid"),
		field.Int64("root"),
		field.Int64("parent"),
		field.Int32("count"),
		field.Int32("rcount"),
		field.Int32("acount"),
		field.Int32("likes"),
		field.Int32("hates"),
		field.Int32("report_count"),
		field.Int8("state"),
		field.Time("ctime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("mtime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the ReplyIndex.
func (ReplyIndex) Edges() []ent.Edge {
	return nil
}
