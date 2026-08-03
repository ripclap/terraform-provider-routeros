resource "routeros_tool_graphing" "graphing" {
  page_refresh = "10s"
  store_every  = "5min"
}
