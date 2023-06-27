// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchCheckLayerAvailability struct {
}

func (*validateOpBatchCheckLayerAvailability) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchCheckLayerAvailability) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchCheckLayerAvailabilityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchCheckLayerAvailabilityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDeleteImage struct {
}

func (*validateOpBatchDeleteImage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDeleteImage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDeleteImageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDeleteImageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCompleteLayerUpload struct {
}

func (*validateOpCompleteLayerUpload) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCompleteLayerUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CompleteLayerUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCompleteLayerUploadInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRepository struct {
}

func (*validateOpCreateRepository) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRepositoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRepository struct {
}

func (*validateOpDeleteRepository) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRepository) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRepositoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRepositoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRepositoryPolicy struct {
}

func (*validateOpDeleteRepositoryPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRepositoryPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRepositoryPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRepositoryPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeImages struct {
}

func (*validateOpDescribeImages) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeImages) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeImagesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeImagesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeImageTags struct {
}

func (*validateOpDescribeImageTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeImageTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeImageTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeImageTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRepositoryCatalogData struct {
}

func (*validateOpGetRepositoryCatalogData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRepositoryCatalogData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRepositoryCatalogDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRepositoryCatalogDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRepositoryPolicy struct {
}

func (*validateOpGetRepositoryPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRepositoryPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRepositoryPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRepositoryPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInitiateLayerUpload struct {
}

func (*validateOpInitiateLayerUpload) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInitiateLayerUpload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InitiateLayerUploadInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInitiateLayerUploadInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutImage struct {
}

func (*validateOpPutImage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutImage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutImageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutImageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRepositoryCatalogData struct {
}

func (*validateOpPutRepositoryCatalogData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRepositoryCatalogData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRepositoryCatalogDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRepositoryCatalogDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetRepositoryPolicy struct {
}

func (*validateOpSetRepositoryPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetRepositoryPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetRepositoryPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetRepositoryPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUploadLayerPart struct {
}

func (*validateOpUploadLayerPart) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUploadLayerPart) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UploadLayerPartInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUploadLayerPartInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchCheckLayerAvailabilityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchCheckLayerAvailability{}, middleware.After)
}

func addOpBatchDeleteImageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchDeleteImage{}, middleware.After)
}

func addOpCompleteLayerUploadValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCompleteLayerUpload{}, middleware.After)
}

func addOpCreateRepositoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRepository{}, middleware.After)
}

func addOpDeleteRepositoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRepository{}, middleware.After)
}

func addOpDeleteRepositoryPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRepositoryPolicy{}, middleware.After)
}

func addOpDescribeImagesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeImages{}, middleware.After)
}

func addOpDescribeImageTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeImageTags{}, middleware.After)
}

func addOpGetRepositoryCatalogDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRepositoryCatalogData{}, middleware.After)
}

func addOpGetRepositoryPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRepositoryPolicy{}, middleware.After)
}

func addOpInitiateLayerUploadValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInitiateLayerUpload{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutImageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutImage{}, middleware.After)
}

func addOpPutRepositoryCatalogDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRepositoryCatalogData{}, middleware.After)
}

func addOpSetRepositoryPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetRepositoryPolicy{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUploadLayerPartValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUploadLayerPart{}, middleware.After)
}

func validateOpBatchCheckLayerAvailabilityInput(v *BatchCheckLayerAvailabilityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchCheckLayerAvailabilityInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.LayerDigests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LayerDigests"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchDeleteImageInput(v *BatchDeleteImageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDeleteImageInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.ImageIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImageIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCompleteLayerUploadInput(v *CompleteLayerUploadInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CompleteLayerUploadInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.UploadId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UploadId"))
	}
	if v.LayerDigests == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LayerDigests"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRepositoryInput(v *CreateRepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRepositoryInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRepositoryInput(v *DeleteRepositoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRepositoryInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRepositoryPolicyInput(v *DeleteRepositoryPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRepositoryPolicyInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeImagesInput(v *DescribeImagesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeImagesInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeImageTagsInput(v *DescribeImageTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeImageTagsInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRepositoryCatalogDataInput(v *GetRepositoryCatalogDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRepositoryCatalogDataInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRepositoryPolicyInput(v *GetRepositoryPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRepositoryPolicyInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInitiateLayerUploadInput(v *InitiateLayerUploadInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InitiateLayerUploadInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutImageInput(v *PutImageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutImageInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.ImageManifest == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ImageManifest"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRepositoryCatalogDataInput(v *PutRepositoryCatalogDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRepositoryCatalogDataInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.CatalogData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CatalogData"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetRepositoryPolicyInput(v *SetRepositoryPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetRepositoryPolicyInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.PolicyText == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyText"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUploadLayerPartInput(v *UploadLayerPartInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UploadLayerPartInput"}
	if v.RepositoryName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RepositoryName"))
	}
	if v.UploadId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UploadId"))
	}
	if v.PartFirstByte == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartFirstByte"))
	}
	if v.PartLastByte == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartLastByte"))
	}
	if v.LayerPartBlob == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LayerPartBlob"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
