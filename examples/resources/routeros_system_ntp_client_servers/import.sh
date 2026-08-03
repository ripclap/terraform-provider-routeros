#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/system/ntp/client/servers get [print show-ids]]
terraform import routeros_system_ntp_client_servers.servers "*1"
