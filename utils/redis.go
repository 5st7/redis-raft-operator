package utils

import (
	"fmt"
	dbv1alpha1 "redis-raft-operator/api/v1alpha1"
)

type RedisPod struct {
	PodName   string
	Namespace string
}

func getRedisHostname(redisInfo RedisPod, cr *dbv1alpha1.RedisRaft) string {
	fqdn := fmt.Sprintf("%s.%s-headless.%s.svc", redisInfo.PodName, cr.ObjectMeta.Name, cr.Namespace)
	return fqdn
}
