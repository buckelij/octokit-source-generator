package enterprises

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemAuditLogStreamsItemWithStream_PutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // This setting pauses or resumes a stream.
    enabled *bool
    // The vendor_specific property
    vendor_specific ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable
}
// ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific composed type wrapper for classes i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable
type ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific struct {
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable
    amazonS3AccessKeysConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable
    amazonS3OidcConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable
    azureBlobConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable
    azureHubConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable
    datadogConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable
    googleCloudConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable
    splunkConfig i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable
}
// NewItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific instantiates a new ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific and sets the default values.
func NewItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific()(*ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) {
    m := &ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific{
    }
    return m
}
// CreateItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific()
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
    return result, nil
}
// GetAmazonS3AccessKeysConfig gets the amazonS3AccessKeysConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable
// returns a AmazonS3AccessKeysConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetAmazonS3AccessKeysConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable) {
    return m.amazonS3AccessKeysConfig
}
// GetAmazonS3OidcConfig gets the amazonS3OidcConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable
// returns a AmazonS3OidcConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetAmazonS3OidcConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable) {
    return m.amazonS3OidcConfig
}
// GetAzureBlobConfig gets the azureBlobConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable
// returns a AzureBlobConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetAzureBlobConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable) {
    return m.azureBlobConfig
}
// GetAzureHubConfig gets the azureHubConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable
// returns a AzureHubConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetAzureHubConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable) {
    return m.azureHubConfig
}
// GetDatadogConfig gets the datadogConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable
// returns a DatadogConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetDatadogConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable) {
    return m.datadogConfig
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetAmazonS3AccessKeysConfig() != nil {
        return m.GetAmazonS3AccessKeysConfig().GetFieldDeserializers()
    } else if m.GetAmazonS3OidcConfig() != nil {
        return m.GetAmazonS3OidcConfig().GetFieldDeserializers()
    } else if m.GetAzureBlobConfig() != nil {
        return m.GetAzureBlobConfig().GetFieldDeserializers()
    } else if m.GetAzureHubConfig() != nil {
        return m.GetAzureHubConfig().GetFieldDeserializers()
    } else if m.GetDatadogConfig() != nil {
        return m.GetDatadogConfig().GetFieldDeserializers()
    } else if m.GetGoogleCloudConfig() != nil {
        return m.GetGoogleCloudConfig().GetFieldDeserializers()
    } else if m.GetSplunkConfig() != nil {
        return m.GetSplunkConfig().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetGoogleCloudConfig gets the googleCloudConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable
// returns a GoogleCloudConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetGoogleCloudConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable) {
    return m.googleCloudConfig
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetIsComposedType()(bool) {
    return true
}
// GetSplunkConfig gets the splunkConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable
// returns a SplunkConfigable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) GetSplunkConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable) {
    return m.splunkConfig
}
// Serialize serializes information the current object
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAmazonS3AccessKeysConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAmazonS3AccessKeysConfig())
        if err != nil {
            return err
        }
    } else if m.GetAmazonS3OidcConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAmazonS3OidcConfig())
        if err != nil {
            return err
        }
    } else if m.GetAzureBlobConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAzureBlobConfig())
        if err != nil {
            return err
        }
    } else if m.GetAzureHubConfig() != nil {
        err := writer.WriteObjectValue("", m.GetAzureHubConfig())
        if err != nil {
            return err
        }
    } else if m.GetDatadogConfig() != nil {
        err := writer.WriteObjectValue("", m.GetDatadogConfig())
        if err != nil {
            return err
        }
    } else if m.GetGoogleCloudConfig() != nil {
        err := writer.WriteObjectValue("", m.GetGoogleCloudConfig())
        if err != nil {
            return err
        }
    } else if m.GetSplunkConfig() != nil {
        err := writer.WriteObjectValue("", m.GetSplunkConfig())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmazonS3AccessKeysConfig sets the amazonS3AccessKeysConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetAmazonS3AccessKeysConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable)() {
    m.amazonS3AccessKeysConfig = value
}
// SetAmazonS3OidcConfig sets the amazonS3OidcConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetAmazonS3OidcConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable)() {
    m.amazonS3OidcConfig = value
}
// SetAzureBlobConfig sets the azureBlobConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetAzureBlobConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable)() {
    m.azureBlobConfig = value
}
// SetAzureHubConfig sets the azureHubConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetAzureHubConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable)() {
    m.azureHubConfig = value
}
// SetDatadogConfig sets the datadogConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetDatadogConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable)() {
    m.datadogConfig = value
}
// SetGoogleCloudConfig sets the googleCloudConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetGoogleCloudConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable)() {
    m.googleCloudConfig = value
}
// SetSplunkConfig sets the splunkConfig property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specific) SetSplunkConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable)() {
    m.splunkConfig = value
}
type ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmazonS3AccessKeysConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable)
    GetAmazonS3OidcConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable)
    GetAzureBlobConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable)
    GetAzureHubConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable)
    GetDatadogConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable)
    GetGoogleCloudConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable)
    GetSplunkConfig()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable)
    SetAmazonS3AccessKeysConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3AccessKeysConfigable)()
    SetAmazonS3OidcConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AmazonS3OidcConfigable)()
    SetAzureBlobConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureBlobConfigable)()
    SetAzureHubConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AzureHubConfigable)()
    SetDatadogConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DatadogConfigable)()
    SetGoogleCloudConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.GoogleCloudConfigable)()
    SetSplunkConfig(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SplunkConfigable)()
}
// NewItemAuditLogStreamsItemWithStream_PutRequestBody instantiates a new ItemAuditLogStreamsItemWithStream_PutRequestBody and sets the default values.
func NewItemAuditLogStreamsItemWithStream_PutRequestBody()(*ItemAuditLogStreamsItemWithStream_PutRequestBody) {
    m := &ItemAuditLogStreamsItemWithStream_PutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAuditLogStreamsItemWithStream_PutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAuditLogStreamsItemWithStream_PutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAuditLogStreamsItemWithStream_PutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. This setting pauses or resumes a stream.
// returns a *bool when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["vendor_specific"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorSpecific(val.(ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable))
        }
        return nil
    }
    return res
}
// GetVendorSpecific gets the vendor_specific property value. The vendor_specific property
// returns a ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable when successful
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) GetVendorSpecific()(ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable) {
    return m.vendor_specific
}
// Serialize serializes information the current object
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("vendor_specific", m.GetVendorSpecific())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. This setting pauses or resumes a stream.
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetVendorSpecific sets the vendor_specific property value. The vendor_specific property
func (m *ItemAuditLogStreamsItemWithStream_PutRequestBody) SetVendorSpecific(value ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable)() {
    m.vendor_specific = value
}
type ItemAuditLogStreamsItemWithStream_PutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetVendorSpecific()(ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable)
    SetEnabled(value *bool)()
    SetVendorSpecific(value ItemAuditLogStreamsItemWithStream_PutRequestBody_WithStream_PutRequestBody_vendor_specificable)()
}
