# AccelByte Go SDK

## Setup

Add the following to your project's `go.mod`:

```
require (
	github.com/AccelByte/accelbyte-go-sdk {VERSION}
)
```

Replace `{VERSION}` with a specific release version tag. When starting a new project, 
using the latest release version is recommended.

## Usage

### Interacting with Justice HTTP Endpoints

We will use `iam-sdk` as an example.

The `iam-sdk` contains 2 directories:
 - `iamclient` contains the logic to make requests. 
 - `iamclientmodels` contains the models such as request and response models.

Client must create a struct that implement following interface:

- `ConfigRepository` is responsible to store configuration.
- `TokenRepository` is responsible to store access token.

For more details, see [samples/cli/pkg/repository](samples/cli/pkg/repository) for more details.

Depending on the HTTP endpoint authorization, there are a few ways to make HTTP requests in the wrapper:

- Basic authorization

  ```go
  // iam service
  o.Client.OAuth20.TokenGrantV3(input, client.BasicAuth(clientId, clientSecret))
  ```

- Bearer token authorization

  ```go
  // iam service
  u.Client.Users.AdminAddUserRoleV3(input, client.BearerToken(*accessToken.AccessToken))
  ```

- No authorization

  ```go
  // platform service
  i.Client.Item.PublicGetItem(input)
  ```

### Interacting with Justice WebSocket Endpoints

To interact with Justice services which use WebSocket endpoints e.g. Justice Lobby Service, client should implement `connectionutils/ConnectionManager` interface. 

`ConnectionManager` manages WebSocket connection that save, get and close the WebSocket connection. In other words, client should maintain WebSocket connection using `ConnectionManager`. For reference, see [samples/cli/pkg/utils/connectionManager.go](samples/cli/pkg/utils/connectionManager.go).

For a working example, see [samples/cli](samples/cli).