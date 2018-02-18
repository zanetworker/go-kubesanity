package network

import (
	"fmt"
	"os"

	"github.com/zanetworker/go-kubesanity/pkg/log"
	typev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func newKubeClient() v1.CoreV1Interface {
	kubeconfigPath, ok := os.LookupEnv("KUBECONFIG_PATH")
	if !ok {
		log.FatalS("KUBECONFIG_PATH was not set")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	kubeclient := clientset.CoreV1()

	return kubeclient
}

//CheckDuplicatePodIP checks if two pods have the same IP in all namespaces
func CheckDuplicatePodIP() (bool, error) {
	kubeclient := newKubeClient()

	podIPs := make(map[string]typev1.Pod)
	podList, err := kubeclient.Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Error(err.Error())
	}

	for _, pod := range podList.Items {
		otherPod, ipExists := podIPs[pod.Status.PodIP]
		if ipExists {
			return true, fmt.Errorf("Duplicate Service IP Address: %s/%s -> %s and %s/%s -> %s",
				pod.ObjectMeta.Namespace,
				pod.ObjectMeta.Name,
				pod.Status.PodIP,
				otherPod.ObjectMeta.Namespace,
				otherPod.ObjectMeta.Name,
				otherPod.Status.PodIP,
			)
		}
		podIPs[pod.Status.PodIP] = pod
	}
	log.Info("No duplicate pod IPs found!")
	return false, nil
}

//CheckDuplicateServiceIP checks if two services have the same IP in all namespaces
func CheckDuplicateServiceIP() (bool, error) {

	kubeclient := newKubeClient()

	serviceIPs := make(map[string]typev1.Service)
	serviceList, err := kubeclient.Services("").List(metav1.ListOptions{})

	if err != nil {
		log.Error(err.Error())
	}

	for _, service := range serviceList.Items {
		otherService, ipExists := serviceIPs[service.Spec.ClusterIP]
		if ipExists {
			return true, fmt.Errorf("Duplicate Service IP Address: %s/%s -> %s and %s/%s -> %s",
				service.ObjectMeta.Namespace,
				service.ObjectMeta.Name,
				service.Spec.ClusterIP,
				otherService.ObjectMeta.Namespace,
				otherService.ObjectMeta.Name,
				otherService.Spec.ClusterIP,
			)
		}
		serviceIPs[service.Spec.ClusterIP] = service
	}
	log.Info("No duplicate service IPs found!")
	return false, nil
}
