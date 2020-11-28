/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//  This is the HCS Schema version of the HvSocket configuration. The VMWP version is  located in Config.Devices.IC in V1.
type HvSocketSystemConfig struct {

	//  SDDL string that HvSocket will check before allowing a host process to bind  to an unlisted service for this specific container/VM (not wildcard binds).
	DefaultBindSecurityDescriptor string `json:"DefaultBindSecurityDescriptor,omitempty"`

	//  SDDL string that HvSocket will check before allowing a host process to connect  to an unlisted service in the VM/container.
	DefaultConnectSecurityDescriptor string `json:"DefaultConnectSecurityDescriptor,omitempty"`

	ServiceTable map[string]HvSocketServiceConfig `json:"ServiceTable,omitempty"`
}
