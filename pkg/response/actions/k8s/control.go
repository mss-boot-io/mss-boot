package k8s

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/k8s"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/5/23 17:53:08
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/5/23 17:53:08
 */

type Control struct {
	opts *Options
}

func (*Control) String() string {
	return "control"
}

func NewControl(opts ...Option) *Control {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return &Control{
		opts: o,
	}
}

func (e *Control) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		if e.opts.Model == nil {
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
			return
		}
		switch c.Request.Method {
		case http.MethodPost:
			e.create(c)
		case http.MethodPut:
			e.update(c)
		default:
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
		}
	}
	chain := gin.HandlersChain{h}
	if e.opts.controlHandlers != nil {
		chain = append(chain, e.opts.controlHandlers...)
	}
	if e.opts.Handlers != nil {
		chain = append(chain, e.opts.Handlers...)
	}
	return chain
}

func (e *Control) create(c *gin.Context) {
	namespace := c.Param("namespace")
	m := pkg.DeepCopy(e.opts.Model)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusBadRequest)
		return
	}
	if e.opts.BeforeCreate != nil {
		if err := e.opts.BeforeCreate(c, k8s.ClientSet, e.opts.Model); err != nil {
			api.AddError(err).Log.Error("BeforeCreate error")
			api.Err(http.StatusInternalServerError)
			return
		}
	}
	object, err := transferObjectToResource(e.opts.ResourceType, m)
	if err != nil {
		api.AddError(err).Log.Error("Transfer error")
		api.Err(http.StatusBadRequest)
		return
	}
	object, err = createResource(c, namespace, object)
	if err != nil {
		api.AddError(err).Log.Error("Create error")
		api.Err(http.StatusInternalServerError)
		return
	}
	if e.opts.AfterCreate != nil {
		err = e.opts.AfterCreate(c, k8s.ClientSet, e.opts.Model)
		if err != nil {
			api.AddError(err).Log.Error("AfterCreate error")
			api.Err(http.StatusInternalServerError)
			return
		}
	}
	api.OK(object)
}

func (e *Control) update(c *gin.Context) {
	namespace := c.Param("namespace")
	m := pkg.DeepCopy(e.opts.Model)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusBadRequest)
		return
	}
	if e.opts.BeforeUpdate != nil {
		if err := e.opts.BeforeUpdate(c, k8s.ClientSet, e.opts.Model); err != nil {
			api.AddError(err).Log.Error("BeforeUpdate error")
			api.Err(http.StatusInternalServerError)
			return
		}
	}
	object, err := transferObjectToResource(e.opts.ResourceType, m)
	if err != nil {
		api.AddError(err).Log.Error("Transfer error")
		api.Err(http.StatusBadRequest)
		return
	}
	object, err = updateResource(c, namespace, object)
	if err != nil {
		api.AddError(err).Log.Error("Update error")
		api.Err(http.StatusInternalServerError)
		return
	}
	if e.opts.AfterUpdate != nil {
		err = e.opts.AfterUpdate(c, k8s.ClientSet, e.opts.Model)
		if err != nil {
			api.AddError(err).Log.Error("AfterUpdate error")
			api.Err(http.StatusInternalServerError)
			return
		}
	}
	api.OK(object)
}

