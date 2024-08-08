// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details about the repository creation templates in a registry. The
// prefixes request parameter can be used to return the details for a specific
// repository creation template.
func (c *Client) DescribeRepositoryCreationTemplates(ctx context.Context, params *DescribeRepositoryCreationTemplatesInput, optFns ...func(*Options)) (*DescribeRepositoryCreationTemplatesOutput, error) {
	if params == nil {
		params = &DescribeRepositoryCreationTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRepositoryCreationTemplates", params, optFns, c.addOperationDescribeRepositoryCreationTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRepositoryCreationTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRepositoryCreationTemplatesInput struct {

	// The maximum number of repository results returned by
	// DescribeRepositoryCreationTemplatesRequest in paginated output. When this
	// parameter is used, DescribeRepositoryCreationTemplatesRequest only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// DescribeRepositoryCreationTemplatesRequest request with the returned nextToken
	// value. This value can be between 1 and 1000. If this parameter is not used, then
	// DescribeRepositoryCreationTemplatesRequest returns up to 100 results and a
	// nextToken value, if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// DescribeRepositoryCreationTemplates request where maxResults was used and the
	// results exceeded the value of that parameter. Pagination continues from the end
	// of the previous results that returned the nextToken value. This value is null
	// when there are no more results to return.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// The repository namespace prefixes associated with the repository creation
	// templates to describe. If this value is not specified, all repository creation
	// templates are returned.
	Prefixes []string

	noSmithyDocumentSerde
}

type DescribeRepositoryCreationTemplatesOutput struct {

	// The nextToken value to include in a future DescribeRepositoryCreationTemplates
	// request. When the results of a DescribeRepositoryCreationTemplates request
	// exceed maxResults , this value can be used to retrieve the next page of results.
	// This value is null when there are no more results to return.
	NextToken *string

	// The registry ID associated with the request.
	RegistryId *string

	// The details of the repository creation templates.
	RepositoryCreationTemplates []types.RepositoryCreationTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRepositoryCreationTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeRepositoryCreationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeRepositoryCreationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRepositoryCreationTemplates"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRepositoryCreationTemplates(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeRepositoryCreationTemplatesPaginatorOptions is the paginator options
// for DescribeRepositoryCreationTemplates
type DescribeRepositoryCreationTemplatesPaginatorOptions struct {
	// The maximum number of repository results returned by
	// DescribeRepositoryCreationTemplatesRequest in paginated output. When this
	// parameter is used, DescribeRepositoryCreationTemplatesRequest only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// DescribeRepositoryCreationTemplatesRequest request with the returned nextToken
	// value. This value can be between 1 and 1000. If this parameter is not used, then
	// DescribeRepositoryCreationTemplatesRequest returns up to 100 results and a
	// nextToken value, if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRepositoryCreationTemplatesPaginator is a paginator for
// DescribeRepositoryCreationTemplates
type DescribeRepositoryCreationTemplatesPaginator struct {
	options   DescribeRepositoryCreationTemplatesPaginatorOptions
	client    DescribeRepositoryCreationTemplatesAPIClient
	params    *DescribeRepositoryCreationTemplatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRepositoryCreationTemplatesPaginator returns a new
// DescribeRepositoryCreationTemplatesPaginator
func NewDescribeRepositoryCreationTemplatesPaginator(client DescribeRepositoryCreationTemplatesAPIClient, params *DescribeRepositoryCreationTemplatesInput, optFns ...func(*DescribeRepositoryCreationTemplatesPaginatorOptions)) *DescribeRepositoryCreationTemplatesPaginator {
	if params == nil {
		params = &DescribeRepositoryCreationTemplatesInput{}
	}

	options := DescribeRepositoryCreationTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRepositoryCreationTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRepositoryCreationTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRepositoryCreationTemplates page.
func (p *DescribeRepositoryCreationTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRepositoryCreationTemplatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeRepositoryCreationTemplates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeRepositoryCreationTemplatesAPIClient is a client that implements the
// DescribeRepositoryCreationTemplates operation.
type DescribeRepositoryCreationTemplatesAPIClient interface {
	DescribeRepositoryCreationTemplates(context.Context, *DescribeRepositoryCreationTemplatesInput, ...func(*Options)) (*DescribeRepositoryCreationTemplatesOutput, error)
}

var _ DescribeRepositoryCreationTemplatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeRepositoryCreationTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRepositoryCreationTemplates",
	}
}
