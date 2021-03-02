package cdk8splus17

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.AddDirectoryOptions",
		reflect.TypeOf((*AddDirectoryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.CommandProbeOptions",
		reflect.TypeOf((*CommandProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.ConfigMap",
		reflect.TypeOf((*ConfigMap)(nil)).Elem(),
		func() interface{} {
			c := configMap{}
			_jsii_.InitJsiiProxy(&c.resource)
			_jsii_.InitJsiiProxy(&c.iConfigMap)
			return &c
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ConfigMapProps",
		reflect.TypeOf((*ConfigMapProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ConfigMapVolumeOptions",
		reflect.TypeOf((*ConfigMapVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		func() interface{} {
			return &container{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ContainerProps",
		reflect.TypeOf((*ContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Deployment",
		reflect.TypeOf((*Deployment)(nil)).Elem(),
		func() interface{} {
			d := deployment{}
			_jsii_.InitJsiiProxy(&d.resource)
			_jsii_.InitJsiiProxy(&d.iPodTemplate)
			return &d
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.DeploymentProps",
		reflect.TypeOf((*DeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.EmptyDirMedium",
		reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": EmptyDirMedium_DEFAULT,
			"MEMORY": EmptyDirMedium_MEMORY,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.EmptyDirVolumeOptions",
		reflect.TypeOf((*EmptyDirVolumeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.EnvValue",
		reflect.TypeOf((*EnvValue)(nil)).Elem(),
		func() interface{} {
			return &envValue{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.EnvValueFromConfigMapOptions",
		reflect.TypeOf((*EnvValueFromConfigMapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.EnvValueFromProcessOptions",
		reflect.TypeOf((*EnvValueFromProcessOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.EnvValueFromSecretOptions",
		reflect.TypeOf((*EnvValueFromSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ExposeOptions",
		reflect.TypeOf((*ExposeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.HttpGetProbeOptions",
		reflect.TypeOf((*HttpGetProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.IConfigMap",
		reflect.TypeOf((*IConfigMap)(nil)).Elem(),
		func() interface{} {
			i := iConfigMap{}
			_jsii_.InitJsiiProxy(&i.iResource)
			return &i
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.IPodSpec",
		reflect.TypeOf((*IPodSpec)(nil)).Elem(),
		func() interface{} {
			return &iPodSpec{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.IPodTemplate",
		reflect.TypeOf((*IPodTemplate)(nil)).Elem(),
		func() interface{} {
			i := iPodTemplate{}
			_jsii_.InitJsiiProxy(&i.iPodSpec)
			return &i
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		func() interface{} {
			return &iResource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.ISecret",
		reflect.TypeOf((*ISecret)(nil)).Elem(),
		func() interface{} {
			i := iSecret{}
			_jsii_.InitJsiiProxy(&i.iResource)
			return &i
		},
	)
	_jsii_.RegisterInterface(
		"cdk8s-plus-17.IServiceAccount",
		reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		func() interface{} {
			i := iServiceAccount{}
			_jsii_.InitJsiiProxy(&i.iResource)
			return &i
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.ImagePullPolicy",
		reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": ImagePullPolicy_ALWAYS,
			"IF_NOT_PRESENT": ImagePullPolicy_IF_NOT_PRESENT,
			"NEVER": ImagePullPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.IngressV1Beta1",
		reflect.TypeOf((*IngressV1Beta1)(nil)).Elem(),
		func() interface{} {
			i := ingressV1Beta1{}
			_jsii_.InitJsiiProxy(&i.resource)
			return &i
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.IngressV1Beta1Backend",
		reflect.TypeOf((*IngressV1Beta1Backend)(nil)).Elem(),
		func() interface{} {
			return &ingressV1Beta1Backend{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.IngressV1Beta1Props",
		reflect.TypeOf((*IngressV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.IngressV1Beta1Rule",
		reflect.TypeOf((*IngressV1Beta1Rule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Job",
		reflect.TypeOf((*Job)(nil)).Elem(),
		func() interface{} {
			j := job{}
			_jsii_.InitJsiiProxy(&j.resource)
			_jsii_.InitJsiiProxy(&j.iPodTemplate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.JobProps",
		reflect.TypeOf((*JobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.MountOptions",
		reflect.TypeOf((*MountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.MountPropagation",
		reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		map[string]interface{}{
			"NONE": MountPropagation_NONE,
			"HOST_TO_CONTAINER": MountPropagation_HOST_TO_CONTAINER,
			"BIDIRECTIONAL": MountPropagation_BIDIRECTIONAL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.PathMapping",
		reflect.TypeOf((*PathMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Pod",
		reflect.TypeOf((*Pod)(nil)).Elem(),
		func() interface{} {
			p := pod{}
			_jsii_.InitJsiiProxy(&p.resource)
			_jsii_.InitJsiiProxy(&p.iPodSpec)
			return &p
		},
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.PodManagementPolicy",
		reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ORDERED_READY": PodManagementPolicy_ORDERED_READY,
			"PARALLEL": PodManagementPolicy_PARALLEL,
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.PodProps",
		reflect.TypeOf((*PodProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.PodSpec",
		reflect.TypeOf((*PodSpec)(nil)).Elem(),
		func() interface{} {
			p := podSpec{}
			_jsii_.InitJsiiProxy(&p.iPodSpec)
			return &p
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.PodSpecProps",
		reflect.TypeOf((*PodSpecProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.PodTemplate",
		reflect.TypeOf((*PodTemplate)(nil)).Elem(),
		func() interface{} {
			p := podTemplate{}
			_jsii_.InitJsiiProxy(&p.podSpec)
			_jsii_.InitJsiiProxy(&p.iPodTemplate)
			return &p
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.PodTemplateProps",
		reflect.TypeOf((*PodTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Probe",
		reflect.TypeOf((*Probe)(nil)).Elem(),
		func() interface{} {
			return &probe{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ProbeOptions",
		reflect.TypeOf((*ProbeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"TCP": Protocol_TCP,
			"UDP": Protocol_UDP,
			"SCTP": Protocol_SCTP,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		func() interface{} {
			r := resource{}
			_jsii_.InitJsiiProxy(&r.Construct)
			_jsii_.InitJsiiProxy(&r.iResource)
			return &r
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.RestartPolicy",
		reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ALWAYS": RestartPolicy_ALWAYS,
			"ON_FAILURE": RestartPolicy_ON_FAILURE,
			"NEVER": RestartPolicy_NEVER,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Secret",
		reflect.TypeOf((*Secret)(nil)).Elem(),
		func() interface{} {
			s := secret{}
			_jsii_.InitJsiiProxy(&s.resource)
			_jsii_.InitJsiiProxy(&s.iSecret)
			return &s
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.SecretProps",
		reflect.TypeOf((*SecretProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.SecretValue",
		reflect.TypeOf((*SecretValue)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		func() interface{} {
			s := service{}
			_jsii_.InitJsiiProxy(&s.resource)
			return &s
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		func() interface{} {
			s := serviceAccount{}
			_jsii_.InitJsiiProxy(&s.resource)
			_jsii_.InitJsiiProxy(&s.iServiceAccount)
			return &s
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ServiceIngressV1BetaBackendOptions",
		reflect.TypeOf((*ServiceIngressV1BetaBackendOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ServicePort",
		reflect.TypeOf((*ServicePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ServicePortOptions",
		reflect.TypeOf((*ServicePortOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s-plus-17.ServiceType",
		reflect.TypeOf((*ServiceType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_IP": ServiceType_CLUSTER_IP,
			"NODE_PORT": ServiceType_NODE_PORT,
			"LOAD_BALANCER": ServiceType_LOAD_BALANCER,
			"EXTERNAL_NAME": ServiceType_EXTERNAL_NAME,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.StatefulSet",
		reflect.TypeOf((*StatefulSet)(nil)).Elem(),
		func() interface{} {
			s := statefulSet{}
			_jsii_.InitJsiiProxy(&s.resource)
			_jsii_.InitJsiiProxy(&s.iPodTemplate)
			return &s
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.StatefulSetProps",
		reflect.TypeOf((*StatefulSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-plus-17.Volume",
		reflect.TypeOf((*Volume)(nil)).Elem(),
		func() interface{} {
			return &volume{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-plus-17.VolumeMount",
		reflect.TypeOf((*VolumeMount)(nil)).Elem(),
	)
}
