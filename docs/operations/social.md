# Social Service Index

&nbsp;

## Operations

### SlotConfig Wrapper:  [SlotConfig](../../services-api/pkg/service/social/slotConfig.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/admin/namespaces/{namespace}/config` | GET | GetNamespaceSlotConfigShort | [GetNamespaceSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [GetNamespaceSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [GetNamespaceSlotConfigShort](../../samples/cli/cmd/social/slotConfig/getNamespaceSlotConfig.go) |
| `/social/admin/namespaces/{namespace}/config` | PUT | UpdateNamespaceSlotConfigShort | [UpdateNamespaceSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [UpdateNamespaceSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [UpdateNamespaceSlotConfigShort](../../samples/cli/cmd/social/slotConfig/updateNamespaceSlotConfig.go) |
| `/social/admin/namespaces/{namespace}/config` | DELETE | DeleteNamespaceSlotConfigShort | [DeleteNamespaceSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [DeleteNamespaceSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [DeleteNamespaceSlotConfigShort](../../samples/cli/cmd/social/slotConfig/deleteNamespaceSlotConfig.go) |
| `/social/admin/namespaces/{namespace}/users/{userId}/config` | GET | GetUserSlotConfigShort | [GetUserSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [GetUserSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [GetUserSlotConfigShort](../../samples/cli/cmd/social/slotConfig/getUserSlotConfig.go) |
| `/social/admin/namespaces/{namespace}/users/{userId}/config` | PUT | UpdateUserSlotConfigShort | [UpdateUserSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [UpdateUserSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [UpdateUserSlotConfigShort](../../samples/cli/cmd/social/slotConfig/updateUserSlotConfig.go) |
| `/social/admin/namespaces/{namespace}/users/{userId}/config` | DELETE | DeleteUserSlotConfigShort | [DeleteUserSlotConfigShort](../../social-sdk/pkg/socialclient/slot_config/slot_config_client.go) | [DeleteUserSlotConfigShort](../../services-api/pkg/service/social/slotConfig.go) | [DeleteUserSlotConfigShort](../../samples/cli/cmd/social/slotConfig/deleteUserSlotConfig.go) |

### GameProfile Wrapper:  [GameProfile](../../services-api/pkg/service/social/gameProfile.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/admin/namespaces/{namespace}/users/{userId}/profiles` | GET | GetUserProfilesShort | [GetUserProfilesShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [GetUserProfilesShort](../../services-api/pkg/service/social/gameProfile.go) | [GetUserProfilesShort](../../samples/cli/cmd/social/gameProfile/getUserProfiles.go) |
| `/social/admin/namespaces/{namespace}/users/{userId}/profiles/{profileId}` | GET | GetProfileShort | [GetProfileShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [GetProfileShort](../../services-api/pkg/service/social/gameProfile.go) | [GetProfileShort](../../samples/cli/cmd/social/gameProfile/getProfile.go) |
| `/social/public/namespaces/{namespace}/profiles` | GET | PublicGetUserGameProfilesShort | [PublicGetUserGameProfilesShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicGetUserGameProfilesShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicGetUserGameProfilesShort](../../samples/cli/cmd/social/gameProfile/publicGetUserGameProfiles.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles` | GET | PublicGetUserProfilesShort | [PublicGetUserProfilesShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicGetUserProfilesShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicGetUserProfilesShort](../../samples/cli/cmd/social/gameProfile/publicGetUserProfiles.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles` | POST | PublicCreateProfileShort | [PublicCreateProfileShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicCreateProfileShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicCreateProfileShort](../../samples/cli/cmd/social/gameProfile/publicCreateProfile.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}` | GET | PublicGetProfileShort | [PublicGetProfileShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicGetProfileShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicGetProfileShort](../../samples/cli/cmd/social/gameProfile/publicGetProfile.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}` | PUT | PublicUpdateProfileShort | [PublicUpdateProfileShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicUpdateProfileShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicUpdateProfileShort](../../samples/cli/cmd/social/gameProfile/publicUpdateProfile.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}` | DELETE | PublicDeleteProfileShort | [PublicDeleteProfileShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicDeleteProfileShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicDeleteProfileShort](../../samples/cli/cmd/social/gameProfile/publicDeleteProfile.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}/attributes/{attributeName}` | GET | PublicGetProfileAttributeShort | [PublicGetProfileAttributeShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicGetProfileAttributeShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicGetProfileAttributeShort](../../samples/cli/cmd/social/gameProfile/publicGetProfileAttribute.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/profiles/{profileId}/attributes/{attributeName}` | PUT | PublicUpdateAttributeShort | [PublicUpdateAttributeShort](../../social-sdk/pkg/socialclient/game_profile/game_profile_client.go) | [PublicUpdateAttributeShort](../../services-api/pkg/service/social/gameProfile.go) | [PublicUpdateAttributeShort](../../samples/cli/cmd/social/gameProfile/publicUpdateAttribute.go) |

### Slot Wrapper:  [Slot](../../services-api/pkg/service/social/slot.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/admin/namespaces/{namespace}/users/{userId}/slots` | GET | GetUserNamespaceSlotsShort | [GetUserNamespaceSlotsShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [GetUserNamespaceSlotsShort](../../services-api/pkg/service/social/slot.go) | [GetUserNamespaceSlotsShort](../../samples/cli/cmd/social/slot/getUserNamespaceSlots.go) |
| `/social/admin/namespaces/{namespace}/users/{userId}/slots/{slotId}` | GET | GetSlotDataShort | [GetSlotDataShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [GetSlotDataShort](../../services-api/pkg/service/social/slot.go) | [GetSlotDataShort](../../samples/cli/cmd/social/slot/getSlotData.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots` | GET | PublicGetUserNamespaceSlotsShort | [PublicGetUserNamespaceSlotsShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicGetUserNamespaceSlotsShort](../../services-api/pkg/service/social/slot.go) | [PublicGetUserNamespaceSlotsShort](../../samples/cli/cmd/social/slot/publicGetUserNamespaceSlots.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots` | POST | PublicCreateUserNamespaceSlotShort | [PublicCreateUserNamespaceSlotShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicCreateUserNamespaceSlotShort](../../services-api/pkg/service/social/slot.go) | [PublicCreateUserNamespaceSlotShort](../../samples/cli/cmd/social/slot/publicCreateUserNamespaceSlot.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots/{slotId}` | GET | PublicGetSlotDataShort | [PublicGetSlotDataShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicGetSlotDataShort](../../services-api/pkg/service/social/slot.go) | [PublicGetSlotDataShort](../../samples/cli/cmd/social/slot/publicGetSlotData.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots/{slotId}` | PUT | PublicUpdateUserNamespaceSlotShort | [PublicUpdateUserNamespaceSlotShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicUpdateUserNamespaceSlotShort](../../services-api/pkg/service/social/slot.go) | [PublicUpdateUserNamespaceSlotShort](../../samples/cli/cmd/social/slot/publicUpdateUserNamespaceSlot.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots/{slotId}` | DELETE | PublicDeleteUserNamespaceSlotShort | [PublicDeleteUserNamespaceSlotShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicDeleteUserNamespaceSlotShort](../../services-api/pkg/service/social/slot.go) | [PublicDeleteUserNamespaceSlotShort](../../samples/cli/cmd/social/slot/publicDeleteUserNamespaceSlot.go) |
| `/social/public/namespaces/{namespace}/users/{userId}/slots/{slotId}/metadata` | PUT | PublicUpdateUserNamespaceSlotMetadataShort | [PublicUpdateUserNamespaceSlotMetadataShort](../../social-sdk/pkg/socialclient/slot/slot_client.go) | [PublicUpdateUserNamespaceSlotMetadataShort](../../services-api/pkg/service/social/slot.go) | [PublicUpdateUserNamespaceSlotMetadataShort](../../samples/cli/cmd/social/slot/publicUpdateUserNamespaceSlotMetadata.go) |

### GlobalStatistic Wrapper:  [GlobalStatistic](../../services-api/pkg/service/social/globalStatistic.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/v1/admin/namespaces/{namespace}/globalstatitems` | GET | GetGlobalStatItemsShort | [GetGlobalStatItemsShort](../../social-sdk/pkg/socialclient/global_statistic/global_statistic_client.go) | [GetGlobalStatItemsShort](../../services-api/pkg/service/social/globalStatistic.go) | [GetGlobalStatItemsShort](../../samples/cli/cmd/social/globalStatistic/getGlobalStatItems.go) |

### UserStatistic Wrapper:  [UserStatistic](../../services-api/pkg/service/social/userStatistic.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/v1/admin/namespaces/{namespace}/statitems/bulk` | GET | BulkFetchStatItemsShort | [BulkFetchStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkFetchStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkFetchStatItemsShort](../../samples/cli/cmd/social/userStatistic/bulkFetchStatItems.go) |
| `/social/v1/admin/namespaces/{namespace}/statitems/value/bulk` | PUT | BulkIncUserStatItemShort | [BulkIncUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkIncUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkIncUserStatItemShort](../../samples/cli/cmd/social/userStatistic/bulkIncUserStatItem.go) |
| `/social/v1/admin/namespaces/{namespace}/statitems/value/bulk` | PATCH | BulkIncUserStatItemValueShort | [BulkIncUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkIncUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkIncUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/bulkIncUserStatItemValue.go) |
| `/social/v1/admin/namespaces/{namespace}/statitems/value/bulk/getOrDefault` | GET | BulkFetchOrDefaultStatItemsShort | [BulkFetchOrDefaultStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkFetchOrDefaultStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkFetchOrDefaultStatItemsShort](../../samples/cli/cmd/social/userStatistic/bulkFetchOrDefaultStatItems.go) |
| `/social/v1/admin/namespaces/{namespace}/statitems/value/reset/bulk` | PUT | BulkResetUserStatItemShort | [BulkResetUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkResetUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkResetUserStatItemShort](../../samples/cli/cmd/social/userStatistic/bulkResetUserStatItem.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/statitems` | GET | GetUserStatItemsShort | [GetUserStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [GetUserStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [GetUserStatItemsShort](../../samples/cli/cmd/social/userStatistic/getUserStatItems.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/bulk` | POST | BulkCreateUserStatItemsShort | [BulkCreateUserStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkCreateUserStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkCreateUserStatItemsShort](../../samples/cli/cmd/social/userStatistic/bulkCreateUserStatItems.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PUT | BulkIncUserStatItem1Short | [BulkIncUserStatItem1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkIncUserStatItem1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkIncUserStatItem1Short](../../samples/cli/cmd/social/userStatistic/bulkIncUserStatItem1.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PATCH | BulkIncUserStatItemValue1Short | [BulkIncUserStatItemValue1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkIncUserStatItemValue1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkIncUserStatItemValue1Short](../../samples/cli/cmd/social/userStatistic/bulkIncUserStatItemValue1.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk` | PUT | BulkResetUserStatItem1Short | [BulkResetUserStatItem1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkResetUserStatItem1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkResetUserStatItem1Short](../../samples/cli/cmd/social/userStatistic/bulkResetUserStatItem1.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems` | POST | CreateUserStatItemShort | [CreateUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [CreateUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [CreateUserStatItemShort](../../samples/cli/cmd/social/userStatistic/createUserStatItem.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems` | DELETE | DeleteUserStatItemsShort | [DeleteUserStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [DeleteUserStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [DeleteUserStatItemsShort](../../samples/cli/cmd/social/userStatistic/deleteUserStatItems.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value` | PATCH | IncUserStatItemValueShort | [IncUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [IncUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [IncUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/incUserStatItemValue.go) |
| `/social/v1/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset` | PUT | ResetUserStatItemValueShort | [ResetUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [ResetUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [ResetUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/resetUserStatItemValue.go) |
| `/social/v1/public/namespaces/{namespace}/statitems/bulk` | GET | BulkFetchStatItems1Short | [BulkFetchStatItems1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkFetchStatItems1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkFetchStatItems1Short](../../samples/cli/cmd/social/userStatistic/bulkFetchStatItems1.go) |
| `/social/v1/public/namespaces/{namespace}/statitems/value/bulk` | PUT | PublicBulkIncUserStatItemShort | [PublicBulkIncUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicBulkIncUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicBulkIncUserStatItemShort](../../samples/cli/cmd/social/userStatistic/publicBulkIncUserStatItem.go) |
| `/social/v1/public/namespaces/{namespace}/statitems/value/bulk` | PATCH | PublicBulkIncUserStatItemValueShort | [PublicBulkIncUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicBulkIncUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicBulkIncUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/publicBulkIncUserStatItemValue.go) |
| `/social/v1/public/namespaces/{namespace}/statitems/value/reset/bulk` | PUT | BulkResetUserStatItem2Short | [BulkResetUserStatItem2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkResetUserStatItem2Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkResetUserStatItem2Short](../../samples/cli/cmd/social/userStatistic/bulkResetUserStatItem2.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems` | GET | PublicQueryUserStatItemsShort | [PublicQueryUserStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicQueryUserStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicQueryUserStatItemsShort](../../samples/cli/cmd/social/userStatistic/publicQueryUserStatItems.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems/bulk` | POST | PublicBulkCreateUserStatItemsShort | [PublicBulkCreateUserStatItemsShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicBulkCreateUserStatItemsShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicBulkCreateUserStatItemsShort](../../samples/cli/cmd/social/userStatistic/publicBulkCreateUserStatItems.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | GET | PublicQueryUserStatItems1Short | [PublicQueryUserStatItems1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicQueryUserStatItems1Short](../../services-api/pkg/service/social/userStatistic.go) | [PublicQueryUserStatItems1Short](../../samples/cli/cmd/social/userStatistic/publicQueryUserStatItems1.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PUT | PublicBulkIncUserStatItem1Short | [PublicBulkIncUserStatItem1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicBulkIncUserStatItem1Short](../../services-api/pkg/service/social/userStatistic.go) | [PublicBulkIncUserStatItem1Short](../../samples/cli/cmd/social/userStatistic/publicBulkIncUserStatItem1.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PATCH | BulkIncUserStatItemValue2Short | [BulkIncUserStatItemValue2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkIncUserStatItemValue2Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkIncUserStatItemValue2Short](../../samples/cli/cmd/social/userStatistic/bulkIncUserStatItemValue2.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk` | PUT | BulkResetUserStatItem3Short | [BulkResetUserStatItem3Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkResetUserStatItem3Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkResetUserStatItem3Short](../../samples/cli/cmd/social/userStatistic/bulkResetUserStatItem3.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems` | POST | PublicCreateUserStatItemShort | [PublicCreateUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicCreateUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicCreateUserStatItemShort](../../samples/cli/cmd/social/userStatistic/publicCreateUserStatItem.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems` | DELETE | DeleteUserStatItems1Short | [DeleteUserStatItems1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [DeleteUserStatItems1Short](../../services-api/pkg/service/social/userStatistic.go) | [DeleteUserStatItems1Short](../../samples/cli/cmd/social/userStatistic/deleteUserStatItems1.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value` | PUT | PublicIncUserStatItemShort | [PublicIncUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicIncUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicIncUserStatItemShort](../../samples/cli/cmd/social/userStatistic/publicIncUserStatItem.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value` | PATCH | PublicIncUserStatItemValueShort | [PublicIncUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicIncUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [PublicIncUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/publicIncUserStatItemValue.go) |
| `/social/v1/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value/reset` | PUT | ResetUserStatItemValue1Short | [ResetUserStatItemValue1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [ResetUserStatItemValue1Short](../../services-api/pkg/service/social/userStatistic.go) | [ResetUserStatItemValue1Short](../../samples/cli/cmd/social/userStatistic/resetUserStatItemValue1.go) |
| `/social/v2/admin/namespaces/{namespace}/statitems/value/bulk` | PUT | BulkUpdateUserStatItemV2Short | [BulkUpdateUserStatItemV2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkUpdateUserStatItemV2Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkUpdateUserStatItemV2Short](../../samples/cli/cmd/social/userStatistic/bulkUpdateUserStatItemV2.go) |
| `/social/v2/admin/namespaces/{namespace}/statitems/value/bulk/getOrDefault` | GET | BulkFetchOrDefaultStatItems1Short | [BulkFetchOrDefaultStatItems1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkFetchOrDefaultStatItems1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkFetchOrDefaultStatItems1Short](../../samples/cli/cmd/social/userStatistic/bulkFetchOrDefaultStatItems1.go) |
| `/social/v2/admin/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PUT | BulkUpdateUserStatItemShort | [BulkUpdateUserStatItemShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkUpdateUserStatItemShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkUpdateUserStatItemShort](../../samples/cli/cmd/social/userStatistic/bulkUpdateUserStatItem.go) |
| `/social/v2/admin/namespaces/{namespace}/users/{userId}/statitems/value/reset/bulk` | PUT | BulkResetUserStatItemValuesShort | [BulkResetUserStatItemValuesShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkResetUserStatItemValuesShort](../../services-api/pkg/service/social/userStatistic.go) | [BulkResetUserStatItemValuesShort](../../samples/cli/cmd/social/userStatistic/bulkResetUserStatItemValues.go) |
| `/social/v2/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems` | DELETE | DeleteUserStatItems2Short | [DeleteUserStatItems2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [DeleteUserStatItems2Short](../../services-api/pkg/service/social/userStatistic.go) | [DeleteUserStatItems2Short](../../samples/cli/cmd/social/userStatistic/deleteUserStatItems2.go) |
| `/social/v2/admin/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value` | PUT | UpdateUserStatItemValueShort | [UpdateUserStatItemValueShort](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [UpdateUserStatItemValueShort](../../services-api/pkg/service/social/userStatistic.go) | [UpdateUserStatItemValueShort](../../samples/cli/cmd/social/userStatistic/updateUserStatItemValue.go) |
| `/social/v2/public/namespaces/{namespace}/statitems/value/bulk` | PUT | BulkUpdateUserStatItem1Short | [BulkUpdateUserStatItem1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkUpdateUserStatItem1Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkUpdateUserStatItem1Short](../../samples/cli/cmd/social/userStatistic/bulkUpdateUserStatItem1.go) |
| `/social/v2/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | GET | PublicQueryUserStatItems2Short | [PublicQueryUserStatItems2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [PublicQueryUserStatItems2Short](../../services-api/pkg/service/social/userStatistic.go) | [PublicQueryUserStatItems2Short](../../samples/cli/cmd/social/userStatistic/publicQueryUserStatItems2.go) |
| `/social/v2/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk` | PUT | BulkUpdateUserStatItem2Short | [BulkUpdateUserStatItem2Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [BulkUpdateUserStatItem2Short](../../services-api/pkg/service/social/userStatistic.go) | [BulkUpdateUserStatItem2Short](../../samples/cli/cmd/social/userStatistic/bulkUpdateUserStatItem2.go) |
| `/social/v2/public/namespaces/{namespace}/users/{userId}/stats/{statCode}/statitems/value` | PUT | UpdateUserStatItemValue1Short | [UpdateUserStatItemValue1Short](../../social-sdk/pkg/socialclient/user_statistic/user_statistic_client.go) | [UpdateUserStatItemValue1Short](../../services-api/pkg/service/social/userStatistic.go) | [UpdateUserStatItemValue1Short](../../samples/cli/cmd/social/userStatistic/updateUserStatItemValue1.go) |

### StatConfiguration Wrapper:  [StatConfiguration](../../services-api/pkg/service/social/statConfiguration.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/social/v1/admin/namespaces/{namespace}/stats` | GET | GetStatsShort | [GetStatsShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [GetStatsShort](../../services-api/pkg/service/social/statConfiguration.go) | [GetStatsShort](../../samples/cli/cmd/social/statConfiguration/getStats.go) |
| `/social/v1/admin/namespaces/{namespace}/stats` | POST | CreateStatShort | [CreateStatShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [CreateStatShort](../../services-api/pkg/service/social/statConfiguration.go) | [CreateStatShort](../../samples/cli/cmd/social/statConfiguration/createStat.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/export` | GET | ExportStatsShort | [ExportStatsShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [ExportStatsShort](../../services-api/pkg/service/social/statConfiguration.go) | [ExportStatsShort](../../samples/cli/cmd/social/statConfiguration/exportStats.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/import` | POST | ImportStatsShort | [ImportStatsShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [ImportStatsShort](../../services-api/pkg/service/social/statConfiguration.go) | [ImportStatsShort](../../samples/cli/cmd/social/statConfiguration/importStats.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/search` | GET | QueryStatsShort | [QueryStatsShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [QueryStatsShort](../../services-api/pkg/service/social/statConfiguration.go) | [QueryStatsShort](../../samples/cli/cmd/social/statConfiguration/queryStats.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/{statCode}` | GET | GetStatShort | [GetStatShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [GetStatShort](../../services-api/pkg/service/social/statConfiguration.go) | [GetStatShort](../../samples/cli/cmd/social/statConfiguration/getStat.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/{statCode}` | DELETE | DeleteStatShort | [DeleteStatShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [DeleteStatShort](../../services-api/pkg/service/social/statConfiguration.go) | [DeleteStatShort](../../samples/cli/cmd/social/statConfiguration/deleteStat.go) |
| `/social/v1/admin/namespaces/{namespace}/stats/{statCode}` | PATCH | UpdateStatShort | [UpdateStatShort](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [UpdateStatShort](../../services-api/pkg/service/social/statConfiguration.go) | [UpdateStatShort](../../samples/cli/cmd/social/statConfiguration/updateStat.go) |
| `/social/v1/public/namespaces/{namespace}/stats` | POST | CreateStat1Short | [CreateStat1Short](../../social-sdk/pkg/socialclient/stat_configuration/stat_configuration_client.go) | [CreateStat1Short](../../services-api/pkg/service/social/statConfiguration.go) | [CreateStat1Short](../../samples/cli/cmd/social/statConfiguration/createStat1.go) |


&nbsp;  

## Models

| Model Struct | Class |
|---|---|
| `A DTO object for resetting user stat items` | [ADTOObjectForResettingUserStatItems ](../../social-sdk/pkg/socialclientmodels/a_dto_object_for_resetting_user_stat_items.go) |
| `A DTO object for user stat item value` | [ADTOObjectForUserStatItemValue ](../../social-sdk/pkg/socialclientmodels/a_dto_object_for_user_stat_item_value.go) |
| `Attribute` | [Attribute ](../../social-sdk/pkg/socialclientmodels/attribute.go) |
| `BulkStatItemCreate` | [BulkStatItemCreate ](../../social-sdk/pkg/socialclientmodels/bulk_stat_item_create.go) |
| `BulkStatItemInc` | [BulkStatItemInc ](../../social-sdk/pkg/socialclientmodels/bulk_stat_item_inc.go) |
| `BulkStatItemOperationResult` | [BulkStatItemOperationResult ](../../social-sdk/pkg/socialclientmodels/bulk_stat_item_operation_result.go) |
| `BulkStatItemReset` | [BulkStatItemReset ](../../social-sdk/pkg/socialclientmodels/bulk_stat_item_reset.go) |
| `BulkStatItemUpdate` | [BulkStatItemUpdate ](../../social-sdk/pkg/socialclientmodels/bulk_stat_item_update.go) |
| `BulkUserStatItemInc` | [BulkUserStatItemInc ](../../social-sdk/pkg/socialclientmodels/bulk_user_stat_item_inc.go) |
| `BulkUserStatItemReset` | [BulkUserStatItemReset ](../../social-sdk/pkg/socialclientmodels/bulk_user_stat_item_reset.go) |
| `BulkUserStatItemUpdate` | [BulkUserStatItemUpdate ](../../social-sdk/pkg/socialclientmodels/bulk_user_stat_item_update.go) |
| `ErrorEntity` | [ErrorEntity ](../../social-sdk/pkg/socialclientmodels/error_entity.go) |
| `FieldValidationError` | [FieldValidationError ](../../social-sdk/pkg/socialclientmodels/field_validation_error.go) |
| `GameProfileHeader` | [GameProfileHeader ](../../social-sdk/pkg/socialclientmodels/game_profile_header.go) |
| `GameProfileInfo` | [GameProfileInfo ](../../social-sdk/pkg/socialclientmodels/game_profile_info.go) |
| `GameProfilePublicInfo` | [GameProfilePublicInfo ](../../social-sdk/pkg/socialclientmodels/game_profile_public_info.go) |
| `GameProfileRequest` | [GameProfileRequest ](../../social-sdk/pkg/socialclientmodels/game_profile_request.go) |
| `GlobalStatItemInfo` | [GlobalStatItemInfo ](../../social-sdk/pkg/socialclientmodels/global_stat_item_info.go) |
| `GlobalStatItemPagingSlicedResult` | [GlobalStatItemPagingSlicedResult ](../../social-sdk/pkg/socialclientmodels/global_stat_item_paging_sliced_result.go) |
| `NamespaceSlotConfigInfo` | [NamespaceSlotConfigInfo ](../../social-sdk/pkg/socialclientmodels/namespace_slot_config_info.go) |
| `Paging` | [Paging ](../../social-sdk/pkg/socialclientmodels/paging.go) |
| `SlotConfigUpdate` | [SlotConfigUpdate ](../../social-sdk/pkg/socialclientmodels/slot_config_update.go) |
| `SlotInfo` | [SlotInfo ](../../social-sdk/pkg/socialclientmodels/slot_info.go) |
| `SlotMetadataUpdate` | [SlotMetadataUpdate ](../../social-sdk/pkg/socialclientmodels/slot_metadata_update.go) |
| `StatCreate` | [StatCreate ](../../social-sdk/pkg/socialclientmodels/stat_create.go) |
| `StatImportInfo` | [StatImportInfo ](../../social-sdk/pkg/socialclientmodels/stat_import_info.go) |
| `StatInfo` | [StatInfo ](../../social-sdk/pkg/socialclientmodels/stat_info.go) |
| `StatItemInc` | [StatItemInc ](../../social-sdk/pkg/socialclientmodels/stat_item_inc.go) |
| `StatItemIncResult` | [StatItemIncResult ](../../social-sdk/pkg/socialclientmodels/stat_item_inc_result.go) |
| `StatItemUpdate` | [StatItemUpdate ](../../social-sdk/pkg/socialclientmodels/stat_item_update.go) |
| `StatPagingSlicedResult` | [StatPagingSlicedResult ](../../social-sdk/pkg/socialclientmodels/stat_paging_sliced_result.go) |
| `StatResetInfo` | [StatResetInfo ](../../social-sdk/pkg/socialclientmodels/stat_reset_info.go) |
| `StatUpdate` | [StatUpdate ](../../social-sdk/pkg/socialclientmodels/stat_update.go) |
| `UserGameProfiles` | [UserGameProfiles ](../../social-sdk/pkg/socialclientmodels/user_game_profiles.go) |
| `UserSlotConfigInfo` | [UserSlotConfigInfo ](../../social-sdk/pkg/socialclientmodels/user_slot_config_info.go) |
| `UserStatItemInfo` | [UserStatItemInfo ](../../social-sdk/pkg/socialclientmodels/user_stat_item_info.go) |
| `UserStatItemPagingSlicedResult` | [UserStatItemPagingSlicedResult ](../../social-sdk/pkg/socialclientmodels/user_stat_item_paging_sliced_result.go) |
| `ValidationErrorEntity` | [ValidationErrorEntity ](../../social-sdk/pkg/socialclientmodels/validation_error_entity.go) |