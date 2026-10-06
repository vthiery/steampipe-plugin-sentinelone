---
organization: Turbot
category: ["security"]
brand_color: "#6100FF"
display_name: SentinelOne
name: sentinelone
description: SentinelOne is an autonomous cybersecurity platform that provides comprehensive protection against various cyber threats.
og_description: Query SentinelOne data with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/vthiery/sentinelone-social-graphic.png"
icon_url: "/images/plugins/vthiery/sentinelone.svg"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# SentinelOne + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[SentinelOne](https://sentinelone.com) provides cloud workload and endpoint security, threat intelligence, and cyberattack response services.

For example:

```sql
select
  alert_id,
  src_ip,
  dst_ip
from 
  sentinelone_alerts
limit 50;
```

## Documentation

- **[Table definitions & examples →](https://github.com/vthiery/steampipe-plugin-sentinelone/tree/main/docs/tables)**

## Get started

### Install

Download and install the latest SentinelOne plugin:

```shell
steampipe plugin install ghcr.io/vthiery/sentinelone
```

### Configuration

Installing the latest sentinelone plugin will create a config file (`~/.steampipe/config/sentinelone.spc`) with a single connection named `sentinelone`:

```hcl
connection "sentinelone" {
  plugin     = "sentinelone"

  # SentinelOne client ID
  # Can also be set with the SENTINELONE_CLIENT_ID environment variable
  # client_id        = "companyname"
  
  # SentinelOne JWT Token
  # Can also be set with the SENTINELONE_API_TOKEN environment variable
  # api_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"

  # HTTP request timeout in seconds for each SentinelOne API call.
  # Increase this if queries time out on large tenants. Default: 30.
  # request_timeout = 30

  # Number of items requested per API page. Defaults to the API maximum
  # of 1000. Lower this if large pages time out or return 502 on your
  # tenant.
  # page_size = 1000
}
```

- `client_id` - (Required) The client ID. Can also be set with the `SENTINELONE_CLIENT_ID` environment variable.
- `api_token` - (Required) The API Access Token. Can also be set with the `SENTINELONE_API_TOKEN` environment variable.
- `request_timeout` - (Optional) HTTP timeout in seconds applied to each API call. Defaults to `30`. Raise it on large tenants where a single page takes longer than that to come back.
- `page_size` - (Optional) Number of items requested per API page, clamped to the API maximum of `1000` (also the default).

  SentinelOne documents `1000` as the maximum page size, but the documented
  maximum is not always the usable one: some tenants' backends return
  `502 Bad Gateway` or stall the connection on large pages for busier
  endpoints. If full table scans fail with a 502 or a hydrate timeout while
  a `limit 1` query succeeds, lower this — `200` is a good starting point.
