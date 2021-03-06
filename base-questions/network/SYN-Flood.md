

## TCP 中 SYN 攻击是什么？如何防止？

SYN Flood 攻击是指利用了 TCP 协议的漏洞，攻击者操控多台傀儡机在短时间内向服务器发送海量伪造的带有 SYN 标志位的数据包，从而造成服务端状态为 `SYN_RECV` 的队列溢出，导致无法正常处理其它正常网络请求的一种 DoS 攻击。

![Alt text](../images/1619766123992.png)

在正常情况下，当服务端收到客户端发来的 SYN 数据包以后，会将 socket 状态更新为 `SYN_RECV` 并保存在监听队列的未完成连接队列中，并向客户端回送 SYN+ACK 数据包，而后客户端向服务端发送 ACK 确认包，服务端收到该确认包以后，socket 状态由 `SYN_RECV` 更新至 `ESTABLISHED`，表示 TCP 连接已建立，并将该 socket 从未完成连接队列中移入至已完成连接队列中，等待应用程序调用 `accept()` 取出。

当服务端遭受 SYN Flood 攻击时，完成连接队列很有可能会溢出，导致服务端无法再处理其它请求。所以这种攻击也被称之为 Dos 攻击，即 Denial Of Service Attack，拒绝服务攻击，使服务器拒绝向正常用户提供服务。

### SYN Flood 攻击防治方法

1. 启用 SYN Cookie

当我们将内核的 `net.ipv4.tcp_syncookies` 设置为 1 时，表示启用 SYN Cookie。此时当服务端收到 SYN 数据包以后，不再将该 socket 保存在未完成连接队列中，并且使用基于时间种子的随机方式进行生成 Seq 序列号，服务端将带有该 Seq 以及 SYN+ACK 的 TCP 数据包回送给客户端。服务端在收到 ACK 确认包以后，会使用 Cookie 检验算法来鉴定 Ack 确认号与发送的 Seq 序列号是否匹配，鉴定通过则完成握手，并将连接放置于已完成连接队列中，失败则直接丢弃。

2. 调整 SYN 最大队列长度

也可以通过设置 `net.ipv4.tcp_max_syn_backlog` 来将未完成连接队列的容量增大，使得半连接队列不会那么快的溢出。

3. 减少 SYN+ACK 的最大重试次数

傀儡机有时根本就不会去响应服务端发送的 SYN+ACK 数据包，那么服务端因为迟迟没有收到客户端发来的 ACK 确认包，会不断地重新发送 SYN+ACK 数据包，系统默认为 5 次。我们可以通过设置 `net.ipv4.tcp_synack_retries` 来调整最大重试次数，比如将其降为 2 或者是 3。

4. 使用 `iptables` 对数据包进行过滤

可以使用 `iptables` 来限制单个 IP 的最大并发数，单个地址的最大连接数，以及单位时间内的连接数。

5. 使用代理服务器进行拦截

一种分工的策略，代理服务器专门进行 DDos 的防护，使后方服务器能够专注于自身的业务，而无需对硬件或内核进行特殊配置。

![Alt text](../images/1620785460173.png)

以 Google Cloud Armor 为例，当用户向服务器发起 SYN 请求时，流量并不会直接进入到服务器，而是首先进入云平台的边缘节点，进行负载均衡、黑/白名单过滤，也就是说，这些边缘节点就像护城河一样为整座城提供防护功能。这样一来，开发人员可专注于实现自己的功能上面，而不需要再去担心网络防护问题。