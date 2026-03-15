package middleware

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
)

type MiddlewareRepository interface {
	UserHasPermission(userID int64, permission string) (bool, error)
}

type MiddlewareRepositoryImpl struct {
	store database.DBStore
}

func NewMiddlewareRepository(store database.DBStore) MiddlewareRepository {
	return &MiddlewareRepositoryImpl{
		store: store,
	}
}

func (r *MiddlewareRepositoryImpl) UserHasPermission(userID int64, permission string) (bool, error) {
	builder := sq.Select("1").
		From("user_roles ur").
		Join("roles r ON r.id = ur.role_id").
		Join("role_permissions rp ON rp.role_id = r.id").
		Join("permissions p ON p.id = rp.permission_id").
		Where(sq.Eq{"ur.user_id": userID}).
		Where(sq.Eq{"p.code": permission}).
		PlaceholderFormat(sq.Question).
		Limit(1)
	
	var exists int
	err := r.store.GetClient().Query(r.store.GetContext(), builder, &exists)
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}