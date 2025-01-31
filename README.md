# Archived
Since 2023, many apps have come out and offer a better user experience than this project. I don't see the point in maintaining it anymore.

# Talk

Talk is a single-page application crafted to converse with AI using voice, replicating the user experience akin to a
native app.

**[Demo](https://talk.legendy4141.com)**  (No registration or login needed. Simply start conversing. For an optimal
experience, open in Chrome)

![figjam-talk-2023-09-22.png](doc/figjam-talk-2023-09-22.png)

<div align="center">

<a  href="https://www.figma.com/file/4Uhzr87YVN1KR5ayz7WDYm/Talk?type=whiteboard&t=OGwMZMSKsWZIwG0T-1" target="_blank">
More details
</a>
</div>

# Highlighted Features

- Broad range of service providers to choose from: ChatGPT, Google Gemini, Elevenlabs, Google Text-toSpeech, Whisper and
  Google
  Speech-to-Text
- Enable voice-driven dialogues
- Modern and stylish user interface
- Unified, standalone binary

# How to use

## 1. Prepare a `talk.yaml` file.

Here is a [simple example](example/talk.simple.example.yaml) utilising ChatGPT, Whisper and
Elevenlabs:

```yaml
speech-to-text:
  whisper: open-ai-01

llm:
  chat-gpt: open-ai-01

text-to-speech:
  elevenlabs: elevenlabs-01

# provide your confidential information below.
creds:
  open-ai-01: "sk-2dwY1IAeEysbnDNuAKJDXofX1IAeEysbnDNuAKJDXofXF5"
  elevenlabs-01: "711sfpb9kk15sds8m4czuk5rozvp43a4"
```

* Not interested in Voice? Give this a try:
```yaml
llm:
  chat-gpt: open-ai-01
creds:
  open-ai-01: "sk-2dwY1IAeEysbnDNuAKJDXofX1IAeEysbnDNuAKJDXofXF5"
```

* Looking to utilise Google Gemini, Google Text-to-Speech and Google Speech-to-Text? Not to worry, we have that covered.
Please refer
to [talk.google.example.yaml](example/talk.google.example.yaml) for more information

* The comprehensive example: [talk.full.example.yaml](example/talk.full.example.yaml)

## 2. Start the application

### Docker

```shell
docker run -it -v ./talk.yaml:/etc/talk/talk.yaml -p 8000:8000 legendy4141/talk
```

### Terraform

Refer to [terraform](example/terraform). The same applies to Kubernetes.

### From scratch

```shell
# clone projects
git clone https://github.com/legendy4141/talk.git legendy4141/talk
git clone https://github.com/legendy4141/talk-web.git legendy4141/talk-web

# build web with yarn and copy; currently using node v20.3.0 
cd legendy4141/talk-web && make copy

# build backend
cd ../talk && make build

# run
./talk --config ./talk.yaml
# or simply `./talk` as it automatically lookup talk.yaml in `/etc/talk/talk.yaml` and `./talk.yaml`
./talk
```

# Advanced usage

### Proxy

We honour `HTTP_PROXY` and `HTTPS_PROXY` env variables. Given that all communication between the Talk server and
service providers occurs via HTTPS, simply employ `HTTPS_PROXY`.

```shell
docker run -it -v ./talk.yaml:/etc/talk/talk.yaml \
-e HTTPS_PROXY=http://192.168.1.105:7890 \
-p 8000:8000 \
legendy4141/talk
```

### Log level

Default log level is `info`, Use env `LOG_LEVEL` to change log level: "debug", "info", "warn", "error", "dpanic", "
panic", and "fatal". e.g.,

```shell
LOG_LEVEL=debug ./talk
```

### HTTPS

`legendy4141/talk` offers three methods for enabling HTTPS.

#### 1. Generate self-signed cert on the fly

Example: [talk.tls.self.signed.example.yaml](example/talk.tls.self.signed.example.yaml)

```yaml
server:
  tls:
    self-signed: true
```

This is handy if you're indifferent to a domain and unconcerned about security, simply desiring to enable
microphone access on browsers.

##### 2. Provide your own TLS

Example: [talk.tls.provided.example.yaml](example/talk.tls.provided.example.yaml)

##### 3. Auto TLS

This configuration example facilitates automatic certificate acquisition from
LetsEncrypt: [talk.tls.auto.example.yaml](example/talk.tls.auto.example.yaml)

Requirements: You should have your personal VPS and domain.

# Troubleshooting

### Why can't I start the recording?

Web browsers safeguard your microphone from being accessed by non-HTTPS websites for security reasons, with the
exceptions being `localhost` and `127.0.0.1`.

Here are some possible solutions:

