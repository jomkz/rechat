// Copyright 2018 Rechat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rechat

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8000
)

// Run will start the server and listen for connections.
func Run() {
	addr := fmt.Sprintf("%s:%d", defaultHost, defaultPort)

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")

	logrus.Infof("Listening on http://%s/", addr)
	http.ListenAndServe(addr, r)
}

// index will handle requests for the index or "home page" for the application.
func index(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Index handler called...")
}
