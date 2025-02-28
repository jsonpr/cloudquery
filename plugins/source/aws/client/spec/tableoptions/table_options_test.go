package tableoptions

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	accessanalyzertypes "github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailtypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	cloudwatchtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	costexplorertypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	inspector2types "github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	securityhubtypes "github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTableOptions_Validate(t *testing.T) {
	tOpts := TableOptions{}
	err := tOpts.Validate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tOpts.CloudTrailEvents = &CloudtrailEvents{
		LookupEventsOpts: []CustomCloudtrailLookupEventsInput{
			{
				LookupEventsInput: cloudtrail.LookupEventsInput{
					EndTime:          nil,
					EventCategory:    "",
					LookupAttributes: nil,
					MaxResults:       nil,
					NextToken:        aws.String("123"),
					StartTime:        nil,
				},
			},
		},
	}

	tOpts.CloudwatchMetrics = CloudwatchMetrics{
		CloudwatchMetric{
			ListMetricsOpts:         CloudwatchListMetricsInput{},
			GetMetricStatisticsOpts: []CloudwatchGetMetricStatisticsInput{},
		},
	}
	err = tOpts.Validate()
	if err == nil {
		t.Fatal("expected error validating cloud_trail_events, got nil")
	}
}

func TestTableOptions_SetDefaults(t *testing.T) {
	opts := &TableOptions{
		SecurityHubFindings: &SecurityHubFindings{
			GetFindingsOpts: []CustomSecurityHubGetFindingsInput{{}},
		},
		ECSTasks: &ECSTasks{
			ListTasksOpts: []CustomECSListTasksInput{{}},
		},
	}
	data, err := json.Marshal(opts)
	require.NoError(t, err)

	require.NotPanics(t, opts.SetDefaults)

	// check something did change
	dataWithDefaults, err := json.Marshal(opts)
	require.NoError(t, err)
	require.NotEqual(t, string(data), string(dataWithDefaults))
}

// TestTableOptionsUnmarshal tests that the TableOptions struct can be unmarshaled from JSON using
// snake_case keys.
func TestTableOptionsUnmarshal(t *testing.T) {
	tableOpts := TableOptions{}
	require.NoError(t, faker.FakeObject(&tableOpts))
	b, err := json.Marshal(tableOpts)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]any{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		t.Fatal(err)
	}
	changeCaseForObject(m, strings.ToLower)
	nb, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var got TableOptions
	err = json.Unmarshal(nb, &got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(tableOpts, got, cmpopts.IgnoreUnexported(
		accessanalyzer.ListFindingsInput{},
		accessanalyzertypes.SortCriteria{},
		accessanalyzertypes.Criterion{},
		cloudwatch.GetMetricStatisticsInput{},
		cloudwatch.ListMetricsInput{},
		cloudwatchtypes.Dimension{},
		cloudwatchtypes.DimensionFilter{},
		cloudtrail.LookupEventsInput{},
		cloudtrailtypes.LookupAttribute{},
		inspector2.ListFindingsInput{},
		inspector2types.StringFilter{},
		inspector2types.DateFilter{},
		inspector2types.NumberFilter{},
		inspector2types.PortRangeFilter{},
		inspector2types.MapFilter{},
		inspector2types.PackageFilter{},
		inspector2types.FilterCriteria{},
		inspector2types.SortCriteria{},
		costexplorertypes.DateInterval{},
		costexplorertypes.Expression{},
		costexplorertypes.CostCategoryValues{},
		costexplorertypes.DimensionValues{},
		costexplorertypes.TagValues{},
		costexplorertypes.GroupDefinition{},
		costexplorer.GetCostAndUsageInput{},
		securityhub.GetFindingsInput{},
		securityhubtypes.AwsSecurityFindingFilters{},
		securityhubtypes.StringFilter{},
		securityhubtypes.NumberFilter{},
		securityhubtypes.DateFilter{},
		securityhubtypes.KeywordFilter{},
		securityhubtypes.MapFilter{},
		securityhubtypes.IpFilter{},
		securityhubtypes.BooleanFilter{},
		securityhubtypes.SortCriterion{},
		ecs.ListTasksInput{},
	)); diff != "" {
		t.Fatalf("mismatch between objects after loading from snake case json: %v", diff)
	}
}

func TestJSONSchema(t *testing.T) {
	schema, err := jsonschema.Generate(TableOptions{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "all null",
			Spec: `{
  "aws_alpha_cloudwatch_metrics": null,
  "aws_cloudtrail_events": null,
  "aws_accessanalyzer_analyzer_findings": null,
  "aws_inspector2_findings": null,
  "aws_alpha_costexplorer_cost_custom": null,
  "aws_securityhub_findings": null,
  "aws_ecs_cluster_tasks": null
}`,
		},
	})
}
