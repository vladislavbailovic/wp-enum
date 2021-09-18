package main

import (
	"fmt"
	"wp-enum/cli"
	"wp-enum/pkg/enum"
	wp_http "wp-enum/pkg/http"
)

func main() {
	params := cli.GetFlags()
	if params.Url == "" {
		panic("URL required")
	}

	kind := enum.EnumerationType(params.Kind)

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	enumeration, err := enum.Enumerate(kind, params.Url)
	if err != nil {
		panic(err)
	}

	res, err := enumeration(client, params.Limit)
	if err != nil {
		panic(err)
	}

	for username, id := range res {
		fmt.Printf("%s:%d\n", username, id)
	}
}
