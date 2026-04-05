package repository

import (
	"context"
	"sync"
	"time"

	"github.com/SiwakornSitti/practice-ai/app/internal/user/domain"
)

type memoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User // map[id]*User
}

// NewMemoryUserRepository creates a new in-memory user repository
func NewMemoryUserRepository() domain.UserRepository {
	return &memoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *memoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if email already exists
	for _, u := range r.users {
		if u.Email == user.Email {
			return domain.ErrEmailAlreadyInUse
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *memoryUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (r *memoryUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (r *memoryUserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user
	return nil
}

func (r *memoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return domain.ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}
