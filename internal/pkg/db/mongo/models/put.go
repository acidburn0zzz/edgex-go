/*******************************************************************************
 * Copyright 2018 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/pkg/errors"
)

type Put struct {
	Action         `bson:",inline"`
	ParameterNames []string `bson:"parameterNames"`
}

func (p *Put) ToContract() contract.Put {
	return contract.Put{
		Action: p.Action.ToContract(),
	}
}

func (p *Put) FromContract(from contract.Put) error {
	action := &Action{}
	err := action.FromContract(from.Action)
	if err != nil {
		return errors.New(err.Error() + " path: " + from.Path)
	}
	return nil
}