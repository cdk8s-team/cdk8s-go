// High level abstractions on top of cdk8s
package cdk8splus17

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/awslabs/cdk8s-go/cdk8splus17/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/awslabs/cdk8s-go/cdk8s"
)

// Options for `configmap.addDirectory()`.
// Experimental.
type AddDirectoryOptions struct {
	// Glob patterns to exclude when adding files.
	// Experimental.
	Exclude []string `json:"exclude"`
	// A prefix to add to all keys in the config map.
	// Experimental.
	KeyPrefix string `json:"keyPrefix"`
}

// Options for `Probe.fromCommand()`.
// Experimental.
type CommandProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	// Experimental.
	FailureThreshold float64 `json:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	// Experimental.
	SuccessThreshold float64 `json:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds"`
}

// ToProbeOptions is a convenience function to obtain a new ProbeOptions from this CommandProbeOptions.
func (c *CommandProbeOptions) ToProbeOptions() ProbeOptions {
	return ProbeOptions {
		FailureThreshold: c.FailureThreshold,
		InitialDelaySeconds: c.InitialDelaySeconds,
		PeriodSeconds: c.PeriodSeconds,
		SuccessThreshold: c.SuccessThreshold,
		TimeoutSeconds: c.TimeoutSeconds,
	}
}

// ConfigMap holds configuration data for pods to consume.
// Experimental.
type ConfigMap interface {
	Resource
	IConfigMap
	ApiObject() cdk8s.ApiObject
	BinaryData() map[string]string
	Data() map[string]string
	AddBinaryData(key string, value string)
	AddData(key string, value string)
	AddDirectory(localDir string, options AddDirectoryOptions)
	AddFile(localFile string, key string)
}

// The jsii proxy struct for ConfigMap
type configMap struct {
	resource // extends cdk8s-plus-17.Resource
	iConfigMap // implements cdk8s-plus-17.IConfigMap
}

func (c *configMap) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		c,
		"apiObject",
		&returns,
	)
	return returns
}

func (c *configMap) BinaryData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"binaryData",
		&returns,
	)
	return returns
}

func (c *configMap) Data() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"data",
		&returns,
	)
	return returns
}

func (c *configMap) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		c,
		"metadata",
		&returns,
	)
	return returns
}

func (c *configMap) Name() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
	)
	return returns
}


func NewConfigMap(scope constructs.Construct, id string, props ConfigMapProps) ConfigMap {
	_init_.Initialize()
	c := configMap{}

	_jsii_.Create(
		"cdk8s-plus-17.ConfigMap",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IConfigMap"},
		[]_jsii_.Override{},
		&c,
	)
	return &c
}

// Represents a ConfigMap created elsewhere.
// Experimental.
func ConfigMap_FromConfigMapName(name string) IConfigMap {
	_init_.Initialize()
	var returns IConfigMap
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.ConfigMap",
		"fromConfigMapName",
		[]interface{}{name},
		true,
		&returns,
	)
	return returns
}

// Adds a binary data entry to the config map.
//
// BinaryData can contain byte
// sequences that are not in the UTF-8 range.
// Experimental.
func (c *configMap) AddBinaryData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBinaryData",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Adds a data entry to the config map.
// Experimental.
func (c *configMap) AddData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addData",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Adds a directory to the ConfigMap.
// Experimental.
func (c *configMap) AddDirectory(localDir string, options AddDirectoryOptions) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDirectory",
		[]interface{}{localDir, options},
		false,
		&returns,
	)
}

// Adds a file to the ConfigMap.
// Experimental.
func (c *configMap) AddFile(localFile string, key string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addFile",
		[]interface{}{localFile, key},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *configMap) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *configMap) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (c *configMap) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		c,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *configMap) ToString() string {
	var returns string
	_jsii_.Invoke(
		c,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Properties for initialization of `ConfigMap`.
// Experimental.
type ConfigMapProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// BinaryData contains the binary data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range. The
	// keys stored in BinaryData must not overlap with the ones in the Data field,
	// this is enforced during validation process. Using this field will require
	// 1.10+ apiserver and kubelet.
	//
	// You can also add binary data using `configMap.addBinaryData()`.
	// Experimental.
	BinaryData map[string]string `json:"binaryData"`
	// Data contains the configuration data.
	//
	// Each key must consist of alphanumeric characters, '-', '_' or '.'. Values
	// with non-UTF-8 byte sequences must use the BinaryData field. The keys
	// stored in Data must not overlap with the keys in the BinaryData field, this
	// is enforced during validation process.
	//
	// You can also add data using `configMap.addData()`.
	// Experimental.
	Data map[string]string `json:"data"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this ConfigMapProps.
func (c *ConfigMapProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: c.Metadata,
	}
}

// Options for the ConfigMap-based volume.
// Experimental.
type ConfigMapVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	// Experimental.
	DefaultMode float64 `json:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the ConfigMap, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	// Experimental.
	Items map[string]PathMapping `json:"items"`
	// The volume name.
	// Experimental.
	Name string `json:"name"`
	// Specify whether the ConfigMap or its keys must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

// A single application container that you want to run within a pod.
// Experimental.
type Container interface {
	Args() []string
	Command() []string
	Env() map[string]EnvValue
	Image() string
	ImagePullPolicy() ImagePullPolicy
	Mounts() []VolumeMount
	Name() string
	Port() float64
	WorkingDir() string
	AddEnv(name string, value EnvValue)
	Mount(path string, volume Volume, options MountOptions)
}

// The jsii proxy struct for Container
type container struct {
	_ byte // padding
}

func (c *container) Args() []string {
	var returns []string
	_jsii_.Get(
		c,
		"args",
		&returns,
	)
	return returns
}

func (c *container) Command() []string {
	var returns []string
	_jsii_.Get(
		c,
		"command",
		&returns,
	)
	return returns
}

func (c *container) Env() map[string]EnvValue {
	var returns map[string]EnvValue
	_jsii_.Get(
		c,
		"env",
		&returns,
	)
	return returns
}

func (c *container) Image() string {
	var returns string
	_jsii_.Get(
		c,
		"image",
		&returns,
	)
	return returns
}

func (c *container) ImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		c,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (c *container) Mounts() []VolumeMount {
	var returns []VolumeMount
	_jsii_.Get(
		c,
		"mounts",
		&returns,
	)
	return returns
}

func (c *container) Name() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
	)
	return returns
}

func (c *container) Port() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"port",
		&returns,
	)
	return returns
}

func (c *container) WorkingDir() string {
	var returns string
	_jsii_.Get(
		c,
		"workingDir",
		&returns,
	)
	return returns
}


func NewContainer(props ContainerProps) Container {
	_init_.Initialize()
	c := container{}

	_jsii_.Create(
		"cdk8s-plus-17.Container",
		[]interface{}{props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&c,
	)
	return &c
}

// Add an environment value to the container.
//
// The variable value can come
// from various dynamic sources such a secrets of config maps.
// See: EnvValue.fromXXX
//
// Experimental.
func (c *container) AddEnv(name string, value EnvValue) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addEnv",
		[]interface{}{name, value},
		false,
		&returns,
	)
}

// Mount a volume to a specific path so that it is accessible by the container.
//
// Every pod that is configured to use this container will autmoatically have access to the volume.
// Experimental.
func (c *container) Mount(path string, volume Volume, options MountOptions) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"mount",
		[]interface{}{path, volume, options},
		false,
		&returns,
	)
}

