# metrics2docs

simple tool to generate a markdown file with info about all documented metrics

# doc format

metrics need single or multi-line comments that are in the form `metric <metricname> is <description>`

Example:

```
// metric metrics_too_old is a counter of points that go back in time.
// E.g. for any given series, when a point has a timestamp
// that is not higher than the timestamp of the last written timestamp for that series.
metricsTooOld met.Count

// metric add_to_saving_chunk is points received - by the primary node - for the most recent chunk
// when that chunk is already being saved (or has been saved).
// this indicates that your GC is actively sealing chunks and saving them before you have the chance to send
// your (infrequent) updates.  The primary won't add them to its in-memory chunks, but secondaries will
// (because they are never in "saving" state for them), see below.
addToSavingChunk met.Count

// metric add_to_saved_chunk is points received - by a secondary node - for the most recent chunk when that chunk
// has already been saved by a primary.  A secondary can add this data to its chunks.
addToSavedChunk met.Count

memToIterDuration met.Timer
persistDuration   met.Timer

metricsActive met.Gauge // metric metrics_active is the amount of currently known metrics (excl rollup series), measured every second
gcMetric      met.Count // metric gc_metric is the amount of times the metrics GC is about to inspect a metric (series)
```

# usage

```
metrics2docs <path-to-codebase>
```

Example:

```
metrics2docs $GOPATH/src/github.com/raintank/metrictank > metrics.md
```
