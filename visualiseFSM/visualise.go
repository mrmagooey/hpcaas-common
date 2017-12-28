package main

import (
	"fmt"

	"github.com/mrmagooey/fsmGraphviz"
	common "github.com/mrmagooey/hpcaas-common"
)

//go:generate /bin/bash -c "go build && ./visualiseFSM | dot -Tpng -Gdpi=300 -o daemonFsm.png"

func main() {
	graphVizString, _ := fsmGraphviz.CreateGraphvizString(common.DaemonStateInitializing.String(), common.DaemonFSM, "Daemon State")
	fmt.Println(graphVizString)
}
