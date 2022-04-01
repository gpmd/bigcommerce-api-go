package bigcommerce

import (
	"net/url"
)

// StoreClient interface handles generic store requests
type StoreClient interface {
	GetAllChannels() ([]Channel, error)
	GetChannels(page int) ([]Channel, bool, error)
	GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
	GetStoreInfo() (StoreInfo, error)
}

// CatalogClient interface handles catalog-related requests
type CatalogClient interface {
	GetAllBrands() ([]Brand, error)
	GetBrands(page int) ([]Brand, bool, error)
	GetAllCategories() ([]Category, error)
	GetCategories(page int) ([]Category, bool, error)
	GetClientRequest(requestURLQuery url.Values) (*ClientRequest, error)
	GetMainThumbnailURL(productID int64) (string, error)
	SetProductFields(fields []string)
	SetProductInclude(subresources []string)
	GetAllProducts() ([]Product, error)
	GetProducts(page int) ([]Product, bool, error)
	GetProductByID(productID int64) (*Product, error)
}

// BlogClient interface handles blog-related requests
type BlogClient interface {
	GetAllPosts(context, xAuthToken string) ([]Post, error)
	GetPosts(page int) ([]Post, bool, error)
}

// CartClient interface handles cart and login related requests
type CartClient interface {
	CreateCart(items []LineItem) (*Cart, error)
	GetCart(cartID string) (*Cart, error)
	CartAddItems(cartID string, items []LineItem) (*Cart, error)
	CartEditItem(cartID string, item LineItem) (*Cart, error)
	CartDeleteItem(cartID string, item LineItem) (*Cart, error)
	CartUpdateCustomerID(cartID, customerID string) (*Cart, error)
	DeleteCart(cartID string) error
}

type CustomerClient interface {
	ValidateCredentials(email, password string) (int64, error)
	CreateAccount(customer *CreateAccountPayload) (*Customer, error)
	CustomerSetFormFields(customerID int64, formFields []FormField) error
	CustomerGetFormFields(customerID int64) ([]FormField, error)
	GetCustomerByID(customerID int64) (*Customer, error)
	GetCustomerByEmail(email string) (*Customer, error)
}

type AddressClient interface {
	CreateAddress(customerID int64, address *Address) (*Address, error)
	UpdateAddress(customerID int64, address *Address) (*Address, error)
	DeleteAddress(customerID int64, addressID int64) error
	GetAddresses(customerID int64) ([]Address, error)
}
