package models

import (
    "encoding/json"
    "testing"
)

func TestTopicListModel(t *testing.T) {
    model := TopicList{        Total: 5,        Topics: []Topic{Topic{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "events",        EmailTotal: 100,        SmsTotal: 100,        PushTotal: 100,        Subscribe: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result TopicList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
