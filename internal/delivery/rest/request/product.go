package request

type Product struct {
	Name          string   `json:"name" validate:"required,string,min=5,max=60"`
	Price         int      `json:"price" validate:"required,numeric,gte=0"`
	ImageUrl      string   `json:"imageUrl" validate:"required,url"`
	Stock         int      `json:"stock" validate:"required,numeric,gte=0"`
	Condition     string   `json:"condition" validate:"required,oneof=new second"`
	Tags          []string `json:"tags" validate:"required,min=0"`
	IsPurchasable bool     `json:"isPurchasable" validate:"required,boolean"`
}

// {
// 	"name": "", // not null, minLength 5, maxLength 60
// 	"price": 10000, // not null, min 0
// 	"imageUrl" : "", // not null, url=true
// 	"stock" : 10, // not null, min 0
// 	"condition": "new | second", // not null, must only accept enum
// 	"tags": [""], // not null, minItems 0
// 	"isPurchasable": true // not null
// }
