# Sessionhistory Service Index

&nbsp;

## Operations

### Operations Wrapper:  [Operations](../../services-api/pkg/service/sessionhistory/operations.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/healthz` | GET | GetHealthcheckInfoShort | [GetHealthcheckInfoShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/operations/operations_client.go) | [GetHealthcheckInfoShort](../../services-api/pkg/service/sessionhistory/operations.go) | [GetHealthcheckInfoShort](../../samples/cli/cmd/sessionhistory/operations/getHealthcheckInfo.go) |
| `/sessionhistory/healthz` | GET | GetHealthcheckInfoV1Short | [GetHealthcheckInfoV1Short](../../sessionhistory-sdk/pkg/sessionhistoryclient/operations/operations_client.go) | [GetHealthcheckInfoV1Short](../../services-api/pkg/service/sessionhistory/operations.go) | [GetHealthcheckInfoV1Short](../../samples/cli/cmd/sessionhistory/operations/getHealthcheckInfoV1.go) |

### Config Wrapper:  [Config](../../services-api/pkg/service/sessionhistory/config.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/sessionhistory/v1/admin/config/log` | GET | AdminGetLogConfigShort | [AdminGetLogConfigShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/config/config_client.go) | [AdminGetLogConfigShort](../../services-api/pkg/service/sessionhistory/config.go) | [AdminGetLogConfigShort](../../samples/cli/cmd/sessionhistory/config/adminGetLogConfig.go) |
| `/sessionhistory/v1/admin/config/log` | PATCH | AdminPatchUpdateLogConfigShort | [AdminPatchUpdateLogConfigShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/config/config_client.go) | [AdminPatchUpdateLogConfigShort](../../services-api/pkg/service/sessionhistory/config.go) | [AdminPatchUpdateLogConfigShort](../../samples/cli/cmd/sessionhistory/config/adminPatchUpdateLogConfig.go) |

