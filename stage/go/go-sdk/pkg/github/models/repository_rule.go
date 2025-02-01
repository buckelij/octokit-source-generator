package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRule composed type wrapper for classes RepositoryRuleBranchNamePatternable, RepositoryRuleCodeScanningable, RepositoryRuleCommitAuthorEmailPatternable, RepositoryRuleCommitMessagePatternable, RepositoryRuleCommitterEmailPatternable, RepositoryRuleCreationable, RepositoryRuleDeletionable, RepositoryRuleMember1able, RepositoryRuleMember2able, RepositoryRuleMember3able, RepositoryRuleMember4able, RepositoryRuleMergeQueueable, RepositoryRuleNonFastForwardable, RepositoryRulePullRequestable, RepositoryRuleRequiredDeploymentsable, RepositoryRuleRequiredLinearHistoryable, RepositoryRuleRequiredSignaturesable, RepositoryRuleRequiredStatusChecksable, RepositoryRuleTagNamePatternable, RepositoryRuleUpdateable, RepositoryRuleWorkflowsable
type RepositoryRule struct {
    // Composed type representation for type RepositoryRuleBranchNamePatternable
    repositoryRuleBranchNamePattern RepositoryRuleBranchNamePatternable
    // Composed type representation for type RepositoryRuleCodeScanningable
    repositoryRuleCodeScanning RepositoryRuleCodeScanningable
    // Composed type representation for type RepositoryRuleCommitAuthorEmailPatternable
    repositoryRuleCommitAuthorEmailPattern RepositoryRuleCommitAuthorEmailPatternable
    // Composed type representation for type RepositoryRuleCommitMessagePatternable
    repositoryRuleCommitMessagePattern RepositoryRuleCommitMessagePatternable
    // Composed type representation for type RepositoryRuleCommitterEmailPatternable
    repositoryRuleCommitterEmailPattern RepositoryRuleCommitterEmailPatternable
    // Composed type representation for type RepositoryRuleCreationable
    repositoryRuleCreation RepositoryRuleCreationable
    // Composed type representation for type RepositoryRuleDeletionable
    repositoryRuleDeletion RepositoryRuleDeletionable
    // Composed type representation for type RepositoryRuleMember1able
    repositoryRuleMember1 RepositoryRuleMember1able
    // Composed type representation for type RepositoryRuleMember2able
    repositoryRuleMember2 RepositoryRuleMember2able
    // Composed type representation for type RepositoryRuleMember3able
    repositoryRuleMember3 RepositoryRuleMember3able
    // Composed type representation for type RepositoryRuleMember4able
    repositoryRuleMember4 RepositoryRuleMember4able
    // Composed type representation for type RepositoryRuleMergeQueueable
    repositoryRuleMergeQueue RepositoryRuleMergeQueueable
    // Composed type representation for type RepositoryRuleNonFastForwardable
    repositoryRuleNonFastForward RepositoryRuleNonFastForwardable
    // Composed type representation for type RepositoryRulePullRequestable
    repositoryRulePullRequest RepositoryRulePullRequestable
    // Composed type representation for type RepositoryRuleRequiredDeploymentsable
    repositoryRuleRequiredDeployments RepositoryRuleRequiredDeploymentsable
    // Composed type representation for type RepositoryRuleRequiredLinearHistoryable
    repositoryRuleRequiredLinearHistory RepositoryRuleRequiredLinearHistoryable
    // Composed type representation for type RepositoryRuleRequiredSignaturesable
    repositoryRuleRequiredSignatures RepositoryRuleRequiredSignaturesable
    // Composed type representation for type RepositoryRuleRequiredStatusChecksable
    repositoryRuleRequiredStatusChecks RepositoryRuleRequiredStatusChecksable
    // Composed type representation for type RepositoryRuleTagNamePatternable
    repositoryRuleTagNamePattern RepositoryRuleTagNamePatternable
    // Composed type representation for type RepositoryRuleUpdateable
    repositoryRuleUpdate RepositoryRuleUpdateable
    // Composed type representation for type RepositoryRuleWorkflowsable
    repositoryRuleWorkflows RepositoryRuleWorkflowsable
}
// NewRepositoryRule instantiates a new RepositoryRule and sets the default values.
func NewRepositoryRule()(*RepositoryRule) {
    m := &RepositoryRule{
    }
    return m
}
// CreateRepositoryRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewRepositoryRule()
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
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetRepositoryRuleBranchNamePattern() != nil {
        return m.GetRepositoryRuleBranchNamePattern().GetFieldDeserializers()
    } else if m.GetRepositoryRuleCodeScanning() != nil {
        return m.GetRepositoryRuleCodeScanning().GetFieldDeserializers()
    } else if m.GetRepositoryRuleCommitAuthorEmailPattern() != nil {
        return m.GetRepositoryRuleCommitAuthorEmailPattern().GetFieldDeserializers()
    } else if m.GetRepositoryRuleCommitMessagePattern() != nil {
        return m.GetRepositoryRuleCommitMessagePattern().GetFieldDeserializers()
    } else if m.GetRepositoryRuleCommitterEmailPattern() != nil {
        return m.GetRepositoryRuleCommitterEmailPattern().GetFieldDeserializers()
    } else if m.GetRepositoryRuleCreation() != nil {
        return m.GetRepositoryRuleCreation().GetFieldDeserializers()
    } else if m.GetRepositoryRuleDeletion() != nil {
        return m.GetRepositoryRuleDeletion().GetFieldDeserializers()
    } else if m.GetRepositoryRuleMember1() != nil {
        return m.GetRepositoryRuleMember1().GetFieldDeserializers()
    } else if m.GetRepositoryRuleMember2() != nil {
        return m.GetRepositoryRuleMember2().GetFieldDeserializers()
    } else if m.GetRepositoryRuleMember3() != nil {
        return m.GetRepositoryRuleMember3().GetFieldDeserializers()
    } else if m.GetRepositoryRuleMember4() != nil {
        return m.GetRepositoryRuleMember4().GetFieldDeserializers()
    } else if m.GetRepositoryRuleMergeQueue() != nil {
        return m.GetRepositoryRuleMergeQueue().GetFieldDeserializers()
    } else if m.GetRepositoryRuleNonFastForward() != nil {
        return m.GetRepositoryRuleNonFastForward().GetFieldDeserializers()
    } else if m.GetRepositoryRulePullRequest() != nil {
        return m.GetRepositoryRulePullRequest().GetFieldDeserializers()
    } else if m.GetRepositoryRuleRequiredDeployments() != nil {
        return m.GetRepositoryRuleRequiredDeployments().GetFieldDeserializers()
    } else if m.GetRepositoryRuleRequiredLinearHistory() != nil {
        return m.GetRepositoryRuleRequiredLinearHistory().GetFieldDeserializers()
    } else if m.GetRepositoryRuleRequiredSignatures() != nil {
        return m.GetRepositoryRuleRequiredSignatures().GetFieldDeserializers()
    } else if m.GetRepositoryRuleRequiredStatusChecks() != nil {
        return m.GetRepositoryRuleRequiredStatusChecks().GetFieldDeserializers()
    } else if m.GetRepositoryRuleTagNamePattern() != nil {
        return m.GetRepositoryRuleTagNamePattern().GetFieldDeserializers()
    } else if m.GetRepositoryRuleUpdate() != nil {
        return m.GetRepositoryRuleUpdate().GetFieldDeserializers()
    } else if m.GetRepositoryRuleWorkflows() != nil {
        return m.GetRepositoryRuleWorkflows().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *RepositoryRule) GetIsComposedType()(bool) {
    return true
}
// GetRepositoryRuleBranchNamePattern gets the repositoryRuleBranchNamePattern property value. Composed type representation for type RepositoryRuleBranchNamePatternable
// returns a RepositoryRuleBranchNamePatternable when successful
func (m *RepositoryRule) GetRepositoryRuleBranchNamePattern()(RepositoryRuleBranchNamePatternable) {
    return m.repositoryRuleBranchNamePattern
}
// GetRepositoryRuleCodeScanning gets the repositoryRuleCodeScanning property value. Composed type representation for type RepositoryRuleCodeScanningable
// returns a RepositoryRuleCodeScanningable when successful
func (m *RepositoryRule) GetRepositoryRuleCodeScanning()(RepositoryRuleCodeScanningable) {
    return m.repositoryRuleCodeScanning
}
// GetRepositoryRuleCommitAuthorEmailPattern gets the repositoryRuleCommitAuthorEmailPattern property value. Composed type representation for type RepositoryRuleCommitAuthorEmailPatternable
// returns a RepositoryRuleCommitAuthorEmailPatternable when successful
func (m *RepositoryRule) GetRepositoryRuleCommitAuthorEmailPattern()(RepositoryRuleCommitAuthorEmailPatternable) {
    return m.repositoryRuleCommitAuthorEmailPattern
}
// GetRepositoryRuleCommitMessagePattern gets the repositoryRuleCommitMessagePattern property value. Composed type representation for type RepositoryRuleCommitMessagePatternable
// returns a RepositoryRuleCommitMessagePatternable when successful
func (m *RepositoryRule) GetRepositoryRuleCommitMessagePattern()(RepositoryRuleCommitMessagePatternable) {
    return m.repositoryRuleCommitMessagePattern
}
// GetRepositoryRuleCommitterEmailPattern gets the repositoryRuleCommitterEmailPattern property value. Composed type representation for type RepositoryRuleCommitterEmailPatternable
// returns a RepositoryRuleCommitterEmailPatternable when successful
func (m *RepositoryRule) GetRepositoryRuleCommitterEmailPattern()(RepositoryRuleCommitterEmailPatternable) {
    return m.repositoryRuleCommitterEmailPattern
}
// GetRepositoryRuleCreation gets the repositoryRuleCreation property value. Composed type representation for type RepositoryRuleCreationable
// returns a RepositoryRuleCreationable when successful
func (m *RepositoryRule) GetRepositoryRuleCreation()(RepositoryRuleCreationable) {
    return m.repositoryRuleCreation
}
// GetRepositoryRuleDeletion gets the repositoryRuleDeletion property value. Composed type representation for type RepositoryRuleDeletionable
// returns a RepositoryRuleDeletionable when successful
func (m *RepositoryRule) GetRepositoryRuleDeletion()(RepositoryRuleDeletionable) {
    return m.repositoryRuleDeletion
}
// GetRepositoryRuleMember1 gets the repositoryRuleMember1 property value. Composed type representation for type RepositoryRuleMember1able
// returns a RepositoryRuleMember1able when successful
func (m *RepositoryRule) GetRepositoryRuleMember1()(RepositoryRuleMember1able) {
    return m.repositoryRuleMember1
}
// GetRepositoryRuleMember2 gets the repositoryRuleMember2 property value. Composed type representation for type RepositoryRuleMember2able
// returns a RepositoryRuleMember2able when successful
func (m *RepositoryRule) GetRepositoryRuleMember2()(RepositoryRuleMember2able) {
    return m.repositoryRuleMember2
}
// GetRepositoryRuleMember3 gets the repositoryRuleMember3 property value. Composed type representation for type RepositoryRuleMember3able
// returns a RepositoryRuleMember3able when successful
func (m *RepositoryRule) GetRepositoryRuleMember3()(RepositoryRuleMember3able) {
    return m.repositoryRuleMember3
}
// GetRepositoryRuleMember4 gets the repositoryRuleMember4 property value. Composed type representation for type RepositoryRuleMember4able
// returns a RepositoryRuleMember4able when successful
func (m *RepositoryRule) GetRepositoryRuleMember4()(RepositoryRuleMember4able) {
    return m.repositoryRuleMember4
}
// GetRepositoryRuleMergeQueue gets the repositoryRuleMergeQueue property value. Composed type representation for type RepositoryRuleMergeQueueable
// returns a RepositoryRuleMergeQueueable when successful
func (m *RepositoryRule) GetRepositoryRuleMergeQueue()(RepositoryRuleMergeQueueable) {
    return m.repositoryRuleMergeQueue
}
// GetRepositoryRuleNonFastForward gets the repositoryRuleNonFastForward property value. Composed type representation for type RepositoryRuleNonFastForwardable
// returns a RepositoryRuleNonFastForwardable when successful
func (m *RepositoryRule) GetRepositoryRuleNonFastForward()(RepositoryRuleNonFastForwardable) {
    return m.repositoryRuleNonFastForward
}
// GetRepositoryRulePullRequest gets the repositoryRulePullRequest property value. Composed type representation for type RepositoryRulePullRequestable
// returns a RepositoryRulePullRequestable when successful
func (m *RepositoryRule) GetRepositoryRulePullRequest()(RepositoryRulePullRequestable) {
    return m.repositoryRulePullRequest
}
// GetRepositoryRuleRequiredDeployments gets the repositoryRuleRequiredDeployments property value. Composed type representation for type RepositoryRuleRequiredDeploymentsable
// returns a RepositoryRuleRequiredDeploymentsable when successful
func (m *RepositoryRule) GetRepositoryRuleRequiredDeployments()(RepositoryRuleRequiredDeploymentsable) {
    return m.repositoryRuleRequiredDeployments
}
// GetRepositoryRuleRequiredLinearHistory gets the repositoryRuleRequiredLinearHistory property value. Composed type representation for type RepositoryRuleRequiredLinearHistoryable
// returns a RepositoryRuleRequiredLinearHistoryable when successful
func (m *RepositoryRule) GetRepositoryRuleRequiredLinearHistory()(RepositoryRuleRequiredLinearHistoryable) {
    return m.repositoryRuleRequiredLinearHistory
}
// GetRepositoryRuleRequiredSignatures gets the repositoryRuleRequiredSignatures property value. Composed type representation for type RepositoryRuleRequiredSignaturesable
// returns a RepositoryRuleRequiredSignaturesable when successful
func (m *RepositoryRule) GetRepositoryRuleRequiredSignatures()(RepositoryRuleRequiredSignaturesable) {
    return m.repositoryRuleRequiredSignatures
}
// GetRepositoryRuleRequiredStatusChecks gets the repositoryRuleRequiredStatusChecks property value. Composed type representation for type RepositoryRuleRequiredStatusChecksable
// returns a RepositoryRuleRequiredStatusChecksable when successful
func (m *RepositoryRule) GetRepositoryRuleRequiredStatusChecks()(RepositoryRuleRequiredStatusChecksable) {
    return m.repositoryRuleRequiredStatusChecks
}
// GetRepositoryRuleTagNamePattern gets the repositoryRuleTagNamePattern property value. Composed type representation for type RepositoryRuleTagNamePatternable
// returns a RepositoryRuleTagNamePatternable when successful
func (m *RepositoryRule) GetRepositoryRuleTagNamePattern()(RepositoryRuleTagNamePatternable) {
    return m.repositoryRuleTagNamePattern
}
// GetRepositoryRuleUpdate gets the repositoryRuleUpdate property value. Composed type representation for type RepositoryRuleUpdateable
// returns a RepositoryRuleUpdateable when successful
func (m *RepositoryRule) GetRepositoryRuleUpdate()(RepositoryRuleUpdateable) {
    return m.repositoryRuleUpdate
}
// GetRepositoryRuleWorkflows gets the repositoryRuleWorkflows property value. Composed type representation for type RepositoryRuleWorkflowsable
// returns a RepositoryRuleWorkflowsable when successful
func (m *RepositoryRule) GetRepositoryRuleWorkflows()(RepositoryRuleWorkflowsable) {
    return m.repositoryRuleWorkflows
}
// Serialize serializes information the current object
func (m *RepositoryRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositoryRuleBranchNamePattern() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleBranchNamePattern())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleCodeScanning() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleCodeScanning())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleCommitAuthorEmailPattern() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleCommitAuthorEmailPattern())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleCommitMessagePattern() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleCommitMessagePattern())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleCommitterEmailPattern() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleCommitterEmailPattern())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleCreation() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleCreation())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleDeletion() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleDeletion())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleMember1() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleMember1())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleMember2() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleMember2())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleMember3() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleMember3())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleMember4() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleMember4())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleMergeQueue() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleMergeQueue())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleNonFastForward() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleNonFastForward())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRulePullRequest() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRulePullRequest())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleRequiredDeployments() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleRequiredDeployments())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleRequiredLinearHistory() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleRequiredLinearHistory())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleRequiredSignatures() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleRequiredSignatures())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleRequiredStatusChecks() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleRequiredStatusChecks())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleTagNamePattern() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleTagNamePattern())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleUpdate() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleUpdate())
        if err != nil {
            return err
        }
    } else if m.GetRepositoryRuleWorkflows() != nil {
        err := writer.WriteObjectValue("", m.GetRepositoryRuleWorkflows())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRepositoryRuleBranchNamePattern sets the repositoryRuleBranchNamePattern property value. Composed type representation for type RepositoryRuleBranchNamePatternable
func (m *RepositoryRule) SetRepositoryRuleBranchNamePattern(value RepositoryRuleBranchNamePatternable)() {
    m.repositoryRuleBranchNamePattern = value
}
// SetRepositoryRuleCodeScanning sets the repositoryRuleCodeScanning property value. Composed type representation for type RepositoryRuleCodeScanningable
func (m *RepositoryRule) SetRepositoryRuleCodeScanning(value RepositoryRuleCodeScanningable)() {
    m.repositoryRuleCodeScanning = value
}
// SetRepositoryRuleCommitAuthorEmailPattern sets the repositoryRuleCommitAuthorEmailPattern property value. Composed type representation for type RepositoryRuleCommitAuthorEmailPatternable
func (m *RepositoryRule) SetRepositoryRuleCommitAuthorEmailPattern(value RepositoryRuleCommitAuthorEmailPatternable)() {
    m.repositoryRuleCommitAuthorEmailPattern = value
}
// SetRepositoryRuleCommitMessagePattern sets the repositoryRuleCommitMessagePattern property value. Composed type representation for type RepositoryRuleCommitMessagePatternable
func (m *RepositoryRule) SetRepositoryRuleCommitMessagePattern(value RepositoryRuleCommitMessagePatternable)() {
    m.repositoryRuleCommitMessagePattern = value
}
// SetRepositoryRuleCommitterEmailPattern sets the repositoryRuleCommitterEmailPattern property value. Composed type representation for type RepositoryRuleCommitterEmailPatternable
func (m *RepositoryRule) SetRepositoryRuleCommitterEmailPattern(value RepositoryRuleCommitterEmailPatternable)() {
    m.repositoryRuleCommitterEmailPattern = value
}
// SetRepositoryRuleCreation sets the repositoryRuleCreation property value. Composed type representation for type RepositoryRuleCreationable
func (m *RepositoryRule) SetRepositoryRuleCreation(value RepositoryRuleCreationable)() {
    m.repositoryRuleCreation = value
}
// SetRepositoryRuleDeletion sets the repositoryRuleDeletion property value. Composed type representation for type RepositoryRuleDeletionable
func (m *RepositoryRule) SetRepositoryRuleDeletion(value RepositoryRuleDeletionable)() {
    m.repositoryRuleDeletion = value
}
// SetRepositoryRuleMember1 sets the repositoryRuleMember1 property value. Composed type representation for type RepositoryRuleMember1able
func (m *RepositoryRule) SetRepositoryRuleMember1(value RepositoryRuleMember1able)() {
    m.repositoryRuleMember1 = value
}
// SetRepositoryRuleMember2 sets the repositoryRuleMember2 property value. Composed type representation for type RepositoryRuleMember2able
func (m *RepositoryRule) SetRepositoryRuleMember2(value RepositoryRuleMember2able)() {
    m.repositoryRuleMember2 = value
}
// SetRepositoryRuleMember3 sets the repositoryRuleMember3 property value. Composed type representation for type RepositoryRuleMember3able
func (m *RepositoryRule) SetRepositoryRuleMember3(value RepositoryRuleMember3able)() {
    m.repositoryRuleMember3 = value
}
// SetRepositoryRuleMember4 sets the repositoryRuleMember4 property value. Composed type representation for type RepositoryRuleMember4able
func (m *RepositoryRule) SetRepositoryRuleMember4(value RepositoryRuleMember4able)() {
    m.repositoryRuleMember4 = value
}
// SetRepositoryRuleMergeQueue sets the repositoryRuleMergeQueue property value. Composed type representation for type RepositoryRuleMergeQueueable
func (m *RepositoryRule) SetRepositoryRuleMergeQueue(value RepositoryRuleMergeQueueable)() {
    m.repositoryRuleMergeQueue = value
}
// SetRepositoryRuleNonFastForward sets the repositoryRuleNonFastForward property value. Composed type representation for type RepositoryRuleNonFastForwardable
func (m *RepositoryRule) SetRepositoryRuleNonFastForward(value RepositoryRuleNonFastForwardable)() {
    m.repositoryRuleNonFastForward = value
}
// SetRepositoryRulePullRequest sets the repositoryRulePullRequest property value. Composed type representation for type RepositoryRulePullRequestable
func (m *RepositoryRule) SetRepositoryRulePullRequest(value RepositoryRulePullRequestable)() {
    m.repositoryRulePullRequest = value
}
// SetRepositoryRuleRequiredDeployments sets the repositoryRuleRequiredDeployments property value. Composed type representation for type RepositoryRuleRequiredDeploymentsable
func (m *RepositoryRule) SetRepositoryRuleRequiredDeployments(value RepositoryRuleRequiredDeploymentsable)() {
    m.repositoryRuleRequiredDeployments = value
}
// SetRepositoryRuleRequiredLinearHistory sets the repositoryRuleRequiredLinearHistory property value. Composed type representation for type RepositoryRuleRequiredLinearHistoryable
func (m *RepositoryRule) SetRepositoryRuleRequiredLinearHistory(value RepositoryRuleRequiredLinearHistoryable)() {
    m.repositoryRuleRequiredLinearHistory = value
}
// SetRepositoryRuleRequiredSignatures sets the repositoryRuleRequiredSignatures property value. Composed type representation for type RepositoryRuleRequiredSignaturesable
func (m *RepositoryRule) SetRepositoryRuleRequiredSignatures(value RepositoryRuleRequiredSignaturesable)() {
    m.repositoryRuleRequiredSignatures = value
}
// SetRepositoryRuleRequiredStatusChecks sets the repositoryRuleRequiredStatusChecks property value. Composed type representation for type RepositoryRuleRequiredStatusChecksable
func (m *RepositoryRule) SetRepositoryRuleRequiredStatusChecks(value RepositoryRuleRequiredStatusChecksable)() {
    m.repositoryRuleRequiredStatusChecks = value
}
// SetRepositoryRuleTagNamePattern sets the repositoryRuleTagNamePattern property value. Composed type representation for type RepositoryRuleTagNamePatternable
func (m *RepositoryRule) SetRepositoryRuleTagNamePattern(value RepositoryRuleTagNamePatternable)() {
    m.repositoryRuleTagNamePattern = value
}
// SetRepositoryRuleUpdate sets the repositoryRuleUpdate property value. Composed type representation for type RepositoryRuleUpdateable
func (m *RepositoryRule) SetRepositoryRuleUpdate(value RepositoryRuleUpdateable)() {
    m.repositoryRuleUpdate = value
}
// SetRepositoryRuleWorkflows sets the repositoryRuleWorkflows property value. Composed type representation for type RepositoryRuleWorkflowsable
func (m *RepositoryRule) SetRepositoryRuleWorkflows(value RepositoryRuleWorkflowsable)() {
    m.repositoryRuleWorkflows = value
}
type RepositoryRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositoryRuleBranchNamePattern()(RepositoryRuleBranchNamePatternable)
    GetRepositoryRuleCodeScanning()(RepositoryRuleCodeScanningable)
    GetRepositoryRuleCommitAuthorEmailPattern()(RepositoryRuleCommitAuthorEmailPatternable)
    GetRepositoryRuleCommitMessagePattern()(RepositoryRuleCommitMessagePatternable)
    GetRepositoryRuleCommitterEmailPattern()(RepositoryRuleCommitterEmailPatternable)
    GetRepositoryRuleCreation()(RepositoryRuleCreationable)
    GetRepositoryRuleDeletion()(RepositoryRuleDeletionable)
    GetRepositoryRuleMember1()(RepositoryRuleMember1able)
    GetRepositoryRuleMember2()(RepositoryRuleMember2able)
    GetRepositoryRuleMember3()(RepositoryRuleMember3able)
    GetRepositoryRuleMember4()(RepositoryRuleMember4able)
    GetRepositoryRuleMergeQueue()(RepositoryRuleMergeQueueable)
    GetRepositoryRuleNonFastForward()(RepositoryRuleNonFastForwardable)
    GetRepositoryRulePullRequest()(RepositoryRulePullRequestable)
    GetRepositoryRuleRequiredDeployments()(RepositoryRuleRequiredDeploymentsable)
    GetRepositoryRuleRequiredLinearHistory()(RepositoryRuleRequiredLinearHistoryable)
    GetRepositoryRuleRequiredSignatures()(RepositoryRuleRequiredSignaturesable)
    GetRepositoryRuleRequiredStatusChecks()(RepositoryRuleRequiredStatusChecksable)
    GetRepositoryRuleTagNamePattern()(RepositoryRuleTagNamePatternable)
    GetRepositoryRuleUpdate()(RepositoryRuleUpdateable)
    GetRepositoryRuleWorkflows()(RepositoryRuleWorkflowsable)
    SetRepositoryRuleBranchNamePattern(value RepositoryRuleBranchNamePatternable)()
    SetRepositoryRuleCodeScanning(value RepositoryRuleCodeScanningable)()
    SetRepositoryRuleCommitAuthorEmailPattern(value RepositoryRuleCommitAuthorEmailPatternable)()
    SetRepositoryRuleCommitMessagePattern(value RepositoryRuleCommitMessagePatternable)()
    SetRepositoryRuleCommitterEmailPattern(value RepositoryRuleCommitterEmailPatternable)()
    SetRepositoryRuleCreation(value RepositoryRuleCreationable)()
    SetRepositoryRuleDeletion(value RepositoryRuleDeletionable)()
    SetRepositoryRuleMember1(value RepositoryRuleMember1able)()
    SetRepositoryRuleMember2(value RepositoryRuleMember2able)()
    SetRepositoryRuleMember3(value RepositoryRuleMember3able)()
    SetRepositoryRuleMember4(value RepositoryRuleMember4able)()
    SetRepositoryRuleMergeQueue(value RepositoryRuleMergeQueueable)()
    SetRepositoryRuleNonFastForward(value RepositoryRuleNonFastForwardable)()
    SetRepositoryRulePullRequest(value RepositoryRulePullRequestable)()
    SetRepositoryRuleRequiredDeployments(value RepositoryRuleRequiredDeploymentsable)()
    SetRepositoryRuleRequiredLinearHistory(value RepositoryRuleRequiredLinearHistoryable)()
    SetRepositoryRuleRequiredSignatures(value RepositoryRuleRequiredSignaturesable)()
    SetRepositoryRuleRequiredStatusChecks(value RepositoryRuleRequiredStatusChecksable)()
    SetRepositoryRuleTagNamePattern(value RepositoryRuleTagNamePatternable)()
    SetRepositoryRuleUpdate(value RepositoryRuleUpdateable)()
    SetRepositoryRuleWorkflows(value RepositoryRuleWorkflowsable)()
}
