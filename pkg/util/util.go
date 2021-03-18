package util

import (
	"bytes"
	"context"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/describe"
)

func StreamLogs(request *rest.Request) (*bytes.Buffer, error){
	return func() (*bytes.Buffer, error) {
		bo := &bytes.Buffer{}
		req, err := request.Stream(context.TODO())
		if err != nil {
			return nil, err
		}
		defer req.Close()

		if _, err = io.Copy(bo, req); err != nil {
			return nil, err
		}
		return bo, nil
	}()

}

func GetNamespace(flags *genericclioptions.ConfigFlags) string{
	namespace, _, err := flags.ToRawKubeConfigLoader().Namespace()
	if err != nil || len(namespace) == 0 {
		namespace = "default"
	}
	return namespace
}

func PrintPodInfo(pod *corev1.Pod, w describe.PrefixWriter) {

	if pod.Spec.NodeName == "" {
		w.Write(describe.LEVEL_0, "Node:\t<none>\n")
	} else {
		w.Write(describe.LEVEL_0, "Node:\t%s\n", pod.Spec.NodeName+"/"+pod.Status.HostIP)
	}
	w.Write(describe.LEVEL_0, "Status:\t%s\n", string(pod.Status.Phase))

	if len(pod.Status.Conditions) > 0 {
		w.Write(describe.LEVEL_0, "Conditions:\n  Type\tStatus\n")
		for _, k := range pod.Status.Conditions {
			w.Write(describe.LEVEL_1, "%v \t%v \n", k.Type, k.Status)

		}
	}
}