// Properties for creating a container.
// Experimental.
type ContainerProps struct {
	// Docker image name.
	// Experimental.
	Image string `json:"image"`
	// Arguments to the entrypoint. The docker image's CMD is used if `command` is not provided.
	//
	// Variable references $(VAR_NAME) are expanded using the container's
	// environment. If a variable cannot be resolved, the reference in the input
	// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	//
	// Cannot be updated.
	// See: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	// Experimental.
	Args []string `json:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
	// Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// Experimental.
	Command []string `json:"command"`
	// List of environment variables to set in the container.
	//
	// Cannot be updated.
	// Experimental.
	Env map[string]EnvValue `json:"env"`
	// Image pull policy for this container.
	// Experimental.
	ImagePullPolicy ImagePullPolicy `json:"imagePullPolicy"`
	// Periodic probe of container liveness.
	//
	// Container will be restarted if the probe fails.
	// Experimental.
	Liveness Probe `json:"liveness"`
	// Name of the container specified as a DNS_LABEL.
	//
	// Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	// Experimental.
	Name string `json:"name"`
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	// Experimental.
	Port float64 `json:"port"`
	// Determines when the container is ready to serve traffic.
	// Experimental.
	Readiness Probe `json:"readiness"`
	// StartupProbe indicates that the Pod has successfully initialized.
	//
	// If specified, no other probes are executed until this completes successfully
	// Experimental.
	Startup Probe `json:"startup"`
	// Pod volumes to mount into the container's filesystem.
	//
	// Cannot be updated.
	// Experimental.
	VolumeMounts []VolumeMount `json:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	// Experimental.
	WorkingDir string `json:"workingDir"`
}

// A Deployment provides declarative updates for Pods and ReplicaSets.
//
// You describe a desired state in a Deployment, and the Deployment Controller changes the actual
// state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove
// existing Deployments and adopt all their resources with new Deployments.
//
// > Note: Do not manage ReplicaSets owned by a Deployment. Consider opening an issue in the main Kubernetes repository if your use case is not covered below.
//
// Use Case
// ---------
//
// The following are typical use cases for Deployments:
//
// - Create a Deployment to rollout a ReplicaSet. The ReplicaSet creates Pods in the background.
//    Check the status of the rollout to see if it succeeds or not.
// - Declare the new state of the Pods by updating the PodTemplateSpec of the Deployment.
//    A new ReplicaSet is created and the Deployment manages moving the Pods from the old ReplicaSet to the new one at a controlled rate.
//    Each new ReplicaSet updates the revision of the Deployment.
// - Rollback to an earlier Deployment revision if the current state of the Deployment is not stable.
//    Each rollback updates the revision of the Deployment.
// - Scale up the Deployment to facilitate more load.
// - Pause the Deployment to apply multiple fixes to its PodTemplateSpec and then resume it to start a new rollout.
// - Use the status of the Deployment as an indicator that a rollout has stuck.
// - Clean up older ReplicaSets that you don't need anymore.
// Experimental.
type Deployment interface {
	Resource
	IPodTemplate
	ApiObject() cdk8s.ApiObject
	Containers() []Container
	LabelSelector() map[string]string
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	Replicas() float64
	RestartPolicy() RestartPolicy
	ServiceAccount() IServiceAccount
	Volumes() []Volume
	AddContainer(container ContainerProps) Container
	AddVolume(volume Volume)
	Expose(port float64, options ExposeOptions) Service
	SelectByLabel(key string, value string)
}

// The jsii proxy struct for Deployment
type deployment struct {
	resource // extends cdk8s-plus-17.Resource
	iPodTemplate // implements cdk8s-plus-17.IPodTemplate
}

func (d *deployment) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		d,
		"apiObject",
		&returns,
	)
	return returns
}

func (d *deployment) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		d,
		"containers",
		&returns,
	)
	return returns
}

func (d *deployment) LabelSelector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"labelSelector",
		&returns,
	)
	return returns
}

func (d *deployment) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		d,
		"podMetadata",
		&returns,
	)
	return returns
}

func (d *deployment) Replicas() float64 {
	var returns float64
	_jsii_.Get(
		d,
		"replicas",
		&returns,
	)
	return returns
}

func (d *deployment) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		d,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (d *deployment) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		d,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (d *deployment) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		d,
		"volumes",
		&returns,
	)
	return returns
}

func (d *deployment) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		d,
		"metadata",
		&returns,
	)
	return returns
}

func (d *deployment) Name() string {
	var returns string
	_jsii_.Get(
		d,
		"name",
		&returns,
	)
	return returns
}


func NewDeployment(scope constructs.Construct, id string, props DeploymentProps) Deployment {
	_init_.Initialize()
	d := deployment{}

	_jsii_.Create(
		"cdk8s-plus-17.Deployment",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&d,
	)
	return &d
}

// Add a container to the pod.
// Experimental.
func (d *deployment) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

// Add a volume to the pod.
// Experimental.
func (d *deployment) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

// Expose a deployment via a service.
//
// This is equivalent to running `kubectl expose deployment <deployment-name>`.
// Experimental.
func (d *deployment) Expose(port float64, options ExposeOptions) Service {
	var returns Service
	_jsii_.Invoke(
		d,
		"expose",
		[]interface{}{port, options},
		true,
		&returns,
	)
	return returns
}

// Configure a label selector to this deployment.
//
// Pods that have the label will be selected by deployments configured with this spec.
// Experimental.
func (d *deployment) SelectByLabel(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"selectByLabel",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *deployment) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *deployment) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *deployment) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		d,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *deployment) ToString() string {
	var returns string
	_jsii_.Invoke(
		d,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Properties for initialization of `Deployment`.
// Experimental.
type DeploymentProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadata `json:"podMetadata"`
	// Automatically allocates a pod selector for this deployment.
	//
	// If this is set to `false` you must define your selector through
	// `deployment.podMetadata.addLabel()` and `deployment.selectByLabel()`.
	// Experimental.
	DefaultSelector bool `json:"defaultSelector"`
	// Number of desired pods.
	// Experimental.
	Replicas float64 `json:"replicas"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this DeploymentProps.
func (d *DeploymentProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: d.Metadata,
	}
}

// ToPodSpecProps is a convenience function to obtain a new PodSpecProps from this DeploymentProps.
func (d *DeploymentProps) ToPodSpecProps() PodSpecProps {
	return PodSpecProps {
		Containers: d.Containers,
		RestartPolicy: d.RestartPolicy,
		ServiceAccount: d.ServiceAccount,
		Volumes: d.Volumes,
	}
}

// ToPodTemplateProps is a convenience function to obtain a new PodTemplateProps from this DeploymentProps.
func (d *DeploymentProps) ToPodTemplateProps() PodTemplateProps {
	return PodTemplateProps {
		Containers: d.Containers,
		RestartPolicy: d.RestartPolicy,
		ServiceAccount: d.ServiceAccount,
		Volumes: d.Volumes,
		PodMetadata: d.PodMetadata,
	}
}

// The medium on which to store the volume.
// Experimental.
type EmptyDirMedium string

const (
	EmptyDirMedium_DEFAULT EmptyDirMedium = "DEFAULT"
	EmptyDirMedium_MEMORY EmptyDirMedium = "MEMORY"
)

// Options for volumes populated with an empty directory.
// Experimental.
type EmptyDirVolumeOptions struct {
	// By default, emptyDir volumes are stored on whatever medium is backing the node - that might be disk or SSD or network storage, depending on your environment.
	//
	// However, you can set the emptyDir.medium field to
	// `EmptyDirMedium.MEMORY` to tell Kubernetes to mount a tmpfs (RAM-backed
	// filesystem) for you instead. While tmpfs is very fast, be aware that unlike
	// disks, tmpfs is cleared on node reboot and any files you write will count
	// against your Container's memory limit.
	// Experimental.
	Medium EmptyDirMedium `json:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// The size
	// limit is also applicable for memory medium. The maximum usage on memory
	// medium EmptyDir would be the minimum value between the SizeLimit specified
	// here and the sum of memory limits of all containers in a pod.
	// Experimental.
	SizeLimit cdk8s.Size `json:"sizeLimit"`
}

// Utility class for creating reading env values from various sources.
// Experimental.
type EnvValue interface {
	Value() interface{}
	ValueFrom() interface{}
}

// The jsii proxy struct for EnvValue
type envValue struct {
	_ byte // padding
}

func (e *envValue) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		e,
		"value",
		&returns,
	)
	return returns
}

