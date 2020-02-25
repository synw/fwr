package main

import (
	"fmt"
	"sync"
	"time"

	wa "github.com/radovskyb/watcher"

	"github.com/synw/fwr/ws"
)

var w = wa.New()

func watch() {
	err := w.AddRecursive("lib")
	if err != nil {
		panic("Can not add path lib")
	}
	err = w.AddRecursive("web")
	if err != nil {
		panic("Can not add path web")
	}
	w.FilterOps(wa.Write, wa.Create, wa.Move, wa.Remove, wa.Rename)
	// lauch listener
	var mux sync.Mutex
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case e := <-w.Event:
				msg := "Change detected in " + e.Path
				fmt.Println(msg)
				wg.Add(1)
				runBuild(&wg, &mux)
				//fmt.Println("Reload")
				ws.SendMsg(msg)
				wg.Wait()
			case err := <-w.Error:
				msg := "Watcher error " + err.Error()
				fmt.Println(msg)
			case <-w.Closed:
				msg := "Watcher closed"
				fmt.Println(msg)
				return
			}
		}
	}()
	fmt.Println("👀 Watching for changes in lib/ and web/")
	// start listening
	err = w.Start(time.Millisecond * 200)
	if err != nil {
		panic("Error starting the watcher")
	}
}

func runBuild(wg *sync.WaitGroup, mux *sync.Mutex) {
	mux.Lock()
	defer mux.Unlock()
	runBuildCmd()
	wg.Done()
}
