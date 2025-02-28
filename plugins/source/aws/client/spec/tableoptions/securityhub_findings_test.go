package tableoptions

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFindings(t *testing.T) {
	u := CustomSecurityHubGetFindingsInput{}
	require.NoError(t, faker.FakeObject(&u))
	api := SecurityHubFindings{
		GetFindingsOpts: []CustomSecurityHubGetFindingsInput{u},
	}
	// Ensure that the validation works as expected
	err := api.Validate()
	assert.EqualError(t, err, "invalid input: cannot set NextToken in GetFindings")

	// Ensure that as soon as the validation passes that there are no unexpected empty or nil fields
	api.GetFindingsOpts[0].NextToken = nil
	err = api.Validate()
	assert.EqualError(t, err, "invalid range: MaxResults must be within range [1-100]")
}

func TestCustomSecurityHubGetFindingsInput_JSONSchemaExtend(t *testing.T) {
	schema, err := jsonschema.Generate(SecurityHubFindings{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "empty get_findings",
			Spec: `{"get_findings":[]}`,
		},
		{
			Name: "null get_findings",
			Spec: `{"get_findings":null}`,
		},
		{
			Name: "bad get_findings",
			Err:  true,
			Spec: `{"get_findings":123}`,
		},
		{
			Name: "empty get_findings entry",
			Spec: `{"get_findings":[{}]}`,
		},
		{
			Name: "null get_findings entry",
			Err:  true,
			Spec: `{"get_findings":[null]}`,
		},
		{
			Name: "bad get_findings entry",
			Err:  true,
			Spec: `{"get_findings":[123]}`,
		},
		{
			Name: "proper get_findings",
			Spec: func() string {
				var input CustomSecurityHubGetFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "get_findings.NextToken is present",
			Err:  true,
			Spec: func() string {
				var input CustomSecurityHubGetFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 10 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input) + `]}`
			}(),
		},
		{
			Name: "get_findings.MaxResults > 100",
			Err:  true,
			Spec: func() string {
				var input CustomSecurityHubGetFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 1000 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
		{
			Name: "get_findings.MaxResults < 1",
			Err:  true,
			Spec: func() string {
				var input CustomSecurityHubGetFindingsInput
				require.NoError(t, faker.FakeObject(&input))
				input.MaxResults = 0 // range 1-100
				return `{"get_findings":[` + jsonschema.WithRemovedKeys(t, &input, "NextToken") + `]}`
			}(),
		},
	})
}