func (e *envValue) ValueFrom() interface{} {
	var returns interface{}
	_jsii_.Get(
		e,
		"valueFrom",
		&returns,
	)
	return returns
}


// Create a value by reading a specific key inside a config map.
// Experimental.
func EnvValue_FromConfigMap(configMap IConfigMap, key string, options EnvValueFromConfigMapOptions) EnvValue {
	_init_.Initialize()
	var returns EnvValue
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.EnvValue",
		"fromConfigMap",
		[]interface{}{configMap, key, options},
		true,
		&returns,
	)
	return returns
}

// Create a value from a key in the current process environment.
// Experimental.
func EnvValue_FromProcess(key string, options EnvValueFromProcessOptions) EnvValue {
	_init_.Initialize()
	var returns EnvValue
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.EnvValue",
		"fromProcess",
		[]interface{}{key, options},
		true,
		&returns,
	)
	return returns
}

// Defines an environment value from a secret JSON value.
// Experimental.
func EnvValue_FromSecretValue(secretValue SecretValue, options EnvValueFromSecretOptions) EnvValue {
	_init_.Initialize()
	var returns EnvValue
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.EnvValue",
		"fromSecretValue",
		[]interface{}{secretValue, options},
		true,
		&returns,
	)
	return returns
}

// Create a value from the given argument.
// Experimental.
func EnvValue_FromValue(value string) EnvValue {
	_init_.Initialize()
	var returns EnvValue
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.EnvValue",
		"fromValue",
		[]interface{}{value},
		true,
		&returns,
	)
	return returns
}

// Options to specify an envionment variable value from a ConfigMap key.
// Experimental.
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

// Options to specify an environment variable value from the process environment.
// Experimental.
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	// Experimental.
	Required bool `json:"required"`
}

// Options to specify an environment variable value from a Secret.
// Experimental.
type EnvValueFromSecretOptions struct {
	// Specify whether the Secret or its key must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

// Options for exposing a deployment via a service.
// Experimental.
type ExposeOptions struct {
	// The name of the service to expose.
	//
	// This will be set on the Service.metadata and must be a DNS_LABEL
	// Experimental.
	Name string `json:"name"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	// Experimental.
	Protocol Protocol `json:"protocol"`
	// The type of the exposed service.
	// Experimental.
	ServiceType ServiceType `json:"serviceType"`
	// The port number the service will redirect to.
	// Experimental.
	TargetPort float64 `json:"targetPort"`
}

// Options for `Probe.fromHttpGet()`.
// Experimental.
type HttpGetProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	// Experimental.
	FailureThreshold float64 `json:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	// Experimental.
	SuccessThreshold float64 `json:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds"`
	// The TCP port to use when sending the GET request.
	// Experimental.
	Port float64 `json:"port"`
}

// ToProbeOptions is a convenience function to obtain a new ProbeOptions from this HttpGetProbeOptions.
func (h *HttpGetProbeOptions) ToProbeOptions() ProbeOptions {
	return ProbeOptions {
		FailureThreshold: h.FailureThreshold,
		InitialDelaySeconds: h.InitialDelaySeconds,
		PeriodSeconds: h.PeriodSeconds,
		SuccessThreshold: h.SuccessThreshold,
		TimeoutSeconds: h.TimeoutSeconds,
	}
}

// Represents a config map.
// Experimental.
type IConfigMap interface {
	IResource
}

// The jsii proxy for IConfigMap
type iConfigMap struct {
	iResource // extends cdk8s-plus-17.IResource
}

// Represents a resource that can be configured with a kuberenets pod spec. (e.g `Deployment`, `Job`, `Pod`, ...).
//
// Use the `PodSpec` class as an implementation helper.
// Experimental.
type IPodSpec interface {
	// Add a container to the pod.
	// Experimental.
	AddContainer(container ContainerProps) Container
	// Add a volume to the pod.
	// Experimental.
	AddVolume(volume Volume)
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers() []Container
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy() RestartPolicy
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount() IServiceAccount
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes() []Volume
}

// The jsii proxy for IPodSpec
type iPodSpec struct {
	_ byte // padding
}

func (i *iPodSpec) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		i,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

func (i *iPodSpec) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

func (i *iPodSpec) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		i,
		"containers",
		&returns,
	)
	return returns
}

func (i *iPodSpec) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		i,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (i *iPodSpec) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		i,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (i *iPodSpec) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		i,
		"volumes",
		&returns,
	)
	return returns
}

// Represents a resource that can be configured with a kuberenets pod template. (e.g `Deployment`, `Job`, ...).
//
// Use the `PodTemplate` class as an implementation helper.
// Experimental.
type IPodTemplate interface {
	IPodSpec
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
}

// The jsii proxy for IPodTemplate
type iPodTemplate struct {
	iPodSpec // extends cdk8s-plus-17.IPodSpec
}

func (i *iPodTemplate) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		i,
		"podMetadata",
		&returns,
	)
	return returns
}

// Represents a resource.
// Experimental.
type IResource interface {
	// The Kubernetes name of this resource.
	// Experimental.
	Name() string
}

// The jsii proxy for IResource
type iResource struct {
	_ byte // padding
}

func (i *iResource) Name() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
	)
	return returns
}

// Experimental.
type ISecret interface {
	IResource
}

// The jsii proxy for ISecret
type iSecret struct {
	iResource // extends cdk8s-plus-17.IResource
}

// Experimental.
type IServiceAccount interface {
	IResource
}

// The jsii proxy for IServiceAccount
type iServiceAccount struct {
	iResource // extends cdk8s-plus-17.IResource
}

// Experimental.
type ImagePullPolicy string

const (
	ImagePullPolicy_ALWAYS ImagePullPolicy = "ALWAYS"
	ImagePullPolicy_IF_NOT_PRESENT ImagePullPolicy = "IF_NOT_PRESENT"
	ImagePullPolicy_NEVER ImagePullPolicy = "NEVER"
)

// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend.
//
// An Ingress can be configured to give services
// externally-reachable urls, load balance traffic, terminate SSL, offer name
// based virtual hosting etc.
// Experimental.
type IngressV1Beta1 interface {
	Resource
	ApiObject() cdk8s.ApiObject
	AddDefaultBackend(backend IngressV1Beta1Backend)
	AddHostDefaultBackend(host string, backend IngressV1Beta1Backend)
	AddHostRule(host string, path string, backend IngressV1Beta1Backend)
	AddRule(path string, backend IngressV1Beta1Backend)
	AddRules(rules IngressV1Beta1Rule)
	OnValidate() []string
}

// The jsii proxy struct for IngressV1Beta1
type ingressV1Beta1 struct {
	resource // extends cdk8s-plus-17.Resource
}

func (i *ingressV1Beta1) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		i,
		"apiObject",
		&returns,
	)
	return returns
}


func NewIngressV1Beta1(scope constructs.Construct, id string, props IngressV1Beta1Props) IngressV1Beta1 {
	_init_.Initialize()
	i := ingressV1Beta1{}

	_jsii_.Create(
		"cdk8s-plus-17.IngressV1Beta1",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&i,
	)
	return &i
}

// Defines the default backend for this ingress.
//
// A default backend capable of
// servicing requests that don't match any rule.
// Experimental.
func (i *ingressV1Beta1) AddDefaultBackend(backend IngressV1Beta1Backend) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addDefaultBackend",
		[]interface{}{backend},
		false,
		&returns,
	)
}

// Specify a default backend for a specific host name.
//
// This backend will be used as a catch-all for requests
// targeted to this host name (the `Host` header matches this value).
// Experimental.
func (i *ingressV1Beta1) AddHostDefaultBackend(host string, backend IngressV1Beta1Backend) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addHostDefaultBackend",
		[]interface{}{host, backend},
		false,
		&returns,
	)
}

