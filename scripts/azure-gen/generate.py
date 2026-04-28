#!/usr/bin/env python3
"""
Generate pike Azure permission mappings from aztft ARM type data,
validated against the live Azure RBAC operation catalog.

Usage:
  ./generate.py validate           # compare generated vs existing pike mappings
  ./generate.py emit <resource>    # print pike JSON for one resource
  ./generate.py emit-all <outdir>  # write JSON for every unmapped resource
"""
import json
import sys
from pathlib import Path

HERE = Path(__file__).parent
ROOT = HERE.parent.parent
EXISTING = ROOT / "src/mapping/azurerm/resource"
EXISTING_DATA = ROOT / "src/mapping/azurerm/data"
MEMBERS = json.loads((ROOT / "parse/azurerm-members.json").read_text())

aztft = json.loads((HERE / "aztft_map.json").read_text())

# RBAC ops are case-insensitive; index lowercase -> canonical
ops_canon = {}
for line in (HERE / "azure_ops.txt").read_text().splitlines():
    line = line.strip()
    if line:
        ops_canon[line.lower()] = line


def arm_type(entry):
    mp = entry.get("management_plane")
    if not mp:
        return None
    provider = mp.get("provider")
    types = mp.get("types") or []
    if not provider or not types:
        return None
    return provider + "/" + "/".join(types)


def valid(action):
    return ops_canon.get(action.lower())


def predict(tf_name):
    """Return predicted pike mapping dict, or None if no ARM type known."""
    entry = aztft.get(tf_name)
    if not entry:
        return None
    rtype = arm_type(entry)
    if not rtype:
        return None

    parts = rtype.split("/")
    apply, plan, destroy, modify = [], [], [], []

    rg = valid("Microsoft.Resources/subscriptions/resourceGroups/read")
    apply.append(rg)
    plan.append(rg)

    # core CRUD on the leaf type
    for verb in ("read", "write", "delete"):
        op = valid(f"{rtype}/{verb}")
        if op:
            apply.append(op)
            if verb == "read":
                plan.append(op)
            if verb == "write":
                modify.append(op)
            if verb == "delete":
                destroy.append(op)

    # parent reads for nested resources (e.g. databaseAccounts/tables -> databaseAccounts/read)
    for i in range(2, len(parts)):
        parent = "/".join(parts[:i])
        op = valid(f"{parent}/read")
        if op and op not in apply:
            apply.append(op)
            plan.append(op)

    # async polling reads only — deliberately NOT adding listKeys/listConnectionStrings
    # (those expose secrets; prefer under-grant and let users add when 403'd)
    for suffix in ("operationStatuses/read", "operationResults/read"):
        op = valid(f"{rtype}/{suffix}")
        if op and op not in apply:
            apply.append(op)

    return {
        "apply": apply,
        "attributes": {"tags": []},
        "destroy": destroy,
        "modify": modify,
        "plan": plan,
    }


