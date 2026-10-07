// SPDX-FileCopyrightText: Copyright The Lima Authors
// SPDX-License-Identifier: Apache-2.0

package limainfo

import (
	"testing"

	"gotest.tools/v3/assert"

	// Register the internal qemu driver.
	_ "github.com/lima-vm/lima/v2/pkg/driver/qemu"
)

func TestDriverFeatures(t *testing.T) {
	features := driverFeatures(t.Context(), "qemu")
	assert.Assert(t, features != nil)
	assert.Assert(t, features.CanSnapshot)

	assert.Assert(t, driverFeatures(t.Context(), "no-such-driver") == nil)
}
