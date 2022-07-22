package avro

const (
	SchemaKeyType   = "type"
	SchemaKeyName   = "name"
	SchemaKeyItems  = "items"
	SchemaKeyFields = "fields"
)

type schema struct {
	Type   any     `json:"type"`
	Name   string  `json:"name"`
	Fields []field `json:"fields"`
}

type field struct {
	Name string   `json:"name"`
	Type []string `json:"type"`
}
