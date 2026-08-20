package main

import (
	"net/http"
	now "khulnasoft/bridge"
)

func main() {
	now.Start(http.HandlerFunc(__NOW_HANDLER_FUNC_NAME))
}
