# Response Package

The `response` package standardizes RESTful API HTTP JSON payload structures across microservices. It ensures every outgoing API response conforms to a unified predictable schema.

## Architecture Flow

1. **Standard Generics**: Relies natively on Go generics (`[T any]`) to maintain robust compile-time type safety over payload wrappers.
2. **Response Factories**: Provides helper functions for emitting standard HTTP statuses (`OK`, `BadRequest`, `NotFound`, `Unauthorized`, etc.) bundled natively with appropriate pre-configured messages and codes.
3. **Unified Schema**:
   * **Base Payload**: Wraps every output inside a top-level `Body` object.
   * **Pagination Payload**: Dedicated wrappers appending `total`, `limit`, and `offset` fields structurally.

## Usage

```go
import "github.com/besanh/go-library/response"
```

### 1. Generating Success Responses

Yield standard `200 OK` JSON responses wrapping data cleanly.

```go
type User struct { ID string }

// Wraps value inside the 'data' field
res := response.OK(&User{ID: "123"})

// Override the default "success" message
resWithMessage := response.OK(&User{ID: "123"}, "User successfully retrieved")

// Yield an empty success response
emptyRes := response.OKOnly("Triggered successfully")
```

### 2. Generating Error Responses

Use explicit helper functions to generate unified structured errors mapping to designated HTTP semantics. Keep things consistent by relying on standard message codes.

```go
if !exists {
    return response.NotFound[User]() // Defaults: code=404, message="not found"
}

if invalidParams {
    return response.BadRequestWithMsg[User]("invalid_email", "The email provided is invalid.")
}
```

### 3. Generating Paginated Responses

Standardize output lists needing index offsets easily.

```go
users := []User{{ID: "1"}, {ID: "2"}}
totalCount := int64(150)
limit := int64(10)
offset := int64(0)

res := response.Pagination(users, totalCount, limit, offset, "Fetched users chunk")
```

**JSON Output Format Example:**
```json
{
  "Body": {
    "code": 200,
    "message": "Fetched users chunk",
    "data": [
      { "ID": "1" },
      { "ID": "2" }
    ],
    "total": 150,
    "limit": 10,
    "offset": 0
  }
}
```
