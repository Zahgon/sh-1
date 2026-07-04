//go:build !linux && !darwin && !windows

package interp

import (
	"io/fs"
	"time"
)

func getAtime(info fs.FileInfo) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
