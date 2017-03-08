package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackActionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackFeatureController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackSupportPocController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"] = append(beego.GlobalControllerRouter["beego/api/controllers:TFeedbackTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
