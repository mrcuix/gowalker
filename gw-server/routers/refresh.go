// Copyright 2013 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

import (
	"os"
	"strings"

	"github.com/Unknwon/gowalker/doc"
	"github.com/astaxie/beego"
)

// RefreshRouter serves search pages.
type RefreshRouter struct {
	beego.Controller
}

// Get implemented Get method for RefreshRouter.
func (this *RefreshRouter) Get() {
	// Set language version.
	curLang := globalSetting(this.Ctx, this.Input(), this.Data)

	// Get argument(s).
	q := this.Input().Get("q")

	// Empty query string shows home page.
	if len(q) == 0 {
		this.Redirect("/", 302)
		return
	}

	// Set properties.
	this.TplNames = "refresh_" + curLang.Lang + ".html"

	pdoc, err := doc.CheckDoc(q, "", doc.REFRESH_REQUEST)
	if err == nil {
		os.Remove("./static/docs/" + pdoc.ImportPath + ".js")

		this.Redirect("/"+q, 302)
		return
	}

	if strings.HasPrefix(err.Error(), "doc.") || strings.HasSuffix(err.Error(), "EOF") {
		this.Data["IsHasError"] = true
		this.Data["ErrMsg"] = strings.Replace(err.Error(),
			doc.GetGithubCredentials(), "<githubCred>", 1)
		return
	}

	this.Data["Path"] = q
	this.Data["LimitTime"] = err.Error()
}
