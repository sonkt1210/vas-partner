package repositories

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type SignOutResult struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type BalanceResult struct {
	DataValue int64 `xml:"dataValue,omitempty" json:"dataValue,omitempty"`

	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	TransID int32 `xml:"transID,omitempty" json:"transID,omitempty"`
}

type SignInResult struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	Token *string `xml:"token,omitempty" json:"token,omitempty"`
}

type CheckTargetObject struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	ProviderCode *string `xml:"providerCode,omitempty" json:"providerCode,omitempty"`
}

type CategoryObject struct {
	CategoryId int32 `xml:"categoryId,omitempty" json:"categoryId,omitempty"`

	CategoryName *string `xml:"categoryName,omitempty" json:"categoryName,omitempty"`
}

type ArrayOfCategoryObject struct {
}

type CategoryListObject struct {
	CategoryList *ArrayOfCategoryObject `xml:"categoryList,omitempty" json:"categoryList,omitempty"`

	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type ProviderObject struct {
	ProviderCode *string `xml:"providerCode,omitempty" json:"providerCode,omitempty"`

	ProviderID int32 `xml:"providerID,omitempty" json:"providerID,omitempty"`

	ProviderName *string `xml:"providerName,omitempty" json:"providerName,omitempty"`
}

type ArrayOfProviderObject struct {
}

type ProviderListObject struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	ProviderList *ArrayOfProviderObject `xml:"providerList,omitempty" json:"providerList,omitempty"`
}

type ProductObject struct {
	ProductId int32 `xml:"productId,omitempty" json:"productId,omitempty"`

	ProductName *string `xml:"productName,omitempty" json:"productName,omitempty"`

	ProductValue int32 `xml:"productValue,omitempty" json:"productValue,omitempty"`
}

type ArrayOfProductObject struct {
}

type ProductListObject struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	ProductList *ArrayOfProductObject `xml:"productList,omitempty" json:"productList,omitempty"`
}

type ArrayOf_xsd_int struct {
}

type LevelResult struct {
	DataValue *ArrayOf_xsd_int `xml:"dataValue,omitempty" json:"dataValue,omitempty"`

	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type ProductInfo struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	ProductID int32 `xml:"productID,omitempty" json:"productID,omitempty"`

	ProductLogo *[]byte `xml:"productLogo,omitempty" json:"productLogo,omitempty"`

	ProductMessage *string `xml:"productMessage,omitempty" json:"productMessage,omitempty"`

	ProductName *string `xml:"productName,omitempty" json:"productName,omitempty"`

	ProductValue int32 `xml:"productValue,omitempty" json:"productValue,omitempty"`
}

type CheckProducts struct {
	Message *string `xml:"message,omitempty" json:"message,omitempty"`

	ProductId int32 `xml:"productId,omitempty" json:"productId,omitempty"`

	ProductName *string `xml:"productName,omitempty" json:"productName,omitempty"`

	Status int32 `xml:"status,omitempty" json:"status,omitempty"`
}

type ArrayOfCheckProducts struct {
}

type CheckProductsDetail struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`

	ProductId int32 `xml:"productId,omitempty" json:"productId,omitempty"`

	Quantity int32 `xml:"quantity,omitempty" json:"quantity,omitempty"`

	Status bool `xml:"status,omitempty" json:"status,omitempty"`
}

type ArrayOfCheckProductsDetail struct {
}

type ChangePasswordResult struct {
	ErrorCode int32 `xml:"errorCode,omitempty" json:"errorCode,omitempty"`

	ErrorMessage *string `xml:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}

type TopupInterface interface {
	SignOut(request *string) (*SignOutResult, error)
	SignOutContext(ctx context.Context, request *string) (*SignOutResult, error)
	CheckConnection() (*string, error)
	CheckConnectionContext(ctx context.Context) (*string, error)
	QueryBalance(request *string) (*BalanceResult, error)
	QueryBalanceContext(ctx context.Context, request *string) (*BalanceResult, error)
	SignInAsPartner(request *string) (*SignInResult, error)
	SignInAsPartnerContext(ctx context.Context, request *string) (*SignInResult, error)
	CheckPhoneNumber(request *string) (*bool, error)
	CheckPhoneNumberContext(ctx context.Context, request *string) (*bool, error)
	CheckTargetAccount(request *string) (*CheckTargetObject, error)
	CheckTargetAccountContext(ctx context.Context, request *string) (*CheckTargetObject, error)
	PartnerGetListCategories(request *string) (*CategoryListObject, error)
	PartnerGetListCategoriesContext(ctx context.Context, request *string) (*CategoryListObject, error)
	PartnerGetListProviderOfCategory(request *string) (*ProviderListObject, error)
	PartnerGetListProviderOfCategoryContext(ctx context.Context, request *string) (*ProviderListObject, error)
	PartnerGetListProductOfCategoryServiceProvider(request *string) (*ProductListObject, error)
	PartnerGetListProductOfCategoryServiceProviderContext(ctx context.Context, request *string) (*ProductListObject, error)
	GetListValueLevels(request *string) (*LevelResult, error)
	GetListValueLevelsContext(ctx context.Context, request *string) (*LevelResult, error)
	GetProductInformation(request *string) (*ProductInfo, error)
	GetProductInformationContext(ctx context.Context, request *string) (*ProductInfo, error)
	PartnerCheckProduct(request *string) (*ArrayOfCheckProducts, error)
	PartnerCheckProductContext(ctx context.Context, request *string) (*ArrayOfCheckProducts, error)
	PartnerCheckProductDetail(request *string) (*ArrayOfCheckProductsDetail, error)
	PartnerCheckProductDetailContext(ctx context.Context, request *string) (*ArrayOfCheckProductsDetail, error)
	PartnerChangePassword(request *string) (*ChangePasswordResult, error)
	PartnerChangePasswordContext(ctx context.Context, request *string) (*ChangePasswordResult, error)
	RequestHandle(request *string) (*string, error)
	RequestHandleContext(ctx context.Context, request *string) (*string, error)
}

