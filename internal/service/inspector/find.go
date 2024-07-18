// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	awstypes "github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

func findSubscriptionsByAssessmentTemplateARN(ctx context.Context, conn *inspector.Client, arn string) ([]*awstypes.Subscription, error) {
	input := &inspector.ListEventSubscriptionsInput{
		ResourceArn: aws.String(arn),
	}

	var results []*awstypes.Subscription

	err := conn.ListEventSubscriptionsPagesWithContext(ctx, input, func(page *inspector.ListEventSubscriptionsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, subscription := range page.Subscriptions {
			if subscription == nil {
				continue
			}

			if aws.ToString(subscription.ResourceArn) == arn {
				results = append(results, subscription)
			}
		}

		return !lastPage
	})

	return results, err
}
