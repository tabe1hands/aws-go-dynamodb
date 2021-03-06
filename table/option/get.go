package option

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// The GetItemInput type is an adapter to change a parameter in
// dynamodb.GetItemInput
type GetItemInput func(req *dynamodb.GetItemInput)

// ConsistentRead enables ConsistentRead in dynamodb.GetItemInput.
func ConsistentRead() GetItemInput {
	return func(req *dynamodb.GetItemInput) {
		req.ConsistentRead = aws.Bool(true)
	}
}
