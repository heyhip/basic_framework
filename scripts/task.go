package scripts

// 定时任务
// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
// * * * * *
// 分 小时 天 月 周
func Execute() {

	//// 清除过期视频、删除视频
	//go func() {
	//	// 定时任务
	//	c := cron.New()
	//
	//	// 每周日，凌晨4到5点，每5分钟执行一次
	//	c.AddFunc("*/5 4-5 * * 0", func() {
	//		// 删除过期未上传的视频
	//		server.UserVideoServer{}.DelExpireVideo()
	//	})
	//
	//	c.Start()
	//}()

	//// 备份数据库
	//go func() {
	//	c := cron.New()
	//
	//	// 执行备份脚本，生成压缩文件
	//	c.AddFunc("0 5 * * *", func() {
	//		commond := "/data/release/bash/hp_mysql_backups.sh"
	//		err := exec.Command("/bin/bash", "-c", commond).Run()
	//		if err != nil {
	//			logging.Error(err.Error())
	//		}
	//	})
	//
	//	// 邮件发送备份文件
	//	c.AddFunc("30 5 * * *", func() {
	//		ymd := time.Unix(time.Now().Unix(), 0).Format(config.Yaml.YMD)
	//		file := fmt.Sprintf("heyhip_%v.sql.zip", ymd)
	//		title := fmt.Sprintf("%v-备份压缩文件", ymd)
	//		gemail.SendEmailFile(title, "备份", true, strings.Split(config.Yaml.BackupSqlEmail, ","), "/data/release/backups/"+file)
	//	})
	//
	//	c.Start()
	//}()

}
