package service

import (
	"assignment_bd/global"
	"assignment_bd/model"
)

func PublishVideo(video *model.Video) error {
	err := global.DB.Create(video).Error
	if err != nil {
		return err
	}
	return nil
}
