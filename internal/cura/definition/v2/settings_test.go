package v2

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"strings"
	"testing"
)

var sampleDefinitions = []string{
	"testdata/fdmprinter.def.json",
	"testdata/fdmextruder.def.json",
}

func getReflectedFields(v interface{}) []string {
	// reflect to json tags to create list of output fields
	fields := []string{}
	for i := 0; i < reflect.TypeOf(v).NumField(); i++ {
		if tag, ok := reflect.TypeOf(v).Field(i).Tag.Lookup("json"); !ok {
			continue
		} else if name := strings.Split(tag, ",")[0]; name == "" {
			continue
		} else {
			fields = append(fields, name)
		}
	}
	return fields
}

func flattenRawSettings(settings map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for n, p := range settings {
		result[n] = p
		if val, ok := p.(map[string]interface{})["children"]; ok {
			subSettings := flattenRawSettings(val.(map[string]interface{}))
			for nn, pp := range subSettings {
				result[nn] = pp
			}
		}
	}
	return result
}

func loadRawData(t *testing.T, fileName string) map[string]interface{} {
	data, err := os.ReadFile(fileName)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	// decode json and find all settings
	def := map[string]interface{}{}
	err = json.Unmarshal(data, &def)
	require.NoError(t, err)

	return def
}

func TestAllSettingPropertiesMapped(t *testing.T) {
	mappedFields := getReflectedFields(Setting{})
	for _, fileName := range sampleDefinitions {
		def := loadRawData(t, fileName)

		// iterate through properties
		require.NotNil(t, def["settings"])
		settings := flattenRawSettings(def["settings"].(map[string]interface{}))
		require.NotEmpty(t, settings)

		// iterate over all settings and collect possible keys
		for _, s := range settings {
			for key := range s.(map[string]interface{}) {
				require.Contains(t, mappedFields, key)
			}
		}
	}
}

func TestAllSettingExportedFieldsAreSet(t *testing.T) {
	var exportedFields []string
	for n := range (Setting{}).AllEntries() {
		exportedFields = append(exportedFields, n)
	}
	// append children to exported fields to match, because in map this field omitted
	exportedFields = append(exportedFields, "children")

	expectedFields := getReflectedFields(Setting{})
	require.ElementsMatch(t, expectedFields, exportedFields)
}

func TestDecodeSettingsWorks(t *testing.T) {
	data, err := os.ReadFile("testdata/settings.sample.json")
	require.NoError(t, err)

	var settings Settings
	err = json.Unmarshal(data, &settings)
	require.NoError(t, err)
}

func TestSettings_Flatten(t *testing.T) {
	// test settings for flattening
	data, err := os.ReadFile("testdata/settings.sample.json")
	require.NoError(t, err)

	var settings Settings
	err = json.Unmarshal(data, &settings)
	require.NoError(t, err)

	assert.Len(t, settings, 2)
	assert.Len(t, settings.Flatten(), 5)
}
