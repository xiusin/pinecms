package markdown

import (
	"github.com/xiusin/pine"
)

var urlPrefix = "/markdown"
var viewsDir = "dcat-page/resources/templates/public/"
var docsDir = "dcat-page/resources/templates/docs"
var configFilePath = "dcat-page/resources/templates/config.json"

const DocMiss = `<div class="the-404">
<div class="contain">
	<div class="content pl-6">
		<h3>You seem to have upset the delicate internal balance of my housekeeper.</h3>
	</div>
</div>
</div>`

func InitRouter(app *pine.Application, router *pine.Router) {
	app.GET(urlPrefix+"/docs/*version", Docs, InjectVar())
	app.GET(urlPrefix+"/*view", Views, InjectVar())
}
