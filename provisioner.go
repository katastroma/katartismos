package katartismos

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Provisioner applies a resource inventory to a cluster.
type Provisioner interface {
	Apply(ctx context.Context, resources []unstructured.Unstructured) error
}
