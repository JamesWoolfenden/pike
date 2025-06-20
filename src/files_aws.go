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

//go:embed mapping/aws/resource/sso/aws_ssoadmin_account_assignment.json
var awsSsoadminAccountAssignment []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_application_access_scope.json
var awsSsoadminApplicationAccessScope []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_application_assignment_configuration.json
var awsSsoadminApplicationAssignmentConfiguration []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_customer_managed_policy_attachment.json
var awsSsoadminCustomerManagedPolicyAttachment []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_instance_access_control_attributes.json
var awsSsoadminInstanceAccessControlAttributes []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_managed_policy_attachment.json
var awsSsoadminManagedPolicyAttachment []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_permission_set_inline_policy.json
var awsSsoadminPermissionSetInlinePolicy []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_permissions_boundary_attachment.json
var awsSsoadminPermissionsBoundaryAttachment []byte

//go:embed mapping/aws/resource/sso/aws_ssoadmin_trusted_token_issuer.json
var awsSsoadminTrustedTokenIssuer []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_account_suppression_attributes.json
var awsSesv2AccountSuppressionAttributes []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_account_vdm_attributes.json
var awsSesv2AccountVdmAttributes []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_configuration_set.json
var awsSesv2ConfigurationSet []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_configuration_set_event_destination.json
var awsSesv2ConfigurationSetEventDestination []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_contact_list.json
var awsSesv2ContactList []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_dedicated_ip_assignment.json
var awsSesv2DedicatedIPAssignment []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_dedicated_ip_pool.json
var awsSesv2DedicatedIPPool []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_email_identity.json
var awsSesv2EmailIdentity []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_email_identity_feedback_attributes.json
var awsSesv2EmailIdentityFeedbackAttributes []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_email_identity_mail_from_attributes.json
var awsSesv2EmailIdentityMailFromAttributes []byte

//go:embed mapping/aws/resource/ses/aws_sesv2_email_identity_policy.json
var awsSesv2EmailIdentityPolicy []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_cluster_snapshot.json
var awsRedshiftClusterSnapshot []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_data_share_authorization.json
var awsRedshiftDataShareAuthorization []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_data_share_consumer_association.json
var awsRedshiftDataShareConsumerAssociation []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_logging.json
var awsRedshiftLogging []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_partner.json
var awsRedshiftPartner []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_resource_policy.json
var awsRedshiftResourcePolicy []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_snapshot_copy.json
var awsRedshiftSnapshotCopy []byte

//go:embed mapping/aws/resource/redshift-data/aws_redshiftdata_statement.json
var awsRedshiftdataStatement []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_custom_domain_association.json
var awsRedshiftserverlessCustomDomainAssociation []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_endpoint_access.json
var awsRedshiftserverlessEndpointAccess []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_resource_policy.json
var awsRedshiftserverlessResourcePolicy []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_snapshot.json
var awsRedshiftserverlessSnapshot []byte

//go:embed mapping/aws/resource/redshift-serverless/aws_redshiftserverless_usage_limit.json
var awsRedshiftserverlessUsageLimit []byte

//go:embed mapping/aws/resource/route53domains/aws_route53domains_delegation_signer_record.json
var awsRoute53DomainsDelegationSignerRecord []byte

//go:embed mapping/aws/resource/route53domains/aws_route53domains_domain.json
var awsRoute53DomainsDomain []byte

//go:embed mapping/aws/resource/route53domains/aws_route53domains_registered_domain.json
var awsRoute53DomainsRegisteredDomain []byte

//go:embed mapping/aws/resource/route53profiles/aws_route53profiles_resource_association.json
var awsRoute53profilesResourceAssociation []byte

//go:embed mapping/aws/resource/s3-express/aws_s3_directory_bucket.json
var awsS3DirectoryBucket []byte

//go:embed mapping/aws/resource/s3/aws_s3_object_copy.json
var awsS3ObjectCopy []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_access_grants_instance_resource_policy.json
var awsS3controlAccessGrantInstanceResourcePolicy []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_bucket.json
var awsS3ControlBucket []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_bucket_lifecycle_configuration.json
var awsS3ControlBucketLifecycleAssociation []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_bucket_policy.json
var awsS3ControlBucketPolicy []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_object_lambda_access_point.json
var awsS3ControlObjectLambdaAccessPoint []byte

