package null

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (t *Time) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	str := aws.StringValue(av.S)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	var err error
	if t.Time, err = time.Parse(time.RFC3339Nano, str); err != nil {
		return awserr.NewUnmarshalError(err, "Can't unmarhshal null.Time", []byte(str))
	}
	t.Valid = true
	return nil
}

func (t Time) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if t.Valid {
		s := t.Time.Format(time.RFC3339Nano)
		av.S = &s
	}
	return nil
}
