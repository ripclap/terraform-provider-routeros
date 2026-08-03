resource "routeros_tool_traffic_generator" "generator" {
  latency_distribution_max = "example"
  measure_out_of_order     = true
  stats_samples_to_keep    = 1
}
