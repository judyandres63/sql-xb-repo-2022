/*
 * Copyright 2020 io.xream.sqlxb
 *
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
package sqlxb

// To build sql, like: SELECT DISTINCT f.id FROM foo f INNER JOIN (SELECT foo_id FROM bar) b ON b.foo_id = f.id
// Sql for MySQL, Clickhouse....
//
// @author Sim
//
type BuilderX struct {
	Builder
	resultKeys   []string
	sbs          []*SourceBuilder
	sourceScript *string
	sourceValues []interface{}
}

func NewBuilderX() *BuilderX {
	return new(BuilderX)
}

func (x *BuilderX) SourceBuilder() *SourceBuilder {
	var sb = SourceBuilder{}
	x.sbs = append(x.sbs, &sb)
	return &sb
}

func (x *BuilderX) SourceScript(script string, arr ...interface{}) *BuilderX {
	x.sourceScript = &script
	if arr != nil {
		for _, v := range arr {
			x.sourceValues = append(x.sourceValues, v)
		}
	}
	return x
}

func (x *BuilderX) ResultKey(resultKey string) *BuilderX {
	x.resultKeys = append(x.resultKeys, resultKey)
	return x
}

func (x *BuilderX) ResultKeys(resultKeys ...string) *BuilderX {
	for _, resultKey := range resultKeys {
		x.resultKeys = append(x.resultKeys, resultKey)
	}
	return x
}

func (x *BuilderX) Having(op Op, k string, v interface{}) *BuilderX {
	if op == nil || k == "" {
		return x
	}
	bb := Bb{
		op:    op(),
		key:   k,
		value: v,
	}
	x.havings = append(x.havings, &bb)
	return x
}

func (x *BuilderX) GroupBy(groupBy string) *BuilderX {
	if groupBy == "" {
		return x
	}
	x.groupBys = append(x.groupBys, groupBy)
	return x
}
