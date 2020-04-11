/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

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

package controllers

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1alpha1 "github.com/oneinfra/oneinfra/apis/cluster/v1alpha1"
)

// ClusterReconciler reconciles a Cluster object
type ClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cluster.oneinfra.ereslibre.es,resources=components,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cluster.oneinfra.ereslibre.es,resources=components/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cluster.oneinfra.ereslibre.es,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cluster.oneinfra.ereslibre.es,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=infra.oneinfra.ereslibre.es,resources=hypervisors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infra.oneinfra.ereslibre.es,resources=hypervisors/status,verbs=get;update;patch

// Reconcile reconciles the cluster resources
func (r *ClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	// TODO: cluster-wide changes will be reconciled here
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the cluster reconciler with mgr manager
func (r *ClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Named("cluster-controller").
		For(&clusterv1alpha1.Cluster{}).
		Complete(r)
}
