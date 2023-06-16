package api

import (
	"github.com/gin-gonic/gin"
	"gorm_gen/dao"
	"gorm_gen/model"
)

type IMessage interface {
	CreateMessageApi(ctx *gin.Context)
	GetMessageListApi(ctx *gin.Context)
}

type Message struct {
	db *dao.Query
}

func (m Message) CreateMessageApi(ctx *gin.Context) {
	if result := m.db.Message.Create(&model.Message{
		UserID:   1,
		Category: 1,
		TargetID: 1,     //int64          `gorm:"column:target_id;type:int;not null;comment:如果是好友就关联到user表的id,如果是群聊就关联到community表id" json:"target_id"`
		Media:    1,     // int64          `gorm:"column:media;type:tinyint;not null;comment:消息按照什么样式展示" json:"media"`
		Content:  "哈哈哈", // string         `gorm:"column:content;type:varchar(255);comment:消息内容体" json:"content"`
		Pic:      "11",  //  string         `gorm:"column:pic;type:varchar(200);comment:图片预览地址" json:"pic"`
		URL:      "11",  // string         `gorm:"column:url;type:varchar(255);comment:服务的url" json:"url"`
		Remark:   "233", //string         `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	}); result != nil {

	}
	ctx.JSON(200, "成功")
}

func (m Message) GetMessageListApi(ctx *gin.Context) {
	dao := m.db.Message
	result, _ := dao.Find()
	ctx.JSON(200, result)
}

func NewMessage(db *dao.Query) IMessage {
	return Message{
		db: db,
	}
}