type topupInterface struct {
	client *soap.Client
}

func NewTopupInterface(client *soap.Client) TopupInterface {
	return &topupInterface{
		client: client,
	}
}

func (service *topupInterface) SignOutContext(ctx context.Context, request *string) (*SignOutResult, error) {
	response := new(SignOutResult)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) SignOut(request *string) (*SignOutResult, error) {
	return service.SignOutContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) CheckConnectionContext(ctx context.Context) (*string, error) {
	response := new(string)
	err := service.client.CallContext(ctx, "''", nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) CheckConnection() (*string, error) {
	return service.CheckConnectionContext(
		context.Background(),
	)
}

func (service *topupInterface) QueryBalanceContext(ctx context.Context, request *string) (*BalanceResult, error) {
	response := new(BalanceResult)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) QueryBalance(request *string) (*BalanceResult, error) {
	return service.QueryBalanceContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) SignInAsPartnerContext(ctx context.Context, request *string) (*SignInResult, error) {
	response := new(SignInResult)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) SignInAsPartner(request *string) (*SignInResult, error) {
	return service.SignInAsPartnerContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) CheckPhoneNumberContext(ctx context.Context, request *string) (*bool, error) {
	response := new(bool)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) CheckPhoneNumber(request *string) (*bool, error) {
	return service.CheckPhoneNumberContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) CheckTargetAccountContext(ctx context.Context, request *string) (*CheckTargetObject, error) {
	response := new(CheckTargetObject)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) CheckTargetAccount(request *string) (*CheckTargetObject, error) {
	return service.CheckTargetAccountContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerGetListCategoriesContext(ctx context.Context, request *string) (*CategoryListObject, error) {
	response := new(CategoryListObject)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerGetListCategories(request *string) (*CategoryListObject, error) {
	return service.PartnerGetListCategoriesContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerGetListProviderOfCategoryContext(ctx context.Context, request *string) (*ProviderListObject, error) {
	response := new(ProviderListObject)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerGetListProviderOfCategory(request *string) (*ProviderListObject, error) {
	return service.PartnerGetListProviderOfCategoryContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerGetListProductOfCategoryServiceProviderContext(ctx context.Context, request *string) (*ProductListObject, error) {
	response := new(ProductListObject)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerGetListProductOfCategoryServiceProvider(request *string) (*ProductListObject, error) {
	return service.PartnerGetListProductOfCategoryServiceProviderContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) GetListValueLevelsContext(ctx context.Context, request *string) (*LevelResult, error) {
	response := new(LevelResult)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) GetListValueLevels(request *string) (*LevelResult, error) {
	return service.GetListValueLevelsContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) GetProductInformationContext(ctx context.Context, request *string) (*ProductInfo, error) {
	response := new(ProductInfo)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) GetProductInformation(request *string) (*ProductInfo, error) {
	return service.GetProductInformationContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerCheckProductContext(ctx context.Context, request *string) (*ArrayOfCheckProducts, error) {
	response := new(ArrayOfCheckProducts)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerCheckProduct(request *string) (*ArrayOfCheckProducts, error) {
	return service.PartnerCheckProductContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerCheckProductDetailContext(ctx context.Context, request *string) (*ArrayOfCheckProductsDetail, error) {
	response := new(ArrayOfCheckProductsDetail)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerCheckProductDetail(request *string) (*ArrayOfCheckProductsDetail, error) {
	return service.PartnerCheckProductDetailContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) PartnerChangePasswordContext(ctx context.Context, request *string) (*ChangePasswordResult, error) {
	response := new(ChangePasswordResult)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) PartnerChangePassword(request *string) (*ChangePasswordResult, error) {
	return service.PartnerChangePasswordContext(
		context.Background(),
		request,
	)
}

func (service *topupInterface) RequestHandleContext(ctx context.Context, request *string) (*string, error) {
	response := new(string)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *topupInterface) RequestHandle(request *string) (*string, error) {
	return service.RequestHandleContext(
		context.Background(),
		request,
	)
}
