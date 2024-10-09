package gofuzzheaders

import (
	"fmt"
	"strings"
	"os"
	"testing"
	corev1 "k8s.io/api/core/v1"
)

func TestGenerateSeed(t *testing.T) {
	pod := &corev1.Pod{}
	c := NewSeedGenerator()
	s := c.GenerateSeed(pod)
	ff := NewConsumer(s)
	ff.GenerateStruct(pod)
	//t.Log(err)
	t.Logf("%+v\n", pod.TypeMeta.Kind)
	t.Logf("%+v\n", pod.TypeMeta.APIVersion)
	t.Logf("%+v\n", pod.ObjectMeta.Name)
	t.Logf("%+v\n", pod.ObjectMeta.GenerateName)
	t.Logf("%+v\n", pod.ObjectMeta.Namespace)
	t.Logf("%+v\n", pod.ObjectMeta.SelfLink)
	t.Logf("%+v\n", pod.ObjectMeta.ResourceVersion)
	t.Logf("%+v\n", pod.ObjectMeta.Generation)
	t.Logf("%+v\n", *pod.ObjectMeta.DeletionGracePeriodSeconds)
	t.Logf("Labels: %+v\n", pod.ObjectMeta.Labels)
	t.Logf("%+v\n", pod.ObjectMeta.Annotations)
	t.Logf("%+v\n", pod.ObjectMeta.OwnerReferences)
	t.Logf("%+v\n", pod.ObjectMeta.Finalizers)
	t.Logf("OwnerReferences: %+v\n", pod.ObjectMeta.OwnerReferences)
	t.Logf("Containers: %+v\n", pod.Spec.Containers)
	//t.Logf("len(pod.Spec.Containers): %d", len(pod.Spec.Containers))
	var sw strings.Builder
	for _, b := range s {
		sw.WriteString(fmt.Sprintf("0x%X ", b))
	}
	//t.Log(s[0:2000])
	t.Log(len(sw.String()))

	// creates a seed file:
	f, err := os.Create("seed1")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(sw.String())
	t.Fatal("Done")
}

func TestPod(t *testing.T) {
	data := []byte{0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x1, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x2, 0x35, 0x61, 0x35, 0x61, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x1, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x2, 0x35, 0x61, 0x35, 0x61, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x1, 0x35, 0x1, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x2, 0x35, 0x61, 0x35, 0x61, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x61, 0x61, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x0, 0x0, 0x0, 0x2, 0x35, 0x61, 0x35, 0x61, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x61, 0x61, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x35, 0x0, 0x0, 0x0, 0x2, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x35, 0x1, 0x35, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x35, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x1, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1, 0x0, 0x0, 0x0, 0x3, 0x41, 0x42, 0x43, 0x1}
	f := NewConsumer(data)
	pod := &corev1.Pod{}
	f.GenerateStruct(pod)
	if pod.TypeMeta.Kind != "ABC" {
		t.Fatal("err")
	}
	if pod.TypeMeta.APIVersion != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.Name != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.GenerateName != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.Namespace != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.SelfLink != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.ResourceVersion != "ABC" {
		t.Fatal("err")
	}
	if pod.ObjectMeta.Generation != 53 {
		t.Fatal("err")
	}
	if *pod.ObjectMeta.DeletionGracePeriodSeconds != 0 {
		t.Fatal("err")		
	}

	t.Logf("%+v\n", *pod.ObjectMeta.DeletionGracePeriodSeconds)
	t.Logf("Labels: %+v\n", pod.ObjectMeta.Labels)
	t.Logf("%+v\n", pod.ObjectMeta.Annotations)
	t.Logf("%+v\n", pod.ObjectMeta.OwnerReferences)
	t.Logf("%+v\n", pod.ObjectMeta.Finalizers)
}