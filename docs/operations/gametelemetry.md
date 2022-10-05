# Gametelemetry Service Index

&nbsp;

## Operations

### Operations Wrapper:  [Operations](../../services-api/pkg/service/gametelemetry/operations.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/game-telemetry/v1/admin/namespaces/{namespace}/events` | GET | GetEventsGameTelemetryV1AdminNamespacesNamespaceEventsGetShort | [GetEventsGameTelemetryV1AdminNamespacesNamespaceEventsGetShort](../../gametelemetry-sdk/pkg/gametelemetryclient/operations/operations_client.go) | [GetEventsGameTelemetryV1AdminNamespacesNamespaceEventsGetShort](../../services-api/pkg/service/gametelemetry/operations.go) | [GetEventsGameTelemetryV1AdminNamespacesNamespaceEventsGetShort](../../samples/cli/cmd/gametelemetry/operations/getEventsGameTelemetryV1AdminNamespacesNamespaceEventsGet.go) |

### Gametelemetry Operations Wrapper:  [GametelemetryOperations](../../services-api/pkg/service/gametelemetry/gametelemetryOperations.go)
| Endpoint | Method | ID | Class | Wrapper | Example |
|---|---|---|---|---|---|
| `/game-telemetry/v1/protected/events` | POST | ProtectedSaveEventsGameTelemetryV1ProtectedEventsPostShort | [ProtectedSaveEventsGameTelemetryV1ProtectedEventsPostShort](../../gametelemetry-sdk/pkg/gametelemetryclient/gametelemetry_operations/gametelemetry_operations_client.go) | [ProtectedSaveEventsGameTelemetryV1ProtectedEventsPostShort](../../services-api/pkg/service/gametelemetry/gametelemetryOperations.go) | [ProtectedSaveEventsGameTelemetryV1ProtectedEventsPostShort](../../samples/cli/cmd/gametelemetry/gametelemetryOperations/protectedSaveEventsGameTelemetryV1ProtectedEventsPost.go) |
| `/game-telemetry/v1/protected/steamIds/{steamId}/playtime` | GET | ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimeGetShort | [ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimeGetShort](../../gametelemetry-sdk/pkg/gametelemetryclient/gametelemetry_operations/gametelemetry_operations_client.go) | [ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimeGetShort](../../services-api/pkg/service/gametelemetry/gametelemetryOperations.go) | [ProtectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimeGetShort](../../samples/cli/cmd/gametelemetry/gametelemetryOperations/protectedGetPlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimeGet.go) |
| `/game-telemetry/v1/protected/steamIds/{steamId}/playtime/{playtime}` | PUT | ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutShort | [ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutShort](../../gametelemetry-sdk/pkg/gametelemetryclient/gametelemetry_operations/gametelemetry_operations_client.go) | [ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutShort](../../services-api/pkg/service/gametelemetry/gametelemetryOperations.go) | [ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutShort](../../samples/cli/cmd/gametelemetry/gametelemetryOperations/protectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePut.go) |


&nbsp;  

## Models

| Model Struct | Class |
|---|---|
| `HTTPValidationError` | [HTTPValidationError ](../../gametelemetry-sdk/pkg/gametelemetryclientmodels/http_validation_error.go) |
| `TelemetryBody` | [TelemetryBody ](../../gametelemetry-sdk/pkg/gametelemetryclientmodels/telemetry_body.go) |
| `ValidationError` | [ValidationError ](../../gametelemetry-sdk/pkg/gametelemetryclientmodels/validation_error.go) |