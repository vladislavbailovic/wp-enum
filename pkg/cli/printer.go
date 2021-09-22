package cli

import (
	"fmt"
	"math"
	"os"
	"strings"
	"wp-user-enum/pkg/data"
)

func Print(results []data.ApiResponse, opts data.Constraints) {
	if opts.Pretty {
		prettyPrint(results)
		return
	}
	if len(results) == 0 {
		return
	}
	for _, result := range results {
		fmt.Printf("%s:%d\n", result.Username, result.UserID)
	}
}

func prettyPrint(results []data.ApiResponse) {
	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "No results")
		return
	}

	nameMax := float64(len("Username"))
	idMax := float64(len("User ID"))
	for _, result := range results {
		nameMax = math.Max(nameMax, float64(len(result.Username)))
		idMax = math.Max(idMax, float64(len(fmt.Sprintf("%d", result.UserID))))
	}

	padWidth := 1

	nameHdiv := fmt.Sprintf(strings.Repeat("-", 2*padWidth+int(nameMax)))
	idHdiv := fmt.Sprintf(strings.Repeat("-", 2*padWidth+int(idMax)))

	fmt.Fprintf(os.Stderr, "+%s+%s+\n", nameHdiv, idHdiv)
	fmt.Fprintf(os.Stderr, "| ")
	fmt.Printf(pad("Username", int(nameMax)))
	fmt.Fprintf(os.Stderr, " | ")
	fmt.Printf(pad("User ID", int(idMax)))
	fmt.Fprintf(os.Stderr, " |")
	fmt.Println()
	fmt.Fprintf(os.Stderr, "+%s+%s+\n", nameHdiv, idHdiv)
	for _, result := range results {
		id := fmt.Sprintf("%d", result.UserID)

		fmt.Fprintf(os.Stderr, "| ")
		fmt.Printf(pad(result.Username, int(nameMax)))
		fmt.Fprintf(os.Stderr, " | ")
		fmt.Printf(pad(id, int(idMax)))
		fmt.Fprintf(os.Stderr, " |")
		fmt.Println()
		fmt.Fprintf(os.Stderr, "+%s+%s+\n", nameHdiv, idHdiv)
	}
}

func pad(what string, max int) string {
	if len(what) >= max {
		return what[:max]
	}
	pad := int((max - len(what)) / 2)
	padding := strings.Repeat(" ", pad)
	padded := fmt.Sprintf("%s%s%s", padding, what, padding)

	if len(padded) < max {
		padded += " "
	}
	return padded
}
