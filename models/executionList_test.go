package models

import (
    "encoding/json"
    "testing"
)

func TestExecutionListModel(t *testing.T) {
    model := ExecutionList{        Total: 5,        Executions: []Execution{Execution{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        FunctionId: "5e5ea6g16897e",        DeploymentId: "5e5ea5c16897e",        Trigger: "http",        Status: "processing",        RequestMethod: "GET",        RequestPath: "/articles?id=5",        RequestHeaders: []Headers{Headers{        Name: "Content-Type",        Value: "application/json",    },
            },        ResponseStatusCode: 200,        ResponseBody: "string",        ResponseHeaders: []Headers{Headers{        Name: "Content-Type",        Value: "application/json",    },
            },        Logs: "string",        Errors: "string",        Duration: 0.4,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ExecutionList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