//go:embed mapping/aws/resource/s3-outposts/aws_s3control_object_lambda_access_point_policy.json
var awsS3ControlObjectLambdaAccessPointPolicy []byte

//go:embed mapping/aws/resource/s3/aws_s3control_storage_lens_configuration.json
var awsS3ControlStorageLensConfiguration []byte

//go:embed mapping/aws/resource/es/aws_opensearch_authorize_vpc_endpoint_access.json
var awsOpensearchAuthorizeVpcEndpointAccess []byte

//go:embed mapping/aws/resource/es/aws_opensearch_domain_saml_options.json
var awsOpensearchDomainSamlOptions []byte

//go:embed mapping/aws/resource/es/aws_opensearch_inbound_connection_accepter.json
var awsOpensearchInboundConnectionAccepter []byte

//go:embed mapping/aws/resource/es/aws_opensearch_outbound_connection.json
var awsOpensearchOutboundConnection []byte

//go:embed mapping/aws/resource/es/aws_opensearch_package.json
var awsOpensearchPackage []byte

//go:embed mapping/aws/resource/es/aws_opensearch_package_association.json
var awsOpensearchPackageAssociation []byte

//go:embed mapping/aws/resource/es/aws_opensearch_vpc_endpoint.json
var awsOpensearchVPCEndpoint []byte

//go:embed mapping/aws/resource/servicediscovery/aws_service_discovery_http_namespace.json
var awsServiceDiscoveryHttpNamespace []byte

//go:embed mapping/aws/resource/servicediscovery/aws_service_discovery_instance.json
var awsServiceDiscoveryInstance []byte

//go:embed mapping/aws/resource/servicediscovery/aws_service_discovery_private_dns_namespace.json
var awsServiceDiscoveryPrivateDNSNamespace []byte

//go:embed mapping/aws/resource/servicediscovery/aws_service_discovery_public_dns_namespace.json
var awsServiceDiscoveryPublicDNSNamespace []byte

//go:embed mapping/aws/resource/servicediscovery/aws_service_discovery_service.json
var awsServiceDiscoveryService []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_account.json
var awsMacieAccount []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_classification_export_configuration.json
var awsMacieClassificationExportConfiguration []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_classification_job.json
var awsMacieClassificationJob []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_invitation_accepter.json
var awsMacieInvitationAccepter []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_member.json
var awsMacieMember []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_organization_admin_account.json
var awsMacieOrganizationAdminAccount []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_policy.json
var awsVpcEndpointPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_private_dns.json
var awsVpcEndpointPrivateDns []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_security_group_association.json
var awsVpcEndpointSecurityGroupAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_service_allowed_principal.json
var awsVpcEndpointServiceAllowedPrincipal []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_service_private_dns_verification.json
var awsVpcEndpointServicePrivateDnsVerification []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipam_organization_admin_account.json
var awsVpcIpamOrganizationAdminAccount []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_ipv6_cidr_block_association.json
var awsVpcIpv6CidrBlockAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_network_performance_metric_subscription.json
var awsVpcNetworkPerformanceMetricSubscription []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_security_group_vpc_association.json
var awsVpcSecurityGroupAssociation []byte

//go:embed mapping/aws/resource/vpc-lattice/aws_vpclattice_service_network_resource_association.json
var awsVpclatticeServiceNetworkResourceAssociation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_endpoint_connection_accepter.json
var awsVpcEndpointConnectionAccepter []byte

//go:embed mapping/aws/resource/iot/aws_iot_event_configurations.json
var awsIotEventConfigurations []byte

//go:embed mapping/aws/resource/iot/aws_iot_indexing_configuration.json
var awsIotIndexingConfiguration []byte

//go:embed mapping/aws/resource/iot/aws_iot_logging_options.json
var awsIotLoggingOptions []byte

//go:embed mapping/aws/resource/iot/aws_iot_policy_attachment.json
var awsIotPolicyAttachment []byte

