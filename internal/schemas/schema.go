package schemas

import (
	"fmt"
	v2 "go-curaprofile-schema-generator/internal/cura/definition/v2"
)

func GenerateSchema(def *v2.Definition) (*JsonSchema, error) {
	schema := &JsonSchema{
		Type:        "object",
		Description: "Definition schema",
		Properties: map[string]JsonSchema{
			"name":     {Type: "string"},
			"version":  {Type: "number", Default: 2},
			"inherits": {Type: "string", Default: "fdmprinter"},
			"metadata": {Type: "object", Properties: map[string]JsonSchema{
				"author":       {Type: "string"},
				"manufacturer": {Type: "string"},
				"visible":      {Type: "boolean", Default: true},
				//"file_formats": {Type: "string", Default: "text/x-gcode"},
				// TODO others
			}},
		},
	}

	//check if settings exists
	if def.Settings == nil {
		return nil, fmt.Errorf("settings do not present in definition")
	}

	// TODO generate metadata

	// generate overrides
	if s, err := generateOverridesSchema(def.Settings); err != nil {
		return nil, err
	} else {
		schema.Properties["overrides"] = *s
	}

	return schema, nil
}
