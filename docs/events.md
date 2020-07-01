## Events

StatusBay provides centralized view with all the resources kinds event that are related to your deployment.

Example:
![Events](/docs/images/events.png)

The list of the events we manage is in [events.yaml](/events.yaml) file.
each event is identified by a top-level resource kind,
and contains multiple (pattern,description) pairs, the first pattern to match displays has its associated descriptions displayed to the user.

How to contribute:
1. fork us and edit [events.yaml](/events.yaml) with event data.
2. open a pull-request.


