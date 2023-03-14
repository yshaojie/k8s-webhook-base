package v1

import (
	"context"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type PodAnnotator struct {
	Client  client.Client
	decoder *admission.Decoder
}

var (
	log = ctrl.Log.WithName("webhook")
)

//+kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,sideEffects=NoneOnDryRun,admissionReviewVersions=v1,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=xy.meteor.io

func (a *PodAnnotator) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}
	err := a.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if req.Kind.Kind != "Pod" {
		return admission.Allowed("not a pod,skip")
	}

	switch req.Operation {
	case admissionv1.Create:
		handleCreate(pod)
	case admissionv1.Update:
		handleUpdate(pod)
	case admissionv1.Delete:
		handleDelete(pod)
	default:
		return admission.Allowed("skip")
	}
	//在 pod 中修改字段
	marshaledPod, err := json.Marshal(pod)
	log.Info(string(marshaledPod))
	log.Info("abc")
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

func handleDelete(pod *corev1.Pod) {

}

func handleUpdate(pod *corev1.Pod) {

}

func handleCreate(pod *corev1.Pod) {

}

func (a *PodAnnotator) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}
