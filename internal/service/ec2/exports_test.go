package ec2

// Exports for use in tests only.
var (
	ResourceInstanceConnectEndpoint  = newResourceInstanceConnectEndpoint
	ResourceSecurityGroupEgressRule  = newResourceSecurityGroupEgressRule
	ResourceSecurityGroupIngressRule = newResourceSecurityGroupIngressRule

	UpdateTags = updateTags
)
