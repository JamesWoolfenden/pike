# Auth for working on Pike with Azure

Unlike AWS you cant name profiles.

You will need two distinct auth identities, one that has the power to update roles, so that you can update the
permissions of the other role that you use to determine the permissions required to create a resource.
You can create the simple role using the contents of <terraform/azurerm/role>; this is the role that you update.

I work in two shells, each with different permissions.

```shell
# sh
export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
export ARM_CLIENT_SECRET="12345678-0000-0000-0000-000000000000"
export ARM_TENANT_ID="10000000-0000-0000-0000-000000000000"
export ARM_SUBSCRIPTION_ID="20000000-0000-0000-0000-000000000000"
```

<https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/service_principal_client_secret>
