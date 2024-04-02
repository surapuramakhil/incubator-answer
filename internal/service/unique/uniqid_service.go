/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package unique

import (
	"context"
	"strconv"
)

// UniqueIDRepo unique id repository
type UniqueIDRepo interface {
	GenUniqueIDStr(ctx context.Context, key string) (uniqueID string, err error)
}

type UniqueIDService struct {
	
}

func NewUniqueIDService() *UniqueIDService {
	return &UniqueIDService{}
}

func (us *UniqueIDService) GetObjectType(uniqueID string) (ObjectType int, err error) {
	objectTypeStr := uniqueID[1:4]
	objectType, err := strconv.Atoi(objectTypeStr)
	if err != nil {
		return 0, err
	}
	return objectType, nil
}
