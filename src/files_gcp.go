package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/google/resource/backend/gcs.json
var gcsBackend []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_level.json
var googleAccessContextManagerAccessLevel []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_levels.json
var googleAccessContextManagerAccessLevels []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_policy.json
var googleAccessContextManagerAccessPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_policy_iam.json
var googleAccessContextManagerAccessPolicyIAM []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_authorized_orgs_desc.json
var googleAccessContextManagerAuthorizedOrgsDesc []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_gcp_user_access_binding.json
var googleAccessContextManagerGcpUserAccessBinding []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter.json
var googleAccessContextManagerServicePerimeter []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeters.json
var googleAccessContextManagerServicePerimeters []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_backup.json
var googleAlloydbBackup []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_cluster.json
var googleAlloydbCluster []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_instance.json
var googleAlloydbInstance []byte

//go:embed mapping/google/resource/alloydb/google_alloydb_user.json
var googleAlloydbUser []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api.json
var googleAPIGatewayAPI []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config.json
var googleAPIGatewayAPIConfig []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_config_iam.json
var googleAPIGatewayAPIConfigIAM []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_api_iam.json
var googleAPIGatewayAPIIAM []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway.json
var googleAPIGatewayGateway []byte

//go:embed mapping/google/resource/apigateway/google_api_gateway_gateway_iam.json
var googleAPIGatewayGatewayIAM []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment.json
var googleApigeeEnvironment []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_binding.json
var googleApigeeEnvironmentIAMBinding []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_member.json
var googleApigeeEnvironmentIAMMember []byte

//go:embed mapping/google/resource/apigee/google_apigee_environment_iam_policy.json
var googleApigeeEnvironmentIAMPolicy []byte

//go:embed mapping/google/resource/apihub/google_apihub_api_hub_instance.json
var googleAPIhubAPIHubInstance []byte

//go:embed mapping/google/resource/apihub/google_apihub_curation.json
var googleAPIhubCuration []byte

//go:embed mapping/google/resource/apihub/google_apihub_host_project_registration.json
var googleAPIhubHostProjectRegistration []byte

//go:embed mapping/google/resource/apihub/google_apihub_plugin.json
var googleAPIhubPlugin []byte

//go:embed mapping/google/resource/apihub/google_apihub_plugin_instance.json
var googleAPIhubPluginInstance []byte

//go:embed mapping/google/resource/apphub/google_apphub_application.json
var googleApphubApplication []byte

//go:embed mapping/google/resource/apphub/google_apphub_service.json
var googleApphubService []byte

//go:embed mapping/google/resource/apphub/google_apphub_service_project_attachment.json
var googleApphubServiceProjectAttachment []byte

//go:embed mapping/google/resource/apphub/google_apphub_workload.json
var googleApphubWorkload []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository.json
var googleArtifactRegistryRepository []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_binding.json
var googleArtifactRegistryRepositoryIAMBinding []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_member.json
var googleArtifactRegistryRepositoryIAMMember []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_repository_iam_policy.json
var googleArtifactRegistryRepositoryIAMPolicy []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_backup_plan.json
var googleBackupDrBackupPlan []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_backup_plan_association.json
var googleBackupDrBackupPlanAssociation []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_backup_vault.json
var googleBackupDrBackupVault []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_management_server.json
var googleBackupDrManagementServer []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_service_config.json
var googleBackupDrServiceConfig []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_app_connection.json
var googleBeyondcorpAppConnection []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_app_connector.json
var googleBeyondcorpAppConnector []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_app_gateway.json
var googleBeyondcorpAppGateway []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application.json
var googleBeyondcorpApplication []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_binding.json
var googleBeyondcorpApplicationIAMBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_member.json
var googleBeyondcorpApplicationIAMMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_application_iam_policy.json
var googleBeyondcorpApplicationIAMPolicy []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway.json
var googleBeyondcorpSecurityGateway []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application.json
var googleBeyondcorpSecurityGatewayApplication []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_binding.json
var googleBeyondcorpSecurityGatewayApplicationIAMBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_member.json
var googleBeyondcorpSecurityGatewayApplicationIAMMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_application_iam_policy.json
var googleBeyondcorpSecurityGatewayApplicationIAMPolicy []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_binding.json
var googleBeyondcorpSecurityGatewayIAMBinding []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_member.json
var googleBeyondcorpSecurityGatewayIAMMember []byte

//go:embed mapping/google/resource/beyondcorp/google_beyondcorp_security_gateway_iam_policy.json
var googleBeyondcorpSecurityGatewayIAMPolicy []byte

//go:embed mapping/google/resource/biglake/google_biglake_catalog.json
var googleBiglakeCatalog []byte

//go:embed mapping/google/resource/biglake/google_biglake_database.json
var googleBiglakeDatabase []byte

//go:embed mapping/google/resource/biglake/google_biglake_table.json
var googleBiglakeTable []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_data_exchange.json
var googleBigqueryAnalyticsHubDataExchange []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_data_exchange_iam.json
var googleBigqueryAnalyticsHubDataExchangeIAM []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing.json
var googleBigqueryAnalyticsHubListing []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing_iam.json
var googleBigqueryAnalyticsHubListingIAM []byte

//go:embed mapping/google/resource/analyticshub/google_bigquery_analytics_hub_listing_subscription.json
var googleBigqueryAnalyticsHubListingSubscription []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_bi_reservation.json
var googleBigqueryBiReservation []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_capacity_commitment.json
var googleBigqueryCapacityCommitment []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_connection.json
var googleBigqueryConnection []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_connection_iam.json
var googleBigqueryConnectionIAM []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_data_transfer_config.json
var googleBigqueryDataTransferConfig []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy.json
var googleBigqueryDatapolicyDataPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_binding.json
var googleBigqueryDatapolicyDataPolicyIAMBinding []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_member.json
var googleBigqueryDatapolicyDataPolicyIAMMember []byte

//go:embed mapping/google/resource/datacatalog/google_bigquery_datapolicy_data_policy_iam_policy.json
var googleBigqueryDatapolicyDataPolicyIAMPolicy []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset.json
var googleBigqueryDataset []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset_access.json
var googleBigqueryDatasetAccess []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_dataset_iam.json
var googleBigqueryDatasetIAM []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_job.json
var googleBigqueryJob []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_reservation.json
var googleBigqueryReservation []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_reservation_assignment.json
var googleBigqueryReservationAssignment []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_routine.json
var googleBigqueryRoutine []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_row_access_policy.json
var googleBigqueryRowAccessPolicy []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_table.json
var googleBigqueryTable []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_table_iam.json
var googleBigqueryTableIAM []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_app_profile.json
var googleBigtableAppProfile []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_authorized_view.json
var googleBigtableAuthorizedView []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_gc_policy.json
var googleBigtableGcPolicy []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_instance.json
var googleBigtableInstance []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_instance_iam.json
var googleBigtableInstanceIAM []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_logical_view.json
var googleBigtableLogicalView []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_materialized_view.json
var googleBigtableMaterializedView []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_schema_bundle.json
var googleBigtableSchemaBundle []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_table.json
var googleBigtableTable []byte

//go:embed mapping/google/resource/bigtable/google_bigtable_table_iam.json
var googleBigtableTableIAM []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_binding.json
var googleBillingAccountIAMBinding []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_member.json
var googleBillingAccountIAMMember []byte

//go:embed mapping/google/resource/billing/google_billing_account_iam_policy.json
var googleBillingAccountIAMPolicy []byte

//go:embed mapping/google/resource/billing/google_billing_budget.json
var googleBillingBudget []byte

//go:embed mapping/google/resource/resourcemanager/google_billing_project_info.json
var googleBillingProjectInfo []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_binding.json
var googleBinaryAuthorizationAttestorIAMBinding []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_member.json
var googleBinaryAuthorizationAttestorIAMMember []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor_iam_policy.json
var googleBinaryAuthorizationAttestorIAMPolicy []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_dns_authorization.json
var googleCertificateManagerDNSAuthorization []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_data_access_label.json
var googleChronicleDataAccessLabel []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_data_access_scope.json
var googleChronicleDataAccessScope []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_reference_list.json
var googleChronicleReferenceList []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_retrohunt.json
var googleChronicleRetrohunt []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_rule.json
var googleChronicleRule []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_rule_deployment.json
var googleChronicleRuleDeployment []byte

//go:embed mapping/google/resource/chronicle/google_chronicle_watchlist.json
var googleChronicleWatchlist []byte

//go:embed mapping/google/resource/run/google_cloud_run_domain_mapping.json
var googleCloudRunDomainMapping []byte

//go:embed mapping/google/resource/run/google_cloud_run_service.json
var googleCloudRunService []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_binding.json
var googleCloudRunServiceIAMBinding []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_member.json
var googleCloudRunServiceIAMMember []byte

//go:embed mapping/google/resource/run/google_cloud_run_service_iam_policy.json
var googleCloudRunServiceIAMPolicy []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_job.json
var googleCloudRunV2Job []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_job_iam.json
var googleCloudRunV2JobIAM []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_service.json
var googleCloudRunV2Service []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_service_iam.json
var googleCloudRunV2ServiceIAM []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool.json
var googleCloudRunV2WorkerPool []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_binding.json
var googleCloudRunV2WorkerPoolIAMBinding []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_member.json
var googleCloudRunV2WorkerPoolIAMMember []byte

//go:embed mapping/google/resource/run/google_cloud_run_v2_worker_pool_iam_policy.json
var googleCloudRunV2WorkerPoolIAMPolicy []byte

//go:embed mapping/google/resource/cloudscheduler/google_cloud_scheduler_job.json
var googleCloudSchedulerJob []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_binding.json
var googleCloudTasksQueueIAMBinding []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_member.json
var googleCloudTasksQueueIAMMember []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue_iam_policy.json
var googleCloudTasksQueueIAMPolicy []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuild_trigger.json
var googleCloudbuildTrigger []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection.json
var googleCloudbuildV2Connection []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_binding.json
var googleCloudbuildV2ConnectionIAMBinding []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_member.json
var googleCloudbuildV2ConnectionIAMMember []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_connection_iam_policy.json
var googleCloudbuildV2ConnectionIAMPolicy []byte

