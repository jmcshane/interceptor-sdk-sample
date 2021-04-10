package interceptor

import (
	"context"

	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var _ triggersv1.InterceptorInterface = (*Interceptor)(nil)

// Interceptor implements a CEL based interceptor that uses CEL expressions
// against the incoming body and headers to match, if the expression returns
// a true value, then the interception is "successful".
type Interceptor struct {
	KubeClientSet kubernetes.Interface
	Logger        *zap.SugaredLogger
}

// NewInterceptor creates a prepopulated Interceptor.
func NewInterceptor(k kubernetes.Interface, l *zap.SugaredLogger) *Interceptor {
	return &Interceptor{
		Logger:        l,
		KubeClientSet: k,
	}
}

func (w *Interceptor) Process(ctx context.Context, r *triggersv1.InterceptorRequest) *triggersv1.InterceptorResponse {
	// TODO: Implement interceptor
	return &triggersv1.InterceptorResponse{
		Continue: true,
	}
}
