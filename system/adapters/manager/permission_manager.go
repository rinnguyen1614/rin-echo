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
		if !resource.IsEmptyPathOrMethod() {
			policy := util.JoinSlice(utils.ToString(roleID), resource.Path, resource.Method)
			policies = append(policies, policy)
		}
	}

	return m.enforcer.AddPolicies(policies)
}

func (m permissionManager) AddPermissionsForRole(roleID uint, resources domain.Resources) (bool, error) {
	var policies [][]string
	for _, resource := range resources {
		if !resource.IsEmptyPathOrMethod() {
			policy := util.JoinSlice(utils.ToString(roleID), resource.Path, resource.Method)
			policies = append(policies, policy)
		}
	}

	return m.enforcer.AddPolicies(policies)
}

func (m permissionManager) RemovePermissionForRole(roleID uint, resource domain.Resource) (bool, error) {
	return m.RemovePermissionForRoles([]uint{roleID}, resource)
}

func (m permissionManager) RemovePermissionForRoles(roleIDs []uint, resource domain.Resource) (bool, error) {
	var policies [][]string
	for _, roleID := range roleIDs {
		policy := util.JoinSlice(utils.ToString(roleID), resource.Path, resource.Method)
		policies = append(policies, policy)
	}

	return m.enforcer.RemovePolicies(policies)
}

func (m permissionManager) RemovePermissionsForRole(roleID uint, resources domain.Resources) (bool, error) {
	var policies [][]string
	for _, resource := range resources {
		policy := util.JoinSlice(utils.ToString(roleID), resource.Path, resource.Method)
		policies = append(policies, policy)
	}

	return m.enforcer.RemovePolicies(policies)
}

func (m permissionManager) UpdatePermissionForRole(roleID uint, oldResource, newResource domain.Resource) (bool, error) {
	return m.UpdatePermissionForRoles([]uint{roleID}, oldResource, newResource)
}

func (m permissionManager) UpdatePermissionForRoles(roleIDs []uint, oldResource, newResource domain.Resource) (bool, error) {
	if newResource.IsEmptyPathOrMethod() {
		return false, nil
	}

	var oldPolices, newPolicies [][]string
	for _, roleID := range roleIDs {
		oldPolicy := util.JoinSlice(utils.ToString(roleID), oldResource.Path, oldResource.Method)
		newPolicy := util.JoinSlice(utils.ToString(roleID), newResource.Path, newResource.Method)

		oldPolices = append(oldPolices, oldPolicy)
		newPolicies = append(newPolicies, newPolicy)
	}

	return m.enforcer.UpdatePolicies(oldPolices, newPolicies)
}
