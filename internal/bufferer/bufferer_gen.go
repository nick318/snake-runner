// Code generated by gonstructor -type Bufferer -init init; DO NOT EDIT.

package bufferer

import "time"

func NewBufferer(
	size int,
	duration time.Duration,
	flush func([]byte),
) *Bufferer {
	r := &Bufferer{
		size:     size,
		duration: duration,
		flush:    flush,
	}

	r.init()

	return r
}