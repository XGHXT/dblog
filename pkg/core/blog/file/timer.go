// Package file provides ...
package file

import (
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/XGHXT/dblog/pkg/cache"
	"github.com/XGHXT/dblog/pkg/config"
	"github.com/XGHXT/dblog/tools"

	"github.com/sirupsen/logrus"
)

var xmlTmpl *template.Template

func init() {
	root := filepath.Join(config.WorkDir, "website", "template", "*.xml")

	var err error
	xmlTmpl, err = template.New("").Funcs(template.FuncMap{
		"dateformat":  tools.DateFormat,
		"imgtonormal": tools.ImgToNormal,
	}).ParseGlob(root)
	if err != nil {
		panic(err)
	}
	generateOpensearch()
	generateRobots()
	generateCrossdomain()
	go timerSitemap()
}

// timerSitemap 定时刷新sitemap
func timerSitemap() {
	tpl := xmlTmpl.Lookup("sitemapTpl.xml")
	if tpl == nil {
		logrus.Info("file: not found: sitemapTpl.xml")
		return
	}

	params := map[string]interface{}{
		"Articles": cache.Ei.Articles,
		"Host":     config.Conf.BlogApp.Host,
	}
	f, err := os.OpenFile("assets/sitemap.xml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Error("file: timerSitemap.OpenFile: ", err)
		return
	}
	defer f.Close()
	err = tpl.Execute(f, params)
	if err != nil {
		logrus.Error("file: timerSitemap.Execute: ", err)
		return
	}
	time.AfterFunc(time.Hour*24, timerSitemap)
}

// generateOpensearch 生成opensearch.xml
func generateOpensearch() {
	tpl := xmlTmpl.Lookup("opensearchTpl.xml")
	if tpl == nil {
		logrus.Info("file: not found: opensearchTpl.xml")
		return
	}
	params := map[string]string{
		"BTitle":   cache.Ei.Blogger.BTitle,
		"SubTitle": cache.Ei.Blogger.SubTitle,
		"Host":     config.Conf.BlogApp.Host,
	}
	f, err := os.OpenFile("assets/opensearch.xml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Error("file: generateOpensearch.OpenFile: ", err)
		return
	}
	defer f.Close()
	err = tpl.Execute(f, params)
	if err != nil {
		logrus.Error("file: generateOpensearch.Execute: ", err)
		return
	}
}

// generateRobots 生成robots.txt
func generateRobots() {
	tpl := xmlTmpl.Lookup("robotsTpl.xml")
	if tpl == nil {
		logrus.Info("file: not found: robotsTpl.xml")
		return
	}
	params := map[string]string{
		"Host": config.Conf.BlogApp.Host,
	}
	f, err := os.OpenFile("assets/robots.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Error("file: generateRobots.OpenFile: ", err)
		return
	}
	defer f.Close()
	err = tpl.Execute(f, params)
	if err != nil {
		logrus.Error("file: generateRobots.Execute: ", err)
		return
	}
}

// generateCrossdomain 生成crossdomain.xml
func generateCrossdomain() {
	tpl := xmlTmpl.Lookup("crossdomainTpl.xml")
	if tpl == nil {
		logrus.Info("file: not found: crossdomainTpl.xml")
		return
	}
	params := map[string]string{
		"Host": config.Conf.BlogApp.Host,
	}
	f, err := os.OpenFile("assets/crossdomain.xml", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Error("file: generateCrossdomain.OpenFile: ", err)
		return
	}
	defer f.Close()
	err = tpl.Execute(f, params)
	if err != nil {
		logrus.Error("file: generateCrossdomain.Execute: ", err)
		return
	}
}
