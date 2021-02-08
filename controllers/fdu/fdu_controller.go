/*
* Copyright (c) 2018,2021 ADLINK Technology Inc.
*
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache Software License 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
*
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors:
*   ADLINK fog05 team, <fog05@adlink-labs.tech>
 */

package fdu

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	fduapi "github.com/eclipse-fog05/fog05-go/pkg/apis/fim/v03alpha1/pb"
	fduv0alpha3 "github.com/eclipse-fog05/fog05-manager/apis/fdu/v0alpha3"
)

// FDUReconciler reconciles a FDU object
type FDUReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=fdu.fog05.io,resources=fdus,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=fdu.fog05.io,resources=fdus/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=fdu.fog05.io,resources=fdus/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=nodes,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FDU object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *FDUReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("fdu", req.NamespacedName)

	// Getting the FDU instance for this reconcile request
	fdu := &fduv0alpha3.FDU{}
	err := r.Client.Get(ctx, req.NamespacedName, fdu)
	log.Info(fmt.Sprintf("Found: %+v", fdu))
	if err != nil {
		if errors.IsNotFound(err) {
			// FDU was not found
			return ctrl.Result{}, nil
		}
		// Other error
		log.Error(err, fmt.Sprintf("Failed to get FDU for %+v", req))
		return ctrl.Result{}, err
	}

	res, err := fduapi.FDUAPI.OnboardFDU(fdu.Spec)
	if err != nil {
		return ctrl.Result{}, err
	}
	log.Info(fmt.Sprintf("Res: %+v", res))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FDUReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&fduv0alpha3.FDU{}).
		Complete(r)
}
