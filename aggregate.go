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
	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/go/pkg/beam/io/bigqueryio"
	"reflect"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"flag"
	"context"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/adrianwit/beam-etl/model"
	"fmt"
	"encoding/json"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
)


func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	project := gcpopts.GetProject(ctx)


	log.Info(ctx, "Running query")
	p := beam.NewPipeline()
	s := p.Root()

	events := bigqueryio.Query(s, project, model.EventDQL, reflect.TypeOf(model.Event{}))
	groupedPerformance := beam.ParDo(s, groupByKey, events)

	beam.GroupByKey(s, groupedPerformance)

	aggregated := beam.CombinePerKey(s, aggregate, groupedPerformance)

	formatted := beam.ParDo(s, demandSupplyPerformanceFormatFn, aggregated)

	textio.Write(s, "/tmp/aggregated_events.json", formatted)
	if err := beamx.Run(ctx, p); err != nil {
		log.Exitf(ctx, "Failed to execute job: %v", err)
	}
}


func groupByKey(event model.Event) (string, model.DemandSupplyPerformance) {
	var performance =  model.DemandSupplyPerformance{
		DemandSideId:event.DemandSideId,
		SupplySideId:event.SupplySideSideId,
		Charge:event.Charge,
		Payment:event.Payment,
	}
	switch event.EventTypeId  {
	case 1:
		performance.EventOneTypeCount = 1
	case 2:
		performance.EventTwoTypeCount = 1
	}
	return fmt.Sprintf("%v-%v", event.DemandSideId, event.SupplySideSideId), performance
}


func aggregate(x, y model.DemandSupplyPerformance) model.DemandSupplyPerformance {
	return model.DemandSupplyPerformance{
		DemandSideId:x.DemandSideId,
		SupplySideId:x.SupplySideId,
		Charge: x.Charge + y.Charge,
		Payment: x.Payment+y.Payment,
		EventOneTypeCount:x.EventOneTypeCount+y.EventOneTypeCount,
		EventTwoTypeCount:x.EventTwoTypeCount+y.EventTwoTypeCount,
	}
}



func demandSupplyPerformanceFormatFn(_ string, row model.DemandSupplyPerformance) string {
	data, err := json.Marshal(row)
	if err == nil {
		return (string(data))
	}
	return ""
}

