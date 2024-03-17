package response

type Product struct {
	Id            string   `json:"productId"`
	Name          string   `json:"name"`
	Price         int      `json:"price"`
	ImageUrl      string   `json:"imageUrl"`
	Stock         int      `json:"stock"`
	Condition     string   `json:"condition"`
	Tags          []string `json:"tags"`
	OwnerId       string   `json:"owner_id"`
	IsPurchasable bool     `json:"ispurchaseable"`
	PurchaseCount int      `json:"purchaseCount"`
}
