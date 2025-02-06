package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/aws/resource/dataexchange/aws_dataexchange_data_set.json
var awsDataexchangeDataSet []byte

//go:embed mapping/aws/resource/dataexchange/aws_dataexchange_revision.json
var awsDataexchangeRevision []byte

//go:embed mapping/aws/resource/devops-guru/aws_devopsguru_event_sources_config.json
var awsDevopsguruEventSourcesConfig []byte

//go:embed mapping/aws/resource/devops-guru/aws_devopsguru_service_integration.json
var awsDevopsguruServiceIntegration []byte

//go:embed mapping/aws/resource/drs/aws_drs_replication_configuration_template.json
var awsDrsReplicationConfigurationTemplate []byte

//go:embed mapping/aws/resource/elastictranscoder/aws_elastictranscoder_pipeline.json
var awsElastictranscoderPipeline []byte

//go:embed mapping/aws/resource/elastictranscoder/aws_elastictranscoder_preset.json
var awsElastictranscoderPreset []byte

//go:embed mapping/aws/resource/kinesisanalytics/aws_kinesis_analytics_application.json
var awsKinesisanalyticsApplication []byte

//go:embed mapping/aws/resource/kinesisanalytics/aws_kinesisanalyticsv2_application_snapshot.json
var awsKinesisanalyticsv2ApplicationSnapshot []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_lf_tag.json
var awsLakeformationLfTag []byte

//go:embed mapping/aws/resource/lakeformation/aws_lakeformation_resource_lf_tags.json
var awsLakeformationResourceLfTags []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_function_recursion_config.json
var awsLambdaFunctionRecursionConfig []byte

//go:embed mapping/aws/resource/lambda/aws_lambda_runtime_management_config.json
var awsLambdaRuntimeManagementConfig []byte

//go:embed mapping/aws/resource/license-manager/aws_licensemanager_association.json
var awsLicensemanagerAssociation []byte

//go:embed mapping/aws/resource/license-manager/aws_licensemanager_grant_accepter.json
var awsLicensemanagerGrantAccepter []byte

//go:embed mapping/aws/resource/mediapackagev2/aws_media_packagev2_channel_group.json
var awsMediaPackagev2ChannelGroup []byte

//go:embed mapping/aws/resource/mediastore/aws_media_store_container.json
var awsMediaStoreContainer []byte

//go:embed mapping/aws/resource/mediastore/aws_media_store_container_policy.json
var awsMediaStoreContainerPolicy []byte

//go:embed mapping/aws/resource/medialive/aws_medialive_channel.json
var awsMedialiveChannel []byte

//go:embed mapping/aws/resource/devicefarm/aws_devicefarm_device_pool.json
var awsDevicefarmDevicePool []byte

//go:embed mapping/aws/resource/devicefarm/aws_devicefarm_instance_profile.json
var awsDevicefarmInstanceProfile []byte

//go:embed mapping/aws/resource/devicefarm/aws_devicefarm_network_profile.json
var awsDevicefarmNetworkProfile []byte

//go:embed mapping/aws/resource/devicefarm/aws_devicefarm_project.json
var awsDevicefarmProject []byte

//go:embed mapping/aws/resource/devicefarm/aws_devicefarm_upload.json
var awsDevicefarmUpload []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_cluster.json
var awsFinspaceKxCluster []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_database.json
var awsFinspaceKxDatabase []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_dataview.json
var awsFinspaceKxDataview []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_scaling_group.json
var awsFinspaceKxScalingGroup []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_user.json
var awsFinspaceKxUser []byte

//go:embed mapping/aws/resource/finspace/aws_finspace_kx_volume.json
var awsFinspaceKxVolume []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_bucket_access_key_access_key.json
var awsLightsailBucketAccessKeyAccessKey []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_bucket_resource_access.json
var awsLightsailBucketResourceAccess []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_container_service.json
var awsLightsailContainerService []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_container_service_deployment_version.json
var awsLightsailContainerServiceDeploymentVersion []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_disk_attachment.json
var awsLightsailDiskAttachment []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_domain.json
var awsLightsailDomain []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_domain_entry.json
var awsLightsailDomainEntry []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb_attachment.json
var awsLightsailLbAttachment []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb_certificate.json
var awsLightsailLbCertificate []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb_certificate_attachment.json
var awsLightsailLbCertificateAttachment []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb_https_redirection_policy.json
var awsLightsailLbHTTPSRedirectionPolicy []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_lb_stickiness_policy.json
var awsLightsailLbStickinessPolicy []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_bgp_peer.json
var awsDxBgpPeer []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_connection.json
var awsDxConnection []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_connection_association.json
var awsDxConnectionAssociation []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_connection_confirmation.json
var awsDxConnectionConfirmation []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_gateway_association_proposal.json
var awsDxGatewayAssociationProposal []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_connection.json
var awsDxHostedConnection []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_private_virtual_interface.json
var awsDxHostedPrivateVirtualInterface []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_private_virtual_interface_accepter.json
var awsDxHostedPrivateVirtualInterfaceAccepter []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_public_virtual_interface.json
var awsDxHostedPublicVirtualInterface []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_public_virtual_interface_accepter.json
var awsDxHostedPublicVirtualInterfaceAccepter []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_hosted_transit_virtual_interface.json
var awsDxHostedTransitVirtualInterface []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_lag.json
var awsDxLag []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_macsec_key_association.json
var awsDxMacsecKeyAssociation []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_private_virtual_interface.json
var awsDxPrivateVirtualInterface []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_public_virtual_interface.json
var awsDxPublicVirtualInterface []byte

