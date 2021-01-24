# Ispdb Api Handler

### About
This is simple go-application for looking imap/pop3/smtp config from ISP DB by Mozilla (Thunderbird autoconfig) with local caching configs and web api.

### Features

Flags to execute
--update-threads=100 - default 100 threads at the pool for downloader cache
--port=37896 - port for api handler
--update-cache=1 - need update cache from ISPDB mozilla

Just get api request:
-> localhost:37896/find/gmail.com
You will get answer with json config OR message - Not Found!
