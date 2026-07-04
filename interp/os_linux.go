package interp

import (
	"io/fs"
	"time"
)

func getAtime(info fs.FileInfo) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Note that these are int32 on 32-bit platforms.
