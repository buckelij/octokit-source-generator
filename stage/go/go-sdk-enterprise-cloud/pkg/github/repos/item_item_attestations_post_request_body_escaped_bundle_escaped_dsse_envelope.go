package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemItemAttestationsPostRequestBody_bundle_dsseEnvelope instantiates a new ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope and sets the default values.
func NewItemItemAttestationsPostRequestBody_bundle_dsseEnvelope()(*ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope) {
    m := &ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemAttestationsPostRequestBody_bundle_dsseEnvelope(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemAttestationsPostRequestBody_bundle_dsseEnvelope) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemItemAttestationsPostRequestBody_bundle_dsseEnvelopeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
