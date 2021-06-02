/**
 * Copyright 2020 IBM Corp.
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

// Package messages ...
package messages

import (
	"github.com/IBM/ibm-csi-common/pkg/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func TestGetCSIError(t *testing.T) {
	MessagesEn = InitMessages()
	ctxLog, reqID := utils.GetContextLogger(context.Background(), false)

	err := GetCSIError(ctxLog, MethodUnimplemented, reqID, nil)
	assert.NotNil(t, err)

	err1 := GetCSIError(ctxLog, MethodUnimplemented, reqID, err)
	assert.NotNil(t, err1)
}

func TestGetCSIMessage(t *testing.T) {
	message := GetCSIMessage(MethodUnimplemented, "create")
	assert.NotNil(t, message)
}
