package organizations

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemCustom_rolesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custom_roles property
    custom_roles []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable
    // The number of custom roles in this organization
    total_count *int32
}
// NewItemCustom_rolesGetResponse instantiates a new ItemCustom_rolesGetResponse and sets the default values.
func NewItemCustom_rolesGetResponse()(*ItemCustom_rolesGetResponse) {
    m := &ItemCustom_rolesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCustom_rolesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCustom_rolesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCustom_rolesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCustom_rolesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomRoles gets the custom_roles property value. The custom_roles property
// returns a []OrganizationCustomRepositoryRoleable when successful
func (m *ItemCustom_rolesGetResponse) GetCustomRoles()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable) {
    return m.custom_roles
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCustom_rolesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateOrganizationCustomRepositoryRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable)
                }
            }
            m.SetCustomRoles(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The number of custom roles in this organization
// returns a *int32 when successful
func (m *ItemCustom_rolesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemCustom_rolesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustomRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomRoles()))
        for i, v := range m.GetCustomRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("custom_roles", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemCustom_rolesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomRoles sets the custom_roles property value. The custom_roles property
func (m *ItemCustom_rolesGetResponse) SetCustomRoles(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable)() {
    m.custom_roles = value
}
// SetTotalCount sets the total_count property value. The number of custom roles in this organization
func (m *ItemCustom_rolesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemCustom_rolesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomRoles()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable)
    GetTotalCount()(*int32)
    SetCustomRoles(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.OrganizationCustomRepositoryRoleable)()
    SetTotalCount(value *int32)()
}