# Datasources with no aztft entry: map to a known ARM type or [] (no RBAC needed).
DATA_OVERRIDES = {
    # client-side / connection-string based — no management-plane RBAC
    "azurerm_client_config": [],
    "azurerm_storage_account_sas": [],
    "azurerm_storage_account_blob_container_sas": [],
    "azurerm_eventhub_sas": [],
    "azurerm_key_vault_encrypted_value": [],
    "azurerm_billing_enrollment_account_scope": [],
    "azurerm_billing_mca_account_scope": [],
    "azurerm_billing_mpa_account_scope": [],
    # direct ARM type
    "azurerm_advisor_recommendations": "Microsoft.Advisor/recommendations",
    "azurerm_blueprint_definition": "Microsoft.Blueprint/blueprints",
    "azurerm_blueprint_published_version": "Microsoft.Blueprint/blueprints/versions",
    "azurerm_cosmosdb_restorable_database_accounts": "Microsoft.DocumentDB/locations/restorableDatabaseAccounts",
    "azurerm_dns_soa_record": "Microsoft.Network/dnsZones/SOA",
    "azurerm_private_dns_soa_record": "Microsoft.Network/privateDnsZones/SOA",
    "azurerm_hdinsight_cluster": "Microsoft.HDInsight/clusters",
    "azurerm_kubernetes_service_versions": "Microsoft.ContainerService/locations/kubernetesVersions",
    "azurerm_kubernetes_node_pool_snapshot": "Microsoft.ContainerService/snapshots",
    "azurerm_location": "Microsoft.Resources/subscriptions/locations",
    "azurerm_network_service_tags": "Microsoft.Network/locations/serviceTags",
    "azurerm_policy_assignment": "Microsoft.Authorization/policyAssignments",
    "azurerm_policy_definition_built_in": "Microsoft.Authorization/policyDefinitions",
    "azurerm_resources": "Microsoft.Resources/subscriptions/resources",
    "azurerm_subscriptions": "Microsoft.Resources/subscriptions",
    "azurerm_traffic_manager_geographical_location": "Microsoft.Network/trafficManagerGeographicHierarchies",
    "azurerm_web_pubsub_private_link_resource": "Microsoft.SignalRService/WebPubSub/privateLinkResources",
    "azurerm_sentinel_alert_rule_template": "Microsoft.SecurityInsights/alertRules",
    "azurerm_oracle_adbs_character_sets": "Oracle.Database/locations/autonomousDatabaseCharacterSets",
    "azurerm_oracle_adbs_national_character_sets": "Oracle.Database/locations/autonomousDatabaseNationalCharacterSets",
    "azurerm_oracle_db_nodes": "Oracle.Database/cloudVmClusters/dbNodes",
    "azurerm_oracle_db_servers": "Oracle.Database/cloudExadataInfrastructures/dbServers",
    "azurerm_oracle_db_system_shapes": "Oracle.Database/locations/dbSystemShapes",
    "azurerm_oracle_gi_versions": "Oracle.Database/locations/giVersions",
    "azurerm_oracle_database_system_versions": "Oracle.Database/locations/systemVersions",
    "azurerm_sentinel_alert_rule_anomaly": "Microsoft.SecurityInsights/alertRules",
    "azurerm_managed_api": "Microsoft.Web/locations/managedApis",
    "azurerm_monitor_diagnostic_categories": "Microsoft.Insights/diagnosticSettings",
    "azurerm_function_app_host_keys": "Microsoft.Web/sites/host/listkeys",
    "azurerm_platform_image": "Microsoft.Compute/locations/publishers",
    "azurerm_private_endpoint_connection": "Microsoft.Network/privateEndpoints",
    "azurerm_private_link_service_endpoint_connections": "Microsoft.Network/privateLinkServices",
    "azurerm_databricks_workspace_private_endpoint_connection": "Microsoft.Databricks/workspaces/privateEndpointConnections",
    "azurerm_log_analytics_workspace_table": "Microsoft.OperationalInsights/workspaces/tables",
    "azurerm_dynatrace_monitor": "Dynatrace.Observability/monitors",
    "azurerm_elastic_san_volume_snapshot": "Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots",
    "azurerm_system_center_virtual_machine_manager_inventory_items": "Microsoft.ScVmm/vmmServers/inventoryItems",
    "azurerm_key_vault_managed_hardware_security_module_key": "Microsoft.KeyVault/managedHSMs/keys",
    "azurerm_key_vault_managed_hardware_security_module_role_definition": "Microsoft.KeyVault/managedHSMs",
    # plurals → singular aztft key
    "azurerm_automation_variables": "@azurerm_automation_variable_string",
    "azurerm_data_factory_trigger_schedules": "@azurerm_data_factory_trigger_schedule",
    "azurerm_images": "@azurerm_image",
    "azurerm_ip_groups": "@azurerm_ip_group",
    "azurerm_managed_disks": "@azurerm_managed_disk",
    "azurerm_role_assignments": "@azurerm_role_assignment",
    "azurerm_shared_image_versions": "@azurerm_shared_image_version",
    "azurerm_storage_containers": "@azurerm_storage_container",
    "azurerm_storage_table_entities": "@azurerm_storage_table_entity",
    "azurerm_public_maintenance_configurations": "@azurerm_maintenance_configuration",
    "azurerm_extended_locations": "@azurerm_extended_location_custom_location",
}


