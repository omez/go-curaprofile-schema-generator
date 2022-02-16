package v2

type Settings map[string]Setting

// Flatten - flattens list of settings
func (s Settings) Flatten() Settings {
	result := Settings{}
	for n, v := range s {
		result[n] = v
		for nn, vv := range v.Children.Flatten() {
			result[nn] = vv
		}
	}
	return result
}

type Setting struct {
	Label                string      `json:"label,omitempty"`
	Type                 string      `json:"type,omitempty"`
	Description          string      `json:"description,omitempty"`
	Icon                 string      `json:"icon,omitempty"`
	SettablePerMesh      interface{} `json:"settable_per_mesh,omitempty"`
	SettablePerExtruder  interface{} `json:"settable_per_extruder,omitempty"`
	SettablePerMeshgroup interface{} `json:"settable_per_meshgroup,omitempty"`
	SettableGlobally     interface{} `json:"settable_globally,omitempty"`
	Enabled              interface{} `json:"enabled,omitempty"`
	Unit                 interface{} `json:"unit,omitempty"`
	Value                interface{} `json:"value,omitempty"`
	DefaultValue         interface{} `json:"default_value,omitempty"`
	MinimumValue         interface{} `json:"minimum_value,omitempty"`
	MaximumValue         interface{} `json:"maximum_value,omitempty"`
	MinimumValueWarning  interface{} `json:"minimum_value_warning,omitempty"`
	MaximumValueWarning  interface{} `json:"maximum_value_warning,omitempty"`
	WarningValue         interface{} `json:"warning_value,omitempty"`
	Options              interface{} `json:"options,omitempty"`
	Resolve              interface{} `json:"resolve,omitempty"`
	LimitToExtruder      interface{} `json:"limit_to_extruder,omitempty"`
	Children             Settings    `json:"children,omitempty"`
}

// AllEntries - returns all entries that could be exported
func (s Setting) AllEntries() map[string]interface{} {
	return map[string]interface{}{
		"label":                  s.Label,
		"type":                   s.Type,
		"description":            s.Description,
		"icon":                   s.Icon,
		"settable_per_mesh":      s.SettablePerMesh,
		"settable_per_extruder":  s.SettablePerExtruder,
		"settable_per_meshgroup": s.SettablePerMeshgroup,
		"settable_globally":      s.SettableGlobally,
		"enabled":                s.SettableGlobally,
		"unit":                   s.Unit,
		"value":                  s.Value,
		"default_value":          s.DefaultValue,
		"minimum_value":          s.MinimumValue,
		"maximum_value":          s.MaximumValue,
		"minimum_value_warning":  s.MinimumValueWarning,
		"maximum_value_warning":  s.MaximumValueWarning,
		"options":                s.Options,
		"warning_value":          s.WarningValue,
		"limit_to_extruder":      s.LimitToExtruder,
		"resolve":                s.Resolve,
	}
}

// Entries - returns map of fields that are not null
func (s Setting) Entries() map[string]interface{} {
	result := map[string]interface{}{}
	for n, v := range s.AllEntries() {
		if v != nil {
			result[n] = v
		}
	}
	return result
}
