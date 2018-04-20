// Copyright 2014 Unknwon
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

package admin

import (
	"github.com/gpmgo/switch/models"
	"github.com/gpmgo/switch/pkg/middleware"
)

func Revisions(ctx *middleware.Context) {
	ctx.Data["PageIsPackages"] = true
	ctx.Data["PageIsPackagesList"] = true
	ctx.HTML(200, "packages/list")
}

func LargeRevisions(ctx *middleware.Context) {
	ctx.Data["PageIsPackages"] = true
	ctx.Data["PageIsPackagesLarges"] = true

	models.GetLocalRevisions()

	ctx.HTML(200, "packages/larges")
}
