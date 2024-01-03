# Server 端

## .env 文件

```ini
DEBUG=1                                     # Optional. 1 or 0, 1 means debug mode on.
LISTEN_ADDR=127.0.0.1:1145                  # Optional, default 0.0.0.0:8080

DB_URL=mysql://root:root@localhost:3306/xxx # Optional

OPENAI=sk-xfeusx233fchwwe239430xxxxxxxxx    # Mandatory. OpenAI API Key
OPENAI_ENDPOINT=http://localhost:5000       # Optional

SERP_DEV=183fjcs92fwewefhwiu382d8uwjcncsk   # Mandatory, serper.dev's API KEY
```