package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Reply holds the schema definition for the Reply entity.
type Reply struct {
	ent.Schema
}

// Fields of the Reply.
func (Reply) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("message"),
		field.String("ats"),
		field.Int64("ip"),
		field.Int8("plat"),
		field.String("device"),
		field.String("version"),
		field.Time("ctime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("mtime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.String("topics"),
		field.String("addr"),
	}
}

// Edges of the Reply.
func (Reply) Edges() []ent.Edge {
	return nil
}
