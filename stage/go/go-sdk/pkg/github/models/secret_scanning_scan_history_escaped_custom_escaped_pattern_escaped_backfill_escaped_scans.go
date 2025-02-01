package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningScanHistory_custom_pattern_backfill_scans struct {
    SecretScanningScan
    // Name of the custom pattern for custom pattern scans
    pattern_name *string
    // Level at which the custom pattern is defined, one of "repository", "organization", or "enterprise"
    pattern_scope *string
}
// NewSecretScanningScanHistory_custom_pattern_backfill_scans instantiates a new SecretScanningScanHistory_custom_pattern_backfill_scans and sets the default values.
func NewSecretScanningScanHistory_custom_pattern_backfill_scans()(*SecretScanningScanHistory_custom_pattern_backfill_scans) {
    m := &SecretScanningScanHistory_custom_pattern_backfill_scans{
        SecretScanningScan: *NewSecretScanningScan(),
    }
    return m
}
// CreateSecretScanningScanHistory_custom_pattern_backfill_scansFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningScanHistory_custom_pattern_backfill_scansFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningScanHistory_custom_pattern_backfill_scans(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecretScanningScan.GetFieldDeserializers()
    res["pattern_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatternName(val)
        }
        return nil
    }
    res["pattern_scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatternScope(val)
        }
        return nil
    }
    return res
}
// GetPatternName gets the pattern_name property value. Name of the custom pattern for custom pattern scans
// returns a *string when successful
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) GetPatternName()(*string) {
    return m.pattern_name
}
// GetPatternScope gets the pattern_scope property value. Level at which the custom pattern is defined, one of "repository", "organization", or "enterprise"
// returns a *string when successful
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) GetPatternScope()(*string) {
    return m.pattern_scope
}
// Serialize serializes information the current object
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecretScanningScan.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("pattern_name", m.GetPatternName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("pattern_scope", m.GetPatternScope())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPatternName sets the pattern_name property value. Name of the custom pattern for custom pattern scans
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) SetPatternName(value *string)() {
    m.pattern_name = value
}
// SetPatternScope sets the pattern_scope property value. Level at which the custom pattern is defined, one of "repository", "organization", or "enterprise"
func (m *SecretScanningScanHistory_custom_pattern_backfill_scans) SetPatternScope(value *string)() {
    m.pattern_scope = value
}
type SecretScanningScanHistory_custom_pattern_backfill_scansable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecretScanningScanable
    GetPatternName()(*string)
    GetPatternScope()(*string)
    SetPatternName(value *string)()
    SetPatternScope(value *string)()
}
