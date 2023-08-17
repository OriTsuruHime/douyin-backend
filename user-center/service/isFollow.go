package service

import (
	"context"
	"fmt"
	"user-center/cache"
	"user-center/dao"
)

type IsFollowService struct {
	UserId       uint `json:"user_id"  form:"user_id"`
	FollowUserId uint `json:"follow_user_id" form:"follow_user_id"`
}

func (service *IsFollowService) IsFollow(ctx context.Context) (bool, error) { //todo: 添加返回结构体
	userDao := dao.NewUserDao(ctx)
	userCache := cache.NewRedisCache(ctx)
	if userCache.IsFollow(ctx, service.UserId, service.FollowUserId) == true { //查找缓存中是否存在
		fmt.Println("缓存中存在该关系")
		return true, nil
	}

	exist, err := userDao.IsFollowLogic(service.UserId, service.FollowUserId)
	if err != nil {
		return false, err
	}
	if exist == true {
		//todo: 添加缓存记录状态
		err = userCache.AddFollow(ctx, service.UserId, service.FollowUserId)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
