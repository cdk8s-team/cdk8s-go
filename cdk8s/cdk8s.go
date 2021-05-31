// Cloud Development Kit for Kubernetes
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/cdk8s-team/cdk8s-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-go/cdk8s/internal"
)

type ApiObject interface {
	constructs.Construct
	ApiGroup() *string
	ApiVersion() *string
	Chart() Chart
	Kind() *string
	Metadata() ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ApiObject
type jsiiProxy_ApiObject struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiObject) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObject) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObject) Chart() Chart {
	var returns Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObject) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObject) Metadata() ApiObjectMetadataDefinition {
	var returns ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines an API object.
func NewApiObject(scope constructs.Construct, id *string, props *ApiObjectProps) ApiObject {
	_init_.Initialize()

	j := jsiiProxy_ApiObject{}

	_jsii_.Create(
		"cdk8s.ApiObject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an API object.
func NewApiObject_Override(a ApiObject, scope constructs.Construct, id *string, props *ApiObjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.ApiObject",
		[]interface{}{scope, id, props},
		a,
	)
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ApiObject_Of(c constructs.IConstruct) ApiObject {
	_init_.Initialize()

	var returns ApiObject

	_jsii_.StaticInvoke(
		"cdk8s.ApiObject",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (a *jsiiProxy_ApiObject) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
func (a *jsiiProxy_ApiObject) AddJsonPatch(ops ...JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
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
func (a *jsiiProxy_ApiObject) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (a *jsiiProxy_ApiObject) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
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
func (a *jsiiProxy_ApiObject) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
func (a *jsiiProxy_ApiObject) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiObject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Metadata associated with this object.
type ApiObjectMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be
	// preserved when modifying objects.
	// See: http://kubernetes.io/docs/user-guide/annotations
	//
	Annotations *map[string]*string `json:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers and services.
	// See: http://kubernetes.io/docs/user-guide/labels
	//
	Labels *map[string]*string `json:"labels"`
	// The unique, namespace-global, name of this object inside the Kubernetes cluster.
	//
	// Normally, you shouldn't specify names for objects and let the CDK generate
	// a name for you that is application-unique. The names CDK generates are
	// composed from the construct path components, separated by dots and a suffix
	// that is based on a hash of the entire path, to ensure uniqueness.
	//
	// You can supply custom name allocation logic by overriding the
	// `chart.generateObjectName` method.
	//
	// If you use an explicit name here, bear in mind that this reduces the
	// composability of your construct because it won't be possible to include
	// more than one instance in any app. Therefore it is highly recommended to
	// leave this unspecified.
	Name *string `json:"name"`
	// Namespace defines the space within each name must be unique.
	//
	// An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty. Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
	Namespace *string `json:"namespace"`
}

// Object metadata.
type ApiObjectMetadataDefinition interface {
	Name() *string
	Namespace() *string
	Add(key *string, value interface{})
	AddAnnotation(key *string, value *string)
	AddLabel(key *string, value *string)
	GetLabel(key *string) *string
	ToJson() interface{}
}

// The jsii proxy struct for ApiObjectMetadataDefinition
type jsiiProxy_ApiObjectMetadataDefinition struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiObjectMetadataDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiObjectMetadataDefinition) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}


func NewApiObjectMetadataDefinition(options *ApiObjectMetadata) ApiObjectMetadataDefinition {
	_init_.Initialize()

	j := jsiiProxy_ApiObjectMetadataDefinition{}

	_jsii_.Create(
		"cdk8s.ApiObjectMetadataDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewApiObjectMetadataDefinition_Override(a ApiObjectMetadataDefinition, options *ApiObjectMetadata) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.ApiObjectMetadataDefinition",
		[]interface{}{options},
		a,
	)
}

// Adds an arbitrary key/value to the object metadata.
func (a *jsiiProxy_ApiObjectMetadataDefinition) Add(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{key, value},
	)
}

// Add an annotation.
func (a *jsiiProxy_ApiObjectMetadataDefinition) AddAnnotation(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"addAnnotation",
		[]interface{}{key, value},
	)
}

// Add a label.
func (a *jsiiProxy_ApiObjectMetadataDefinition) AddLabel(key *string, value *string) {
	_jsii_.InvokeVoid(
		a,
		"addLabel",
		[]interface{}{key, value},
	)
}

// Returns: a value of a label or undefined
func (a *jsiiProxy_ApiObjectMetadataDefinition) GetLabel(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getLabel",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Synthesizes a k8s ObjectMeta for this metadata set.
func (a *jsiiProxy_ApiObjectMetadataDefinition) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for defining API objects.
type ApiObjectProps struct {
	// API version.
	ApiVersion *string `json:"apiVersion"`
	// Resource kind.
	Kind *string `json:"kind"`
	// Object metadata.
	//
	// If `name` is not specified, an app-unique name will be allocated by the
	// framework based on the path of the construct within thes construct tree.
	Metadata *ApiObjectMetadata `json:"metadata"`
}

// Represents a cdk8s application.
type App interface {
	constructs.Construct
	Outdir() *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Synth()
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_App) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}


// Defines an app.
func NewApp(props *AppProps) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"cdk8s.App",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Defines an app.
func NewApp_Override(a App, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.App",
		[]interface{}{props},
		a,
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
func (a *jsiiProxy_App) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (a *jsiiProxy_App) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
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
func (a *jsiiProxy_App) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Synthesizes all manifests to the output directory.
func (a *jsiiProxy_App) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppProps struct {
	// The directory to output Kubernetes manifests.
	Outdir *string `json:"outdir"`
}

type Chart interface {
	constructs.Construct
	Labels() *map[string]*string
	Namespace() *string
	AddDependency(dependencies ...constructs.IConstruct)
	GenerateObjectName(apiObject ApiObject) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() *[]interface{}
	ToString() *string
}

// The jsii proxy struct for Chart
type jsiiProxy_Chart struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Chart) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Chart) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}


func NewChart(scope constructs.Construct, id *string, props *ChartProps) Chart {
	_init_.Initialize()

	j := jsiiProxy_Chart{}

	_jsii_.Create(
		"cdk8s.Chart",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewChart_Override(c Chart, scope constructs.Construct, id *string, props *ChartProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Chart",
		[]interface{}{scope, id, props},
		c,
	)
}

// Finds the chart in which a node is defined.
func Chart_Of(c constructs.IConstruct) Chart {
	_init_.Initialize()

	var returns Chart

	_jsii_.StaticInvoke(
		"cdk8s.Chart",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

// Create a dependency between this Chart and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
func (c *jsiiProxy_Chart) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addDependency",
		args,
	)
}

// Generates a app-unique name for an object given it's construct node path.
//
// Different resource types may have different constraints on names
// (`metadata.name`). The previous version of the name generator was
// compatible with DNS_SUBDOMAIN but not with DNS_LABEL.
//
// For example, `Deployment` names must comply with DNS_SUBDOMAIN while
// `Service` names must comply with DNS_LABEL.
//
// Since there is no formal specification for this, the default name
// generation scheme for kubernetes objects in cdk8s was changed to DNS_LABEL,
// since it’s the common denominator for all kubernetes resources
// (supposedly).
//
// You can override this method if you wish to customize object names at the
// chart level.
func (c *jsiiProxy_Chart) GenerateObjectName(apiObject ApiObject) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generateObjectName",
		[]interface{}{apiObject},
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
func (c *jsiiProxy_Chart) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (c *jsiiProxy_Chart) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
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
func (c *jsiiProxy_Chart) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders this chart to a set of Kubernetes JSON resources.
//
// Returns: array of resource manifests
func (c *jsiiProxy_Chart) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_Chart) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ChartProps struct {
	// Labels to apply to all resources in this chart.
	Labels *map[string]*string `json:"labels"`
	// The default namespace for all objects defined in this chart (directly or indirectly).
	//
	// This namespace will only apply to objects that don't have a
	// `namespace` explicitly defined for them.
	Namespace *string `json:"namespace"`
}

// Represents the dependency graph for a given Node.
//
// This graph includes the dependency relationships between all nodes in the
// node (construct) sub-tree who's root is this Node.
//
// Note that this means that lonely nodes (no dependencies and no dependants) are also included in this graph as
// childless children of the root node of the graph.
//
// The graph does not include cross-scope dependencies. That is, if a child on the current scope depends on a node
// from a different scope, that relationship is not represented in this graph.
type DependencyGraph interface {
	Root() DependencyVertex
	Topology() *[]constructs.IConstruct
}

// The jsii proxy struct for DependencyGraph
type jsiiProxy_DependencyGraph struct {
	_ byte // padding
}

func (j *jsiiProxy_DependencyGraph) Root() DependencyVertex {
	var returns DependencyVertex
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}


func NewDependencyGraph(node constructs.Node) DependencyGraph {
	_init_.Initialize()

	j := jsiiProxy_DependencyGraph{}

	_jsii_.Create(
		"cdk8s.DependencyGraph",
		[]interface{}{node},
		&j,
	)

	return &j
}

func NewDependencyGraph_Override(d DependencyGraph, node constructs.Node) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.DependencyGraph",
		[]interface{}{node},
		d,
	)
}

// See: Vertex.topology()
//
func (d *jsiiProxy_DependencyGraph) Topology() *[]constructs.IConstruct {
	var returns *[]constructs.IConstruct

	_jsii_.Invoke(
		d,
		"topology",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a vertex in the graph.
//
// The value of each vertex is an `IConstruct` that is accessible via the `.value` getter.
type DependencyVertex interface {
	Inbound() *[]DependencyVertex
	Outbound() *[]DependencyVertex
	Value() constructs.IConstruct
	AddChild(dep DependencyVertex)
	Topology() *[]constructs.IConstruct
}

// The jsii proxy struct for DependencyVertex
type jsiiProxy_DependencyVertex struct {
	_ byte // padding
}

func (j *jsiiProxy_DependencyVertex) Inbound() *[]DependencyVertex {
	var returns *[]DependencyVertex
	_jsii_.Get(
		j,
		"inbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DependencyVertex) Outbound() *[]DependencyVertex {
	var returns *[]DependencyVertex
	_jsii_.Get(
		j,
		"outbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DependencyVertex) Value() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewDependencyVertex(value constructs.IConstruct) DependencyVertex {
	_init_.Initialize()

	j := jsiiProxy_DependencyVertex{}

	_jsii_.Create(
		"cdk8s.DependencyVertex",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewDependencyVertex_Override(d DependencyVertex, value constructs.IConstruct) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.DependencyVertex",
		[]interface{}{value},
		d,
	)
}

// Adds a vertex as a dependency of the current node.
//
// Also updates the parents of `dep`, so that it contains this node as a parent.
//
// This operation will fail in case it creates a cycle in the graph.
func (d *jsiiProxy_DependencyVertex) AddChild(dep DependencyVertex) {
	_jsii_.InvokeVoid(
		d,
		"addChild",
		[]interface{}{dep},
	)
}

// Returns a topologically sorted array of the constructs in the sub-graph.
func (d *jsiiProxy_DependencyVertex) Topology() *[]constructs.IConstruct {
	var returns *[]constructs.IConstruct

	_jsii_.Invoke(
		d,
		"topology",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a length of time.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative.
type Duration interface {
	ToDays(opts *TimeConversionOptions) *float64
	ToHours(opts *TimeConversionOptions) *float64
	ToHumanString() *string
	ToIsoString() *string
	ToMilliseconds(opts *TimeConversionOptions) *float64
	ToMinutes(opts *TimeConversionOptions) *float64
	ToSeconds(opts *TimeConversionOptions) *float64
}

// The jsii proxy struct for Duration
type jsiiProxy_Duration struct {
	_ byte // padding
}

// Create a Duration representing an amount of days.
//
// Returns: a new `Duration` representing `amount` Days.
func Duration_Days(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"days",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of hours.
//
// Returns: a new `Duration` representing `amount` Hours.
func Duration_Hours(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"hours",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of milliseconds.
//
// Returns: a new `Duration` representing `amount` ms.
func Duration_Millis(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of minutes.
//
// Returns: a new `Duration` representing `amount` Minutes.
func Duration_Minutes(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"minutes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Parse a period formatted according to the ISO 8601 standard.
//
// Returns: the parsed `Duration`.
// See: https://www.iso.org/fr/standard/70907.html
//
func Duration_Parse(duration *string) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"parse",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Create a Duration representing an amount of seconds.
//
// Returns: a new `Duration` representing `amount` Seconds.
func Duration_Seconds(amount *float64) Duration {
	_init_.Initialize()

	var returns Duration

	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"seconds",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Return the total number of days in this Duration.
//
// Returns: the value of this `Duration` expressed in Days.
func (d *jsiiProxy_Duration) ToDays(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toDays",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return the total number of hours in this Duration.
//
// Returns: the value of this `Duration` expressed in Hours.
func (d *jsiiProxy_Duration) ToHours(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toHours",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Turn this duration into a human-readable string.
func (d *jsiiProxy_Duration) ToHumanString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toHumanString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Return an ISO 8601 representation of this period.
//
// Returns: a string starting with 'PT' describing the period
// See: https://www.iso.org/fr/standard/70907.html
//
func (d *jsiiProxy_Duration) ToIsoString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toIsoString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Return the total number of milliseconds in this Duration.
//
// Returns: the value of this `Duration` expressed in Milliseconds.
func (d *jsiiProxy_Duration) ToMilliseconds(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMilliseconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return the total number of minutes in this Duration.
//
// Returns: the value of this `Duration` expressed in Minutes.
func (d *jsiiProxy_Duration) ToMinutes(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toMinutes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return the total number of seconds in this Duration.
//
// Returns: the value of this `Duration` expressed in Seconds.
func (d *jsiiProxy_Duration) ToSeconds(opts *TimeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"toSeconds",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

type GroupVersionKind struct {
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion *string `json:"apiVersion"`
	// The object kind.
	Kind *string `json:"kind"`
}

// Represents a Helm deployment.
//
// Use this construct to import an existing Helm chart and incorporate it into your constructs.
type Helm interface {
	Include
	ApiObjects() *[]ApiObject
	ReleaseName() *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToString() *string
}

// The jsii proxy struct for Helm
type jsiiProxy_Helm struct {
	jsiiProxy_Include
}

func (j *jsiiProxy_Helm) ApiObjects() *[]ApiObject {
	var returns *[]ApiObject
	_jsii_.Get(
		j,
		"apiObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Helm) ReleaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseName",
		&returns,
	)
	return returns
}


func NewHelm(scope constructs.Construct, id *string, props *HelmProps) Helm {
	_init_.Initialize()

	j := jsiiProxy_Helm{}

	_jsii_.Create(
		"cdk8s.Helm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHelm_Override(h Helm, scope constructs.Construct, id *string, props *HelmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Helm",
		[]interface{}{scope, id, props},
		h,
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
func (h *jsiiProxy_Helm) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (h *jsiiProxy_Helm) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
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
func (h *jsiiProxy_Helm) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (h *jsiiProxy_Helm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `Helm`.
type HelmProps struct {
	// The chart name to use. It can be a chart from a helm repository or a local directory.
	//
	// This name is passed to `helm template` and has all the relevant semantics.
	//
	// TODO: EXAMPLE
	//
	Chart *string `json:"chart"`
	// The local helm executable to use in order to create the manifest the chart.
	HelmExecutable *string `json:"helmExecutable"`
	// Additional flags to add to the `helm` execution.
	HelmFlags *[]*string `json:"helmFlags"`
	// The release name.
	// See: https://helm.sh/docs/intro/using_helm/#three-big-concepts
	//
	ReleaseName *string `json:"releaseName"`
	// Values to pass to the chart.
	Values *map[string]interface{} `json:"values"`
}

type IAnyProducer interface {
	Produce() interface{}
}

// The jsii proxy for IAnyProducer
type jsiiProxy_IAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnyProducer) Produce() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Reads a YAML manifest from a file or a URL and defines all resources as API objects within the defined scope.
//
// The names (`metadata.name`) of imported resources will be preserved as-is
// from the manifest.
type Include interface {
	constructs.Construct
	ApiObjects() *[]ApiObject
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToString() *string
}

// The jsii proxy struct for Include
type jsiiProxy_Include struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Include) ApiObjects() *[]ApiObject {
	var returns *[]ApiObject
	_jsii_.Get(
		j,
		"apiObjects",
		&returns,
	)
	return returns
}


func NewInclude(scope constructs.Construct, id *string, props *IncludeProps) Include {
	_init_.Initialize()

	j := jsiiProxy_Include{}

	_jsii_.Create(
		"cdk8s.Include",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewInclude_Override(i Include, scope constructs.Construct, id *string, props *IncludeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s.Include",
		[]interface{}{scope, id, props},
		i,
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
func (i *jsiiProxy_Include) OnPrepare() {
	_jsii_.InvokeVoid(
		i,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
func (i *jsiiProxy_Include) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"onSynthesize",
		[]interface{}{session},
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
func (i *jsiiProxy_Include) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_Include) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IncludeProps struct {
	// Local file path or URL which includes a Kubernetes YAML manifest.
	//
	// TODO: EXAMPLE
	//
	Url *string `json:"url"`
}

// Utility for applying RFC-6902 JSON-Patch to a document.
//
// Use the the `JsonPatch.apply(doc, ...ops)` function to apply a set of
// operations to a JSON document and return the result.
//
// Operations can be created using the factory methods `JsonPatch.add()`,
// `JsonPatch.remove()`, etc.
//
// TODO: EXAMPLE
//
type JsonPatch interface {
}

// The jsii proxy struct for JsonPatch
type jsiiProxy_JsonPatch struct {
	_ byte // padding
}

// Adds a value to an object or inserts it into an array.
//
// In the case of an
// array, the value is inserted before the given index. The - character can be
// used instead of an index to insert at the end of an array.
//
// TODO: EXAMPLE
//
func JsonPatch_Add(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"add",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

// Applies a set of JSON-Patch (RFC-6902) operations to `document` and returns the result.
//
// Returns: The result document
func JsonPatch_Apply(document interface{}, ops ...JsonPatch) interface{} {
	_init_.Initialize()

	args := []interface{}{document}
	for _, a := range ops {
		args = append(args, a)
	}

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"apply",
		args,
		&returns,
	)

	return returns
}

// Copies a value from one location to another within the JSON document.
//
// Both
// from and path are JSON Pointers.
//
// TODO: EXAMPLE
//
func JsonPatch_Copy(from *string, path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"copy",
		[]interface{}{from, path},
		&returns,
	)

	return returns
}

// Moves a value from one location to the other.
//
// Both from and path are JSON Pointers.
//
// TODO: EXAMPLE
//
func JsonPatch_Move(from *string, path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"move",
		[]interface{}{from, path},
		&returns,
	)

	return returns
}

// Removes a value from an object or array.
//
// TODO: EXAMPLE
//
func JsonPatch_Remove(path *string) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"remove",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Replaces a value.
//
// Equivalent to a “remove” followed by an “add”.
//
// TODO: EXAMPLE
//
func JsonPatch_Replace(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"replace",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

// Tests that the specified value is set in the document.
//
// If the test fails,
// then the patch as a whole should not apply.
//
// TODO: EXAMPLE
//
func JsonPatch_Test(path *string, value interface{}) JsonPatch {
	_init_.Initialize()

	var returns JsonPatch

	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"test",
		[]interface{}{path, value},
		&returns,
	)

	return returns
}

type Lazy interface {
	Produce() interface{}
}

// The jsii proxy struct for Lazy
type jsiiProxy_Lazy struct {
	_ byte // padding
}

func Lazy_Any(producer IAnyProducer) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk8s.Lazy",
		"any",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lazy) Produce() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"produce",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for name generation.
type NameOptions struct {
	// Delimiter to use between components.
	Delimiter *string `json:"delimiter"`
	// Extra components to include in the name.
	Extra *[]*string `json:"extra"`
	// Include a short hash as last part of the name.
	IncludeHash *bool `json:"includeHash"`
	// Maximum allowed length for the name.
	MaxLen *float64 `json:"maxLen"`
}

// Utilities for generating unique and stable names.
type Names interface {
}

// The jsii proxy struct for Names
type jsiiProxy_Names struct {
	_ byte // padding
}

// Generates a unique and stable name compatible DNS_LABEL from RFC-1123 from a path.
//
// The generated name will:
//   - contain at most 63 characters
//   - contain only lowercase alphanumeric characters or ‘-’
//   - start with an alphanumeric character
//   - end with an alphanumeric character
//
// The generated name will have the form:
//   <comp0>-<comp1>-..-<compN>-<short-hash>
//
// Where <comp> are the path components (assuming they are is separated by
// "/").
//
// Note that if the total length is longer than 63 characters, we will trim
// the first components since the last components usually encode more meaning.
func Names_ToDnsLabel(scope constructs.Construct, options *NameOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdk8s.Names",
		"toDnsLabel",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Generates a unique and stable name compatible label key name segment and label value from a path.
//
// The name segment is required and must be 63 characters or less, beginning
// and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-),
// underscores (_), dots (.), and alphanumerics between.
//
// Valid label values must be 63 characters or less and must be empty or
// begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes
// (-), underscores (_), dots (.), and alphanumerics between.
//
// The generated name will have the form:
//   <comp0><delim><comp1><delim>..<delim><compN><delim><short-hash>
//
// Where <comp> are the path components (assuming they are is separated by
// "/").
//
// Note that if the total length is longer than 63 characters, we will trim
// the first components since the last components usually encode more meaning.
func Names_ToLabelValue(scope constructs.Construct, options *NameOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdk8s.Names",
		"toLabelValue",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

// Represents the amount of digital storage.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative.
//
// When the amount is passed as a token, unit conversion is not possible.
type Size interface {
	ToGibibytes(opts *SizeConversionOptions) *float64
	ToKibibytes(opts *SizeConversionOptions) *float64
	ToMebibytes(opts *SizeConversionOptions) *float64
	ToPebibytes(opts *SizeConversionOptions) *float64
	ToTebibytes(opts *SizeConversionOptions) *float64
}

// The jsii proxy struct for Size
type jsiiProxy_Size struct {
	_ byte // padding
}

// Create a Storage representing an amount gibibytes.
//
// 1 GiB = 1024 MiB
func Size_Gibibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"gibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount kibibytes.
//
// 1 KiB = 1024 bytes
func Size_Kibibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"kibibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount mebibytes.
//
// 1 MiB = 1024 KiB
func Size_Mebibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"mebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount pebibytes.
//
// 1 PiB = 1024 TiB
func Size_Pebibyte(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"pebibyte",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Create a Storage representing an amount tebibytes.
//
// 1 TiB = 1024 GiB
func Size_Tebibytes(amount *float64) Size {
	_init_.Initialize()

	var returns Size

	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"tebibytes",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Return this storage as a total number of gibibytes.
func (s *jsiiProxy_Size) ToGibibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toGibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return this storage as a total number of kibibytes.
func (s *jsiiProxy_Size) ToKibibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toKibibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return this storage as a total number of mebibytes.
func (s *jsiiProxy_Size) ToMebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toMebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return this storage as a total number of pebibytes.
func (s *jsiiProxy_Size) ToPebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toPebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Return this storage as a total number of tebibytes.
func (s *jsiiProxy_Size) ToTebibytes(opts *SizeConversionOptions) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"toTebibytes",
		[]interface{}{opts},
		&returns,
	)

	return returns
}

// Options for how to convert time to a different unit.
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	Rounding SizeRoundingBehavior `json:"rounding"`
}

// Rounding behaviour when converting between units of `Size`.
type SizeRoundingBehavior string

const (
	SizeRoundingBehavior_FAIL SizeRoundingBehavior = "FAIL"
	SizeRoundingBehavior_FLOOR SizeRoundingBehavior = "FLOOR"
	SizeRoundingBehavior_NONE SizeRoundingBehavior = "NONE"
)

// Testing utilities for cdk8s applications.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Returns an app for testing with the following properties: - Output directory is a temp dir.
func Testing_App() App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"app",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns: a Chart that can be used for tests
func Testing_Chart() Chart {
	_init_.Initialize()

	var returns Chart

	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"chart",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns the Kubernetes manifest synthesized from this chart.
func Testing_Synth(chart Chart) *[]interface{} {
	_init_.Initialize()

	var returns *[]interface{}

	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"synth",
		[]interface{}{chart},
		&returns,
	)

	return returns
}

// Options for how to convert time to a different unit.
type TimeConversionOptions struct {
	// If `true`, conversions into a larger time unit (e.g. `Seconds` to `Minutes`) will fail if the result is not an integer.
	Integral *bool `json:"integral"`
}

// YAML utilities.
type Yaml interface {
}

// The jsii proxy struct for Yaml
type jsiiProxy_Yaml struct {
	_ byte // padding
}

// Downloads a set of YAML documents (k8s manifest for example) from a URL or a file and returns them as javascript objects.
//
// Empty documents are filtered out.
//
// Returns: an array of objects, each represents a document inside the YAML
func Yaml_Load(urlOrFile *string) *[]interface{} {
	_init_.Initialize()

	var returns *[]interface{}

	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"load",
		[]interface{}{urlOrFile},
		&returns,
	)

	return returns
}

// Saves a set of objects as a multi-document YAML file.
func Yaml_Save(filePath *string, docs *[]interface{}) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"cdk8s.Yaml",
		"save",
		[]interface{}{filePath, docs},
	)
}

// Stringify a document into yaml.
func Yaml_Stringify(doc interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"stringify",
		[]interface{}{doc},
		&returns,
	)

	return returns
}

// Saves a set of YAML documents into a temp file (in /tmp).
//
// Returns: the path to the temporary file
func Yaml_Tmp(docs *[]interface{}) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"tmp",
		[]interface{}{docs},
		&returns,
	)

	return returns
}

