// High level abstractions on top of cdk8s
package cdk8splus17

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/awslabs/cdk8s-go/cdk8splus17/jsii"
	"reflect"
	"github.com/awslabs/cdk8s-go/cdk8s"
	"github.com/aws/constructs-go/constructs/v3"
)

// AddDirectoryOptionsIface is the public interface for the custom type AddDirectoryOptions
// Experimental.
type AddDirectoryOptionsIface interface {
	GetExclude() []string
	GetKeyPrefix() string
}

// Options for `configmap.addDirectory()`.
// Experimental.
// Struct proxy
type AddDirectoryOptions struct {
	// Glob patterns to exclude when adding files.
	// Experimental.
	Exclude []string `json:"exclude"`
	// A prefix to add to all keys in the config map.
	// Experimental.
	KeyPrefix string `json:"keyPrefix"`
}

func (a *AddDirectoryOptions) GetExclude() []string {
	var returns []string
	_jsii_.Get(
		a,
		"exclude",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (a *AddDirectoryOptions) GetKeyPrefix() string {
	var returns string
	_jsii_.Get(
		a,
		"keyPrefix",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// CommandProbeOptionsIface is the public interface for the custom type CommandProbeOptions
// Experimental.
type CommandProbeOptionsIface interface {
	GetFailureThreshold() float64
	GetInitialDelaySeconds() cdk8s.DurationIface
	GetPeriodSeconds() cdk8s.DurationIface
	GetSuccessThreshold() float64
	GetTimeoutSeconds() cdk8s.DurationIface
}

// Options for `Probe.fromCommand()`.
// Experimental.
// Struct proxy
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
	InitialDelaySeconds cdk8s.DurationIface `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.DurationIface `json:"periodSeconds"`
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
	TimeoutSeconds cdk8s.DurationIface `json:"timeoutSeconds"`
}

func (c *CommandProbeOptions) GetFailureThreshold() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"failureThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *CommandProbeOptions) GetInitialDelaySeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		c,
		"initialDelaySeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (c *CommandProbeOptions) GetPeriodSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		c,
		"periodSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (c *CommandProbeOptions) GetSuccessThreshold() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"successThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *CommandProbeOptions) GetTimeoutSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		c,
		"timeoutSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type ConfigMapIface interface {
	constructs.IConstructIface
	IResourceIface
	IConfigMapIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetBinaryData() map[string]string
	GetData() map[string]string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddBinaryData(key string, value string)
	AddData(key string, value string)
	AddDirectory(localDir string, options AddDirectoryOptionsIface)
	AddFile(localFile string, key string)
}

// ConfigMap holds configuration data for pods to consume.
// Experimental.
// Struct proxy
type ConfigMap struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// The binary data associated with this config map.
	//
	// Returns a copy. To add data records, use `addBinaryData()` or `addData()`.
	// Experimental.
	BinaryData map[string]string `json:"binaryData"`
	// The data associated with this config map.
	//
	// Returns an copy. To add data records, use `addData()` or `addBinaryData()`.
	// Experimental.
	Data map[string]string `json:"data"`
}

func (c *ConfigMap) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		c,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMap) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		c,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMap) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConfigMap) GetBinaryData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"binaryData",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMap) GetData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"data",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


func NewConfigMap(scope constructs.ConstructIface, id string, props ConfigMapPropsIface) ConfigMapIface {
	_init_.Initialize()
	self := ConfigMap{}
	_jsii_.Create(
		"cdk8s-plus-17.ConfigMap",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IConfigMap"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ConfigMap_FromConfigMapName(name string) IConfigMapIface {
	_init_.Initialize()
	var returns IConfigMapIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.ConfigMap",
		"fromConfigMapName",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IConfigMapIface)(nil)).Elem(): reflect.TypeOf((*IConfigMap)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMap) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConfigMap) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConfigMap) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		c,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMap) ToString() string {
	var returns string
	_jsii_.Invoke(
		c,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConfigMap) AddBinaryData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addBinaryData",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConfigMap) AddData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addData",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConfigMap) AddDirectory(localDir string, options AddDirectoryOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDirectory",
		[]interface{}{localDir, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *ConfigMap) AddFile(localFile string, key string) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addFile",
		[]interface{}{localFile, key},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ConfigMapPropsIface is the public interface for the custom type ConfigMapProps
// Experimental.
type ConfigMapPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetBinaryData() map[string]string
	GetData() map[string]string
}

// Properties for initialization of `ConfigMap`.
// Experimental.
// Struct proxy
type ConfigMapProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
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

func (c *ConfigMapProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		c,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMapProps) GetBinaryData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"binaryData",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMapProps) GetData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"data",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}


// ConfigMapVolumeOptionsIface is the public interface for the custom type ConfigMapVolumeOptions
// Experimental.
type ConfigMapVolumeOptionsIface interface {
	GetDefaultMode() float64
	GetItems() map[string]PathMappingIface
	GetName() string
	GetOptional() bool
}

// Options for the ConfigMap-based volume.
// Experimental.
// Struct proxy
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
	Items map[string]PathMappingIface `json:"items"`
	// The volume name.
	// Experimental.
	Name string `json:"name"`
	// Specify whether the ConfigMap or its keys must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

func (c *ConfigMapVolumeOptions) GetDefaultMode() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"defaultMode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConfigMapVolumeOptions) GetItems() map[string]PathMappingIface {
	var returns map[string]PathMappingIface
	_jsii_.Get(
		c,
		"items",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PathMappingIface)(nil)).Elem(): reflect.TypeOf((*PathMapping)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ConfigMapVolumeOptions) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ConfigMapVolumeOptions) GetOptional() bool {
	var returns bool
	_jsii_.Get(
		c,
		"optional",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type ContainerIface interface {
	GetEnv() map[string]EnvValueIface
	GetImage() string
	GetImagePullPolicy() ImagePullPolicy
	GetMounts() []VolumeMountIface
	GetName() string
	GetArgs() []string
	GetCommand() []string
	GetPort() float64
	GetWorkingDir() string
	AddEnv(name string, value EnvValueIface)
	Mount(path string, volume VolumeIface, options MountOptionsIface)
}

// A single application container that you want to run within a pod.
// Experimental.
// Struct proxy
type Container struct {
	// The environment variables for this container.
	//
	// Returns a copy. To add environment variables use `addEnv()`.
	// Experimental.
	Env map[string]EnvValueIface `json:"env"`
	// The container image.
	// Experimental.
	Image string `json:"image"`
	// Image pull policy for this container.
	// Experimental.
	ImagePullPolicy ImagePullPolicy `json:"imagePullPolicy"`
	// Volume mounts configured for this container.
	// Experimental.
	Mounts []VolumeMountIface `json:"mounts"`
	// The name of the container.
	// Experimental.
	Name string `json:"name"`
	// Arguments to the entrypoint.
	//
	// Returns: a copy of the arguments array, cannot be modified.
	// Experimental.
	Args []string `json:"args"`
	// Entrypoint array (the command to execute when the container starts).
	//
	// Returns: a copy of the entrypoint array, cannot be modified
	// Experimental.
	Command []string `json:"command"`
	// The port this container exposes.
	// Experimental.
	Port float64 `json:"port"`
	// The working directory inside the container.
	// Experimental.
	WorkingDir string `json:"workingDir"`
}

func (c *Container) GetEnv() map[string]EnvValueIface {
	var returns map[string]EnvValueIface
	_jsii_.Get(
		c,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Container) GetImage() string {
	var returns string
	_jsii_.Get(
		c,
		"image",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *Container) GetImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		c,
		"imagePullPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(): reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Container) GetMounts() []VolumeMountIface {
	var returns []VolumeMountIface
	_jsii_.Get(
		c,
		"mounts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeMountIface)(nil)).Elem(): reflect.TypeOf((*VolumeMount)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Container) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *Container) GetArgs() []string {
	var returns []string
	_jsii_.Get(
		c,
		"args",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Container) GetCommand() []string {
	var returns []string
	_jsii_.Get(
		c,
		"command",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *Container) GetPort() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *Container) GetWorkingDir() string {
	var returns string
	_jsii_.Get(
		c,
		"workingDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewContainer(props ContainerPropsIface) ContainerIface {
	_init_.Initialize()
	self := Container{}
	_jsii_.Create(
		"cdk8s-plus-17.Container",
		[]interface{}{props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (c *Container) AddEnv(name string, value EnvValueIface) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addEnv",
		[]interface{}{name, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (c *Container) Mount(path string, volume VolumeIface, options MountOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"mount",
		[]interface{}{path, volume, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ContainerPropsIface is the public interface for the custom type ContainerProps
// Experimental.
type ContainerPropsIface interface {
	GetImage() string
	GetArgs() []string
	GetCommand() []string
	GetEnv() map[string]EnvValueIface
	GetImagePullPolicy() ImagePullPolicy
	GetLiveness() ProbeIface
	GetName() string
	GetPort() float64
	GetReadiness() ProbeIface
	GetStartup() ProbeIface
	GetVolumeMounts() []VolumeMountIface
	GetWorkingDir() string
}

// Properties for creating a container.
// Experimental.
// Struct proxy
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
	Env map[string]EnvValueIface `json:"env"`
	// Image pull policy for this container.
	// Experimental.
	ImagePullPolicy ImagePullPolicy `json:"imagePullPolicy"`
	// Periodic probe of container liveness.
	//
	// Container will be restarted if the probe fails.
	// Experimental.
	Liveness ProbeIface `json:"liveness"`
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
	Readiness ProbeIface `json:"readiness"`
	// StartupProbe indicates that the Pod has successfully initialized.
	//
	// If specified, no other probes are executed until this completes successfully
	// Experimental.
	Startup ProbeIface `json:"startup"`
	// Pod volumes to mount into the container's filesystem.
	//
	// Cannot be updated.
	// Experimental.
	VolumeMounts []VolumeMountIface `json:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	// Experimental.
	WorkingDir string `json:"workingDir"`
}

func (c *ContainerProps) GetImage() string {
	var returns string
	_jsii_.Get(
		c,
		"image",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ContainerProps) GetArgs() []string {
	var returns []string
	_jsii_.Get(
		c,
		"args",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetCommand() []string {
	var returns []string
	_jsii_.Get(
		c,
		"command",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetEnv() map[string]EnvValueIface {
	var returns map[string]EnvValueIface
	_jsii_.Get(
		c,
		"env",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetImagePullPolicy() ImagePullPolicy {
	var returns ImagePullPolicy
	_jsii_.Get(
		c,
		"imagePullPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(): reflect.TypeOf((*ImagePullPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetLiveness() ProbeIface {
	var returns ProbeIface
	_jsii_.Get(
		c,
		"liveness",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProbeIface)(nil)).Elem(): reflect.TypeOf((*Probe)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetName() string {
	var returns string
	_jsii_.Get(
		c,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ContainerProps) GetPort() float64 {
	var returns float64
	_jsii_.Get(
		c,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (c *ContainerProps) GetReadiness() ProbeIface {
	var returns ProbeIface
	_jsii_.Get(
		c,
		"readiness",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProbeIface)(nil)).Elem(): reflect.TypeOf((*Probe)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetStartup() ProbeIface {
	var returns ProbeIface
	_jsii_.Get(
		c,
		"startup",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProbeIface)(nil)).Elem(): reflect.TypeOf((*Probe)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetVolumeMounts() []VolumeMountIface {
	var returns []VolumeMountIface
	_jsii_.Get(
		c,
		"volumeMounts",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeMountIface)(nil)).Elem(): reflect.TypeOf((*VolumeMount)(nil)).Elem(),
		},
	)
	return returns
}

func (c *ContainerProps) GetWorkingDir() string {
	var returns string
	_jsii_.Get(
		c,
		"workingDir",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type DeploymentIface interface {
	constructs.IConstructIface
	IResourceIface
	IPodTemplateIface
	IPodSpecIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetContainers() []ContainerIface
	GetLabelSelector() map[string]string
	GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetReplicas() float64
	GetVolumes() []VolumeIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
	Expose(port float64, options ExposeOptionsIface) ServiceIface
	SelectByLabel(key string, value string)
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
// Struct proxy
type Deployment struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// The labels this deployment will match against in order to select pods.
	//
	// Returns a a copy. Use `selectByLabel()` to add labels.
	// Experimental.
	LabelSelector map[string]string `json:"labelSelector"`
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataDefinitionIface `json:"podMetadata"`
	// Number of desired pods.
	// Experimental.
	Replicas float64 `json:"replicas"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
}

func (d *Deployment) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		d,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		d,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetName() string {
	var returns string
	_jsii_.Get(
		d,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Deployment) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		d,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetLabelSelector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		d,
		"labelSelector",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		d,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetReplicas() float64 {
	var returns float64
	_jsii_.Get(
		d,
		"replicas",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Deployment) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		d,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		d,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		d,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}


func NewDeployment(scope constructs.ConstructIface, id string, props DeploymentPropsIface) DeploymentIface {
	_init_.Initialize()
	self := Deployment{}
	_jsii_.Create(
		"cdk8s-plus-17.Deployment",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (d *Deployment) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Deployment) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Deployment) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		d,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) ToString() string {
	var returns string
	_jsii_.Invoke(
		d,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *Deployment) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		d,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (d *Deployment) Expose(port float64, options ExposeOptionsIface) ServiceIface {
	var returns ServiceIface
	_jsii_.Invoke(
		d,
		"expose",
		[]interface{}{port, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServiceIface)(nil)).Elem(): reflect.TypeOf((*Service)(nil)).Elem(),
		},
	)
	return returns
}

func (d *Deployment) SelectByLabel(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"selectByLabel",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// DeploymentPropsIface is the public interface for the custom type DeploymentProps
// Experimental.
type DeploymentPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
	GetPodMetadata() cdk8s.ApiObjectMetadataIface
	GetDefaultSelector() bool
	GetReplicas() float64
}

// Properties for initialization of `Deployment`.
// Experimental.
// Struct proxy
type DeploymentProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataIface `json:"podMetadata"`
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

func (d *DeploymentProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		d,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		d,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		d,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		d,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		d,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetPodMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		d,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (d *DeploymentProps) GetDefaultSelector() bool {
	var returns bool
	_jsii_.Get(
		d,
		"defaultSelector",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (d *DeploymentProps) GetReplicas() float64 {
	var returns float64
	_jsii_.Get(
		d,
		"replicas",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// The medium on which to store the volume.
// Experimental.
type EmptyDirMedium string

const (
	EmptyDirMediumDefault EmptyDirMedium = "DEFAULT"
	EmptyDirMediumMemory EmptyDirMedium = "MEMORY"
)

// EmptyDirVolumeOptionsIface is the public interface for the custom type EmptyDirVolumeOptions
// Experimental.
type EmptyDirVolumeOptionsIface interface {
	GetMedium() EmptyDirMedium
	GetSizeLimit() cdk8s.SizeIface
}

// Options for volumes populated with an empty directory.
// Experimental.
// Struct proxy
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
	SizeLimit cdk8s.SizeIface `json:"sizeLimit"`
}

func (e *EmptyDirVolumeOptions) GetMedium() EmptyDirMedium {
	var returns EmptyDirMedium
	_jsii_.Get(
		e,
		"medium",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(): reflect.TypeOf((*EmptyDirMedium)(nil)).Elem(),
		},
	)
	return returns
}

func (e *EmptyDirVolumeOptions) GetSizeLimit() cdk8s.SizeIface {
	var returns cdk8s.SizeIface
	_jsii_.Get(
		e,
		"sizeLimit",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.SizeIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Size)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type EnvValueIface interface {
	GetValue() interface{}
	GetValueFrom() interface{}
}

// Utility class for creating reading env values from various sources.
// Experimental.
// Struct proxy
type EnvValue struct {
	// Experimental.
	Value interface{} `json:"value"`
	// Experimental.
	ValueFrom interface{} `json:"valueFrom"`
}

func (e *EnvValue) GetValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		e,
		"value",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (e *EnvValue) GetValueFrom() interface{} {
	var returns interface{}
	_jsii_.Get(
		e,
		"valueFrom",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func EnvValue_FromConfigMap(configMap IConfigMapIface, key string, options EnvValueFromConfigMapOptionsIface) EnvValueIface {
	_init_.Initialize()
	var returns EnvValueIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.EnvValue",
		"fromConfigMap",
		[]interface{}{configMap, key, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

func EnvValue_FromProcess(key string, options EnvValueFromProcessOptionsIface) EnvValueIface {
	_init_.Initialize()
	var returns EnvValueIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.EnvValue",
		"fromProcess",
		[]interface{}{key, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

func EnvValue_FromSecretValue(secretValue SecretValueIface, options EnvValueFromSecretOptionsIface) EnvValueIface {
	_init_.Initialize()
	var returns EnvValueIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.EnvValue",
		"fromSecretValue",
		[]interface{}{secretValue, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

func EnvValue_FromValue(value string) EnvValueIface {
	_init_.Initialize()
	var returns EnvValueIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.EnvValue",
		"fromValue",
		[]interface{}{value},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*EnvValueIface)(nil)).Elem(): reflect.TypeOf((*EnvValue)(nil)).Elem(),
		},
	)
	return returns
}

// EnvValueFromConfigMapOptionsIface is the public interface for the custom type EnvValueFromConfigMapOptions
// Experimental.
type EnvValueFromConfigMapOptionsIface interface {
	GetOptional() bool
}

// Options to specify an envionment variable value from a ConfigMap key.
// Experimental.
// Struct proxy
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

func (e *EnvValueFromConfigMapOptions) GetOptional() bool {
	var returns bool
	_jsii_.Get(
		e,
		"optional",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// EnvValueFromProcessOptionsIface is the public interface for the custom type EnvValueFromProcessOptions
// Experimental.
type EnvValueFromProcessOptionsIface interface {
	GetRequired() bool
}

// Options to specify an environment variable value from the process environment.
// Experimental.
// Struct proxy
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	// Experimental.
	Required bool `json:"required"`
}

func (e *EnvValueFromProcessOptions) GetRequired() bool {
	var returns bool
	_jsii_.Get(
		e,
		"required",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// EnvValueFromSecretOptionsIface is the public interface for the custom type EnvValueFromSecretOptions
// Experimental.
type EnvValueFromSecretOptionsIface interface {
	GetOptional() bool
}

// Options to specify an environment variable value from a Secret.
// Experimental.
// Struct proxy
type EnvValueFromSecretOptions struct {
	// Specify whether the Secret or its key must be defined.
	// Experimental.
	Optional bool `json:"optional"`
}

func (e *EnvValueFromSecretOptions) GetOptional() bool {
	var returns bool
	_jsii_.Get(
		e,
		"optional",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ExposeOptionsIface is the public interface for the custom type ExposeOptions
// Experimental.
type ExposeOptionsIface interface {
	GetName() string
	GetProtocol() Protocol
	GetServiceType() ServiceType
	GetTargetPort() float64
}

// Options for exposing a deployment via a service.
// Experimental.
// Struct proxy
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

func (e *ExposeOptions) GetName() string {
	var returns string
	_jsii_.Get(
		e,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (e *ExposeOptions) GetProtocol() Protocol {
	var returns Protocol
	_jsii_.Get(
		e,
		"protocol",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*Protocol)(nil)).Elem(): reflect.TypeOf((*Protocol)(nil)).Elem(),
		},
	)
	return returns
}

func (e *ExposeOptions) GetServiceType() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		e,
		"serviceType",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServiceType)(nil)).Elem(): reflect.TypeOf((*ServiceType)(nil)).Elem(),
		},
	)
	return returns
}

func (e *ExposeOptions) GetTargetPort() float64 {
	var returns float64
	_jsii_.Get(
		e,
		"targetPort",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// HttpGetProbeOptionsIface is the public interface for the custom type HttpGetProbeOptions
// Experimental.
type HttpGetProbeOptionsIface interface {
	GetFailureThreshold() float64
	GetInitialDelaySeconds() cdk8s.DurationIface
	GetPeriodSeconds() cdk8s.DurationIface
	GetSuccessThreshold() float64
	GetTimeoutSeconds() cdk8s.DurationIface
	GetPort() float64
}

// Options for `Probe.fromHttpGet()`.
// Experimental.
// Struct proxy
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
	InitialDelaySeconds cdk8s.DurationIface `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.DurationIface `json:"periodSeconds"`
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
	TimeoutSeconds cdk8s.DurationIface `json:"timeoutSeconds"`
	// The TCP port to use when sending the GET request.
	// Experimental.
	Port float64 `json:"port"`
}

func (h *HttpGetProbeOptions) GetFailureThreshold() float64 {
	var returns float64
	_jsii_.Get(
		h,
		"failureThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HttpGetProbeOptions) GetInitialDelaySeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		h,
		"initialDelaySeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (h *HttpGetProbeOptions) GetPeriodSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		h,
		"periodSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (h *HttpGetProbeOptions) GetSuccessThreshold() float64 {
	var returns float64
	_jsii_.Get(
		h,
		"successThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HttpGetProbeOptions) GetTimeoutSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		h,
		"timeoutSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (h *HttpGetProbeOptions) GetPort() float64 {
	var returns float64
	_jsii_.Get(
		h,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Represents a config map.
// Experimental.
type IConfigMapIface interface {
	IResourceIface
}

type IConfigMap struct {}

func (i *IConfigMap) GetName() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Represents a resource that can be configured with a kuberenets pod spec. (e.g `Deployment`, `Job`, `Pod`, ...).
//
// Use the `PodSpec` class as an implementation helper.
// Experimental.
type IPodSpecIface interface {
	// Add a container to the pod.
	// Experimental.
	AddContainer(container ContainerPropsIface) ContainerIface
	// Add a volume to the pod.
	// Experimental.
	AddVolume(volume VolumeIface)
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	GetContainers() []ContainerIface
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	GetVolumes() []VolumeIface
	// Restart policy for all containers within the pod.
	// Experimental.
	GetRestartPolicy() RestartPolicy
	// The service account used to run this pod.
	// Experimental.
	GetServiceAccount() IServiceAccountIface
}

type IPodSpec struct {}

func (i *IPodSpec) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		i,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodSpec) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IPodSpec) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		i,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodSpec) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		i,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodSpec) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		i,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodSpec) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		i,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

// Represents a resource that can be configured with a kuberenets pod template. (e.g `Deployment`, `Job`, ...).
//
// Use the `PodTemplate` class as an implementation helper.
// Experimental.
type IPodTemplateIface interface {
	IPodSpecIface
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface
}

type IPodTemplate struct {}

func (i *IPodTemplate) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		i,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodTemplate) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IPodTemplate) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		i,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodTemplate) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		i,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodTemplate) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		i,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodTemplate) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		i,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IPodTemplate) GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		i,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

// Represents a resource.
// Experimental.
type IResourceIface interface {
	// The Kubernetes name of this resource.
	// Experimental.
	GetName() string
}

type IResource struct {}

func (i *IResource) GetName() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Experimental.
type ISecretIface interface {
	IResourceIface
}

type ISecret struct {}

func (i *ISecret) GetName() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Experimental.
type IServiceAccountIface interface {
	IResourceIface
}

type IServiceAccount struct {}

func (i *IServiceAccount) GetName() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// Experimental.
type ImagePullPolicy string

const (
	ImagePullPolicyAlways ImagePullPolicy = "ALWAYS"
	ImagePullPolicyIfNotPresent ImagePullPolicy = "IF_NOT_PRESENT"
	ImagePullPolicyNever ImagePullPolicy = "NEVER"
)

// Class interface
type IngressV1Beta1Iface interface {
	constructs.IConstructIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddDefaultBackend(backend IngressV1Beta1BackendIface)
	AddHostDefaultBackend(host string, backend IngressV1Beta1BackendIface)
	AddHostRule(host string, path string, backend IngressV1Beta1BackendIface)
	AddRule(path string, backend IngressV1Beta1BackendIface)
	AddRules(rules IngressV1Beta1RuleIface)
}

// Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend.
//
// An Ingress can be configured to give services
// externally-reachable urls, load balance traffic, terminate SSL, offer name
// based virtual hosting etc.
// Experimental.
// Struct proxy
type IngressV1Beta1 struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
}

func (i *IngressV1Beta1) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		i,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		i,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1) GetName() string {
	var returns string
	_jsii_.Get(
		i,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewIngressV1Beta1(scope constructs.ConstructIface, id string, props IngressV1Beta1PropsIface) IngressV1Beta1Iface {
	_init_.Initialize()
	self := IngressV1Beta1{}
	_jsii_.Create(
		"cdk8s-plus-17.IngressV1Beta1",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (i *IngressV1Beta1) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		i,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1) ToString() string {
	var returns string
	_jsii_.Invoke(
		i,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IngressV1Beta1) AddDefaultBackend(backend IngressV1Beta1BackendIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addDefaultBackend",
		[]interface{}{backend},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) AddHostDefaultBackend(host string, backend IngressV1Beta1BackendIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addHostDefaultBackend",
		[]interface{}{host, backend},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) AddHostRule(host string, path string, backend IngressV1Beta1BackendIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addHostRule",
		[]interface{}{host, path, backend},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) AddRule(path string, backend IngressV1Beta1BackendIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addRule",
		[]interface{}{path, backend},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (i *IngressV1Beta1) AddRules(rules IngressV1Beta1RuleIface) {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"addRules",
		[]interface{}{rules},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type IngressV1Beta1BackendIface interface {
}

// The backend for an ingress path.
// Experimental.
// Struct proxy
type IngressV1Beta1Backend struct {
}

func IngressV1Beta1Backend_FromService(service ServiceIface, options ServiceIngressV1BetaBackendOptionsIface) IngressV1Beta1BackendIface {
	_init_.Initialize()
	var returns IngressV1Beta1BackendIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.IngressV1Beta1Backend",
		"fromService",
		[]interface{}{service, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IngressV1Beta1BackendIface)(nil)).Elem(): reflect.TypeOf((*IngressV1Beta1Backend)(nil)).Elem(),
		},
	)
	return returns
}

// IngressV1Beta1PropsIface is the public interface for the custom type IngressV1Beta1Props
// Experimental.
type IngressV1Beta1PropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetDefaultBackend() IngressV1Beta1BackendIface
	GetRules() []IngressV1Beta1RuleIface
}

// Properties for `Ingress`.
// Experimental.
// Struct proxy
type IngressV1Beta1Props struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// The default backend services requests that do not match any rule.
	//
	// Using this option or the `addDefaultBackend()` method is equivalent to
	// adding a rule with both `path` and `host` undefined.
	// Experimental.
	DefaultBackend IngressV1Beta1BackendIface `json:"defaultBackend"`
	// Routing rules for this ingress.
	//
	// Each rule must define an `IngressBackend` that will receive the requests
	// that match this rule. If both `host` and `path` are not specifiec, this
	// backend will be used as the default backend of the ingress.
	//
	// You can also add rules later using `addRule()`, `addHostRule()`,
	// `addDefaultBackend()` and `addHostDefaultBackend()`.
	// Experimental.
	Rules []IngressV1Beta1RuleIface `json:"rules"`
}

func (i *IngressV1Beta1Props) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		i,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1Props) GetDefaultBackend() IngressV1Beta1BackendIface {
	var returns IngressV1Beta1BackendIface
	_jsii_.Get(
		i,
		"defaultBackend",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IngressV1Beta1BackendIface)(nil)).Elem(): reflect.TypeOf((*IngressV1Beta1Backend)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1Props) GetRules() []IngressV1Beta1RuleIface {
	var returns []IngressV1Beta1RuleIface
	_jsii_.Get(
		i,
		"rules",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IngressV1Beta1RuleIface)(nil)).Elem(): reflect.TypeOf((*IngressV1Beta1Rule)(nil)).Elem(),
		},
	)
	return returns
}


// IngressV1Beta1RuleIface is the public interface for the custom type IngressV1Beta1Rule
// Experimental.
type IngressV1Beta1RuleIface interface {
	GetBackend() IngressV1Beta1BackendIface
	GetHost() string
	GetPath() string
}

// Represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match,
// then routed to the backend associated with the matching path.
// Experimental.
// Struct proxy
type IngressV1Beta1Rule struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	// Experimental.
	Backend IngressV1Beta1BackendIface `json:"backend"`
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

func (i *IngressV1Beta1Rule) GetBackend() IngressV1Beta1BackendIface {
	var returns IngressV1Beta1BackendIface
	_jsii_.Get(
		i,
		"backend",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IngressV1Beta1BackendIface)(nil)).Elem(): reflect.TypeOf((*IngressV1Beta1Backend)(nil)).Elem(),
		},
	)
	return returns
}

func (i *IngressV1Beta1Rule) GetHost() string {
	var returns string
	_jsii_.Get(
		i,
		"host",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (i *IngressV1Beta1Rule) GetPath() string {
	var returns string
	_jsii_.Get(
		i,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type JobIface interface {
	constructs.IConstructIface
	IResourceIface
	IPodTemplateIface
	IPodSpecIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetContainers() []ContainerIface
	GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetVolumes() []VolumeIface
	GetActiveDeadline() cdk8s.DurationIface
	GetBackoffLimit() float64
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetTtlAfterFinished() cdk8s.DurationIface
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
}

// A Job creates one or more Pods and ensures that a specified number of them successfully terminate.
//
// As pods successfully complete,
// the Job tracks the successful completions. When a specified number of successful completions is reached, the task (ie, Job) is complete.
// Deleting a Job will clean up the Pods it created. A simple case is to create one Job object in order to reliably run one Pod to completion.
// The Job object will start a new Pod if the first Pod fails or is deleted (for example due to a node hardware failure or a node reboot).
// You can also use a Job to run multiple Pods in parallel.
// Experimental.
// Struct proxy
type Job struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataDefinitionIface `json:"podMetadata"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Duration before job is terminated.
	//
	// If undefined, there is no deadline.
	// Experimental.
	ActiveDeadline cdk8s.DurationIface `json:"activeDeadline"`
	// Number of retries before marking failed.
	// Experimental.
	BackoffLimit float64 `json:"backoffLimit"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// TTL before the job is deleted after it is finished.
	// Experimental.
	TtlAfterFinished cdk8s.DurationIface `json:"ttlAfterFinished"`
}

func (j *Job) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		j,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		j,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetName() string {
	var returns string
	_jsii_.Get(
		j,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *Job) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		j,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		j,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetActiveDeadline() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		j,
		"activeDeadline",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetBackoffLimit() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *Job) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) GetTtlAfterFinished() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		j,
		"ttlAfterFinished",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}


func NewJob(scope constructs.ConstructIface, id string, props JobPropsIface) JobIface {
	_init_.Initialize()
	self := Job{}
	_jsii_.Create(
		"cdk8s-plus-17.Job",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (j *Job) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Job) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (j *Job) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		j,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) ToString() string {
	var returns string
	_jsii_.Invoke(
		j,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *Job) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		j,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (j *Job) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		j,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// JobPropsIface is the public interface for the custom type JobProps
// Experimental.
type JobPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
	GetPodMetadata() cdk8s.ApiObjectMetadataIface
	GetActiveDeadline() cdk8s.DurationIface
	GetBackoffLimit() float64
	GetTtlAfterFinished() cdk8s.DurationIface
}

// Properties for initialization of `Job`.
// Experimental.
// Struct proxy
type JobProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataIface `json:"podMetadata"`
	// Specifies the duration the job may be active before the system tries to terminate it.
	// Experimental.
	ActiveDeadline cdk8s.DurationIface `json:"activeDeadline"`
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
	TtlAfterFinished cdk8s.DurationIface `json:"ttlAfterFinished"`
}

func (j *JobProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		j,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		j,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		j,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetPodMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		j,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetActiveDeadline() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		j,
		"activeDeadline",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (j *JobProps) GetBackoffLimit() float64 {
	var returns float64
	_jsii_.Get(
		j,
		"backoffLimit",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (j *JobProps) GetTtlAfterFinished() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		j,
		"ttlAfterFinished",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}


// MountOptionsIface is the public interface for the custom type MountOptions
// Experimental.
type MountOptionsIface interface {
	GetPropagation() MountPropagation
	GetReadOnly() bool
	GetSubPath() string
	GetSubPathExpr() string
}

// Options for mounts.
// Experimental.
// Struct proxy
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

func (m *MountOptions) GetPropagation() MountPropagation {
	var returns MountPropagation
	_jsii_.Get(
		m,
		"propagation",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MountPropagation)(nil)).Elem(): reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		},
	)
	return returns
}

