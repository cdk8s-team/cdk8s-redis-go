//go:build no_runtime_type_checking

package cdk8sredis

// Building without runtime type checking enabled, so all the below just return nil

func validateRedis_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRedisParameters(scope constructs.Construct, id *string, options *RedisOptions) error {
	return nil
}

