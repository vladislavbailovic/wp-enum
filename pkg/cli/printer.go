package cli

import (
	"fmt"
	"os"
	"strings"
	"wp-enum/pkg/data"
)

func formatLine(parts []string) string {
	return strings.Join(parts, "\t")
}

func Print(results []data.ApiResponse, opts data.Constraints) {
	if len(results) == 0 {
		return
	}
	fmt.Fprintln(os.Stderr, formatLine([]string{"Username", "User ID"}))
	for _, result := range results {
		fmt.Println(formatLine([]string{
			result.Name,
			fmt.Sprintf("%d", result.Id),
		}))
	}
}
