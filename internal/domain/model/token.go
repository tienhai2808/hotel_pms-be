package model

import "time"

type Token struct {
	ID        int64      `gorm:"type:bigint;primaryKey" json:"id"`
	UserID    int64      `gorm:"type:bigint;not null;index:tokens_user_id_user_agent_expires_at_idx,priority:1" json:"user_id"`
	Token     string     `gorm:"type:varchar(255);not null;uniqueIndex:tokens_token_key" json:"token"`
	UserAgent string     `gorm:"type:varchar(255);not null;index:tokens_user_id_user_agent_expires_at_idx,priority:2" json:"user_agent"`
	RevokedAt *time.Time `json:"revoked_at"`
	ExpiresAt time.Time  `gorm:"not null;index:tokens_user_id_user_agent_expires_at_idx,priority:3" json:"expires_at"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:fk_tokens_user,OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
