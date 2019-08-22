package null

import (
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var testDate, _ = time.Parse(time.RFC3339, "2016-05-03T17:06:26.209072Z")

func TestTime_DynamoDBAttributeValue(t *testing.T) {
	tests := []struct {
		name string
		date Time
		args map[string]*dynamodb.AttributeValue
	}{
		{
			"Null",
			Time{},
			map[string]*dynamodb.AttributeValue{
				"d": {S: nil},
			},
		},
		{
			"Valid",
			TimeFrom(testDate),
			map[string]*dynamodb.AttributeValue{
				"d": {S: aws.String("2016-05-03T17:06:26.209072Z")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual, err := dynamodbattribute.Marshal(tt.date); err != nil {
				t.Errorf("UnmarshalDynamoDBAttributeValue() error = %v", err)
			} else {
				if e, a := tt.args["d"], actual; !reflect.DeepEqual(e, a) {
					t.Errorf("UnmarshalDynamoDBAttributeValue() expect = %v, got = %v", e, a)
				}

				// if err := tt.date.UnmarshalDynamoDBAttributeValue(tt.args["d"]); (err != nil) {
				// 	t.Errorf("UnmarshalDynamoDBAttributeValue() error = %v", err)
				// }
			}
		})
	}
}