// Adds an ingress rule applied to requests to a specific host and a specific HTTP path (the `Host` header matches this value).
// Experimental.
func (i *ingressV1Beta1) AddHostRule(host string, path string, backend IngressV1Beta1Backend) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addHostRule",
		[]interface{}{host, path, backend},
		false,
		&returns,
	)
}

// Adds an ingress rule applied to requests sent to a specific HTTP path.
// Experimental.
func (i *ingressV1Beta1) AddRule(path string, backend IngressV1Beta1Backend) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addRule",
		[]interface{}{path, backend},
		false,
		&returns,
	)
}

// Adds rules to this ingress.
// Experimental.
func (i *ingressV1Beta1) AddRules(rules IngressV1Beta1Rule) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addRules",
		[]interface{}{rules},
		false,
		&returns,
	)
}

// (deprecated) Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
// Experimental.
func (i *ingressV1Beta1) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		i,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// The backend for an ingress path.
// Experimental.
type IngressV1Beta1Backend interface {
}

// The jsii proxy struct for IngressV1Beta1Backend
type ingressV1Beta1Backend struct {
	_ byte // padding
}

// A Kubernetes `Service` to use as the backend for this path.
// Experimental.
func IngressV1Beta1Backend_FromService(service Service, options ServiceIngressV1BetaBackendOptions) IngressV1Beta1Backend {
	_init_.Initialize()
	var returns IngressV1Beta1Backend
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.IngressV1Beta1Backend",
		"fromService",
		[]interface{}{service, options},
		true,
		&returns,
	)
	return returns
}

// Properties for `Ingress`.
// Experimental.
type IngressV1Beta1Props struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// The default backend services requests that do not match any rule.
	//
	// Using this option or the `addDefaultBackend()` method is equivalent to
	// adding a rule with both `path` and `host` undefined.
	// Experimental.
	DefaultBackend IngressV1Beta1Backend `json:"defaultBackend"`
	// Routing rules for this ingress.
	//
	// Each rule must define an `IngressBackend` that will receive the requests
	// that match this rule. If both `host` and `path` are not specifiec, this
	// backend will be used as the default backend of the ingress.
	//
	// You can also add rules later using `addRule()`, `addHostRule()`,
	// `addDefaultBackend()` and `addHostDefaultBackend()`.
	// Experimental.
	Rules []IngressV1Beta1Rule `json:"rules"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this IngressV1Beta1Props.
func (i *IngressV1Beta1Props) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: i.Metadata,
	}
}

// Represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match,
// then routed to the backend associated with the matching path.
// Experimental.
type IngressV1Beta1Rule struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	// Experimental.
	Backend IngressV1Beta1Backend `json:"backend"`
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986.
	//
	// Note the following deviations from the "host" part of the URI as
	// defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue
	// can only apply to the IP in the Spec of the parent Ingress. 2. The `:`
	// delimiter is not respected because ports are not allowed. Currently the
	// port of an Ingress is implicitly :80 for http and :443 for https. Both
	// these may change in the future. Incoming requests are matched against the
	// host before the IngressRuleValue.
	// Experimental.
	Host string `json:"host"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'.
	// Experimental.
	Path string `json:"path"`
}

// A Job creates one or more Pods and ensures that a specified number of them successfully terminate.
//
// As pods successfully complete,
// the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete.
// Deleting a Job will clean up the Pods it created. A simple case is to create one Job object in order to reliably run one Pod to completion.
// The Job object will start a new Pod if the first Pod fails or is deleted (for example due to a node hardware failure or a node reboot).
// You can also use a Job to run multiple Pods in parallel.
// Experimental.
type Job interface {
	Resource
	IPodTemplate
	ActiveDeadline() cdk8s.Duration
	ApiObject() cdk8s.ApiObject
	BackoffLimit() float64
	Containers() []Container
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	RestartPolicy() RestartPolicy
	ServiceAccount() IServiceAccount
	TtlAfterFinished() cdk8s.Duration
	Volumes() []Volume
	AddContainer(container ContainerProps) Container
	AddVolume(volume Volume)
}

// The jsii proxy struct for Job
type job struct {
	resource // extends cdk8s-plus-17.Resource
	iPodTemplate // implements cdk8s-plus-17.IPodTemplate
}

func (j *job) ActiveDeadline() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"activeDeadline",
		&returns,
	)
	return returns
}

func (j *job) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
	)
	return returns
}

func (j *job) BackoffLimit() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
	)
	return returns
}

func (j *job) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *job) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
	)
	return returns
}

func (j *job) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *job) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *job) TtlAfterFinished() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"ttlAfterFinished",
		&returns,
	)
	return returns
}

func (j *job) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *job) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *job) Name() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewJob(scope constructs.Construct, id string, props JobProps) Job {
	_init_.Initialize()
	j := job{}

	_jsii_.Create(
		"cdk8s-plus-17.Job",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&j,
	)
	return &j
}

// Add a container to the pod.
// Experimental.
func (j *job) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		j,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

// Add a volume to the pod.
// Experimental.
func (j *job) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (j *job) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (j *job) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (j *job) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		j,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (j *job) ToString() string {
	var returns string
	_jsii_.Invoke(
		j,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Properties for initialization of `Job`.
// Experimental.
type JobProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadata `json:"podMetadata"`
	// Specifies the duration the job may be active before the system tries to terminate it.
	// Experimental.
	ActiveDeadline cdk8s.Duration `json:"activeDeadline"`
	// Specifies the number of retries before marking this job failed.
	// Experimental.
	BackoffLimit float64 `json:"backoffLimit"`
	// Limits the lifetime of a Job that has finished execution (either Complete or Failed).
	//
	// If this field is set, after the Job finishes, it is eligible to
	// be automatically deleted. When the Job is being deleted, its lifecycle
	// guarantees (e.g. finalizers) will be honored. If this field is set to zero,
	// the Job becomes eligible to be deleted immediately after it finishes. This
	// field is alpha-level and is only honored by servers that enable the
	// `TTLAfterFinished` feature.
	// Experimental.
	TtlAfterFinished cdk8s.Duration `json:"ttlAfterFinished"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this JobProps.
func (j *JobProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: j.Metadata,
	}
}

// ToPodSpecProps is a convenience function to obtain a new PodSpecProps from this JobProps.
func (j *JobProps) ToPodSpecProps() PodSpecProps {
	return PodSpecProps {
		Containers: j.Containers,
		RestartPolicy: j.RestartPolicy,
		ServiceAccount: j.ServiceAccount,
		Volumes: j.Volumes,
	}
}

// ToPodTemplateProps is a convenience function to obtain a new PodTemplateProps from this JobProps.
func (j *JobProps) ToPodTemplateProps() PodTemplateProps {
	return PodTemplateProps {
		Containers: j.Containers,
		RestartPolicy: j.RestartPolicy,
		ServiceAccount: j.ServiceAccount,
		Volumes: j.Volumes,
		PodMetadata: j.PodMetadata,
	}
}

// Options for mounts.
// Experimental.
type MountOptions struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	//
	// This field is beta in 1.10.
	// Experimental.
	Propagation MountPropagation `json:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	// Experimental.
	ReadOnly bool `json:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	// Experimental.
	SubPath string `json:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root). SubPathExpr and SubPath are mutually exclusive. This field
	// is beta in 1.15.
	//
	// `subPathExpr` and `subPath` are mutually exclusive. This field is beta in
	// 1.15.
	// Experimental.
	SubPathExpr string `json:"subPathExpr"`
}

// Experimental.
type MountPropagation string

const (
	MountPropagation_NONE MountPropagation = "NONE"
	MountPropagation_HOST_TO_CONTAINER MountPropagation = "HOST_TO_CONTAINER"
	MountPropagation_BIDIRECTIONAL MountPropagation = "BIDIRECTIONAL"
)

