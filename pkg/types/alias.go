/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

import (
	"mosn.io/api"
	"mosn.io/pkg/buffer"
	"sync"
)

// use alias to keep compatiable
type IoBuffer = buffer.IoBuffer

type Protocol = api.Protocol

type HeaderMap = api.HeaderMap

type HostInfo = api.HostInfo

type RequestInfo = api.RequestInfo

type Route = api.Route

type StreamBuffer struct {
	buffer.IoBuffer
	End chan int
	mu  sync.Mutex
}

func (b *StreamBuffer) SynAppend(data []byte) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Append(data)
}

func (b *StreamBuffer) SynLen() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Len()
}

func (b *StreamBuffer) SynRead(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Read(p)
}
