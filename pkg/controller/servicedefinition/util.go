package servicedefinition

import (
	"context"
	servicedefinitionv1alpha1 "github.com/AESthetix256/servicedefinition/pkg/apis/servicedefinition/v1alpha1"
	"github.com/AESthetix256/servicedefinition/pkg/controller/servicedefinition/serviceContainerDefinitions"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"k8s.io/apimachinery/pkg/api/errors"
)

func (r *ReconcileServicedefinition) CreateIfNotExists(reqLogger logr.Logger, pod *corev1.Pod, service *corev1.Service, ingress *networkingv1.Ingress) (reconcile.Result, error) {

	foundPod := &corev1.Pod{}
	foundSvc := &corev1.Service{}
	foundIngress := &networkingv1.Ingress{}


	err := r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, foundPod)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
		reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", foundPod.Namespace, "Pod.Name", foundPod.Name)
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundSvc)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Pod.Namespace", service.Namespace, "Pod.Name", service.Name)
		err = r.client.Create(context.TODO(), service)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
		reqLogger.Info("Skip reconcile: Service already exists", "Pod.Namespace", foundSvc.Namespace, "Pod.Name", foundSvc.Name)
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, foundIngress)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Ingress", "Pod.Namespace", ingress.Namespace, "Pod.Name", ingress.Name)
		err = r.client.Create(context.TODO(), ingress)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
		reqLogger.Info("Skip reconcile: Ingress already exists", "Pod.Namespace", foundIngress.Namespace, "Pod.Name", foundIngress.Name)
	}



	return reconcile.Result{}, nil
}

func ValidateService (reqLogger logr.Logger, service *servicedefinitionv1alpha1.Servicedefinition) (bool, error) {

	return true, nil

	/*reqLogger.Info("Checking if the input is correct")

	reqLogger.V(10).Info("Checking if all are empty")
	//check if all empty
	if (service.Spec.CMS.Name == "" && service.Spec.Database.Name == "" && service.Spec.Service.Name == "" && service.Spec.Webserver.Name == "") {
		// false
	}

	//check if CMS and FTP is not empty
	if (service.Spec.CMS.Name != "" && service.Spec.Service.Name != "" ) {
		// false
	}


	// check if Webserver or Database and CMS or FTP or all is not empty
	if ((service.Spec.Webserver.Name != "" || service.Spec.Database.Name != "") && (service.Spec.CMS.Name != "" || service.Spec.Service.Name != "")) {
		// false
	}

	//check if owner is not empty
	if (service.Spec.Owner == "") {
		// false
	}*/

}

func ValidateServiceInput (reqLogger logr.Logger, service *servicedefinitionv1alpha1.Servicedefinition) (bool, error) {

	return true, nil
}

func ReturnPodForService (reqLogger logr.Logger, service *servicedefinitionv1alpha1.Servicedefinition) (*corev1.Pod, *corev1.Service, *networkingv1.Ingress, error) {

	valid, _ := ValidateServiceInput(reqLogger, service)
	if(!valid) {
		return nil, nil, nil, errors.NewBadRequest("Service already exists")
	}

	if service.Spec.CMS.Name == "Typo3" {
		reqLogger.V(3).Info("Namespace of service: " + service.Namespace)

		return 	serviceContainerDefinitions.NewTypo3PodForCR(service),
		      	serviceContainerDefinitions.NewTypo3ServiceForCR(service),
		      	serviceContainerDefinitions.NewTypo3IngressForCR(service),
				nil
	}

	if service.Spec.Service.Name == "FTP" {
		return nil, nil, nil, nil
	}

	return nil, nil, nil, errors.NewBadRequest("Service not found")
}



