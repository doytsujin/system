/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"github.com/knative/pkg/kmeta"
	servingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	projectriffv1alpha1 "github.com/projectriff/system/pkg/apis/projectriff/v1alpha1"
	"github.com/projectriff/system/pkg/reconciler/v1alpha1/application/resources/names"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeConfiguration creates a Configuration from a Application object.
func MakeConfiguration(application *projectriffv1alpha1.Application) (*servingv1alpha1.Configuration, error) {
	configuration := &servingv1alpha1.Configuration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.Configuration(application),
			Namespace: application.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(application),
			},
			Labels: makeLabels(application),
		},
		Spec: servingv1alpha1.ConfigurationSpec{
			RevisionTemplate: servingv1alpha1.RevisionTemplateSpec{
				Spec: servingv1alpha1.RevisionSpec{
					Container: corev1.Container{
						Image:     application.Spec.Image,
						EnvFrom:   application.Spec.Run.EnvFrom,
						Env:       application.Spec.Run.Env,
						Resources: application.Spec.Run.Resources,
					},
				},
			},
		},
	}

	return configuration, nil
}