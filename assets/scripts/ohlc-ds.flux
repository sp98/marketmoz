input = {bucket: "${INPUT_BUCKET}", measurement: "${INPUT_MEASUREMENT}"}
output = {bucket: "${OUTPUT_BUCKET}", measurement: "${OUTPUT_MEASUREMENT}"}

time = now()
DATA = from(bucket: input.bucket)
    |> range(start: -task.every)
    |> filter(fn: (r) => r["_measurement"] == input.measurement)
    |> filter(fn: (r) => r["_field"] == "LastPrice")
    |> filter(fn: (r) => r["stock"] == "nifty")
    |> drop(columns: ["_start", "_stop"])

DATA
    |> first()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Open")
    |> set(key: "_measurement", value: output.measurement)
    |> to(bucket: output.bucket)
DATA
    |> max()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "High")
    |> set(key: "_measurement", value: output.measurement)
    |> to(bucket: output.bucket)
DATA
    |> min()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Low")
    |> set(key: "_measurement", value: output.measurement)
    |> to(bucket: output.bucket)
DATA
    |> last()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Close")
    |> set(key: "_measurement", value: output.measurement)
    |> to(bucket: output.bucket)
