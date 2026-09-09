package gormutil

const (
	defaultOffset = 0
	defaultLimit  = 20
)

// OffsetLimit holds pagination parameters with defaults applied.
type OffsetLimit struct {
	Offset int
	Limit  int
}

// DePointer dereferences optional int64 pointers and applies defaults.
func DePointer(offset, limit *int64) OffsetLimit {
	ol := OffsetLimit{Offset: defaultOffset, Limit: defaultLimit}
	if offset != nil {
		ol.Offset = int(*offset)
	}
	if limit != nil {
		ol.Limit = int(*limit)
	}
	return ol
}
