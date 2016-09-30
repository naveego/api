Data Quality API
----------------

Data quality data is made accessible through the use of the Data Qaulity API.

## Trending Exceptions

Naveego keeps statistical data about the overall health of the customers 
data quality.  Each time a quality check runs, the platform will use historical
run information to calculate a z-score.  

Additional information on z-scores: [https://en.wikipedia.org/wiki/Standard_score](https://en.wikipedia.org/wiki/Standard_score)

Using the trending API endpoint, you can retrieve the top 5 trending exceptions based on 
z-score.  This list will represent the quality checks that have changed most drastically
with respect to recent history.

Example Request:
```json
GET /v3/dataquality/trending
```

Example Response:
```json
{
  // An array of trending exception data points
  "data": [
    {
      "id": "14f0c95f-3455-477b-a82b-f30cd5c798ac",
      "runId": "9543d716-4310-43e4-a632-f4d82de8d321",
      "status": "open",
      "state": "complete",
      "transferState": "complete",
      "startedAt": "2016-09-29T13:00:00.042Z",
      "finishedAt": "2016-09-29T13:00:02.073Z",
      "transferCompletedAt": "2016-09-29T13:00:02.53Z",
      "impact": "multi-department",
      "severity": "critical",
      "assignedTo": "jdoe",
      "category": "somecategory",
      "class": "accuracy",
      "object": "customer",
      "property": "name",
      "tags": [
          "customers"
      ],
      "exceptionCount": 2,
      "population": 11866,
      "runtime": 537,
      "commentCount": 0,
      "source": {
        "key": "300acebf-ac33-b94c-b640-e5aaf8ee08f5",
        "name": "CRM"
      },
      "query": {
        "key": "14f0c95f-3455-477b-a82b-f30cd5c798ac",
        "name": "Customers Missing Name"
      },
      "rule": {
        "key": "08d2e0a4-aa24-b4ca-4981-3a1d34000c9f",
        "name": "Customers Missing Name"
      },
      "syncClient": {
        "key": "9fefdf35-f6dd-4c41-998f-e578ea3f3a3e",
        "name": "My SyncClient"
      },
      // This represents the trending data
      "zScores": {
        "exceptions": 5.1994694689574521,
        "population": 2.7557215666932531,
        "runtimes": 0.13356788519040333
      },
      // This can be used to determine the direction
      // of change.  For example a positive number means
      // the trend is increasing while a negative means
      // it is decreasing.
      "slopes": {
        "exceptions": 2,
        "population": 23,
        "runtimes": -24
      },
      // An array of recent exception counts sorted
      // from oldest to most recent. 
      "exceptionCounts": [
        0,
        0,
        ...
        2
      ],
      // An arry of recent populations sorted
      // from oldest to most recent.
      "populations": [
        11815,
        11823,
        ...
        11866
      ],
      // An array of recent runtimes sorted
      // from oldest to most recent.
      "runtimes": [
        314,
        567,
        ...
        537
      ],
      // An array of the previous runids sorted
      // from oldest to most recent.
      "runIds": [
        "6dabb49e-1296-4027-ba52-d56288c9d3a2",
        "fcb743a5-dee4-4a2a-9df8-e63676cdd348",
        ...
        "9543d716-4310-43e4-a632-f4d82de8d321"
      ],
      "dates": [
        "2016-08-21T13:00:00.025Z",
        "2016-08-26T13:00:00.031Z",
        ...
        "2016-09-29T13:00:00.042Z"
      ],
      "modified": "2016-09-29T13:00:02.295Z",
      "created": "2016-06-17T13:00:30.351Z"
    },
    ...
  ]
}
```

