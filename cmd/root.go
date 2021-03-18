package cmd

import (
	"bytes"
	"fmt"
	"github.com/HamzaZo/kubectl-pod-viewer/pkg/pods"
	"github.com/HamzaZo/kubectl-pod-viewer/pkg/util"
	"github.com/spf13/cobra"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/describe"
	"text/tabwriter"
)

var (
	viewerExample = `
    # full view of pod 
    %[1]s pod-viewer <pod-name>

    # view pod in different namespace
    %[1]s pod-viewer <pod-name> -n/--namespace <ns>

    # view pod in different context
    %[1]s pod-viewer <pod-name> --context <ctx>

    # view pod by providing kubeconfig
    %[1]s pod-viewer <pod-name> --kubeconfig <kcfg>

    # view logs of container nginx in pod frontend
    %[1]s pod-viewer <frontend> -c nginx

    # display only the most recent 20 line of output in pod frontend
    %[1]s pod-viewer <frontend> -t/--tail=20 

`
)

type ViewerPodOptions struct {
	configFlags *genericclioptions.ConfigFlags
	ioStreams   genericclioptions.IOStreams

	kubeClient             kubernetes.Interface
	userSpecifiedPodName   string
	userSpecifiedNamespace string
	userSpecifiedTail      bool

	container string
	tailLines int64
	args      []string
}

func NewPodViewerOption(streams genericclioptions.IOStreams) *ViewerPodOptions {
	return &ViewerPodOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		ioStreams:   streams,
	}
}

//NewCmdPodViewer runs the pod-view root command
func NewCmdPodViewer(streams genericclioptions.IOStreams) *cobra.Command {
	v := NewPodViewerOption(streams)
	cmd := &cobra.Command{
		Use:          "pod-viewer [pod-name] [flags]",
		Short:        "A Full view of kubernetes pod",
		Example:      fmt.Sprintf(viewerExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := v.Complete(cmd, args); err != nil {
				return err
			}
			if err := v.Validate(); err != nil {
				return err
			}
			err := v.Run(cmd)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().Int64P("tail", "t", v.tailLines, "Lines of recent log file to display")
	cmd.Flags().StringVarP(&v.container, "container", "c", v.container, "Print the logs of this container")
	v.configFlags.AddFlags(cmd.Flags())

	return cmd
}

func (v *ViewerPodOptions) Complete(cmd *cobra.Command, args []string) error {
	v.args = args
	if len(args) > 0 {
		if len(v.userSpecifiedPodName) > 0 {
			return fmt.Errorf("cannot specify multiple pods, only one pod name is allowed")
		}

		v.userSpecifiedPodName = args[0]
	}

	cfg, err := v.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	v.kubeClient, err = kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}

	v.container, err = cmd.Flags().GetString("container")
	if err != nil {
		return err
	}

	v.tailLines, err = cmd.Flags().GetInt64("tail")
	if err != nil {
		return err
	}

	v.userSpecifiedNamespace = util.GetNamespace(v.configFlags)

	return nil
}

//LogOptions defines log options for a given pod
func (v *ViewerPodOptions) LogOptions(cmd *cobra.Command) *corev1.PodLogOptions {
	v.userSpecifiedTail = cmd.Flag("tail").Changed

	logOptions := &corev1.PodLogOptions{}
	if v.userSpecifiedTail {
		logOptions.Container = v.container
		logOptions.TailLines = &v.tailLines
	} else {
		logOptions.Container = v.container
	}

	return logOptions
}

//Validate the root cmd arguments
func (v *ViewerPodOptions) Validate() error {
	if len(v.args) == 0 {
		return fmt.Errorf("at least one argument is required")
	}

	if len(v.args) > 1 {
		return fmt.Errorf("only one argument is allowed ")
	}
	return nil
}

// Overview creates a full overview of kubernetes pods
func (v ViewerPodOptions) Overview(podName, namespace string, client kubernetes.Interface, cmd *cobra.Command) (string, error) {
	buf := &bytes.Buffer{}

	obj, err := pods.Get(client, podName, namespace)
	if err != nil {
		return "", err
	}

	events, err := pods.SearchEvents(client, namespace, obj)
	if err != nil {
		return "", err
	}
	log := pods.Logs(client, namespace, podName, v.LogOptions(cmd))

	s, err := util.StreamLogs(log)
	if err != nil {
		return "", err
	}

	return func(out io.Writer) (string, error) {
		m := tabwriter.NewWriter(buf, 0, 8, 2, ' ', 0)
		w := describe.NewPrefixWriter(m)
		w.Write(describe.LEVEL_0, "Name:\t%s\n", podName)
		w.Write(describe.LEVEL_0, "Namespace:\t%s\n", namespace)
		util.PrintPodInfo(obj, w)
		if events != nil {
			describe.DescribeEvents(events, w)
		}
		w.Write(describe.LEVEL_0, "Logs:  \n%v", s.String())
		m.Flush()
		return buf.String(), nil
	}(buf)

}

func (v ViewerPodOptions) Run(cmd *cobra.Command) error {
	str, err := v.Overview(v.userSpecifiedPodName, v.userSpecifiedNamespace, v.kubeClient, cmd)
	if err != nil {
		return err
	}
	fmt.Fprintf(v.ioStreams.Out, str)
	return nil
}
