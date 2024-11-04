package predicate

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const (
	labelSelectorKeyPauseReconciler = "cloud.pingcap.com/migration-pause-reconciler"
)

func SetupPredicate() (predicate.Predicate, error) {
	labelSelector := metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      labelSelectorKeyPauseReconciler,
				Operator: metav1.LabelSelectorOpDoesNotExist,
			},
		},
	}

	p, err := predicate.LabelSelectorPredicate(labelSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to create predicate: %v", err)
	}
	return p, nil
}
