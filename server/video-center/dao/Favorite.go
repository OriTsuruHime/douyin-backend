package dao

import (
	"context"
	"gorm.io/gorm"
)

type Fav struct {
	UserId  uint
	VideoId int64
	Video   Video `gorm:"foreignKey:video_id"`
}

func (u *Fav) TableName() string {
	return "FavRelation"
}

// CreateFav 点赞
func CreateFav(ctx context.Context, videoId int64, userId int64) error {
	tx := DB.WithContext(ctx).Begin() // 开启事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	fav := Fav{
		UserId:  uint(userId),
		VideoId: videoId,
	}
	err := tx.Model(&Fav{}).Create(&fav).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// DeleteFav 取消点赞
func DeleteFav(ctx context.Context, videoId int64, userId int64) error {
	tx := DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Model(&Fav{}).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&Fav{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// IsFavorite 判断是否点赞
func IsFavorite(ctx context.Context, videoId int64, userId int64) (bool, error) {
	var fav Fav
	result := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).First(&fav)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

// GetFavoriteCount 获取用户点赞数量
func GetFavoriteCount(ctx context.Context, userId int64) (int32, error) {
	var favCount int64
	result := DB.WithContext(ctx).Model(&Fav{}).Where("user_id = ?", userId).Count(&favCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return int32(favCount), nil
}

// GetReceivedFavoriteCount 获取用户收到的点赞数量
func GetReceivedFavoriteCount(ctx context.Context, userId int64) (int32, error) {
	var video []Video
	result := DB.Table("videos").Where("author_id = ?", userId).Select("favorite_count").Find(&video)
	if result.Error != nil {
		return 0, result.Error
	}
	var favCount int64
	for _, v := range video {
		favCount += v.FavoriteCount
	}
	return int32(favCount), nil
}

// GetSingleVideoFavoriteCount 获取单个视频点赞数量
func GetSingleVideoFavoriteCount(ctx context.Context, videoId int64) (int32, error) {
	var favCount int32
	result := DB.Table("videos").Where("video_id = ?", videoId).Select("favorite_count").Scan(&favCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return favCount, nil
}

// ListFav 获取用户喜欢列表
func ListFav(ctx context.Context, userId int64) []Video {
	var favs []Fav
	DB.WithContext(ctx).Where("user_id = ? ", userId).Find(&favs)
	var videoIDs []int64
	for _, rel := range favs {
		videoIDs = append(videoIDs, rel.VideoId)
	}
	if len(videoIDs) == 0 {
		return []Video{}
	}
	var videos []Video
	DB.WithContext(ctx).Find(&videos, videoIDs)
	return videos
}
