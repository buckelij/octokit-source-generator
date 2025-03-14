package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueLabelsArray composed type wrapper for classes []IssueLabelsArrayMember1able, []string
type IssueLabelsArray struct {
    // Composed type representation for type []IssueLabelsArrayMember1able
    issueLabelsArrayMember1 []IssueLabelsArrayMember1able
    // Composed type representation for type []string
    string []string
}
// NewIssueLabelsArray instantiates a new IssueLabelsArray and sets the default values.
func NewIssueLabelsArray()(*IssueLabelsArray) {
    m := &IssueLabelsArray{
    }
    return m
}
// CreateIssueLabelsArrayFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueLabelsArrayFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewIssueLabelsArray()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetCollectionOfObjectValues(CreateIssueLabelsArrayMember1FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]IssueLabelsArrayMember1able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(IssueLabelsArrayMember1able)
            }
        }
        result.SetIssueLabelsArrayMember1(cast)
    } else if val, err := parseNode.GetCollectionOfPrimitiveValues("string"); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]string, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = *(v.(*string))
            }
        }
        result.SetString(cast)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueLabelsArray) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *IssueLabelsArray) GetIsComposedType()(bool) {
    return true
}
// GetIssueLabelsArrayMember1 gets the issueLabelsArrayMember1 property value. Composed type representation for type []IssueLabelsArrayMember1able
// returns a []IssueLabelsArrayMember1able when successful
func (m *IssueLabelsArray) GetIssueLabelsArrayMember1()([]IssueLabelsArrayMember1able) {
    return m.issueLabelsArrayMember1
}
// GetString gets the string property value. Composed type representation for type []string
// returns a []string when successful
func (m *IssueLabelsArray) GetString()([]string) {
    return m.string
}
// Serialize serializes information the current object
func (m *IssueLabelsArray) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIssueLabelsArrayMember1() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIssueLabelsArrayMember1()))
        for i, v := range m.GetIssueLabelsArrayMember1() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteCollectionOfStringValues("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIssueLabelsArrayMember1 sets the issueLabelsArrayMember1 property value. Composed type representation for type []IssueLabelsArrayMember1able
func (m *IssueLabelsArray) SetIssueLabelsArrayMember1(value []IssueLabelsArrayMember1able)() {
    m.issueLabelsArrayMember1 = value
}
// SetString sets the string property value. Composed type representation for type []string
func (m *IssueLabelsArray) SetString(value []string)() {
    m.string = value
}
type IssueLabelsArrayable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssueLabelsArrayMember1()([]IssueLabelsArrayMember1able)
    GetString()([]string)
    SetIssueLabelsArrayMember1(value []IssueLabelsArrayMember1able)()
    SetString(value []string)()
}
