# Telemetry

StatusBay application collects runtime metrics related to the performance of the system.

## How to setup?



to receive metrics in
* datadog you need to pass dogstatsd_addr in the config file 
* prometheus you need to pass prometheus_retention_time_sec in the config file and have it be > 0
* statsd you need to pass statsd_address in the config file
* statsite you need to pass statsite_address in the config file 


## List of metrics

* `statusbay.runtime.num_goroutines` the number of active goroutines
* `statusbay.runtime.alloc_bytes` the number of allocated bytes by the process
* `statusbay.runtime.malloc_count` the cumulative count of heap objects allocated
* `statusbay.runtime.free_count` the cumulative count of heap objects freed
* `statusbay.runtime.heap_objects` the number of currently allocated heap objects
* `statusbay.runtime.total_gc_pause_ns` the cumulative nanoseconds in GC stop-the-world pauses since the program started
* `statusbay.runtime.total_gc_runs` the number of completed GC cycles
