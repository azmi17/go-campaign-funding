package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

// Untuk membuat instance dari Struct Repository agar bisa diakses dari luar package
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {

	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {

	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error // Preload untuk me-load sautu relasi dari campaign => campaignImages
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
