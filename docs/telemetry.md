# Telemetry

StatusBay application collects runtime metrics related to the performance of the system.

## How to setup?



to receive metrics in
* datadog you need to pass dogstatsd_addr in the config file 
* prometheus you need to pass prometheus_retention_time_sec in the config file and have it be > 0
* statsd you need to pass statsd_address in the config file
* statsite you need to pass statsite_address in the config file 


## List of metrics

statusbay.runtime.num_goroutines
statusbay.runtime.alloc_bytes
statusbay.runtime.malloc_count
statusbay.runtime.free_count
statusbay.runtime.heap_objects
statusbay.runtime.total_gc_pause_ns
statusbay.runtime.total_gc_runs
