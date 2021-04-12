package main

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"knative.dev/pkg/signals"

	"github.com/jmcshane/interceptor-sdk/pkg/interceptor"
	"github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"github.com/tektoncd/triggers/pkg/interceptors/sdk"
)

func main() {
	ctx := signals.NewContext()

	sdk.InterceptorMainWithConfig(ctx, "interceptors", map[string]func(kubernetes.Interface, *zap.SugaredLogger) v1alpha1.InterceptorInterface{
		"example": interceptor.NewInterceptor,
	})
}
