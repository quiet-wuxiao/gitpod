package wsmanagerbridge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gitpod-io/gitpod/installer/pkg/common"
)

// WorkspaceCluster from components/gitpod-protocol/src/workspace-cluster.ts
type WorkspaceCluster struct {
	Name                 string                `json:"name"`
	URL                  string                `json:"url"`
	TLS                  common.TLS            `json:"tls"`
	State                WorkspaceClusterState `json:"state"`
	MaxScore             int32                 `json:"maxScore"`
	Score                int32                 `json:"score"`
	Govern               bool                  `json:"govern"`
	AdmissionConstraints []AdmissionConstraint `json:"admissionConstraints"`
}

// WorkspaceClusterState from components/gitpod-protocol/src/workspace-cluster.ts
type WorkspaceClusterState string

const (
	WorkspaceClusterStateAvailable WorkspaceClusterState = "available"
	WorkspaceClusterStateCordoned  WorkspaceClusterState = "cordoned"
	WorkspaceClusterStateDraining  WorkspaceClusterState = "draining"
)

type AdmissionConstraint struct {
	Type       AdmissionConstraintType       `json:"type"`
	Permission AdmissionConstraintPermission `json:"permission"`
}

type AdmissionConstraintType string

const (
	AdmissionConstraintFeaturePreview AdmissionConstraintType = "has-feature-preview"
	AdmissionConstraintHasRole        AdmissionConstraintType = "has-permission"
)

type AdmissionConstraintPermission string

const (
	AdmissionConstraintPermissionMonitor             AdmissionConstraintPermission = "monitor"
	AdmissionConstraintPermissionEnforcement         AdmissionConstraintPermission = "enforcement"
	AdmissionConstraintPermissionPrivilegedWS        AdmissionConstraintPermission = "privileged-ws"
	AdmissionConstraintPermissionRegistryAccess      AdmissionConstraintPermission = "registry-access"
	AdmissionConstraintPermissionAdminUsers          AdmissionConstraintPermission = "admin-users"
	AdmissionConstraintPermissionAdminWorkspaces     AdmissionConstraintPermission = "admin-workspaces"
	AdmissionConstraintPermissionAdminApi            AdmissionConstraintPermission = "admin-api"
	AdmissionConstraintPermissionIDESettings         AdmissionConstraintPermission = "ide-settings"
	AdmissionConstraintPermissionNewWorkspaceCluster AdmissionConstraintPermission = "new-workspace-cluster"
	AdmissionConstraintPermissionTeamsAndProjects    AdmissionConstraintPermission = "teams-and-projects"
)

func GenerateWorkspaceManagerListForEnvVar(cluster *WorkspaceCluster) (*string, error) {
	fc, err := json.MarshalIndent(cluster, "", " ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal WorkspaceManagerList config: %w", err)
	}

	str := base64.StdEncoding.EncodeToString(fc)

	return &str, nil
}
