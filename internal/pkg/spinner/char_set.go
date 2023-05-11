package spinner

const (
	UpDownCharsIndex int = iota
	DanceCharsIndex
	DefaultCharsIndex
	RodsCharsIndex
	DotsCharsIndex
	PipeCharsIndex
)

var CharSets = map[int][]string{
	UpDownCharsIndex:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	DanceCharsIndex:   {"▖", "▘", "▝", "▗"},
	DefaultCharsIndex: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	RodsCharsIndex:    {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	DotsCharsIndex:    {".", "..", "...", "....", ".....", "......", ".....", "....", "...", "..", "."},
	PipeCharsIndex:    {"|", "/", "-", "\\"},
}
