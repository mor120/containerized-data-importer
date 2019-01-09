/*
Copyright 2018 The CDI Authors.

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

package cluster

import (
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
)

const (
	apiServerResourceName = "cdi-apiserver"
)

func createAPIServerResources(args *FactoryArgs) []Resource {
	return []Resource{
		createAPIServerClusterRoleBinding(args.Namespace),
		createAPIServerClusterRole(),
		createAPIServerAuthClusterRoleBinding(args.Namespace),
	}
}

func createAPIServerClusterRoleBinding(namespace string) *rbacv1beta1.ClusterRoleBinding {
	return createClusterRoleBinding(apiServerResourceName, apiServerResourceName, apiServerResourceName, namespace)
}

func createAPIServerClusterRole() *rbacv1beta1.ClusterRole {
	clusterRole := createClusterRole(apiServerResourceName)
	clusterRole.Rules = []rbacv1beta1.PolicyRule{
		{
			APIGroups: []string{
				"admissionregistration.k8s.io",
			},
			Resources: []string{
				"validatingwebhookconfigurations",
			},
			Verbs: []string{
				"get",
				"create",
				"update",
			},
		},
		{
			APIGroups: []string{
				"apiregistration.k8s.io",
			},
			Resources: []string{
				"apiservices",
			},
			Verbs: []string{
				"get",
				"create",
				"update",
			},
		},
		{
			APIGroups: []string{
				"",
			},
			Resources: []string{
				"pods",
				"persistentvolumeclaims",
			},
			Verbs: []string{
				"get",
				"list",
			},
		},
	}
	return clusterRole
}

func createAPIServerAuthClusterRoleBinding(namespace string) *rbacv1beta1.ClusterRoleBinding {
	return createClusterRoleBinding("cdi-apiserver-auth-delegator", "system:auth-delegator", apiServerResourceName, namespace)
}
