input = {bucket: "test", measurement: "nifty-1m", every: 30m}

from(bucket: input.bucket)
    |> range(start: -input.every)
    |> filter(fn: (r) => r["_measurement"] == input.measurement)
    |> filter(fn: (r) => r["_field"] == "Close" or r["_field"] == "High" or r["_field"] == "Low" or r["_field"] == "Open")
