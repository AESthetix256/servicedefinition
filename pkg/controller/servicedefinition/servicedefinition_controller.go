package servicedefinition

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	servicedefinitionv1alpha1 "servicedefinition/pkg/apis/servicedefinition/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_servicedefinition")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Servicedefinition Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileServicedefinition{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("servicedefinition-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Servicedefinition
	err = c.Watch(&source.Kind{Type: &servicedefinitionv1alpha1.Servicedefinition{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Servicedefinition
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &servicedefinitionv1alpha1.Servicedefinition{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileServicedefinition implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileServicedefinition{}

// ReconcileServicedefinition reconciles a Servicedefinition object
type ReconcileServicedefinition struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Servicedefinition object and makes changes based on the state read
// and what is in the Servicedefinition.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileServicedefinition) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Servicedefinition")

	// Fetch the Servicedefinition instance
	instance := &servicedefinitionv1alpha1.Servicedefinition{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Überprüfen ob instanz passt
	valid, err := ValidateService(reqLogger,instance)
	if valid {

		// Identify Service

		// Define a new Pod object
		pod, service, ingress, _ := ReturnPodForService(reqLogger, instance)

		r.CreateIfNotExists(reqLogger, pod, service, ingress)

		return reconcile.Result{}, nil
	}

	// nicht ok, nicht erstellen und löschen, error messages
	return reconcile.Result{}, err
}

// TODO: Set Servicedefinition instance as the owner and controller
/*if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
	return reconcile.Result{}, err
}*/