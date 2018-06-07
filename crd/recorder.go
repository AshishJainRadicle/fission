/*
Copyright 2016 The Fission Authors.

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

package crd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	//"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type (
	RecorderInterface interface {
		Create(*Recorder) (*Recorder, error)
		Get(name string) (*Recorder, error)
		Update(*Recorder) (*Recorder, error)
		List(opts metav1.ListOptions) (*RecorderList, error)
		Watch(opts metav1.ListOptions) (watch.Interface, error)
	}

	recorderClient struct {
		client *rest.RESTClient
		namespace string
	}
)

/*
type (
	MessageQueueTriggerInterface interface {
		Create(*MessageQueueTrigger) (*MessageQueueTrigger, error)
		Get(name string) (*MessageQueueTrigger, error)
		Update(*MessageQueueTrigger) (*MessageQueueTrigger, error)
		Delete(name string, options *metav1.DeleteOptions) error
		List(opts metav1.ListOptions) (*MessageQueueTriggerList, error)
		Watch(opts metav1.ListOptions) (watch.Interface, error)
	}

	messageQueueTriggerClient struct {
		client    *rest.RESTClient
		namespace string
	}
)
*/

func MakeRecorderInterface(crdClient *rest.RESTClient, namespace string) RecorderInterface {
	return &recorderClient{
		client: crdClient,
		namespace: namespace,
	}
}

/*
func MakeMessageQueueTriggerInterface(crdClient *rest.RESTClient, namespace string) MessageQueueTriggerInterface {
	return &messageQueueTriggerClient{
		client:    crdClient,
		namespace: namespace,
	}
}
*/

func (rc *recorderClient) Create(f *Recorder) (*Recorder, error){
	return nil, nil
}

/*
func (fc *messageQueueTriggerClient) Create(f *MessageQueueTrigger) (*MessageQueueTrigger, error) {
	var result MessageQueueTrigger
	err := fc.client.Post().
		Resource("messagequeuetriggers").
		Namespace(fc.namespace).
		Body(f).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
*/

func (rc *recorderClient) Get(name string) (*Recorder, error) {
	return nil, nil
}

/*
func (fc *messageQueueTriggerClient) Get(name string) (*MessageQueueTrigger, error) {
	var result MessageQueueTrigger
	err := fc.client.Get().
		Resource("messagequeuetriggers").
		Namespace(fc.namespace).
		Name(name).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
*/

func (rc *recorderClient) Update(f *Recorder) (*Recorder, error) {
	return nil, nil
}

/*
func (fc *messageQueueTriggerClient) Update(f *MessageQueueTrigger) (*MessageQueueTrigger, error) {
	var result MessageQueueTrigger
	err := fc.client.Put().
		Resource("messagequeuetriggers").
		Namespace(fc.namespace).
		Name(f.Metadata.Name).
		Body(f).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
*/

func (rc *recorderClient) Delete(name, string, opts *metav1.DeleteOptions) error {
	return nil
}

/*
func (fc *messageQueueTriggerClient) Delete(name string, opts *metav1.DeleteOptions) error {
	return fc.client.Delete().
		Namespace(fc.namespace).
		Resource("messagequeuetriggers").
		Name(name).
		Body(opts).
		Do().
		Error()
}
*/

func (rc *recorderClient) List(opts metav1.ListOptions) (*RecorderList, error) {
	return nil, nil
}
/*
func (fc *messageQueueTriggerClient) List(opts metav1.ListOptions) (*MessageQueueTriggerList, error) {
	var result MessageQueueTriggerList
	err := fc.client.Get().
		Namespace(fc.namespace).
		Resource("messagequeuetriggers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
*/


func (rc *recorderClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return nil, nil
}

/*
func (fc *messageQueueTriggerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return fc.client.Get().
		Prefix("watch").
		Namespace(fc.namespace).
		Resource("messagequeuetriggers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
*/
