package main

import (
	"flag"
	"libstorage"
	"os"
	"path/filepath"
	"runtime"
	"util/file"
	"util/logger"
	"validate"
	"app"
)

// 当客户端下载文件的时候，如果文件尚未在组内全部同步完成，
// 并且恰好访问到没有同步完成的机器时，客户端会将请求重定向到文件原始服务器
// exp: /G001(组)/01(原始服务器实例ID)/M[S](单片or多片)/{MD5}[.ext]
// 文件的原始名称需要客户端自行记录（可能未来加上服务端记录功能）
// TODO support detect total file size for assigning
// TODO Optimize log information
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.RUN_WITH = 1
	abs, _ := filepath.Abs(os.Args[0])
	s, _ := filepath.Split(abs)
	s = file.FixPath(s)
	var confPath = flag.String("c", s+string(filepath.Separator)+".."+string(filepath.Separator)+"conf"+string(filepath.Separator)+"storage.conf", "custom config file")
	flag.Parse()
	logger.Info("using config file:", *confPath)
	m, e := file.ReadPropFile(*confPath)
	if e == nil {
		validate.Check(m, app.RUN_WITH)
		for k, v := range m {
			logger.Debug(k, "=", v)
		}
		libstorage.StartService(m)
	} else {
		logger.Fatal("error read file:", e)
	}
}
