package domain

type PermissionManager interface {
	AddRole(roleID uint) (bool, error)

	AddRoleForUser(userID uint, roleID uint) (bool, error)

	AddRolesForUser(userID uint, roleIDs []uint) (bool, error)

	DeleteRoleForUser(userID uint, roleID uint) (bool, error)

	DeleteRolesForUser(userID uint, roleIDs []uint) (bool, error)

	AddPermissionForRole(roleID uint, resource Resource) (bool, error)

	AddPermissionForRoles(roleIDs []uint, resource Resource) (bool, error)

	AddPermissionsForRole(roleID uint, resources Resources) (bool, error)

	RemovePermissionForRole(roleID uint, resource Resource) (bool, error)

	RemovePermissionForRoles(roleIDs []uint, resource Resource) (bool, error)

	RemovePermissionsForRole(roleID uint, resources Resources) (bool, error)

	UpdatePermissionForRole(roleID uint, oldResource, newResource Resource) (bool, error)

	UpdatePermissionForRoles(roleIDs []uint, oldResource, newResource Resource) (bool, error)
}
