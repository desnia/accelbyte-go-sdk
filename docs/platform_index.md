# Platform Service Index

&nbsp;

## Operations

### FulfillmentScript Wrapper:  [FulfillmentScript](../services-api/pkg/service/platform/fulfillmentScript.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/fulfillment/scripts` | GET | ListFulfillmentScriptsShort | [ListFulfillmentScriptsShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [ListFulfillmentScriptsShort](../services-api/pkg/service/platform/fulfillmentScript.go) |
| `/platform/admin/fulfillment/scripts/tests/eval` | POST | TestFulfillmentScriptEvalShort | [TestFulfillmentScriptEvalShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [TestFulfillmentScriptEvalShort](../services-api/pkg/service/platform/fulfillmentScript.go) |
| `/platform/admin/fulfillment/scripts/{id}` | GET | GetFulfillmentScriptShort | [GetFulfillmentScriptShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [GetFulfillmentScriptShort](../services-api/pkg/service/platform/fulfillmentScript.go) |
| `/platform/admin/fulfillment/scripts/{id}` | POST | CreateFulfillmentScriptShort | [CreateFulfillmentScriptShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [CreateFulfillmentScriptShort](../services-api/pkg/service/platform/fulfillmentScript.go) |
| `/platform/admin/fulfillment/scripts/{id}` | DELETE | DeleteFulfillmentScriptShort | [DeleteFulfillmentScriptShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [DeleteFulfillmentScriptShort](../services-api/pkg/service/platform/fulfillmentScript.go) |
| `/platform/admin/fulfillment/scripts/{id}` | PATCH | UpdateFulfillmentScriptShort | [UpdateFulfillmentScriptShort](../platform-sdk/pkg/platformclient/fulfillment_script/fulfillment_script_client.go) | [UpdateFulfillmentScriptShort](../services-api/pkg/service/platform/fulfillmentScript.go) |

### Campaign Wrapper:  [Campaign](../services-api/pkg/service/platform/campaign.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/campaigns` | GET | QueryCampaignsShort | [QueryCampaignsShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [QueryCampaignsShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/campaigns` | POST | CreateCampaignShort | [CreateCampaignShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [CreateCampaignShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/campaigns/{campaignId}` | GET | GetCampaignShort | [GetCampaignShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [GetCampaignShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/campaigns/{campaignId}` | PUT | UpdateCampaignShort | [UpdateCampaignShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [UpdateCampaignShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/campaigns/{campaignId}/dynamic` | GET | GetCampaignDynamicShort | [GetCampaignDynamicShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [GetCampaignDynamicShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}` | GET | QueryCodesShort | [QueryCodesShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [QueryCodesShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}` | POST | CreateCodesShort | [CreateCodesShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [CreateCodesShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/codes.csv` | GET | DownloadShort | [DownloadShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [DownloadShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/disable/bulk` | PUT | BulkDisableCodesShort | [BulkDisableCodesShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [BulkDisableCodesShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/enable/bulk` | PUT | BulkEnableCodesShort | [BulkEnableCodesShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [BulkEnableCodesShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/history` | GET | QueryRedeemHistoryShort | [QueryRedeemHistoryShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [QueryRedeemHistoryShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/{code}` | GET | GetCodeShort | [GetCodeShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [GetCodeShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/{code}/disable` | PUT | DisableCodeShort | [DisableCodeShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [DisableCodeShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/codes/{code}/enable` | PUT | EnableCodeShort | [EnableCodeShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [EnableCodeShort](../services-api/pkg/service/platform/campaign.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/redemption` | POST | ApplyUserRedemptionShort | [ApplyUserRedemptionShort](../platform-sdk/pkg/platformclient/campaign/campaign_client.go) | [ApplyUserRedemptionShort](../services-api/pkg/service/platform/campaign.go) |

### Category Wrapper:  [Category](../services-api/pkg/service/platform/category.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/categories` | GET | GetRootCategoriesShort | [GetRootCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [GetRootCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories` | POST | CreateCategoryShort | [CreateCategoryShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [CreateCategoryShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/basic` | GET | ListCategoriesBasicShort | [ListCategoriesBasicShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [ListCategoriesBasicShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/{categoryPath}` | GET | GetCategoryShort | [GetCategoryShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [GetCategoryShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/{categoryPath}` | PUT | UpdateCategoryShort | [UpdateCategoryShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [UpdateCategoryShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/{categoryPath}` | DELETE | DeleteCategoryShort | [DeleteCategoryShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [DeleteCategoryShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/{categoryPath}/children` | GET | GetChildCategoriesShort | [GetChildCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [GetChildCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/admin/namespaces/{namespace}/categories/{categoryPath}/descendants` | GET | GetDescendantCategoriesShort | [GetDescendantCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [GetDescendantCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/public/namespaces/{namespace}/categories` | GET | PublicGetRootCategoriesShort | [PublicGetRootCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [PublicGetRootCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/public/namespaces/{namespace}/categories/download` | GET | DownloadCategoriesShort | [DownloadCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [DownloadCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/public/namespaces/{namespace}/categories/{categoryPath}` | GET | PublicGetCategoryShort | [PublicGetCategoryShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [PublicGetCategoryShort](../services-api/pkg/service/platform/category.go) |
| `/platform/public/namespaces/{namespace}/categories/{categoryPath}/children` | GET | PublicGetChildCategoriesShort | [PublicGetChildCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [PublicGetChildCategoriesShort](../services-api/pkg/service/platform/category.go) |
| `/platform/public/namespaces/{namespace}/categories/{categoryPath}/descendants` | GET | PublicGetDescendantCategoriesShort | [PublicGetDescendantCategoriesShort](../platform-sdk/pkg/platformclient/category/category_client.go) | [PublicGetDescendantCategoriesShort](../services-api/pkg/service/platform/category.go) |

### Currency Wrapper:  [Currency](../services-api/pkg/service/platform/currency.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/currencies` | GET | ListCurrenciesShort | [ListCurrenciesShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [ListCurrenciesShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/admin/namespaces/{namespace}/currencies` | POST | CreateCurrencyShort | [CreateCurrencyShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [CreateCurrencyShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/admin/namespaces/{namespace}/currencies/{currencyCode}` | PUT | UpdateCurrencyShort | [UpdateCurrencyShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [UpdateCurrencyShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/admin/namespaces/{namespace}/currencies/{currencyCode}` | DELETE | DeleteCurrencyShort | [DeleteCurrencyShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [DeleteCurrencyShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/admin/namespaces/{namespace}/currencies/{currencyCode}/config` | GET | GetCurrencyConfigShort | [GetCurrencyConfigShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [GetCurrencyConfigShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/admin/namespaces/{namespace}/currencies/{currencyCode}/summary` | GET | GetCurrencySummaryShort | [GetCurrencySummaryShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [GetCurrencySummaryShort](../services-api/pkg/service/platform/currency.go) |
| `/platform/public/namespaces/{namespace}/currencies` | GET | PublicListCurrenciesShort | [PublicListCurrenciesShort](../platform-sdk/pkg/platformclient/currency/currency_client.go) | [PublicListCurrenciesShort](../services-api/pkg/service/platform/currency.go) |

### DLC Wrapper:  [DLC](../services-api/pkg/service/platform/dlc.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/dlc/config/item` | GET | GetDLCItemConfigShort | [GetDLCItemConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [GetDLCItemConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/admin/namespaces/{namespace}/dlc/config/item` | PUT | UpdateDLCItemConfigShort | [UpdateDLCItemConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [UpdateDLCItemConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/admin/namespaces/{namespace}/dlc/config/item` | DELETE | DeleteDLCItemConfigShort | [DeleteDLCItemConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [DeleteDLCItemConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/admin/namespaces/{namespace}/dlc/config/platformMap` | GET | GetPlatformDLCConfigShort | [GetPlatformDLCConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [GetPlatformDLCConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/admin/namespaces/{namespace}/dlc/config/platformMap` | PUT | UpdatePlatformDLCConfigShort | [UpdatePlatformDLCConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [UpdatePlatformDLCConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/admin/namespaces/{namespace}/dlc/config/platformMap` | DELETE | DeletePlatformDLCConfigShort | [DeletePlatformDLCConfigShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [DeletePlatformDLCConfigShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/dlc/psn/sync` | PUT | PublicSyncPsnDlcInventoryShort | [PublicSyncPsnDlcInventoryShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [PublicSyncPsnDlcInventoryShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/dlc/steam/sync` | PUT | SyncSteamDLCShort | [SyncSteamDLCShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [SyncSteamDLCShort](../services-api/pkg/service/platform/dlc.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/dlc/xbl/sync` | PUT | SyncXboxDLCShort | [SyncXboxDLCShort](../platform-sdk/pkg/platformclient/dlc/dlc_client.go) | [SyncXboxDLCShort](../services-api/pkg/service/platform/dlc.go) |

### Entitlement Wrapper:  [Entitlement](../services-api/pkg/service/platform/entitlement.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/entitlements` | GET | QueryEntitlementsShort | [QueryEntitlementsShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [QueryEntitlementsShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/entitlements/{entitlementId}` | GET | GetEntitlementShort | [GetEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements` | GET | QueryUserEntitlementsShort | [QueryUserEntitlementsShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [QueryUserEntitlementsShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements` | POST | GrantUserEntitlementShort | [GrantUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GrantUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/byAppId` | GET | GetUserAppEntitlementByAppIdShort | [GetUserAppEntitlementByAppIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserAppEntitlementByAppIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/byAppType` | GET | QueryUserEntitlementsByAppTypeShort | [QueryUserEntitlementsByAppTypeShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [QueryUserEntitlementsByAppTypeShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/byItemId` | GET | GetUserEntitlementByItemIdShort | [GetUserEntitlementByItemIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementByItemIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/bySku` | GET | GetUserEntitlementBySkuShort | [GetUserEntitlementBySkuShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementBySkuShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/any` | GET | ExistsAnyUserActiveEntitlementShort | [ExistsAnyUserActiveEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [ExistsAnyUserActiveEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/anyOf` | GET | ExistsAnyUserActiveEntitlementByItemIdsShort | [ExistsAnyUserActiveEntitlementByItemIdsShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [ExistsAnyUserActiveEntitlementByItemIdsShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId` | GET | GetUserAppEntitlementOwnershipByAppIdShort | [GetUserAppEntitlementOwnershipByAppIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserAppEntitlementOwnershipByAppIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/byItemId` | GET | GetUserEntitlementOwnershipByItemIdShort | [GetUserEntitlementOwnershipByItemIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementOwnershipByItemIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/ownership/bySku` | GET | GetUserEntitlementOwnershipBySkuShort | [GetUserEntitlementOwnershipBySkuShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementOwnershipBySkuShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/revoke/byIds` | PUT | RevokeUserEntitlementsShort | [RevokeUserEntitlementsShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [RevokeUserEntitlementsShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}` | GET | GetUserEntitlementShort | [GetUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}` | PUT | UpdateUserEntitlementShort | [UpdateUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [UpdateUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/decrement` | PUT | ConsumeUserEntitlementShort | [ConsumeUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [ConsumeUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/disable` | PUT | DisableUserEntitlementShort | [DisableUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [DisableUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/enable` | PUT | EnableUserEntitlementShort | [EnableUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [EnableUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/history` | GET | GetUserEntitlementHistoriesShort | [GetUserEntitlementHistoriesShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [GetUserEntitlementHistoriesShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/revoke` | PUT | RevokeUserEntitlementShort | [RevokeUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [RevokeUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/me/entitlements/ownership/any` | GET | PublicExistsAnyMyActiveEntitlementShort | [PublicExistsAnyMyActiveEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicExistsAnyMyActiveEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/me/entitlements/ownership/byAppId` | GET | PublicGetMyAppEntitlementOwnershipByAppIdShort | [PublicGetMyAppEntitlementOwnershipByAppIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetMyAppEntitlementOwnershipByAppIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/me/entitlements/ownership/byItemId` | GET | PublicGetMyEntitlementOwnershipByItemIdShort | [PublicGetMyEntitlementOwnershipByItemIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetMyEntitlementOwnershipByItemIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/me/entitlements/ownership/bySku` | GET | PublicGetMyEntitlementOwnershipBySkuShort | [PublicGetMyEntitlementOwnershipBySkuShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetMyEntitlementOwnershipBySkuShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/me/entitlements/ownershipToken` | GET | PublicGetEntitlementOwnershipTokenShort | [PublicGetEntitlementOwnershipTokenShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetEntitlementOwnershipTokenShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements` | GET | PublicQueryUserEntitlementsShort | [PublicQueryUserEntitlementsShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicQueryUserEntitlementsShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/byAppId` | GET | PublicGetUserAppEntitlementByAppIdShort | [PublicGetUserAppEntitlementByAppIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserAppEntitlementByAppIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/byAppType` | GET | PublicQueryUserEntitlementsByAppTypeShort | [PublicQueryUserEntitlementsByAppTypeShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicQueryUserEntitlementsByAppTypeShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/byItemId` | GET | PublicGetUserEntitlementByItemIdShort | [PublicGetUserEntitlementByItemIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserEntitlementByItemIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/bySku` | GET | PublicGetUserEntitlementBySkuShort | [PublicGetUserEntitlementBySkuShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserEntitlementBySkuShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/ownership/any` | GET | PublicExistsAnyUserActiveEntitlementShort | [PublicExistsAnyUserActiveEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicExistsAnyUserActiveEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId` | GET | PublicGetUserAppEntitlementOwnershipByAppIdShort | [PublicGetUserAppEntitlementOwnershipByAppIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserAppEntitlementOwnershipByAppIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/ownership/byItemId` | GET | PublicGetUserEntitlementOwnershipByItemIdShort | [PublicGetUserEntitlementOwnershipByItemIdShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserEntitlementOwnershipByItemIdShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/ownership/bySku` | GET | PublicGetUserEntitlementOwnershipBySkuShort | [PublicGetUserEntitlementOwnershipBySkuShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserEntitlementOwnershipBySkuShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}` | GET | PublicGetUserEntitlementShort | [PublicGetUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicGetUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/entitlements/{entitlementId}/decrement` | PUT | PublicConsumeUserEntitlementShort | [PublicConsumeUserEntitlementShort](../platform-sdk/pkg/platformclient/entitlement/entitlement_client.go) | [PublicConsumeUserEntitlementShort](../services-api/pkg/service/platform/entitlement.go) |

### Fulfillment Wrapper:  [Fulfillment](../services-api/pkg/service/platform/fulfillment.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/fulfillment/history` | GET | QueryFulfillmentHistoriesShort | [QueryFulfillmentHistoriesShort](../platform-sdk/pkg/platformclient/fulfillment/fulfillment_client.go) | [QueryFulfillmentHistoriesShort](../services-api/pkg/service/platform/fulfillment.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/fulfillment` | POST | FulfillItemShort | [FulfillItemShort](../platform-sdk/pkg/platformclient/fulfillment/fulfillment_client.go) | [FulfillItemShort](../services-api/pkg/service/platform/fulfillment.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/fulfillment/code` | POST | RedeemCodeShort | [RedeemCodeShort](../platform-sdk/pkg/platformclient/fulfillment/fulfillment_client.go) | [RedeemCodeShort](../services-api/pkg/service/platform/fulfillment.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/fulfillment/rewards` | POST | FulfillRewardsShort | [FulfillRewardsShort](../platform-sdk/pkg/platformclient/fulfillment/fulfillment_client.go) | [FulfillRewardsShort](../services-api/pkg/service/platform/fulfillment.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/fulfillment/code` | POST | PublicRedeemCodeShort | [PublicRedeemCodeShort](../platform-sdk/pkg/platformclient/fulfillment/fulfillment_client.go) | [PublicRedeemCodeShort](../services-api/pkg/service/platform/fulfillment.go) |

### IAP Wrapper:  [IAP](../services-api/pkg/service/platform/iap.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/iap/config/apple` | GET | GetAppleIAPConfigShort | [GetAppleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetAppleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/apple` | PUT | UpdateAppleIAPConfigShort | [UpdateAppleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateAppleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/apple` | DELETE | DeleteAppleIAPConfigShort | [DeleteAppleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteAppleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/epicgames` | GET | GetEpicGamesIAPConfigShort | [GetEpicGamesIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetEpicGamesIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/epicgames` | PUT | UpdateEpicGamesIAPConfigShort | [UpdateEpicGamesIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateEpicGamesIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/epicgames` | DELETE | DeleteEpicGamesIAPConfigShort | [DeleteEpicGamesIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteEpicGamesIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/google` | GET | GetGoogleIAPConfigShort | [GetGoogleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetGoogleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/google` | PUT | UpdateGoogleIAPConfigShort | [UpdateGoogleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateGoogleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/google` | DELETE | DeleteGoogleIAPConfigShort | [DeleteGoogleIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteGoogleIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/google/cert` | PUT | UpdateGoogleP12FileShort | [UpdateGoogleP12FileShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateGoogleP12FileShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/item` | GET | GetIAPItemConfigShort | [GetIAPItemConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetIAPItemConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/item` | PUT | UpdateIAPItemConfigShort | [UpdateIAPItemConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateIAPItemConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/item` | DELETE | DeleteIAPItemConfigShort | [DeleteIAPItemConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteIAPItemConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/playstation` | GET | GetPlayStationIAPConfigShort | [GetPlayStationIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetPlayStationIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/playstation` | PUT | UpdatePlaystationIAPConfigShort | [UpdatePlaystationIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdatePlaystationIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/playstation` | DELETE | DeletePlaystationIAPConfigShort | [DeletePlaystationIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeletePlaystationIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/stadia` | GET | GetStadiaIAPConfigShort | [GetStadiaIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetStadiaIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/stadia` | DELETE | DeleteStadiaIAPConfigShort | [DeleteStadiaIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteStadiaIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/stadia/cert` | PUT | UpdateStadiaJsonConfigFileShort | [UpdateStadiaJsonConfigFileShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateStadiaJsonConfigFileShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/steam` | GET | GetSteamIAPConfigShort | [GetSteamIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetSteamIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/steam` | PUT | UpdateSteamIAPConfigShort | [UpdateSteamIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateSteamIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/steam` | DELETE | DeleteSteamIAPConfigShort | [DeleteSteamIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteSteamIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/twitch` | GET | GetTwitchIAPConfigShort | [GetTwitchIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetTwitchIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/twitch` | PUT | UpdateTwitchIAPConfigShort | [UpdateTwitchIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateTwitchIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/twitch` | DELETE | DeleteTwitchIAPConfigShort | [DeleteTwitchIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteTwitchIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/xbl` | GET | GetXblIAPConfigShort | [GetXblIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [GetXblIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/xbl` | PUT | UpdateXblIAPConfigShort | [UpdateXblIAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateXblIAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/xbl` | DELETE | DeleteXblAPConfigShort | [DeleteXblAPConfigShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [DeleteXblAPConfigShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/iap/config/xbl/cert` | PUT | UpdateXblBPCertFileShort | [UpdateXblBPCertFileShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [UpdateXblBPCertFileShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/iap` | GET | QueryUserIAPOrdersShort | [QueryUserIAPOrdersShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [QueryUserIAPOrdersShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/iap/all` | GET | QueryAllUserIAPOrdersShort | [QueryAllUserIAPOrdersShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [QueryAllUserIAPOrdersShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/iap/mock/receipt` | PUT | MockFulfillIAPItemShort | [MockFulfillIAPItemShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [MockFulfillIAPItemShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/apple/receipt` | PUT | PublicFulfillAppleIAPItemShort | [PublicFulfillAppleIAPItemShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [PublicFulfillAppleIAPItemShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/epicgames/sync` | PUT | SyncEpicGamesInventoryShort | [SyncEpicGamesInventoryShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [SyncEpicGamesInventoryShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/google/receipt` | PUT | PublicFulfillGoogleIAPItemShort | [PublicFulfillGoogleIAPItemShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [PublicFulfillGoogleIAPItemShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/psn/sync` | PUT | PublicReconcilePlayStationStoreShort | [PublicReconcilePlayStationStoreShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [PublicReconcilePlayStationStoreShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/stadia/sync` | PUT | SyncStadiaEntitlementShort | [SyncStadiaEntitlementShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [SyncStadiaEntitlementShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/steam/sync` | PUT | SyncSteamInventoryShort | [SyncSteamInventoryShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [SyncSteamInventoryShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/twitch/sync` | PUT | SyncTwitchDropsEntitlementShort | [SyncTwitchDropsEntitlementShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [SyncTwitchDropsEntitlementShort](../services-api/pkg/service/platform/iap.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/iap/xbl/sync` | PUT | SyncXboxInventoryShort | [SyncXboxInventoryShort](../platform-sdk/pkg/platformclient/iap/iap_client.go) | [SyncXboxInventoryShort](../services-api/pkg/service/platform/iap.go) |

### Item Wrapper:  [Item](../services-api/pkg/service/platform/item.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/items` | PUT | SyncInGameItemShort | [SyncInGameItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [SyncInGameItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items` | POST | CreateItemShort | [CreateItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [CreateItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/byAppId` | GET | GetItemByAppIdShort | [GetItemByAppIdShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetItemByAppIdShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/byCriteria` | GET | QueryItemsShort | [QueryItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [QueryItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/byFeatures/basic` | GET | ListBasicItemsByFeaturesShort | [ListBasicItemsByFeaturesShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [ListBasicItemsByFeaturesShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/bySku` | GET | GetItemBySkuShort | [GetItemBySkuShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetItemBySkuShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/bySku/locale` | GET | GetLocaleItemBySkuShort | [GetLocaleItemBySkuShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetLocaleItemBySkuShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/itemId/bySku` | GET | GetItemIdBySkuShort | [GetItemIdBySkuShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetItemIdBySkuShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/locale/byIds` | GET | BulkGetLocaleItemsShort | [BulkGetLocaleItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [BulkGetLocaleItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/search` | GET | SearchItemsShort | [SearchItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [SearchItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/uncategorized` | GET | QueryUncategorizedItemsShort | [QueryUncategorizedItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [QueryUncategorizedItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}` | GET | GetItemShort | [GetItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}` | PUT | UpdateItemShort | [UpdateItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [UpdateItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}` | DELETE | DeleteItemShort | [DeleteItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [DeleteItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/acquire` | PUT | AcquireItemShort | [AcquireItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [AcquireItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/app` | GET | GetAppShort | [GetAppShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetAppShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/app` | PUT | UpdateAppShort | [UpdateAppShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [UpdateAppShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/disable` | PUT | DisableItemShort | [DisableItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [DisableItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/dynamic` | GET | GetItemDynamicDataShort | [GetItemDynamicDataShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetItemDynamicDataShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/enable` | PUT | EnableItemShort | [EnableItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [EnableItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/features/{feature}` | PUT | FeatureItemShort | [FeatureItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [FeatureItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/features/{feature}` | DELETE | DefeatureItemShort | [DefeatureItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [DefeatureItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/locale` | GET | GetLocaleItemShort | [GetLocaleItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [GetLocaleItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/admin/namespaces/{namespace}/items/{itemId}/return` | PUT | ReturnItemShort | [ReturnItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [ReturnItemShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/byAppId` | GET | PublicGetItemByAppIdShort | [PublicGetItemByAppIdShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicGetItemByAppIdShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/byCriteria` | GET | PublicQueryItemsShort | [PublicQueryItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicQueryItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/bySku` | GET | PublicGetItemBySkuShort | [PublicGetItemBySkuShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicGetItemBySkuShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/locale/byIds` | GET | PublicBulkGetItemsShort | [PublicBulkGetItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicBulkGetItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/search` | GET | PublicSearchItemsShort | [PublicSearchItemsShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicSearchItemsShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/{itemId}/app/locale` | GET | PublicGetAppShort | [PublicGetAppShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicGetAppShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/{itemId}/dynamic` | GET | PublicGetItemDynamicDataShort | [PublicGetItemDynamicDataShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicGetItemDynamicDataShort](../services-api/pkg/service/platform/item.go) |
| `/platform/public/namespaces/{namespace}/items/{itemId}/locale` | GET | PublicGetItemShort | [PublicGetItemShort](../platform-sdk/pkg/platformclient/item/item_client.go) | [PublicGetItemShort](../services-api/pkg/service/platform/item.go) |

### KeyGroup Wrapper:  [KeyGroup](../services-api/pkg/service/platform/keyGroup.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/keygroups` | GET | QueryKeyGroupsShort | [QueryKeyGroupsShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [QueryKeyGroupsShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups` | POST | CreateKeyGroupShort | [CreateKeyGroupShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [CreateKeyGroupShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}` | GET | GetKeyGroupShort | [GetKeyGroupShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [GetKeyGroupShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}` | PUT | UpdateKeyGroupShort | [UpdateKeyGroupShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [UpdateKeyGroupShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}/dynamic` | GET | GetKeyGroupDynamicShort | [GetKeyGroupDynamicShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [GetKeyGroupDynamicShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}/keys` | GET | ListKeysShort | [ListKeysShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [ListKeysShort](../services-api/pkg/service/platform/keyGroup.go) |
| `/platform/admin/namespaces/{namespace}/keygroups/{keyGroupId}/keys` | POST | UploadKeysShort | [UploadKeysShort](../platform-sdk/pkg/platformclient/key_group/key_group_client.go) | [UploadKeysShort](../services-api/pkg/service/platform/keyGroup.go) |

### Order Wrapper:  [Order](../services-api/pkg/service/platform/order.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/orders` | GET | QueryOrdersShort | [QueryOrdersShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [QueryOrdersShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/orders/stats` | GET | GetOrderStatisticsShort | [GetOrderStatisticsShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [GetOrderStatisticsShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/orders/{orderNo}` | GET | GetOrderShort | [GetOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [GetOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/orders/{orderNo}/refund` | PUT | RefundOrderShort | [RefundOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [RefundOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders` | GET | QueryUserOrdersShort | [QueryUserOrdersShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [QueryUserOrdersShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/countOfItem` | GET | CountOfPurchasedItemShort | [CountOfPurchasedItemShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [CountOfPurchasedItemShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}` | GET | GetUserOrderShort | [GetUserOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [GetUserOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}` | PUT | UpdateUserOrderStatusShort | [UpdateUserOrderStatusShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [UpdateUserOrderStatusShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/fulfill` | PUT | FulfillUserOrderShort | [FulfillUserOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [FulfillUserOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/grant` | GET | GetUserOrderGrantShort | [GetUserOrderGrantShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [GetUserOrderGrantShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/history` | GET | GetUserOrderHistoriesShort | [GetUserOrderHistoriesShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [GetUserOrderHistoriesShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/notifications` | POST | ProcessUserOrderNotificationShort | [ProcessUserOrderNotificationShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [ProcessUserOrderNotificationShort](../services-api/pkg/service/platform/order.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf` | GET | DownloadUserOrderReceiptShort | [DownloadUserOrderReceiptShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [DownloadUserOrderReceiptShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders` | GET | PublicQueryUserOrdersShort | [PublicQueryUserOrdersShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicQueryUserOrdersShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders` | POST | PublicCreateUserOrderShort | [PublicCreateUserOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicCreateUserOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}` | GET | PublicGetUserOrderShort | [PublicGetUserOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicGetUserOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/cancel` | PUT | PublicCancelUserOrderShort | [PublicCancelUserOrderShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicCancelUserOrderShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/history` | GET | PublicGetUserOrderHistoriesShort | [PublicGetUserOrderHistoriesShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicGetUserOrderHistoriesShort](../services-api/pkg/service/platform/order.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/orders/{orderNo}/receipt.pdf` | GET | PublicDownloadUserOrderReceiptShort | [PublicDownloadUserOrderReceiptShort](../platform-sdk/pkg/platformclient/order/order_client.go) | [PublicDownloadUserOrderReceiptShort](../services-api/pkg/service/platform/order.go) |

### PaymentCallbackConfig Wrapper:  [PaymentCallbackConfig](../services-api/pkg/service/platform/paymentCallbackConfig.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/payment/config/callback` | GET | GetPaymentCallbackConfigShort | [GetPaymentCallbackConfigShort](../platform-sdk/pkg/platformclient/payment_callback_config/payment_callback_config_client.go) | [GetPaymentCallbackConfigShort](../services-api/pkg/service/platform/paymentCallbackConfig.go) |
| `/platform/admin/namespaces/{namespace}/payment/config/callback` | PUT | UpdatePaymentCallbackConfigShort | [UpdatePaymentCallbackConfigShort](../platform-sdk/pkg/platformclient/payment_callback_config/payment_callback_config_client.go) | [UpdatePaymentCallbackConfigShort](../services-api/pkg/service/platform/paymentCallbackConfig.go) |

### Payment Wrapper:  [Payment](../services-api/pkg/service/platform/payment.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/payment/notifications` | GET | QueryPaymentNotificationsShort | [QueryPaymentNotificationsShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [QueryPaymentNotificationsShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders` | GET | QueryPaymentOrdersShort | [QueryPaymentOrdersShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [QueryPaymentOrdersShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/byExtTxId` | GET | ListExtOrderNoByExtTxIdShort | [ListExtOrderNoByExtTxIdShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [ListExtOrderNoByExtTxIdShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}` | GET | GetPaymentOrderShort | [GetPaymentOrderShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [GetPaymentOrderShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}` | PUT | ChargePaymentOrderShort | [ChargePaymentOrderShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [ChargePaymentOrderShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}/simulate-notification` | PUT | SimulatePaymentOrderNotificationShort | [SimulatePaymentOrderNotificationShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [SimulatePaymentOrderNotificationShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}/status` | GET | GetPaymentOrderChargeStatusShort | [GetPaymentOrderChargeStatusShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [GetPaymentOrderChargeStatusShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/payment/orders` | POST | CreateUserPaymentOrderShort | [CreateUserPaymentOrderShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [CreateUserPaymentOrderShort](../services-api/pkg/service/platform/payment.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/payment/orders/{paymentOrderNo}/refund` | PUT | RefundUserPaymentOrderShort | [RefundUserPaymentOrderShort](../platform-sdk/pkg/platformclient/payment/payment_client.go) | [RefundUserPaymentOrderShort](../services-api/pkg/service/platform/payment.go) |

### Payment(Dedicated) Wrapper:  [PaymentDedicated](../services-api/pkg/service/platform/paymentDedicated.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/payment/orders` | POST | CreatePaymentOrderByDedicatedShort | [CreatePaymentOrderByDedicatedShort](../platform-sdk/pkg/platformclient/payment_dedicated/payment_dedicated_client.go) | [CreatePaymentOrderByDedicatedShort](../services-api/pkg/service/platform/paymentDedicated.go) |
| `/platform/admin/namespaces/{namespace}/payment/orders/{paymentOrderNo}/refund` | PUT | RefundPaymentOrderByDedicatedShort | [RefundPaymentOrderByDedicatedShort](../platform-sdk/pkg/platformclient/payment_dedicated/payment_dedicated_client.go) | [RefundPaymentOrderByDedicatedShort](../services-api/pkg/service/platform/paymentDedicated.go) |
| `/platform/admin/payment/orders` | GET | SyncPaymentOrdersShort | [SyncPaymentOrdersShort](../platform-sdk/pkg/platformclient/payment_dedicated/payment_dedicated_client.go) | [SyncPaymentOrdersShort](../services-api/pkg/service/platform/paymentDedicated.go) |

### Reward Wrapper:  [Reward](../services-api/pkg/service/platform/reward.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/rewards` | POST | CreateRewardShort | [CreateRewardShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [CreateRewardShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/byCriteria` | GET | QueryRewardsShort | [QueryRewardsShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [QueryRewardsShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/export` | GET | ExportRewardsShort | [ExportRewardsShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [ExportRewardsShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/import` | POST | ImportRewardsShort | [ImportRewardsShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [ImportRewardsShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/{rewardId}` | GET | GetRewardShort | [GetRewardShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [GetRewardShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/{rewardId}` | PUT | UpdateRewardShort | [UpdateRewardShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [UpdateRewardShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/{rewardId}` | DELETE | DeleteRewardShort | [DeleteRewardShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [DeleteRewardShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/admin/namespaces/{namespace}/rewards/{rewardId}/match` | PUT | CheckEventConditionShort | [CheckEventConditionShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [CheckEventConditionShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/public/namespaces/{namespace}/rewards/byCode` | GET | GetRewardByCodeShort | [GetRewardByCodeShort](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [GetRewardByCodeShort](../services-api/pkg/service/platform/reward.go) |
| `/platform/public/namespaces/{namespace}/rewards/byCriteria` | GET | QueryRewards1Short | [QueryRewards1Short](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [QueryRewards1Short](../services-api/pkg/service/platform/reward.go) |
| `/platform/public/namespaces/{namespace}/rewards/{rewardId}` | GET | GetReward1Short | [GetReward1Short](../platform-sdk/pkg/platformclient/reward/reward_client.go) | [GetReward1Short](../services-api/pkg/service/platform/reward.go) |

### Store Wrapper:  [Store](../services-api/pkg/service/platform/store.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/stores` | GET | ListStoresShort | [ListStoresShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [ListStoresShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores` | POST | CreateStoreShort | [CreateStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [CreateStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/import` | PUT | ImportStoreShort | [ImportStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [ImportStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/published` | GET | GetPublishedStoreShort | [GetPublishedStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [GetPublishedStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/published` | DELETE | DeletePublishedStoreShort | [DeletePublishedStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [DeletePublishedStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/published/backup` | GET | GetPublishedStoreBackupShort | [GetPublishedStoreBackupShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [GetPublishedStoreBackupShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/published/rollback` | PUT | RollbackPublishedStoreShort | [RollbackPublishedStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [RollbackPublishedStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/{storeId}` | GET | GetStoreShort | [GetStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [GetStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/{storeId}` | PUT | UpdateStoreShort | [UpdateStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [UpdateStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/{storeId}` | DELETE | DeleteStoreShort | [DeleteStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [DeleteStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/{storeId}/clone` | PUT | CloneStoreShort | [CloneStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [CloneStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/admin/namespaces/{namespace}/stores/{storeId}/export` | GET | ExportStoreShort | [ExportStoreShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [ExportStoreShort](../services-api/pkg/service/platform/store.go) |
| `/platform/public/namespaces/{namespace}/stores` | GET | PublicListStoresShort | [PublicListStoresShort](../platform-sdk/pkg/platformclient/store/store_client.go) | [PublicListStoresShort](../services-api/pkg/service/platform/store.go) |

### Subscription Wrapper:  [Subscription](../services-api/pkg/service/platform/subscription.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/subscriptions` | GET | QuerySubscriptionsShort | [QuerySubscriptionsShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [QuerySubscriptionsShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/subscriptions/{subscriptionId}/recurring` | PUT | RecurringChargeSubscriptionShort | [RecurringChargeSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [RecurringChargeSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions` | GET | QueryUserSubscriptionsShort | [QueryUserSubscriptionsShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [QueryUserSubscriptionsShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/activities` | GET | GetUserSubscriptionActivitiesShort | [GetUserSubscriptionActivitiesShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [GetUserSubscriptionActivitiesShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/platformSubscribe` | POST | PlatformSubscribeSubscriptionShort | [PlatformSubscribeSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PlatformSubscribeSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId` | GET | CheckUserSubscriptionSubscribableByItemIdShort | [CheckUserSubscriptionSubscribableByItemIdShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [CheckUserSubscriptionSubscribableByItemIdShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}` | GET | GetUserSubscriptionShort | [GetUserSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [GetUserSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}` | DELETE | DeleteUserSubscriptionShort | [DeleteUserSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [DeleteUserSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/cancel` | PUT | CancelSubscriptionShort | [CancelSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [CancelSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/grant` | PUT | GrantDaysToSubscriptionShort | [GrantDaysToSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [GrantDaysToSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history` | GET | GetUserSubscriptionBillingHistoriesShort | [GetUserSubscriptionBillingHistoriesShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [GetUserSubscriptionBillingHistoriesShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/notifications` | POST | ProcessUserSubscriptionNotificationShort | [ProcessUserSubscriptionNotificationShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [ProcessUserSubscriptionNotificationShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions` | GET | PublicQueryUserSubscriptionsShort | [PublicQueryUserSubscriptionsShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicQueryUserSubscriptionsShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions` | POST | PublicSubscribeSubscriptionShort | [PublicSubscribeSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicSubscribeSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId` | GET | PublicCheckUserSubscriptionSubscribableByItemIdShort | [PublicCheckUserSubscriptionSubscribableByItemIdShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicCheckUserSubscriptionSubscribableByItemIdShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}` | GET | PublicGetUserSubscriptionShort | [PublicGetUserSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicGetUserSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/billingAccount` | PUT | PublicChangeSubscriptionBillingAccountShort | [PublicChangeSubscriptionBillingAccountShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicChangeSubscriptionBillingAccountShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/cancel` | PUT | PublicCancelSubscriptionShort | [PublicCancelSubscriptionShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicCancelSubscriptionShort](../services-api/pkg/service/platform/subscription.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/subscriptions/{subscriptionId}/history` | GET | PublicGetUserSubscriptionBillingHistoriesShort | [PublicGetUserSubscriptionBillingHistoriesShort](../platform-sdk/pkg/platformclient/subscription/subscription_client.go) | [PublicGetUserSubscriptionBillingHistoriesShort](../services-api/pkg/service/platform/subscription.go) |

### Ticket Wrapper:  [Ticket](../services-api/pkg/service/platform/ticket.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/tickets/{boothName}` | GET | GetTicketDynamicShort | [GetTicketDynamicShort](../platform-sdk/pkg/platformclient/ticket/ticket_client.go) | [GetTicketDynamicShort](../services-api/pkg/service/platform/ticket.go) |
| `/platform/admin/namespaces/{namespace}/tickets/{boothName}/decrement` | PUT | DecreaseTicketSaleShort | [DecreaseTicketSaleShort](../platform-sdk/pkg/platformclient/ticket/ticket_client.go) | [DecreaseTicketSaleShort](../services-api/pkg/service/platform/ticket.go) |
| `/platform/admin/namespaces/{namespace}/tickets/{boothName}/id` | GET | GetTicketBoothIDShort | [GetTicketBoothIDShort](../platform-sdk/pkg/platformclient/ticket/ticket_client.go) | [GetTicketBoothIDShort](../services-api/pkg/service/platform/ticket.go) |
| `/platform/admin/namespaces/{namespace}/tickets/{boothName}/increment` | PUT | IncreaseTicketSaleShort | [IncreaseTicketSaleShort](../platform-sdk/pkg/platformclient/ticket/ticket_client.go) | [IncreaseTicketSaleShort](../services-api/pkg/service/platform/ticket.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/tickets/{boothName}` | POST | AcquireUserTicketShort | [AcquireUserTicketShort](../platform-sdk/pkg/platformclient/ticket/ticket_client.go) | [AcquireUserTicketShort](../services-api/pkg/service/platform/ticket.go) |

### Anonymization Wrapper:  [Anonymization](../services-api/pkg/service/platform/anonymization.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/campaign` | DELETE | AnonymizeCampaignShort | [AnonymizeCampaignShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeCampaignShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/entitlements` | DELETE | AnonymizeEntitlementShort | [AnonymizeEntitlementShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeEntitlementShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/fulfillment` | DELETE | AnonymizeFulfillmentShort | [AnonymizeFulfillmentShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeFulfillmentShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/integrations` | DELETE | AnonymizeIntegrationShort | [AnonymizeIntegrationShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeIntegrationShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/orders` | DELETE | AnonymizeOrderShort | [AnonymizeOrderShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeOrderShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/payment` | DELETE | AnonymizePaymentShort | [AnonymizePaymentShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizePaymentShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/subscriptions` | DELETE | AnonymizeSubscriptionShort | [AnonymizeSubscriptionShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeSubscriptionShort](../services-api/pkg/service/platform/anonymization.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/anonymization/wallets` | DELETE | AnonymizeWalletShort | [AnonymizeWalletShort](../platform-sdk/pkg/platformclient/anonymization/anonymization_client.go) | [AnonymizeWalletShort](../services-api/pkg/service/platform/anonymization.go) |

### Wallet Wrapper:  [Wallet](../services-api/pkg/service/platform/wallet.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/check` | GET | CheckWalletShort | [CheckWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [CheckWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/credit` | PUT | CreditUserWalletShort | [CreditUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [CreditUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/payment` | PUT | PayWithUserWalletShort | [PayWithUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [PayWithUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{walletId}` | GET | GetUserWalletShort | [GetUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [GetUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{walletId}/debit` | PUT | DebitUserWalletShort | [DebitUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [DebitUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{walletId}/disable` | PUT | DisableUserWalletShort | [DisableUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [DisableUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{walletId}/enable` | PUT | EnableUserWalletShort | [EnableUserWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [EnableUserWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/users/{userId}/wallets/{walletId}/transactions` | GET | ListUserWalletTransactionsShort | [ListUserWalletTransactionsShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [ListUserWalletTransactionsShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/wallets` | GET | QueryWalletsShort | [QueryWalletsShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [QueryWalletsShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/admin/namespaces/{namespace}/wallets/{walletId}` | GET | GetWalletShort | [GetWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [GetWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/public/namespaces/{namespace}/users/me/wallets/{currencyCode}` | GET | PublicGetMyWalletShort | [PublicGetMyWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [PublicGetMyWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}` | GET | PublicGetWalletShort | [PublicGetWalletShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [PublicGetWalletShort](../services-api/pkg/service/platform/wallet.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/transactions` | GET | PublicListUserWalletTransactionsShort | [PublicListUserWalletTransactionsShort](../platform-sdk/pkg/platformclient/wallet/wallet_client.go) | [PublicListUserWalletTransactionsShort](../services-api/pkg/service/platform/wallet.go) |

### Order(Dedicated) Wrapper:  [OrderDedicated](../services-api/pkg/service/platform/orderDedicated.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/orders` | GET | SyncOrdersShort | [SyncOrdersShort](../platform-sdk/pkg/platformclient/order_dedicated/order_dedicated_client.go) | [SyncOrdersShort](../services-api/pkg/service/platform/orderDedicated.go) |

### PaymentConfig Wrapper:  [PaymentConfig](../services-api/pkg/service/platform/paymentConfig.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/admin/payment/config/merchant/adyenconfig/test` | POST | TestAdyenConfigShort | [TestAdyenConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestAdyenConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/alipayconfig/test` | POST | TestAliPayConfigShort | [TestAliPayConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestAliPayConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/checkoutconfig/test` | POST | TestCheckoutConfigShort | [TestCheckoutConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestCheckoutConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/matched` | GET | DebugMatchedPaymentMerchantConfigShort | [DebugMatchedPaymentMerchantConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [DebugMatchedPaymentMerchantConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/paypalconfig/test` | POST | TestPayPalConfigShort | [TestPayPalConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestPayPalConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/stripeconfig/test` | POST | TestStripeConfigShort | [TestStripeConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestStripeConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/wxpayconfig/test` | POST | TestWxPayConfigShort | [TestWxPayConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestWxPayConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/xsollaconfig/test` | POST | TestXsollaConfigShort | [TestXsollaConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestXsollaConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}` | GET | GetPaymentMerchantConfigShort | [GetPaymentMerchantConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [GetPaymentMerchantConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/adyenconfig` | PUT | UpdateAdyenConfigShort | [UpdateAdyenConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateAdyenConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/adyenconfig/test` | GET | TestAdyenConfigByIdShort | [TestAdyenConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestAdyenConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/alipayconfig` | PUT | UpdateAliPayConfigShort | [UpdateAliPayConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateAliPayConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/alipayconfig/test` | GET | TestAliPayConfigByIdShort | [TestAliPayConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestAliPayConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/checkoutconfig` | PUT | UpdateCheckoutConfigShort | [UpdateCheckoutConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateCheckoutConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/checkoutconfig/test` | GET | TestCheckoutConfigByIdShort | [TestCheckoutConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestCheckoutConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/paypalconfig` | PUT | UpdatePayPalConfigShort | [UpdatePayPalConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdatePayPalConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/paypalconfig/test` | GET | TestPayPalConfigByIdShort | [TestPayPalConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestPayPalConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/stripeconfig` | PUT | UpdateStripeConfigShort | [UpdateStripeConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateStripeConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/stripeconfig/test` | GET | TestStripeConfigByIdShort | [TestStripeConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestStripeConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/wxpayconfig` | PUT | UpdateWxPayConfigShort | [UpdateWxPayConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateWxPayConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/wxpayconfig/cert` | PUT | UpdateWxPayConfigCertShort | [UpdateWxPayConfigCertShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateWxPayConfigCertShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/wxpayconfig/test` | GET | TestWxPayConfigByIdShort | [TestWxPayConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestWxPayConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/xsollaconfig` | PUT | UpdateXsollaConfigShort | [UpdateXsollaConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateXsollaConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/xsollaconfig/test` | GET | TestXsollaConfigByIdShort | [TestXsollaConfigByIdShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [TestXsollaConfigByIdShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/merchant/{id}/xsollauiconfig` | PUT | UpdateXsollaUIConfigShort | [UpdateXsollaUIConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdateXsollaUIConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider` | GET | QueryPaymentProviderConfigShort | [QueryPaymentProviderConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [QueryPaymentProviderConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider` | POST | CreatePaymentProviderConfigShort | [CreatePaymentProviderConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [CreatePaymentProviderConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider/aggregate` | GET | GetAggregatePaymentProvidersShort | [GetAggregatePaymentProvidersShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [GetAggregatePaymentProvidersShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider/matched` | GET | DebugMatchedPaymentProviderConfigShort | [DebugMatchedPaymentProviderConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [DebugMatchedPaymentProviderConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider/special` | GET | GetSpecialPaymentProvidersShort | [GetSpecialPaymentProvidersShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [GetSpecialPaymentProvidersShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider/{id}` | PUT | UpdatePaymentProviderConfigShort | [UpdatePaymentProviderConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdatePaymentProviderConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/provider/{id}` | DELETE | DeletePaymentProviderConfigShort | [DeletePaymentProviderConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [DeletePaymentProviderConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/tax` | GET | GetPaymentTaxConfigShort | [GetPaymentTaxConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [GetPaymentTaxConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |
| `/platform/admin/payment/config/tax` | PUT | UpdatePaymentTaxConfigShort | [UpdatePaymentTaxConfigShort](../platform-sdk/pkg/platformclient/payment_config/payment_config_client.go) | [UpdatePaymentTaxConfigShort](../services-api/pkg/service/platform/paymentConfig.go) |

### PaymentStation Wrapper:  [PaymentStation](../services-api/pkg/service/platform/paymentStation.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/public/namespaces/{namespace}/payment/customization` | GET | GetPaymentCustomizationShort | [GetPaymentCustomizationShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [GetPaymentCustomizationShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/link` | POST | PublicGetPaymentUrlShort | [PublicGetPaymentUrlShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicGetPaymentUrlShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/methods` | GET | PublicGetPaymentMethodsShort | [PublicGetPaymentMethodsShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicGetPaymentMethodsShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/orders/{paymentOrderNo}/info` | GET | PublicGetUnpaidPaymentOrderShort | [PublicGetUnpaidPaymentOrderShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicGetUnpaidPaymentOrderShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/orders/{paymentOrderNo}/pay` | POST | PayShort | [PayShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PayShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/orders/{paymentOrderNo}/status` | GET | PublicCheckPaymentOrderPaidStatusShort | [PublicCheckPaymentOrderPaidStatusShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicCheckPaymentOrderPaidStatusShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/publicconfig` | GET | GetPaymentPublicConfigShort | [GetPaymentPublicConfigShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [GetPaymentPublicConfigShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/qrcode` | GET | PublicGetQRCodeShort | [PublicGetQRCodeShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicGetQRCodeShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/returnurl` | GET | PublicNormalizePaymentReturnUrlShort | [PublicNormalizePaymentReturnUrlShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [PublicNormalizePaymentReturnUrlShort](../services-api/pkg/service/platform/paymentStation.go) |
| `/platform/public/namespaces/{namespace}/payment/tax` | GET | GetPaymentTaxValueShort | [GetPaymentTaxValueShort](../platform-sdk/pkg/platformclient/payment_station/payment_station_client.go) | [GetPaymentTaxValueShort](../services-api/pkg/service/platform/paymentStation.go) |

### PaymentAccount Wrapper:  [PaymentAccount](../services-api/pkg/service/platform/paymentAccount.go)
| Endpoint | Method | ID | Class | Wrapper |
|---|---|---|---|---|
| `/platform/public/namespaces/{namespace}/users/{userId}/payment/accounts` | GET | PublicGetPaymentAccountsShort | [PublicGetPaymentAccountsShort](../platform-sdk/pkg/platformclient/payment_account/payment_account_client.go) | [PublicGetPaymentAccountsShort](../services-api/pkg/service/platform/paymentAccount.go) |
| `/platform/public/namespaces/{namespace}/users/{userId}/payment/accounts/{type}/{id}` | DELETE | PublicDeletePaymentAccountShort | [PublicDeletePaymentAccountShort](../platform-sdk/pkg/platformclient/payment_account/payment_account_client.go) | [PublicDeletePaymentAccountShort](../services-api/pkg/service/platform/paymentAccount.go) |


&nbsp;  

## Models

| Model Struct | Class |
|---|---|
| `AdditionalData` | [AdditionalData ](../platform-sdk/pkg/platformclientmodels/additional_data.go) |
| `AdyenConfig` | [AdyenConfig ](../platform-sdk/pkg/platformclientmodels/adyen_config.go) |
| `AliPayConfig` | [AliPayConfig ](../platform-sdk/pkg/platformclientmodels/ali_pay_config.go) |
| `AppEntitlementInfo` | [AppEntitlementInfo ](../platform-sdk/pkg/platformclientmodels/app_entitlement_info.go) |
| `AppEntitlementPagingSlicedResult` | [AppEntitlementPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/app_entitlement_paging_sliced_result.go) |
| `AppInfo` | [AppInfo ](../platform-sdk/pkg/platformclientmodels/app_info.go) |
| `AppLocalization` | [AppLocalization ](../platform-sdk/pkg/platformclientmodels/app_localization.go) |
| `AppUpdate` | [AppUpdate ](../platform-sdk/pkg/platformclientmodels/app_update.go) |
| `AppleIAPConfigInfo` | [AppleIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/apple_iap_config_info.go) |
| `AppleIAPConfigRequest` | [AppleIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/apple_iap_config_request.go) |
| `AppleIAPReceipt` | [AppleIAPReceipt ](../platform-sdk/pkg/platformclientmodels/apple_iap_receipt.go) |
| `BasicCategoryInfo` | [BasicCategoryInfo ](../platform-sdk/pkg/platformclientmodels/basic_category_info.go) |
| `BasicItem` | [BasicItem ](../platform-sdk/pkg/platformclientmodels/basic_item.go) |
| `BillingAccount` | [BillingAccount ](../platform-sdk/pkg/platformclientmodels/billing_account.go) |
| `BillingHistoryInfo` | [BillingHistoryInfo ](../platform-sdk/pkg/platformclientmodels/billing_history_info.go) |
| `BillingHistoryPagingSlicedResult` | [BillingHistoryPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/billing_history_paging_sliced_result.go) |
| `BulkOperationResult` | [BulkOperationResult ](../platform-sdk/pkg/platformclientmodels/bulk_operation_result.go) |
| `BundledItemInfo` | [BundledItemInfo ](../platform-sdk/pkg/platformclientmodels/bundled_item_info.go) |
| `CampaignCreate` | [CampaignCreate ](../platform-sdk/pkg/platformclientmodels/campaign_create.go) |
| `CampaignDynamicInfo` | [CampaignDynamicInfo ](../platform-sdk/pkg/platformclientmodels/campaign_dynamic_info.go) |
| `CampaignInfo` | [CampaignInfo ](../platform-sdk/pkg/platformclientmodels/campaign_info.go) |
| `CampaignPagingSlicedResult` | [CampaignPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/campaign_paging_sliced_result.go) |
| `CampaignUpdate` | [CampaignUpdate ](../platform-sdk/pkg/platformclientmodels/campaign_update.go) |
| `CancelRequest` | [CancelRequest ](../platform-sdk/pkg/platformclientmodels/cancel_request.go) |
| `CategoryCreate` | [CategoryCreate ](../platform-sdk/pkg/platformclientmodels/category_create.go) |
| `CategoryInfo` | [CategoryInfo ](../platform-sdk/pkg/platformclientmodels/category_info.go) |
| `CategoryUpdate` | [CategoryUpdate ](../platform-sdk/pkg/platformclientmodels/category_update.go) |
| `CheckoutConfig` | [CheckoutConfig ](../platform-sdk/pkg/platformclientmodels/checkout_config.go) |
| `CodeCreate` | [CodeCreate ](../platform-sdk/pkg/platformclientmodels/code_create.go) |
| `CodeCreateResult` | [CodeCreateResult ](../platform-sdk/pkg/platformclientmodels/code_create_result.go) |
| `CodeInfo` | [CodeInfo ](../platform-sdk/pkg/platformclientmodels/code_info.go) |
| `CodeInfoPagingSlicedResult` | [CodeInfoPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/code_info_paging_sliced_result.go) |
| `ConditionMatchResult` | [ConditionMatchResult ](../platform-sdk/pkg/platformclientmodels/condition_match_result.go) |
| `CreditRequest` | [CreditRequest ](../platform-sdk/pkg/platformclientmodels/credit_request.go) |
| `CreditSummary` | [CreditSummary ](../platform-sdk/pkg/platformclientmodels/credit_summary.go) |
| `CurrencyConfig` | [CurrencyConfig ](../platform-sdk/pkg/platformclientmodels/currency_config.go) |
| `CurrencyCreate` | [CurrencyCreate ](../platform-sdk/pkg/platformclientmodels/currency_create.go) |
| `CurrencyInfo` | [CurrencyInfo ](../platform-sdk/pkg/platformclientmodels/currency_info.go) |
| `CurrencySummary` | [CurrencySummary ](../platform-sdk/pkg/platformclientmodels/currency_summary.go) |
| `CurrencyUpdate` | [CurrencyUpdate ](../platform-sdk/pkg/platformclientmodels/currency_update.go) |
| `Customization` | [Customization ](../platform-sdk/pkg/platformclientmodels/customization.go) |
| `DLCItem` | [DLCItem ](../platform-sdk/pkg/platformclientmodels/dlc_item.go) |
| `DLCItemConfigInfo` | [DLCItemConfigInfo ](../platform-sdk/pkg/platformclientmodels/dlc_item_config_info.go) |
| `DLCItemConfigUpdate` | [DLCItemConfigUpdate ](../platform-sdk/pkg/platformclientmodels/dlc_item_config_update.go) |
| `DebitRequest` | [DebitRequest ](../platform-sdk/pkg/platformclientmodels/debit_request.go) |
| `EntitlementDecrement` | [EntitlementDecrement ](../platform-sdk/pkg/platformclientmodels/entitlement_decrement.go) |
| `EntitlementGrant` | [EntitlementGrant ](../platform-sdk/pkg/platformclientmodels/entitlement_grant.go) |
| `EntitlementHistoryInfo` | [EntitlementHistoryInfo ](../platform-sdk/pkg/platformclientmodels/entitlement_history_info.go) |
| `EntitlementInfo` | [EntitlementInfo ](../platform-sdk/pkg/platformclientmodels/entitlement_info.go) |
| `EntitlementPagingSlicedResult` | [EntitlementPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/entitlement_paging_sliced_result.go) |
| `EntitlementSummary` | [EntitlementSummary ](../platform-sdk/pkg/platformclientmodels/entitlement_summary.go) |
| `EntitlementUpdate` | [EntitlementUpdate ](../platform-sdk/pkg/platformclientmodels/entitlement_update.go) |
| `EpicGamesIAPConfigInfo` | [EpicGamesIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/epic_games_iap_config_info.go) |
| `EpicGamesIAPConfigRequest` | [EpicGamesIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/epic_games_iap_config_request.go) |
| `EpicGamesReconcileRequest` | [EpicGamesReconcileRequest ](../platform-sdk/pkg/platformclientmodels/epic_games_reconcile_request.go) |
| `EpicGamesReconcileResult` | [EpicGamesReconcileResult ](../platform-sdk/pkg/platformclientmodels/epic_games_reconcile_result.go) |
| `ErrorEntity` | [ErrorEntity ](../platform-sdk/pkg/platformclientmodels/error_entity.go) |
| `EventPayload` | [EventPayload ](../platform-sdk/pkg/platformclientmodels/event_payload.go) |
| `ExternalPaymentOrderCreate` | [ExternalPaymentOrderCreate ](../platform-sdk/pkg/platformclientmodels/external_payment_order_create.go) |
| `FieldValidationError` | [FieldValidationError ](../platform-sdk/pkg/platformclientmodels/field_validation_error.go) |
| `FulfillCodeRequest` | [FulfillCodeRequest ](../platform-sdk/pkg/platformclientmodels/fulfill_code_request.go) |
| `FulfillmentError` | [FulfillmentError ](../platform-sdk/pkg/platformclientmodels/fulfillment_error.go) |
| `FulfillmentHistoryInfo` | [FulfillmentHistoryInfo ](../platform-sdk/pkg/platformclientmodels/fulfillment_history_info.go) |
| `FulfillmentHistoryPagingSlicedResult` | [FulfillmentHistoryPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/fulfillment_history_paging_sliced_result.go) |
| `FulfillmentItem` | [FulfillmentItem ](../platform-sdk/pkg/platformclientmodels/fulfillment_item.go) |
| `FulfillmentRequest` | [FulfillmentRequest ](../platform-sdk/pkg/platformclientmodels/fulfillment_request.go) |
| `FulfillmentResult` | [FulfillmentResult ](../platform-sdk/pkg/platformclientmodels/fulfillment_result.go) |
| `FulfillmentScriptContext` | [FulfillmentScriptContext ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_context.go) |
| `FulfillmentScriptCreate` | [FulfillmentScriptCreate ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_create.go) |
| `FulfillmentScriptEvalTestRequest` | [FulfillmentScriptEvalTestRequest ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_eval_test_request.go) |
| `FulfillmentScriptEvalTestResult` | [FulfillmentScriptEvalTestResult ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_eval_test_result.go) |
| `FulfillmentScriptInfo` | [FulfillmentScriptInfo ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_info.go) |
| `FulfillmentScriptUpdate` | [FulfillmentScriptUpdate ](../platform-sdk/pkg/platformclientmodels/fulfillment_script_update.go) |
| `FullAppInfo` | [FullAppInfo ](../platform-sdk/pkg/platformclientmodels/full_app_info.go) |
| `FullCategoryInfo` | [FullCategoryInfo ](../platform-sdk/pkg/platformclientmodels/full_category_info.go) |
| `FullItemInfo` | [FullItemInfo ](../platform-sdk/pkg/platformclientmodels/full_item_info.go) |
| `FullItemPagingSlicedResult` | [FullItemPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/full_item_paging_sliced_result.go) |
| `GoogleIAPConfigInfo` | [GoogleIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/google_iap_config_info.go) |
| `GoogleIAPConfigRequest` | [GoogleIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/google_iap_config_request.go) |
| `GoogleIAPReceipt` | [GoogleIAPReceipt ](../platform-sdk/pkg/platformclientmodels/google_iap_receipt.go) |
| `GoogleReceiptResolveResult` | [GoogleReceiptResolveResult ](../platform-sdk/pkg/platformclientmodels/google_receipt_resolve_result.go) |
| `GrantSubscriptionDaysRequest` | [GrantSubscriptionDaysRequest ](../platform-sdk/pkg/platformclientmodels/grant_subscription_days_request.go) |
| `HierarchicalCategoryInfo` | [HierarchicalCategoryInfo ](../platform-sdk/pkg/platformclientmodels/hierarchical_category_info.go) |
| `IAPItemConfigInfo` | [IAPItemConfigInfo ](../platform-sdk/pkg/platformclientmodels/iap_item_config_info.go) |
| `IAPItemConfigUpdate` | [IAPItemConfigUpdate ](../platform-sdk/pkg/platformclientmodels/iap_item_config_update.go) |
| `IAPItemEntry` | [IAPItemEntry ](../platform-sdk/pkg/platformclientmodels/iap_item_entry.go) |
| `IAPOrderInfo` | [IAPOrderInfo ](../platform-sdk/pkg/platformclientmodels/iap_order_info.go) |
| `IAPOrderPagingSlicedResult` | [IAPOrderPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/iap_order_paging_sliced_result.go) |
| `Image` | [Image ](../platform-sdk/pkg/platformclientmodels/image.go) |
| `InGameItemSync` | [InGameItemSync ](../platform-sdk/pkg/platformclientmodels/in_game_item_sync.go) |
| `ItemAcquireRequest` | [ItemAcquireRequest ](../platform-sdk/pkg/platformclientmodels/item_acquire_request.go) |
| `ItemAcquireResult` | [ItemAcquireResult ](../platform-sdk/pkg/platformclientmodels/item_acquire_result.go) |
| `ItemCreate` | [ItemCreate ](../platform-sdk/pkg/platformclientmodels/item_create.go) |
| `ItemDynamicDataInfo` | [ItemDynamicDataInfo ](../platform-sdk/pkg/platformclientmodels/item_dynamic_data_info.go) |
| `ItemId` | [ItemId ](../platform-sdk/pkg/platformclientmodels/item_id.go) |
| `ItemInfo` | [ItemInfo ](../platform-sdk/pkg/platformclientmodels/item_info.go) |
| `ItemPagingSlicedResult` | [ItemPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/item_paging_sliced_result.go) |
| `ItemReturnRequest` | [ItemReturnRequest ](../platform-sdk/pkg/platformclientmodels/item_return_request.go) |
| `ItemSnapshot` | [ItemSnapshot ](../platform-sdk/pkg/platformclientmodels/item_snapshot.go) |
| `ItemUpdate` | [ItemUpdate ](../platform-sdk/pkg/platformclientmodels/item_update.go) |
| `KeyGroupCreate` | [KeyGroupCreate ](../platform-sdk/pkg/platformclientmodels/key_group_create.go) |
| `KeyGroupDynamicInfo` | [KeyGroupDynamicInfo ](../platform-sdk/pkg/platformclientmodels/key_group_dynamic_info.go) |
| `KeyGroupInfo` | [KeyGroupInfo ](../platform-sdk/pkg/platformclientmodels/key_group_info.go) |
| `KeyGroupPagingSlicedResult` | [KeyGroupPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/key_group_paging_sliced_result.go) |
| `KeyGroupUpdate` | [KeyGroupUpdate ](../platform-sdk/pkg/platformclientmodels/key_group_update.go) |
| `KeyInfo` | [KeyInfo ](../platform-sdk/pkg/platformclientmodels/key_info.go) |
| `KeyPagingSliceResult` | [KeyPagingSliceResult ](../platform-sdk/pkg/platformclientmodels/key_paging_slice_result.go) |
| `Localization` | [Localization ](../platform-sdk/pkg/platformclientmodels/localization.go) |
| `MockIAPReceipt` | [MockIAPReceipt ](../platform-sdk/pkg/platformclientmodels/mock_iap_receipt.go) |
| `NotificationProcessResult` | [NotificationProcessResult ](../platform-sdk/pkg/platformclientmodels/notification_process_result.go) |
| `Order` | [Order ](../platform-sdk/pkg/platformclientmodels/order.go) |
| `OrderCreate` | [OrderCreate ](../platform-sdk/pkg/platformclientmodels/order_create.go) |
| `OrderGrantInfo` | [OrderGrantInfo ](../platform-sdk/pkg/platformclientmodels/order_grant_info.go) |
| `OrderHistoryInfo` | [OrderHistoryInfo ](../platform-sdk/pkg/platformclientmodels/order_history_info.go) |
| `OrderInfo` | [OrderInfo ](../platform-sdk/pkg/platformclientmodels/order_info.go) |
| `OrderPagingResult` | [OrderPagingResult ](../platform-sdk/pkg/platformclientmodels/order_paging_result.go) |
| `OrderPagingSlicedResult` | [OrderPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/order_paging_sliced_result.go) |
| `OrderRefundCreate` | [OrderRefundCreate ](../platform-sdk/pkg/platformclientmodels/order_refund_create.go) |
| `OrderStatistics` | [OrderStatistics ](../platform-sdk/pkg/platformclientmodels/order_statistics.go) |
| `OrderSummary` | [OrderSummary ](../platform-sdk/pkg/platformclientmodels/order_summary.go) |
| `OrderSyncResult` | [OrderSyncResult ](../platform-sdk/pkg/platformclientmodels/order_sync_result.go) |
| `OrderUpdate` | [OrderUpdate ](../platform-sdk/pkg/platformclientmodels/order_update.go) |
| `Ownership` | [Ownership ](../platform-sdk/pkg/platformclientmodels/ownership.go) |
| `OwnershipToken` | [OwnershipToken ](../platform-sdk/pkg/platformclientmodels/ownership_token.go) |
| `Paging` | [Paging ](../platform-sdk/pkg/platformclientmodels/paging.go) |
| `PayPalConfig` | [PayPalConfig ](../platform-sdk/pkg/platformclientmodels/pay_pal_config.go) |
| `PaymentAccount` | [PaymentAccount ](../platform-sdk/pkg/platformclientmodels/payment_account.go) |
| `PaymentCallbackConfigInfo` | [PaymentCallbackConfigInfo ](../platform-sdk/pkg/platformclientmodels/payment_callback_config_info.go) |
| `PaymentCallbackConfigUpdate` | [PaymentCallbackConfigUpdate ](../platform-sdk/pkg/platformclientmodels/payment_callback_config_update.go) |
| `PaymentMerchantConfigInfo` | [PaymentMerchantConfigInfo ](../platform-sdk/pkg/platformclientmodels/payment_merchant_config_info.go) |
| `PaymentMethod` | [PaymentMethod ](../platform-sdk/pkg/platformclientmodels/payment_method.go) |
| `PaymentNotificationInfo` | [PaymentNotificationInfo ](../platform-sdk/pkg/platformclientmodels/payment_notification_info.go) |
| `PaymentNotificationPagingSlicedResult` | [PaymentNotificationPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/payment_notification_paging_sliced_result.go) |
| `PaymentOrder` | [PaymentOrder ](../platform-sdk/pkg/platformclientmodels/payment_order.go) |
| `PaymentOrderChargeRequest` | [PaymentOrderChargeRequest ](../platform-sdk/pkg/platformclientmodels/payment_order_charge_request.go) |
| `PaymentOrderChargeStatus` | [PaymentOrderChargeStatus ](../platform-sdk/pkg/platformclientmodels/payment_order_charge_status.go) |
| `PaymentOrderCreate` | [PaymentOrderCreate ](../platform-sdk/pkg/platformclientmodels/payment_order_create.go) |
| `PaymentOrderCreateResult` | [PaymentOrderCreateResult ](../platform-sdk/pkg/platformclientmodels/payment_order_create_result.go) |
| `PaymentOrderDetails` | [PaymentOrderDetails ](../platform-sdk/pkg/platformclientmodels/payment_order_details.go) |
| `PaymentOrderInfo` | [PaymentOrderInfo ](../platform-sdk/pkg/platformclientmodels/payment_order_info.go) |
| `PaymentOrderNotifySimulation` | [PaymentOrderNotifySimulation ](../platform-sdk/pkg/platformclientmodels/payment_order_notify_simulation.go) |
| `PaymentOrderPagingSlicedResult` | [PaymentOrderPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/payment_order_paging_sliced_result.go) |
| `PaymentOrderPaidResult` | [PaymentOrderPaidResult ](../platform-sdk/pkg/platformclientmodels/payment_order_paid_result.go) |
| `PaymentOrderRefund` | [PaymentOrderRefund ](../platform-sdk/pkg/platformclientmodels/payment_order_refund.go) |
| `PaymentOrderRefundResult` | [PaymentOrderRefundResult ](../platform-sdk/pkg/platformclientmodels/payment_order_refund_result.go) |
| `PaymentOrderSyncResult` | [PaymentOrderSyncResult ](../platform-sdk/pkg/platformclientmodels/payment_order_sync_result.go) |
| `PaymentProcessResult` | [PaymentProcessResult ](../platform-sdk/pkg/platformclientmodels/payment_process_result.go) |
| `PaymentProviderConfigEdit` | [PaymentProviderConfigEdit ](../platform-sdk/pkg/platformclientmodels/payment_provider_config_edit.go) |
| `PaymentProviderConfigInfo` | [PaymentProviderConfigInfo ](../platform-sdk/pkg/platformclientmodels/payment_provider_config_info.go) |
| `PaymentProviderConfigPagingSlicedResult` | [PaymentProviderConfigPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/payment_provider_config_paging_sliced_result.go) |
| `PaymentRequest` | [PaymentRequest ](../platform-sdk/pkg/platformclientmodels/payment_request.go) |
| `PaymentTaxConfigEdit` | [PaymentTaxConfigEdit ](../platform-sdk/pkg/platformclientmodels/payment_tax_config_edit.go) |
| `PaymentTaxConfigInfo` | [PaymentTaxConfigInfo ](../platform-sdk/pkg/platformclientmodels/payment_tax_config_info.go) |
| `PaymentToken` | [PaymentToken ](../platform-sdk/pkg/platformclientmodels/payment_token.go) |
| `PaymentUrl` | [PaymentUrl ](../platform-sdk/pkg/platformclientmodels/payment_url.go) |
| `PaymentUrlCreate` | [PaymentUrlCreate ](../platform-sdk/pkg/platformclientmodels/payment_url_create.go) |
| `PlatformDLCConfigInfo` | [PlatformDLCConfigInfo ](../platform-sdk/pkg/platformclientmodels/platform_dlc_config_info.go) |
| `PlatformDLCConfigUpdate` | [PlatformDLCConfigUpdate ](../platform-sdk/pkg/platformclientmodels/platform_dlc_config_update.go) |
| `PlatformDlcEntry` | [PlatformDlcEntry ](../platform-sdk/pkg/platformclientmodels/platform_dlc_entry.go) |
| `PlatformReward` | [PlatformReward ](../platform-sdk/pkg/platformclientmodels/platform_reward.go) |
| `PlatformRewardCurrency` | [PlatformRewardCurrency ](../platform-sdk/pkg/platformclientmodels/platform_reward_currency.go) |
| `PlatformRewardItem` | [PlatformRewardItem ](../platform-sdk/pkg/platformclientmodels/platform_reward_item.go) |
| `PlatformSubscribeRequest` | [PlatformSubscribeRequest ](../platform-sdk/pkg/platformclientmodels/platform_subscribe_request.go) |
| `PlayStationDLCSyncRequest` | [PlayStationDLCSyncRequest ](../platform-sdk/pkg/platformclientmodels/play_station_dlc_sync_request.go) |
| `PlayStationIAPConfigInfo` | [PlayStationIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/play_station_iap_config_info.go) |
| `PlayStationReconcileRequest` | [PlayStationReconcileRequest ](../platform-sdk/pkg/platformclientmodels/play_station_reconcile_request.go) |
| `PlayStationReconcileResult` | [PlayStationReconcileResult ](../platform-sdk/pkg/platformclientmodels/play_station_reconcile_result.go) |
| `PlaystationIAPConfigRequest` | [PlaystationIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/playstation_iap_config_request.go) |
| `PopulatedItemInfo` | [PopulatedItemInfo ](../platform-sdk/pkg/platformclientmodels/populated_item_info.go) |
| `PurchasedItemCount` | [PurchasedItemCount ](../platform-sdk/pkg/platformclientmodels/purchased_item_count.go) |
| `Recurring` | [Recurring ](../platform-sdk/pkg/platformclientmodels/recurring.go) |
| `RecurringChargeResult` | [RecurringChargeResult ](../platform-sdk/pkg/platformclientmodels/recurring_charge_result.go) |
| `RedeemHistoryInfo` | [RedeemHistoryInfo ](../platform-sdk/pkg/platformclientmodels/redeem_history_info.go) |
| `RedeemHistoryPagingSlicedResult` | [RedeemHistoryPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/redeem_history_paging_sliced_result.go) |
| `RedeemRequest` | [RedeemRequest ](../platform-sdk/pkg/platformclientmodels/redeem_request.go) |
| `RedeemResult` | [RedeemResult ](../platform-sdk/pkg/platformclientmodels/redeem_result.go) |
| `RedeemableItem` | [RedeemableItem ](../platform-sdk/pkg/platformclientmodels/redeemable_item.go) |
| `RegionDataItem` | [RegionDataItem ](../platform-sdk/pkg/platformclientmodels/region_data_item.go) |
| `Requirement` | [Requirement ](../platform-sdk/pkg/platformclientmodels/requirement.go) |
| `RewardCondition` | [RewardCondition ](../platform-sdk/pkg/platformclientmodels/reward_condition.go) |
| `RewardCreate` | [RewardCreate ](../platform-sdk/pkg/platformclientmodels/reward_create.go) |
| `RewardInfo` | [RewardInfo ](../platform-sdk/pkg/platformclientmodels/reward_info.go) |
| `RewardItem` | [RewardItem ](../platform-sdk/pkg/platformclientmodels/reward_item.go) |
| `RewardPagingSlicedResult` | [RewardPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/reward_paging_sliced_result.go) |
| `RewardUpdate` | [RewardUpdate ](../platform-sdk/pkg/platformclientmodels/reward_update.go) |
| `RewardsRequest` | [RewardsRequest ](../platform-sdk/pkg/platformclientmodels/rewards_request.go) |
| `Slide` | [Slide ](../platform-sdk/pkg/platformclientmodels/slide.go) |
| `StackableEntitlementInfo` | [StackableEntitlementInfo ](../platform-sdk/pkg/platformclientmodels/stackable_entitlement_info.go) |
| `StadiaIAPConfigInfo` | [StadiaIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/stadia_iap_config_info.go) |
| `StadiaSyncRequest` | [StadiaSyncRequest ](../platform-sdk/pkg/platformclientmodels/stadia_sync_request.go) |
| `SteamDLCSyncRequest` | [SteamDLCSyncRequest ](../platform-sdk/pkg/platformclientmodels/steam_dlc_sync_request.go) |
| `SteamIAPConfig` | [SteamIAPConfig ](../platform-sdk/pkg/platformclientmodels/steam_iap_config.go) |
| `SteamIAPConfigInfo` | [SteamIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/steam_iap_config_info.go) |
| `SteamIAPConfigRequest` | [SteamIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/steam_iap_config_request.go) |
| `SteamSyncRequest` | [SteamSyncRequest ](../platform-sdk/pkg/platformclientmodels/steam_sync_request.go) |
| `StoreBackupInfo` | [StoreBackupInfo ](../platform-sdk/pkg/platformclientmodels/store_backup_info.go) |
| `StoreCreate` | [StoreCreate ](../platform-sdk/pkg/platformclientmodels/store_create.go) |
| `StoreInfo` | [StoreInfo ](../platform-sdk/pkg/platformclientmodels/store_info.go) |
| `StoreUpdate` | [StoreUpdate ](../platform-sdk/pkg/platformclientmodels/store_update.go) |
| `StripeConfig` | [StripeConfig ](../platform-sdk/pkg/platformclientmodels/stripe_config.go) |
| `Subscribable` | [Subscribable ](../platform-sdk/pkg/platformclientmodels/subscribable.go) |
| `SubscribeRequest` | [SubscribeRequest ](../platform-sdk/pkg/platformclientmodels/subscribe_request.go) |
| `SubscriptionActivityInfo` | [SubscriptionActivityInfo ](../platform-sdk/pkg/platformclientmodels/subscription_activity_info.go) |
| `SubscriptionActivityPagingSlicedResult` | [SubscriptionActivityPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/subscription_activity_paging_sliced_result.go) |
| `SubscriptionInfo` | [SubscriptionInfo ](../platform-sdk/pkg/platformclientmodels/subscription_info.go) |
| `SubscriptionPagingSlicedResult` | [SubscriptionPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/subscription_paging_sliced_result.go) |
| `SubscriptionSummary` | [SubscriptionSummary ](../platform-sdk/pkg/platformclientmodels/subscription_summary.go) |
| `TaxResult` | [TaxResult ](../platform-sdk/pkg/platformclientmodels/tax_result.go) |
| `TestResult` | [TestResult ](../platform-sdk/pkg/platformclientmodels/test_result.go) |
| `TicketAcquireRequest` | [TicketAcquireRequest ](../platform-sdk/pkg/platformclientmodels/ticket_acquire_request.go) |
| `TicketAcquireResult` | [TicketAcquireResult ](../platform-sdk/pkg/platformclientmodels/ticket_acquire_result.go) |
| `TicketBoothID` | [TicketBoothID ](../platform-sdk/pkg/platformclientmodels/ticket_booth_id.go) |
| `TicketDynamicInfo` | [TicketDynamicInfo ](../platform-sdk/pkg/platformclientmodels/ticket_dynamic_info.go) |
| `TicketSaleDecrementRequest` | [TicketSaleDecrementRequest ](../platform-sdk/pkg/platformclientmodels/ticket_sale_decrement_request.go) |
| `TicketSaleIncrementRequest` | [TicketSaleIncrementRequest ](../platform-sdk/pkg/platformclientmodels/ticket_sale_increment_request.go) |
| `TicketSaleIncrementResult` | [TicketSaleIncrementResult ](../platform-sdk/pkg/platformclientmodels/ticket_sale_increment_result.go) |
| `TimedOwnership` | [TimedOwnership ](../platform-sdk/pkg/platformclientmodels/timed_ownership.go) |
| `TradeNotification` | [TradeNotification ](../platform-sdk/pkg/platformclientmodels/trade_notification.go) |
| `Transaction` | [Transaction ](../platform-sdk/pkg/platformclientmodels/transaction.go) |
| `TwitchIAPConfigInfo` | [TwitchIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/twitch_iap_config_info.go) |
| `TwitchIAPConfigRequest` | [TwitchIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/twitch_iap_config_request.go) |
| `TwitchSyncRequest` | [TwitchSyncRequest ](../platform-sdk/pkg/platformclientmodels/twitch_sync_request.go) |
| `ValidationErrorEntity` | [ValidationErrorEntity ](../platform-sdk/pkg/platformclientmodels/validation_error_entity.go) |
| `WalletInfo` | [WalletInfo ](../platform-sdk/pkg/platformclientmodels/wallet_info.go) |
| `WalletPagingSlicedResult` | [WalletPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/wallet_paging_sliced_result.go) |
| `WalletTransactionInfo` | [WalletTransactionInfo ](../platform-sdk/pkg/platformclientmodels/wallet_transaction_info.go) |
| `WalletTransactionPagingSlicedResult` | [WalletTransactionPagingSlicedResult ](../platform-sdk/pkg/platformclientmodels/wallet_transaction_paging_sliced_result.go) |
| `WxPayConfigInfo` | [WxPayConfigInfo ](../platform-sdk/pkg/platformclientmodels/wx_pay_config_info.go) |
| `WxPayConfigRequest` | [WxPayConfigRequest ](../platform-sdk/pkg/platformclientmodels/wx_pay_config_request.go) |
| `XblDLCSyncRequest` | [XblDLCSyncRequest ](../platform-sdk/pkg/platformclientmodels/xbl_dlc_sync_request.go) |
| `XblIAPConfigInfo` | [XblIAPConfigInfo ](../platform-sdk/pkg/platformclientmodels/xbl_iap_config_info.go) |
| `XblIAPConfigRequest` | [XblIAPConfigRequest ](../platform-sdk/pkg/platformclientmodels/xbl_iap_config_request.go) |
| `XblReconcileRequest` | [XblReconcileRequest ](../platform-sdk/pkg/platformclientmodels/xbl_reconcile_request.go) |
| `XblReconcileResult` | [XblReconcileResult ](../platform-sdk/pkg/platformclientmodels/xbl_reconcile_result.go) |
| `XsollaConfig` | [XsollaConfig ](../platform-sdk/pkg/platformclientmodels/xsolla_config.go) |
| `XsollaPaywallConfig` | [XsollaPaywallConfig ](../platform-sdk/pkg/platformclientmodels/xsolla_paywall_config.go) |
| `XsollaPaywallConfigRequest` | [XsollaPaywallConfigRequest ](../platform-sdk/pkg/platformclientmodels/xsolla_paywall_config_request.go) |