// Maps a string key to a path within a volume.
// Experimental.
type PathMapping struct {
	// The relative path of the file to map the key to.
	//
	// May not be an absolute
	// path. May not contain the path element '..'. May not start with the string
	// '..'.
	// Experimental.
	Path string `json:"path"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be
	// in conflict with other options that affect the file mode, like fsGroup, and
	// the result can be other mode bits set.
	// Experimental.
	Mode float64 `json:"mode"`
}

// Pod is a collection of containers that can run on a host.
//
// This resource is
// created by clients and scheduled onto hosts.
// Experimental.
type Pod interface {
	Resource
	IPodSpec
	ApiObject() cdk8s.ApiObject
	Containers() []Container
	RestartPolicy() RestartPolicy
	ServiceAccount() IServiceAccount
	Volumes() []Volume
	AddContainer(container ContainerProps) Container
	AddVolume(volume Volume)
}

// The jsii proxy struct for Pod
type pod struct {
	resource // extends cdk8s-plus-17.Resource
	iPodSpec // implements cdk8s-plus-17.IPodSpec
}

func (p *pod) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		p,
		"apiObject",
		&returns,
	)
	return returns
}

func (p *pod) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		p,
		"containers",
		&returns,
	)
	return returns
}

func (p *pod) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (p *pod) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (p *pod) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		p,
		"volumes",
		&returns,
	)
	return returns
}

func (p *pod) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		p,
		"metadata",
		&returns,
	)
	return returns
}

func (p *pod) Name() string {
	var returns string
	_jsii_.Get(
		p,
		"name",
		&returns,
	)
	return returns
}


func NewPod(scope constructs.Construct, id string, props PodProps) Pod {
	_init_.Initialize()
	p := pod{}

	_jsii_.Create(
		"cdk8s-plus-17.Pod",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodSpec"},
		[]_jsii_.Override{},
		&p,
	)
	return &p
}

// Add a container to the pod.
// Experimental.
func (p *pod) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

// Add a volume to the pod.
// Experimental.
func (p *pod) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (p *pod) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *pod) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (p *pod) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		p,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *pod) ToString() string {
	var returns string
	_jsii_.Invoke(
		p,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
//
// The default policy is `OrderedReady`, where pods are created in increasing order
// (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before
// continuing. When scaling down, the pods are removed in the opposite order.
//
// The alternative policy is `Parallel` which will create pods in parallel to match the
// desired scale without waiting, and on scale down will delete all pods at once.
// Experimental.
type PodManagementPolicy string

const (
	PodManagementPolicy_ORDERED_READY PodManagementPolicy = "ORDERED_READY"
	PodManagementPolicy_PARALLEL PodManagementPolicy = "PARALLEL"
)

// Properties for initialization of `Pod`.
// Experimental.
type PodProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this PodProps.
func (p *PodProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: p.Metadata,
	}
}

// ToPodSpecProps is a convenience function to obtain a new PodSpecProps from this PodProps.
func (p *PodProps) ToPodSpecProps() PodSpecProps {
	return PodSpecProps {
		Containers: p.Containers,
		RestartPolicy: p.RestartPolicy,
		ServiceAccount: p.ServiceAccount,
		Volumes: p.Volumes,
	}
}

// Provides read/write capabilities ontop of a `PodSpecProps`.
// Experimental.
type PodSpec interface {
	IPodSpec
	Containers() []Container
	RestartPolicy() RestartPolicy
	ServiceAccount() IServiceAccount
	Volumes() []Volume
	AddContainer(container ContainerProps) Container
	AddVolume(volume Volume)
}

// The jsii proxy struct for PodSpec
type podSpec struct {
	iPodSpec // implements cdk8s-plus-17.IPodSpec
}

func (p *podSpec) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		p,
		"containers",
		&returns,
	)
	return returns
}

func (p *podSpec) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (p *podSpec) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (p *podSpec) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		p,
		"volumes",
		&returns,
	)
	return returns
}


func NewPodSpec(props PodSpecProps) PodSpec {
	_init_.Initialize()
	p := podSpec{}

	_jsii_.Create(
		"cdk8s-plus-17.PodSpec",
		[]interface{}{props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodSpec"},
		[]_jsii_.Override{},
		&p,
	)
	return &p
}

// Add a container to the pod.
// Experimental.
func (p *podSpec) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

// Add a volume to the pod.
// Experimental.
func (p *podSpec) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

// Properties of a `PodSpec`.
// Experimental.
type PodSpecProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
}

// Provides read/write capabilities ontop of a `PodTemplateProps`.
// Experimental.
type PodTemplate interface {
	PodSpec
	IPodTemplate
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
}

// The jsii proxy struct for PodTemplate
type podTemplate struct {
	podSpec // extends cdk8s-plus-17.PodSpec
	iPodTemplate // implements cdk8s-plus-17.IPodTemplate
}

func (p *podTemplate) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		p,
		"podMetadata",
		&returns,
	)
	return returns
}


func NewPodTemplate(props PodTemplateProps) PodTemplate {
	_init_.Initialize()
	p := podTemplate{}

	_jsii_.Create(
		"cdk8s-plus-17.PodTemplate",
		[]interface{}{props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&p,
	)
	return &p
}

// Properties of a `PodTemplate`.
//
// Adds metadata information on top of the spec.
// Experimental.
type PodTemplateProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadata `json:"podMetadata"`
}

// ToPodSpecProps is a convenience function to obtain a new PodSpecProps from this PodTemplateProps.
func (p *PodTemplateProps) ToPodSpecProps() PodSpecProps {
	return PodSpecProps {
		Containers: p.Containers,
		RestartPolicy: p.RestartPolicy,
		ServiceAccount: p.ServiceAccount,
		Volumes: p.Volumes,
	}
}

// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
// Experimental.
type Probe interface {
}

// The jsii proxy struct for Probe
type probe struct {
	_ byte // padding
}

func NewProbe() Probe {
	_init_.Initialize()
	p := probe{}

	_jsii_.Create(
		"cdk8s-plus-17.Probe",
		[]interface{}{},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&p,
	)
	return &p
}

// Defines a probe based on a command which is executed within the container.
// Experimental.
func Probe_FromCommand(command []string, options CommandProbeOptions) Probe {
	_init_.Initialize()
	var returns Probe
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.Probe",
		"fromCommand",
		[]interface{}{command, options},
		true,
		&returns,
	)
	return returns
}

// Defines a probe based on an HTTP GET request to the IP address of the container.
// Experimental.
func Probe_FromHttpGet(path string, options HttpGetProbeOptions) Probe {
	_init_.Initialize()
	var returns Probe
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.Probe",
		"fromHttpGet",
		[]interface{}{path, options},
		true,
		&returns,
	)
	return returns
}

// Probe options.
// Experimental.
type ProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	// Experimental.
	FailureThreshold float64 `json:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	InitialDelaySeconds cdk8s.Duration `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.Duration `json:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	// Experimental.
	SuccessThreshold float64 `json:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Experimental.
	TimeoutSeconds cdk8s.Duration `json:"timeoutSeconds"`
}

// Experimental.
type Protocol string

const (
	Protocol_TCP Protocol = "TCP"
	Protocol_UDP Protocol = "UDP"
	Protocol_SCTP Protocol = "SCTP"
)

// Base class for all Kubernetes objects in stdk8s.
//
// Represents a single
// resource.
// Experimental.
type Resource interface {
	constructs.Construct
	IResource
	ApiObject() cdk8s.ApiObject
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() string
}

// The jsii proxy struct for Resource
type resource struct {
	constructs.Construct // extends constructs.Construct
	iResource // implements cdk8s-plus-17.IResource
}

func (r *resource) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		r,
		"apiObject",
		&returns,
	)
	return returns
}

func (r *resource) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		r,
		"metadata",
		&returns,
	)
	return returns
}

func (r *resource) Name() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
	)
	return returns
}


// Creates a new construct node.
func NewResource(scope constructs.Construct, id string, options constructs.ConstructOptions) Resource {
	_init_.Initialize()
	r := resource{}

	_jsii_.Create(
		"cdk8s-plus-17.Resource",
		[]interface{}{scope, id, options},
		[]_jsii_.FQN{"cdk8s-plus-17.IResource"},
		[]_jsii_.Override{},
		&r,
	)
	return &r
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (r *resource) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *resource) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (r *resource) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		r,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (r *resource) ToString() string {
	var returns string
	_jsii_.Invoke(
		r,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Initialization properties for resources.
// Experimental.
type ResourceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
}

// Restart policy for all containers within the pod.
// Experimental.
type RestartPolicy string

const (
	RestartPolicy_ALWAYS RestartPolicy = "ALWAYS"
	RestartPolicy_ON_FAILURE RestartPolicy = "ON_FAILURE"
	RestartPolicy_NEVER RestartPolicy = "NEVER"
)

// Kubernetes Secrets let you store and manage sensitive information, such as passwords, OAuth tokens, and ssh keys.
//
// Storing confidential information in a
// Secret is safer and more flexible than putting it verbatim in a Pod
// definition or in a container image.
// See: https://kubernetes.io/docs/concepts/configuration/secret
//
// Experimental.
type Secret interface {
	Resource
	ISecret
	ApiObject() cdk8s.ApiObject
	AddStringData(key string, value string)
	GetStringData(key string) string
}

// The jsii proxy struct for Secret
type secret struct {
	resource // extends cdk8s-plus-17.Resource
	iSecret // implements cdk8s-plus-17.ISecret
}

func (s *secret) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
	)
	return returns
}

func (s *secret) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		s,
		"metadata",
		&returns,
	)
	return returns
}

func (s *secret) Name() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
	)
	return returns
}


func NewSecret(scope constructs.Construct, id string, props SecretProps) Secret {
	_init_.Initialize()
	s := secret{}

	_jsii_.Create(
		"cdk8s-plus-17.Secret",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.ISecret"},
		[]_jsii_.Override{},
		&s,
	)
	return &s
}

// Imports a secret from the cluster as a reference.
// Experimental.
func Secret_FromSecretName(name string) ISecret {
	_init_.Initialize()
	var returns ISecret
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.Secret",
		"fromSecretName",
		[]interface{}{name},
		true,
		&returns,
	)
	return returns
}

// Adds a string data field to the secert.
// Experimental.
func (s *secret) AddStringData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addStringData",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Gets a string data by key or undefined.
// Experimental.
func (s *secret) GetStringData(key string) string {
	var returns string
	_jsii_.Invoke(
		s,
		"getStringData",
		[]interface{}{key},
		true,
		&returns,
	)
	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *secret) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *secret) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *secret) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *secret) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Experimental.
type SecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// stringData allows specifying non-binary secret data in string form.
	//
	// It is
	// provided as a write-only convenience method. All keys and values are merged
	// into the data field on write, overwriting any existing values. It is never
	// output when reading from the API.
	// Experimental.
	StringData map[string]string `json:"stringData"`
	// Optional type associated with the secret.
	//
	// Used to facilitate programmatic
	// handling of secret data by various controllers.
	// Experimental.
	Type string `json:"type"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this SecretProps.
func (s *SecretProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: s.Metadata,
	}
}

