package rbac

import (
	"context"

	"github.com/go-qbit/qerror"
	"github.com/go-qbit/rbac"

	"riceboards/authctx"
)

const (
	RoleAdmin Role = iota
	RoleUser
)

type Role uint8

type RbacStorage struct {
	//RolesPermissions *roles_permissions.RolesPermissions
	//Users            *users.Users
}

func (s *RbacStorage) AddRole(ctx context.Context, caption string) (interface{}, error) {
	return nil, qerror.Errorf("Unimplemented")
}

func (s *RbacStorage) GetRoles(ctx context.Context, ids ...interface{}) ([]rbac.RoleDescription, error) {
	return []rbac.RoleDescription{
		{Id: RoleAdmin, Caption: "Admin"},
		{Id: RoleUser, Caption: "User"},
	}, nil
}

func (s *RbacStorage) AddRolesPermissions(ctx context.Context, rolesPermissions ...rbac.RolePermission) error {
	return qerror.Errorf("Unimplemented")
	//_, err := s.RolesPermissions.AddFromStructs(ctx, rolesPermissions, model.AddOptions{})
	//return err
}

func (s *RbacStorage) RevokeRolesPermissions(ctx context.Context, rolesPermissions ...rbac.RolePermission) error {
	return qerror.Errorf("Unimplemented")
	//rpfilters := make([]model.IExpression, 0, len(rolesPermissions))
	//for _, rp := range rolesPermissions {
	//	rpfilters = append(rpfilters, expr.And(
	//		expr.Eq(s.RolesPermissions.FieldExpr("role_id"), expr.Value(rp.RoleId)),
	//		expr.Eq(s.RolesPermissions.FieldExpr("permission_id"), expr.Value(rp.PermissionId)),
	//	))
	//}
	//
	//var filter model.IExpression
	//switch len(rpfilters) {
	//case 0:
	//	return nil
	//case 1:
	//	filter = rpfilters[0]
	//case 2:
	//	filter = expr.Or(rpfilters[0], rpfilters[1])
	//default:
	//	filter = expr.Or(rpfilters[0], rpfilters[1], rpfilters[2:]...)
	//}
	//
	//return s.RolesPermissions.Delete(ctx, filter)
}

func (s *RbacStorage) GetRolesPermissions(ctx context.Context, rolesIds ...interface{}) ([]string, error) {
	if len(rolesIds) == 0 {
		return []string{}, nil
	}

	/*
		var rolesPerms []struct {
			RoleId       string
			PermissionId string
		}

		filter := expr.In(s.RolesPermissions.FieldExpr("role_id"))
		for _, roleId := range rolesIds {
			filter.Add(expr.Value(roleId))
		}

		if err := s.RolesPermissions.GetAllToStruct(context.Background(), &rolesPerms, model.GetAllOptions{Filter: filter}); err != nil {
			return nil, err
		}

		rolesPermsMap := make(map[string][]string)
		for _, rp := range rolesPerms {
			rolesPermsMap[rp.RoleId] = append(rolesPermsMap[rp.RoleId], rp.PermissionId)
		}

		res := make([]string, 0, len(rolesPermsMap))

		for _, perms := range rolesPermsMap {
			for _, permissionId := range perms {
				res = append(res, permissionId)
			}
		}
	*/

	var res []string
	for _, roleId := range rolesIds {
		switch roleId.(Role) {
		case RoleAdmin:
			for _, group := range rbac.GetAllPermissionsGroups() {
				for _, p := range group.GetAllPermissions() {
					res = append(res, p.GetFullId())
				}
			}

		case RoleUser:
			res = append(res)
		}
	}

	return res, nil
}

func (s *RbacStorage) AddUserRoles(ctx context.Context, usersRoles ...rbac.UserRole) error {
	return qerror.Errorf("Unimplemented")
}

func (s *RbacStorage) RevokeUserRoles(ctx context.Context, usersRoles ...rbac.UserRole) error {
	return qerror.Errorf("Unimplemented")
}

func (s *RbacStorage) GetUserRoles(ctx context.Context, userId interface{}) ([]interface{}, error) {
	if userId != nil {
		return nil, qerror.Errorf("Getting roles for not current user is not implemented")
	}

	user := authctx.FromContext(ctx)
	if user == nil { // No user, no roles
		return nil, nil
	}

	if user.IsAdmin {
		return []interface{}{RoleAdmin}, nil
	}

	return []interface{}{RoleUser}, nil
}