//go:embed mapping/aws/resource/iot/aws_iot_thing_group_membership.json
var awsIotThingGroupMembership []byte

//go:embed mapping/aws/resource/iot/aws_iot_thing_principal_attachment.json
var awsIotThingPrincipalAttachment []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_listener_certificate.json
var awsLbListenerCertificate []byte

//go:embed mapping/aws/resource/elasticloadbalancing/aws_lb_ssl_negotiation_policy.json
var awsLbSslNegotiationPolicy []byte

//go:embed mapping/aws/resource/ec2/aws_main_route_table_association.json
var awsMainRouteTableAssociation []byte

//go:embed mapping/aws/resource/memorydb/aws_memorydb_multi_region_cluster.json
var awsMemorydbMultiRegionCluster []byte

//go:embed mapping/aws/resource/kafka/aws_msk_single_scram_secret_association.json
var awsMskSingleScramSecretAssociation []byte

//go:embed mapping/aws/resource/kafkaconnect/aws_mskconnect_custom_plugin.json
var awsMskconnectCustomPlugin []byte

//go:embed mapping/aws/resource/kafkaconnect/aws_mskconnect_worker_configuration.json
var awsMskconnectWorkerConfiguration []byte

//go:embed mapping/aws/resource/ec2/aws_network_interface_sg_attachment.json
var awsNetworkInterfaceSgAttachment []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_attachment_accepter.json
var awsNetworkmanagerAttachmentAccepter []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_connection.json
var awsNetworkmanagerConnection []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_core_network_policy_attachment.json
var awsNetworkmanagerCoreNetworkPolicyAttachment []byte

//go:embed mapping/aws/resource/networkmanager/aws_networkmanager_transit_gateway_connect_peer_association.json
var awsNetworkmanagerTransitGatewayConnectPeerAssociation []byte

//go:embed mapping/aws/resource/networkmonitor/aws_networkmonitor_monitor.json
var awsNetworkmonitorMonitor []byte

//go:embed mapping/aws/resource/networkmonitor/aws_networkmonitor_probe.json
var awsNetworkmonitorProbe []byte

//go:embed mapping/aws/resource/organizations/aws_organizations_delegated_administrator.json
var awsOrganizationsDelegatedAdministrator []byte

//go:embed mapping/aws/resource/aps/aws_prometheus_rule_group_namespace.json
var awsPrometheusRuleGroupNamespace []byte

//go:embed mapping/aws/resource/qldb/aws_qldb_ledger.json
var awsQlbdLedger []byte

//go:embed mapping/aws/resource/ram/aws_ram_resource_share_accepter.json
var awsRamResourceShareAccepter []byte

//go:embed mapping/aws/resource/ram/aws_ram_sharing_with_organization.json
var awsRamSharingWithOrganization []byte

//go:embed mapping/aws/resource/rbin/aws_rbin_rule.json
var awsRbinRule []byte

//go:embed mapping/aws/resource/rds/aws_rds_instance_state.json
var awsRdsInstanceState []byte

//go:embed mapping/aws/resource/resiliencehub/aws_resiliencehub_resiliency_policy.json
var awsResilienceResiliencePolicy []byte

//go:embed mapping/aws/resource/resource-groups/aws_resourcegroups_resource.json
var awsResourcegroupsResource []byte

//go:embed mapping/aws/resource/rum/aws_rum_metrics_destination.json
var awsRumMetricsDestination []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_invite_accepter.json
var awsSecurityHubInviteAccepter []byte

//go:embed mapping/aws/resource/securityhub/aws_securityhub_member.json
var awsSecurityhubMember []byte

//go:embed mapping/aws/resource/serverlessrepo/aws_serverlessapplicationrepository_cloudformation_stack.json
var awsServerlessapplicationrepositoryCloudformationStack []byte

//go:embed mapping/aws/resource/servicecatalog/aws_servicecatalog_provisioning_artifact.json
var awsServiceCatalogProvisioningArtifact []byte

//go:embed mapping/aws/resource/shield/aws_shield_application_layer_automatic_response.json
var awsShieldApplicationLayerAutomaticsResponse []byte

