package enum

const (
	// SuperAdminUser is super admin in enterprise, who cannot be modified
	// Only super admin can change user type in enterprise
	// Also, super admin can use all module of system
	SuperAdminUser = 0

	// AdminUser is normal admin user in enterprise, who can add user into
	// enterprise as a normal user, and user all module of system
	AdminUser = 1

	// NormalUser is normal user in enterprise, who can only use modules
	// by privilege setting
	NormalUser = 2
)
