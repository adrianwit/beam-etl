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
	"context"
	"flag"
	"reflect"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/io/bigqueryio"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/apache/beam/sdks/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"encoding/json"
	"github.com/adrianwit/beam-etl/model"
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
	locations := bigqueryio.Query(s, project, model.LocationDQL, reflect.TypeOf(model.Location{}))

	joinedEvents := joinEvents(s, events, locations)

	formatted := beam.ParDo(s, eventFormatFn, joinedEvents)
	textio.Write(s, "/tmp/joined_events.json", formatted)
	if err := beamx.Run(ctx, p); err != nil {
		log.Exitf(ctx, "Failed to execute job: %v", err)
	}
}

func joinEvents(s beam.Scope, events, locations beam.PCollection) beam.PCollection {
	joined := beam.CoGroupByKey(s,
		beam.ParDo(s, eventByLocationID, events),
		beam.ParDo(s, locationByID, locations))
	return beam.ParDo(s, joinFn, joined)
}


func joinFn(locationId int, events, locations func(*model.Event) bool, emit func(int, model.Event)) {
	var locationEvent = model.Event{}
	locations(&locationEvent) // grab first (and only) location, if any
	var event = model.Event{}
	for events(&event) {
		event.Location = locationEvent.Location
		emit(locationId, event)
	}
}

func eventByLocationID(event model.Event) (int, model.Event) {
	return event.LocationID, event
}
func locationByID(location model.Location) (int, model.Event) {
	return location.ID, model.Event{Location: location}
}

func eventFormatFn(_ int, row model.Event) string {
	data, err := json.Marshal(row)
	if err == nil {
		return (string(data))
	}
	return ""
}
