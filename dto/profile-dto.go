package dto

type ProfileUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Pic1        string `json:"pic1" form:"pic1" binding:"required"`
	Pic2        string `json:"pic2" form:"pic2" binding:"required"`
	Pic3        string `json:"pic3" form:"pic3" binding:"required"`
	Pic4        string `json:"pic4" form:"pic4" binding:"required"`
	Pic5        string `json:"pic5" form:"pic5" binding:"required"`
	Pic6        string `json:"pic6" form:"pic6" binding:"required"`
	Pic7        string `json:"pic7" form:"pic7" binding:"required"`
	Fields      string `json:"fields" form:"fields" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type ProfileCreateDTO struct {
	Description string `json:"description" form:"description" binding:"required"`
	Pic1        string `json:"pic1" form:"pic1" binding:"required"`
	Pic2        string `json:"pic2" form:"pic2" binding:"required"`
	Pic3        string `json:"pic3" form:"pic3" binding:"required"`
	Pic4        string `json:"pic4" form:"pic4" binding:"required"`
	Pic5        string `json:"pic5" form:"pic5" binding:"required"`
	Pic6        string `json:"pic6" form:"pic6" binding:"required"`
	Pic7        string `json:"pic7" form:"pic7" binding:"required"`
	Fields      string `json:"fields" form:"fields" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
