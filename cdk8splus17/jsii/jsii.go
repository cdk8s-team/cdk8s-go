package jsii

import (
	rt "github.com/aws/jsii-runtime-go"
	"sync"
	// Initialization endpoints of dependencies
	constructs "github.com/aws/constructs-go/constructs/v3/jsii"
	cdk8s "github.com/awslabs/cdk8s-go/cdk8s/jsii"
)

var once sync.Once

// Initialize performs the necessary work for the enclosing
// module to be loaded in the jsii kernel.
func Initialize() {
	once.Do(func(){
		// Ensure all dependencies are initialized
		cdk8s.Initialize()
		constructs.Initialize()

		// Load this library into the kernel
		rt.Load("cdk8s-plus-17", "1.0.0-beta.8", tarball)
	})
}
