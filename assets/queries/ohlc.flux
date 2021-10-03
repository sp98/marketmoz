input = {bucket: "${INPUT_BUCKET}", measurement: "${INPUT_MEASUREMENT}", every: "${INPUT_EVERY}"}

from(bucket: input.bucket)
    |> range(start: -input.every)
    |> tail(n:100)
    |> filter(fn: (r) => r["_measurement"] == input.measurement)
    |> filter(fn: (r) => r["_field"] == "Close" or r["_field"] == "High" or r["_field"] == "Low" or r["_field"] == "Open")
