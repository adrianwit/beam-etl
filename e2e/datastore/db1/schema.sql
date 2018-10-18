CREATE OR REPLACE TABLE events (
  id INT64,
  sessionId STRING,
  eventTypeId INT64,
  demandSideId INT64,
  subjectId INT64,
  supplySideSideId INT64,
  locationId INT64,
  charge FLOAT64,
  payment FLOAT64,
  timestamp TIMESTAMP
);

CREATE OR REPLACE TABLE locations (
  id INT64,
  typeId int64,
  charge FLOAT64,
  revenueShare FLOAT64
);

