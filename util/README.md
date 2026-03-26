# Util Package

The `util` package provides a centralized library of generic, cross-cutting helper functions required across various microservices and domains. It reduces code duplication and standardizes primitive manipulation.

## Core Capabilities

The `IUtil` interface guarantees access to the following domain suites:

### 1. Type Casting & Parsing
Safely transform values between unstructured (`any`) types and core primitives (`string`, `float64`, `int64`, etc.) using explicit reflection and JSON marshaling fallbacks.
- `ParseAnyToString`, `MustParseAnyToString`
- `ParseStringToAny`, `ParseAnyToAny`
- `ParseFloat64`, `ParseInt64`, `ParseFloat64With2Decimal`

### 2. Time Operations
Time manipulation unified via the `github.com/dromara/carbon/v2` module.
- `ConvertMillisToTimeString`
- `ParseStringToTime`
- `GetEndOfDay`
- `CheckFromAndToDateValid`

### 3. Protobuf & UUID Converters
Helper functions transforming domains standard into Protobuf/gRPC compatible models or canonical primitive strings.
- `UUIDPtrToStringPtr`, `StringPtrToUUIDPtr`
- `ConvertStringToTimestampPb`

### 4. Text & Randomization
- `GenerateRandomString`: Accepts custom rune dictionary arrays.
- `DecodeUnicode`: Normalizes text encoding.

### 5. Document Building
Includes stubs and definitions for converting structured rows and headers into streamable CSV or Excel `bytes.Buffer` exports gracefully.

## Usage
The `util.IUtil` interface encapsulates all logic, enabling mocking and dependency injection during testing.

```go
import "github.com/besanh/go-library/util"

// Safely parse an interface to string
strVal := myUtilSvc.MustParseAnyToString(someInterface)

// Generate a random string of length 10
rnd := myUtilSvc.GenerateRandomString(10, util.LETTER_RUNES)
```