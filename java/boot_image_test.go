// Copyright (C) 2021 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package java

import (
	"testing"
)

// Contains some simple tests for boot_image logic, additional tests can be found in
// apex/boot_image_test.go as the ART boot image requires modules from the ART apex.

func TestUnknownBootImage(t *testing.T) {
	testJavaError(t, "image_name: Unknown image name \\\"unknown\\\", expected one of art, boot", `
		boot_image {
			name: "unknown-boot-image",
			image_name: "unknown",
		}
`)
}

func TestUnknownPrebuiltBootImage(t *testing.T) {
	testJavaError(t, "image_name: Unknown image name \\\"unknown\\\", expected one of art, boot", `
		prebuilt_boot_image {
			name: "unknown-boot-image",
			image_name: "unknown",
		}
`)
}
