package models

import (
    "encoding/json"
    "testing"
)

func TestEmailTemplateListModel(t *testing.T) {
    model := EmailTemplateList{        Total: 5,        Templates: []EmailTemplate{EmailTemplate{        TemplateId: "verification",        Locale: "en_us",        Message: "Click on the link to verify your account.",        SenderName: "My User",        SenderEmail: "mail@appwrite.io",        ReplyToEmail: "emails@appwrite.io",        ReplyToName: "Support Team",        Subject: "Please verify your email address",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result EmailTemplateList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
