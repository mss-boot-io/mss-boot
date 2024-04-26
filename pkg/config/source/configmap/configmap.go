package configmap

import (
	"context"
	"io/fs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/4/26 17:19:21
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/4/26 17:19:21
 */

// Source is a k8s configmap source
type Source struct {
	opt  *source.Options
	done chan struct{}
}

// Open a file for reading
func (s *Source) Open(name string) (fs.File, error) {
	return nil, nil
}

// ReadFile read file
func (s *Source) ReadFile(name string) (rb []byte, err error) {
	cm, err := s.opt.Clientset.CoreV1().
		ConfigMaps(s.opt.Namespace).
		Get(context.TODO(), s.opt.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	for i := range source.Extends {
		if v, ok := cm.Data[name+"."+string(source.Extends[i])]; ok {
			s.opt.Extend = source.Extends[i]
			return []byte(v), nil
		}
	}
	return nil, err
}

// Watch configmap
func (s *Source) Watch(e source.Entity, unm func([]byte, any) error) error {
	w, err := s.opt.Clientset.CoreV1().ConfigMaps(s.opt.Namespace).Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: "metadata.name=" + s.opt.Name,
	})
	if err != nil {
		return err
	}
	defer w.Stop()
	if s.done == nil {
		s.done = make(chan struct{})
	}
	for {
		select {
		case <-s.done:
			return nil
		case event := <-w.ResultChan():
			if event.Type == watch.Modified {
				cm, ok := event.Object.(*corev1.ConfigMap)
				if !ok {
					continue
				}
				for i := range source.Extends {
					if v, ok := cm.Data[s.opt.Name+"."+string(source.Extends[i])]; ok {
						s.opt.Extend = source.Extends[i]
						if err = unm([]byte(v), e); err != nil {
							return err
						}
					}
				}
			}
		}
	}
}
