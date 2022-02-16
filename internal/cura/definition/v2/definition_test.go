package v2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAllMetadataFieldsMapped(t *testing.T) {
	mappedFields := getReflectedFields(Metadata{})
	for _, fileName := range sampleDefinitions {
		def := loadRawData(t, fileName)
		// check fields
		require.NotEmpty(t, def["metadata"])
		for n := range def["metadata"].(map[string]interface{}) {
			require.Contains(t, mappedFields, n)
		}
	}
}
