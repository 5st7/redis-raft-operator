package utils

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	dbv1alpha1 "redis-raft-operator/api/v1alpha1"
)

func NewStatefulSetSpec(instance *dbv1alpha1.RedisRaft) *appsv1.StatefulSet {
	labels := map[string]string{
		"app":        "redisraft",
		"controller": instance.Name,
	}

	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(instance, dbv1alpha1.GroupVersion.WithKind("RedisRaft")),
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &instance.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "redis",
							Image:   "redislabs/ng-redis-raft:latest",
							Command: []string{"redis-server"},
							Args: []string{
								"--port 6379",
								"--dbfilename raft.rdb",
								"--loadmodule /redisraft.so",
								"--raft.log-filename raftlog.db",
								"--raft.addr localhost:6379",
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "redis",
									ContainerPort: redisPort,
								},
							},
						},
					},
				},
			},
		},
	}
}
