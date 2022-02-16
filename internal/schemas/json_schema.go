package schemas

type JsonSchema struct {
	Type                 string                `json:"type,omitempty"`
	Title                string                `json:"title,omitempty"`
	Description          string                `json:"description,omitempty"`
	AdditionalProperties *bool                 `json:"additionalProperties,omitempty"`
	Default              interface{}           `json:"default,omitempty"`
	Enum                 []interface{}         `json:"enum,omitempty"`
	Properties           map[string]JsonSchema `json:"properties,omitempty"`
}
