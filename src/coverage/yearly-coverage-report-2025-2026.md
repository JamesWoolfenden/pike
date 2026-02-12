# Pike Cloud Coverage Report: Year-over-Year Analysis (Feb 2025 - Feb 2026)

## Executive Summary

This report analyzes the changes in Pike's Terraform resource coverage across AWS, Google Cloud Platform (GCP), and Azure over the past year (February 2025 to February 2026).

### Key Findings

- **AWS**: Coverage decreased from 100% to 93.96% as new resources were added to Terraform faster than Pike could implement them
- **GCP**: Coverage increased dramatically from 18.13% to 78.50%, showing significant development effort
- **Azure**: Coverage remained minimal at ~5%, showing limited focus on this platform

## Detailed Platform Analysis

### Amazon Web Services (AWS)

| Metric                    | Feb 2025 | Jul 2025 | Jan 2026 | Feb 2026 | Change (YoY) |
|---------------------------|----------|----------|----------|----------|--------------|
| **Resources Coverage**    | 100.00%  | 99.80%   | 94.78%   | 93.96%   | **-6.04%**   |
| **Resources Supported**   | N/A      | N/A      | 1,508    | 1,508    | -            |
| **Total Resources**       | N/A      | N/A      | 1,591    | 1,605    | -            |
| **Resources Missing**     | 0        | 3        | 83       | 97       | **+97**      |
| **Datasource Coverage**   | 100.00%  | 100.00%  | 99.06%   | 99.84%   | **-0.16%**   |
| **Datasources Supported** | N/A      | N/A      | 632      | 641      | -            |
| **Total Datasources**     | N/A      | N/A      | 638      | 642      | -            |
| **Datasources Missing**   | 0        | 0        | 6        | 1        | **+1**       |

**Analysis:**

- AWS started at 100% coverage in February 2025
- By July 2025, coverage remained high at 99.80% with only 3 missing resources
- Current coverage is 93.96%, indicating AWS Terraform provider has added 97 new resources in the past year
- Despite maintaining 1,508 supported resources, the total available resources grew from 1,591 to 1,605
- Datasource coverage remains excellent at 99.84% (only 1 missing out of 642)

**Notable Missing Resources (Sample):**

- aws_apigatewayv2_routing_rule
- aws_bedrockagent_flow
- aws_billing_view
- aws_cloudfront_anycast_ip_list
- aws_cloudwatch_log_transformer
- aws_controltower_baseline
- Multiple new bedrock agent core resources
- Multiple new notification and monitoring resources

### Google Cloud Platform (GCP)

| Metric                    | Feb 2025 | Jul 2025 | Jan 2026 | Feb 2026 | Change (YoY) |
|---------------------------|----------|----------|----------|----------|--------------|
| **Resources Coverage**    | 18.13%   | 50.68%   | 68.90%   | 78.50%   | **+60.37%**  |
| **Resources Supported**   | N/A      | N/A      | 884      | 1,015    | **+131**     |
| **Total Resources**       | N/A      | N/A      | 1,283    | 1,293    | -            |
| **Resources Missing**     | ~980     | ~588     | 399      | 278      | **-702**     |
| **Datasource Coverage**   | 71.05%   | 99.74%   | 100.00%  | 100.00%  | **+28.95%**  |
| **Datasources Supported** | N/A      | N/A      | 410      | 411      | -            |
| **Total Datasources**     | N/A      | N/A      | 410      | 411      | -            |
| **Datasources Missing**   | ~350     | ~1       | 0        | 0        | **0**        |

**Analysis:**

- GCP showed the most dramatic improvement of all platforms
- Resource coverage increased from 18.13% to 78.50% - a gain of **60.37 percentage points**
- Between Jan and Feb 2026 alone, 131 new resources were added (884 → 1,015)
- Datasource coverage reached 100% (up from 71.05%)
- This represents significant engineering effort focused on GCP platform

**Notable Missing Resources (Sample):**

- google_cloud_ids_endpoint
- google_cloud_quotas_quota_preference
- google_cloudbuild_bitbucket_server_config
- Multiple compute engine resources (autoscaler, firewalls, VPN)
- Multiple container/GKE related resources
- Firebase and network security resources

### Microsoft Azure

