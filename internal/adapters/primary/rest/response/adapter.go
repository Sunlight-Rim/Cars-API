package response

import authEntity "cars/internal/entities/auth"

func NewSignup(r *authEntity.UsecaseResSignup) *Signup {
	return &Signup{ID: r.ID}
}
