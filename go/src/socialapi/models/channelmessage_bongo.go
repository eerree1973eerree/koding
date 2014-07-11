package models

import (

	"github.com/koding/bongo"
)

func (c *ChannelMessage) BeforeCreate() error {
	c.DeletedAt = ZeroDate()

	return c.MarkIfExempt()
}

func (c *ChannelMessage) BeforeUpdate() error {
	c.DeletedAt = ZeroDate()

	return c.MarkIfExempt()
}

func (c *ChannelMessage) AfterCreate() {
	bongo.B.AfterCreate(c)
}

func (c *ChannelMessage) AfterUpdate() {
	bongo.B.AfterUpdate(c)
}

func (c ChannelMessage) AfterDelete() {
	bongo.B.AfterDelete(c)
}

func (c ChannelMessage) GetId() int64 {
	return c.Id
}

func (c ChannelMessage) TableName() string {
	return "api.channel_message"
}

func NewChannelMessage() *ChannelMessage {
	return &ChannelMessage{}
}

func (c *ChannelMessage) Update() error {
	if err := bodyLenCheck(c.Body); err != nil {
		return err
	}
	// only update body
	err := bongo.B.UpdatePartial(c,
		map[string]interface{}{
			"body": c.Body,
		},
	)
	return err
}

func (c *ChannelMessage) Create() error {
	if err := bodyLenCheck(c.Body); err != nil {
		return err
	}

	var err error
	c, err = Slugify(c)
	if err != nil {
		return err
	}

	return bongo.B.Create(c)
}

func (c *ChannelMessage) ById(id int64) error {
	return bongo.B.ById(c, id)
}

func (c *ChannelMessage) One(q *bongo.Query) error {
	return bongo.B.One(c, c, q)
}

func (c *ChannelMessage) Some(data interface{}, q *bongo.Query) error {
	return bongo.B.Some(c, data, q)
}

func (c *ChannelMessage) UpdateMulti(rest ...map[string]interface{}) error {
	return bongo.B.UpdateMulti(c, rest...)
}

func (c *ChannelMessage) CountWithQuery(q *bongo.Query) (int, error) {
	return bongo.B.CountWithQuery(c, q)
}

func (c *ChannelMessage) Delete() error {
	return bongo.B.Delete(c)
}
