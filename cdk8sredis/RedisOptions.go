package cdk8sredis


// Experimental.
type RedisOptions struct {
	// Extra labels to associate with resources.
	// Default: - none.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Number of replicas.
	// Default: 2.
	//
	// Experimental.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

