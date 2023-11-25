package main

import (
	"sync"

	"github.com/Ubn-Jr/hirob-be-core/internal/http"
	"github.com/Ubn-Jr/hirob-be-core/internal/mqtt"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)

	go http.SetupHTTPServer()

	go mqtt.Subscribe()

	wg.Wait()

}
