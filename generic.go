// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"flag"
	"context"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/adrianwit/beam-etl/model"
	"reflect"
	"github.com/apache/beam/sdks/go/pkg/beam/io/databaseio"
	"fmt"
)


func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	log.Info(ctx, "Running query")
	p := beam.NewPipeline()
	s := p.Root()


	subjects := databaseio.Query(s, "mysql", "root:dev@tcp(127.0.0.1:3306)/db2?parseTime=true", model.SubjectDQL, reflect.TypeOf(model.Record{}))

	fmt.Printf("len: %v\n", (subjects))
	//textio.Write(s, "/tmp/filtered_subjects.json", subjects)
	if err := beamx.Run(ctx, p); err != nil {
		log.Exitf(ctx, "Failed to execute job: %v", err)
	}



}


