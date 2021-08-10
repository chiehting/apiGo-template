package models

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// User table struct
type User struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	Email             string    `gorm:"unique" json:"email"`
	EncryptedPassword []byte    `json:"encryptedPassword"`
	SignInCount       uint8     `json:"signInCount"`
	SignInAt          time.Time `json:"signInAt"`
	SignInIP          string    `json:"signInIp"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	Name              string    `gorm:"index" json:"name"`
	LockedAt          time.Time `gorm:"default:0000-00-00 00:00:00" json:"lockedAt"`
	Username          string    `gorm:"unique" json:"username"`
	State             int       `json:"state"`
	CreatedByID       int       `json:"createdById"`
	Location          string    `json:"location"`
	Note              string    `json:"note"`
	External          uint8     `json:"external"`
	Organization      string    `json:"organization"`
	PreferredLanguage string    `json:"preferredLanguage"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	UserType          uint8     `json:"userType"`
	login             `gorm:"-"`
}

type login struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

var _ = User{}

// TableName sets the insert table name for this struct type
func (user *User) TableName() string {
	return "Users"
}

// Create is insert data to users table
func (user *User) Create() error {
	return _db.Create(user).Error
}

// Get is select login user
func (user *User) Get() error {
	var result = _db.First(&user, User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		State:    user.State,
	})
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	return result.Error
}

// Updates is update data when login, like sign_in_ip or sign_in_count
func (user *User) Updates() error {
	return _db.Model(&user).Updates(user).Error
}

// SetSignIn is update data when login, like sign_in_ip or sign_in_count
func (user *User) SetSignIn() error {
	return _db.Model(&user).Select("sign_in_at", "sign_in_ip", "sign_in_count").Updates(user).Error
}

// GetCache is get user info from cache
func (user *User) GetCache() map[string]string {
	return _cache.HGetAll(_ctx, user.getCacheKey(0)).Val()
}

// DelCache is to remove user cache
func (user *User) DelCache() error {
	return _cache.Del(_ctx, user.getCacheKey(0)).Err()
}

// SetExpireTime is set cache expire time
func (user *User) SetExpireTime(time time.Duration) error {
	return _cache.Expire(_ctx, user.getCacheKey(0), time).Err()
}

// SetCache is set data to cache
func (user *User) SetCache() error {
	var cacheKey = user.getCacheKey(0)
	var data = map[string]interface{}{
		"Token":             user.Token,
		"Email":             user.Email,
		"Username":          user.Username,
		"PreferredLanguage": user.PreferredLanguage,
		"FirstName":         user.FirstName,
		"LastName":          user.LastName,
		"SignInIP":          user.SignInIP,
	}
	return _cache.HSet(_ctx, cacheKey, data).Err()
}

func (user *User) getCacheKey(cacheKey int) string {
	return _cachePrefix + [...]string{
		"user-" + strconv.Itoa(user.ID),
	}[cacheKey]
}
