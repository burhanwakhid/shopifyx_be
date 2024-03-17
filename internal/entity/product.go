package entity

type Product struct {
	Id            string   `db:"id"`
	Name          string   `db:"name"`
	Price         int      `db:"price"`
	ImageUrl      string   `db:"imageUrl"`
	Stock         int      `db:"stock"`
	Condition     string   `db:"condition"`
	Tags          []string `db:"tags"`
	OwnerId       string   `db:"owner_id"`
	IsPurchasable bool     `db:"ispurchaseable"`
	PurchaseCount int      `db:"purchaseCount"`
	SellerName    string   `db:"seller_name"`
}

type Seller struct {
	Name             string `db:"name"`
	ProductSoldTotal string
	Banks            *[]Bank
}

type ProductDetail struct {
	Product *Product
	Seller  *Seller
}
