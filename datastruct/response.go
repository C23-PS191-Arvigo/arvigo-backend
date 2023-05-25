package datastruct

import "time"

type (
	LoginRegisterResponse struct {
		UserID uint64 `json:"user_id"`
		Token  string `json:"token"`
	}

	UserDetailResponse struct {
		UserDetail
		Address UserAddress `json:"address"`
	}

	FaceShapeResponse struct {
		ImageUrl string `json:"image_url"`
		Result   string `json:"result"`
	}

	UserDetail struct {
		ID                        uint64     `json:"id"`
		Email                     string     `json:"email"`
		RoleID                    uint64     `json:"role_id"`
		RoleName                  string     `json:"role_name"`
		FullName                  string     `json:"full_name"`
		Gender                    string     `json:"gender"`
		DateOfBirth               *time.Time `json:"date_of_birth"`
		PlaceOfBirth              string     `json:"place_of_birth"`
		IsCompletePersonalityTest bool       `json:"is_complete_personality_test"`
		IsCompleteFaceTest        bool       `json:"is_complete_face_test"`
		PersonalityID             bool       `json:"personality_id"`
		FaceShapeTagID            bool       `json:"face_shape_tag_id"`
		IsVerified                bool       `json:"is_verified"`
		Avatar                    string     `json:"avatar"`
		AddressID                 uint64     `json:"addresses_id"`
		MerchantID                uint64     `json:"merchant_id"`
	}

	UserAddress struct {
		Street      string `json:"street"`
		Province    string `json:"province"`
		City        string `json:"city"`
		District    string `json:"district"`
		SubDistrict string `json:"sub_district"`
		PostalCode  uint64 `json:"postal_code"`
	}

	InitialProductResponse struct {
		InitialProduct
		Images   []string                `json:"images"`
		Variants []InitialProductVariant `json:"variants"`
	}

	InitialProduct struct {
		ID           uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Name         string `gorm:"column:name" json:"name"`
		Description  string `gorm:"column:description" json:"description"`
		LinkExternal string `gorm:"column:link_external" json:"link_external"`
		CategoryName string `gorm:"column:category_name" json:"category_name"`
		BrandName    string `gorm:"column:brand_name" json:"brand_name"`
		Images       string `gorm:"column:images" json:"-"`
	}

	InitialProductVariant struct {
		Name             string `gorm:"column:name" json:"name"`
		LinkAR           string `gorm:"column:link_ar" json:"link_ar"`
		IsPrimaryVariant bool   `gorm:"column:is_primary_variant" json:"is_primary_variant"`
		ProductID        uint64 `gorm:"column:product_id" json:"-"`
	}

	BrandResponse struct {
		Name  string `gorm:"column:name" json:"name"`
		Image string `gorm:"column:image" json:"image"`
	}

	CategoryResponse struct {
		Name string `gorm:"column:name" json:"name"`
	}

	ProductRecommendationResponse struct {
		ID          uint64 `gorm:"column:id" json:"id"`
		Name        string `gorm:"column:name" json:"name"`
		Description string `gorm:"column:description" json:"description"`
		Category    string `gorm:"column:category" json:"category"`
		Brand       string `gorm:"column:brand" json:"brand"`
		Tags        string `gorm:"column:tags" json:"tags"`
		Merchants   string `gorm:"column:merchants" json:"merchants"`
		Clicked     uint64 `gorm:"column:clicked" json:"clicked"`
		MerchantIDs string `gorm:"column:merchant_id" json:"-"`
		ProductIDs  string `gorm:"column:linked_product" json:"-"`
	}
)
