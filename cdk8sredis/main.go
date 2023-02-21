package cdk8sredis

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s-redis.Redis",
		reflect.TypeOf((*Redis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "primaryHost", GoGetter: "PrimaryHost"},
			_jsii_.MemberProperty{JsiiProperty: "replicaHost", GoGetter: "ReplicaHost"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Redis{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-redis.RedisOptions",
		reflect.TypeOf((*RedisOptions)(nil)).Elem(),
	)
}
