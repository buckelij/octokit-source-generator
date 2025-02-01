package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookConfigInsecureSsl composed type wrapper for classes float64, string
type WebhookConfigInsecureSsl struct {
    // Composed type representation for type float64
    double *float64
    // Composed type representation for type string
    string *string
}
// NewWebhookConfigInsecureSsl instantiates a new WebhookConfigInsecureSsl and sets the default values.
func NewWebhookConfigInsecureSsl()(*WebhookConfigInsecureSsl) {
    m := &WebhookConfigInsecureSsl{
    }
    return m
}
// CreateWebhookConfigInsecureSslFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebhookConfigInsecureSslFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWebhookConfigInsecureSsl()
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
    if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetDouble()(*float64) {
    return m.double
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebhookConfigInsecureSsl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WebhookConfigInsecureSsl) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *WebhookConfigInsecureSsl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetDouble(value *float64)() {
    m.double = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetString(value *string)() {
    m.string = value
}
type WebhookConfigInsecureSslable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDouble()(*float64)
    GetString()(*string)
    SetDouble(value *float64)()
    SetString(value *string)()
}