def resolve_data_type(tf_name):
    if tf_name in DATA_OVERRIDES:
        v = DATA_OVERRIDES[tf_name]
        if isinstance(v, list):
            return "", v
        if v.startswith("@"):
            entry = aztft.get(v[1:])
            return (arm_type(entry), None) if entry else (None, None)
        return v, None
    entry = aztft.get(tf_name)
    if entry:
        return arm_type(entry), None
    # last-ditch: strip trailing 's'
    if tf_name.endswith("s"):
        entry = aztft.get(tf_name[:-1])
        if entry:
            return arm_type(entry), None
    return None, None


def predict_data(tf_name):
    """Datasource mapping: plan-only, read permissions on type + parents."""
    rtype, literal = resolve_data_type(tf_name)
    if literal is not None:
        return {"apply": [], "attributes": {"tags": []}, "destroy": [], "modify": [], "plan": literal}
    if not rtype:
        return None

    parts = rtype.split("/")
    plan = []
    op = valid(f"{rtype}/read")
    if op:
        plan.append(op)
    for i in range(2, len(parts)):
        op = valid(f"{'/'.join(parts[:i])}/read")
        if op and op not in plan:
            plan.append(op)
    if not plan:
        return None

    return {
        "apply": [],
        "attributes": {"tags": []},
        "destroy": [],
        "modify": [],
        "plan": plan,
    }


def load_existing():
    out = {}
    for p in EXISTING.rglob("azurerm_*.json"):
        try:
            data = json.loads(p.read_text())[0]
        except Exception:
            continue
        out[p.stem] = {
            "path": p,
            "apply": set(a for a in data.get("apply", []) if a),
        }
    return out


def validate():
    existing = load_existing()
    rows = []
    tot_pred = tot_actual = tot_hit = 0
    matched = 0

    for name, info in sorted(existing.items()):
        actual = {a.lower(): a for a in info["apply"]}
        gen = predict(name)
        if gen is None:
            rows.append((name, "NO ARM MAP", 0, len(actual), 0, 0))
            continue
        matched += 1
        pred = {p.lower(): p for p in gen["apply"]}
        hit = pred.keys() & actual.keys()
        miss = {actual[k] for k in actual.keys() - pred.keys()}
        extra = {pred[k] for k in pred.keys() - actual.keys()}
        tot_pred += len(pred)
        tot_actual += len(actual)
        tot_hit += len(hit)
        rows.append((name, "ok", len(hit), len(actual), len(miss), len(extra)))
        if miss or extra:
            print(f"\n{name}  [{info['path'].relative_to(ROOT)}]")
            print(f"  predicted {len(pred)}, actual {len(actual)}, overlap {len(hit)}")
            for m in sorted(miss):
                print(f"  - MISSED  {m}")
            for e in sorted(extra):
                print(f"  + EXTRA   {e}")

    print("\n" + "=" * 70)
    print(f"resources with ARM map: {matched}/{len(existing)}")
    if tot_actual:
        recall = 100 * tot_hit / tot_actual
        print(f"recall    (predicted ∩ actual / actual)   : {tot_hit}/{tot_actual} = {recall:.1f}%")
    if tot_pred:
        prec = 100 * tot_hit / tot_pred
        print(f"precision (predicted ∩ actual / predicted): {tot_hit}/{tot_pred} = {prec:.1f}%")
    perfect = sum(1 for r in rows if r[1] == "ok" and r[4] == 0 and r[5] == 0)
    print(f"exact matches: {perfect}/{matched}")