func (m *MountOptions) GetReadOnly() bool {
	var returns bool
	_jsii_.Get(
		m,
		"readOnly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MountOptions) GetSubPath() string {
	var returns string
	_jsii_.Get(
		m,
		"subPath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (m *MountOptions) GetSubPathExpr() string {
	var returns string
	_jsii_.Get(
		m,
		"subPathExpr",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Experimental.
type MountPropagation string

const (
	MountPropagationNone MountPropagation = "NONE"
	MountPropagationHostToContainer MountPropagation = "HOST_TO_CONTAINER"
	MountPropagationBidirectional MountPropagation = "BIDIRECTIONAL"
)

// PathMappingIface is the public interface for the custom type PathMapping
// Experimental.
type PathMappingIface interface {
	GetPath() string
	GetMode() float64
}

// Maps a string key to a path within a volume.
// Experimental.
// Struct proxy
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

func (p *PathMapping) GetPath() string {
	var returns string
	_jsii_.Get(
		p,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *PathMapping) GetMode() float64 {
	var returns float64
	_jsii_.Get(
		p,
		"mode",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type PodIface interface {
	constructs.IConstructIface
	IResourceIface
	IPodSpecIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetContainers() []ContainerIface
	GetVolumes() []VolumeIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
}

// Pod is a collection of containers that can run on a host.
//
// This resource is
// created by clients and scheduled onto hosts.
// Experimental.
// Struct proxy
type Pod struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
}

func (p *Pod) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		p,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		p,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) GetName() string {
	var returns string
	_jsii_.Get(
		p,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pod) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}


func NewPod(scope constructs.ConstructIface, id string, props PodPropsIface) PodIface {
	_init_.Initialize()
	self := Pod{}
	_jsii_.Create(
		"cdk8s-plus-17.Pod",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodSpec"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *Pod) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pod) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (p *Pod) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		p,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) ToString() string {
	var returns string
	_jsii_.Invoke(
		p,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *Pod) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *Pod) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
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
	PodManagementPolicyOrderedReady PodManagementPolicy = "ORDERED_READY"
	PodManagementPolicyParallel PodManagementPolicy = "PARALLEL"
)

// PodPropsIface is the public interface for the custom type PodProps
// Experimental.
type PodPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
}

// Properties for initialization of `Pod`.
// Experimental.
// Struct proxy
type PodProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
}

func (p *PodProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		p,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type PodSpecIface interface {
	IPodSpecIface
	GetContainers() []ContainerIface
	GetVolumes() []VolumeIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
}

// Provides read/write capabilities ontop of a `PodSpecProps`.
// Experimental.
// Struct proxy
type PodSpec struct {
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
}

func (p *PodSpec) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpec) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpec) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpec) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}


