FROM golang:1.19-alpine
ENV TERM xterm-256color
ENV COLORTERM truecolor
RUN go install github.com/cedricblondeau/world-cup-2022-cli-dashboard@latest
ENTRYPOINT world-cup-2022-cli-dashboard
