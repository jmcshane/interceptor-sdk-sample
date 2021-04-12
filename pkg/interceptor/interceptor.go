package interceptor

import (
	"context"

	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	triggersv1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var _ triggersv1.InterceptorInterface = (*Interceptor)(nil)

type Interceptor struct {
	KubeClientSet kubernetes.Interface
	Logger        *zap.SugaredLogger
}

// NewInterceptor creates a prepopulated Interceptor.
func NewInterceptor(k kubernetes.Interface, l *zap.SugaredLogger) v1alpha1.InterceptorInterface {
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
