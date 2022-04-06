// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lambdaiface provides an interface to enable mocking the AWS Lambda service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lambdaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaAPI provides an interface to enable mocking the
// lambda.Lambda service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Lambda.
//    func myFunc(svc lambdaiface.LambdaAPI) bool {
//        // Make svc.AddLayerVersionPermission request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := lambda.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockLambdaClient struct {
//        lambdaiface.LambdaAPI
//    }
//    func (m *mockLambdaClient) AddLayerVersionPermission(input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockLambdaClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LambdaAPI interface {
	AddLayerVersionPermission(*lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error)
	AddLayerVersionPermissionWithContext(aws.Context, *lambda.AddLayerVersionPermissionInput, ...request.Option) (*lambda.AddLayerVersionPermissionOutput, error)
	AddLayerVersionPermissionRequest(*lambda.AddLayerVersionPermissionInput) (*request.Request, *lambda.AddLayerVersionPermissionOutput)

	AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)
	AddPermissionWithContext(aws.Context, *lambda.AddPermissionInput, ...request.Option) (*lambda.AddPermissionOutput, error)
	AddPermissionRequest(*lambda.AddPermissionInput) (*request.Request, *lambda.AddPermissionOutput)

	CreateAlias(*lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)
	CreateAliasWithContext(aws.Context, *lambda.CreateAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	CreateAliasRequest(*lambda.CreateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	CreateCodeSigningConfig(*lambda.CreateCodeSigningConfigInput) (*lambda.CreateCodeSigningConfigOutput, error)
	CreateCodeSigningConfigWithContext(aws.Context, *lambda.CreateCodeSigningConfigInput, ...request.Option) (*lambda.CreateCodeSigningConfigOutput, error)
	CreateCodeSigningConfigRequest(*lambda.CreateCodeSigningConfigInput) (*request.Request, *lambda.CreateCodeSigningConfigOutput)

	CreateEventSourceMapping(*lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingWithContext(aws.Context, *lambda.CreateEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingRequest(*lambda.CreateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	CreateFunction(*lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
	CreateFunctionWithContext(aws.Context, *lambda.CreateFunctionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	CreateFunctionRequest(*lambda.CreateFunctionInput) (*request.Request, *lambda.FunctionConfiguration)

	CreateFunctionUrlConfig(*lambda.CreateFunctionUrlConfigInput) (*lambda.CreateFunctionUrlConfigOutput, error)
	CreateFunctionUrlConfigWithContext(aws.Context, *lambda.CreateFunctionUrlConfigInput, ...request.Option) (*lambda.CreateFunctionUrlConfigOutput, error)
	CreateFunctionUrlConfigRequest(*lambda.CreateFunctionUrlConfigInput) (*request.Request, *lambda.CreateFunctionUrlConfigOutput)

	DeleteAlias(*lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)
	DeleteAliasWithContext(aws.Context, *lambda.DeleteAliasInput, ...request.Option) (*lambda.DeleteAliasOutput, error)
	DeleteAliasRequest(*lambda.DeleteAliasInput) (*request.Request, *lambda.DeleteAliasOutput)

	DeleteCodeSigningConfig(*lambda.DeleteCodeSigningConfigInput) (*lambda.DeleteCodeSigningConfigOutput, error)
	DeleteCodeSigningConfigWithContext(aws.Context, *lambda.DeleteCodeSigningConfigInput, ...request.Option) (*lambda.DeleteCodeSigningConfigOutput, error)
	DeleteCodeSigningConfigRequest(*lambda.DeleteCodeSigningConfigInput) (*request.Request, *lambda.DeleteCodeSigningConfigOutput)

	DeleteEventSourceMapping(*lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingWithContext(aws.Context, *lambda.DeleteEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingRequest(*lambda.DeleteEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	DeleteFunction(*lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionWithContext(aws.Context, *lambda.DeleteFunctionInput, ...request.Option) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionRequest(*lambda.DeleteFunctionInput) (*request.Request, *lambda.DeleteFunctionOutput)

	DeleteFunctionCodeSigningConfig(*lambda.DeleteFunctionCodeSigningConfigInput) (*lambda.DeleteFunctionCodeSigningConfigOutput, error)
	DeleteFunctionCodeSigningConfigWithContext(aws.Context, *lambda.DeleteFunctionCodeSigningConfigInput, ...request.Option) (*lambda.DeleteFunctionCodeSigningConfigOutput, error)
	DeleteFunctionCodeSigningConfigRequest(*lambda.DeleteFunctionCodeSigningConfigInput) (*request.Request, *lambda.DeleteFunctionCodeSigningConfigOutput)

	DeleteFunctionConcurrency(*lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyWithContext(aws.Context, *lambda.DeleteFunctionConcurrencyInput, ...request.Option) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyRequest(*lambda.DeleteFunctionConcurrencyInput) (*request.Request, *lambda.DeleteFunctionConcurrencyOutput)

	DeleteFunctionEventInvokeConfig(*lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	DeleteFunctionEventInvokeConfigWithContext(aws.Context, *lambda.DeleteFunctionEventInvokeConfigInput, ...request.Option) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	DeleteFunctionEventInvokeConfigRequest(*lambda.DeleteFunctionEventInvokeConfigInput) (*request.Request, *lambda.DeleteFunctionEventInvokeConfigOutput)

	DeleteFunctionUrlConfig(*lambda.DeleteFunctionUrlConfigInput) (*lambda.DeleteFunctionUrlConfigOutput, error)
	DeleteFunctionUrlConfigWithContext(aws.Context, *lambda.DeleteFunctionUrlConfigInput, ...request.Option) (*lambda.DeleteFunctionUrlConfigOutput, error)
	DeleteFunctionUrlConfigRequest(*lambda.DeleteFunctionUrlConfigInput) (*request.Request, *lambda.DeleteFunctionUrlConfigOutput)

	DeleteLayerVersion(*lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error)
	DeleteLayerVersionWithContext(aws.Context, *lambda.DeleteLayerVersionInput, ...request.Option) (*lambda.DeleteLayerVersionOutput, error)
	DeleteLayerVersionRequest(*lambda.DeleteLayerVersionInput) (*request.Request, *lambda.DeleteLayerVersionOutput)

	DeleteProvisionedConcurrencyConfig(*lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	DeleteProvisionedConcurrencyConfigWithContext(aws.Context, *lambda.DeleteProvisionedConcurrencyConfigInput, ...request.Option) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	DeleteProvisionedConcurrencyConfigRequest(*lambda.DeleteProvisionedConcurrencyConfigInput) (*request.Request, *lambda.DeleteProvisionedConcurrencyConfigOutput)

	GetAccountSettings(*lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsWithContext(aws.Context, *lambda.GetAccountSettingsInput, ...request.Option) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsRequest(*lambda.GetAccountSettingsInput) (*request.Request, *lambda.GetAccountSettingsOutput)

	GetAlias(*lambda.GetAliasInput) (*lambda.AliasConfiguration, error)
	GetAliasWithContext(aws.Context, *lambda.GetAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	GetAliasRequest(*lambda.GetAliasInput) (*request.Request, *lambda.AliasConfiguration)

	GetCodeSigningConfig(*lambda.GetCodeSigningConfigInput) (*lambda.GetCodeSigningConfigOutput, error)
	GetCodeSigningConfigWithContext(aws.Context, *lambda.GetCodeSigningConfigInput, ...request.Option) (*lambda.GetCodeSigningConfigOutput, error)
	GetCodeSigningConfigRequest(*lambda.GetCodeSigningConfigInput) (*request.Request, *lambda.GetCodeSigningConfigOutput)

	GetEventSourceMapping(*lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingWithContext(aws.Context, *lambda.GetEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingRequest(*lambda.GetEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	GetFunction(*lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
	GetFunctionWithContext(aws.Context, *lambda.GetFunctionInput, ...request.Option) (*lambda.GetFunctionOutput, error)
	GetFunctionRequest(*lambda.GetFunctionInput) (*request.Request, *lambda.GetFunctionOutput)

	GetFunctionCodeSigningConfig(*lambda.GetFunctionCodeSigningConfigInput) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionCodeSigningConfigWithContext(aws.Context, *lambda.GetFunctionCodeSigningConfigInput, ...request.Option) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionCodeSigningConfigRequest(*lambda.GetFunctionCodeSigningConfigInput) (*request.Request, *lambda.GetFunctionCodeSigningConfigOutput)

	GetFunctionConcurrency(*lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConcurrencyWithContext(aws.Context, *lambda.GetFunctionConcurrencyInput, ...request.Option) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConcurrencyRequest(*lambda.GetFunctionConcurrencyInput) (*request.Request, *lambda.GetFunctionConcurrencyOutput)

	GetFunctionConfiguration(*lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationWithContext(aws.Context, *lambda.GetFunctionConfigurationInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationRequest(*lambda.GetFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)

	GetFunctionEventInvokeConfig(*lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionEventInvokeConfigWithContext(aws.Context, *lambda.GetFunctionEventInvokeConfigInput, ...request.Option) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionEventInvokeConfigRequest(*lambda.GetFunctionEventInvokeConfigInput) (*request.Request, *lambda.GetFunctionEventInvokeConfigOutput)

	GetFunctionUrlConfig(*lambda.GetFunctionUrlConfigInput) (*lambda.GetFunctionUrlConfigOutput, error)
	GetFunctionUrlConfigWithContext(aws.Context, *lambda.GetFunctionUrlConfigInput, ...request.Option) (*lambda.GetFunctionUrlConfigOutput, error)
	GetFunctionUrlConfigRequest(*lambda.GetFunctionUrlConfigInput) (*request.Request, *lambda.GetFunctionUrlConfigOutput)

	GetLayerVersion(*lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionWithContext(aws.Context, *lambda.GetLayerVersionInput, ...request.Option) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionRequest(*lambda.GetLayerVersionInput) (*request.Request, *lambda.GetLayerVersionOutput)

	GetLayerVersionByArn(*lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionByArnWithContext(aws.Context, *lambda.GetLayerVersionByArnInput, ...request.Option) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionByArnRequest(*lambda.GetLayerVersionByArnInput) (*request.Request, *lambda.GetLayerVersionByArnOutput)

	GetLayerVersionPolicy(*lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error)
	GetLayerVersionPolicyWithContext(aws.Context, *lambda.GetLayerVersionPolicyInput, ...request.Option) (*lambda.GetLayerVersionPolicyOutput, error)
	GetLayerVersionPolicyRequest(*lambda.GetLayerVersionPolicyInput) (*request.Request, *lambda.GetLayerVersionPolicyOutput)

	GetPolicy(*lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *lambda.GetPolicyInput, ...request.Option) (*lambda.GetPolicyOutput, error)
	GetPolicyRequest(*lambda.GetPolicyInput) (*request.Request, *lambda.GetPolicyOutput)

	GetProvisionedConcurrencyConfig(*lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetProvisionedConcurrencyConfigWithContext(aws.Context, *lambda.GetProvisionedConcurrencyConfigInput, ...request.Option) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetProvisionedConcurrencyConfigRequest(*lambda.GetProvisionedConcurrencyConfigInput) (*request.Request, *lambda.GetProvisionedConcurrencyConfigOutput)

	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)
	InvokeWithContext(aws.Context, *lambda.InvokeInput, ...request.Option) (*lambda.InvokeOutput, error)
	InvokeRequest(*lambda.InvokeInput) (*request.Request, *lambda.InvokeOutput)

	InvokeAsync(*lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error)
	InvokeAsyncWithContext(aws.Context, *lambda.InvokeAsyncInput, ...request.Option) (*lambda.InvokeAsyncOutput, error)
	InvokeAsyncRequest(*lambda.InvokeAsyncInput) (*request.Request, *lambda.InvokeAsyncOutput)

	ListAliases(*lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)
	ListAliasesWithContext(aws.Context, *lambda.ListAliasesInput, ...request.Option) (*lambda.ListAliasesOutput, error)
	ListAliasesRequest(*lambda.ListAliasesInput) (*request.Request, *lambda.ListAliasesOutput)

	ListAliasesPages(*lambda.ListAliasesInput, func(*lambda.ListAliasesOutput, bool) bool) error
	ListAliasesPagesWithContext(aws.Context, *lambda.ListAliasesInput, func(*lambda.ListAliasesOutput, bool) bool, ...request.Option) error

	ListCodeSigningConfigs(*lambda.ListCodeSigningConfigsInput) (*lambda.ListCodeSigningConfigsOutput, error)
	ListCodeSigningConfigsWithContext(aws.Context, *lambda.ListCodeSigningConfigsInput, ...request.Option) (*lambda.ListCodeSigningConfigsOutput, error)
	ListCodeSigningConfigsRequest(*lambda.ListCodeSigningConfigsInput) (*request.Request, *lambda.ListCodeSigningConfigsOutput)

	ListCodeSigningConfigsPages(*lambda.ListCodeSigningConfigsInput, func(*lambda.ListCodeSigningConfigsOutput, bool) bool) error
	ListCodeSigningConfigsPagesWithContext(aws.Context, *lambda.ListCodeSigningConfigsInput, func(*lambda.ListCodeSigningConfigsOutput, bool) bool, ...request.Option) error

	ListEventSourceMappings(*lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsWithContext(aws.Context, *lambda.ListEventSourceMappingsInput, ...request.Option) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsRequest(*lambda.ListEventSourceMappingsInput) (*request.Request, *lambda.ListEventSourceMappingsOutput)

	ListEventSourceMappingsPages(*lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool) error
	ListEventSourceMappingsPagesWithContext(aws.Context, *lambda.ListEventSourceMappingsInput, func(*lambda.ListEventSourceMappingsOutput, bool) bool, ...request.Option) error

	ListFunctionEventInvokeConfigs(*lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionEventInvokeConfigsWithContext(aws.Context, *lambda.ListFunctionEventInvokeConfigsInput, ...request.Option) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionEventInvokeConfigsRequest(*lambda.ListFunctionEventInvokeConfigsInput) (*request.Request, *lambda.ListFunctionEventInvokeConfigsOutput)

	ListFunctionEventInvokeConfigsPages(*lambda.ListFunctionEventInvokeConfigsInput, func(*lambda.ListFunctionEventInvokeConfigsOutput, bool) bool) error
	ListFunctionEventInvokeConfigsPagesWithContext(aws.Context, *lambda.ListFunctionEventInvokeConfigsInput, func(*lambda.ListFunctionEventInvokeConfigsOutput, bool) bool, ...request.Option) error

	ListFunctionUrlConfigs(*lambda.ListFunctionUrlConfigsInput) (*lambda.ListFunctionUrlConfigsOutput, error)
	ListFunctionUrlConfigsWithContext(aws.Context, *lambda.ListFunctionUrlConfigsInput, ...request.Option) (*lambda.ListFunctionUrlConfigsOutput, error)
	ListFunctionUrlConfigsRequest(*lambda.ListFunctionUrlConfigsInput) (*request.Request, *lambda.ListFunctionUrlConfigsOutput)

	ListFunctionUrlConfigsPages(*lambda.ListFunctionUrlConfigsInput, func(*lambda.ListFunctionUrlConfigsOutput, bool) bool) error
	ListFunctionUrlConfigsPagesWithContext(aws.Context, *lambda.ListFunctionUrlConfigsInput, func(*lambda.ListFunctionUrlConfigsOutput, bool) bool, ...request.Option) error

	ListFunctions(*lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
	ListFunctionsWithContext(aws.Context, *lambda.ListFunctionsInput, ...request.Option) (*lambda.ListFunctionsOutput, error)
	ListFunctionsRequest(*lambda.ListFunctionsInput) (*request.Request, *lambda.ListFunctionsOutput)

	ListFunctionsPages(*lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool) error
	ListFunctionsPagesWithContext(aws.Context, *lambda.ListFunctionsInput, func(*lambda.ListFunctionsOutput, bool) bool, ...request.Option) error

	ListFunctionsByCodeSigningConfig(*lambda.ListFunctionsByCodeSigningConfigInput) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListFunctionsByCodeSigningConfigWithContext(aws.Context, *lambda.ListFunctionsByCodeSigningConfigInput, ...request.Option) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListFunctionsByCodeSigningConfigRequest(*lambda.ListFunctionsByCodeSigningConfigInput) (*request.Request, *lambda.ListFunctionsByCodeSigningConfigOutput)

	ListFunctionsByCodeSigningConfigPages(*lambda.ListFunctionsByCodeSigningConfigInput, func(*lambda.ListFunctionsByCodeSigningConfigOutput, bool) bool) error
	ListFunctionsByCodeSigningConfigPagesWithContext(aws.Context, *lambda.ListFunctionsByCodeSigningConfigInput, func(*lambda.ListFunctionsByCodeSigningConfigOutput, bool) bool, ...request.Option) error

	ListLayerVersions(*lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error)
	ListLayerVersionsWithContext(aws.Context, *lambda.ListLayerVersionsInput, ...request.Option) (*lambda.ListLayerVersionsOutput, error)
	ListLayerVersionsRequest(*lambda.ListLayerVersionsInput) (*request.Request, *lambda.ListLayerVersionsOutput)

	ListLayerVersionsPages(*lambda.ListLayerVersionsInput, func(*lambda.ListLayerVersionsOutput, bool) bool) error
	ListLayerVersionsPagesWithContext(aws.Context, *lambda.ListLayerVersionsInput, func(*lambda.ListLayerVersionsOutput, bool) bool, ...request.Option) error

	ListLayers(*lambda.ListLayersInput) (*lambda.ListLayersOutput, error)
	ListLayersWithContext(aws.Context, *lambda.ListLayersInput, ...request.Option) (*lambda.ListLayersOutput, error)
	ListLayersRequest(*lambda.ListLayersInput) (*request.Request, *lambda.ListLayersOutput)

	ListLayersPages(*lambda.ListLayersInput, func(*lambda.ListLayersOutput, bool) bool) error
	ListLayersPagesWithContext(aws.Context, *lambda.ListLayersInput, func(*lambda.ListLayersOutput, bool) bool, ...request.Option) error

	ListProvisionedConcurrencyConfigs(*lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListProvisionedConcurrencyConfigsWithContext(aws.Context, *lambda.ListProvisionedConcurrencyConfigsInput, ...request.Option) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListProvisionedConcurrencyConfigsRequest(*lambda.ListProvisionedConcurrencyConfigsInput) (*request.Request, *lambda.ListProvisionedConcurrencyConfigsOutput)

	ListProvisionedConcurrencyConfigsPages(*lambda.ListProvisionedConcurrencyConfigsInput, func(*lambda.ListProvisionedConcurrencyConfigsOutput, bool) bool) error
	ListProvisionedConcurrencyConfigsPagesWithContext(aws.Context, *lambda.ListProvisionedConcurrencyConfigsInput, func(*lambda.ListProvisionedConcurrencyConfigsOutput, bool) bool, ...request.Option) error

	ListTags(*lambda.ListTagsInput) (*lambda.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *lambda.ListTagsInput, ...request.Option) (*lambda.ListTagsOutput, error)
	ListTagsRequest(*lambda.ListTagsInput) (*request.Request, *lambda.ListTagsOutput)

	ListVersionsByFunction(*lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionWithContext(aws.Context, *lambda.ListVersionsByFunctionInput, ...request.Option) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionRequest(*lambda.ListVersionsByFunctionInput) (*request.Request, *lambda.ListVersionsByFunctionOutput)

	ListVersionsByFunctionPages(*lambda.ListVersionsByFunctionInput, func(*lambda.ListVersionsByFunctionOutput, bool) bool) error
	ListVersionsByFunctionPagesWithContext(aws.Context, *lambda.ListVersionsByFunctionInput, func(*lambda.ListVersionsByFunctionOutput, bool) bool, ...request.Option) error

	PublishLayerVersion(*lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error)
	PublishLayerVersionWithContext(aws.Context, *lambda.PublishLayerVersionInput, ...request.Option) (*lambda.PublishLayerVersionOutput, error)
	PublishLayerVersionRequest(*lambda.PublishLayerVersionInput) (*request.Request, *lambda.PublishLayerVersionOutput)

	PublishVersion(*lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)
	PublishVersionWithContext(aws.Context, *lambda.PublishVersionInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	PublishVersionRequest(*lambda.PublishVersionInput) (*request.Request, *lambda.FunctionConfiguration)

	PutFunctionCodeSigningConfig(*lambda.PutFunctionCodeSigningConfigInput) (*lambda.PutFunctionCodeSigningConfigOutput, error)
	PutFunctionCodeSigningConfigWithContext(aws.Context, *lambda.PutFunctionCodeSigningConfigInput, ...request.Option) (*lambda.PutFunctionCodeSigningConfigOutput, error)
	PutFunctionCodeSigningConfigRequest(*lambda.PutFunctionCodeSigningConfigInput) (*request.Request, *lambda.PutFunctionCodeSigningConfigOutput)

	PutFunctionConcurrency(*lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyWithContext(aws.Context, *lambda.PutFunctionConcurrencyInput, ...request.Option) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyRequest(*lambda.PutFunctionConcurrencyInput) (*request.Request, *lambda.PutFunctionConcurrencyOutput)

	PutFunctionEventInvokeConfig(*lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	PutFunctionEventInvokeConfigWithContext(aws.Context, *lambda.PutFunctionEventInvokeConfigInput, ...request.Option) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	PutFunctionEventInvokeConfigRequest(*lambda.PutFunctionEventInvokeConfigInput) (*request.Request, *lambda.PutFunctionEventInvokeConfigOutput)

	PutProvisionedConcurrencyConfig(*lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	PutProvisionedConcurrencyConfigWithContext(aws.Context, *lambda.PutProvisionedConcurrencyConfigInput, ...request.Option) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	PutProvisionedConcurrencyConfigRequest(*lambda.PutProvisionedConcurrencyConfigInput) (*request.Request, *lambda.PutProvisionedConcurrencyConfigOutput)

	RemoveLayerVersionPermission(*lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error)
	RemoveLayerVersionPermissionWithContext(aws.Context, *lambda.RemoveLayerVersionPermissionInput, ...request.Option) (*lambda.RemoveLayerVersionPermissionOutput, error)
	RemoveLayerVersionPermissionRequest(*lambda.RemoveLayerVersionPermissionInput) (*request.Request, *lambda.RemoveLayerVersionPermissionOutput)

	RemovePermission(*lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *lambda.RemovePermissionInput, ...request.Option) (*lambda.RemovePermissionOutput, error)
	RemovePermissionRequest(*lambda.RemovePermissionInput) (*request.Request, *lambda.RemovePermissionOutput)

	TagResource(*lambda.TagResourceInput) (*lambda.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lambda.TagResourceInput, ...request.Option) (*lambda.TagResourceOutput, error)
	TagResourceRequest(*lambda.TagResourceInput) (*request.Request, *lambda.TagResourceOutput)

	UntagResource(*lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lambda.UntagResourceInput, ...request.Option) (*lambda.UntagResourceOutput, error)
	UntagResourceRequest(*lambda.UntagResourceInput) (*request.Request, *lambda.UntagResourceOutput)

	UpdateAlias(*lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)
	UpdateAliasWithContext(aws.Context, *lambda.UpdateAliasInput, ...request.Option) (*lambda.AliasConfiguration, error)
	UpdateAliasRequest(*lambda.UpdateAliasInput) (*request.Request, *lambda.AliasConfiguration)

	UpdateCodeSigningConfig(*lambda.UpdateCodeSigningConfigInput) (*lambda.UpdateCodeSigningConfigOutput, error)
	UpdateCodeSigningConfigWithContext(aws.Context, *lambda.UpdateCodeSigningConfigInput, ...request.Option) (*lambda.UpdateCodeSigningConfigOutput, error)
	UpdateCodeSigningConfigRequest(*lambda.UpdateCodeSigningConfigInput) (*request.Request, *lambda.UpdateCodeSigningConfigOutput)

	UpdateEventSourceMapping(*lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingWithContext(aws.Context, *lambda.UpdateEventSourceMappingInput, ...request.Option) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingRequest(*lambda.UpdateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration)

	UpdateFunctionCode(*lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeWithContext(aws.Context, *lambda.UpdateFunctionCodeInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeRequest(*lambda.UpdateFunctionCodeInput) (*request.Request, *lambda.FunctionConfiguration)

	UpdateFunctionConfiguration(*lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationWithContext(aws.Context, *lambda.UpdateFunctionConfigurationInput, ...request.Option) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationRequest(*lambda.UpdateFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration)

	UpdateFunctionEventInvokeConfig(*lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
	UpdateFunctionEventInvokeConfigWithContext(aws.Context, *lambda.UpdateFunctionEventInvokeConfigInput, ...request.Option) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
	UpdateFunctionEventInvokeConfigRequest(*lambda.UpdateFunctionEventInvokeConfigInput) (*request.Request, *lambda.UpdateFunctionEventInvokeConfigOutput)

	UpdateFunctionUrlConfig(*lambda.UpdateFunctionUrlConfigInput) (*lambda.UpdateFunctionUrlConfigOutput, error)
	UpdateFunctionUrlConfigWithContext(aws.Context, *lambda.UpdateFunctionUrlConfigInput, ...request.Option) (*lambda.UpdateFunctionUrlConfigOutput, error)
	UpdateFunctionUrlConfigRequest(*lambda.UpdateFunctionUrlConfigInput) (*request.Request, *lambda.UpdateFunctionUrlConfigOutput)

	WaitUntilFunctionActive(*lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionActiveWithContext(aws.Context, *lambda.GetFunctionConfigurationInput, ...request.WaiterOption) error

	WaitUntilFunctionActiveV2(*lambda.GetFunctionInput) error
	WaitUntilFunctionActiveV2WithContext(aws.Context, *lambda.GetFunctionInput, ...request.WaiterOption) error

	WaitUntilFunctionExists(*lambda.GetFunctionInput) error
	WaitUntilFunctionExistsWithContext(aws.Context, *lambda.GetFunctionInput, ...request.WaiterOption) error

	WaitUntilFunctionUpdated(*lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionUpdatedWithContext(aws.Context, *lambda.GetFunctionConfigurationInput, ...request.WaiterOption) error

	WaitUntilFunctionUpdatedV2(*lambda.GetFunctionInput) error
	WaitUntilFunctionUpdatedV2WithContext(aws.Context, *lambda.GetFunctionInput, ...request.WaiterOption) error
}

var _ LambdaAPI = (*lambda.Lambda)(nil)
