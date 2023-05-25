/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CenterPhase string

const (
	DataCenterReady   CenterPhase = "Ready"
	DataCenterUnReady CenterPhase = "UnReady"
)

type ResourceInfo struct {
	//可分配资源总量
	MemAllocatable int64 `json:"memAllocatable,omitempty"`
	//存储资源总量
	MemCapacity int64 `json:"memCapacity,omitempty"`
	//可创建数据卷数量
	VolumeCntAllocatable int `json:"volumeCntAllocatable,omitempty"`
	//存储卷总量
	VolumeCntCapacity int `json:"volumeCntCapacity,omitempty"`
	//预留资源
	ReserveResources v1.ResourceList `json:"reserveResources,omitempty"`
}

type SchedulerParams struct {
	Label      string `json:"label,omitempty" protobuf:"bytes,1,opt,name=label"`
	DirtyLabel string `json:"dirtyLabel,omitempty" protobuf:"bytes,2,opt,name=dirtyLabel"`
}

// DataCenterSpec defines the desired state of DataCenter
type DataCenterSpec struct {
	ResourceInfo ResourceInfo    `json:"resourceInfo,omitempty" protobuf:"bytes,1,opt,name=resourceInfo"`
	Scheduler    SchedulerParams `json:"scheduler,omitempty" protobuf:"bytes,2,opt,name=scheduler"`
}

type Idle struct {
	MemAllocatable       int64 `json:"memAllocatable"`
	VolumeCntAllocatable int   `json:"volumeCntAllocatable"`
}

// DataCenterStatus defines the observed state of DataCenter
type DataCenterStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of dataCenter
	Idle        Idle        `json:"idle,omitempty"`
	Status      CenterPhase `json:"status,omitempty"`
	UpdatedTime metav1.Time `json:"updatedTime,omitempty"`
	Message     string      `json:"message,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=dataCenter,scope=DataCenter
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DataCenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCenterSpec   `json:"spec,omitempty"`
	Status DataCenterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataCenterList contains a list of DataCenter
type DataCenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCenter `json:"items"`
}