//go:embed mapping/aws/resource/directconnect/aws_dx_transit_virtual_interface.json
var awsDxTransitVirtualInterface []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_conditional_forwarder.json
var awsDirectoryServiceConditionalForwarder []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_radius_settings.json
var awsDirectoryServiceRadiusSettings []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_region.json
var awsDirectoryServiceRegion []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_shared_directory.json
var awsDirectoryServiceSharedDirectory []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_shared_directory_accepter.json
var awsDirectoryServiceSharedDirectoryAccepter []byte

//go:embed mapping/aws/resource/directoryservice/aws_directory_service_trust.json
var awsDirectoryServiceTrust []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_application.json
var awsOpsworksApplication []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_custom_layer.json
var awsOpsworksCustomLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_ecs_cluster_layer.json
var awsOpsworksEcsClusterLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_ganglia_layer.json
var awsOpsworksGangliaLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_haproxy_layer.json
var awsOpsworksHaproxyLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_instance.json
var awsOpsworksInstance []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_java_app_layer.json
var awsOpsworksJavaAppLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_memcached_layer.json
var awsOpsworksMemcachedLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_mysql_layer.json
var awsOpsworksMysqlLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_nodejs_app_layer.json
var awsOpsworksNodejsAppLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_permission.json
var awsOpsworksPermission []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_php_app_layer.json
var awsOpsworksPhpAppLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_rails_app_layer.json
var awsOpsworksRailsAppLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_rds_db_instance.json
var awsOpsworksRdsDBInstance []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_stack.json
var awsOpsworksStack []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_static_web_layer.json
var awsOpsworksStaticWebLayer []byte

//go:embed mapping/aws/resource/opsworks/aws_opsworks_user_profile.json
var awsOpsworksUserProfile []byte

//go:embed mapping/aws/resource/kafka/aws_mskconnect_connector.json
var awsMskconnectConnector []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_kinesis_streaming_destination.json
var awsDynamodbKinesisStreamingDestination []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_resource_policy.json
var awsDynamodbResourcePolicy []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_table_export.json
var awsDynamodbTableExport []byte

//go:embed mapping/aws/resource/dynamodb/aws_dynamodb_table_replica.json
var awsDynamodbTableReplica []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_snapshot_import.json
var awsEbsSnapshotImport []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_capacity_block_reservation.json
var awsEc2CapacityBlockReservation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_instance_metadata_defaults.json
var awsEc2InstanceMetadataDefaults []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_instance_state.json
var awsEc2InstanceState []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_managed_prefix_list_entry.json
var awsEc2ManagedPrefixListEntry []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_subnet_cidr_reservation.json
var awsEc2SubnetCidrReservation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_connect.json
var awsEc2TransitGatewayConnectPeer []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_default_route_table_association.json
var awsEc2TransitGatewayDefaultRouteTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_route_table_propagation.json
var awsEc2TransitGatewayDefaultRouteTablePropagation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_peering_attachment_accepter.json
var awsEc2TransitGatewayPeeringAttachmentAccepter []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_policy_table.json
var awsEc2TransitGatewayPolicyTable []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_policy_table_association.json
var awsEc2TransitGatewayPolicyTableAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_prefix_list_reference.json
var awsEc2TransitGatewayPrefixListReference []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_transit_gateway_vpc_attachment_accepter.json
var awsEc2TransitGatewayVpcAttachmentAccepter []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_detector_feature.json
var awsGuarddutyDetectorFeature []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_invite_accepter.json
var awsGuarddutyInviteAcceptor []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_member_detector_feature.json
var awsGuarddutyMemberDetectorFeature []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_organization_admin_account.json
var awsGuarddutyOrganizationAdminAccount []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_organization_configuration.json
var awsGuarddutyOrganizationConfiguration []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_publishing_destination.json
var awsGuarddutyPublishingDestination []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_bot.json
var awsLexv2ModelsBot []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_bot_locale.json
var awsLexv2ModelsBotLocale []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_bot_version.json
var awsLexv2ModelsBotVersion []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_intent.json
var awsLexv2ModelsIntent []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_slot.json
var awsLexv2ModelsSlot []byte

//go:embed mapping/aws/resource/lex/aws_lexv2models_slot_type.json
var awsLexv2ModelsSlotType []byte

//go:embed mapping/aws/resource/guardduty/aws_guardduty_organization_configuration_feature.json
var awsGuarddutyOrganizationConfigurationFeature []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_adm_channel.json
var awsPinpointAdmChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_apns_channel.json
var awsPinpointApnsChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_apns_sandbox_channel.json
var awsPinpointApnsSandboxChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_apns_voip_channel.json
var awsPinpointApnsVoipChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_apns_voip_sandbox_channel.json
var awsPinpointApnsVoipSandboxChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_app.json
var awsPinpointApp []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_baidu_channel.json
var awsPinpointBaiduChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_email_channel.json
var awsPinpointEmailChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_email_template.json
var awsPinpointEmailTemplate []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_event_stream.json
var awsPinpointEventStream []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_gcm_channel.json
var awsPinpointGcmChannel []byte

//go:embed mapping/aws/resource/mobiletargeting/aws_pinpoint_sms_channel.json
var awsPinpointSmsChannel []byte

//go:embed mapping/aws/resource/sms-voice/aws_pinpointsmsvoicev2_configuration_set.json
var awsPinpointsmsvoicev2ConfigurationSet []byte

//go:embed mapping/aws/resource/sms-voice/aws_pinpointsmsvoicev2_opt_out_list.json
var awsPinpointsmsvoicev2OptOutList []byte

//go:embed mapping/aws/resource/sms-voice/aws_pinpointsmsvoicev2_phone_number.json
var awsPinpointsmsvoicev2PhoneNumber []byte
