
resource "google_project_iam_custom_role" "terraform_pike" {
  project     = "pike-477416"
  role_id     = "terraform_pike"
  title       = "terraform_pike"
  description = "A user with least privileges"
  permissions = [
    # google_agent_registry_agent
    "agentregistry.agents.get",
    "agentregistry.agents.list",

    # google_agent_registry_endpoint
    "agentregistry.endpoints.get",
    "agentregistry.endpoints.list",

    # google_agent_registry_mcp_server
    "agentregistry.mcpServers.get",
    "agentregistry.mcpServers.list",

    # google_apigee_instance
    "apigee.instances.get",

    # google_bigquery_routine_iam_policy
    # non required

    # google_compute_firewall_policy_iam_policy
    # google_compute_network_firewall_policy_iam_policy
    "compute.firewallPolicies.getIamPolicy",

    # google_compute_instance_groups
    "compute.instanceGroups.list",

    # google_compute_region_network_firewall_policy_iam_policy
    "compute.regionFirewallPolicies.getIamPolicy",

    # google_compute_region_target_http_proxy
    "compute.regionTargetHttpProxies.get",

    # google_compute_region_target_https_proxy
    "compute.regionTargetHttpsProxies.get",

    # google_compute_service_attachment
    "compute.serviceAttachments.get",

    # google_compute_target_http_proxy
    "compute.targetHttpProxies.get",

    # google_compute_target_https_proxy
    "compute.targetHttpsProxies.get",

    # google_dataplex_data_product_iam_policy
    "dataplex.dataProducts.getIamPolicy",

    # google_discovery_engine_search_engine_iam_policy
    "discoveryengine.engines.getIamPolicy",

    # google_iap_agent_registry_iam_policy
    # google_iap_location_web_iam_policy
    "iap.web.getIamPolicy",

    # google_oracle_database_goldengate_connection_types
    "oracledatabase.goldenGateConnectionTypes.list",

    # google_oracle_database_goldengate_deployment_environments
    "oracledatabase.goldenGateDeploymentEnvironments.list",

    # google_oracle_database_goldengate_deployment_types
    "oracledatabase.goldenGateDeploymentTypes.list",

    # google_oracle_database_goldengate_deployment_versions
    "oracledatabase.goldenGateDeploymentVersions.list",

    # google_storage_control_project_intelligence_finding
    # google_storage_control_project_intelligence_finding_revision
    # "storage.intelligenceFindings.get",

    # google_storage_control_project_intelligence_findings
    # google_storage_control_project_intelligence_finding_revisions
    # "storage.intelligenceFindings.list",

    # google_storage_control_folder_intelligence_findings_summary
    # google_storage_control_organization_intelligence_findings_summary
    # google_storage_control_project_intelligence_findings_summary
    # "storage.intelligenceFindings.summarize",
  ]
}
