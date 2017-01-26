package main

import (
	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/predicates"
//	"k8s.io/kubernetes/plugin/pkg/scheduler/algorithm/priorities"
	 "github.com/golang/glog"
)

func main() {
	_ = predicates.EBSVolumeFilter
//	_ = priorities.BalancedResourceAllocationMap

	glog.Info("abcd")
}