1. Enable [HTTPS](#https). Particularly, you
   can [Generate self-signed cert on the fly](#generate-self-signed-cert-on-the-fly) in a mere second.
2. Run Talk through a reverse proxy like Nginx and set up TLS within this service.
3. In Chrome, go to `chrome://flags/`, find `Insecure origins treated as secure`, and enable it:
   <br>
   <img src="./doc/image/chrome-microphone-access.jpg" alt="Markdownify" width="600">
   <br>

# Browser compatibility

|            | [Arc](https://arc.net/) | Chrome | FireFox | Edge | Safari |
|:----------:|:-----------------------:|:------:|:-------:|:----:|:------:|
| Microphone |            ✅            |   ✅    |    ✅    |  ❌   |   ❌    |
|     UI     |            ✅            |   ✅    |    ✅    |  ✅   |   ❌    |

# Q&A

**Q: Why not use TypeScript for both the frontend and backend development?**

A:

* When I embarked on this project, I was largely inspired by [Hugh](https://github.com/IgnoranceAI/hugh), a project
  primarily coded in Python, supplemented with HTML and a touch of JavaScript. To broaden the horizons of text-to-speech
  providers, I revamped the backend logic using Go, transforming it into a Go-based project.
* Crafting backend logic with Go feels incredibly intuitive—it distills everything down to a single binary.
* Moreover, my skills in frontend development were somewhat rudimentary at that time.

**Q: Will a mobile browser-friendly version be made available?**

A: Streamlining the website for mobile usage would be a time-intensive endeavour and, given my current time constraints,
it isn't the primary concern. As it stands, the site performs optimally on desktop browsers based on the Chromium
Engine, with certain limitations on browsers such as Safari.

# Roadmap

- [x] Google TTS
- [x] Google STT
- [x] OpenAI Whisper STT
- [x] Setting language, speed, stability, etc
- [x] Choose voice
- [x] Docker image
- [x] Server Side Events(SSE)
- [x] More LLMs other than ChatGPT
- [x] Download and import text history
- [x] Download chat MP3
- [x] Prompt template
- [ ] Dark mode

# Contributing

We're in the midst of a dynamic development stage for this project and warmly invite new contributors.

[CONTRIBUTING.md](CONTRIBUTING.md)

# Credits

### Front-end

* [React](https://github.com/facebook/react): The library for web and native user interfaces
* [vite](https://github.com/vitejs/vite): Next generation frontend tooling. It's fast!
* [valtio](https://github.com/pmndrs/valtio): Valtio makes proxy-state simple for React and Vanilla
* [wavesurfer.js](https://github.com/katspaugh/wavesurfer.js): Audio waveform player
* [granim.js](https://github.com/sarcadass/granim.js): Create fluid and interactive gradient animations with this small
  javascript library.
* [virtual](https://github.com/tanstack/virtual): Headless UI for Virtualizing Large Element Lists in JS/TS, React,
  Solid, Vue and Svelte
* [markdown-it](https://github.com/markdown-it/markdown-it): Markdown parser, done right. 100% CommonMark support,
  extensions, syntax plugins & high speed
* [highlight.js](https://github.com/highlightjs/highlight.js): JavaScript syntax highlighter with language
  auto-detection and zero dependencies.

### Back-end

* This project draws inspiration from [Hugh](https://github.com/IgnoranceAI/hugh), a remarkable tool that enables
  seamless communication with AI using minimal code.
* [go-openai](https://github.com/sashabaranov/go-openai): OpenAI ChatGPT, GPT-3, GPT-4, DALL·E, Whisper API wrapper for
  Go.
* [google-cloud-go](https://github.com/googleapis/google-cloud-go): Google Cloud Client Libraries for Go. Thanks to
  [googleapis](https://github.com/googleapis) for the prompt response to our concern.
* [echo](https://github.com/labstack/echo): High performance, minimalist Go web framework
* [elevenlabs-go](https://github.com/haguro/elevenlabs-go): A Go API client library for the ElevenLabs speech synthesis
* [r3labs/sse](https://github.com/r3labs/sse/): Server Sent Events server and client for Golang platform.

### Design

* [wikiart.org](https://www.wikiart.org): Wikiart is a great place to find art online. Most wallpapers of Talk come
  from WikiArt.org
* [Arc](https://arc.net/): Arc is the Chrome replacement I’ve been waiting
  for -- [THE VERGE](https://www.theverge.com/23462235/arc-web-browser-review)
* [grainy-gradients](https://github.com/cjimmy/grainy-gradients): Thanks to [cjimmy](https://github.com/cjimmy/) for his
  amazing [tutorial](https://css-tricks.com/grainy-gradients/) on noise and gradient background
* [Signal-Desktop](https://github.com/signalapp/Signal-Desktop)
  and [Signal-iOS](https://github.com/signalapp/Signal-iOS): Private messengers. Much of the inspiration for the UI
  comes from Signal.

We would also like to thank all other open-source projects and communities not listed here for their valuable
contributions to our project.
