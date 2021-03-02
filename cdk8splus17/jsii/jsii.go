package jsii

import (
	_          "embed"
	"sync"

	_jsii_     "github.com/aws/jsii-runtime-go"

	constructs "github.com/aws/constructs-go/constructs/v3/jsii"
	cdk8s      "github.com/awslabs/cdk8s-go/cdk8s/jsii"
)

//go:embed cdk8s-plus-17-1.0.0-beta.10.tgz
var tarball []byte
var once    sync.Once

// Initialize performs the necessary work for the enclosing
// module to be loaded in the jsii kernel.
func Initialize() {
	once.Do(func(){
		// Ensure all dependencies are initialized
		cdk8s.Initialize()
		constructs.Initialize()

		// Load this library into the kernel
		_jsii_.Load("cdk8s-plus-17", "1.0.0-beta.10", tarball)
	})
}
