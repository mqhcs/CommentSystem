package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// ReplyArea holds the schema definition for the ReplyArea entity.
type ReplyArea struct {
	ent.Schema
}

// Fields of the ReplyArea.
func (ReplyArea) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("oid"),
		field.Int8("type"),
		field.Int64("mid"),
		field.Int32("count"),
		field.Int32("root_count"),
		field.Int32("acount"),
		field.Int8("state"),
		field.Time("ctime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("mtime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Int32("attrs"),
		field.String("meta"),
	}
}

// Edges of the ReplyArea.
func (ReplyArea) Edges() []ent.Edge {
	return nil
}
