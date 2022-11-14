package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	corev1 "k8s.io/api/core/v1"
)

const (
	redisPort = 6379
)

func generateServiceSpec(serviceMeta metav1.ObjectMeta, ownerRef metav1.OwnerReference, headless bool) *corev1.Service {
	service := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: serviceMeta,
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: "",
			Selector:  serviceMeta.GetLabels(),
			Ports: []corev1.ServicePort{
				{
					Name:       "redis",
					Port:       redisPort,
					TargetPort: intstr.FromInt(redisPort),
					Protocol:   corev1.ProtocolTCP,
				},
			},
		},
	}
	if headless {
		service.Spec.ClusterIP = "None"
	}

	service.SetOwnerReferences(append(service.GetOwnerReferences(), ownerRef))
	return service
}
