package cdk8sredis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-redis-go/cdk8sredis/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-redis-go/cdk8sredis/internal"
)

// Experimental.
type Redis interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The DNS host for the primary service.
	// Experimental.
	PrimaryHost() *string
	// The DNS host for the replica service.
	// Experimental.
	ReplicaHost() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Redis
type jsiiProxy_Redis struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Redis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) PrimaryHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Redis) ReplicaHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaHost",
		&returns,
	)
	return returns
}


// Experimental.
func NewRedis(scope constructs.Construct, id *string, options *RedisOptions) Redis {
	_init_.Initialize()

	if err := validateNewRedisParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Redis{}

	_jsii_.Create(
		"cdk8s-redis.Redis",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRedis_Override(r Redis, scope constructs.Construct, id *string, options *RedisOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-redis.Redis",
		[]interface{}{scope, id, options},
		r,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Redis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedis_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-redis.Redis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Redis) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

