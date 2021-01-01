package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	wa "github.com/radovskyb/watcher"

	"github.com/synw/fwr/ws"
)

var w = wa.New()
var building = false
var changeRequestedWhileBuilding = false

func watch(verbose bool) {
	err := w.AddRecursive("./lib")
	if err != nil {
		panic("Can not add path lib")
	}
	err = w.AddRecursive("./web")
	if err != nil {
		panic("Can not add path web")
	}
	w.FilterOps(wa.Write, wa.Create, wa.Move, wa.Remove, wa.Rename)
	// lauch listener
	var mux sync.Mutex
	go func() {
		for {
			select {
			case e := <-w.Event:
				noBuild := strings.HasSuffix(e.Path, "generated_plugin_registrant.dart")
				if noBuild == false {
					if building == false {
						go build(&mux, e.Path, verbose)
					} else {
						if verbose {
							fmt.Println("Change requested while building, delaying next build")
						}
						changeRequestedWhileBuilding = true
					}
				}
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

func build(mux *sync.Mutex, path string, verbose bool) {
	building = true
	msg := "Change detected in " + path
	if verbose {
		fmt.Println(msg)
	}
	if verbose {
		fmt.Println("Running build ...")
	}
	mux.Lock()
	runBuildCmd()
	mux.Unlock()
	if verbose {
		fmt.Println("Build done")
	}
	ws.SendMsg(msg)
	building = false
	if changeRequestedWhileBuilding == true {
		changeRequestedWhileBuilding = false

		build(mux, path, verbose)
	}
}
