// +build dev

package server

import (
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

// Initialize crawler entry
func Initialize() (err error) {
	m := macaron.Classic()
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	m.Use(gzip.Gziper())
	m.Use(session.Sessioner())
	m.Get("/get", func(ctx *macaron.Context) {
		ctx.Data["Title"] = "jeremy"
		ctx.HTML(200, "index") // 200 为响应码
	})
	m.Run()
	return
}
