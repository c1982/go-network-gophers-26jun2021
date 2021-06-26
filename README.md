# go-network-gophers-26jun2021
https://kommunity.com/goturkiye/events/go-ile-network-programlama-97094b41


# Go Network Programiing

- Oğuzhan YILMAZ @ Masomo

## Content

- TCP Server
- TCP Client
- Custom Protocol
- Profile
- Optimzation
    - Trace
    - CPU Profil
    - Connection Pool (sync.Pool, WorkerPool)
    - Epoll + Netpoll (non-blocking)
    - SO_REUSEPORT 
    - eBPF
    - ARM (1.16)
- Finish

client <---------------> server
  |                         |
  |__________stream w/r_____|


0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string
type    | length  | data

Default:
Duration time: 2.076635674s

512byte
duration time: 1.285418311s

unsafepointer
duration time: 1.223559957s

disable pprof
duration time: 1.205238147s