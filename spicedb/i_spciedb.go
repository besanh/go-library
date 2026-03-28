package spicedb

import (
	"context"
	"fmt"

	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

type ISpiceDB interface {
	WriteSchema(ctx context.Context, schemaText string) error
}

func (s *SpiceClient) WriteSchema(ctx context.Context, schemaText string) error {
	request := &pb.WriteSchemaRequest{
		Schema: schemaText,
	}

	_, err := s.client.SchemaServiceClient.WriteSchema(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to write schema: %w", err)
	}

	return nil
}
