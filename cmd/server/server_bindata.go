// +build !dev

package server

import (
	"github.com/go-macaron/bindata"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"

	"github.com/tosone/panx/bindata/public"
	"github.com/tosone/panx/bindata/templates"
)

// Initialize crawler entry
func Initialize() (err error) {
	m := macaron.Classic()
	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      public.Asset,
				AssetDir:   public.AssetDir,
				AssetNames: public.AssetNames,
				AssetInfo:  templates.AssetInfo,
			}),
		},
	))
	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			TemplateFileSystem: bindata.Templates(bindata.Options{
				Asset:      templates.Asset,
				AssetDir:   templates.AssetDir,
				AssetInfo:  templates.AssetInfo,
				AssetNames: templates.AssetNames,
			}),
		},
	))
	m.Use(gzip.Gziper())
	m.Use(session.Sessioner())
	m.Get("/get", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "jeremy"
		ctx.HTML(200, "index") // 200 为响应码
	})
	m.Run()
	return
}
