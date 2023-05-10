package create

import (
	"fmt"

	pb "github.com/schollz/progressbar/v3"
)

func newBar(total int, stage, title string) *pb.ProgressBar {
	return pb.NewOptions(total,
		pb.OptionEnableColorCodes(true),
		pb.OptionSetWidth(15),
		pb.OptionShowCount(),
		pb.OptionSetDescription(fmt.Sprintf("[cyan][%s][reset] %s ...", stage, title)),
		pb.OptionSetTheme(pb.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		pb.OptionOnCompletion(func() {
			fmt.Println()
		}),
	)
}