// Represents a specific value in JSON secret.
// Experimental.
type SecretValue struct {
	// The JSON key.
	// Experimental.
	Key string `json:"key"`
	// The secret.
	// Experimental.
	Secret ISecret `json:"secret"`
}

// An abstract way to expose an application running on a set of Pods as a network service.
//
// With Kubernetes you don't need to modify your application to use an unfamiliar service discovery mechanism.
// Kubernetes gives Pods their own IP addresses and a single DNS name for a set of Pods, and can load-balance across them.
//
// For example, consider a stateless image-processing backend which is running with 3 replicas. Those replicas are fungible—frontends do not care which backend they use.
// While the actual Pods that compose the backend set may change, the frontend clients should not need to be aware of that,
// nor should they need to keep track of the set of backends themselves.
// The Service abstraction enables this decoupling.
//
// If you're able to use Kubernetes APIs for service discovery in your application, you can query the API server for Endpoints,
// that get updated whenever the set of Pods in a Service changes. For non-native applications, Kubernetes offers ways to place a network port
// or load balancer in between your application and the backend Pods.
// Experimental.
type Service interface {
	Resource
	ApiObject() cdk8s.ApiObject
	ClusterIp() string
	ExternalName() string
	Ports() []ServicePort
	Selector() map[string]string
	Type() ServiceType
	AddDeployment(deployment Deployment, port float64, options ServicePortOptions)
	AddSelector(label string, value string)
	Serve(port float64, options ServicePortOptions)
}

// The jsii proxy struct for Service
type service struct {
	resource // extends cdk8s-plus-17.Resource
}

func (s *service) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
	)
	return returns
}

func (s *service) ClusterIp() string {
	var returns string
	_jsii_.Get(
		s,
		"clusterIP",
		&returns,
	)
	return returns
}

func (s *service) ExternalName() string {
	var returns string
	_jsii_.Get(
		s,
		"externalName",
		&returns,
	)
	return returns
}

func (s *service) Ports() []ServicePort {
	var returns []ServicePort
	_jsii_.Get(
		s,
		"ports",
		&returns,
	)
	return returns
}

func (s *service) Selector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"selector",
		&returns,
	)
	return returns
}

func (s *service) Type() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		s,
		"type",
		&returns,
	)
	return returns
}


func NewService(scope constructs.Construct, id string, props ServiceProps) Service {
	_init_.Initialize()
	s := service{}

	_jsii_.Create(
		"cdk8s-plus-17.Service",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&s,
	)
	return &s
}

// Associate a deployment to this service.
//
// If not targetPort is specific in the portOptions, then requests will be routed
// to the port exposed by the first container in the deployment's pods.
// The deployment's `labelSelector` will be used to select pods.
// Experimental.
func (s *service) AddDeployment(deployment Deployment, port float64, options ServicePortOptions) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addDeployment",
		[]interface{}{deployment, port, options},
		false,
		&returns,
	)
}

// Services defined using this spec will select pods according the provided label.
// Experimental.
func (s *service) AddSelector(label string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addSelector",
		[]interface{}{label, value},
		false,
		&returns,
	)
}

// Configure a port the service will bind to.
//
// This method can be called multiple times.
// Experimental.
func (s *service) Serve(port float64, options ServicePortOptions) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"serve",
		[]interface{}{port, options},
		false,
		&returns,
	)
}

// A service account provides an identity for processes that run in a Pod.
//
// When you (a human) access the cluster (for example, using kubectl), you are
// authenticated by the apiserver as a particular User Account (currently this
// is usually admin, unless your cluster administrator has customized your
// cluster). Processes in containers inside pods can also contact the apiserver.
// When they do, they are authenticated as a particular Service Account (for
// example, default).
// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account
//
// Experimental.
type ServiceAccount interface {
	Resource
	IServiceAccount
	ApiObject() cdk8s.ApiObject
	Secrets() []ISecret
	AddSecret(secret ISecret)
}

// The jsii proxy struct for ServiceAccount
type serviceAccount struct {
	resource // extends cdk8s-plus-17.Resource
	iServiceAccount // implements cdk8s-plus-17.IServiceAccount
}

func (s *serviceAccount) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
	)
	return returns
}

func (s *serviceAccount) Secrets() []ISecret {
	var returns []ISecret
	_jsii_.Get(
		s,
		"secrets",
		&returns,
	)
	return returns
}

func (s *serviceAccount) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		s,
		"metadata",
		&returns,
	)
	return returns
}

func (s *serviceAccount) Name() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
	)
	return returns
}


