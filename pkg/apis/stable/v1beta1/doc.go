//拷贝官方示例代码 修改groupName部分，让他和我们的group[crd.yaml部分内容]相同

// +k8s:deepcopy-gen=package
// +groupName=stable.example.com

// Package v1beta1 is the v1beta1 version of the API.
package v1beta1 // import "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
