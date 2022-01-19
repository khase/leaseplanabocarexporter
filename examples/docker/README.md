# Overview

This docker image can be used to regulary fetch the leaseplan api and when the pipeline detects any changes send a notification via telegram

# Usage

In order to be able to fetch the data the container needs a `.leaseplanabocar` config with a valid token mounted at `/root/.leaseplanabocar`.
For simplification just execute the `LeaseplanAbocarExporter login` locally and mount your generated config into the container.
- Linux: 
`-v ~/.leaseplanabocar:/root/.leaseplanabocar`
- Windows: 
`-v C:\Users\[your user]\.leaseplanabocar:/root/.leaseplanabocar`

The Scripts store their generated data at `/opt/exporter/cache` when you want to inspect these files just mount any folder at this position:
- Linux: 
`-v ~/leaseplanCache:/opt/exporter/cache`
- Windows: 
`-v C:\TEMP\leaseplanCache:/opt/exporter/cache`

For the Telegram Notifications a telegram `API_KEY` as well as your `CHAT_ID` need to be provided.

To create a new bot chat with `@botfather` and follow the `/newbot` workflow. At the end you will be provided with an API_KEY for your newly created Bot. Just inject the Key into the Environment of the Docker Container `-e API_KEY=[your api key]`

To get your personal chat id just start a chat with `@username_to_id_bot` the bot will reply with your ID that you can inject into the Environment of the Docker Container as well `-e CHAT_ID=[your chat id]`

The complete command will look like this:
- Linux: 
```bash
docker run --name=leaseplan_monitor -itd -v ~/leaseplanCache:/opt/exporter/cache -v ~/leaseplanCache:/opt/exporter/cache -e API_KEY=[your api key] -e CHAT_ID=[your chat id] khase/leaseplanabocarexporter
```
- Windows: 
```bash
docker run --name=leaseplan_monitor -itd -v C:\Users\[your user]\.leaseplanabocar:/root/.leaseplanabocar -v C:\TEMP\leaseplanCache:/opt/exporter/cache -e API_KEY=[your api key] -e CHAT_ID=[your chat id] khase/leaseplanabocarexporter
```
