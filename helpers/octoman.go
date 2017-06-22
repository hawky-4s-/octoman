package helpers

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"html/template"
)

var (
	// CommitHash contains the current Git revision. Use make to build to make
	// sure this gets set.
	CommitHash string

	// BuildDate contains the date of the current build.
	BuildDate string

	octomanInfo *OctomanInfo
)

func init() {
	octomanInfo = &OctomanInfo{
		Version:    CurrentOctomanVersion.String(),
		CommitHash: CommitHash,
		BuildDate:  BuildDate,
		Generator:  template.HTML(fmt.Sprintf(`<meta name="generator" content="Octoman %s" />`, CurrentOctomanVersion.String())),
	}
}

// OctomanVersion represents the Octoman build version.
type OctomanVersion struct {
	// Major and minor version.
	Number float32

	// Increment this for bug releases
	PatchLevel int

	// OctomanVersionSuffix is the suffix used in the Octoman version string.
	// It will be blank for release versions.
	Suffix string
}

func (v OctomanVersion) String() string {
	return octomanVersion(v.Number, v.PatchLevel, v.Suffix)
}

// ReleaseVersion represents the release version.
func (v OctomanVersion) ReleaseVersion() OctomanVersion {
	v.Suffix = ""
	return v
}

// Next returns the next Octoman release version.
func (v OctomanVersion) Next() OctomanVersion {
	return OctomanVersion{Number: v.Number + 0.01}
}

// Pre returns the previous Octoman release version.
func (v OctomanVersion) Prev() OctomanVersion {
	return OctomanVersion{Number: v.Number - 0.01}
}

// NextPatchLevel returns the next patch/bugfix Octoman version.
// This will be a patch increment on the previous Octoman version.
func (v OctomanVersion) NextPatchLevel(level int) OctomanVersion {
	return OctomanVersion{Number: v.Number - 0.01, PatchLevel: level}
}

// CurrentOctomanVersion represents the current build version.
// This should be the only one.
var CurrentOctomanVersion = OctomanVersion{
	Number:     0.01,
	PatchLevel: 0,
	Suffix:     "-DEV",
}

func octomanVersion(version float32, patchVersion int, suffix string) string {
	if patchVersion > 0 {
		return fmt.Sprintf("%.2f.%d%s", version, patchVersion, suffix)
	}
	return fmt.Sprintf("%.2f%s", version, suffix)
}

// CompareVersion compares the given version string or number against the
// running Octoman version.
// It returns -1 if the given version is less than, 0 if equal and 1 if greater than
// the running version.
func CompareVersion(version interface{}) int {
	return compareVersions(CurrentOctomanVersion.Number, CurrentOctomanVersion.PatchLevel, version)
}

func compareVersions(inVersion float32, inPatchVersion int, in interface{}) int {
	switch d := in.(type) {
	case float64:
		return compareFloatVersions(inVersion, float32(d))
	case float32:
		return compareFloatVersions(inVersion, d)
	case int:
		return compareFloatVersions(inVersion, float32(d))
	case int32:
		return compareFloatVersions(inVersion, float32(d))
	case int64:
		return compareFloatVersions(inVersion, float32(d))
	default:
		s, err := cast.ToStringE(in)
		if err != nil {
			return -1
		}

		var (
			v float32
			p int
		)

		if strings.Count(s, ".") == 2 {
			li := strings.LastIndex(s, ".")
			p = cast.ToInt(s[li+1:])
			s = s[:li]
		}

		v = float32(cast.ToFloat64(s))

		if v == inVersion && p == inPatchVersion {
			return 0
		}

		if v < inVersion || (v == inVersion && p < inPatchVersion) {
			return -1
		}

		return 1
	}
}

func compareFloatVersions(version float32, v float32) int {
	if v == version {
		return 0
	}
	if v < version {
		return -1
	}
	return 1
}

// OctomanInfo contains information about the current Octoman environment
type OctomanInfo struct {
	Version    string
	Generator  template.HTML
	CommitHash string
	BuildDate  string
}