### Game Session Detail Wrapper:  [GameSessionDetail](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/sessionhistory/v1/admin/namespaces/{namespace}/gamesessions` | GET | AdminQueryGameSessionDetailShort | [AdminQueryGameSessionDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminQueryGameSessionDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminQueryGameSessionDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminQueryGameSessionDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/gamesessions/{sessionId}` | GET | GetGameSessionDetailShort | [GetGameSessionDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [GetGameSessionDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [GetGameSessionDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/getGameSessionDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/matchmaking` | GET | AdminQueryMatchmakingDetailShort | [AdminQueryMatchmakingDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminQueryMatchmakingDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminQueryMatchmakingDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminQueryMatchmakingDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/matchmaking/session/{sessionId}` | GET | AdminGetMatchmakingDetailBySessionIDShort | [AdminGetMatchmakingDetailBySessionIDShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminGetMatchmakingDetailBySessionIDShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminGetMatchmakingDetailBySessionIDShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminGetMatchmakingDetailBySessionID.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/matchmaking/ticket/{ticketId}` | GET | AdminGetMatchmakingDetailByTicketIDShort | [AdminGetMatchmakingDetailByTicketIDShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminGetMatchmakingDetailByTicketIDShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminGetMatchmakingDetailByTicketIDShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminGetMatchmakingDetailByTicketID.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/parties` | GET | AdminQueryPartyDetailShort | [AdminQueryPartyDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminQueryPartyDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminQueryPartyDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminQueryPartyDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/parties/{sessionId}` | GET | GetPartyDetailShort | [GetPartyDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [GetPartyDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [GetPartyDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/getPartyDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/tickets` | GET | AdminQueryTicketDetailShort | [AdminQueryTicketDetailShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminQueryTicketDetailShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminQueryTicketDetailShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminQueryTicketDetail.go) |
| `/sessionhistory/v1/admin/namespaces/{namespace}/tickets/{ticketId}` | GET | AdminTicketDetailGetByTicketIDShort | [AdminTicketDetailGetByTicketIDShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [AdminTicketDetailGetByTicketIDShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [AdminTicketDetailGetByTicketIDShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/adminTicketDetailGetByTicketID.go) |
| `/sessionhistory/v1/public/namespaces/{namespace}/users/me/gamesessions` | GET | PublicQueryGameSessionMeShort | [PublicQueryGameSessionMeShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/game_session_detail/game_session_detail_client.go) | [PublicQueryGameSessionMeShort](../../services-api/pkg/service/sessionhistory/gameSessionDetail.go) | [PublicQueryGameSessionMeShort](../../samples/cli/cmd/sessionhistory/gameSessionDetail/publicQueryGameSessionMe.go) |

### XRay Wrapper:  [XRay](../../services-api/pkg/service/sessionhistory/xRay.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/match-pools/{poolName}` | GET | QueryXrayMatchPoolShort | [QueryXrayMatchPoolShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryXrayMatchPoolShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryXrayMatchPoolShort](../../samples/cli/cmd/sessionhistory/xRay/queryXrayMatchPool.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/match-pools/{poolName}/pods/{podName}/ticks` | GET | QueryDetailTickMatchPoolShort | [QueryDetailTickMatchPoolShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryDetailTickMatchPoolShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryDetailTickMatchPoolShort](../../samples/cli/cmd/sessionhistory/xRay/queryDetailTickMatchPool.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/match-pools/{poolName}/pods/{podName}/ticks/{tickId}/matches` | GET | QueryDetailTickMatchPoolMatchesShort | [QueryDetailTickMatchPoolMatchesShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryDetailTickMatchPoolMatchesShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryDetailTickMatchPoolMatchesShort](../../samples/cli/cmd/sessionhistory/xRay/queryDetailTickMatchPoolMatches.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/match-pools/{poolName}/pods/{podName}/ticks/{tickId}/tickets` | GET | QueryDetailTickMatchPoolTicketShort | [QueryDetailTickMatchPoolTicketShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryDetailTickMatchPoolTicketShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryDetailTickMatchPoolTicketShort](../../samples/cli/cmd/sessionhistory/xRay/queryDetailTickMatchPoolTicket.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/matches/{matchId}/histories` | GET | QueryMatchHistoriesShort | [QueryMatchHistoriesShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryMatchHistoriesShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryMatchHistoriesShort](../../samples/cli/cmd/sessionhistory/xRay/queryMatchHistories.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/matches/{matchId}/ticket-histories` | GET | QueryMatchTicketHistoriesShort | [QueryMatchTicketHistoriesShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryMatchTicketHistoriesShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryMatchTicketHistoriesShort](../../samples/cli/cmd/sessionhistory/xRay/queryMatchTicketHistories.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/matches/{matchId}/tickets` | GET | QueryXrayMatchShort | [QueryXrayMatchShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryXrayMatchShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryXrayMatchShort](../../samples/cli/cmd/sessionhistory/xRay/queryXrayMatch.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/acquiring-ds` | GET | QueryAcquiringDSShort | [QueryAcquiringDSShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryAcquiringDSShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryAcquiringDSShort](../../samples/cli/cmd/sessionhistory/xRay/queryAcquiringDS.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/acquiring-ds-wait-time-avg` | GET | QueryAcquiringDSWaitTimeAvgShort | [QueryAcquiringDSWaitTimeAvgShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryAcquiringDSWaitTimeAvgShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryAcquiringDSWaitTimeAvgShort](../../samples/cli/cmd/sessionhistory/xRay/queryAcquiringDSWaitTimeAvg.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/match-length-duration-avg` | GET | QueryMatchLengthDurationpAvgShort | [QueryMatchLengthDurationpAvgShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryMatchLengthDurationpAvgShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryMatchLengthDurationpAvgShort](../../samples/cli/cmd/sessionhistory/xRay/queryMatchLengthDurationpAvg.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/match-length-duration-p99` | GET | QueryMatchLengthDurationp99Short | [QueryMatchLengthDurationp99Short](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryMatchLengthDurationp99Short](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryMatchLengthDurationp99Short](../../samples/cli/cmd/sessionhistory/xRay/queryMatchLengthDurationp99.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-active-session` | GET | QueryTotalActiveSessionShort | [QueryTotalActiveSessionShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalActiveSessionShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalActiveSessionShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalActiveSession.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-match` | GET | QueryTotalMatchmakingMatchShort | [QueryTotalMatchmakingMatchShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalMatchmakingMatchShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalMatchmakingMatchShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalMatchmakingMatch.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-player-persession-avg` | GET | QueryTotalPlayerPersessionShort | [QueryTotalPlayerPersessionShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalPlayerPersessionShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalPlayerPersessionShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalPlayerPersession.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-ticket-canceled` | GET | QueryTotalMatchmakingCanceledShort | [QueryTotalMatchmakingCanceledShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalMatchmakingCanceledShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalMatchmakingCanceledShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalMatchmakingCanceled.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-ticket-created` | GET | QueryTotalMatchmakingCreatedShort | [QueryTotalMatchmakingCreatedShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalMatchmakingCreatedShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalMatchmakingCreatedShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalMatchmakingCreated.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-ticket-expired` | GET | QueryTotalMatchmakingExpiredShort | [QueryTotalMatchmakingExpiredShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalMatchmakingExpiredShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalMatchmakingExpiredShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalMatchmakingExpired.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/metrics/total-ticket-match` | GET | QueryTotalMatchmakingMatchTicketShort | [QueryTotalMatchmakingMatchTicketShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryTotalMatchmakingMatchTicketShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryTotalMatchmakingMatchTicketShort](../../samples/cli/cmd/sessionhistory/xRay/queryTotalMatchmakingMatchTicket.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/tickets` | POST | CreateXrayTicketObservabilityShort | [CreateXrayTicketObservabilityShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [CreateXrayTicketObservabilityShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [CreateXrayTicketObservabilityShort](../../samples/cli/cmd/sessionhistory/xRay/createXrayTicketObservability.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/tickets/bulk` | POST | CreateXrayBulkTicketObservabilityShort | [CreateXrayBulkTicketObservabilityShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [CreateXrayBulkTicketObservabilityShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [CreateXrayBulkTicketObservabilityShort](../../samples/cli/cmd/sessionhistory/xRay/createXrayBulkTicketObservability.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/tickets/{ticketId}` | GET | QueryXrayTimelineByTicketIDShort | [QueryXrayTimelineByTicketIDShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryXrayTimelineByTicketIDShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryXrayTimelineByTicketIDShort](../../samples/cli/cmd/sessionhistory/xRay/queryXrayTimelineByTicketID.go) |
| `/sessionhistory/v2/admin/namespaces/{namespace}/xray/users/{userId}/tickets` | GET | QueryXrayTimelineByUserIDShort | [QueryXrayTimelineByUserIDShort](../../sessionhistory-sdk/pkg/sessionhistoryclient/x_ray/x_ray_client.go) | [QueryXrayTimelineByUserIDShort](../../services-api/pkg/service/sessionhistory/xRay.go) | [QueryXrayTimelineByUserIDShort](../../samples/cli/cmd/sessionhistory/xRay/queryXrayTimelineByUserID.go) |


&nbsp;  

## Models

| Model Struct | Class |
|---|---|
| `apimodels.AcquiringDS` | [ApimodelsAcquiringDS ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_acquiring_d_s.go) |
| `apimodels.AcquiringDsWaitTime` | [ApimodelsAcquiringDsWaitTime ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_acquiring_ds_wait_time.go) |
| `apimodels.CanceledMatchmakingTicket` | [ApimodelsCanceledMatchmakingTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_canceled_matchmaking_ticket.go) |
| `apimodels.CreatedMatchmakingTicket` | [ApimodelsCreatedMatchmakingTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_created_matchmaking_ticket.go) |
| `apimodels.EventMatchHistory` | [ApimodelsEventMatchHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_event_match_history.go) |
| `apimodels.ExpiredMatchmakingTicket` | [ApimodelsExpiredMatchmakingTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_expired_matchmaking_ticket.go) |
| `apimodels.GameSessionDetail` | [ApimodelsGameSessionDetail ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_game_session_detail.go) |
| `apimodels.GameSessionDetailQueryResponse` | [ApimodelsGameSessionDetailQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_game_session_detail_query_response.go) |
| `apimodels.History` | [ApimodelsHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_history.go) |
| `apimodels.MatchLengthDuration` | [ApimodelsMatchLengthDuration ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_match_length_duration.go) |
| `apimodels.MatchMatchmaking` | [ApimodelsMatchMatchmaking ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_match_matchmaking.go) |
| `apimodels.MatchMatchmakingTicket` | [ApimodelsMatchMatchmakingTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_match_matchmaking_ticket.go) |
| `apimodels.MatchmakingDetail` | [ApimodelsMatchmakingDetail ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_matchmaking_detail.go) |
| `apimodels.MatchmakingDetailQueryResponse` | [ApimodelsMatchmakingDetailQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_matchmaking_detail_query_response.go) |
| `apimodels.MatchmakingHistory` | [ApimodelsMatchmakingHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_matchmaking_history.go) |
| `apimodels.Pagination` | [ApimodelsPagination ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_pagination.go) |
| `apimodels.PartyDetail` | [ApimodelsPartyDetail ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_party_detail.go) |
| `apimodels.PartyDetailQueryResponse` | [ApimodelsPartyDetailQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_party_detail_query_response.go) |
| `apimodels.PartyHistory` | [ApimodelsPartyHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_party_history.go) |
| `apimodels.TicketDetailQueryResponse` | [ApimodelsTicketDetailQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_ticket_detail_query_response.go) |
| `apimodels.TicketObservabilityDetail` | [ApimodelsTicketObservabilityDetail ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_ticket_observability_detail.go) |
| `apimodels.TicketObservabilityHistory` | [ApimodelsTicketObservabilityHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_ticket_observability_history.go) |
| `apimodels.TotalActiveSession` | [ApimodelsTotalActiveSession ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_total_active_session.go) |
| `apimodels.TotalPlayerPersession` | [ApimodelsTotalPlayerPersession ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_total_player_persession.go) |
| `apimodels.XRayAcquiringDsQueryResponse` | [ApimodelsXRayAcquiringDsQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_acquiring_ds_query_response.go) |
| `apimodels.XRayAcquiringDsWaitTimeQueryResponse` | [ApimodelsXRayAcquiringDsWaitTimeQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_acquiring_ds_wait_time_query_response.go) |
| `apimodels.XRayBulkTicketObservabilityRequest` | [ApimodelsXRayBulkTicketObservabilityRequest ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_bulk_ticket_observability_request.go) |
| `apimodels.XRayBulkTicketObservabilityResponse` | [ApimodelsXRayBulkTicketObservabilityResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_bulk_ticket_observability_response.go) |
| `apimodels.XRayCanceledMatchmakingTicketQueryResponse` | [ApimodelsXRayCanceledMatchmakingTicketQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_canceled_matchmaking_ticket_query_response.go) |
| `apimodels.XRayCreatedMatchmakingTicketQueryResponse` | [ApimodelsXRayCreatedMatchmakingTicketQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_created_matchmaking_ticket_query_response.go) |
| `apimodels.XRayExpiredMatchmakingTicketQueryResponse` | [ApimodelsXRayExpiredMatchmakingTicketQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_expired_matchmaking_ticket_query_response.go) |
| `apimodels.XRayMatchHistorQueryResponse` | [ApimodelsXRayMatchHistorQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_histor_query_response.go) |
| `apimodels.XRayMatchLengthDurationQueryResponse` | [ApimodelsXRayMatchLengthDurationQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_length_duration_query_response.go) |
| `apimodels.XRayMatchMatchmakingQueryResponse` | [ApimodelsXRayMatchMatchmakingQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_matchmaking_query_response.go) |
| `apimodels.XRayMatchMatchmakingTicketQueryResponse` | [ApimodelsXRayMatchMatchmakingTicketQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_matchmaking_ticket_query_response.go) |
| `apimodels.XRayMatchPoolPodTickMatchResponse` | [ApimodelsXRayMatchPoolPodTickMatchResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_pod_tick_match_response.go) |
| `apimodels.XRayMatchPoolPodTickQueryResponse` | [ApimodelsXRayMatchPoolPodTickQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_pod_tick_query_response.go) |
| `apimodels.XRayMatchPoolPodTickResult` | [ApimodelsXRayMatchPoolPodTickResult ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_pod_tick_result.go) |
| `apimodels.XRayMatchPoolPodTickTicketResponse` | [ApimodelsXRayMatchPoolPodTickTicketResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_pod_tick_ticket_response.go) |
| `apimodels.XRayMatchPoolQueryResponse` | [ApimodelsXRayMatchPoolQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_query_response.go) |
| `apimodels.XRayMatchPoolResult` | [ApimodelsXRayMatchPoolResult ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_pool_result.go) |
| `apimodels.XRayMatchTicketHistory` | [ApimodelsXRayMatchTicketHistory ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_ticket_history.go) |
| `apimodels.XRayMatchTicketHistoryQueryResponse` | [ApimodelsXRayMatchTicketHistoryQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_match_ticket_history_query_response.go) |
| `apimodels.XRayMatchesQueryResponse` | [ApimodelsXRayMatchesQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_matches_query_response.go) |
| `apimodels.XRayTicketMatchesResult` | [ApimodelsXRayTicketMatchesResult ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_ticket_matches_result.go) |
| `apimodels.XRayTicketObservabilityRequest` | [ApimodelsXRayTicketObservabilityRequest ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_ticket_observability_request.go) |
| `apimodels.XRayTicketObservabilityResponse` | [ApimodelsXRayTicketObservabilityResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_ticket_observability_response.go) |
| `apimodels.XRayTicketQueryResponse` | [ApimodelsXRayTicketQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_ticket_query_response.go) |
| `apimodels.XRayTicketResult` | [ApimodelsXRayTicketResult ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_ticket_result.go) |
| `apimodels.XRayTotalActiveSessionQueryResponse` | [ApimodelsXRayTotalActiveSessionQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_total_active_session_query_response.go) |
| `apimodels.XRayTotalPlayerPersessionAVGQueryResponse` | [ApimodelsXRayTotalPlayerPersessionAVGQueryResponse ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/apimodels_x_ray_total_player_persession_a_v_g_query_response.go) |
| `logconfig.Configuration` | [LogconfigConfiguration ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/logconfig_configuration.go) |
| `models.AllianceRule` | [ModelsAllianceRule ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_alliance_rule.go) |
| `models.BackfillProposal` | [ModelsBackfillProposal ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_backfill_proposal.go) |
| `models.BackfillTicket` | [ModelsBackfillTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_backfill_ticket.go) |
| `models.DSInformation` | [ModelsDSInformation ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_d_s_information.go) |
| `models.GameServer` | [ModelsGameServer ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_game_server.go) |
| `models.GameSession` | [ModelsGameSession ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_game_session.go) |
| `models.GameSessionTeam` | [ModelsGameSessionTeam ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_game_session_team.go) |
| `models.Match` | [ModelsMatch ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_match.go) |
| `models.MatchTicket` | [ModelsMatchTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_match_ticket.go) |
| `models.MatchingAlly` | [ModelsMatchingAlly ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_matching_ally.go) |
| `models.MatchingParty` | [ModelsMatchingParty ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_matching_party.go) |
| `models.MatchingRule` | [ModelsMatchingRule ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_matching_rule.go) |
| `models.MatchmakingResult` | [ModelsMatchmakingResult ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_matchmaking_result.go) |
| `models.Party` | [ModelsParty ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_party.go) |
| `models.PartyMember` | [ModelsPartyMember ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_party_member.go) |
| `models.PartyMembers` | [ModelsPartyMembers ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_party_members.go) |
| `models.PartyTeam` | [ModelsPartyTeam ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_party_team.go) |
| `models.PlayerData` | [ModelsPlayerData ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_player_data.go) |
| `models.ProposedProposal` | [ModelsProposedProposal ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_proposed_proposal.go) |
| `models.SessionConfig` | [ModelsSessionConfig ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_session_config.go) |
| `models.SessionConfiguration` | [ModelsSessionConfiguration ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_session_configuration.go) |
| `models.Team` | [ModelsTeam ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_team.go) |
| `models.Ticket` | [ModelsTicket ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_ticket.go) |
| `models.TicketData` | [ModelsTicketData ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_ticket_data.go) |
| `models.TicketObservability` | [ModelsTicketObservability ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_ticket_observability.go) |
| `models.TicketStatus` | [ModelsTicketStatus ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_ticket_status.go) |
| `models.User` | [ModelsUser ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/models_user.go) |
| `response.Error` | [ResponseError ](../../sessionhistory-sdk/pkg/sessionhistoryclientmodels/response_error.go) |