//go:embed mapping/google/resource/cloudbuild/google_cloudbuildv2_repository.json
var googleCloudbuildV2Repository []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_binding.json
var googleClouddeployCustomTargetTypeIAMBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_member.json
var googleClouddeployCustomTargetTypeIAMMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type_iam_policy.json
var googleClouddeployCustomTargetTypeIAMPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_binding.json
var googleClouddeployDeliveryPIPelineIAMBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_member.json
var googleClouddeployDeliveryPIPelineIAMMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline_iam_policy.json
var googleClouddeployDeliveryPIPelineIAMPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_binding.json
var googleClouddeployTargetIAMBinding []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_member.json
var googleClouddeployTargetIAMMember []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target_iam_policy.json
var googleClouddeployTargetIAMPolicy []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function.json
var googleCloudfunctionsFunction []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function_iam_binding.json
var googleCloudfunctionsFunctionIAMBinding []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions_function_iam_policy.json
var googleCloudfunctionsFunctionIAMPolicy []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function.json
var googleCloudfunctions2Function []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_binding.json
var googleCloudfunctions2FunctionIAMBinding []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_member.json
var googleCloudfunctions2FunctionIAMMember []byte

//go:embed mapping/google/resource/cloudfunctions/google_cloudfunctions2_function_iam_policy.json
var googleCloudfunctions2FunctionIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_binding.json
var googleColabRuntimeTemplateIAMBinding []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_member.json
var googleColabRuntimeTemplateIAMMember []byte

//go:embed mapping/google/resource/aiplatform/google_colab_runtime_template_iam_policy.json
var googleColabRuntimeTemplateIAMPolicy []byte

//go:embed mapping/google/resource/composer/google_composer_environment.json
var googleComposerEnvironment []byte

//go:embed mapping/google/resource/composer/google_composer_user_workloads_config_map.json
var googleComposerUserWorkloadsConfigMap []byte

//go:embed mapping/google/resource/composer/google_composer_user_workloads_secret.json
var googleComposerUserWorkloadsSecret []byte

//go:embed mapping/google/resource/compute/google_compute_address.json
var googleComputeAddress []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket.json
var googleComputeBackendBucket []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_binding.json
var googleComputeBackendBucketIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_member.json
var googleComputeBackendBucketIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_backend_bucket_iam_policy.json
var googleComputeBackendBucketIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service.json
var googleComputeBackendService []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_binding.json
var googleComputeBackendServiceIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_member.json
var googleComputeBackendServiceIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_backend_service_iam_policy.json
var googleComputeBackendServiceIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_binding.json
var googleComputeDiskIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_member.json
var googleComputeDiskIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_disk_iam_policy.json
var googleComputeDiskIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_firewall.json
var googleComputeFirewall []byte

//go:embed mapping/google/resource/compute/google_compute_forwarding_rule.json
var googleComputeForwardingRule []byte

//go:embed mapping/google/resource/compute/google_compute_global_address.json
var googleComputeGlobalAddress []byte

//go:embed mapping/google/resource/compute/google_compute_global_forwarding_rule.json
var googleComputeGlobalForwardingRule []byte

//go:embed mapping/google/resource/compute/google_compute_global_network_endpoint_group.json
var googleComputeGlobalNetworkEndpointGroup []byte

//go:embed mapping/google/resource/compute/google_compute_health_check.json
var googleComputeHealthCheck []byte

//go:embed mapping/google/resource/compute/google_compute_http_health_check.json
var googleComputeHttpHealthCheck []byte

//go:embed mapping/google/resource/compute/google_compute_https_health_check.json
var googleComputeHttpsHealthCheck []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_binding.json
var googleComputeImageIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_member.json
var googleComputeImageIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_image_iam_policy.json
var googleComputeImageIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_instance.json
var googleComputeInstance []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_binding.json
var googleComputeInstanceIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_member.json
var googleComputeInstanceIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_instance_iam_policy.json
var googleComputeInstanceIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template.json
var googleComputeInstanceTemplate []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_binding.json
var googleComputeInstanceTemplateIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_member.json
var googleComputeInstanceTemplateIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_instance_template_iam_policy.json
var googleComputeInstanceTemplateIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_binding.json
var googleComputeInstantSnapshotIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_member.json
var googleComputeInstantSnapshotIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_instant_snapshot_iam_policy.json
var googleComputeInstantSnapshotIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_binding.json
var googleComputeMachineImageIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_member.json
var googleComputeMachineImageIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image_iam_policy.json
var googleComputeMachineImageIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_network.json
var googleComputeNetwork []byte

//go:embed mapping/google/resource/compute/google_compute_network_attachment.json
var googleComputeNetworkAttachment []byte

//go:embed mapping/google/resource/compute/google_compute_network_endpoint_group.json
var googleComputeNetworkEndpointGroup []byte

//go:embed mapping/google/resource/compute/google_compute_project_metadata_item.json
var googleComputeProjectMetadataItem []byte

//go:embed mapping/google/resource/compute/google_compute_router.json
var googleComputeRouter []byte

//go:embed mapping/google/resource/compute/google_compute_router_nat.json
var googleComputeRouterNat []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service.json
var googleComputeRegionBackendService []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_binding.json
var googleComputeRegionBackendServiceIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_member.json
var googleComputeRegionBackendServiceIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_region_backend_service_iam_policy.json
var googleComputeRegionBackendServiceIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_binding.json
var googleComputeRegionDiskIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_member.json
var googleComputeRegionDiskIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_region_disk_iam_policy.json
var googleComputeRegionDiskIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_region_health_check.json
var googleComputeRegionHealthCheck []byte

//go:embed mapping/google/resource/compute/google_compute_region_network_endpoint_group.json
var googleComputeRegionNetworkEndpointGroup []byte

//go:embed mapping/google/resource/compute/google_compute_region_ssl_certificate.json
var googleComputeRegionSslCertificate []byte

//go:embed mapping/google/resource/compute/google_compute_region_target_http_proxy.json
var googleComputeRegionTargetHttpProxy []byte

//go:embed mapping/google/resource/compute/google_compute_region_target_https_proxy.json
var googleComputeRegionTargetHttpsProxy []byte

//go:embed mapping/google/resource/compute/google_compute_region_target_tcp_proxy.json
var googleComputeRegionTargetTcpProxy []byte

//go:embed mapping/google/resource/compute/google_compute_region_url_map.json
var googleComputeRegionUrlMap []byte

//go:embed mapping/google/resource/compute/google_compute_security_policy.json
var googleComputeSecurityPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_binding.json
var googleComputeSnapshotIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_member.json
var googleComputeSnapshotIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_snapshot_iam_policy.json
var googleComputeSnapshotIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_binding.json
var googleComputeStoragePoolIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_member.json
var googleComputeStoragePoolIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_storage_pool_iam_policy.json
var googleComputeStoragePoolIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork.json
var googleComputeSubnetwork []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_binding.json
var googleComputeSubnetworkIAMBinding []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_member.json
var googleComputeSubnetworkIAMMember []byte

//go:embed mapping/google/resource/compute/google_compute_subnetwork_iam_policy.json
var googleComputeSubnetworkIAMPolicy []byte

//go:embed mapping/google/resource/compute/google_compute_target_http_proxy.json
var googleComputeTargetHttpProxy []byte

//go:embed mapping/google/resource/compute/google_compute_target_https_proxy.json
var googleComputeTargetHttpsProxy []byte

//go:embed mapping/google/resource/compute/google_compute_url_map.json
var googleComputeUrlMap []byte

//go:embed mapping/google/resource/contactcenterinsights/google_contact_center_insights_analysis_rule.json
var googleContactCenterInsightsAnalysisRule []byte

//go:embed mapping/google/resource/contactcenterinsights/google_contact_center_insights_view.json
var googleContactCenterInsightsView []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_binding.json
var googleContainerAnalysisNoteIAMBinding []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_member.json
var googleContainerAnalysisNoteIAMMember []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note_iam_policy.json
var googleContainerAnalysisNoteIAMPolicy []byte

//go:embed mapping/google/resource/container/google_container_cluster.json
var googleContainerCluster []byte

//go:embed mapping/google/resource/container/google_container_node_pool.json
var googleContainerNodePool []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry.json
var googleDataCatalogEntry []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group.json
var googleDataCatalogEntryGroup []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group_iam_binding.json
var googleDataCatalogEntryGroupIAMBinding []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group_iam_member.json
var googleDataCatalogEntryGroupIAMMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_entry_group_iam_policy.json
var googleDataCatalogEntryGroupIAMPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag.json
var googleDataCatalogPolicyTag []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag_iam_binding.json
var googleDataCatalogPolicyTagIAMBinding []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag_iam_member.json
var googleDataCatalogPolicyTagIAMMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_policy_tag_iam_policy.json
var googleDataCatalogPolicyTagIAMPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag.json
var googleDataCatalogTag []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template.json
var googleDataCatalogTagTemplate []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_binding.json
var googleDataCatalogTagTemplateIAMBinding []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_member.json
var googleDataCatalogTagTemplateIAMMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_tag_template_iam_policy.json
var googleDataCatalogTagTemplateIAMPolicy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy.json
var googleDataCatalogTaxonomy []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy_iam_binding.json
var googleDataCatalogTaxonomyIAMBinding []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy_iam_member.json
var googleDataCatalogTaxonomyIAMMember []byte

//go:embed mapping/google/resource/datacatalog/google_data_catalog_taxonomy_iam_policy.json
var googleDataCatalogTaxonomyIAMPolicy []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_binding.json
var googleDataFusionInstanceIAMBinding []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_member.json
var googleDataFusionInstanceIAMMember []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance_iam_policy.json
var googleDataFusionInstanceIAMPolicy []byte

//go:embed mapping/google/resource/dataflow/google_dataflow_job.json
var googleDataflowJob []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository.json
var googleDataformRepository []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_binding.json
var googleDataformRepositoryIAMBinding []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_member.json
var googleDataformRepositoryIAMMember []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_iam_policy.json
var googleDataformRepositoryIAMPolicy []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_release_config.json
var googleDataformRepositoryReleaseConfig []byte

