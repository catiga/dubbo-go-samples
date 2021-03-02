// +build integration

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

package integration

import (
	"github.com/apache/dubbo-go/config"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/metadata/service/inmemory"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

import (
	"os"
	"testing"
	"time"
)

var cat = new(CatService)
var dog = new(DogService)
var tiger = new(TigerService)
var lion = new(LionService)

func TestMain(m *testing.M) {
	config.SetConsumerService(cat)
	config.SetConsumerService(dog)
	config.SetConsumerService(tiger)
	config.SetConsumerService(lion)
	config.Load()
	time.Sleep(3 * time.Second)

	os.Exit(m.Run())
}

type CatService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (c *CatService) Reference() string {
	return "CatService"
}

type DogService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (d *DogService) Reference() string {
	return "DogService"
}

type TigerService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (t *TigerService) Reference() string {
	return "TigerService"
}

type LionService struct {
	GetId   func() (int, error)
	GetName func() (string, error)
	Yell    func() (string, error)
}

func (l *LionService) Reference() string {
	return "LionService"
}
