package descriptor_test

import (
	"slices"
	"testing"

	"github.com/go-4devs/console/output/descriptor"
)

func TestDescriptorsSorted(t *testing.T) {
	t.Parallel()

	names := descriptor.Descriptors()
	if !slices.IsSorted(names) {
		t.Fatalf("descriptors not sorted: %v", names)
	}

	if !slices.Contains(names, descriptor.FormatTxt) {
		t.Fatalf("missing %s: %v", descriptor.FormatTxt, names)
	}
}
