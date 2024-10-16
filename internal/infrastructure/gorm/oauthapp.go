package gorm

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type OAuthApp struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Provider     string `gorm:"index"`
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresAt    time.Time
	Scopes       pq.StringArray `gorm:"type:text[]"`
	OwnerID      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (OAuthApp) TableName() string {
	return "oauth_apps"
}

func (dao OAuthApp) ToDomain() *oauthapp.OAuthApp {

	return &oauthapp.OAuthApp{
		ID:           dao.ID,
		Provider:     dao.Provider,
		AccessToken:  dao.AccessToken,
		RefreshToken: dao.RefreshToken,
		ExpiresAt:    dao.ExpiresAt,
		TokenType:    dao.TokenType,
		Scopes:       dao.Scopes,
		OwnerID:      dao.OwnerID,
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
	}
}

func newOAuthAppFromDomain(app *oauthapp.OAuthApp) *OAuthApp {
	return &OAuthApp{
		ID:           app.ID,
		Provider:     app.Provider,
		AccessToken:  app.AccessToken,
		RefreshToken: app.RefreshToken,
		ExpiresAt:    app.ExpiresAt,
		Scopes:       app.Scopes,
		TokenType:    app.TokenType,
		OwnerID:      app.OwnerID,
		CreatedAt:    app.CreatedAt,
		UpdatedAt:    app.UpdatedAt,
	}
}

type OAuthAppRepository struct {
	db *gorm.DB
}

func NewOAuthAppRepository(db *gorm.DB) *OAuthAppRepository {
	return &OAuthAppRepository{
		db: db,
	}
}

func (r OAuthAppRepository) Find(ctx context.Context, ownerID string, provider string) (*oauthapp.OAuthApp, error) {

	var app OAuthApp
	if err := r.db.WithContext(ctx).Where("provider = ? AND owner_id = ?", provider, ownerID).Take(&app).Error; err != nil {
		return nil, err
	}

	return app.ToDomain(), nil
}

func (r OAuthAppRepository) ListForOwner(ctx context.Context, ownerID string) ([]*oauthapp.OAuthApp, error) {

	var apps []*OAuthApp

	if err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&apps).Error; err != nil {
		return nil, err
	}

	var result []*oauthapp.OAuthApp
	for _, app := range apps {
		result = append(result, app.ToDomain())
	}

	return result, nil
}

func (r OAuthAppRepository) ListForProvider(ctx context.Context, providerID string) ([]*oauthapp.OAuthApp, error) {

	var apps []*OAuthApp

	if err := r.db.WithContext(ctx).Where("provider = ?", providerID).Find(&apps).Error; err != nil {
		return nil, err
	}

	var result []*oauthapp.OAuthApp
	for _, app := range apps {
		result = append(result, app.ToDomain())
	}

	return result, nil
}

func (r OAuthAppRepository) Create(ctx context.Context, app *oauthapp.OAuthApp) error {
	if err := r.db.WithContext(ctx).Create(newOAuthAppFromDomain(app)).Error; err != nil {
		return err
	}

	return nil
}

func (r OAuthAppRepository) UpdateByID(ctx context.Context, id uint, updateFn func(app *oauthapp.OAuthApp) error) error {
	var app OAuthApp
	if err := r.db.WithContext(ctx).Find(&app, "id = ?", id).Error; err != nil {
		return err
	}

	domainObject := app.ToDomain()
	if err := updateFn(domainObject); err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Save(newOAuthAppFromDomain(app.ToDomain())).Error; err != nil {
		return err
	}

	return nil
}

func (r OAuthAppRepository) Delete(ctx context.Context, id uint) error {
	q := r.db.WithContext(ctx).Where("id = ?", id).Delete(&OAuthApp{})

	if q.Error != nil {
		return fmt.Errorf("failed to delete provider: %w", q.Error)
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