//go:embed mapping/aws/resource/shield/aws_shield_drt_access_log_bucket_association.json
var awsShieldDrtAccessLogBucketAssociation []byte

//go:embed mapping/aws/resource/shield/aws_shield_drt_access_role_arn_association.json
var awsShieldDrtAccessRoleArnAssociation []byte

//go:embed mapping/aws/resource/shield/aws_shield_protection_health_check_association.json
var awsShieldProtectionHealthCheckAssociation []byte

//go:embed mapping/aws/resource/timestream/aws_timestreamquery_scheduled_query.json
var awsTimestreamqueryScheduledQuery []byte

//go:embed mapping/aws/resource/verified-access/aws_verifiedaccess_instance_trust_provider_attachment.json
var awsVerifiedaccessInstanceTrustProviderAttachment []byte

//go:embed mapping/aws/resource/verifiedpermissions/aws_verifiedpermissions_schema.json
var awsVerifiedpermissionsSchema []byte

//go:embed mapping/aws/resource/cognito-idp/aws_cognito_managed_user_pool_client.json
var awsCognitoManagedUserPoolClient []byte

//go:embed mapping/aws/resource/ec2/aws_ebs_fast_snapshot_restore.json
var awsEbsFastSnapshotRestore []byte

//go:embed mapping/aws/resource/ecr/aws_ecr_account_setting.json
var awsEcrAccountSetting []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_account_setting_default.json
var awsEcsAccountSettingDefault []byte

//go:embed mapping/aws/resource/ecs/aws_ecs_cluster_capacity_providers.json
var awsEcsClusterCapacityProviders []byte

//go:embed mapping/aws/resource/ec2/aws_eip_domain_name.json
var awsEipDomainName []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_reserved_cache_node.json
var awsElasticacheReservedCacheNode []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_serverless_cache.json
var awsElasticacheServerlessCache []byte

//go:embed mapping/aws/resource/elasticache/aws_elasticache_user_group_association.json
var awsElasticacheUserGroupAssociation []byte

//go:embed mapping/aws/resource/iam/aws_iam_group_policy_attachments_exclusive.json
var awsIamGroupPolicyAttachmentExclusive []byte

//go:embed mapping/aws/resource/iam/aws_iam_organizations_features.json
var awsIamOrganizationsFeatures []byte

//go:embed mapping/aws/resource/iam/aws_iam_role_policy_attachments_exclusive.json
var awsIamRolePolicyAttachmentsExclusive []byte

//go:embed mapping/aws/resource/iam/aws_iam_user_policy_attachments_exclusive.json
var awsIamUserPolicyAttachmentsExclusive []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_workspace_saml_configuration.json
var awsGrafanaWorkspaceSamlConfiguration []byte

//go:embed mapping/aws/resource/grafana/aws_grafana_role_association.json
var awsGrafanaRoleAssociation []byte

//go:embed mapping/aws/resource/es/aws_elasticsearch_domain_saml_options.json
var awsElasticsearchDomainSamlOptions []byte

//go:embed mapping/aws/resource/es/aws_opensearch_vpc_endpoint.json
var awsElasticsearchVpcEndpoint []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_contributor_insight_rule.json
var awsCloudwatchContributorInsightRule []byte

//go:embed mapping/aws/resource/cloudwatch/aws_cloudwatch_contributor_managed_insight_rule.json
var awsCloudwatchContributorManagedInsightRule []byte

//go:embed mapping/aws/resource/qbusiness/aws_qbusiness_application.json
var awsQbusinessApplication []byte

//go:embed mapping/aws/resource/dataexchange/aws_dataexchange_event_action.json
var awsDataexchangeEventAction []byte

//go:embed mapping/aws/resource/macie2/aws_macie2_organization_configuration.json
var awsMacie2OrganizationConfiguration []byte

//go:embed mapping/aws/resource/neptune-graph/aws_neptunegraph_graph.json
var awsNeptunegraphGraph []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_role_membership.json
var awsQuicksightRoleMembership []byte