func NewPodSpec(props PodSpecPropsIface) PodSpecIface {
	_init_.Initialize()
	self := PodSpec{}
	_jsii_.Create(
		"cdk8s-plus-17.PodSpec",
		[]interface{}{props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodSpec"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *PodSpec) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpec) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// PodSpecPropsIface is the public interface for the custom type PodSpecProps
// Experimental.
type PodSpecPropsIface interface {
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
}

// Properties of a `PodSpec`.
// Experimental.
// Struct proxy
type PodSpecProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
}

func (p *PodSpecProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpecProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpecProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodSpecProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type PodTemplateIface interface {
	IPodSpecIface
	IPodTemplateIface
	IPodSpecIface
	GetContainers() []ContainerIface
	GetVolumes() []VolumeIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
}

// Provides read/write capabilities ontop of a `PodTemplateProps`.
// Experimental.
// Struct proxy
type PodTemplate struct {
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataDefinitionIface `json:"podMetadata"`
}

func (p *PodTemplate) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplate) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplate) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplate) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplate) GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		p,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}


func NewPodTemplate(props PodTemplatePropsIface) PodTemplateIface {
	_init_.Initialize()
	self := PodTemplate{}
	_jsii_.Create(
		"cdk8s-plus-17.PodTemplate",
		[]interface{}{props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (p *PodTemplate) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		p,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplate) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		p,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// PodTemplatePropsIface is the public interface for the custom type PodTemplateProps
// Experimental.
type PodTemplatePropsIface interface {
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
	GetPodMetadata() cdk8s.ApiObjectMetadataIface
}

// Properties of a `PodTemplate`.
//
// Adds metadata information on top of the spec.
// Experimental.
// Struct proxy
type PodTemplateProps struct {
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataIface `json:"podMetadata"`
}

func (p *PodTemplateProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		p,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplateProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		p,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplateProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		p,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplateProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		p,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (p *PodTemplateProps) GetPodMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		p,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type ProbeIface interface {
}

// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
// Experimental.
// Struct proxy
type Probe struct {
}

func NewProbe() ProbeIface {
	_init_.Initialize()
	self := Probe{}
	_jsii_.Create(
		"cdk8s-plus-17.Probe",
		[]interface{}{},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Probe_FromCommand(command []string, options CommandProbeOptionsIface) ProbeIface {
	_init_.Initialize()
	var returns ProbeIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.Probe",
		"fromCommand",
		[]interface{}{command, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProbeIface)(nil)).Elem(): reflect.TypeOf((*Probe)(nil)).Elem(),
		},
	)
	return returns
}

func Probe_FromHttpGet(path string, options HttpGetProbeOptionsIface) ProbeIface {
	_init_.Initialize()
	var returns ProbeIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.Probe",
		"fromHttpGet",
		[]interface{}{path, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ProbeIface)(nil)).Elem(): reflect.TypeOf((*Probe)(nil)).Elem(),
		},
	)
	return returns
}

