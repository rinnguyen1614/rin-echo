package adapters

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

type RBACCasbin struct {
	enforcer *casbin.SyncedEnforcer
}

func NewRBACCasbin(e *casbin.SyncedEnforcer) RBACCasbin {
	return RBACCasbin{
		enforcer: e,
	}
}

func (re RBACCasbin) AddRole(role string) (bool, error) {
	return re.enforcer.AddNamedGroupingPolicy("g", role)
}

func (re RBACCasbin) AddRoleChildrenInParent(parent string, roles ...string) (bool, error) {
	var rules [][]string
	for _, role := range roles {
		rules = append(rules, []string{parent, role})
	}
	return re.enforcer.AddNamedGroupingPolicies("g", rules)
}

func (re RBACCasbin) AddChildreInParentResource(parent string, resources ...string) (bool, error) {
	var rules [][]string
	for _, rs := range resources {
		rules = append(rules, []string{parent, rs})
	}
	return re.enforcer.AddNamedGroupingPolicies("g2", rules)
}

func (re RBACCasbin) AddPermissionForRole(name string, permission ...string) (bool, error) {
	return re.enforcer.AddPolicy(util.JoinSlice(name, permission...))
}

func (re RBACCasbin) GetAllRoles() []string {
	return re.enforcer.GetAllRoles()
}

func (re RBACCasbin) GetRolesForUser(name string) ([]string, error) {
	return re.enforcer.GetRolesForUser(name)
}

func (re RBACCasbin) GetUsersForRole(name string) ([]string, error) {
	return re.enforcer.GetUsersForRole(name)
}

func (re RBACCasbin) HasRoleForUser(name string, role string) (bool, error) {
	return re.enforcer.HasRoleForUser(name, role)
}

func (re RBACCasbin) AddRoleForUser(name string, role string) (bool, error) {
	return re.enforcer.AddRoleForUser(name, role)
}

func (re RBACCasbin) AddRolesForUser(name string, roles []string) (bool, error) {
	return re.enforcer.AddRolesForUser(name, roles)
}

func (re RBACCasbin) DeleteRoleForUser(name string, role string) (bool, error) {
	return re.enforcer.DeleteRoleForUser(name, role)
}

func (re RBACCasbin) DeleteRolesForUser(name string) (bool, error) {
	return re.enforcer.DeleteRolesForUser(name)
}

func (re RBACCasbin) DeleteUser(name string) (bool, error) {
	return re.enforcer.DeleteUser(name)
}

func (re RBACCasbin) DeleteRole(name string) (bool, error) {
	return re.enforcer.DeleteRole(name)
}

func (re RBACCasbin) DeletePermission(name string) (bool, error) {
	return re.enforcer.DeletePermission(name)
}

func (re RBACCasbin) AddPermissionForUser(name string, permission ...string) (bool, error) {
	return re.enforcer.AddPermissionForUser(name, permission...)
}

func (re RBACCasbin) DeletePermissionForUser(name string, permission ...string) (bool, error) {
	return re.enforcer.DeletePermissionForUser(name, permission...)
}

func (re RBACCasbin) DeletePermissionsForUser(name string) (bool, error) {
	return re.enforcer.DeletePermissionsForUser(name)
}

func (re RBACCasbin) GetPermissionsForUser(name string) [][]string {
	return re.enforcer.GetPermissionsForUser(name)
}

func (re RBACCasbin) HasPermissionForUser(name string, permission string) bool {
	return re.enforcer.HasPermissionForUser(name, permission)
}

func (re RBACCasbin) GetImplicitRolesForUser(name string) ([]string, error) {
	return re.enforcer.GetImplicitRolesForUser(name)
}

func (re RBACCasbin) GetImplicitPermissionsForUser(name string) ([][]string, error) {
	return re.enforcer.GetImplicitPermissionsForUser(name)
}

func (re RBACCasbin) GetImplicitUsersForPermission(name string) ([]string, error) {
	return re.enforcer.GetImplicitUsersForPermission(name)
}