//go:embed mapping/google/resource/dataform/google_dataform_repository_workflow_config.json
var googleDataformRepositoryWorkflowConfig []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type.json
var googleDataplexAspectType []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_binding.json
var googleDataplexAspectTypeIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_member.json
var googleDataplexAspectTypeIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_aspect_type_iam_policy.json
var googleDataplexAspectTypeIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset.json
var googleDataplexAsset []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_binding.json
var googleDataplexAssetIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_member.json
var googleDataplexAssetIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_asset_iam_policy.json
var googleDataplexAssetIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan.json
var googleDataplexDatascan []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_binding.json
var googleDataplexDatascanIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_member.json
var googleDataplexDatascanIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_datascan_iam_policy.json
var googleDataplexDatascanIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry.json
var googleDataplexEntry []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group.json
var googleDataplexEntryGroup []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_binding.json
var googleDataplexEntryGroupIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_member.json
var googleDataplexEntryGroupIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_group_iam_policy.json
var googleDataplexEntryGroupIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_type.json
var googleDataplexEntryType []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_type_iam_binding.json
var googleDataplexEntryTypeIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_type_iam_member.json
var googleDataplexEntryTypeIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_type_iam_policy.json
var googleDataplexEntryTypeIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary.json
var googleDataplexGlossary []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_category.json
var googleDataplexGlossaryCategory []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_binding.json
var googleDataplexGlossaryIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_member.json
var googleDataplexGlossaryIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_iam_policy.json
var googleDataplexGlossaryIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_glossary_term.json
var googleDataplexGlossaryTerm []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake.json
var googleDataplexLake []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_binding.json
var googleDataplexLakeIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_member.json
var googleDataplexLakeIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_lake_iam_policy.json
var googleDataplexLakeIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task.json
var googleDataplexTask []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task_iam_binding.json
var googleDataplexTaskIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task_iam_member.json
var googleDataplexTaskIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_task_iam_policy.json
var googleDataplexTaskIAMPolicy []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone.json
var googleDataplexZone []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone_iam_binding.json
var googleDataplexZoneIAMBinding []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone_iam_member.json
var googleDataplexZoneIAMMember []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_zone_iam_policy.json
var googleDataplexZoneIAMPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy.json
var googleDataprocAutoscalingPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_binding.json
var googleDataprocAutoscalingPolicyIAMBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_member.json
var googleDataprocAutoscalingPolicyIAMMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_autoscaling_policy_iam_policy.json
var googleDataprocAutoscalingPolicyIAMPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_batch.json
var googleDataprocBatch []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster.json
var googleDataprocCluster []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_binding.json
var googleDataprocClusterIAMBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_member.json
var googleDataprocClusterIAMMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_cluster_iam_policy.json
var googleDataprocClusterIAMPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job.json
var googleDataprocJob []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_binding.json
var googleDataprocJobIAMBinding []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_member.json
var googleDataprocJobIAMMember []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_job_iam_policy.json
var googleDataprocJobIAMPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_binding.json
var googleDataprocMetastoreDatabaseIAMBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_member.json
var googleDataprocMetastoreDatabaseIAMMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_database_iam_policy.json
var googleDataprocMetastoreDatabaseIAMPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_binding.json
var googleDataprocMetastoreFederationIAMBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_member.json
var googleDataprocMetastoreFederationIAMMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_federation_iam_policy.json
var googleDataprocMetastoreFederationIAMPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_service_iam_binding.json
var googleDataprocMetastoreServiceIAMBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_service_iam_member.json
var googleDataprocMetastoreServiceIAMMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_service_iam_policy.json
var googleDataprocMetastoreServiceIAMPolicy []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_table_iam_binding.json
var googleDataprocMetastoreTableIAMBinding []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_table_iam_member.json
var googleDataprocMetastoreTableIAMMember []byte

//go:embed mapping/google/resource/metastore/google_dataproc_metastore_table_iam_policy.json
var googleDataprocMetastoreTableIAMPolicy []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_session_template.json
var googleDataprocSessionTemplate []byte

//go:embed mapping/google/resource/dataproc/google_dataproc_workflow_template.json
var googleDataprocWorkflowTemplate []byte

//go:embed mapping/google/resource/iam/google_default_service_accounts.json
var googleDefaultServiceAccounts []byte

//go:embed mapping/google/resource/deploymentmanager/google_deployment_manager_deployment.json
var googleDeploymentManagerDeployment []byte

//go:embed mapping/google/resource/developerconnect/google_developer_connect_account_connector.json
var googleDeveloperConnectAccountConnector []byte

//go:embed mapping/google/resource/developerconnect/google_developer_connect_connection.json
var googleDeveloperConnectConnection []byte

//go:embed mapping/google/resource/developerconnect/google_developer_connect_git_repository_link.json
var googleDeveloperConnectGitRepositoryLink []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_agent.json
var googleDialogflowAgent []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_conversation_profile.json
var googleDialogflowConversationProfile []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_agent.json
var googleDialogflowCxAgent []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_entity_type.json
var googleDialogflowCxEntityType []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_environment.json
var googleDialogflowCxEnvironment []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_flow.json
var googleDialogflowCxFlow []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_generative_settings.json
var googleDialogflowCxGenerativeSettings []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_generator.json
var googleDialogflowCxGenerator []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_intent.json
var googleDialogflowCxIntent []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_page.json
var googleDialogflowCxPage []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_playbook.json
var googleDialogflowCxPlaybook []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_security_settings.json
var googleDialogflowCxSecuritySettings []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_tool.json
var googleDialogflowCxTool []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_version.json
var googleDialogflowCxVersion []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_cx_webhook.json
var googleDialogflowCxWebhook []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_encryption_spec.json
var googleDialogflowEncryptionSpec []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_entity_type.json
var googleDialogflowEntityType []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_fulfillment.json
var googleDialogflowFulfillment []byte

//go:embed mapping/google/resource/dialogflow/google_dialogflow_intent.json
var googleDialogflowIntent []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone.json
var googleDNSManagedZone []byte

//go:embed mapping/google/resource/dns/google_dns_policy.json
var googleDNSPolicy []byte

//go:embed mapping/google/resource/dns/google_dns_record_set.json
var googleDNSRecordSet []byte

//go:embed mapping/google/resource/dns/google_dns_response_policy.json
var googleDNSResponsePolicy []byte

//go:embed mapping/google/resource/dns/google_dns_response_policy_rule.json
var googleDNSResponsePolicyRule []byte

//go:embed mapping/google/resource/edgecontainer/google_edgecontainer_cluster.json
var googleEdgecontainerCluster []byte

//go:embed mapping/google/resource/edgecontainer/google_edgecontainer_node_pool.json
var googleEdgecontainerNodePool []byte

//go:embed mapping/google/resource/edgecontainer/google_edgecontainer_vpn_connection.json
var googleEdgecontainerVpnConnection []byte

//go:embed mapping/google/resource/edgenetwork/google_edgenetwork_interconnect_attachment.json
var googleEdgenetworkInterconnectAttachment []byte

//go:embed mapping/google/resource/edgenetwork/google_edgenetwork_network.json
var googleEdgenetworkNetwork []byte

//go:embed mapping/google/resource/edgenetwork/google_edgenetwork_subnet.json
var googleEdgenetworkSubnet []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_binding.json
var googleEndpointsServiceConsumersIAMBinding []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_member.json
var googleEndpointsServiceConsumersIAMMember []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_consumers_iam_policy.json
var googleEndpointsServiceConsumersIAMPolicy []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_binding.json
var googleEndpointsServiceIAMBinding []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_member.json
var googleEndpointsServiceIAMMember []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service_iam_policy.json
var googleEndpointsServiceIAMPolicy []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_channel.json
var googleEventarcChannel []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_enrollment.json
var googleEventarcEnrollment []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_google_api_source.json
var googleEventarcGoogleAPISource []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_google_channel_config.json
var googleEventarcGoogleChannelConfig []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_message_bus.json
var googleEventarcMessageBus []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_pipeline.json
var googleEventarcPIPeline []byte

//go:embed mapping/google/resource/eventarc/google_eventarc_trigger.json
var googleEventarcTrigger []byte

//go:embed mapping/google/resource/firebase/google_firebase_android_app.json
var googleFirebaseAndroIDApp []byte

//go:embed mapping/google/resource/firebase/google_firebase_apple_app.json
var googleFirebaseAppleApp []byte

//go:embed mapping/google/resource/firebasedatabase/google_firebase_database_instance.json
var googleFirebaseDatabaseInstance []byte

//go:embed mapping/google/resource/firebasehosting/google_firebase_hosting_site.json
var googleFirebaseHostingSite []byte

//go:embed mapping/google/resource/firebase/google_firebase_project.json
var googleFirebaseProject []byte

//go:embed mapping/google/resource/firebasestorage/google_firebase_storage_bucket.json
var googleFirebaseStorageBucket []byte

//go:embed mapping/google/resource/firebase/google_firebase_web_app.json
var googleFirebaseWebApp []byte

//go:embed mapping/google/resource/firebaserules/google_firebaserules_release.json
var googleFirebaserulesRelease []byte

//go:embed mapping/google/resource/firebaserules/google_firebaserules_ruleset.json
var googleFirebaserulesRuleset []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_binding.json
var googleFolderIAMBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_member.json
var googleFolderIAMMember []byte

//go:embed mapping/google/resource/resourcemanager/google_folder_iam_policy.json
var googleFolderIAMPolicy []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_code_repository_index.json
var googleGeminiCodeRepositoryIndex []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_code_tools_setting.json
var googleGeminiCodeToolsSetting []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_code_tools_setting_binding.json
var googleGeminiCodeToolsSettingBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_data_sharing_with_google_setting.json
var googleGeminiDataSharingWithGoogleSetting []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_data_sharing_with_google_setting_binding.json
var googleGeminiDataSharingWithGoogleSettingBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_gemini_gcp_enablement_setting.json
var googleGeminiGeminiGcpEnablementSetting []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_gemini_gcp_enablement_setting_binding.json
var googleGeminiGeminiGcpEnablementSettingBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_logging_setting.json
var googleGeminiLoggingSetting []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_logging_setting_binding.json
var googleGeminiLoggingSettingBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_release_channel_setting.json
var googleGeminiReleaseChannelSetting []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_release_channel_setting_binding.json
var googleGeminiReleaseChannelSettingBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group.json
var googleGeminiRepositoryGroup []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_binding.json
var googleGeminiRepositoryGroupIAMBinding []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_member.json
var googleGeminiRepositoryGroupIAMMember []byte

