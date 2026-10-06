connection "sentinelone" {
  plugin = "xybytes/sentinelone"

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
