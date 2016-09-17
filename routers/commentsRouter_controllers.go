package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sekolah/controllers:MakanController"] = append(beego.GlobalControllerRouter["sekolah/controllers:MakanController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:MakanController"] = append(beego.GlobalControllerRouter["sekolah/controllers:MakanController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:MakanController"] = append(beego.GlobalControllerRouter["sekolah/controllers:MakanController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:MakanController"] = append(beego.GlobalControllerRouter["sekolah/controllers:MakanController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:MakanController"] = append(beego.GlobalControllerRouter["sekolah/controllers:MakanController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:PelajarController"] = append(beego.GlobalControllerRouter["sekolah/controllers:PelajarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:PelajarController"] = append(beego.GlobalControllerRouter["sekolah/controllers:PelajarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:PelajarController"] = append(beego.GlobalControllerRouter["sekolah/controllers:PelajarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:PelajarController"] = append(beego.GlobalControllerRouter["sekolah/controllers:PelajarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sekolah/controllers:PelajarController"] = append(beego.GlobalControllerRouter["sekolah/controllers:PelajarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