//go:embed mapping/google/resource/cloudaicompanion/google_gemini_repository_group_iam_policy.json
var googleGeminiRepositoryGroupIAMPolicy []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_binding.json
var googleGkeBackupBackupPlanIAMBinding []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_member.json
var googleGkeBackupBackupPlanIAMMember []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_backup_plan_iam_policy.json
var googleGkeBackupBackupPlanIAMPolicy []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_binding.json
var googleGkeBackupRestorePlanIAMBinding []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_member.json
var googleGkeBackupRestorePlanIAMMember []byte

//go:embed mapping/google/resource/gkebackup/google_gke_backup_restore_plan_iam_policy.json
var googleGkeBackupRestorePlanIAMPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_binding.json
var googleGkeHubFeatureIAMBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_member.json
var googleGkeHubFeatureIAMMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_feature_iam_policy.json
var googleGkeHubFeatureIAMPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_binding.json
var googleGkeHubMembershIPIAMBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_member.json
var googleGkeHubMembershIPIAMMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_membership_iam_policy.json
var googleGkeHubMembershIPIAMPolicy []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_binding.json
var googleGkeHubScopeIAMBinding []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_member.json
var googleGkeHubScopeIAMMember []byte

//go:embed mapping/google/resource/gkehub/google_gke_hub_scope_iam_policy.json
var googleGkeHubScopeIAMPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_binding.json
var googleHealthcareConsentStoreIAMBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_member.json
var googleHealthcareConsentStoreIAMMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store_iam_policy.json
var googleHealthcareConsentStoreIAMPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_binding.json
var googleHealthcareDatasetIAMBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_member.json
var googleHealthcareDatasetIAMMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset_iam_policy.json
var googleHealthcareDatasetIAMPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_binding.json
var googleHealthcareDicomStoreIAMBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_member.json
var googleHealthcareDicomStoreIAMMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store_iam_policy.json
var googleHealthcareDicomStoreIAMPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_binding.json
var googleHealthcareFhirStoreIAMBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_member.json
var googleHealthcareFhirStoreIAMMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store_iam_policy.json
var googleHealthcareFhirStoreIAMPolicy []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_binding.json
var googleHealthcareHl7V2StoreIAMBinding []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_member.json
var googleHealthcareHl7V2StoreIAMMember []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store_iam_policy.json
var googleHealthcareHl7V2StoreIAMPolicy []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool.json
var googleIAMWorkforcePool []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool_iam_binding.json
var googleIAMWorkforcePoolIAMBinding []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool_iam_member.json
var googleIAMWorkforcePoolIAMMember []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool_iam_policy.json
var googleIAMWorkforcePoolIAMPolicy []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool_provider.json
var googleIAMWorkforcePoolProvIDer []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workforce_pool_provider_key.json
var googleIAMWorkforcePoolProvIDerKey []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool.json
var googleIAMWorkloadIdentityPool []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_binding.json
var googleIAMWorkloadIdentityPoolIAMBinding []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_member.json
var googleIAMWorkloadIdentityPoolIAMMember []byte

//go:embed mapping/google/resource/iam.googleapis.com/google_iam_workload_identity_pool_iam_policy.json
var googleIAMWorkloadIdentityPoolIAMPolicy []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool_provider.json
var googleIAMWorkloadIdentityPoolProvIDer []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_binding.json
var googleIapAppEngineServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_member.json
var googleIapAppEngineServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_service_iam_policy.json
var googleIapAppEngineServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_binding.json
var googleIapAppEngineVersionIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_member.json
var googleIapAppEngineVersionIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_app_engine_version_iam_policy.json
var googleIapAppEngineVersionIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_settings.json
var googleIapSettings []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group.json
var googleIapTunnelDestGroup []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_binding.json
var googleIapTunnelDestGroupIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_member.json
var googleIapTunnelDestGroupIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_dest_group_iam_policy.json
var googleIapTunnelDestGroupIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_binding.json
var googleIapTunnelIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_member.json
var googleIapTunnelIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_iam_policy.json
var googleIapTunnelIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_binding.json
var googleIapTunnelInstanceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_member.json
var googleIapTunnelInstanceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_tunnel_instance_iam_policy.json
var googleIapTunnelInstanceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_binding.json
var googleIapWebBackendServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_member.json
var googleIapWebBackendServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_backend_service_iam_policy.json
var googleIapWebBackendServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_binding.json
var googleIapWebCloudRunServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_member.json
var googleIapWebCloudRunServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_cloud_run_service_iam_policy.json
var googleIapWebCloudRunServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_forwarding_rule_service_iam_binding.json
var googleIapWebForwardingRuleServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_forwarding_rule_service_iam_member.json
var googleIapWebForwardingRuleServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_forwarding_rule_service_iam_policy.json
var googleIapWebForwardingRuleServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_binding.json
var googleIapWebIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_member.json
var googleIapWebIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_iam_policy.json
var googleIapWebIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_binding.json
var googleIapWebRegionBackendServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_member.json
var googleIapWebRegionBackendServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_backend_service_iam_policy.json
var googleIapWebRegionBackendServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_forwarding_rule_service_iam_binding.json
var googleIapWebRegionForwardingRuleServiceIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_forwarding_rule_service_iam_member.json
var googleIapWebRegionForwardingRuleServiceIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_region_forwarding_rule_service_iam_policy.json
var googleIapWebRegionForwardingRuleServiceIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_binding.json
var googleIapWebTypeAppEngineIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_member.json
var googleIapWebTypeAppEngineIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_app_engine_iam_policy.json
var googleIapWebTypeAppEngineIAMPolicy []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_binding.json
var googleIapWebTypeComputeIAMBinding []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_member.json
var googleIapWebTypeComputeIAMMember []byte

//go:embed mapping/google/resource/iap/google_iap_web_type_compute_iam_policy.json
var googleIapWebTypeComputeIAMPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key.json
var googleKMSCryptoKey []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_binding.json
var googleKMSCryptoKeyIAMBinding []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_member.json
var googleKMSCryptoKeyIAMMember []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_iam_policy.json
var googleKMSCryptoKeyIAMPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_crypto_key_version.json
var googleKMSCryptoKeyVersion []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_binding.json
var googleKMSEkmConnectionIAMBinding []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_member.json
var googleKMSEkmConnectionIAMMember []byte

//go:embed mapping/google/resource/kms/google_kms_ekm_connection_iam_policy.json
var googleKMSEkmConnectionIAMPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_handle.json
var googleKMSKeyHandle []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring.json
var googleKMSKeyRing []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_binding.json
var googleKMSKeyRingIAMBinding []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_member.json
var googleKMSKeyRingIAMMember []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_iam_policy.json
var googleKMSKeyRingIAMPolicy []byte

//go:embed mapping/google/resource/cloudkms/google_kms_key_ring_import_job.json
var googleKMSKeyRingImportJob []byte

//go:embed mapping/google/resource/cloudkms/google_kms_secret_ciphertext.json
var googleKMSSecretCIPhertext []byte

//go:embed mapping/google/resource/logging/google_logging_billing_account_exclusion.json
var googleLoggingBillingAccountExclusion []byte

//go:embed mapping/google/resource/logging/google_logging_billing_account_sink.json
var googleLoggingBillingAccountSink []byte

//go:embed mapping/google/resource/logging/google_logging_folder_exclusion.json
var googleLoggingFolderExclusion []byte

//go:embed mapping/google/resource/logging/google_logging_folder_settings.json
var googleLoggingFolderSettings []byte

//go:embed mapping/google/resource/logging/google_logging_folder_sink.json
var googleLoggingFolderSink []byte

//go:embed mapping/google/resource/logging/google_logging_linked_dataset.json
var googleLoggingLinkedDataset []byte

//go:embed mapping/google/resource/logging/google_logging_log_scope.json
var googleLoggingLogScope []byte

//go:embed mapping/google/resource/logging/google_logging_log_view.json
var googleLoggingLogView []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_binding.json
var googleLoggingLogViewIAMBinding []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_member.json
var googleLoggingLogViewIAMMember []byte

//go:embed mapping/google/resource/logging/google_logging_log_view_iam_policy.json
var googleLoggingLogViewIAMPolicy []byte

//go:embed mapping/google/resource/logging/google_logging_metric.json
var googleLoggingMetric []byte

//go:embed mapping/google/resource/logging/google_logging_organization_exclusion.json
var googleLoggingOrganizationExclusion []byte

//go:embed mapping/google/resource/logging/google_logging_organization_settings.json
var googleLoggingOrganizationSettings []byte

//go:embed mapping/google/resource/logging/google_logging_organization_sink.json
var googleLoggingOrganizationSink []byte

//go:embed mapping/google/resource/logging/google_logging_project_exclusion.json
var googleLoggingProjectExclusion []byte

//go:embed mapping/google/resource/logging/google_logging_project_sink.json
var googleLoggingProjectSink []byte

//go:embed mapping/google/resource/managedkafka/google_managed_kafka_acl.json
var googleManagedKafkaAcl []byte

//go:embed mapping/google/resource/managedkafka/google_managed_kafka_cluster.json
var googleManagedKafkaCluster []byte

//go:embed mapping/google/resource/managedkafka/google_managed_kafka_connect_cluster.json
var googleManagedKafkaConnectCluster []byte

//go:embed mapping/google/resource/managedkafka/google_managed_kafka_connector.json
var googleManagedKafkaConnector []byte

//go:embed mapping/google/resource/managedkafka/google_managed_kafka_topic.json
var googleManagedKafkaTopic []byte

//go:embed mapping/google/resource/memcache/google_memcache_instance.json
var googleMemcacheInstance []byte

//go:embed mapping/google/resource/memorystore/google_memorystore_instance.json
var googleMemorystoreInstance []byte

