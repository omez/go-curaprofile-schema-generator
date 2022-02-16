package schemas

import v2 "go-curaprofile-schema-generator/internal/cura/definition/v2"

// generateOverridesSchema creates override schema properties as [name]: [property schema]
func generateOverridesSchema(settings v2.Settings) (*JsonSchema, error) {
	schema := &JsonSchema{
		Type:                 "object",
		Description:          "Overrides for cura settings",
		Properties:           map[string]JsonSchema{},
		AdditionalProperties: &[]bool{false}[0],
	}

	for name, setting := range settings.Flatten() {
		overridePropertySchema, err := generateOverridePropertySchema(setting)
		if err != nil {
			return nil, err
		}
		schema.Properties[name] = *overridePropertySchema
	}

	return schema, nil
}

// generateOverridePropertySchema generates schema for each setting
func generateOverridePropertySchema(setting v2.Setting) (*JsonSchema, error) {
	schema := &JsonSchema{
		Type:                 "object",
		Title:                setting.Label,
		Description:          setting.Description,
		Properties:           map[string]JsonSchema{},
		AdditionalProperties: &[]bool{false}[0],
	}

	skippedFields := []string{
		"label",
		"description",
		"icon",
	}

	for n, v := range setting.AllEntries() {
		skipped := false
		for _, f := range skippedFields {
			if n == f {
				skipped = true
				break
			}
		}
		if skipped {
			continue
		}

		// filter metadata fields
		schema.Properties[n] = JsonSchema{
			Default: v,
		}
	}

	return schema, nil
}