// ProbeOptionsIface is the public interface for the custom type ProbeOptions
// Experimental.
type ProbeOptionsIface interface {
	GetFailureThreshold() float64
	GetInitialDelaySeconds() cdk8s.DurationIface
	GetPeriodSeconds() cdk8s.DurationIface
	GetSuccessThreshold() float64
	GetTimeoutSeconds() cdk8s.DurationIface
}

// Probe options.
// Experimental.
// Struct proxy
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
	InitialDelaySeconds cdk8s.DurationIface `json:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Experimental.
	PeriodSeconds cdk8s.DurationIface `json:"periodSeconds"`
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
	TimeoutSeconds cdk8s.DurationIface `json:"timeoutSeconds"`
}

func (p *ProbeOptions) GetFailureThreshold() float64 {
	var returns float64
	_jsii_.Get(
		p,
		"failureThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProbeOptions) GetInitialDelaySeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		p,
		"initialDelaySeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (p *ProbeOptions) GetPeriodSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		p,
		"periodSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}

func (p *ProbeOptions) GetSuccessThreshold() float64 {
	var returns float64
	_jsii_.Get(
		p,
		"successThreshold",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (p *ProbeOptions) GetTimeoutSeconds() cdk8s.DurationIface {
	var returns cdk8s.DurationIface
	_jsii_.Get(
		p,
		"timeoutSeconds",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.DurationIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.Duration)(nil)).Elem(),
		},
	)
	return returns
}


// Experimental.
type Protocol string

const (
	ProtocolTcp Protocol = "TCP"
	ProtocolUdp Protocol = "UDP"
	ProtocolSctp Protocol = "SCTP"
)

// Class interface
type ResourceIface interface {
	constructs.IConstructIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
}

// Base class for all Kubernetes objects in stdk8s.
//
// Represents a single
// resource.
// Experimental.
// Struct proxy
type Resource struct {
	// The underlying cdk8s API object.
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
}

func (r *Resource) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		r,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (r *Resource) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		r,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (r *Resource) GetName() string {
	var returns string
	_jsii_.Get(
		r,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Creates a new construct node.
func NewResource(scope constructs.ConstructIface, id string, options constructs.ConstructOptionsIface) ResourceIface {
	_init_.Initialize()
	self := Resource{}
	_jsii_.Create(
		"cdk8s-plus-17.Resource",
		[]interface{}{scope, id, options},
		[]_jsii_.FQN{"cdk8s-plus-17.IResource"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (r *Resource) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *Resource) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		r,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (r *Resource) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		r,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (r *Resource) ToString() string {
	var returns string
	_jsii_.Invoke(
		r,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// ResourcePropsIface is the public interface for the custom type ResourceProps
// Experimental.
type ResourcePropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
}

// Initialization properties for resources.
// Experimental.
// Struct proxy
type ResourceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
}

func (r *ResourceProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		r,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}


// Restart policy for all containers within the pod.
// Experimental.
type RestartPolicy string

const (
	RestartPolicyAlways RestartPolicy = "ALWAYS"
	RestartPolicyOnFailure RestartPolicy = "ON_FAILURE"
	RestartPolicyNever RestartPolicy = "NEVER"
)

// Class interface
type SecretIface interface {
	constructs.IConstructIface
	IResourceIface
	ISecretIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddStringData(key string, value string)
	GetStringData(key string) string
}

// Kubernetes Secrets let you store and manage sensitive information, such as passwords, OAuth tokens, and ssh keys.
//
// Storing confidential information in a
// Secret is safer and more flexible than putting it verbatim in a Pod
// definition or in a container image.
// See: https://kubernetes.io/docs/concepts/configuration/secret
//
// Experimental.
// Struct proxy
type Secret struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
}

func (s *Secret) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Secret) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Secret) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewSecret(scope constructs.ConstructIface, id string, props SecretPropsIface) SecretIface {
	_init_.Initialize()
	self := Secret{}
	_jsii_.Create(
		"cdk8s-plus-17.Secret",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.ISecret"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Secret_FromSecretName(name string) ISecretIface {
	_init_.Initialize()
	var returns ISecretIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.Secret",
		"fromSecretName",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ISecretIface)(nil)).Elem(): reflect.TypeOf((*ISecret)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Secret) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Secret) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Secret) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Secret) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Secret) AddStringData(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addStringData",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Secret) GetStringData(key string) string {
	var returns string
	_jsii_.Invoke(
		s,
		"getStringData",
		[]interface{}{key},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// SecretPropsIface is the public interface for the custom type SecretProps
// Experimental.
type SecretPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetStringData() map[string]string
	GetType() string
}

// Experimental.
// Struct proxy
type SecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
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

func (s *SecretProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (s *SecretProps) GetStringData() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"stringData",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *SecretProps) GetType() string {
	var returns string
	_jsii_.Get(
		s,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// SecretValueIface is the public interface for the custom type SecretValue
// Experimental.
type SecretValueIface interface {
	GetKey() string
	GetSecret() ISecretIface
}

// Represents a specific value in JSON secret.
// Experimental.
// Struct proxy
type SecretValue struct {
	// The JSON key.
	// Experimental.
	Key string `json:"key"`
	// The secret.
	// Experimental.
	Secret ISecretIface `json:"secret"`
}

func (s *SecretValue) GetKey() string {
	var returns string
	_jsii_.Get(
		s,
		"key",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *SecretValue) GetSecret() ISecretIface {
	var returns ISecretIface
	_jsii_.Get(
		s,
		"secret",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ISecretIface)(nil)).Elem(): reflect.TypeOf((*ISecret)(nil)).Elem(),
		},
	)
	return returns
}


// Class interface
type ServiceIface interface {
	constructs.IConstructIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetPorts() []ServicePortIface
	GetSelector() map[string]string
	GetType() ServiceType
	GetClusterIp() string
	GetExternalName() string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddDeployment(deployment DeploymentIface, port float64, options ServicePortOptionsIface)
	AddSelector(label string, value string)
	Serve(port float64, options ServicePortOptionsIface)
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
// Struct proxy
type Service struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// Ports for this service.
	//
	// Use `serve()` to expose additional service ports.
	// Experimental.
	Ports []ServicePortIface `json:"ports"`
	// Returns the labels which are used to select pods for this service.
	// Experimental.
	Selector map[string]string `json:"selector"`
	// Determines how the Service is exposed.
	// Experimental.
	Type ServiceType `json:"type"`
	// The IP address of the service and is usually assigned randomly by the master.
	// Experimental.
	ClusterIp string `json:"clusterIP"`
	// The externalName to be used for EXTERNAL_NAME types.
	// Experimental.
	ExternalName string `json:"externalName"`
}

func (s *Service) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Service) GetPorts() []ServicePortIface {
	var returns []ServicePortIface
	_jsii_.Get(
		s,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServicePortIface)(nil)).Elem(): reflect.TypeOf((*ServicePort)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) GetSelector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"selector",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) GetType() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		s,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServiceType)(nil)).Elem(): reflect.TypeOf((*ServiceType)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) GetClusterIp() string {
	var returns string
	_jsii_.Get(
		s,
		"clusterIP",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Service) GetExternalName() string {
	var returns string
	_jsii_.Get(
		s,
		"externalName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewService(scope constructs.ConstructIface, id string, props ServicePropsIface) ServiceIface {
	_init_.Initialize()
	self := Service{}
	_jsii_.Create(
		"cdk8s-plus-17.Service",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (s *Service) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Service) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Service) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *Service) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *Service) AddDeployment(deployment DeploymentIface, port float64, options ServicePortOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addDeployment",
		[]interface{}{deployment, port, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Service) AddSelector(label string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addSelector",
		[]interface{}{label, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *Service) Serve(port float64, options ServicePortOptionsIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"serve",
		[]interface{}{port, options},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// Class interface
type ServiceAccountIface interface {
	constructs.IConstructIface
	IResourceIface
	IServiceAccountIface
	IResourceIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetSecrets() []ISecretIface
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddSecret(secret ISecretIface)
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
// Struct proxy
type ServiceAccount struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// List of secrets allowed to be used by pods running using this service account.
	//
	// Returns a copy. To add a secret, use `addSecret()`.
	// Experimental.
	Secrets []ISecretIface `json:"secrets"`
}

func (s *ServiceAccount) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceAccount) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceAccount) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServiceAccount) GetSecrets() []ISecretIface {
	var returns []ISecretIface
	_jsii_.Get(
		s,
		"secrets",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ISecretIface)(nil)).Elem(): reflect.TypeOf((*ISecret)(nil)).Elem(),
		},
	)
	return returns
}


func NewServiceAccount(scope constructs.ConstructIface, id string, props ServiceAccountPropsIface) ServiceAccountIface {
	_init_.Initialize()
	self := ServiceAccount{}
	_jsii_.Create(
		"cdk8s-plus-17.ServiceAccount",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IServiceAccount"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func ServiceAccount_FromServiceAccountName(name string) IServiceAccountIface {
	_init_.Initialize()
	var returns IServiceAccountIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.ServiceAccount",
		"fromServiceAccountName",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceAccount) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *ServiceAccount) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *ServiceAccount) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceAccount) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServiceAccount) AddSecret(secret ISecretIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addSecret",
		[]interface{}{secret},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// ServiceAccountPropsIface is the public interface for the custom type ServiceAccountProps
// Experimental.
type ServiceAccountPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetSecrets() []ISecretIface
}

// Properties for initialization of `ServiceAccount`.
//
// Properties for initialization of `ServiceAccount`.
// Experimental.
// Struct proxy
type ServiceAccountProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// List of secrets allowed to be used by pods running using this ServiceAccount.
	// See: https://kubernetes.io/docs/concepts/configuration/secret
	//
	// Experimental.
	Secrets []ISecretIface `json:"secrets"`
}

func (s *ServiceAccountProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceAccountProps) GetSecrets() []ISecretIface {
	var returns []ISecretIface
	_jsii_.Get(
		s,
		"secrets",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ISecretIface)(nil)).Elem(): reflect.TypeOf((*ISecret)(nil)).Elem(),
		},
	)
	return returns
}


// ServiceIngressV1BetaBackendOptionsIface is the public interface for the custom type ServiceIngressV1BetaBackendOptions
// Experimental.
type ServiceIngressV1BetaBackendOptionsIface interface {
	GetPort() float64
}

// Options for setting up backends for ingress rules.
// Experimental.
// Struct proxy
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

func (s *ServiceIngressV1BetaBackendOptions) GetPort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ServicePortIface is the public interface for the custom type ServicePort
// Experimental.
type ServicePortIface interface {
	GetName() string
	GetNodePort() float64
	GetProtocol() Protocol
	GetTargetPort() float64
	GetPort() float64
}

// Definition of a service port.
// Experimental.
// Struct proxy
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

func (s *ServicePort) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServicePort) GetNodePort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"nodePort",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServicePort) GetProtocol() Protocol {
	var returns Protocol
	_jsii_.Get(
		s,
		"protocol",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*Protocol)(nil)).Elem(): reflect.TypeOf((*Protocol)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServicePort) GetTargetPort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"targetPort",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServicePort) GetPort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"port",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ServicePortOptionsIface is the public interface for the custom type ServicePortOptions
// Experimental.
type ServicePortOptionsIface interface {
	GetName() string
	GetNodePort() float64
	GetProtocol() Protocol
	GetTargetPort() float64
}

// Experimental.
// Struct proxy
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

func (s *ServicePortOptions) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServicePortOptions) GetNodePort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"nodePort",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServicePortOptions) GetProtocol() Protocol {
	var returns Protocol
	_jsii_.Get(
		s,
		"protocol",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*Protocol)(nil)).Elem(): reflect.TypeOf((*Protocol)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServicePortOptions) GetTargetPort() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"targetPort",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// ServicePropsIface is the public interface for the custom type ServiceProps
// Experimental.
type ServicePropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetClusterIp() string
	GetExternalIPs() []string
	GetExternalName() string
	GetLoadBalancerSourceRanges() []string
	GetPorts() []ServicePortIface
	GetType() ServiceType
}

// Properties for initialization of `Service`.
// Experimental.
// Struct proxy
type ServiceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
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
	Ports []ServicePortIface `json:"ports"`
	// Determines how the Service is exposed.
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// Experimental.
	Type ServiceType `json:"type"`
}

func (s *ServiceProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceProps) GetClusterIp() string {
	var returns string
	_jsii_.Get(
		s,
		"clusterIP",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServiceProps) GetExternalIPs() []string {
	var returns []string
	_jsii_.Get(
		s,
		"externalIPs",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceProps) GetExternalName() string {
	var returns string
	_jsii_.Get(
		s,
		"externalName",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *ServiceProps) GetLoadBalancerSourceRanges() []string {
	var returns []string
	_jsii_.Get(
		s,
		"loadBalancerSourceRanges",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceProps) GetPorts() []ServicePortIface {
	var returns []ServicePortIface
	_jsii_.Get(
		s,
		"ports",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServicePortIface)(nil)).Elem(): reflect.TypeOf((*ServicePort)(nil)).Elem(),
		},
	)
	return returns
}

func (s *ServiceProps) GetType() ServiceType {
	var returns ServiceType
	_jsii_.Get(
		s,
		"type",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServiceType)(nil)).Elem(): reflect.TypeOf((*ServiceType)(nil)).Elem(),
		},
	)
	return returns
}


// For some parts of your application (for example, frontends) you may want to expose a Service onto an external IP address, that's outside of your cluster.
//
// Kubernetes ServiceTypes allow you to specify what kind of Service you want.
// The default is ClusterIP.
// Experimental.
type ServiceType string

const (
	ServiceTypeClusterIp ServiceType = "CLUSTER_IP"
	ServiceTypeNodePort ServiceType = "NODE_PORT"
	ServiceTypeLoadBalancer ServiceType = "LOAD_BALANCER"
	ServiceTypeExternalName ServiceType = "EXTERNAL_NAME"
)

// Class interface
type StatefulSetIface interface {
	constructs.IConstructIface
	IResourceIface
	IPodTemplateIface
	IPodSpecIface
	GetApiObject() cdk8s.ApiObjectIface
	GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetName() string
	GetContainers() []ContainerIface
	GetLabelSelector() map[string]string
	GetPodManagementPolicy() PodManagementPolicy
	GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface
	GetReplicas() float64
	GetVolumes() []VolumeIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSessionIface)
	OnValidate() []string
	ToString() string
	AddContainer(container ContainerPropsIface) ContainerIface
	AddVolume(volume VolumeIface)
	SelectByLabel(key string, value string)
}

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
// Struct proxy
type StatefulSet struct {
	// The underlying cdk8s API object.
	// See: base.Resource.apiObject
	//
	// Experimental.
	ApiObject cdk8s.ApiObjectIface `json:"apiObject"`
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataDefinitionIface `json:"metadata"`
	// The name of this API object.
	// Experimental.
	Name string `json:"name"`
	// The containers belonging to the pod.
	//
	// Use `addContainer` to add containers.
	// Experimental.
	Containers []ContainerIface `json:"containers"`
	// The labels this statefulset will match against in order to select pods.
	//
	// Returns a a copy. Use `selectByLabel()` to add labels.
	// Experimental.
	LabelSelector map[string]string `json:"labelSelector"`
	// Management policy to use for the set.
	// Experimental.
	PodManagementPolicy PodManagementPolicy `json:"podManagementPolicy"`
	// Provides read/write access to the underlying pod metadata of the resource.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataDefinitionIface `json:"podMetadata"`
	// Number of desired pods.
	// Experimental.
	Replicas float64 `json:"replicas"`
	// The volumes associated with this pod.
	//
	// Use `addVolume` to add volumes.
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// Restart policy for all containers within the pod.
	// Experimental.
	RestartPolicy RestartPolicy `json:"restartPolicy"`
	// The service account used to run this pod.
	// Experimental.
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
}

func (s *StatefulSet) GetApiObject() cdk8s.ApiObjectIface {
	var returns cdk8s.ApiObjectIface
	_jsii_.Get(
		s,
		"apiObject",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObject)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetName() string {
	var returns string
	_jsii_.Get(
		s,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *StatefulSet) GetContainers() []ContainerIface {
	var returns []ContainerIface
	_jsii_.Get(
		s,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetLabelSelector() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		s,
		"labelSelector",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetPodManagementPolicy() PodManagementPolicy {
	var returns PodManagementPolicy
	_jsii_.Get(
		s,
		"podManagementPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(): reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetPodMetadata() cdk8s.ApiObjectMetadataDefinitionIface {
	var returns cdk8s.ApiObjectMetadataDefinitionIface
	_jsii_.Get(
		s,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinitionIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadataDefinition)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetReplicas() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"replicas",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *StatefulSet) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		s,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		s,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		s,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}


func NewStatefulSet(scope constructs.ConstructIface, id string, props StatefulSetPropsIface) StatefulSetIface {
	_init_.Initialize()
	self := StatefulSet{}
	_jsii_.Create(
		"cdk8s-plus-17.StatefulSet",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{"cdk8s-plus-17.IPodTemplate"},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (s *StatefulSet) OnPrepare() {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onPrepare",
		[]interface{}{},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *StatefulSet) OnSynthesize(session constructs.ISynthesisSessionIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"onSynthesize",
		[]interface{}{session},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *StatefulSet) OnValidate() []string {
	var returns []string
	_jsii_.Invoke(
		s,
		"onValidate",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*string)(nil)).Elem(): reflect.TypeOf((*string)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) ToString() string {
	var returns string
	_jsii_.Invoke(
		s,
		"toString",
		[]interface{}{},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *StatefulSet) AddContainer(container ContainerPropsIface) ContainerIface {
	var returns ContainerIface
	_jsii_.Invoke(
		s,
		"addContainer",
		[]interface{}{container},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerIface)(nil)).Elem(): reflect.TypeOf((*Container)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSet) AddVolume(volume VolumeIface) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"addVolume",
		[]interface{}{volume},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

func (s *StatefulSet) SelectByLabel(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		s,
		"selectByLabel",
		[]interface{}{key, value},
		false,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
}

// StatefulSetPropsIface is the public interface for the custom type StatefulSetProps
// Experimental.
type StatefulSetPropsIface interface {
	GetMetadata() cdk8s.ApiObjectMetadataIface
	GetContainers() []ContainerPropsIface
	GetRestartPolicy() RestartPolicy
	GetServiceAccount() IServiceAccountIface
	GetVolumes() []VolumeIface
	GetPodMetadata() cdk8s.ApiObjectMetadataIface
	GetService() ServiceIface
	GetDefaultSelector() bool
	GetPodManagementPolicy() PodManagementPolicy
	GetReplicas() float64
}

// Properties for initialization of `StatefulSet`.
// Experimental.
// Struct proxy
type StatefulSetProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	// Experimental.
	Metadata cdk8s.ApiObjectMetadataIface `json:"metadata"`
	// List of containers belonging to the pod.
	//
	// Containers cannot currently be
	// added or removed. There must be at least one container in a Pod.
	//
	// You can add additionnal containers using `podSpec.addContainer()`
	// Experimental.
	Containers []ContainerPropsIface `json:"containers"`
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
	ServiceAccount IServiceAccountIface `json:"serviceAccount"`
	// List of volumes that can be mounted by containers belonging to the pod.
	//
	// You can also add volumes later using `podSpec.addVolume()`
	// See: https://kubernetes.io/docs/concepts/storage/volumes
	//
	// Experimental.
	Volumes []VolumeIface `json:"volumes"`
	// The pod metadata.
	// Experimental.
	PodMetadata cdk8s.ApiObjectMetadataIface `json:"podMetadata"`
	// Service to associate with the statefulset.
	// Experimental.
	Service ServiceIface `json:"service"`
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

func (s *StatefulSetProps) GetMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		s,
		"metadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetContainers() []ContainerPropsIface {
	var returns []ContainerPropsIface
	_jsii_.Get(
		s,
		"containers",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ContainerPropsIface)(nil)).Elem(): reflect.TypeOf((*ContainerProps)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetRestartPolicy() RestartPolicy {
	var returns RestartPolicy
	_jsii_.Get(
		s,
		"restartPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*RestartPolicy)(nil)).Elem(): reflect.TypeOf((*RestartPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetServiceAccount() IServiceAccountIface {
	var returns IServiceAccountIface
	_jsii_.Get(
		s,
		"serviceAccount",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*IServiceAccountIface)(nil)).Elem(): reflect.TypeOf((*IServiceAccount)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetVolumes() []VolumeIface {
	var returns []VolumeIface
	_jsii_.Get(
		s,
		"volumes",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetPodMetadata() cdk8s.ApiObjectMetadataIface {
	var returns cdk8s.ApiObjectMetadataIface
	_jsii_.Get(
		s,
		"podMetadata",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*cdk8s.ApiObjectMetadataIface)(nil)).Elem(): reflect.TypeOf((*cdk8s.ApiObjectMetadata)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetService() ServiceIface {
	var returns ServiceIface
	_jsii_.Get(
		s,
		"service",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*ServiceIface)(nil)).Elem(): reflect.TypeOf((*Service)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetDefaultSelector() bool {
	var returns bool
	_jsii_.Get(
		s,
		"defaultSelector",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (s *StatefulSetProps) GetPodManagementPolicy() PodManagementPolicy {
	var returns PodManagementPolicy
	_jsii_.Get(
		s,
		"podManagementPolicy",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(): reflect.TypeOf((*PodManagementPolicy)(nil)).Elem(),
		},
	)
	return returns
}

func (s *StatefulSetProps) GetReplicas() float64 {
	var returns float64
	_jsii_.Get(
		s,
		"replicas",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


// Class interface
type VolumeIface interface {
	GetName() string
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
// Struct proxy
type Volume struct {
	// Experimental.
	Name string `json:"name"`
}

func (v *Volume) GetName() string {
	var returns string
	_jsii_.Get(
		v,
		"name",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


func NewVolume(name string, config interface{}) VolumeIface {
	_init_.Initialize()
	self := Volume{}
	_jsii_.Create(
		"cdk8s-plus-17.Volume",
		[]interface{}{name, config},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func Volume_FromConfigMap(configMap IConfigMapIface, options ConfigMapVolumeOptionsIface) VolumeIface {
	_init_.Initialize()
	var returns VolumeIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.Volume",
		"fromConfigMap",
		[]interface{}{configMap, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

func Volume_FromEmptyDir(name string, options EmptyDirVolumeOptionsIface) VolumeIface {
	_init_.Initialize()
	var returns VolumeIface
	_jsii_.InvokeStatic(
		"cdk8s-plus-17.Volume",
		"fromEmptyDir",
		[]interface{}{name, options},
		true,
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}

// VolumeMountIface is the public interface for the custom type VolumeMount
// Experimental.
type VolumeMountIface interface {
	GetPropagation() MountPropagation
	GetReadOnly() bool
	GetSubPath() string
	GetSubPathExpr() string
	GetPath() string
	GetVolume() VolumeIface
}

// Mount a volume from the pod to the container.
// Experimental.
// Struct proxy
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
	Volume VolumeIface `json:"volume"`
}

func (v *VolumeMount) GetPropagation() MountPropagation {
	var returns MountPropagation
	_jsii_.Get(
		v,
		"propagation",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*MountPropagation)(nil)).Elem(): reflect.TypeOf((*MountPropagation)(nil)).Elem(),
		},
	)
	return returns
}

func (v *VolumeMount) GetReadOnly() bool {
	var returns bool
	_jsii_.Get(
		v,
		"readOnly",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VolumeMount) GetSubPath() string {
	var returns string
	_jsii_.Get(
		v,
		"subPath",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VolumeMount) GetSubPathExpr() string {
	var returns string
	_jsii_.Get(
		v,
		"subPathExpr",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VolumeMount) GetPath() string {
	var returns string
	_jsii_.Get(
		v,
		"path",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (v *VolumeMount) GetVolume() VolumeIface {
	var returns VolumeIface
	_jsii_.Get(
		v,
		"volume",
		&returns,
		map[reflect.Type]reflect.Type{
			reflect.TypeOf((*VolumeIface)(nil)).Elem(): reflect.TypeOf((*Volume)(nil)).Elem(),
		},
	)
	return returns
}