//go:embed mapping/google/resource/modelarmor/google_model_armor_floorsetting.json
var googleModelArmorFloorsetting []byte

//go:embed mapping/google/resource/modelarmor/google_model_armor_template.json
var googleModelArmorTemplate []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_alert_policy.json
var googleMonitoringAlertPolicy []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_custom_service.json
var googleMonitoringCustomService []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_dashboard.json
var googleMonitoringDashboard []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_group.json
var googleMonitoringGroup []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_metric_descriptor.json
var googleMonitoringMetricDescrIPtor []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_monitored_project.json
var googleMonitoringMonitoredProject []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_notification_channel.json
var googleMonitoringNotificationChannel []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_service.json
var googleMonitoringService []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_slo.json
var googleMonitoringSlo []byte

//go:embed mapping/google/resource/monitoring/google_monitoring_uptime_check_config.json
var googleMonitoringUptimeCheckConfig []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_binding.json
var googleNetworkSecurityAddressGroupIAMBinding []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_member.json
var googleNetworkSecurityAddressGroupIAMMember []byte

//go:embed mapping/google/resource/networksecurity/google_network_security_address_group_iam_policy.json
var googleNetworkSecurityAddressGroupIAMPolicy []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_environment.json
var googleNotebooksEnvironment []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance.json
var googleNotebooksInstance []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_binding.json
var googleNotebooksInstanceIAMBinding []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_member.json
var googleNotebooksInstanceIAMMember []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_instance_iam_policy.json
var googleNotebooksInstanceIAMPolicy []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime.json
var googleNotebooksRuntime []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_binding.json
var googleNotebooksRuntimeIAMBinding []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_member.json
var googleNotebooksRuntimeIAMMember []byte

//go:embed mapping/google/resource/notebooks/google_notebooks_runtime_iam_policy.json
var googleNotebooksRuntimeIAMPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_binding.json
var googleOrganizationIAMBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_member.json
var googleOrganizationIAMMember []byte

//go:embed mapping/google/resource/resourcemanager/google_organization_iam_policy.json
var googleOrganizationIAMPolicy []byte

//go:embed mapping/google/resource/osconfig/google_os_config_guest_policies.json
var googleOsConfigGuestPolicies []byte

//go:embed mapping/google/resource/osconfig/google_os_config_os_policy_assignment.json
var googleOsConfigOsPolicyAssignment []byte

//go:embed mapping/google/resource/osconfig/google_os_config_patch_deployment.json
var googleOsConfigPatchDeployment []byte

//go:embed mapping/google/resource/osconfig/google_os_config_v2_policy_orchestrator.json
var googleOsConfigV2PolicyOrchestrator []byte

//go:embed mapping/google/resource/osconfig/google_os_config_v2_policy_orchestrator_for_folder.json
var googleOsConfigV2PolicyOrchestratorForFolder []byte

//go:embed mapping/google/resource/osconfig/google_os_config_v2_policy_orchestrator_for_organization.json
var googleOsConfigV2PolicyOrchestratorForOrganization []byte

//go:embed mapping/google/resource/parallelstore/google_parallelstore_instance.json
var googleParallelstoreInstance []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool.json
var googlePrivatecaCaPool []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_binding.json
var googlePrivatecaCaPoolIAMBinding []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_member.json
var googlePrivatecaCaPoolIAMMember []byte

//go:embed mapping/google/resource/privateca/google_privateca_ca_pool_iam_policy.json
var googlePrivatecaCaPoolIAMPolicy []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate.json
var googlePrivatecaCertificate []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_authority.json
var googlePrivatecaCertificateAuthority []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template.json
var googlePrivatecaCertificateTemplate []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_binding.json
var googlePrivatecaCertificateTemplateIAMBinding []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_member.json
var googlePrivatecaCertificateTemplateIAMMember []byte

//go:embed mapping/google/resource/privateca/google_privateca_certificate_template_iam_policy.json
var googlePrivatecaCertificateTemplateIAMPolicy []byte

//go:embed mapping/google/resource/privilegedaccessmanager/google_privileged_access_manager_entitlement.json
var googlePrivilegedAccessManagerEntitlement []byte

//go:embed mapping/google/resource/resourcemanager/google_project.json
var googleProject []byte

//go:embed mapping/google/resource/accessapproval/google_project_access_approval_settings.json
var googleProjectAccessApprovalSettings []byte

//go:embed mapping/google/resource/iam/google_project_default_service_accounts.json
var googleProjectDefaultServiceAccounts []byte

//go:embed mapping/google/resource/iam/google_project_iam_audit_config.json
var googleProjectIAMAuditConfig []byte

//go:embed mapping/google/resource/resourcemanager/google_project_iam_binding.json
var googleProjectIAMBinding []byte

//go:embed mapping/google/resource/iam/google_project_iam_custom_role.json
var googleProjectIAMCustomRole []byte

//go:embed mapping/google/resource/resourcemanager/google_project_iam_member_remove.json
var googleProjectIAMMemberRemove []byte

//go:embed mapping/google/resource/orgpolicy/google_project_organization_policy.json
var googleProjectOrganizationPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_project_service.json
var googleProjectService []byte

//go:embed mapping/google/resource/compute/google_project_usage_export_bucket.json
var googleProjectUsageExportBucket []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_reservation.json
var googlePubsubLiteReservation []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_subscription.json
var googlePubsubLiteSubscrIPtion []byte

//go:embed mapping/google/resource/pubsublite/google_pubsub_lite_topic.json
var googlePubsubLiteTopic []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema.json
var googlePubsubSchema []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_binding.json
var googlePubsubSchemaIAMBinding []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_member.json
var googlePubsubSchemaIAMMember []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_schema_iam_policy.json
var googlePubsubSchemaIAMPolicy []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription.json
var googlePubsubSubscrIPtion []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_binding.json
var googlePubsubSubscrIPtionIAMBinding []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_member.json
var googlePubsubSubscrIPtionIAMMember []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_subscription_iam_policy.json
var googlePubsubSubscrIPtionIAMPolicy []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic.json
var googlePubsubTopic []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic_iam_binding.json
var googlePubsubTopicIAMBinding []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic_iam_member.json
var googlePubsubTopicIAMMember []byte

//go:embed mapping/google/resource/pubsub/google_pubsub_topic_iam_policy.json
var googlePubsubTopicIAMPolicy []byte

//go:embed mapping/google/resource/redis/google_redis_cluster.json
var googleRedisCluster []byte

//go:embed mapping/google/resource/redis/google_redis_cluster_user_created_connections.json
var googleRedisClusterUserCreatedConnections []byte

//go:embed mapping/google/resource/redis/google_redis_instance.json
var googleRedisInstance []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config.json
var googleRuntimeconfigConfig []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_binding.json
var googleRuntimeconfigConfigIAMBinding []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_member.json
var googleRuntimeconfigConfigIAMMember []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_config_iam_policy.json
var googleRuntimeconfigConfigIAMPolicy []byte

//go:embed mapping/google/resource/runtimeconfig/google_runtimeconfig_variable.json
var googleRuntimeconfigVariable []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_binding.json
var googleSccSourceIAMBinding []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_member.json
var googleSccSourceIAMMember []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source_iam_policy.json
var googleSccSourceIAMPolicy []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_binding.json
var googleSccV2OrganizationSourceIAMBinding []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_member.json
var googleSccV2OrganizationSourceIAMMember []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_source_iam_policy.json
var googleSccV2OrganizationSourceIAMPolicy []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret.json
var googleSecretManagerRegionalSecret []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_binding.json
var googleSecretManagerRegionalSecretIAMBinding []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_member.json
var googleSecretManagerRegionalSecretIAMMember []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_iam_policy.json
var googleSecretManagerRegionalSecretIAMPolicy []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_regional_secret_version.json
var googleSecretManagerRegionalSecretVersion []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret.json
var googleSecretManagerSecret []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_iam_binding.json
var googleSecretManagerSecretIAMBinding []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_iam_policy.json
var googleSecretManagerSecretIAMPolicy []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_iam_member.json
var googleSecretManagerSecretIAMMember []byte

//go:embed mapping/google/resource/secretmanager/google_secret_manager_secret_version.json
var googleSecretManagerSecretVersion []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_branch_rule.json
var googleSecureSourceManagerBranchRule []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance.json
var googleSecureSourceManagerInstance []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_binding.json
var googleSecureSourceManagerInstanceIAMBinding []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_member.json
var googleSecureSourceManagerInstanceIAMMember []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_instance_iam_policy.json
var googleSecureSourceManagerInstanceIAMPolicy []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository.json
var googleSecureSourceManagerRepository []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_binding.json
var googleSecureSourceManagerRepositoryIAMBinding []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_member.json
var googleSecureSourceManagerRepositoryIAMMember []byte

//go:embed mapping/google/resource/securesourcemanager/google_secure_source_manager_repository_iam_policy.json
var googleSecureSourceManagerRepositoryIAMPolicy []byte

//go:embed mapping/google/resource/iam/google_service_account.json
var googleServiceAccount []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_binding.json
var googleServiceAccountIAMBinding []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_member.json
var googleServiceAccountIAMMember []byte

//go:embed mapping/google/resource/iam/google_service_account_iam_policy.json
var googleServiceAccountIAMPolicy []byte

//go:embed mapping/google/resource/iam/google_service_account_key.json
var googleServiceAccountKey []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_endpoint.json
var googleServiceDirectoryEndpoint []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace.json
var googleServiceDirectoryNamespace []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_binding.json
var googleServiceDirectoryNamespaceIAMBinding []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_member.json
var googleServiceDirectoryNamespaceIAMMember []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_namespace_iam_policy.json
var googleServiceDirectoryNamespaceIAMPolicy []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service.json
var googleServiceDirectoryService []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_binding.json
var googleServiceDirectoryServiceIAMBinding []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_member.json
var googleServiceDirectoryServiceIAMMember []byte

//go:embed mapping/google/resource/servicedirectory/google_service_directory_service_iam_policy.json
var googleServiceDirectoryServiceIAMPolicy []byte

//go:embed mapping/google/resource/servicenetworking/google_service_networking_connection.json
var googleServiceNetworkingConnection []byte

