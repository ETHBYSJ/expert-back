package conf

// 默认值
var SystemConfig = &SystemConf{}

// 默认配置
func getDefault() {
	SystemConfig.System.Listen = ":8888"

	SystemConfig.File.Download.Recommend.Path = "./static/download/recommend.docx"
	SystemConfig.File.Download.Recommend.Name = "长三角区域教育评价变革协作联盟专家库成员推荐汇总表.docx"
	SystemConfig.File.Download.Apply.Path = "./static/download/apply.docx"
	SystemConfig.File.Download.Apply.Name = "长三角区域教育评价变革协作联盟专家库成员申请表.docx"
	SystemConfig.File.Upload.Recommend.Path = "./static/upload/recommend"
	SystemConfig.File.Upload.Apply.Path = "./static/upload/apply"
	SystemConfig.File.Upload.Picture.Path = "./static/upload/picture"

	SystemConfig.Database.Name = "expert"
	SystemConfig.Database.Connection = "mongodb://202.120.39.3:27017"
}
