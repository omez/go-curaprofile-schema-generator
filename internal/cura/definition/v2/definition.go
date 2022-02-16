package v2

type Definition struct {
	Name     string   `json:"name"`
	Version  int      `json:"version"`
	Metadata Metadata `json:"metadata"`
	Settings Settings `json:"settings"`
}

type Metadata struct {
	Type                      string      `json:"type"`
	Author                    string      `json:"author"`
	Manufacturer              string      `json:"manufacturer"`
	SettingVersion            int         `json:"setting_version"`
	FileFormats               string      `json:"file_formats"`
	Visible                   bool        `json:"visible"`
	HasMaterials              bool        `json:"has_materials"`
	HasVariants               bool        `json:"has_variants"`
	HasMachineQuality         bool        `json:"has_machine_quality"`
	PreferredMaterial         string      `json:"preferred_material"`
	PreferredQualityType      string      `json:"preferred_quality_type"`
	MachineExtruderTrains     interface{} `json:"machine_extruder_trains"`
	SupportsUsbConnection     bool        `json:"supports_usb_connection"`
	SupportsNetworkConnection bool        `json:"supports_network_connection"`
	Position                  interface{} `json:"position"`
}
