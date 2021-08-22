taskName = "${TASK_NAME}"
from = {bucket: "${FROM_BUCKET}", measurement: "${FROM_MEASUREMENT}"}
to = {bucket: "${TO_BUCKET}", measurement: "${TO_MEASUREMENT}"}
dsTime = "${DS_TIME}"

option task = {name: taskName, every: dsTime}

time = now()
DATA = from(bucket: from.bucket)
    |> range(start: -task.every)
    |> filter(fn: (r) => r["_measurement"] == from.measurement)
    |> filter(fn: (r) => r["_field"] == "LastPrice")
    |> filter(fn: (r) => r["stock"] == "nifty")
    |> drop(columns: ["_start", "_stop"])

DATA
    |> first()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Open")
    |> set(key: "_measurement", value: to.measurement)
    |> to(bucket: to.bucket)
DATA
    |> max()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "High")
    |> set(key: "_measurement", value: to.measurement)
    |> to(bucket: to.bucket)
DATA
    |> min()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Low")
    |> set(key: "_measurement", value: to.measurement)
    |> to(bucket: to.bucket)
DATA
    |> last()
    |> map(fn: (r) => ({r with _time: time}))
    |> set(key: "_field", value: "Close")
    |> set(key: "_measurement", value: to.measurement)
    |> to(bucket: to.bucket)
