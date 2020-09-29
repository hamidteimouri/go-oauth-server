package main

import (
	"fmt"
	"github.com/hamidteimouri/go-oauth-server/pkg/bootstrap"
)

func main() {
	/* run the application */
	bootstrap.Run()
	fmt.Println("application is running now... enjoy it!")
}
