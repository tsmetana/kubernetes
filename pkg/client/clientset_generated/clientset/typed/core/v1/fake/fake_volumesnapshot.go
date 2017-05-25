/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "k8s.io/kubernetes/pkg/api/v1"
)

// FakeVolumeSnapshots implements VolumeSnapshotInterface
type FakeVolumeSnapshots struct {
	Fake *FakeCoreV1
	ns   string
}

var volumesnapshotsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "volumesnapshots"}

var volumesnapshotsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "VolumeSnapshot"}

func (c *FakeVolumeSnapshots) Create(volumeSnapshot *v1.VolumeSnapshot) (result *v1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) Update(volumeSnapshot *v1.VolumeSnapshot) (result *v1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) UpdateStatus(volumeSnapshot *v1.VolumeSnapshot) (*v1.VolumeSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotsResource, "status", c.ns, volumeSnapshot), &v1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) Delete(name string, options *meta_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesnapshotsResource, c.ns, name), &v1.VolumeSnapshot{})

	return err
}

func (c *FakeVolumeSnapshots) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1.VolumeSnapshotList{})
	return err
}

func (c *FakeVolumeSnapshots) Get(name string, options meta_v1.GetOptions) (result *v1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotsResource, c.ns, name), &v1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) List(opts meta_v1.ListOptions) (result *v1.VolumeSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotsResource, volumesnapshotsKind, c.ns, opts), &v1.VolumeSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VolumeSnapshotList{}
	for _, item := range obj.(*v1.VolumeSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshots.
func (c *FakeVolumeSnapshots) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched volumeSnapshot.
func (c *FakeVolumeSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotsResource, c.ns, name, data, subresources...), &v1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VolumeSnapshot), err
}
