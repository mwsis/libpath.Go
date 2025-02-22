package libpath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 1
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x4001
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, VersionMajor)
	require.Equal(t, Expected_VersionMinor, VersionMinor)
	require.Equal(t, Expected_VersionPatch, VersionPatch)
	require.Equal(t, Expected_VersionAB, VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000000100004001), Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.1.0-alpha1", VersionString())
}
