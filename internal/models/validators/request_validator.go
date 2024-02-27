package validators

type RegisterRequestModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRequestModel struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type BookAddUpdateRequestModel struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Content     []byte `json:"content"`
}

//type BookUpdateRequestModel struct {
//	ISBN        string `json:"isbn" gorm:"primaryKey"`
//	Title       string `json:"title" gorm:"optional"`
//	Author      string `json:"author"`
//	Description string `json:"description"`
//}

// CartAddUpdateRequestModel used for validating incoming request
type CartAddUpdateRequestModel struct {
	ISBN     string `json:"isbn"`
	Quantity int    `json:"quantity"`
}

type ReviewAddRequest struct {
	ISBN   string `json:"isbn"`
	Review string `json:"review"`
}

type PurchaseRequest struct {
	ISBN string `json:"isbn"`
}
type AdminRequest struct {
	AdminCode string `json:"admin_code"`
}
