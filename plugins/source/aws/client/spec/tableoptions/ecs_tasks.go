package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type ECSTasks struct {
	ListTasksOpts []CustomECSListTasksInput `json:"list_tasks,omitempty"`
}

type CustomECSListTasksInput struct {
	ecs.ListTasksInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomECSListTasksInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (s *CustomECSListTasksInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &s.ListTasksInput)
}

// JSONSchemaExtend is required to remove `NextToken` & `Cluster`, as well as add default for `MaxResults`.
func (CustomECSListTasksInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("NextToken")
	sc.Properties.Delete("Cluster")

	sc.Properties.Value("MaxResults").Default = 100
}

func (s *ECSTasks) validateListTasks() error {
	for _, opt := range s.ListTasksOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in ListTasks")
		}

		if aws.ToString(opt.Cluster) != "" {
			return errors.New("invalid input: cannot set Cluster in ListTasks")
		}
	}
	return nil
}

func (s *ECSTasks) SetDefaults() {
	for i := 0; i < len(s.ListTasksOpts); i++ {
		if aws.ToInt32(s.ListTasksOpts[i].MaxResults) == 0 {
			s.ListTasksOpts[i].MaxResults = aws.Int32(100)
		}
	}
}

func (s *ECSTasks) Validate() error {
	return s.validateListTasks()
}
