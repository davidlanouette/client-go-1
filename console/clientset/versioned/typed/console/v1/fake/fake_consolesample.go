// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/console/v1"
	consolev1 "github.com/openshift/client-go/console/applyconfigurations/console/v1"
	typedconsolev1 "github.com/openshift/client-go/console/clientset/versioned/typed/console/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeConsoleSamples implements ConsoleSampleInterface
type fakeConsoleSamples struct {
	*gentype.FakeClientWithListAndApply[*v1.ConsoleSample, *v1.ConsoleSampleList, *consolev1.ConsoleSampleApplyConfiguration]
	Fake *FakeConsoleV1
}

func newFakeConsoleSamples(fake *FakeConsoleV1) typedconsolev1.ConsoleSampleInterface {
	return &fakeConsoleSamples{
		gentype.NewFakeClientWithListAndApply[*v1.ConsoleSample, *v1.ConsoleSampleList, *consolev1.ConsoleSampleApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("consolesamples"),
			v1.SchemeGroupVersion.WithKind("ConsoleSample"),
			func() *v1.ConsoleSample { return &v1.ConsoleSample{} },
			func() *v1.ConsoleSampleList { return &v1.ConsoleSampleList{} },
			func(dst, src *v1.ConsoleSampleList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ConsoleSampleList) []*v1.ConsoleSample { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ConsoleSampleList, items []*v1.ConsoleSample) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
