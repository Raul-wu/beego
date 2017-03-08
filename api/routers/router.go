// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/t_feedback",
			beego.NSInclude(
				&controllers.TFeedbackController{},
			),
		),

		beego.NSNamespace("/t_feedback_action",
			beego.NSInclude(
				&controllers.TFeedbackActionController{},
			),
		),

		beego.NSNamespace("/t_feedback_feature",
			beego.NSInclude(
				&controllers.TFeedbackFeatureController{},
			),
		),

		beego.NSNamespace("/t_feedback_status",
			beego.NSInclude(
				&controllers.TFeedbackStatusController{},
			),
		),

		beego.NSNamespace("/t_feedback_support_poc",
			beego.NSInclude(
				&controllers.TFeedbackSupportPocController{},
			),
		),

		beego.NSNamespace("/t_feedback_type",
			beego.NSInclude(
				&controllers.TFeedbackTypeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
