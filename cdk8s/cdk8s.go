// Cloud Development Kit for Kubernetes
package cdk8s

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/awslabs/cdk8s-go/cdk8s/jsii"

	"github.com/aws/constructs-go/constructs/v3"
)

// Experimental.
type ApiObject interface {
	constructs.Construct
	ApiGroup() string
	ApiVersion() string
	Chart() Chart
	Kind() string
	Metadata() ApiObjectMetadataDefinition
	Name() string
	AddDependency(dependencies constructs.IConstruct)
	AddJsonPatch(ops JsonPatch)
	ToJson() interface{}
}

// The jsii proxy struct for ApiObject
type apiObject struct {
	constructs.Construct // extends constructs.Construct
}

func (a *apiObject) ApiGroup() string {
	var returns string
	_jsii_.Get(
		a,
		"apiGroup",
		&returns,
	)
	return returns
}

func (a *apiObject) ApiVersion() string {
	var returns string
	_jsii_.Get(
		a,
		"apiVersion",
		&returns,
	)
	return returns
}

func (a *apiObject) Chart() Chart {
	var returns Chart
	_jsii_.Get(
		a,
		"chart",
		&returns,
	)
	return returns
}

func (a *apiObject) Kind() string {
	var returns string
	_jsii_.Get(
		a,
		"kind",
		&returns,
	)
	return returns
}

func (a *apiObject) Metadata() ApiObjectMetadataDefinition {
	var returns ApiObjectMetadataDefinition
	_jsii_.Get(
		a,
		"metadata",
		&returns,
	)
	return returns
}

func (a *apiObject) Name() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
	)
	return returns
}


