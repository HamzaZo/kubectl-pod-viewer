package pods

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/scheme"
)

func Get(client kubernetes.Interface, podName, namespace string) (*corev1.Pod, error) {
	return client.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
}

func SearchEvents(client kubernetes.Interface, namespace string, obj runtime.Object) (*corev1.EventList, error) {
	return client.CoreV1().Events(namespace).Search(scheme.Scheme, obj)
}

func Logs(client kubernetes.Interface, namespace, podName string, Options *corev1.PodLogOptions) *rest.Request {
	return client.CoreV1().Pods(namespace).GetLogs(podName, Options)
}
