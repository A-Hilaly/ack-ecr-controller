// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package repository

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktypes "github.com/aws-controllers-k8s/runtime/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sapirt "k8s.io/apimachinery/pkg/runtime"
	k8sctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	svcapitypes "github.com/aws-controllers-k8s/ecr-controller/apis/v2"
)

const (
	finalizerString = "finalizers.ecr.services.k8s.aws/Repository"
)

var (
	resourceGK = metav1.GroupKind{
		Group: "ecr.services.k8s.aws",
		Kind:  "Repository",
	}
)

// resourceDescriptor implements the
// `aws-service-operator-k8s/pkg/types.AWSResourceDescriptor` interface
type resourceDescriptor struct {
}

// GroupKind returns a Kubernetes metav1.GroupKind struct that describes the
// API Group and Kind of CRs described by the descriptor
func (d *resourceDescriptor) GroupKind() *metav1.GroupKind {
	return &resourceGK
}

// EmptyRuntimeObject returns an empty object prototype that may be used in
// apimachinery and k8s client operations
func (d *resourceDescriptor) EmptyRuntimeObject() k8sapirt.Object {
	return &svcapitypes.Repository{}
}

// ResourceFromRuntimeObject returns an AWSResource that has been initialized
// with the supplied runtime.Object
func (d *resourceDescriptor) ResourceFromRuntimeObject(
	obj k8sapirt.Object,
) acktypes.AWSResource {
	return &resource{
		ko: obj.(*svcapitypes.Repository),
	}
}

// Delta returns an `ackcompare.Delta` object containing the difference between
// one `AWSResource` and another.
func (d *resourceDescriptor) Delta(a, b acktypes.AWSResource) *ackcompare.Delta {
	return newResourceDelta(a.(*resource), b.(*resource))
}

// UpdateCRStatus accepts an AWSResource object and changes the Status
// sub-object of the AWSResource's Kubernetes custom resource (CR) and
// returns whether any changes were made
func (d *resourceDescriptor) UpdateCRStatus(
	res acktypes.AWSResource,
) (bool, error) {
	updated := true
	return updated, nil
}

// IsManaged returns true if the supplied AWSResource is under the management
// of an ACK service controller. What this means in practice is that the
// underlying custom resource (CR) in the AWSResource has had a
// resource-specific finalizer associated with it.
func (d *resourceDescriptor) IsManaged(
	res acktypes.AWSResource,
) bool {
	obj := res.RuntimeMetaObject()
	if obj == nil {
		// Should not happen. If it does, there is a bug in the code
		panic("nil RuntimeMetaObject in AWSResource")
	}
	// Remove use of custom code once
	// https://github.com/kubernetes-sigs/controller-runtime/issues/994 is
	// fixed. This should be able to be:
	//
	// return k8sctrlutil.ContainsFinalizer(obj, finalizerString)
	return containsFinalizer(obj, finalizerString)
}

// Remove once https://github.com/kubernetes-sigs/controller-runtime/issues/994
// is fixed.
func containsFinalizer(obj acktypes.RuntimeMetaObject, finalizer string) bool {
	f := obj.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return true
		}
	}
	return false
}

// MarkManaged places the supplied resource under the management of ACK.  What
// this typically means is that the resource manager will decorate the
// underlying custom resource (CR) with a finalizer that indicates ACK is
// managing the resource and the underlying CR may not be deleted until ACK is
// finished cleaning up any backend AWS service resources associated with the
// CR.
func (d *resourceDescriptor) MarkManaged(
	res acktypes.AWSResource,
) {
	obj := res.RuntimeMetaObject()
	if obj == nil {
		// Should not happen. If it does, there is a bug in the code
		panic("nil RuntimeMetaObject in AWSResource")
	}
	k8sctrlutil.AddFinalizer(obj, finalizerString)
}

// MarkUnmanaged removes the supplied resource from management by ACK.  What
// this typically means is that the resource manager will remove a finalizer
// underlying custom resource (CR) that indicates ACK is managing the resource.
// This will allow the Kubernetes API server to delete the underlying CR.
func (d *resourceDescriptor) MarkUnmanaged(
	res acktypes.AWSResource,
) {
	obj := res.RuntimeMetaObject()
	if obj == nil {
		// Should not happen. If it does, there is a bug in the code
		panic("nil RuntimeMetaObject in AWSResource")
	}
	k8sctrlutil.RemoveFinalizer(obj, finalizerString)
}

// MarkAdopted places descriptors on the custom resource that indicate the
// resource was not created from within ACK.
func (d *resourceDescriptor) MarkAdopted(
	res acktypes.AWSResource,
) {
	obj := res.RuntimeMetaObject()
	if obj == nil {
		// Should not happen. If it does, there is a bug in the code
		panic("nil RuntimeMetaObject in AWSResource")
	}
	curr := obj.GetAnnotations()
	if curr == nil {
		curr = make(map[string]string)
	}
	curr[ackv1alpha1.AnnotationAdopted] = "true"
	obj.SetAnnotations(curr)
}
