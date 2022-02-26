package tweet

import "gorm.io/gorm"

type Repository interface {
	Save(tweet Tweet) (Tweet, error)
	Update(tweet Tweet) (Tweet, error)
	GetById(id uint) (Tweet, error)
	GetByUserId(userId uint) ([]Tweet, error)
	Get() ([]Tweet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(tweet Tweet) (Tweet, error) {
	if err := r.db.Create(&tweet).Error; err != nil {
		return tweet, err
	}

	return tweet, nil
}

func (r *repository) Update(tweet Tweet) (Tweet, error) {
	if err := r.db.Save(&tweet).Error; err != nil {
		return tweet, err
	}

	return tweet, nil
}

func (r *repository) GetById(id uint) (Tweet, error) {
	var tweet Tweet
	if err := r.db.Where("id = ?", id).Find(&tweet).Error; err != nil {
		return tweet, err
	}

	return tweet, nil
}

func (r *repository) GetByUserId(userId uint) ([]Tweet, error) {
	var tweets []Tweet
	if err := r.db.Preload("User").Where("user_id = ?", userId).Find(&tweets).Error; err != nil {
		return tweets, err
	}

	return tweets, nil
}

func (r *repository) Get() ([]Tweet, error) {
	var tweets []Tweet
	if err := r.db.Preload("User").Find(&tweets).Error; err != nil {
		return tweets, err
	}

	return tweets, nil
}
