package validators

type RegisterResponseModel struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}

type LoginResponseModel struct {
	Token     string `json:"token"`
	User      User   `json:"user"`
	ExpiresIn int64  `json:"expires_in"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type DeactivationResponseModel struct {
	Message string `json:"message"`
}

type DeleteResponseModel struct {
	Message string `json:"message"`
}

type InventoryResponse struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type InventoryListResponse struct {
	Items []InventoryResponse `json:"items"`
}

type CartResponse struct {
	ISBN     string `json:"isbn"`
	Quantity int    `json:"quantity"`
}

type ReviewResponse struct {
	ISBN   string `json:"isbn"`
	Review string `json:"review"`
}

type CartItemResponse struct {
	ID       uint   `json:"id"`
	ISBN     string `json:"isbn"`
	Quantity int    `json:"quantity"`
}

type PurchaseResponse struct {
	Message string `json:"message"`
	ISBN    string `json:"isbn"`
}
