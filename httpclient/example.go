package go_httpclient

import (
	gohttp "github.com/Alejandro-Rojas/Http-Client/httpclient"
)

func exampleUsage() {
	client := gohttp.New()

	client.GET()

}
