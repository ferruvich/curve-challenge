package middleware

import (
	"context"

	"github.com/ferruvich/curve-prepaid-card/internal/model"
	"github.com/ferruvich/curve-prepaid-card/internal/repo"
)

// Card represents the card middleware interface
type Card interface {
	Create(context.Context, string) (*model.Card, error)
	GetCard(context.Context, string) (*model.Card, error)
	Deposit(context.Context, string, float64) error
}

// CardMiddleware is the Card implementation
type CardMiddleware struct {
	repo repo.Card
}

// NewCardMiddleware returns a new middleware for user
func NewCardMiddleware(ctx context.Context) (Card, error) {
	repo, err := repo.NewCardRepo(ctx)
	if err != nil {
		return nil, err
	}

	return &CardMiddleware{
		repo: repo,
	}, nil
}

// Create creates and returns new merchants, or an error if there is any
func (m *CardMiddleware) Create(ctx context.Context, ownerID string) (*model.Card, error) {

	card, err := model.NewCard(ownerID)
	if err != nil {
		return nil, err
	}

	if err = m.repo.Write(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}

// GetCard gets and returns an existing card, or an error if there is any
func (m *CardMiddleware) GetCard(ctx context.Context, cardID string) (*model.Card, error) {

	card, err := m.repo.Read(ctx, cardID)
	if err != nil {
		return nil, err
	}

	return card, nil
}

// Deposit topups a card, adding some money
func (m *CardMiddleware) Deposit(ctx context.Context, cardID string, amount float64) error {

	card, err := m.repo.Read(ctx, cardID)
	if err != nil {
		return err
	}

	card.IncrementAccountBalance(amount)

	if err = m.repo.Update(ctx, card); err != nil {
		return err
	}

	return nil
}
