# [float32.app](https://float32.app)：现代 AI 驱动的搜索助手

float32 是一个基于大语言模型驱动的搜索助手，它可以帮助你快速找到你想要的答案。依赖 RAG 技术，float32 可以获取互联网信息以提供更准确的答复。

> [!TIP]
> 尝试一下 [float32.app](https://float32.app)。
> 如果你对 float32 的 prompt 感兴趣，可以看看 [📁 ./prompt/*.promptc](prompt/)。
> 里面的 `.promptc` 文件记录了你感兴趣的东西。

## Acknowledged

> [!NOTE]
> <a href="https://limit.dev"><img width="50" height="50" align="left" style="float: left; margin: 0 0px 0 0; margin-left: 10px;" alt="Limit-LAB logo" src="https://avatars.githubusercontent.com/u/116355161?s=200&v=4"></a>
> This project is affiliated to [Limit-LAB](https://github.com/Limit-LAB).  
> Special thanks to [@ZincCat](https://github.com/zinccat).

## Server 端

float32.app 支持自托管，你可以在本地搭建一个 float32.app 服务。服务相关代码可以参考 [backend/exec/svr/...](backend/exec/svr)

为了能运行，你需要
- OpenAI API 服务  
  **必须。** 包括 API Key 和 EndPoint（如适用）
- Serper.dev 服务  
  **必须。** 用于获取搜索引擎结果。包括一个 API Key。
- MySQL 数据库  
  **可选。** 用于历史服务，如没有 MySQL 数据库，则历史服务/分享服务不可用。

> [!TIP]
> 一键部署？试试看 [deploy.sh](deploy.sh)。
> ```bash
> bash deploy.sh
> ```

### .env 文件

.env 文件用于配置服务，你可以参考 [backend/exec/svr/README.md](backend/exec/svr/README.md) 的描述与 [backend/exec/svr/init.go](backend/exec/svr/init.go) 中的实现代码。

```env
DEBUG=1                                     # Optional. 1 or 0, 1 means debug mode on.
LISTEN_ADDR=127.0.0.1:1145                  # Optional. default 0.0.0.0:8080
DB_URL=mysql://root:root@localhost:3306/xxx # Optional
OPENAI=sk-xfeusx233fchwwe239430xxxxxxxxx    # Mandatory. OpenAI API Key
OPENAI_ENDPOINT=http://localhost:5000       # Optional
SERP_DEV=183fjcs92fwewefhwiu382d8uwjcncsk   # Mandatory. serper.dev's API KEY
ALLOW_ORIGINS=https://float32.app           # Optional, CORS, split by space
```

## Prompt 与 PromptC 文件

float32.app 使用 promptc 标准来实践 prompt 开发的解耦。请参阅 [promptc.dev](https://promptc.dev/) 与 [promptc-go](https://github.com/promptc/promptc-go) 获得更多信息。

所有 float32.app 使用的 prompt 都位于 [prompt](prompt) 目录下。它们包括：

- [`code.promptc`](prompt/code.promptc)：代码相关 prompt
- [`med.promptc`](prompt/med.promptc)：医学相关的 prompt

目前所有的 prompt 是基于基础模板 [`base.promptc`](prompt/base.promptc) 使用 `sed` 与 [`generate.sh`](prompt/generate.sh) 生成的。

> [!WARNING]
> 如果你是 macOS，则需要安装 `gsed` 以运行 `generate.sh`。  
> ```bash
> brew install gnu-sed
> ```

## 前端

前端使用 pnpm + React + Vite + MobX + TDesign 的结构。请使用以下命令以启动开发服务器：

```bash
cd frontend
bash tdesign.sh # 获取 TDesign 资源
pnpm i
pnpm dev
```
