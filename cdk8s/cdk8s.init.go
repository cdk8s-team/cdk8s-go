package cdk8s

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s.ApiObject",
		reflect.TypeOf((*ApiObject)(nil)).Elem(),
		func() interface{} {
			a := apiObject{}
			_jsii_.InitJsiiProxy(&a.Construct)
			return &a
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.ApiObjectMetadata",
		reflect.TypeOf((*ApiObjectMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.ApiObjectMetadataDefinition",
		reflect.TypeOf((*ApiObjectMetadataDefinition)(nil)).Elem(),
		func() interface{} {
			return &apiObjectMetadataDefinition{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.ApiObjectProps",
		reflect.TypeOf((*ApiObjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		func() interface{} {
			a := app{}
			_jsii_.InitJsiiProxy(&a.Construct)
			return &a
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.AppProps",
		reflect.TypeOf((*AppProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.Chart",
		reflect.TypeOf((*Chart)(nil)).Elem(),
		func() interface{} {
			c := chart{}
			_jsii_.InitJsiiProxy(&c.Construct)
			return &c
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.ChartProps",
		reflect.TypeOf((*ChartProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.DependencyGraph",
		reflect.TypeOf((*DependencyGraph)(nil)).Elem(),
		func() interface{} {
			return &dependencyGraph{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.DependencyVertex",
		reflect.TypeOf((*DependencyVertex)(nil)).Elem(),
		func() interface{} {
			return &dependencyVertex{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.Duration",
		reflect.TypeOf((*Duration)(nil)).Elem(),
		func() interface{} {
			return &duration{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.GroupVersionKind",
		reflect.TypeOf((*GroupVersionKind)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.Helm",
		reflect.TypeOf((*Helm)(nil)).Elem(),
		func() interface{} {
			h := helm{}
			_jsii_.InitJsiiProxy(&h.include)
			return &h
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.HelmProps",
		reflect.TypeOf((*HelmProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk8s.IAnyProducer",
		reflect.TypeOf((*IAnyProducer)(nil)).Elem(),
		func() interface{} {
			return &iAnyProducer{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.Include",
		reflect.TypeOf((*Include)(nil)).Elem(),
		func() interface{} {
			i := include{}
			_jsii_.InitJsiiProxy(&i.Construct)
			return &i
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.IncludeProps",
		reflect.TypeOf((*IncludeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.JsonPatch",
		reflect.TypeOf((*JsonPatch)(nil)).Elem(),
		func() interface{} {
			return &jsonPatch{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.Lazy",
		reflect.TypeOf((*Lazy)(nil)).Elem(),
		func() interface{} {
			return &lazy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.NameOptions",
		reflect.TypeOf((*NameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.Names",
		reflect.TypeOf((*Names)(nil)).Elem(),
		func() interface{} {
			return &names{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.Size",
		reflect.TypeOf((*Size)(nil)).Elem(),
		func() interface{} {
			return &size{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.SizeConversionOptions",
		reflect.TypeOf((*SizeConversionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk8s.SizeRoundingBehavior",
		reflect.TypeOf((*SizeRoundingBehavior)(nil)).Elem(),
		map[string]interface{}{
			"FAIL": SizeRoundingBehavior_FAIL,
			"FLOOR": SizeRoundingBehavior_FLOOR,
			"NONE": SizeRoundingBehavior_NONE,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s.Testing",
		reflect.TypeOf((*Testing)(nil)).Elem(),
		func() interface{} {
			return &testing{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s.TimeConversionOptions",
		reflect.TypeOf((*TimeConversionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s.Yaml",
		reflect.TypeOf((*Yaml)(nil)).Elem(),
		func() interface{} {
			return &yaml{}
		},
	)
}
