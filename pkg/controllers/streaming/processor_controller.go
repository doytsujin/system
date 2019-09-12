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

package streaming

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	streamingv1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
)

// ProcessorReconciler reconciles a Processor object
type ProcessorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=streaming.projectriff.io,resources=processors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=streaming.projectriff.io,resources=processors/status,verbs=get;update;patch

func (r *ProcessorReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("processor", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *ProcessorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&streamingv1alpha1.Processor{}).
		Complete(r)
}