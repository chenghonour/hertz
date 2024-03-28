/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package generator

import (
	"regexp"
	"testing"
)

func TestMatchInsertPointPatternNew(t *testing.T) {
	regRegisterV3 = regexp.MustCompile(insertPointPatternNew)
	
	ValidPatternFunc := func(s string) bool {
		subIndexReg := regRegisterV3.FindSubmatchIndex([]byte(s))
		if len(subIndexReg) != 2 || subIndexReg[0] < 1 {
			return false
		}
		return true
	}

	type args struct {
		name        string
		fileContent string
	}
	tests := []struct {
		name      string
		args      args
		WantValid bool
	}{
		{"TestInsertPointPatternNoSpace", args{"test", `// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	admin "export-go/biz/router/admin"
	export "export-go/biz/router/export"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	admin.Register(r)
}`}, true},
		{"TestInsertPointPatternHasSpace", args{"test", `// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	admin "export-go/biz/router/admin"
	export "export-go/biz/router/export"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	// INSERT_POINT: DO NOT DELETE THIS LINE!
	admin.Register(r)
}`}, true},
		{"TestInsertPointPatternWrongFormat", args{"test", `// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	admin "export-go/biz/router/admin"
	export "export-go/biz/router/export"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	admin.Register(r)
}`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPatternFunc(tt.args.fileContent); got != tt.WantValid {
				t.Errorf("wrong format %s: insert-point '%s' not found", tt.args.fileContent, insertPointNew)
			}
		})
	}
}
