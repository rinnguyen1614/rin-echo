package manager

import (
	"rin-echo/common/utils"
	"rin-echo/system/domain"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

type permissionManager struct {
	enforcer *casbin.SyncedEnforcer
}

func NewPermissionManager(e *casbin.SyncedEnforcer) domain.PermissionManager {
	return &permissionManager{
		enforcer: e,
	}
}

func (m permissionManager) HasPermissionForRole(roleID uint, resource domain.Resource) bool {
	return m.enforcer.HasPermissionForUser(utils.ToString(roleID), resource.Object, resource.Action)
}

func (m permissionManager) AddRole(roleID uint) (bool, error) {
	return m.enforcer.AddNamedGroupingPolicy("g", utils.ToString(roleID))
}

func (re permissionManager) AddRoleForUser(userID, roleID uint) (bool, error) {
	return re.enforcer.AddRoleForUser(utils.ToString(userID), utils.ToString(roleID))
}

func (re permissionManager) AddRolesForUser(userID uint, roleIDs []uint) (bool, error) {
	var roles []string
	for _, roleID := range roleIDs {
		roles = append(roles, utils.ToString(roleID))
	}
	return re.enforcer.AddRolesForUser(utils.ToString(userID), roles)
}

func (re permissionManager) DeleteRoleForUser(userID, roleID uint) (bool, error) {
	return re.enforcer.DeleteRoleForUser(utils.ToString(userID), utils.ToString(roleID))
}

func (re permissionManager) DeleteRolesForUser(userID uint, roleIDs []uint) (bool, error) {
	for _, roleID := range roleIDs {
		if ok, err := re.DeleteRoleForUser(userID, roleID); err != nil {
			return ok, err
		}
	}
	return true, nil
}

func (m permissionManager) AddPermissionForRole(roleID uint, resource domain.Resource) (bool, error) {
	return m.AddPermissionForRoles([]uint{roleID}, resource)
}

func (m permissionManager) AddPermissionForRoles(roleIDs []uint, resource domain.Resource) (bool, error) {
	var policies [][]string
	for _, roleID := range roleIDs {
		if !resource.IsEmptyObjectOrAction() {
			policy := util.JoinSlice(utils.ToString(roleID), resource.Object, resource.Action)
			policies = append(policies, policy)
		}
	}

	if len(policies) > 0 {
		m.enforcer.AddPolicies(policies)
	}
	return true, nil
}

func (m permissionManager) AddPermissionsForRole(roleID uint, resources domain.Resources) (bool, error) {
	var policies [][]string
	for _, resource := range resources {
		if !resource.IsEmptyObjectOrAction() {
			policy := util.JoinSlice(utils.ToString(roleID), resource.Object, resource.Action)
			policies = append(policies, policy)
		}
	}

	if len(policies) > 0 {
		m.enforcer.AddPolicies(policies)
	}
	return true, nil
}

func (m permissionManager) RemovePermissionForRole(roleID uint, resource domain.Resource) (bool, error) {
	return m.RemovePermissionForRoles([]uint{roleID}, resource)
}

func (m permissionManager) RemovePermissionForRoles(roleIDs []uint, resource domain.Resource) (bool, error) {
	if resource.IsEmptyObjectOrAction() {
		return true, nil
	}

	var policies [][]string
	for _, roleID := range roleIDs {
		policy := util.JoinSlice(utils.ToString(roleID), resource.Object, resource.Action)
		policies = append(policies, policy)
	}

	return m.enforcer.RemovePolicies(policies)
}

func (m permissionManager) RemovePermissionsForRole(roleID uint, resources domain.Resources) (bool, error) {
	var policies [][]string
	for _, resource := range resources {
		if !resource.IsEmptyObjectOrAction() {
			policy := util.JoinSlice(utils.ToString(roleID), resource.Object, resource.Action)
			policies = append(policies, policy)
		}
	}

	return m.enforcer.RemovePolicies(policies)
}

func (m permissionManager) UpdatePermissionForRole(roleID uint, oldResource, newResource domain.Resource) (bool, error) {
	return m.UpdatePermissionForRoles([]uint{roleID}, oldResource, newResource)
}

func (m permissionManager) UpdatePermissionForRoles(roleIDs []uint, oldResource, newResource domain.Resource) (bool, error) {
	if oldResource.IsEmptyObjectOrAction() {
		return m.AddPermissionForRoles(roleIDs, newResource)
	}

	if newResource.IsEmptyObjectOrAction() {
		return m.RemovePermissionForRoles(roleIDs, oldResource)
	}

	var oldPolices, newPolicies [][]string
	for _, roleID := range roleIDs {
		oldPolicy := util.JoinSlice(utils.ToString(roleID), oldResource.Object, oldResource.Action)
		newPolicy := util.JoinSlice(utils.ToString(roleID), newResource.Object, newResource.Action)

		oldPolices = append(oldPolices, oldPolicy)
		newPolicies = append(newPolicies, newPolicy)
	}

	return m.enforcer.UpdatePolicies(oldPolices, newPolicies)
}