// Defines an API object.
func NewApiObject(scope constructs.Construct, id string, props ApiObjectProps) ApiObject {
	_init_.Initialize()
	a := apiObject{}

	_jsii_.Create(
		"cdk8s.ApiObject",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&a,
	)
	return &a
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func ApiObject_Of(c constructs.IConstruct) ApiObject {
	_init_.Initialize()
	var returns ApiObject
	_jsii_.StaticInvoke(
		"cdk8s.ApiObject",
		"of",
		[]interface{}{c},
		true,
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (a *apiObject) AddDependency(dependencies constructs.IConstruct) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addDependency",
		[]interface{}{dependencies},
		false,
		&returns,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
// Experimental.
func (a *apiObject) AddJsonPatch(ops JsonPatch) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addJsonPatch",
		[]interface{}{ops},
		false,
		&returns,
	)
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
// Experimental.
func (a *apiObject) ToJson() interface{} {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"toJson",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Metadata associated with this object.
// Experimental.
type ApiObjectMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be
	// preserved when modifying objects.
	// See: http://kubernetes.io/docs/user-guide/annotations
	//
	// Experimental.
	Annotations map[string]string `json:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers and services.
	// See: http://kubernetes.io/docs/user-guide/labels
	//
	// Experimental.
	Labels map[string]string `json:"labels"`
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
	// Experimental.
	Name string `json:"name"`
	// Namespace defines the space within each name must be unique.
	//
	// An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty. Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
	// Experimental.
	Namespace string `json:"namespace"`
}

// Object metadata.
// Experimental.
type ApiObjectMetadataDefinition interface {
	Name() string
	Namespace() string
	Add(key string, value interface{})
	AddAnnotation(key string, value string)
	AddLabel(key string, value string)
	GetLabel(key string) string
	ToJson() interface{}
}

// The jsii proxy struct for ApiObjectMetadataDefinition
type apiObjectMetadataDefinition struct {
	_ byte // padding
}

func (a *apiObjectMetadataDefinition) Name() string {
	var returns string
	_jsii_.Get(
		a,
		"name",
		&returns,
	)
	return returns
}

func (a *apiObjectMetadataDefinition) Namespace() string {
	var returns string
	_jsii_.Get(
		a,
		"namespace",
		&returns,
	)
	return returns
}


func NewApiObjectMetadataDefinition(options ApiObjectMetadata) ApiObjectMetadataDefinition {
	_init_.Initialize()
	a := apiObjectMetadataDefinition{}

	_jsii_.Create(
		"cdk8s.ApiObjectMetadataDefinition",
		[]interface{}{options},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&a,
	)
	return &a
}

// Adds an arbitrary key/value to the object metadata.
// Experimental.
func (a *apiObjectMetadataDefinition) Add(key string, value interface{}) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"add",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Add an annotation.
// Experimental.
func (a *apiObjectMetadataDefinition) AddAnnotation(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addAnnotation",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Add a label.
// Experimental.
func (a *apiObjectMetadataDefinition) AddLabel(key string, value string) {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"addLabel",
		[]interface{}{key, value},
		false,
		&returns,
	)
}

// Returns: a value of a label or undefined
// Experimental.
func (a *apiObjectMetadataDefinition) GetLabel(key string) string {
	var returns string
	_jsii_.Invoke(
		a,
		"getLabel",
		[]interface{}{key},
		true,
		&returns,
	)
	return returns
}

// Synthesizes a k8s ObjectMeta for this metadata set.
// Experimental.
func (a *apiObjectMetadataDefinition) ToJson() interface{} {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"toJson",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Options for defining API objects.
// Experimental.
type ApiObjectProps struct {
	// API version.
	// Experimental.
	ApiVersion string `json:"apiVersion"`
	// Resource kind.
	// Experimental.
	Kind string `json:"kind"`
	// Object metadata.
	//
	// If `name` is not specified, an app-unique name will be allocated by the
	// framework based on the path of the construct within thes construct tree.
	// Experimental.
	Metadata ApiObjectMetadata `json:"metadata"`
}

// Represents a cdk8s application.
// Experimental.
type App interface {
	constructs.Construct
	Outdir() string
	Synth()
}

// The jsii proxy struct for App
type app struct {
	constructs.Construct // extends constructs.Construct
}

func (a *app) Outdir() string {
	var returns string
	_jsii_.Get(
		a,
		"outdir",
		&returns,
	)
	return returns
}


// Defines an app.
func NewApp(props AppProps) App {
	_init_.Initialize()
	a := app{}

	_jsii_.Create(
		"cdk8s.App",
		[]interface{}{props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&a,
	)
	return &a
}

// Synthesizes all manifests to the output directory.
// Experimental.
func (a *app) Synth() {
	var returns interface{}
	_jsii_.Invoke(
		a,
		"synth",
		[]interface{}{},
		false,
		&returns,
	)
}

// Experimental.
type AppProps struct {
	// The directory to output Kubernetes manifests.
	// Experimental.
	Outdir string `json:"outdir"`
}

// Experimental.
type Chart interface {
	constructs.Construct
	Labels() map[string]string
	Namespace() string
	AddDependency(dependencies constructs.IConstruct)
	GenerateObjectName(apiObject ApiObject) string
	ToJson() []interface{}
}

// The jsii proxy struct for Chart
type chart struct {
	constructs.Construct // extends constructs.Construct
}

func (c *chart) Labels() map[string]string {
	var returns map[string]string
	_jsii_.Get(
		c,
		"labels",
		&returns,
	)
	return returns
}

func (c *chart) Namespace() string {
	var returns string
	_jsii_.Get(
		c,
		"namespace",
		&returns,
	)
	return returns
}


func NewChart(scope constructs.Construct, id string, props ChartProps) Chart {
	_init_.Initialize()
	c := chart{}

	_jsii_.Create(
		"cdk8s.Chart",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&c,
	)
	return &c
}

// Finds the chart in which a node is defined.
// Experimental.
func Chart_Of(c constructs.IConstruct) Chart {
	_init_.Initialize()
	var returns Chart
	_jsii_.StaticInvoke(
		"cdk8s.Chart",
		"of",
		[]interface{}{c},
		true,
		&returns,
	)
	return returns
}

// Create a dependency between this Chart and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (c *chart) AddDependency(dependencies constructs.IConstruct) {
	var returns interface{}
	_jsii_.Invoke(
		c,
		"addDependency",
		[]interface{}{dependencies},
		false,
		&returns,
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
// Experimental.
func (c *chart) GenerateObjectName(apiObject ApiObject) string {
	var returns string
	_jsii_.Invoke(
		c,
		"generateObjectName",
		[]interface{}{apiObject},
		true,
		&returns,
	)
	return returns
}

// Renders this chart to a set of Kubernetes JSON resources.
//
// Returns: array of resource manifests
// Experimental.
func (c *chart) ToJson() []interface{} {
	var returns []interface{}
	_jsii_.Invoke(
		c,
		"toJson",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Experimental.
type ChartProps struct {
	// Labels to apply to all resources in this chart.
	// Experimental.
	Labels map[string]string `json:"labels"`
	// The default namespace for all objects defined in this chart (directly or indirectly).
	//
	// This namespace will only apply to objects that don't have a
	// `namespace` explicitly defined for them.
	// Experimental.
	Namespace string `json:"namespace"`
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
// Experimental.
type DependencyGraph interface {
	Root() DependencyVertex
	Topology() []constructs.IConstruct
}

// The jsii proxy struct for DependencyGraph
type dependencyGraph struct {
	_ byte // padding
}

func (d *dependencyGraph) Root() DependencyVertex {
	var returns DependencyVertex
	_jsii_.Get(
		d,
		"root",
		&returns,
	)
	return returns
}


func NewDependencyGraph(node constructs.Node) DependencyGraph {
	_init_.Initialize()
	d := dependencyGraph{}

	_jsii_.Create(
		"cdk8s.DependencyGraph",
		[]interface{}{node},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&d,
	)
	return &d
}

// See: Vertex.topology()
//
// Experimental.
func (d *dependencyGraph) Topology() []constructs.IConstruct {
	var returns []constructs.IConstruct
	_jsii_.Invoke(
		d,
		"topology",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Represents a vertex in the graph.
//
// The value of each vertex is an `IConstruct` that is accessible via the `.value` getter.
// Experimental.
type DependencyVertex interface {
	Inbound() []DependencyVertex
	Outbound() []DependencyVertex
	Value() constructs.IConstruct
	AddChild(dep DependencyVertex)
	Topology() []constructs.IConstruct
}

// The jsii proxy struct for DependencyVertex
type dependencyVertex struct {
	_ byte // padding
}

func (d *dependencyVertex) Inbound() []DependencyVertex {
	var returns []DependencyVertex
	_jsii_.Get(
		d,
		"inbound",
		&returns,
	)
	return returns
}

func (d *dependencyVertex) Outbound() []DependencyVertex {
	var returns []DependencyVertex
	_jsii_.Get(
		d,
		"outbound",
		&returns,
	)
	return returns
}

func (d *dependencyVertex) Value() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		d,
		"value",
		&returns,
	)
	return returns
}


func NewDependencyVertex(value constructs.IConstruct) DependencyVertex {
	_init_.Initialize()
	d := dependencyVertex{}

	_jsii_.Create(
		"cdk8s.DependencyVertex",
		[]interface{}{value},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&d,
	)
	return &d
}

// Adds a vertex as a dependency of the current node.
//
// Also updates the parents of `dep`, so that it contains this node as a parent.
//
// This operation will fail in case it creates a cycle in the graph.
// Experimental.
func (d *dependencyVertex) AddChild(dep DependencyVertex) {
	var returns interface{}
	_jsii_.Invoke(
		d,
		"addChild",
		[]interface{}{dep},
		false,
		&returns,
	)
}

// Returns a topologically sorted array of the constructs in the sub-graph.
// Experimental.
func (d *dependencyVertex) Topology() []constructs.IConstruct {
	var returns []constructs.IConstruct
	_jsii_.Invoke(
		d,
		"topology",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Represents a length of time.
//
// The amount can be specified either as a literal value (e.g: `10`) which
// cannot be negative.
// Experimental.
type Duration interface {
	ToDays(opts TimeConversionOptions) float64
	ToHours(opts TimeConversionOptions) float64
	ToHumanString() string
	ToIsoString() string
	ToMilliseconds(opts TimeConversionOptions) float64
	ToMinutes(opts TimeConversionOptions) float64
	ToSeconds(opts TimeConversionOptions) float64
}

// The jsii proxy struct for Duration
type duration struct {
	_ byte // padding
}

// Create a Duration representing an amount of days.
//
// Returns: a new `Duration` representing `amount` Days.
// Experimental.
func Duration_Days(amount float64) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"days",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Duration representing an amount of hours.
//
// Returns: a new `Duration` representing `amount` Hours.
// Experimental.
func Duration_Hours(amount float64) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"hours",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Duration representing an amount of milliseconds.
//
// Returns: a new `Duration` representing `amount` ms.
// Experimental.
func Duration_Millis(amount float64) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"millis",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Duration representing an amount of minutes.
//
// Returns: a new `Duration` representing `amount` Minutes.
// Experimental.
func Duration_Minutes(amount float64) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"minutes",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Parse a period formatted according to the ISO 8601 standard.
//
// Returns: the parsed `Duration`.
// See: https://www.iso.org/fr/standard/70907.html
//
// Experimental.
func Duration_Parse(duration string) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"parse",
		[]interface{}{duration},
		true,
		&returns,
	)
	return returns
}

// Create a Duration representing an amount of seconds.
//
// Returns: a new `Duration` representing `amount` Seconds.
// Experimental.
func Duration_Seconds(amount float64) Duration {
	_init_.Initialize()
	var returns Duration
	_jsii_.StaticInvoke(
		"cdk8s.Duration",
		"seconds",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Return the total number of days in this Duration.
//
// Returns: the value of this `Duration` expressed in Days.
// Experimental.
func (d *duration) ToDays(opts TimeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		d,
		"toDays",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return the total number of hours in this Duration.
//
// Returns: the value of this `Duration` expressed in Hours.
// Experimental.
func (d *duration) ToHours(opts TimeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		d,
		"toHours",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Turn this duration into a human-readable string.
// Experimental.
func (d *duration) ToHumanString() string {
	var returns string
	_jsii_.Invoke(
		d,
		"toHumanString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Return an ISO 8601 representation of this period.
//
// Returns: a string starting with 'PT' describing the period
// See: https://www.iso.org/fr/standard/70907.html
//
// Experimental.
func (d *duration) ToIsoString() string {
	var returns string
	_jsii_.Invoke(
		d,
		"toIsoString",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Return the total number of milliseconds in this Duration.
//
// Returns: the value of this `Duration` expressed in Milliseconds.
// Experimental.
func (d *duration) ToMilliseconds(opts TimeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		d,
		"toMilliseconds",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return the total number of minutes in this Duration.
//
// Returns: the value of this `Duration` expressed in Minutes.
// Experimental.
func (d *duration) ToMinutes(opts TimeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		d,
		"toMinutes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return the total number of seconds in this Duration.
//
// Returns: the value of this `Duration` expressed in Seconds.
// Experimental.
func (d *duration) ToSeconds(opts TimeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		d,
		"toSeconds",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Experimental.
type GroupVersionKind struct {
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	// Experimental.
	ApiVersion string `json:"apiVersion"`
	// The object kind.
	// Experimental.
	Kind string `json:"kind"`
}

// Represents a Helm deployment.
//
// Use this construct to import an existing Helm chart and incorporate it into your constructs.
// Experimental.
type Helm interface {
	Include
	ReleaseName() string
}

// The jsii proxy struct for Helm
type helm struct {
	include // extends cdk8s.Include
}

func (h *helm) ReleaseName() string {
	var returns string
	_jsii_.Get(
		h,
		"releaseName",
		&returns,
	)
	return returns
}


func NewHelm(scope constructs.Construct, id string, props HelmProps) Helm {
	_init_.Initialize()
	h := helm{}

	_jsii_.Create(
		"cdk8s.Helm",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&h,
	)
	return &h
}

// Options for `Helm`.
// Experimental.
type HelmProps struct {
	// The chart name to use. It can be a chart from a helm repository or a local directory.
	//
	// This name is passed to `helm template` and has all the relevant semantics.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Chart string `json:"chart"`
	// The local helm executable to use in order to create the manifest the chart.
	// Experimental.
	HelmExecutable string `json:"helmExecutable"`
	// Additional flags to add to the `helm` execution.
	// Experimental.
	HelmFlags []string `json:"helmFlags"`
	// The release name.
	// See: https://helm.sh/docs/intro/using_helm/#three-big-concepts
	//
	// Experimental.
	ReleaseName string `json:"releaseName"`
	// Values to pass to the chart.
	// Experimental.
	Values map[string]interface{} `json:"values"`
}

// Experimental.
type IAnyProducer interface {
	// Experimental.
	Produce() interface{}
}

// The jsii proxy for IAnyProducer
type iAnyProducer struct {
	_ byte // padding
}

func (i *iAnyProducer) Produce() interface{} {
	var returns interface{}
	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Reads a YAML manifest from a file or a URL and defines all resources as API objects within the defined scope.
//
// The names (`metadata.name`) of imported resources will be preserved as-is
// from the manifest.
// Experimental.
type Include interface {
	constructs.Construct
	ApiObjects() []ApiObject
}

// The jsii proxy struct for Include
type include struct {
	constructs.Construct // extends constructs.Construct
}

func (i *include) ApiObjects() []ApiObject {
	var returns []ApiObject
	_jsii_.Get(
		i,
		"apiObjects",
		&returns,
	)
	return returns
}


func NewInclude(scope constructs.Construct, id string, props IncludeProps) Include {
	_init_.Initialize()
	i := include{}

	_jsii_.Create(
		"cdk8s.Include",
		[]interface{}{scope, id, props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&i,
	)
	return &i
}

// Experimental.
type IncludeProps struct {
	// Local file path or URL which includes a Kubernetes YAML manifest.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Url string `json:"url"`
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
// Experimental.
type JsonPatch interface {
}

// The jsii proxy struct for JsonPatch
type jsonPatch struct {
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
// Experimental.
func JsonPatch_Add(path string, value interface{}) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"add",
		[]interface{}{path, value},
		true,
		&returns,
	)
	return returns
}

// Applies a set of JSON-Patch (RFC-6902) operations to `document` and returns the result.
//
// Returns: The result document
// Experimental.
func JsonPatch_Apply(document interface{}, ops JsonPatch) interface{} {
	_init_.Initialize()
	var returns interface{}
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"apply",
		[]interface{}{document, ops},
		true,
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
// Experimental.
func JsonPatch_Copy(from string, path string) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"copy",
		[]interface{}{from, path},
		true,
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
// Experimental.
func JsonPatch_Move(from string, path string) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"move",
		[]interface{}{from, path},
		true,
		&returns,
	)
	return returns
}

// Removes a value from an object or array.
//
// TODO: EXAMPLE
//
// Experimental.
func JsonPatch_Remove(path string) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"remove",
		[]interface{}{path},
		true,
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
// Experimental.
func JsonPatch_Replace(path string, value interface{}) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"replace",
		[]interface{}{path, value},
		true,
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
// Experimental.
func JsonPatch_Test(path string, value interface{}) JsonPatch {
	_init_.Initialize()
	var returns JsonPatch
	_jsii_.StaticInvoke(
		"cdk8s.JsonPatch",
		"test",
		[]interface{}{path, value},
		true,
		&returns,
	)
	return returns
}

// Experimental.
type Lazy interface {
	Produce() interface{}
}

// The jsii proxy struct for Lazy
type lazy struct {
	_ byte // padding
}

// Experimental.
func Lazy_Any(producer IAnyProducer) interface{} {
	_init_.Initialize()
	var returns interface{}
	_jsii_.StaticInvoke(
		"cdk8s.Lazy",
		"any",
		[]interface{}{producer},
		true,
		&returns,
	)
	return returns
}

// Experimental.
func (l *lazy) Produce() interface{} {
	var returns interface{}
	_jsii_.Invoke(
		l,
		"produce",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Options for name generation.
// Experimental.
type NameOptions struct {
	// Delimiter to use between components.
	// Experimental.
	Delimiter string `json:"delimiter"`
	// Extra components to include in the name.
	// Experimental.
	Extra []string `json:"extra"`
	// Include a short hash as last part of the name.
	// Experimental.
	IncludeHash bool `json:"includeHash"`
	// Maximum allowed length for the name.
	// Experimental.
	MaxLen float64 `json:"maxLen"`
}

// Utilities for generating unique and stable names.
// Experimental.
type Names interface {
}

// The jsii proxy struct for Names
type names struct {
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
// Experimental.
func Names_ToDnsLabel(scope constructs.Construct, options NameOptions) string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticInvoke(
		"cdk8s.Names",
		"toDnsLabel",
		[]interface{}{scope, options},
		true,
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
// Experimental.
func Names_ToLabelValue(scope constructs.Construct, options NameOptions) string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticInvoke(
		"cdk8s.Names",
		"toLabelValue",
		[]interface{}{scope, options},
		true,
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
// Experimental.
type Size interface {
	ToGibibytes(opts SizeConversionOptions) float64
	ToKibibytes(opts SizeConversionOptions) float64
	ToMebibytes(opts SizeConversionOptions) float64
	ToPebibytes(opts SizeConversionOptions) float64
	ToTebibytes(opts SizeConversionOptions) float64
}

// The jsii proxy struct for Size
type size struct {
	_ byte // padding
}

// Create a Storage representing an amount gibibytes.
//
// 1 GiB = 1024 MiB
// Experimental.
func Size_Gibibytes(amount float64) Size {
	_init_.Initialize()
	var returns Size
	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"gibibytes",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Storage representing an amount kibibytes.
//
// 1 KiB = 1024 bytes
// Experimental.
func Size_Kibibytes(amount float64) Size {
	_init_.Initialize()
	var returns Size
	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"kibibytes",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Storage representing an amount mebibytes.
//
// 1 MiB = 1024 KiB
// Experimental.
func Size_Mebibytes(amount float64) Size {
	_init_.Initialize()
	var returns Size
	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"mebibytes",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Storage representing an amount pebibytes.
//
// 1 PiB = 1024 TiB
// Experimental.
func Size_Pebibyte(amount float64) Size {
	_init_.Initialize()
	var returns Size
	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"pebibyte",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Create a Storage representing an amount tebibytes.
//
// 1 TiB = 1024 GiB
// Experimental.
func Size_Tebibytes(amount float64) Size {
	_init_.Initialize()
	var returns Size
	_jsii_.StaticInvoke(
		"cdk8s.Size",
		"tebibytes",
		[]interface{}{amount},
		true,
		&returns,
	)
	return returns
}

// Return this storage as a total number of gibibytes.
// Experimental.
func (s *size) ToGibibytes(opts SizeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		s,
		"toGibibytes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return this storage as a total number of kibibytes.
// Experimental.
func (s *size) ToKibibytes(opts SizeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		s,
		"toKibibytes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return this storage as a total number of mebibytes.
// Experimental.
func (s *size) ToMebibytes(opts SizeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		s,
		"toMebibytes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return this storage as a total number of pebibytes.
// Experimental.
func (s *size) ToPebibytes(opts SizeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		s,
		"toPebibytes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Return this storage as a total number of tebibytes.
// Experimental.
func (s *size) ToTebibytes(opts SizeConversionOptions) float64 {
	var returns float64
	_jsii_.Invoke(
		s,
		"toTebibytes",
		[]interface{}{opts},
		true,
		&returns,
	)
	return returns
}

// Options for how to convert time to a different unit.
// Experimental.
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	// Experimental.
	Rounding SizeRoundingBehavior `json:"rounding"`
}

// Rounding behaviour when converting between units of `Size`.
// Experimental.
type SizeRoundingBehavior string

const (
	SizeRoundingBehavior_FAIL SizeRoundingBehavior = "FAIL"
	SizeRoundingBehavior_FLOOR SizeRoundingBehavior = "FLOOR"
	SizeRoundingBehavior_NONE SizeRoundingBehavior = "NONE"
)

// Testing utilities for cdk8s applications.
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type testing struct {
	_ byte // padding
}

// Returns an app for testing with the following properties: - Output directory is a temp dir.
// Experimental.
func Testing_App() App {
	_init_.Initialize()
	var returns App
	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"app",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns: a Chart that can be used for tests
// Experimental.
func Testing_Chart() Chart {
	_init_.Initialize()
	var returns Chart
	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"chart",
		[]interface{}{},
		true,
		&returns,
	)
	return returns
}

// Returns the Kubernetes manifest synthesized from this chart.
// Experimental.
func Testing_Synth(chart Chart) []interface{} {
	_init_.Initialize()
	var returns []interface{}
	_jsii_.StaticInvoke(
		"cdk8s.Testing",
		"synth",
		[]interface{}{chart},
		true,
		&returns,
	)
	return returns
}

// Options for how to convert time to a different unit.
// Experimental.
type TimeConversionOptions struct {
	// If `true`, conversions into a larger time unit (e.g. `Seconds` to `Minutes`) will fail if the result is not an integer.
	// Experimental.
	Integral bool `json:"integral"`
}

// YAML utilities.
// Experimental.
type Yaml interface {
}

// The jsii proxy struct for Yaml
type yaml struct {
	_ byte // padding
}

// Downloads a set of YAML documents (k8s manifest for example) from a URL or a file and returns them as javascript objects.
//
// Empty documents are filtered out.
//
// Returns: an array of objects, each represents a document inside the YAML
// Experimental.
func Yaml_Load(urlOrFile string) []interface{} {
	_init_.Initialize()
	var returns []interface{}
	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"load",
		[]interface{}{urlOrFile},
		true,
		&returns,
	)
	return returns
}

// Saves a set of objects as a multi-document YAML file.
// Experimental.
func Yaml_Save(filePath string, docs []interface{}) {
	_init_.Initialize()
	var returns interface{}
	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"save",
		[]interface{}{filePath, docs},
		false,
		&returns,
	)
}

// Saves a set of YAML documents into a temp file (in /tmp).
//
// Returns: the path to the temporary file
// Experimental.
func Yaml_Tmp(docs []interface{}) string {
	_init_.Initialize()
	var returns string
	_jsii_.StaticInvoke(
		"cdk8s.Yaml",
		"tmp",
		[]interface{}{docs},
		true,
		&returns,
	)
	return returns
}