def emit(name):
    gen = predict(name)
    if gen is None:
        print(f"no ARM type mapping for {name}", file=sys.stderr)
        sys.exit(1)
    print(json.dumps([gen], indent=2))


def emit_all(outdir):
    outdir = Path(outdir)
    existing = set(load_existing())
    written = skipped = nocrud = 0
    for name in sorted(aztft):
        if name in existing:
            continue
        entry = aztft[name]
        rtype = arm_type(entry)
        gen = predict(name)
        if gen is None:
            skipped += 1
            continue
        # require at least one core CRUD op beyond the RG boilerplate
        if len(gen["apply"]) <= 1:
            nocrud += 1
            continue
        provider = rtype.split("/")[0]
        sub = provider.split(".", 1)[1].lower() if "." in provider else provider.lower()
        d = outdir / sub
        d.mkdir(parents=True, exist_ok=True)
        (d / f"{name}.json").write_text(json.dumps([gen], indent=2) + "\n")
        written += 1
    print(f"wrote {written}, skipped {skipped} (no ARM map), {nocrud} (no valid CRUD ops)")


def load_existing_data():
    out = {}
    for p in EXISTING_DATA.rglob("azurerm_*.json"):
        try:
            data = json.loads(p.read_text())[0]
        except Exception:
            continue
        out[p.stem] = set(a for a in data.get("plan", []) if a)
    return out


def validate_data():
    existing = load_existing_data()
    tot_pred = tot_actual = tot_hit = matched = 0
    for name, actual_set in sorted(existing.items()):
        actual = {a.lower(): a for a in actual_set}
        gen = predict_data(name)
        if gen is None:
            continue
        matched += 1
        pred = {p.lower(): p for p in gen["plan"]}
        hit = pred.keys() & actual.keys()
        miss = {actual[k] for k in actual.keys() - pred.keys()}
        extra = {pred[k] for k in pred.keys() - actual.keys()}
        tot_pred += len(pred)
        tot_actual += len(actual)
        tot_hit += len(hit)
        if miss or extra:
            print(f"{name}")
            for m in sorted(miss):
                print(f"  - MISSED  {m}")
            for e in sorted(extra):
                print(f"  + EXTRA   {e}")
    print("=" * 70)
    print(f"datasources with ARM map: {matched}/{len(existing)}")
    if tot_actual:
        print(f"recall   : {tot_hit}/{tot_actual} = {100*tot_hit/tot_actual:.1f}%")
    if tot_pred:
        print(f"precision: {tot_hit}/{tot_pred} = {100*tot_hit/tot_pred:.1f}%")


def emit_all_data(outdir):
    outdir = Path(outdir)
    existing = set(load_existing_data())
    written = nomap = 0
    for name in sorted(MEMBERS.get("dataSources", [])):
        if name in existing:
            continue
        gen = predict_data(name)
        if gen is None:
            nomap += 1
            continue
        if gen["plan"]:
            provider = gen["plan"][0].split("/")[0]
            sub = provider.split(".", 1)[1].lower() if "." in provider else provider.lower()
        else:
            sub = "noop"
        d = outdir / sub
        d.mkdir(parents=True, exist_ok=True)
        (d / f"{name}.json").write_text(json.dumps([gen], indent=2) + "\n")
        written += 1
    print(f"wrote {written}, skipped {nomap} (no ARM map / no read op)")


if __name__ == "__main__":
    cmd = sys.argv[1] if len(sys.argv) > 1 else "validate"
    if cmd == "validate":
        validate()
    elif cmd == "validate-data":
        validate_data()
    elif cmd == "emit":
        emit(sys.argv[2])
    elif cmd == "emit-all":
        emit_all(sys.argv[2])
    elif cmd == "emit-all-data":
        emit_all_data(sys.argv[2])
    else:
        print(__doc__)
        sys.exit(1)
