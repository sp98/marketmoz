input = {bucket: "${INPUT_BUCKET}", measurement: "${INPUT_MEASUREMENT}", every: ${INPUT_EVERY}}

from(bucket: input.bucket)
    |> range(start: -input.every)
    |> filter(fn: (r) => r["_measurement"] == input.measurement)
    |> filter(fn: (r) => r["_field"] == "LastPrice")
    |> last()
