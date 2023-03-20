package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisSpec `json:"spec"`
	//改crd暂时不需要status字段,在上边代码生成部分要 添加 // +genclient:noStatus
	//Status            RedisStatus `json:"status"`
}
type RedisSpec struct {
	// 参考 redis.yaml 文件的 spec部分
	Image    string `json:"image"`
	Port     int    `json:"port"`
	Target   int    `json:"target"`
	Password string `json:"password"`
}

//暂时不需要Redisstatus结构体
//type RedisStatus struct {
//	// 参考 redis.yaml 文件的 status部分
//	Ready bool `json:"ready"`
//}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Redis `json:"items"`
}
