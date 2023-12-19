package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/azurerm/data/resourcegroups/azurerm_resource_group.json
var dataAzurermResourceGroup []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault.json
var dataAzurermKeyVault []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_key.json
var dataAzurermKeyVaultKey []byte

//go:embed mapping/azurerm/data/network/azurerm_subnet.json
var dataAzurermSubnet []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network.json
var dataAzurermVirtualNetwork []byte

//go:embed mapping/azurerm/data/resources/azurerm_subscription.json
var dataAzurermSubscription []byte

//go:embed mapping/azurerm/data/network/azurerm_network_watcher.json
var dataAzurermNetworkWatcher []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_account.json
var dataAzurermStorageAccount []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate.json
var dataAzurermKeyVaultCertificate []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate_data.json
var dataAzurermKeyVaultCertificateData []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate_issuer.json
var dataAzurermKeyVaultCertificateIssuer []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificates.json
var dataAzurermKeyVaultCertificates []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_managed_hardware_security_module.json
var dataAzurermKeyVaultManagedHardwareSecurityModule []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_secret.json
var dataAzurermKeyVaultSecret []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_secrets.json
var dataAzurermKeyVaultSecrets []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service.json
var dataAzurermAppService []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_certificate.json
var dataAzurermAppServiceCertificate []byte

//go:embed mapping/azurerm/data/resources/azurerm_app_service_certificate_order.json
var dataAzurermAppServiceCertificateOrder []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_environment.json
var dataAzurermAppServiceEnvironment []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_environment_v3.json
var dataAzurermAppServiceEnvironmentV3 []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_plan.json
var dataAzurermAppServicePlan []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ip.json
var dataAzurermPublicIP []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ip_prefix.json
var dataAzurermPublicIPPrefix []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ips.json
var dataAzurermPublicIps []byte

//go:embed mapping/azurerm/data/web/azurerm_windows_function_app.json
var dataAzurermWindowsFunctionApp []byte

//go:embed mapping/azurerm/data/web/azurerm_windows_web_app.json
var dataAzurermWindowsWebApp []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_blob.json
var dataAzurermStorageBlob []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_container.json
var dataAzurermStorageContainer []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_encryption_scope.json
var dataAzurermStorageEncryptionScope []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_management_policy.json
var dataAzurermStorageManagementPolicy []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_share.json
var dataAzurermStorageShare []byte

//go:embed mapping/azurerm/data/storagesync/azurerm_storage_sync.json
var dataAzurermStorageSync []byte

//go:embed mapping/azurerm/data/storagesync/azurerm_storage_sync_group.json
var dataAzurermStorageSyncGroup []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_table_entity.json
var dataAzurermStorageTableEntity []byte

//go:embed mapping/azurerm/data/compute/azurerm_ssh_public_key.json
var dataAzurermSSHPublicKey []byte

//go:embed mapping/azurerm/data/streamanalytics/azurerm_stream_analytics_job.json
var dataAzurermStreamAnalyticsJob []byte

//go:embed mapping/azurerm/data/resources/azurerm_subscription_template_deployment.json
var dataAzurermSubscriptionTemplateDeployment []byte

//go:embed mapping/azurerm/data/synapase/azurerm_synapse_workspace.json
var dataAzurermSynapseWorkspace []byte

//go:embed mapping/azurerm/data/resources/azurerm_template_spec_version.json
var dataAzurermTemplateSpecVersion []byte

//go:embed mapping/azurerm/data/resources/azurerm_tenant_template_deployment.json
var dataAzurermTenantTemplateDeployment []byte

//go:embed mapping/azurerm/data/network/azurerm_traffic_manager_profile.json
var dataAzurermTrafficManagerProfile []byte

//go:embed mapping/azurerm/data/managedidentites/azurerm_user_assigned_identity.json
var dataAzurermUserAssignedIdentity []byte

//go:embed mapping/azurerm/data/desktopvirtualization/azurerm_virtual_desktop_host_pool.json
var dataAzurermVirtualDesktopHostPool []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_hub.json
var dataAzurermVirtualHub []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_hub_connection.json
var dataAzurermVirtualHubConnection []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_hub_route_table.json
var dataAzurermVirtualHubRouteTable []byte

//go:embed mapping/azurerm/data/compute/azurerm_virtual_machine.json
var dataAzurermVirtualMachine []byte

//go:embed mapping/azurerm/data/compute/azurerm_virtual_machine_scale_set.json
var dataAzurermVirtualNetworkScaleSet []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network_gateway.json
var dataAzurermVirtualNetworkGateway []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network_gateway_connection.json
var dataAzurermVirtualNetworkGatewayConnection []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_wan.json
var dataAzurermVirtualWan []byte

//go:embed mapping/azurerm/data/avs/azurerm_vmware_private_cloud.json
var dataAzurermVmwarePrivateCloud []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network_gateway.json
var dataAzurermVpnGateway []byte

//go:embed mapping/azurerm/data/network/azurerm_web_application_firewall_policy.json
var dataAzurermWebApplicationFirewallPolicy []byte

//go:embed mapping/azurerm/data/signalrservice/azurerm_web_pubsub.json
var dataAzurermWebpubsub []byte

//go:embed mapping/azurerm/data/cache/azurerm_redis_cache.json
var dataAzurermRedisCache []byte

//go:embed mapping/azurerm/data/cache/azurerm_redis_enterprise_database.json
var dataAzurermRedisEnterpriseDatabase []byte

//go:embed mapping/azurerm/data/resources/azurerm_resource_group_template_deployment.json
var dataAzurermResourceGroupTemplateDeployment []byte

