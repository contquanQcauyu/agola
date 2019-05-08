// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package action

import (
	"net/http"

	"github.com/pkg/errors"
	csapi "github.com/sorintlab/agola/internal/services/configstore/api"
	"github.com/sorintlab/agola/internal/services/gateway/common"
	rsapi "github.com/sorintlab/agola/internal/services/runservice/api"
	"github.com/sorintlab/agola/internal/util"

	"go.uber.org/zap"
)

type ActionHandler struct {
	log               *zap.SugaredLogger
	sd                *common.TokenSigningData
	configstoreClient *csapi.Client
	runserviceClient  *rsapi.Client
	agolaID           string
	apiExposedURL     string
	webExposedURL     string
}

func NewActionHandler(logger *zap.Logger, sd *common.TokenSigningData, configstoreClient *csapi.Client, runserviceClient *rsapi.Client, agolaID, apiExposedURL, webExposedURL string) *ActionHandler {
	return &ActionHandler{
		log:               logger.Sugar(),
		sd:                sd,
		configstoreClient: configstoreClient,
		runserviceClient:  runserviceClient,
		agolaID:           agolaID,
		apiExposedURL:     apiExposedURL,
		webExposedURL:     webExposedURL,
	}
}

func ErrFromRemote(resp *http.Response, err error) error {
	if err == nil {
		return nil
	}

	if resp != nil {
		switch resp.StatusCode {
		// remove wrapping from errors sent to client
		case http.StatusBadRequest:
			return util.NewErrBadRequest(errors.Cause(err))
		case http.StatusNotFound:
			return util.NewErrNotFound(errors.Cause(err))
		}
	}

	return err
}