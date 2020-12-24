package resolvers

import (
	"products/data/database/controllers"

	"github.com/graphql-go/graphql"
)

// ProductResolver type.
type ProductResolver struct {
	DB *controllers.DBProductInstance
}

// IndexProducts function returns a list of products that the user_id
// corresponds with the provided user_id in arguments.
func (resolver *ProductResolver) IndexProducts(params graphql.ResolveParams) (interface{}, error) {
	products, err := resolver.DB.IndexProducts(params.Args["user_id"].(string))

	if err != nil {
		return nil, err
	}

	return products, nil
}

// AddProduct function adds a product with the provided data and return
// the product row registered in the database.
func (resolver *ProductResolver) AddProduct(params graphql.ResolveParams) (interface{}, error) {
	product, err := resolver.DB.AddProduct(params.Args)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct updates the product data that id corresponds with the
// provided id in arguments.
func (resolver *ProductResolver) UpdateProduct(params graphql.ResolveParams) (interface{}, error) {
	product, err := resolver.DB.UpdateProduct(params.Args)

	if err != nil {
		return nil, err
	}

	return product, err
}

// DeleteProduct function deletes a product in database by id provided
// in arguments.
func (resolver *ProductResolver) DeleteProduct(params graphql.ResolveParams) (interface{}, error) {
	result, err := resolver.DB.DeleteProduct(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	return result, nil
}
