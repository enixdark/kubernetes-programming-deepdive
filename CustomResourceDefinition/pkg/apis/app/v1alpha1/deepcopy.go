package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *SslConfig) DeepCopyInto(out *SslConfig) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = SslConfigSpec{
		Cert: in.Spec.Cert,
		Key: in.Spec.Key,
		Domain: in.Spec.Domain,
	}
	out.Status = SslConfigStatus{
		State: in.Status.State,
		Message: in.Status.Message,
	}
}

func (in *SslConfig) DeepCopyObject() runtime.Object {
	out := SslConfig{}
	in.DeepCopyInto(&out)
	return &out
}

func (in* SslConfigList) DeepCopyObject() runtime.Object {
	out := SslConfigList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]SslConfig, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}