//go:embed mapping/aws/resource/rds/aws_rds_shard_group.json
var awsRdsShardGroup []byte

//go:embed mapping/aws/resource/xray/aws_xray_resource_policy.json
var awsXrayResourcePolicy []byte

//go:embed mapping/aws/resource/ec2/aws_network_interface_permission.json
var awsNetworkInterfacePermission []byte

//go:embed mapping/aws/resource/route53/aws_route53_records_exclusive.json
var awsRoute53RecordsExclusive []byte

//go:embed mapping/aws/resource/athena/aws_athena_capacity_reservation.json
var awsAthenaCapacityReservation []byte

//go:embed mapping/aws/resource/redshift/aws_redshift_integration.json
var awsRedshiftIntegration []byte

//go:embed mapping/aws/resource/backup/aws_backup_restore_testing_plan.json
var awsBackupRestoreTestingPlan []byte

//go:embed mapping/aws/resource/ec2/aws_ec2_default_credit_specification.json
var awsEc2DefaultCreditSpecification []byte

//go:embed mapping/aws/resource/bedrock/aws_bedrockagent_prompt.json
var awsBedrockagentPrompt []byte

//go:embed mapping/aws/resource/cloudfront-keyvaluestore/aws_cloudfrontkeyvaluestore_keys_exclusive.json
var awsCloudfrontkeyvaluestoreKeysExclusive []byte

//go:embed mapping/aws/resource/dataexchange/aws_dataexchange_revision_assets.json
var awsDataexchangeRevisionAssets []byte

//go:embed mapping/aws/resource/inspector2/aws_inspector2_filter.json
var awsInspector2Filter []byte

//go:embed mapping/aws/resource/wafv2/aws_wafv2_api_key.json
var awsWafv2ApiKey []byte

//go:embed mapping/aws/resource/notifications/aws_notifications_channel_association.json
var awsNotificationsChannelAssociation []byte

//go:embed mapping/aws/resource/notifications/aws_notifications_event_rule.json
var awsNotificationsEventRule []byte

//go:embed mapping/aws/resource/notifications/aws_notifications_notification_configuration.json
var awsNotificationsNotificationConfiguration []byte

//go:embed mapping/aws/resource/notifications/aws_notifications_notification_hub.json
var awsNotificationsNotificationHub []byte

//go:embed mapping/aws/resource/notifications-contacts/aws_notificationscontacts_email_contact.json
var awsNotificationscontactsEmailContact []byte

//go:embed mapping/aws/resource/quicksight/aws_quicksight_account_settings.json
var awsQuicksightAccountSettings []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_browser_settings.json
var awsWorkspaceswebBrowserSettings []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_network_settings.json
var awsWorkspaceswebNetworkSettings []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_user_settings.json
var awsWorkspaceswebUserSettings []byte

//go:embed mapping/aws/resource/dsql/aws_dsql_cluster.json
var awsDsqlCluster []byte

//go:embed mapping/aws/resource/dsql/aws_dsql_cluster_peering.json
var awsDsqlClusterPeering []byte

//go:embed mapping/aws/resource/lightsail/aws_lightsail_bucket_access_key_access_key.json
var awsLightsailBucketAccessKey []byte

//go:embed mapping/aws/resource/aps/aws_prometheus_workspace_configuration.json
var awsPrometheusWorkspaceConfiguration []byte

//go:embed mapping/aws/resource/s3control/aws_s3control_directory_bucket_access_point_scope.json
var awsS3controlDirectoryBucketAccessPointScope []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_route_server.json
var awsVpcRouteServer []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_route_server_propagation.json
var awsVpcRouteServerPropagation []byte

//go:embed mapping/aws/resource/ec2/aws_vpc_route_server_vpc_association.json
var awsVpcRouteServerVpcAssociation []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_data_protection_settings.json
var awsWorkspaceswebDataProtectionSettings []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_ip_access_settings.json
var awsWorkspaceswebIpAccessSettings []byte

//go:embed mapping/aws/resource/workspaces-web/aws_workspacesweb_user_access_logging_settings.json
var awsWorkspaceswebUserAccessLoggingSettings []byte
