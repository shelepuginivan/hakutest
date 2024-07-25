package security

import mapset "github.com/deckarep/golang-set/v2"

// HasPermissions reports whether user has permission for the action based on
// roles.
func HasPermissions(c *Credentials, requiredRoles []string) bool {
	userRoles := mapset.NewSet(c.Roles...)
	wantRoles := mapset.NewSet(requiredRoles...)

	return userRoles.IsSuperset(wantRoles)
}
