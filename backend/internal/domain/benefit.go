package domain

// Benefit представляет льготу
type Benefit struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Discount    int    `json:"discount"` // процент скидки
	IsActive    bool   `json:"is_active"`
}

// BenefitCategory представляет категории льгот
type BenefitCategory string

const (
	BenefitCategoryLargeFamily BenefitCategory = "large_family"      // 30%
	BenefitCategoryDisabled    BenefitCategory = "disabled"          // 30%
	BenefitCategoryRepressed   BenefitCategory = "repressed"         // 50%
	BenefitCategoryAfghan      BenefitCategory = "afghan_veteran"     // 100%
	BenefitCategoryWWII        BenefitCategory = "wwii_veteran"      // 100%
)