func NewServiceAccount(scope constructs.Construct, id string, props ServiceAccountProps) ServiceAccount {
	_init_.Initialize()
	s := serviceAccount{}

	_jsii_.Create(
		"cdk8s-plus-17.ServiceAccount",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IServiceAccount"},
		[]_jsii_.Override{},
		&s,
	)
	return &s
}

// Imports a service account from the cluster as a reference.
// Experimental.
func ServiceAccount_FromServiceAccountName(name string) IServiceAccount {
	_init_.Initialize()
	var returns IServiceAccount
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.ServiceAccount",
		"fromServiceAccountName",
		[]interface{}{name},
		true,
		&returns,
	)
	return returns
}

// Allow a secret to be accessed by pods using this service account.
// Experimental.
func (s *serviceAccount) AddSecret(secret ISecret) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addSecret",
		[]interface{}{secret},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *serviceAccount) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *serviceAccount) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *serviceAccount) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *serviceAccount) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Properties for initialization of `ServiceAccount`.
//
// Properties for initialization of `ServiceAccount`.
// Experimental.
type ServiceAccountProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// List of secrets allowed to be used by pods running using this ServiceAccount.
	// See: https://kubernetes.io/docs/concepts/configuration/secret
	//
	// Experimental.
	Secrets []ISecret `json:"secrets"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this ServiceAccountProps.
func (s *ServiceAccountProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: s.Metadata,
	}
}

// Options for setting up backends for ingress rules.
// Experimental.
type ServiceIngressV1BetaBackendOptions struct {
	// The port to use to access the service.
	//
	// - This option will fail if the service does not expose any ports.
	// - If the service exposes multiple ports, this option must be specified.
	// - If the service exposes a single port, this option is optional and if
	//    specified, it must be the same port exposed by the service.
	// Experimental.
	Port float64 `json:"port"`
}

// Definition of a service port.
// Experimental.
type ServicePort struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	// Experimental.
	Name string `json:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	// Experimental.
	NodePort float64 `json:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	// Experimental.
	Protocol Protocol `json:"protocol"`
	// The port number the service will redirect to.
	// Experimental.
	TargetPort float64 `json:"targetPort"`
	// The port number the service will bind to.
	// Experimental.
	Port float64 `json:"port"`
}

// ToServicePortOptions is a convenience function to obtain a new ServicePortOptions from this ServicePort.
func (s *ServicePort) ToServicePortOptions() ServicePortOptions {
	return ServicePortOptions {
		Name: s.Name,
		NodePort: s.NodePort,
		Protocol: s.Protocol,
		TargetPort: s.TargetPort,
	}
}

// Experimental.
type ServicePortOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	// Experimental.
	Name string `json:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	// Experimental.
	NodePort float64 `json:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	// Experimental.
	Protocol Protocol `json:"protocol"`
	// The port number the service will redirect to.
	// Experimental.
	TargetPort float64 `json:"targetPort"`
}

// Properties for initialization of `Service`.
// Experimental.
type ServiceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// The IP address of the service and is usually assigned randomly by the master.
	//
	// If an address is specified manually and is not in use by others, it
	// will be allocated to the service; otherwise, creation of the service will
	// fail. This field can not be changed through updates. Valid values are
	// "None", empty string (""), or a valid IP address. "None" can be specified
	// for headless services when proxying is not required. Only applies to types
	// ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	//
	// Experimental.
	ClusterIp string `json:"clusterIP"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user
	// is responsible for ensuring that traffic arrives at a node with this IP. A
	// common example is external load-balancers that are not part of the
	// Kubernetes system.
	// Experimental.
	ExternalIPs []string `json:"externalIPs"`
	// The externalName to be used when ServiceType.EXTERNAL_NAME is set.
	// Experimental.
	ExternalName string `json:"externalName"`
	// A list of CIDR IP addresses, if specified and supported by the platform, will restrict traffic through the cloud-provider load-balancer to the specified client IPs.
	//
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/
	// Experimental.
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges"`
	// The port exposed by this service.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// Experimental.
	Ports []ServicePort `json:"ports"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// Experimental.
	Type ServiceType `json:"type"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this ServiceProps.
func (s *ServiceProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: s.Metadata,
	}
}

// For some parts of your application (for example, frontends) you may want to expose a Service onto an external IP address, that's outside of your cluster.
//
// Kubernetes ServiceTypes allow you to specify what kind of Service you want.
// The default is ClusterIP.
// Experimental.
type ServiceType string

const (
	ServiceType_CLUSTER_IP ServiceType = "CLUSTER_IP"
	ServiceType_NODE_PORT ServiceType = "NODE_PORT"
	ServiceType_LOAD_BALANCER ServiceType = "LOAD_BALANCER"
	ServiceType_EXTERNAL_NAME ServiceType = "EXTERNAL_NAME"
)

// StatefulSet is the workload API object used to manage stateful applications.
//
// Manages the deployment and scaling of a set of Pods, and provides guarantees
// about the ordering and uniqueness of these Pods.
//
// Like a Deployment, a StatefulSet manages Pods that are based on an identical
// container spec. Unlike a Deployment, a StatefulSet maintains a sticky identity
// for each of their Pods. These pods are created from the same spec, but are not
// interchangeable: each has a persistent identifier that it maintains across any
// rescheduling.
//
// If you want to use storage volumes to provide persistence for your workload, you
// can use a StatefulSet as part of the solution. Although individual Pods in a StatefulSet
// are susceptible to failure, the persistent Pod identifiers make it easier to match existing
// volumes to the new Pods that replace any that have failed.
//
// Using StatefulSets
// ------------------
// StatefulSets are valuable for applications that require one or more of the following.
//
// - Stable, unique network identifiers.
// - Stable, persistent storage.
// - Ordered, graceful deployment and scaling.
// - Ordered, automated rolling updates.
// Experimental.
type StatefulSet interface {
	Resource
	IPodTemplate
	ApiObject() cdk8s.ApiObject
	Containers() []Container
	LabelSelector() map[string]string
	PodManagementPolicy() PodManagementPolicy
	PodMetadata() cdk8s.ApiObjectMetadataDefinition
	Replicas() float64
	RestartPolicy() RestartPolicy
	ServiceAccount() IServiceAccount
	Volumes() []Volume
	AddContainer(container ContainerProps) Container
	AddVolume(volume Volume)
	SelectByLabel(key string, value string)
}

// The jsii proxy struct for StatefulSet
type statefulSet struct {
	resource // extends cdk8s-plus-17.Resource
	iPodTemplate // implements cdk8s-plus-17.IPodTemplate
}

func (s *statefulSet) ApiObject() cdk8s.ApiObject {
	var returns cdk8s.ApiObject
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
	)
	return returns
}

func (s *statefulSet) Containers() []Container {
	var returns []Container
	_jsii_.Get(
		s,
		"containers",
		&returns,
	)
	return returns
}

func (s *statefulSet) LabelSelector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"labelSelector",
		&returns,
	)
	return returns
}

func (s *statefulSet) PodManagementPolicy() PodManagementPolicy {
	var returns PodManagementPolicy
	_jsii_.Get(
		s,
		"podManagementPolicy",
		&returns,
	)
	return returns
}

func (s *statefulSet) PodMetadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		s,
		"podMetadata",
		&returns,
	)
	return returns
}

func (s *statefulSet) Replicas() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"replicas",
		&returns,
	)
	return returns
}

func (s *statefulSet) RestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		s,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (s *statefulSet) ServiceAccount() IServiceAccount {
	var returns IServiceAccount
	_jsii_.Get(
		s,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (s *statefulSet) Volumes() []Volume {
	var returns []Volume
	_jsii_.Get(
		s,
		"volumes",
		&returns,
	)
	return returns
}

func (s *statefulSet) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		s,
		"metadata",
		&returns,
	)
	return returns
}

func (s *statefulSet) Name() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
	)
	return returns
}


