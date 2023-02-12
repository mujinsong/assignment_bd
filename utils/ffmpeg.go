package utils

import (
	"bytes"
	"fmt"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
获取视频封面图片路径
*/
func GetVideoName(videoPath string) string {
	names := filepath.Base(filepath.ToSlash(videoPath))
	names = strings.Split(names, ".")[0]
	return path.Join("static", "image", names+".jpg")
}

/*
截取视频第一帧为视频封面 返回图片路径
*/
func GetSnapshot(videoPath string) string {
	buf := bytes.NewBuffer(nil)
	videoPath = filepath.ToSlash("./" + videoPath)
	ffmpeg.Input(videoPath, ffmpeg.KwArgs{
		"loglevel": "quiet",
	}).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", int(1))}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	coverdata, err := imaging.Decode(buf)
	covername := "./" + GetVideoName(videoPath)
	err = imaging.Save(coverdata, covername)
	if err != nil {
		//println("failed to save image:", err)
	}
	return covername[2:]
}

/*
生成随机视频名
*/
func RandVideoName(filename string) string {
	videoType := GetVideoType(filename)
	return RandStr(10) + videoType
}

/*
获取视频类型 如.mp4
*/
func GetVideoType(filename string) string {
	return filename[strings.LastIndex(filename, "."):]
}
