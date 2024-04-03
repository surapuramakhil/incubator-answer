package migrations

import (
	"context"
	"time"

	"xorm.io/xorm"
)

func addColumnObjectTypeInTagReltable(ctx context.Context, x *xorm.Engine) error {

	type TagRel struct {
		ID        int64     `xorm:"not null pk autoincr BIGINT(20) id"`
		CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
		UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
		ObjectID  string    `xorm:"not null INDEX UNIQUE(s) BIGINT(20) object_id"`
		ObjectType int       `xorm:"not null INDEX UNIQUE(s) INT(11) object_type"`
		TagID  string `xorm:"not null INDEX UNIQUE(s) BIGINT(20) tag_id"`
		Status int    `xorm:"not null default 1 INT(11) status"`
	}

	// set object type to 1 for existing records for all sql dialects
	err := x.Sync(new(TagRel))

	if err != nil {
		return err
	}

	_, err = x.Exec("UPDATE tag_rel SET object_type = 1")
	
	return err
}