//go:embed mapping/azurerm/data/network/azurerm_route_filter.json
var dataAzurermRouteFilter []byte

//go:embed mapping/azurerm/data/network/azurerm_route_table.json
var dataAzurermRouteTable []byte

//go:embed mapping/azurerm/data/securityinsights/azurerm_sentinel_alert_rule.json
var dataAzurermSentinelAlertRule []byte

//go:embed mapping/azurerm/data/web/azurerm_service_plan.json
var dataAzurermServicePlan []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_namespace.json
var dataAzurermServicebusNamespace []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_namespace_authorization_rule.json
var dataAzurermServicebusNamespaceAuthorizationRule []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_queue.json
var dataAzurermServicebusQueue []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_namespace_authorization_rule.json
var dataAzurermServicebusQueueAuthorizationRule []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_subscription.json
var dataAzurermServicebusSubscription []byte

//go:embed mapping/azurerm/data/servicebus/azurerm_servicebus_topic.json
var dataAzurermServicebusTopic []byte

//go:embed mapping/azurerm/data/compute/azurerm_shared_image.json
var dataAzurermSharedImage []byte

//go:embed mapping/azurerm/data/compute/azurerm_shared_image_gallery.json
var dataAzurermSharedImageGallery []byte

//go:embed mapping/azurerm/data/signalrservice/azurerm_signalr_service.json
var dataAzurermSignalrService []byte

//go:embed mapping/azurerm/data/recoveryservices/azurerm_site_recovery_fabric.json
var dataAzurermSiteRecoveryFabric []byte

//go:embed mapping/azurerm/data/recoveryservices/azurerm_site_recovery_protection_container.json
var dataAzurermSiteRecoveryProtectionContainer []byte

//go:embed mapping/azurerm/data/recoveryservices/azurerm_site_recovery_replication_policy.json
var dataAzurermSiteRecoveryReplicationPolicy []byte

//go:embed mapping/azurerm/data/appplatform/azurerm_spring_cloud_app.json
var dataAzurermSpringCloudApp []byte

//go:embed mapping/azurerm/data/appplatform/azurerm_spring_cloud_service.json
var dataAzurermSpringCloudService []byte

//go:embed mapping/azurerm/data/sql/azurerm_sql_database.json
var dataAzurermSqlDatabase []byte

//go:embed mapping/azurerm/data/sql/azurerm_sql_managed_instance.json
var dataAzurermSqlManagedInstance []byte

//go:embed mapping/azurerm/data/sql/azurerm_sql_server.json
var dataAzurermSqlServer []byte

//go:embed mapping/azurerm/data/azurestackhci/azurerm_stack_hci_cluster.json
var dataAzurermStackHCICluster []byte

////go:embed mapping/azurerm/data/
//var dataAzurermSnapshot []byte
//
////go:embed mapping/azurerm/data/
//var dataAzurermSourceControlToken []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management.json
var dataAzurermApiManagement []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_api.json
var dataAzurermApiManagementApi []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_api_version_set.json
var dataAzurermApiManagementVersionSet []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_gateway.json
var dataAzurermApiManagementGateway []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_gateway_host_name_configuration.json
var dataAzurermApiManagementGatewayHostNameConfiguration []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_group.json
var dataAzurermApiManagementGroup []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_product.json
var dataAzurermApiManagmentProduct []byte

//go:embed mapping/azurerm/data/apimanagement/azurerm_api_management_user.json
var dataAzurermApiManagementUser []byte

//go:embed mapping/azurerm/data/appconfiguration/azurerm_app_configuration.json
var dataAzurermAppConfiguration []byte

//go:embed mapping/azurerm/data/appconfiguration/azurerm_app_configuration_key.json
var dataAzurermAppConfigurationKey []byte

//go:embed mapping/azurerm/data/appconfiguration/azurerm_app_configuration_keys.json
var dataAzurermAppConfigurationKeys []byte

//go:embed mapping/azurerm/data/network/azurerm_application_gateway.json
var dataAzurermApplicationGateway []byte

//go:embed mapping/azurerm/data/insights/azurerm_application_insights.json
var dataAzurermApplicationInsights []byte

//go:embed mapping/azurerm/data/network/azurerm_application_security_group.json
var dataAzurermApplicationSecurityGroup []byte

//go:embed mapping/azurerm/data/hybridcompute/azurerm_arc_machine.json
var dataAzurermArcMachine []byte

//go:embed mapping/azurerm/data/automation/azurerm_automation_account.json
var dataAzurermAutomationAccount []byte

//go:embed mapping/azurerm/data/automation/azurerm_automation_variable.json
var dataAzurermAutomationVariable []byte

//go:embed mapping/azurerm/data/compute/azurerm_availability_set.json
var dataAzurermAvailabilitySet []byte

//go:embed mapping/azurerm/data/recoveryservices/azurerm_backup_policy_file_share.json
var dataAzurermBackupPolicyFileShare []byte

//go:embed mapping/azurerm/data/recoveryservices/azurerm_backup_policy_vm.json
var dataAzurermBackupPolicyVm []byte

//go:embed mapping/azurerm/data/network/azurerm_bastion_host.json
var dataAzurermBastionHost []byte

//go:embed mapping/azurerm/data/batch/azurerm_batch_account.json
var dataAzurermBatchAccount []byte

//go:embed mapping/azurerm/data/batch/azurerm_batch_application.json
var dataAzurermBatchApplication []byte

//go:embed mapping/azurerm/data/batch/azurerm_batch_certificate.json
var dataAzurermBatchCertificate []byte

//go:embed mapping/azurerm/data/batch/azurerm_batch_pool.json
var dataAzurermBatchPool []byte