//go:embed mapping/google/resource/source/google_soucerepo_repository.json
var googleSoucerepoRepository []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_binding.json
var googleSourcerepoRepositoryIAMBinding []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_member.json
var googleSourcerepoRepositoryIAMMember []byte

//go:embed mapping/google/resource/source/google_sourcerepo_repository_iam_policy.json
var googleSourcerepoRepositoryIAMPolicy []byte

//go:embed mapping/google/resource/spanner/google_spanner_backup_schedule.json
var googleSpannerBackupSchedule []byte

//go:embed mapping/google/resource/spanner/google_spanner_database.json
var googleSpannerDatabase []byte

//go:embed mapping/google/resource/spanner/google_spanner_database_iam.json
var googleSpannerDatabaseIAM []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance.json
var googleSpannerInstance []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_config.json
var googleSpannerInstanceConfig []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_iam.json
var googleSpannerInstanceIAM []byte

//go:embed mapping/google/resource/spanner/google_spanner_instance_partition.json
var googleSpannerInstancePartition []byte

//go:embed mapping/google/resource/cloudsql/google_sql_database.json
var googleSQLDatabase []byte

//go:embed mapping/google/resource/cloudsql/google_sql_database_instance.json
var googleSQLDatabaseInstance []byte

//go:embed mapping/google/resource/cloudsql/google_sql_ssl_cert.json
var googleSQLSslCert []byte

//go:embed mapping/google/resource/cloudsql/google_sql_user.json
var googleSQLUser []byte

//go:embed mapping/google/resource/storage/google_storage_bucket.json
var googleStorageBucket []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_access_control.json
var googleStorageBucketAccessControl []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_acl.json
var googleStorageBucketAcl []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_binding.json
var googleStorageBucketIAMBinding []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_member.json
var googleStorageBucketIAMMember []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_iam_policy.json
var googleStorageBucketIAMPolicy []byte

//go:embed mapping/google/resource/storage/google_storage_bucket_object.json
var googleStorageBucketObject []byte

//go:embed mapping/google/resource/storage/google_storage_control_folder_intelligence_config.json
var googleStorageControlFolderIntelligenceConfig []byte

//go:embed mapping/google/resource/storage/google_storage_control_organization_intelligence_config.json
var googleStorageControlOrganizationIntelligenceConfig []byte

//go:embed mapping/google/resource/storage/google_storage_control_project_intelligence_config.json
var googleStorageControlProjectIntelligenceConfig []byte

//go:embed mapping/google/resource/storage/google_storage_default_object_access_control.json
var googleStorageDefaultObjectAccessControl []byte

//go:embed mapping/google/resource/storage/google_storage_default_object_acl.json
var googleStorageDefaultObjectAcl []byte

//go:embed mapping/google/resource/storage/google_storage_hmac_key.json
var googleStorageHmacKey []byte

//go:embed mapping/google/resource/storageinsights/google_storage_insights_dataset_config.json
var googleStorageInsightsDatasetConfig []byte

//go:embed mapping/google/resource/storage/google_storage_insights_report_config.json
var googleStorageInsightsReportConfig []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_binding.json
var googleStorageManagedFolderIAMBinding []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_member.json
var googleStorageManagedFolderIAMMember []byte

//go:embed mapping/google/resource/storage/google_storage_managed_folder_iam_policy.json
var googleStorageManagedFolderIAMPolicy []byte

//go:embed mapping/google/resource/storage/google_storage_object_access_control.json
var googleStorageObjectAccessControl []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_location_tag_binding.json
var googleTagsLocationTagBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_binding.json
var googleTagsTagBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key.json
var googleTagsTagKey []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_binding.json
var googleTagsTagKeyIAMBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_member.json
var googleTagsTagKeyIAMMember []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_key_iam_policy.json
var googleTagsTagKeyIAMPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value.json
var googleTagsTagValue []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_binding.json
var googleTagsTagValueIAMBinding []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_member.json
var googleTagsTagValueIAMMember []byte

//go:embed mapping/google/resource/resourcemanager/google_tags_tag_value_iam_policy.json
var googleTagsTagValueIAMPolicy []byte

//go:embed mapping/google/resource/tpu/google_tpu_node.json
var googleTpuNode []byte

//go:embed mapping/google/resource/tpu/google_tpu_v2_queued_resource.json
var googleTpuV2QueuedResource []byte

//go:embed mapping/google/resource/tpu/google_tpu_v2_vm.json
var googleTpuV2Vm []byte

//go:embed mapping/google/resource/transcoder/google_transcoder_job.json
var googleTranscoderJob []byte

//go:embed mapping/google/resource/transcoder/google_transcoder_job_template.json
var googleTranscoderJobTemplate []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_cache_config.json
var googleVertexAiCacheConfig []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_dataset.json
var googleVertexAIDataset []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_deployment_resource_pool.json
var googleVertexAIDeploymentResourcePool []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_endpoint.json
var googleVertexAiEndpoint []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_endpoint_iam.json
var googleVertexAiEndpointIAM []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_endpoint_with_model_garden_deployment.json
var googleVertexAiEndpointWithModelGardenDeployment []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group.json
var googleVertexAiFeatureGroup []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_feature.json
var googleVertexAiFeatureGroupFeature []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_binding.json
var googleVertexAiFeatureGroupIAMBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_member.json
var googleVertexAiFeatureGroupIAMMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_group_iam_policy.json
var googleVertexAiFeatureGroupIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store.json
var googleVertexAiFeatureOnlineStore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview.json
var googleVertexAiFeatureOnlineStoreFeatureview []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_binding.json
var googleVertexAiFeatureOnlineStoreFeatureviewIAMBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_member.json
var googleVertexAiFeatureOnlineStoreFeatureviewIAMMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_featureview_iam_policy.json
var googleVertexAiFeatureOnlineStoreFeatureviewIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_binding.json
var googleVertexAiFeatureOnlineStoreIAMBinding []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_member.json
var googleVertexAiFeatureOnlineStoreIAMMember []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_feature_online_store_iam_policy.json
var googleVertexAiFeatureOnlineStoreIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore.json
var googleVertexAiFeaturestore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype.json
var googleVertexAiFeaturestoreEntitytype []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype_feature.json
var googleVertexAiFeaturestoreEntitytypeFeature []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype_iam.json
var googleVertexAiFeaturestoreEntitytypeIAM []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_entitytype_iam_policy.json
var googleVertexAiFeaturestoreEntitytypeIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_iam.json
var googleVertexAiFeaturestoreIAM []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_featurestore_iam_policy.json
var googleVertexAiFeaturestoreIAMPolicy []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_index.json
var googleVertexAiIndex []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_index_endpoint.json
var googleVertexAiIndexEndpoint []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_index_endpoint_deployed_index.json
var googleVertexAiIndexEndpointDeployedIndex []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_metadata_store.json
var googleVertexAiMetadataStore []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_rag_engine_config.json
var googleVertexAiRagEngineConfig []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_reasoning_engine.json
var googleVertexAiReasoningEngine []byte

//go:embed mapping/google/resource/aiplatform/google_vertex_ai_tensorboard.json
var googleVertexAiTensorboard []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_cluster.json
var googleVmwareengineCluster []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_external_access_rule.json
var googleVmwareengineExternalAccessRule []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_external_address.json
var googleVmwareengineExternalAddress []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_network.json
var googleVmwareengineNetwork []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_network_peering.json
var googleVmwareengineNetworkPeering []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_network_policy.json
var googleVmwareengineNetworkPolicy []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_private_cloud.json
var googleVmwareenginePrivateCloud []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_subnet.json
var googleVmwareengineSubnet []byte

//go:embed mapping/google/resource/vpcaccess/google_vpc_access_connector.json
var googleVPCAccessConnector []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance.json
var googleWorkbenchInstance []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_binding.json
var googleWorkbenchInstanceIAMBinding []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_member.json
var googleWorkbenchInstanceIAMMember []byte

//go:embed mapping/google/resource/notebooks/google_workbench_instance_iam_policy.json
var googleWorkbenchInstanceIAMPolicy []byte

//go:embed mapping/google/resource/workflows/google_workflows_workflow.json
var googleWorkflowsWorkflow []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation.json
var googleWorkstationsWorkstation []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_cluster.json
var googleWorkstationsWorkstationCluster []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config.json
var googleWorkstationsWorkstationConfig []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_binding.json
var googleWorkstationsWorkstationConfigIAMBinding []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_member.json
var googleWorkstationsWorkstationConfigIAMMember []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_config_iam_policy.json
var googleWorkstationsWorkstationConfigIAMPolicy []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_binding.json
var googleWorkstationsWorkstationIAMBinding []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_member.json
var googleWorkstationsWorkstationIAMMember []byte

//go:embed mapping/google/resource/workstations/google_workstations_workstation_iam_policy.json
var googleWorkstationsWorkstationIAMPolicy []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_binding.json
var googleDNSManagedZoneIAMBinding []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_member.json
var googleDNSManagedZoneIAMMember []byte

//go:embed mapping/google/resource/dns/google_dns_managed_zone_iam_policy.json
var googleDNSManagedZoneIAMPolicy []byte

//go:embed mapping/google/resource/resourcemanager/google_project_iam_policy.json
var googleProjectIAMPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_access_level_condition.json
var googleAccessContextManagerAccessLevelCondition []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_dry_run_egress_policy.json
var googleAccessContextManagerServicePerimeterDryRunEgressPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_dry_run_ingress_policy.json
var googleAccessContextManagerServicePerimeterDryRunIngressPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_dry_run_resource.json
var googleAccessContextManagerServicePerimeterDryRunResource []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_egress_policy.json
var googleAccessContextManagerServicePerimeterEgressPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_ingress_policy.json
var googleAccessContextManagerServicePerimeterIngressPolicy []byte

//go:embed mapping/google/resource/accesscontextmanager/google_access_context_manager_service_perimeter_resource.json
var googleAccessContextManagerServicePerimeterResource []byte

//go:embed mapping/google/resource/managedidentities/google_active_directory_domain.json
var googleActiveDirectoryDomain []byte

