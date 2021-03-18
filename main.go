package main

import (
	"github.com/HamzaZo/kubectl-pod-viewer/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth" //required for auth
	"os"
)

func main() {
	root := cmd.NewCmdPodViewer(genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}

}
