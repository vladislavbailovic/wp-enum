package main

import (
	"wp-user-enum/pkg/cli"
	"wp-user-enum/pkg/data"
	"wp-user-enum/pkg/enum"
	wp_http "wp-user-enum/pkg/http"
)

func main() {
	params := cli.GetFlags()
	if params.URL == "" {
		panic("URL required")
	}

	kind := data.EnumerationType(params.Kind)

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	if params.RandomUA {
		ua := wp_http.NewRandomUA()
		client.SetAgent(&ua)
	}
	enumeration, err := enum.Enumerate(kind, params.URL)
	if err != nil {
		panic(err)
	}

	res, err := enumeration(client, params)
	if err != nil {
		panic(err)
	}

	cli.Print(res, params)
}
