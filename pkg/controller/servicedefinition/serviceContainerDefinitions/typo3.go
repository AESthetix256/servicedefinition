package serviceContainerDefinitions

import (
	servicedefinitionv1alpha1 "github.com/AESthetix256/servicedefinition/pkg/apis/servicedefinition/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewTypo3IngressForCR(cr *servicedefinitionv1alpha1.Servicedefinition) *networkingv1.Ingress {
	labels := map[string]string{
		"app": cr.Name,
	}

	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-ingress",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: networkingv1.IngressSpec{
			Rules:   []networkingv1.IngressRule{
				{
					Host:             cr.Spec.Owner + ".diplo.link",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path: "/",
									Backend: networkingv1.IngressBackend{
										ServiceName: cr.Name + "-svc-np",
										ServicePort: intstr.IntOrString{
											Type: intstr.Int,
											IntVal:80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func NewTypo3ServiceForCR(cr *servicedefinitionv1alpha1.Servicedefinition) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name,
	}

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-svc-np",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec:       corev1.ServiceSpec{
			Selector: map[string]string {
				"app": cr.Name,
			},
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort {
				{
					Port: 80,
					Protocol: corev1.ProtocolTCP,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
		},
	}
}

func NewTypo3PodForCR(cr *servicedefinitionv1alpha1.Servicedefinition) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "typo3-web",
					Image:   "martinhelmich/typo3:9.5",
					Ports: []corev1.ContainerPort{
						{
							Name:          "http",
							ContainerPort: 80,
							Protocol:      "TCP",
							HostIP:        "",
						},
					},

				},
				{
					Name: 		"typo3-db",
					Image: 		"mariadb:latest",
					Args: []string{"--character-set-server=utf8", "--collation-server=utf8_unicode_ci"},
					// Command: 	[]string{"mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci"},
					Env: []corev1.EnvVar{
						{
							Name:	"MYSQL_ROOT_PASSWORD",
							Value: 	"password",
						},
						{
							Name:	"MYSQL_USER",
							Value: 	"user",
						},
						{
							Name: 	"MYSQL_PASSWORD",
							Value: 	"password",
						},
						{
							Name: 	"MYSQL_DATABASE",
							Value: 	"database",
						},
					},
				},
			},
		},
	}
}
