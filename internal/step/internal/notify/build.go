package notify

import (
	"time"
)

type Build struct {
	Name     string
	Status   string
	Url      string
	Steps    []string
	Created  time.Time
	Finished time.Time
	Elapsed  time.Duration
}