//go:embed mapping/google/resource/managedidentities/google_active_directory_domain_trust.json
var googleActiveDirectoryDomainTrust []byte

//go:embed mapping/google/resource/managedidentities/google_active_directory_peering.json
var googleActiveDirectoryPeering []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_certificate.json
var googleCertificateManagerCertificate []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_certificate_issuance_config.json
var googleCertificateManagerCertificateIssuanceConfig []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_certificate_map.json
var googleCertificateManagerCertificxateMap []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_certificate_map_entry.json
var googleCertificateManagerCertificateMapEntry []byte

//go:embed mapping/google/resource/certificatemanager/google_certificate_manager_trust_config.json
var googleCertificateManagerTrustConfig []byte

//go:embed mapping/google/resource/datafusion/google_data_fusion_instance.json
var googleDataFusionInstance []byte

//go:embed mapping/google/resource/servicemanagement/google_endpoints_service.json
var googleEndpointsService []byte

//go:embed mapping/google/resource/essentialcontacts/google_essential_contacts_contact.json
var googleEssentialContactsContact []byte

//go:embed mapping/google/resource/iam/google_iam_access_boundary_policy.json
var googleIamAccessBoundaryPolicy []byte

//go:embed mapping/google/resource/iam/google_iam_deny_policy.json
var googleIamDenyPolicy []byte

//go:embed mapping/google/resource/iam/google_iam_oauth_client.json
var googleIamOauthClient []byte

//go:embed mapping/google/resource/iam/google_iam_oauth_client_credential.json
var googleIamOauthClientCredential []byte

//go:embed mapping/google/resource/iam/google_iam_principal_access_boundary_policy.json
var googleIamPrincipalAccessBoundaryPolicy []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool_managed_identity.json
var googleIamWorkloadIdentityPoolManagedIdentity []byte

//go:embed mapping/google/resource/iam/google_iam_workload_identity_pool_namespace.json
var googleIamWorkloadIdentityPoolNamespace []byte

//go:embed mapping/google/resource/compute/google_compute_route.json
var googleComputeRoute []byte

//go:embed mapping/google/resource/compute/google_compute_router_nat_address.json
var googleComputeRouterNatAddress []byte

//go:embed mapping/google/resource/compute/google_compute_router_peer.json
var googleComputeRouterPeer []byte

//go:embed mapping/google/resource/compute/google_compute_router_route_policy.json
var googleComputeRouterRoutePolicy []byte

//go:embed mapping/google/resource/compute/google_compute_router_interface.json
var googleComputeRouterInterface []byte

//go:embed mapping/google/resource/file/google_filestore_backup.json
var googleFilestoreBackup []byte

//go:embed mapping/google/resource/file/google_filestore_instance.json
var googleFilestoreInstance []byte

//go:embed mapping/google/resource/file/google_filestore_snapshot.json
var googleFilestoreSnapshot []byte

//go:embed mapping/google/resource/datastore/google_firestore_backup_schedule.json
var googleFirestoreBackupSchedule []byte

//go:embed mapping/google/resource/datastore/google_firestore_database.json
var googleFirestoreDatabase []byte

//go:embed mapping/google/resource/datastore/google_firestore_document.json
var googleFirestoreDocument []byte

//go:embed mapping/google/resource/datastore/google_firestore_field.json
var googleFirestoreField []byte

//go:embed mapping/google/resource/datastore/google_firestore_index.json
var googleFirestoreIndex []byte

//go:embed mapping/google/resource/datastore/google_firestore_user_creds.json
var googleFirestoreUserCreds []byte

//go:embed mapping/google/resource/compute/google_compute_image.json
var googleComputeImage []byte

//go:embed mapping/google/resource/compute/google_compute_disk.json
var googleComputeDisk []byte

//go:embed mapping/google/resource/compute/google_compute_machine_image.json
var googleComputeMachineImage []byte

//go:embed mapping/google/resource/netapp/google_netapp_backup_policy.json
var googleNetappBackupPolicy []byte

//go:embed mapping/google/resource/observability/google_observability_trace_scope.json
var googleObservabilityTraceScope []byte

//go:embed mapping/google/resource/cloudsql/google_sql_source_representation_instance.json
var googleSqlSourceRepresentationInstance []byte

//go:embed mapping/google/resource/biglake/google_biglake_iceberg_catalog.json
var googleBiglakeIcebergCatalog []byte

//go:embed mapping/google/resource/biglake/google_biglake_iceberg_catalog_iam.json
var googleBiglakeIcebergCatalogIam []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_datapolicyv2_data_policy.json
var googleBigqueryDatapolicyv2DataPolicy []byte

//go:embed mapping/google/resource/bigquery/google_bigquery_datapolicyv2_data_policy_iam.json
var googleBigqueryDatapolicyv2DataPolicyIam []byte

//go:embed mapping/google/resource/apphub/google_apphub_boundary.json
var googleApphubBoundary []byte

//go:embed mapping/google/resource/artifactregistry/google_artifact_registry_vpcsc_config.json
var googleArtifactRegistryVpcscConfig []byte

//go:embed mapping/google/resource/assuredworkloads/google_assured_workloads_workload.json
var googleAssuredWorkloadsWorkload []byte

//go:embed mapping/google/resource/backupdr/google_backup_dr_restore_workload.json
var googleBackupDrRestoreWorkload []byte

//go:embed mapping/google/resource/billing/google_billing_subaccount.json
var googleBillingSubaccount []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_attestor.json
var googleBinaryAuthorizationAttestor []byte

//go:embed mapping/google/resource/containeranalysis/google_container_analysis_note.json
var googleContainerAnalysisNote []byte

//go:embed mapping/google/resource/binaryauthorization/google_binary_authorization_policy.json
var googleBinaryAuthorizationPolicy []byte

//go:embed mapping/google/resource/blockchainnodeengine/google_blockchain_node_engine_blockchain_nodes.json
var googleBlockchainNodeEngineBlockchainNodes []byte

//go:embed mapping/google/resource/looker/google_looker_instance.json
var googleLookerInstance []byte

//go:embed mapping/google/resource/lustre/google_lustre_instance.json
var googleLustreInstance []byte

//go:embed mapping/google/resource/privilegedaccessmanager/google_privileged_access_manager_settings.json
var googlePrivilegedAccessManagerSettings []byte

//go:embed mapping/google/resource/publicca/google_public_ca_external_account_key.json
var googlePublicCaExternalAccountKey []byte

//go:embed mapping/google/resource/recaptchaenterprise/google_recaptcha_enterprise_key.json
var googleRecaptchaEnterpriseKey []byte

//go:embed mapping/google/resource/developerconnect/google_developer_connect_insights_config.json
var googleDeveloperConnectInsightsConfig []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_consent_store.json
var googleHealthcareConsentStore []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dataset.json
var googleHealthcareDataset []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_dicom_store.json
var googleHealthcareDicomStore []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_fhir_store.json
var googleHealthcareFhirStore []byte

//go:embed mapping/google/resource/healthcare/google_healthcare_hl7_v2_store.json
var googleHealthcareHl7V2Store []byte

//go:embed mapping/google/resource/cloudkms/google_kms_ekm_connection.json
var googleKmsEkmConnection []byte

//go:embed mapping/google/resource/cloudkms/google_kms_folder_kaj_policy_config.json
var googleKmsFolderKajPolicyConfig []byte

//go:embed mapping/google/resource/cloudkms/google_kms_organization_kaj_policy_config.json
var googleKmsOrganizationKajPolicyConfig []byte

//go:embed mapping/google/resource/cloudkms/google_kms_project_kaj_policy_config.json
var googleKmsProjectKajPolicyConfig []byte

//go:embed mapping/google/resource/memorystore/google_memorystore_instance_desired_user_created_endpoints.json
var googleMemorystoreInstanceDesiredUserCreatedEndpoints []byte

//go:embed mapping/google/resource/migrationcenter/google_migration_center_group.json
var googleMigrationCenterGroup []byte

//go:embed mapping/google/resource/migrationcenter/google_migration_center_preference_set.json
var googleMigrationCenterPreferenceSet []byte

//go:embed mapping/google/resource/apigee/google_apigee_addons_config.json
var GoogleApigeeAddonsConfig []byte

//go:embed mapping/google/resource/apigee/google_apigee_api_deployment.json
var GoogleApigeeApiDeployment []byte

//go:embed mapping/google/resource/apigee/google_apigee_api_product.json
var GoogleApigeeApiProduct []byte

//go:embed mapping/google/resource/apigee/google_apigee_app_group.json
var GoogleApigeeAppGroup []byte

//go:embed mapping/google/resource/apigee/google_apigee_developer.json
var googleApigeeDeveloper []byte

//go:embed mapping/google/resource/apigee/google_apigee_developer_app.json
var googleApigeeDeveloperApp []byte

//go:embed mapping/google/resource/apigee/google_apigee_dns_zone.json
var googleApigeeDnsZone []byte

//go:embed mapping/google/resource/apigee/google_apigee_envgroup.json
var googleApigeeEnvgroup []byte

//go:embed mapping/google/resource/apigee/google_apigee_envgroup_attachment.json
var GoogleApigeeEnvgroupAttachment []byte

//go:embed mapping/google/resource/apigee/google_apigee_instance.json
var googleApigeeInstance []byte

//go:embed mapping/google/resource/apigee/google_apigee_instance_attachment.json
var googleApigeeInstanceAttachment []byte

//go:embed mapping/google/resource/apigee/google_apigee_nat_address.json
var googleApigeeNatAddress []byte

//go:embed mapping/google/resource/apigee/google_apigee_organization.json
var googleApigeeOrganization []byte

//go:embed mapping/google/resource/apigee/google_apigee_security_action.json
var googleApigeeSecurityAction []byte

//go:embed mapping/google/resource/apigee/google_apigee_security_feedback.json
var googleApigeeSecurityFeedback []byte

//go:embed mapping/google/resource/apigee/google_apigee_security_monitoring_condition.json
var googleApigeeSecurityMonitoringCondition []byte

//go:embed mapping/google/resource/apigee/google_apigee_security_profile_v2.json
var googleApigeeSecurityProfileV2 []byte