| Metric                    | Feb 2025 | Jul 2025 | Jan 2026 | Feb 2026 | Change (YoY) |
|---------------------------|----------|----------|----------|----------|--------------|
| **Resources Coverage**    | 4.56%    | 4.48%    | 5.09%    | 5.12%    | **+0.56%**   |
| **Resources Supported**   | N/A      | N/A      | 57       | 57       | -            |
| **Total Resources**       | N/A      | N/A      | 1,120    | 1,114    | -            |
| **Resources Missing**     | ~1,251   | ~1,140   | 1,063    | 1,057    | **-194**     |
| **Datasource Coverage**   | 35.57%   | 33.80%   | 30.58%   | 31.28%   | **-4.29%**   |
| **Datasources Supported** | N/A      | N/A      | 122      | 122      | -            |
| **Total Datasources**     | N/A      | N/A      | 399      | 390      | -            |
| **Datasources Missing**   | ~808     | ~264     | 277      | 268      | **-540**     |

**Analysis:**

- Azure coverage remained minimal throughout the year
- Resource coverage increased only slightly from 4.56% to 5.12%
- Only 57 resources supported out of 1,114 total resources
- Datasource coverage actually decreased from 35.57% to 31.28%
- This indicates Azure is not a priority platform for Pike development

**Notable Missing Resources:**

- Over 1,000 Azure resources remain unsupported
- Major gaps across all service areas
- Limited movement in supported resource count

## Platform Comparison

### Coverage Trends

```shell
Platform    Feb 2025    Feb 2026    Change      Trend
========================================================
AWS         100.00%     93.96%      -6.04%      Declining (new resources outpacing implementation)
GCP          18.13%     78.50%     +60.37%      Rapidly Growing
Azure         4.56%      5.12%      +0.56%      Stagnant
```

### Development Focus

Based on the coverage changes, the development priorities over the past year were:

1. **Primary Focus: Google Cloud Platform**
   - Added 131+ new resources in the last month alone
   - Achieved 100% datasource coverage
   - More than quadrupled resource coverage (18% → 78%)

2. **Maintenance Mode: AWS**
   - Maintaining existing 1,508 resources
   - Keeping up with some new Terraform provider additions
   - Still maintaining near-complete coverage (93.96%)

3. **Minimal Focus: Azure**
   - Only 57 resources supported
   - No significant additions over the year
   - 95% of Azure resources remain unsupported

## Resource Growth Analysis

### Total Resources Available (Terraform Providers)

| Platform | Jan 2026                          | Feb 2026                          | Monthly Growth |
|----------|-----------------------------------|-----------------------------------|----------------|
| AWS      | 1,591 Resources / 638 Datasources | 1,605 Resources / 642 Datasources | +14 / +4       |
| GCP      | 1,283 Resources / 410 Datasources | 1,293 Resources / 411 Datasources | +10 / +1       |
| Azure    | 1,120 Resources / 399 Datasources | 1,114 Resources / 390 Datasources | -6 / -9        |

**Note:** Azure showing negative growth suggests resource consolidation or reclassification in the Terraform provider.

## Recommendations

### For AWS

- Continue monitoring new resource additions
- Prioritize implementing the 97 missing resources, particularly:
  - Bedrock Agent Core resources (AI/ML services)
  - CloudFront advanced features
  - Control Tower baseline resources
  - New networking features (VPC encryption, route servers)

### For GCP

- Continue the strong momentum
- Focus on the remaining 278 resources to reach 90%+ coverage
- Priority areas appear to be:
  - Compute Engine advanced features
  - GKE/Container resources
  - Network security resources
  - Firebase services

### For Azure

- Strategic decision needed: invest significantly or discontinue
- Current 5% coverage is insufficient for production use
- Would require substantial engineering effort to reach parity with AWS/GCP

## Methodology

This report was generated by analyzing git history of coverage files in the Pike repository:

- Baseline: Commit 1da865c2 (Feb 12, 2025)
- Mid-point: Commit 027bbbfb (Jul 13, 2025)
- Latest: Commits cc7e628c/c664026c (Feb 9, 2026)

Coverage percentages and resource counts were extracted from the markdown files in `src/coverage/`.

---

**Report Generated:** February 10, 2026
**Repository:** pike (<https://github.com/JamesWoolfenden/pike>)
**Analysis Period:** February 2025 - February 2026
