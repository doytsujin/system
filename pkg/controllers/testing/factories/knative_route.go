/*
Copyright 2019 the original author or authors.

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

package factories

import (
	"fmt"

	"github.com/projectriff/system/pkg/apis"
	knativeservingv1 "github.com/projectriff/system/pkg/apis/thirdparty/knative/serving/v1"
)

type knativeRoute struct {
	target *knativeservingv1.Route
}

func KnativeRoute(seed ...*knativeservingv1.Route) *knativeRoute {
	var target *knativeservingv1.Route
	switch len(seed) {
	case 0:
		target = &knativeservingv1.Route{}
	case 1:
		target = seed[0]
	default:
		panic(fmt.Errorf("expected exactly zero or one seed, got %v", seed))
	}
	return &knativeRoute{
		target: target,
	}
}

func (f *knativeRoute) deepCopy() *knativeRoute {
	return KnativeRoute(f.target.DeepCopy())
}

func (f *knativeRoute) Get() *knativeservingv1.Route {
	return f.deepCopy().target
}

func (f *knativeRoute) Mutate(m func(*knativeservingv1.Route)) *knativeRoute {
	f = f.deepCopy()
	m(f.target)
	return f
}

func (f *knativeRoute) NamespaceName(namespace, name string) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		route.ObjectMeta.Namespace = namespace
		route.ObjectMeta.Name = name
	})
}

func (f *knativeRoute) ObjectMeta(nf func(ObjectMeta)) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		omf := objectMeta(route.ObjectMeta)
		nf(omf)
		route.ObjectMeta = omf.Get()
	})
}

func (f *knativeRoute) StatusConditions(conditions ...apis.Condition) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		route.Status.Conditions = conditions
	})
}

func (f *knativeRoute) StatusObservedGeneration(generation int64) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		route.Status.ObservedGeneration = generation
	})
}

func (f *knativeRoute) StatusAddressURL(url string) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		route.Status.Address = &apis.Addressable{
			URL: url,
		}
	})
}

func (f *knativeRoute) StatusURL(url string) *knativeRoute {
	return f.Mutate(func(route *knativeservingv1.Route) {
		route.Status.URL = url
	})
}