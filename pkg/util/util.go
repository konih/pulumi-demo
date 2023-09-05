package util

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func ConvertToMapInput(input map[string]string) pulumi.MapInput {
	out := pulumi.Map{}
	for k, v := range input {
		out[k] = pulumi.String(v)
	}
	return out
}