func transferObjectToResource(resourceType ResourceType, object any) (any, error) {
	rb, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	var result any
	// set result to object as ResourceType
	switch resourceType {
	case Deployment:
		result = &appsv1.Deployment{}
	case ConfigMap:
		result = &corev1.ConfigMap{}
	case Secret:
		result = &corev1.Secret{}
	case Service:
		result = &corev1.Service{}
	case Pod:
		result = &corev1.Pod{}
	case StatefulSet:
		result = &appsv1.StatefulSet{}
	case Job:
		result = &batchv1.Job{}
	case CronJob:
		result = &batchv1.CronJob{}
	case DaemonSet:
		result = &appsv1.DaemonSet{}
	case Ingress:
		result = &networkingv1.Ingress{}
	case ResourceQuota:
		result = &corev1.ResourceQuota{}
	case LimitRange:
		result = &corev1.LimitRange{}
	default:
		return nil, errors.New("not support resource type")
	}
	err = json.Unmarshal(rb, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateResource(ctx context.Context, namespace string, object any) (any, error) {
	switch object.(type) {
	case *appsv1.Deployment:
		return k8s.ClientSet.AppsV1().Deployments(namespace).
			Update(ctx, object.(*appsv1.Deployment), metav1.UpdateOptions{})
	case *corev1.ConfigMap:
		return k8s.ClientSet.CoreV1().ConfigMaps(namespace).
			Update(ctx, object.(*corev1.ConfigMap), metav1.UpdateOptions{})
	case *corev1.Secret:
		return k8s.ClientSet.CoreV1().Secrets(namespace).
			Update(ctx, object.(*corev1.Secret), metav1.UpdateOptions{})
	case *corev1.Service:
		return k8s.ClientSet.CoreV1().Services(namespace).
			Update(ctx, object.(*corev1.Service), metav1.UpdateOptions{})
	case *corev1.Pod:
		return k8s.ClientSet.CoreV1().Pods(namespace).
			Update(ctx, object.(*corev1.Pod), metav1.UpdateOptions{})
	case *appsv1.StatefulSet:
		return k8s.ClientSet.AppsV1().StatefulSets(namespace).
			Update(ctx, object.(*appsv1.StatefulSet), metav1.UpdateOptions{})
	case *batchv1.Job:
		return k8s.ClientSet.BatchV1().Jobs(namespace).
			Update(ctx, object.(*batchv1.Job), metav1.UpdateOptions{})
	case *batchv1.CronJob:
		return k8s.ClientSet.BatchV1().CronJobs(namespace).
			Update(ctx, object.(*batchv1.CronJob), metav1.UpdateOptions{})
	case *appsv1.DaemonSet:
		return k8s.ClientSet.AppsV1().DaemonSets(namespace).
			Update(ctx, object.(*appsv1.DaemonSet), metav1.UpdateOptions{})
	case *networkingv1.Ingress:
		return k8s.ClientSet.NetworkingV1().Ingresses(namespace).
			Update(ctx, object.(*networkingv1.Ingress), metav1.UpdateOptions{})
	case *corev1.ResourceQuota:
		return k8s.ClientSet.CoreV1().ResourceQuotas(namespace).
			Update(ctx, object.(*corev1.ResourceQuota), metav1.UpdateOptions{})
	case *corev1.LimitRange:
		return k8s.ClientSet.CoreV1().LimitRanges(namespace).
			Update(ctx, object.(*corev1.LimitRange), metav1.UpdateOptions{})
	case *corev1.PersistentVolume:
		return k8s.ClientSet.CoreV1().PersistentVolumes().
			Update(ctx, object.(*corev1.PersistentVolume), metav1.UpdateOptions{})
	case *corev1.PersistentVolumeClaim:
		return k8s.ClientSet.CoreV1().PersistentVolumeClaims(namespace).
			Update(ctx, object.(*corev1.PersistentVolumeClaim), metav1.UpdateOptions{})
	case *corev1.Namespace:
		return k8s.ClientSet.CoreV1().Namespaces().
			Update(ctx, object.(*corev1.Namespace), metav1.UpdateOptions{})
	case *storagev1.StorageClass:
		return k8s.ClientSet.StorageV1().StorageClasses().
			Update(ctx, object.(*storagev1.StorageClass), metav1.UpdateOptions{})
	case *networkingv1.IngressClass:
		return k8s.ClientSet.NetworkingV1().IngressClasses().
			Update(ctx, object.(*networkingv1.IngressClass), metav1.UpdateOptions{})
	}
	return nil, errors.New("not support resource type")
}

func createResource(ctx context.Context, namespace string, object any) (any, error) {
	switch object.(type) {
	case *appsv1.Deployment:
		return k8s.ClientSet.AppsV1().Deployments(namespace).Create(ctx, object.(*appsv1.Deployment), metav1.CreateOptions{})
	case *corev1.ConfigMap:
		return k8s.ClientSet.CoreV1().ConfigMaps(namespace).Create(ctx, object.(*corev1.ConfigMap), metav1.CreateOptions{})
	case *corev1.Secret:
		return k8s.ClientSet.CoreV1().Secrets(namespace).Create(ctx, object.(*corev1.Secret), metav1.CreateOptions{})
	case *corev1.Service:
		return k8s.ClientSet.CoreV1().Services(namespace).Create(ctx, object.(*corev1.Service), metav1.CreateOptions{})
	case *corev1.Pod:
		return k8s.ClientSet.CoreV1().Pods(namespace).Create(ctx, object.(*corev1.Pod), metav1.CreateOptions{})
	case *appsv1.StatefulSet:
		return k8s.ClientSet.AppsV1().StatefulSets(namespace).Create(ctx, object.(*appsv1.StatefulSet), metav1.CreateOptions{})
	case *batchv1.Job:
		return k8s.ClientSet.BatchV1().Jobs(namespace).Create(ctx, object.(*batchv1.Job), metav1.CreateOptions{})
	case *batchv1.CronJob:
		return k8s.ClientSet.BatchV1().CronJobs(namespace).Create(ctx, object.(*batchv1.CronJob), metav1.CreateOptions{})
	case *appsv1.DaemonSet:
		return k8s.ClientSet.AppsV1().DaemonSets(namespace).Create(ctx, object.(*appsv1.DaemonSet), metav1.CreateOptions{})
	case *networkingv1.Ingress:
		return k8s.ClientSet.NetworkingV1().Ingresses(namespace).Create(ctx, object.(*networkingv1.Ingress), metav1.CreateOptions{})
	case *corev1.ResourceQuota:
		return k8s.ClientSet.CoreV1().ResourceQuotas(namespace).Create(ctx, object.(*corev1.ResourceQuota), metav1.CreateOptions{})
	case *corev1.LimitRange:
		return k8s.ClientSet.CoreV1().LimitRanges(namespace).Create(ctx, object.(*corev1.LimitRange), metav1.CreateOptions{})
	case *corev1.PersistentVolume:
		return k8s.ClientSet.CoreV1().PersistentVolumes().Create(ctx, object.(*corev1.PersistentVolume), metav1.CreateOptions{})
	case *corev1.PersistentVolumeClaim:
		return k8s.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Create(ctx, object.(*corev1.PersistentVolumeClaim), metav1.CreateOptions{})
	case *corev1.Namespace:
		return k8s.ClientSet.CoreV1().Namespaces().Create(ctx, object.(*corev1.Namespace), metav1.CreateOptions{})
	case *storagev1.StorageClass:
		return k8s.ClientSet.StorageV1().StorageClasses().Create(ctx, object.(*storagev1.StorageClass), metav1.CreateOptions{})
	case *networkingv1.IngressClass:
		return k8s.ClientSet.NetworkingV1().IngressClasses().Create(ctx, object.(*networkingv1.IngressClass), metav1.CreateOptions{})
	}
	return nil, errors.New("not support resource type")
}
