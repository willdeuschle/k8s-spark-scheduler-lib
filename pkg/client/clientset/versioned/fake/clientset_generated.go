// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned"
	scalerv1alpha1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/scaler/v1alpha1"
	fakescalerv1alpha1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/scaler/v1alpha1/fake"
	scalerv1alpha2 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/scaler/v1alpha2"
	fakescalerv1alpha2 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/scaler/v1alpha2/fake"
	sparkschedulerv1beta1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/sparkscheduler/v1beta1"
	fakesparkschedulerv1beta1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/sparkscheduler/v1beta1/fake"
	sparkschedulerv1beta2 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/sparkscheduler/v1beta2"
	fakesparkschedulerv1beta2 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned/typed/sparkscheduler/v1beta2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ScalerV1alpha1 retrieves the ScalerV1alpha1Client
func (c *Clientset) ScalerV1alpha1() scalerv1alpha1.ScalerV1alpha1Interface {
	return &fakescalerv1alpha1.FakeScalerV1alpha1{Fake: &c.Fake}
}

// ScalerV1alpha2 retrieves the ScalerV1alpha2Client
func (c *Clientset) ScalerV1alpha2() scalerv1alpha2.ScalerV1alpha2Interface {
	return &fakescalerv1alpha2.FakeScalerV1alpha2{Fake: &c.Fake}
}

// SparkschedulerV1beta1 retrieves the SparkschedulerV1beta1Client
func (c *Clientset) SparkschedulerV1beta1() sparkschedulerv1beta1.SparkschedulerV1beta1Interface {
	return &fakesparkschedulerv1beta1.FakeSparkschedulerV1beta1{Fake: &c.Fake}
}

// SparkschedulerV1beta2 retrieves the SparkschedulerV1beta2Client
func (c *Clientset) SparkschedulerV1beta2() sparkschedulerv1beta2.SparkschedulerV1beta2Interface {
	return &fakesparkschedulerv1beta2.FakeSparkschedulerV1beta2{Fake: &c.Fake}
}
