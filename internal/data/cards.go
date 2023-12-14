package data

import "time"

type Card struct {
	ID        int64      `json:"id"` // Unique integer id for the company
	ProductID int64      `json:"productId"`
	SetID     int        `json:"setId"`
	SetCode   string     `json:"setCode"`
	Name      string     `json:"productUrlName"` // name of card
	URL       string     `json:"url"`            // URL location where resource is located
	Version   int32      `json:"version"`        // updated each time a record is updated
	CreatedAt *time.Time `json:"-"`              // created timestamp for the data

	PricePoints Pricepoints `json:"pricePoints"`
	Latest      LatestSales `json:"latestSales"`
}

type Pricepoints []struct {
	PrintingType       string  `json:"printingType"`
	MarketPrice        float64 `json:"marketPrice"`
	BuylistMarketPrice float64 `json:"buylistMarketPrice"`
	ListedMedianPrice  float64 `json:"listedMedianPrice"`
}

type LatestSales struct {
	ResultCount  int `json:"resultCount"`
	TotalResults int `json:"totalResults"`
	Data         []struct {
		Condition       string    `json:"condition"`
		Variant         string    `json:"variant"`
		Language        string    `json:"language"`
		Quantity        int       `json:"quantity"`
		Title           string    `json:"title"`
		ListingType     string    `json:"listingType"`
		CustomListingID string    `json:"customListingId"`
		PurchasePrice   float64   `json:"purchasePrice"`
		ShippingPrice   float64   `json:"shippingPrice"`
		OrderDate       time.Time `json:"orderDate"`
	} `json:"data"`
}
