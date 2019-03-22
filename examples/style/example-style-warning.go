package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
)

func main() {
	// creating default console styler
	io := style.NewConsoleGoStyler()

	// according to my terminal size (default: 120)
	io.SetMaxLineLength(80)

	// use simple strings for short messages
	io.Warning("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec nisl nec mi blandit malesuada. Nunc augue risus, posuere vitae feugiat quis, pulvinar non ligula.")

	// consider using arrays when displaying long messages
	io.WarningArray([]string{
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
	})
}