//go:embed mapping/google/resource/apigee/google_apigee_target_server.json
var googleApigeeTargetServer []byte

//go:embed mapping/google/resource/appengine/google_app_engine_application.json
var googleAppEngineApplication []byte

//go:embed mapping/google/resource/appengine/google_app_engine_application_url_dispatch_rules.json
var googleAppEngineApplicationUrlDispacthRules []byte

//go:embed mapping/google/resource/appengine/google_app_engine_domain_mapping.json
var googleAppEngineDomainMapping []byte

//go:embed mapping/google/resource/appengine/google_app_engine_firewall_rule.json
var googleAppEngineFirewallRule []byte

//go:embed mapping/google/resource/appengine/google_app_engine_flexible_app_version.json
var googleAppEngineFlexibleAppVersion []byte

//go:embed mapping/google/resource/appengine/google_app_engine_service_network_settings.json
var googleAppEngineServiceNetworksSettings []byte

//go:embed mapping/google/resource/appengine/google_app_engine_service_split_traffic.json
var googleAppEngineServiceSplitTraffic []byte

//go:embed mapping/google/resource/appengine/google_app_engine_standard_app_version.json
var googleAppEngineStandardAppVersion []byte

//go:embed mapping/google/resource/apikeys/google_apikeys_key.json
var googleApikeysKey []byte

//go:embed mapping/google/resource/ces/google_ces_agent.json
var googleCesAgent []byte

//go:embed mapping/google/resource/ces/google_ces_app.json
var googleCesApp []byte

//go:embed mapping/google/resource/ces/google_ces_app_version.json
var googleCesAppVersion []byte

//go:embed mapping/google/resource/ces/google_ces_deployment.json
var googleCesDeployment []byte

//go:embed mapping/google/resource/ces/google_ces_example.json
var googleCesExample []byte

//go:embed mapping/google/resource/ces/google_ces_guardrail.json
var googleCesGuardrail []byte

//go:embed mapping/google/resource/ces/google_ces_tool.json
var googleCesTool []byte

//go:embed mapping/google/resource/ces/google_ces_toolset.json
var googleCesToolset []byte

//go:embed mapping/google/resource/cloudasset/google_cloud_asset_folder_feed.json
var googleCloudAssetFolderFeed []byte

//go:embed mapping/google/resource/cloudasset/google_cloud_asset_organization_feed.json
var googleClousAssetOrganizationFeed []byte

//go:embed mapping/google/resource/cloudasset/google_cloud_asset_project_feed.json
var googleCloudAssetProjectFeed []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_acl_config.json
var googleDiscoveryEngineAclConfig []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_assistant.json
var googleDiscoveryEngineAssistant []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_data_store.json
var googleDiscoveryEngineDataStore []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_chat_engine.json
var googleDiscoveryEngineChatEngine []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_cmek_config.json
var googleDiscoveryEngineCmekConfig []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_control.json
var googleDiscoveryEngineControl []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_data_connector.json
var googleDiscoveryEngineDataConnector []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_license_config.json
var googleDiscoveryEngineLicenseConfig []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_recommendation_engine.json
var googleDiscoveryEngineRecommendationEngine []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_schema.json
var googleDiscoveryEngineSchema []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_search_engine.json
var googleDiscoveryEngineSearchEmgine []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_sitemap.json
var googleDiscoveryEngineSitemap []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_target_site.json
var googleDiscoveryEngineTargetSite []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_user_store.json
var googleDiscoveryEngineUserStore []byte

//go:embed mapping/google/resource/discoveryengine/google_discovery_engine_widget_config.json
var googleDiscoveryEngineWidgetConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_event_threat_detection_custom_module.json
var googleSccEventThreatDetectionCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_folder_custom_module.json
var googleSccFolderCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_folder_notification_config.json
var googleSccFolderNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_folder_scc_big_query_export.json
var googleSccFolderSccBigQueryExport []byte

//go:embed mapping/google/resource/securitycenter/google_scc_management_folder_security_health_analytics_custom_module.json
var googleSccManagementFolderSecurityHealthAnalyticsCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_management_organization_event_threat_detection_custom_module.json
var googleSccManagementOrganizationEventThreatDetectionCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_management_organization_security_health_analytics_custom_module.json
var googleSccManagementOrganizationSecurityHealthAnalyticsCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_management_project_security_health_analytics_custom_module.json
var googleSccManagementProjectSecurityHealthAnalyticsCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_mute_config.json
var googleSccMuteConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_notification_config.json
var googleSccNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_organization_custom_module.json
var googleSccOrganizationCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_organization_scc_big_query_export.json
var googleSccOrganizationSccBigQueryExport []byte

//go:embed mapping/google/resource/securitycenter/google_scc_project_custom_module.json
var googleSccProjectCustomModule []byte

//go:embed mapping/google/resource/securitycenter/google_scc_project_notification_config.json
var googleSccProjectNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_project_scc_big_query_export.json
var googleSccProjectSccBigQueryExport []byte

//go:embed mapping/google/resource/securitycenter/google_scc_source.json
var googleSccSource []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_folder_mute_config.json
var googleSccV2folderMuteConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_folder_notification_config.json
var googleSccV2folderNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_folder_scc_big_query_export.json
var googleSccV2folderSccBigQueryExport []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_mute_config.json
var googleSccV2OrganizationMuteConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_notification_config.json
var googleSccV2OrganizationNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_scc_big_query_export.json
var googleSccV2OrganizationSccBigQueryExport []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_organization_scc_big_query_exports.json
var googleSccV2OrganizationSccBigQueryExports []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_project_mute_config.json
var googleSccV2ProjectMuteConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_project_notification_config.json
var googleSccV2ProjectNotificationConfig []byte

//go:embed mapping/google/resource/securitycenter/google_scc_v2_project_scc_big_query_export.json
var googleSccV2ProjectSccBigQueryExport []byte

//go:embed mapping/google/resource/dlp/google_data_loss_prevention_deidentify_template.json
var googleDataLossPreventionDeindentifyTemplate []byte

//go:embed mapping/google/resource/dlp/google_data_loss_prevention_discovery_config.json
var googleDataLossPreventionDiscoveryConfig []byte

//go:embed mapping/google/resource/dlp/google_data_loss_prevention_inspect_template.json
var googleDataLossPreventionInspectTemplate []byte

//go:embed mapping/google/resource/dlp/google_data_loss_prevention_job_trigger.json
var googleDataLossPreventionJobTrigger []byte

//go:embed mapping/google/resource/dlp/google_data_loss_prevention_stored_info_type.json
var googleDataLossPreventionStoredInfoType []byte

//go:embed mapping/google/resource/parametermanager/google_parameter_manager_parameter.json
var googleParameterManagerParameter []byte

//go:embed mapping/google/resource/parametermanager/google_parameter_manager_parameter_version.json
var googleParameterManagerParameterVersion []byte

//go:embed mapping/google/resource/parametermanager/google_parameter_manager_regional_parameter.json
var googleParameterManagerRegionalParameter []byte

//go:embed mapping/google/resource/parametermanager/google_parameter_manager_regional_parameter_version.json
var googleParameterManagerRegionalParameterVersion []byte

//go:embed mapping/google/resource/datapipelines/google_data_pipeline_pipeline.json
var googleDataPipelinePipeline []byte

//go:embed mapping/google/resource/datamigration/google_database_migration_service_connection_profile.json
var googleDatabaseMigrationServiceConnectionProfile []byte

//go:embed mapping/google/resource/datamigration/google_database_migration_service_migration_job.json
var googleDatabaseMigrationServiceMigrationJob []byte

//go:embed mapping/google/resource/datamigration/google_database_migration_service_private_connection.json
var googleDatabaseMigrationServicePrivateConnection []byte

//go:embed mapping/google/resource/dataflow/google_dataflow_flex_template_job.json
var googleDataflowFlexTemplateJob []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_data_product.json
var googleDataplexDataProduct []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_data_asset.json
var googleDataplexDataAsset []byte

//go:embed mapping/google/resource/dataplex/google_dataplex_entry_link.json
var googleDataplexEntryLink []byte

//go:embed mapping/google/resource/vmwareengine/google_vmwareengine_datastore.json
var googleVmwareengineDatastore []byte

//go:embed mapping/google/resource/cloudtasks/google_cloud_tasks_queue.json
var googleCloudTasksQueue []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_automation.json
var googleClouddeployAutomation []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_custom_target_type.json
var googleClouddeployCustomTargetType []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_delivery_pipeline.json
var googleClouddeployDeliveryPipeline []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_deploy_policy.json
var googleClouddeployDeployPolicy []byte

//go:embed mapping/google/resource/clouddeploy/google_clouddeploy_target.json
var googleClouddeployTarget []byte

//go:embed mapping/google/resource/datastream/google_datastream_connection_profile.json
var googleDatastreamConnectionProfile []byte

//go:embed mapping/google/resource/datastream/google_datastream_private_connection.json
var googleDatastreamPrivateConnection []byte

//go:embed mapping/google/resource/datastream/google_datastream_stream.json
var googleDatastreamStream []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_autonomous_database.json
var googleOracleDatabaseAutonomousDatabase []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_cloud_exadata_infrastructure.json
var googleOracleDatabaseCloudExadataInfrastructure []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_cloud_vm_cluster.json
var googleOracleDatabaseCloudVmCluster []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_db_system.json
var googleOracleDatabaseDbSystem []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_exascale_db_storage_vault.json
var googleOracleDatabaseExascaleDbStorageVault []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_odb_network.json
var googleOracleDatabaseOdbNetwork []byte

//go:embed mapping/google/resource/oracledatabase/google_oracle_database_odb_subnet.json
var googleOracleDatabaseOdbSubnet []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_release.json
var googleSaasRuntimeRelease []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_rollout_kind.json
var googleSaasRuntimeRolloutKind []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_saas.json
var googleSaasRuntimeSaas []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_tenant.json
var googleSaasRuntimeTenant []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_unit.json
var googleSaasRuntimeUnit []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_unit_kind.json
var googleSaasRuntimeUnitKind []byte

//go:embed mapping/google/resource/saasservicemgmt/google_saas_runtime_unit_operation.json
var googleSaasRuntimeUnitOperation []byte
