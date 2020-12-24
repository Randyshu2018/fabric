/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package runner

import (
	"time"

	"github.com/Randyshu2018/fabric/integration/helpers"
)

const DefaultStartTimeout = 45 * time.Second

// DefaultNamer is the default naming function.
var DefaultNamer NameFunc = helpers.UniqueName

// A NameFunc is used to generate container names.
type NameFunc func() string