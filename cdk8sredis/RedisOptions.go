// Basic implementation of a Redis construct for cdk8s.
package cdk8sredis


// Experimental.
type RedisOptions struct {
	// Extra labels to associate with resources.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Number of replicas.
	// Experimental.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

