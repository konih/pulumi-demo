package util

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestConvertToMapInput(t *testing.T) {
	testCases := []struct {
		name   string
		input  map[string]string
		expect pulumi.Map
	}{
		{
			name:   "Empty Map",
			input:  map[string]string{},
			expect: pulumi.Map{},
		},
		{
			name: "Non-Empty Map",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expect: pulumi.Map{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := ConvertToMapInput(tc.input)
			assert.Equal(t, tc.expect, output)
		})
	}
}