func NewStatefulSet(scope constructs.Construct, id string, props StatefulSetProps) StatefulSet {
	_init_.Initialize()
	s := statefulSet{}

	_jsii_.Create(
		"cdk8s-plus-17.StatefulSet",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&s,
	)
	return &s
}

// Add a container to the pod.
// Experimental.
func (s *statefulSet) AddContainer(container ContainerProps) Container {
	var returns Container
	_jsii_.Invoke(
		s,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
	)
	return returns
}

// Add a volume to the pod.
// Experimental.
func (s *statefulSet) AddVolume(volume Volume) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
	)
}

// Configure a label selector to this deployment.
//
// Pods that have the label will be selected by deployments configured with this spec.
// Experimental.
func (s *statefulSet) SelectByLabel(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"selectByLabel",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *statefulSet) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *statefulSet) OnSynthesize(session constructs.ISynthesisSession) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *statefulSet) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *statefulSet) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Properties for initialization of `StatefulSet`.
// Experimental.
type StatefulSetProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadata `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerProps `json:"containers"`
	// Restart policy for all containers within the pod.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
	//
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// A service account provides an identity for processes that run in a Pod.
	//
	// When you (a human) access the cluster (for example, using kubectl), you are
	// authenticated by the apiserver as a particular User Account (currently this
	// is usually admin, unless your cluster administrator has customized your
	// cluster). Processes in containers inside pods can also contact the
	// apiserver. When they do, they are authenticated as a particular Service
	// Account (for example, default).
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	//
	// Experimental.
	ServiceAccount IServiceAccount `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []Volume `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadata `json:"podMetadata"`
	// Service to associate with the statefulset.
	// Experimental.
	Service Service `json:"service"`
	// Automatically allocates a pod selector for this statefulset.
	//
	// If this is set to `false` you must define your selector through
	// `statefulset.podMetadata.addLabel()` and `statefulset.selectByLabel()`.
	// Experimental.
	DefaultSelector bool `json:"defaultSelector"`
	// Pod management policy to use for this statefulset.
	// Experimental.
	PodManagementPolicy PodManagementPolicy `json:"podManagementPolicy"`
	// Number of desired pods.
	// Experimental.
	Replicas float64 `json:"replicas"`
}

// ToResourceProps is a convenience function to obtain a new ResourceProps from this StatefulSetProps.
func (s *StatefulSetProps) ToResourceProps() ResourceProps {
	return ResourceProps {
		Metadata: s.Metadata,
	}
}

// ToPodSpecProps is a convenience function to obtain a new PodSpecProps from this StatefulSetProps.
func (s *StatefulSetProps) ToPodSpecProps() PodSpecProps {
	return PodSpecProps {
		Containers: s.Containers,
		RestartPolicy: s.RestartPolicy,
		ServiceAccount: s.ServiceAccount,
		Volumes: s.Volumes,
	}
}

// ToPodTemplateProps is a convenience function to obtain a new PodTemplateProps from this StatefulSetProps.
func (s *StatefulSetProps) ToPodTemplateProps() PodTemplateProps {
	return PodTemplateProps {
		Containers: s.Containers,
		RestartPolicy: s.RestartPolicy,
		ServiceAccount: s.ServiceAccount,
		Volumes: s.Volumes,
		PodMetadata: s.PodMetadata,
	}
}

// Volume represents a named volume in a pod that may be accessed by any container in the pod.
//
// Docker also has a concept of volumes, though it is somewhat looser and less
// managed. In Docker, a volume is simply a directory on disk or in another
// Container. Lifetimes are not managed and until very recently there were only
// local-disk-backed volumes. Docker now provides volume drivers, but the
// functionality is very limited for now (e.g. as of Docker 1.7 only one volume
// driver is allowed per Container and there is no way to pass parameters to
// volumes).
//
// A Kubernetes volume, on the other hand, has an explicit lifetime - the same
// as the Pod that encloses it. Consequently, a volume outlives any Containers
// that run within the Pod, and data is preserved across Container restarts. Of
// course, when a Pod ceases to exist, the volume will cease to exist, too.
// Perhaps more importantly than this, Kubernetes supports many types of
// volumes, and a Pod can use any number of them simultaneously.
//
// At its core, a volume is just a directory, possibly with some data in it,
// which is accessible to the Containers in a Pod. How that directory comes to
// be, the medium that backs it, and the contents of it are determined by the
// particular volume type used.
//
// To use a volume, a Pod specifies what volumes to provide for the Pod (the
// .spec.volumes field) and where to mount those into Containers (the
// .spec.containers[*].volumeMounts field).
//
// A process in a container sees a filesystem view composed from their Docker
// image and volumes. The Docker image is at the root of the filesystem
// hierarchy, and any volumes are mounted at the specified paths within the
// image. Volumes can not mount onto other volumes
// Experimental.
type Volume interface {
	Name() string
}

// The jsii proxy struct for Volume
type volume struct {
	_ byte // padding
}

func (v *volume) Name() string {
	var returns string
	_jsii_.Get(
		v,
		"name",
		&returns,
	)
	return returns
}


func NewVolume(name string, config interface{}) Volume {
	_init_.Initialize()
	v := volume{}

	_jsii_.Create(
		"cdk8s-plus-17.Volume",
		[]interface{}{name, config},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&v,
	)
	return &v
}

// Populate the volume from a ConfigMap.
//
// The configMap resource provides a way to inject configuration data into
// Pods. The data stored in a ConfigMap object can be referenced in a volume
// of type configMap and then consumed by containerized applications running
// in a Pod.
//
// When referencing a configMap object, you can simply provide its name in the
// volume to reference it. You can also customize the path to use for a
// specific entry in the ConfigMap.
// Experimental.
func Volume_FromConfigMap(configMap IConfigMap, options ConfigMapVolumeOptions) Volume {
	_init_.Initialize()
	var returns Volume
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.Volume",
		"fromConfigMap",
		[]interface{}{configMap, options},
		true,
		&returns,
	)
	return returns
}

// An emptyDir volume is first created when a Pod is assigned to a Node, and exists as long as that Pod is running on that node.
//
// As the name says, it is
// initially empty. Containers in the Pod can all read and write the same
// files in the emptyDir volume, though that volume can be mounted at the same
// or different paths in each Container. When a Pod is removed from a node for
// any reason, the data in the emptyDir is deleted forever.
// See: http://kubernetes.io/docs/user-guide/volumes#emptydir
//
// Experimental.
func Volume_FromEmptyDir(name string, options EmptyDirVolumeOptions) Volume {
	_init_.Initialize()
	var returns Volume
	_jsii_.StaticInvoke(
		"cdk8s-plus-17.Volume",
		"fromEmptyDir",
		[]interface{}{name, options},
		true,
		&returns,
	)
	return returns
}

// Mount a volume from the pod to the container.
// Experimental.
type VolumeMount struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	//
	// This field is beta in 1.10.
	// Experimental.
	Propagation MountPropagation `json:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	// Experimental.
	ReadOnly bool `json:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	// Experimental.
	SubPath string `json:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root). SubPathExpr and SubPath are mutually exclusive. This field
	// is beta in 1.15.
	//
	// `subPathExpr` and `subPath` are mutually exclusive. This field is beta in
	// 1.15.
	// Experimental.
	SubPathExpr string `json:"subPathExpr"`
	// Path within the container at which the volume should be mounted.
	//
	// Must not
	// contain ':'.
	// Experimental.
	Path string `json:"path"`
	// The volume to mount.
	// Experimental.
	Volume Volume `json:"volume"`
}

// ToMountOptions is a convenience function to obtain a new MountOptions from this VolumeMount.
func (v *VolumeMount) ToMountOptions() MountOptions {
	return MountOptions {
		Propagation: v.Propagation,
		ReadOnly: v.ReadOnly,
		SubPath: v.SubPath,
		SubPathExpr: v.SubPathExpr,
	}
}

