package appsclient

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

//func int32Ptr(int322 int32) *int32 {
//	return &int322
//}

func initiate() *kubernetes.Clientset {
	kubeconfigPath := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)

	if err != nil {
		log.Fatalf("Error building config file")
	}

	clientset := kubernetes.NewForConfigOrDie(config)

	return clientset
}

func CreateDeployment(name string, replicas int32) {

	log.Println("Creating deployment of AppsCodeServer...\n")

	kubeconfig := initiate()

	varAppsV1 := kubeconfig.AppsV1()

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"api": "latest",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"api": "latest",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
					Labels: map[string]string{
						"api": "latest",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            name,
							Image:           "masudjuly02/appscodeserver",
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "apps-port",
									ContainerPort: 8080,
									Protocol:      "TCP",
								},
							},
						},
					},
					RestartPolicy: "Always",
				},
			},
		},
	}

	_, err := varAppsV1.Deployments("default").Create(deployment)

	if err != nil {
		panic(err)
	}

	log.Printf("Deployment `%s` created successfully...!\n", name)
}

func CreateService(name string) {
	log.Printf("Creating service `%s` ...\n", name)

	kubeconfig := initiate()
	varCoreV1 := kubeconfig.CoreV1()

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"api": "latest",
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"api": "latest",
			},
			Ports: []corev1.ServicePort{
				{
					Port: 8080,
					TargetPort: intstr.FromInt(8080),
					Protocol: "TCP",
				},
			},
			Type: "NodePort",
		},
	}

	_, err := varCoreV1.Services("default").Create(service)

	if err != nil {
		panic(err)
	}

	log.Printf("Created service `%s` successfully\n", name)
}

func UpdateDeployment(name string, replicas int32) {
	log.Printf("Scaling deployment `%s` to %d replicas\n", name, replicas)

	kubeconfig := initiate()

	varAppsV1 := kubeconfig.AppsV1()

	deploy, err := varAppsV1.Deployments("default").Get(name, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}

	deploy.Spec.Replicas = &replicas

	_, err = varAppsV1.Deployments("default").Update(deploy)

	if err != nil {
		panic(err)
	}

	log.Println("Scaling completed successfully")
}

func IngressService(host, name string) {
	log.Printf("Creating Ingress of service `%s...`\n", name)

	kubeconfig := initiate()
	varExtensionV1Beta1 := kubeconfig.ExtensionsV1beta1()

	ingress := &extv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ingress-"+name,
		},
		Spec: extv1beta1.IngressSpec{
			Rules: []extv1beta1.IngressRule{
				{
					Host: host,
					IngressRuleValue: extv1beta1.IngressRuleValue{
						HTTP: &extv1beta1.HTTPIngressRuleValue{
							Paths: []extv1beta1.HTTPIngressPath{
								{
									Path: "/",
									Backend: extv1beta1.IngressBackend{
										ServiceName: name,
										ServicePort: intstr.FromInt(8080),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := varExtensionV1Beta1.Ingresses("default").Create(ingress)

	if err != nil {
		panic(err)
	}

	log.Printf("Created Ingress of service `%s` successfully\n", name)
}

func DeleteDeployment(name string){
	log.Println("Deleting everything related to this Deployment...")

	kubeconfig := initiate()
	varAppsV1 := kubeconfig.AppsV1()
	varCoreV1 := kubeconfig.CoreV1()
	varExtensionV1Beta1 := kubeconfig.ExtensionsV1beta1()

	if err := varAppsV1.Deployments("default").Delete(name, &metav1.DeleteOptions{}); err != nil {
		panic(err)
	}

	if err := varCoreV1.Services("default").Delete(name, &metav1.DeleteOptions{}); err != nil {
		panic(err)
	}

	if err := varExtensionV1Beta1.Ingresses("default").Delete("ingress-"+name, &metav1.DeleteOptions{}); err != nil {
		panic(err)
	}

	log.Printf("Deleted deployment `%s` successfully\n", name)
}