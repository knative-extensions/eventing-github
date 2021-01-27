/*
Copyright 2021 The Knative Authors

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

package common

import (
	"context"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// SecretFrom gets the value of the Secret key referenced by sel from the
// Kubernetes cluster.
func SecretFrom(ctx context.Context, secretCli clientcorev1.SecretInterface, sel *corev1.SecretKeySelector) (string, error) {
	if sel == nil {
		return "", errors.New("missing Secret key selector")
	}

	secret, err := secretCli.Get(ctx, sel.Name, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("getting Secret from cluster: %w", err)
	}

	secretVal, ok := secret.Data[sel.Key]
	if !ok {
		return "", fmt.Errorf("key %q not found in Secret %q", sel.Key, sel.Name)
	}

	return string(secretVal), nil
}
