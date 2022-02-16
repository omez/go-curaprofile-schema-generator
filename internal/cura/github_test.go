package cura

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetReleases(t *testing.T) {
	releases, err := GetReleases()
	require.NoError(t, err)
	assert.NotEmpty(t, releases)
	t.Logf("%+v", releases)
}

func TestGetLatestRelease(t *testing.T) {
	release, err := GetLatestRelease()
	require.NoError(t, err)
	assert.NotEmpty(t, release)
	t.Logf("%+v", release)
}
