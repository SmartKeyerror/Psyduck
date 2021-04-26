#include <signal.h>
#include <errno.h>
#include <unistd.h>
#include <stdio.h>


void signal_handler(int sig) {
    // errno 虽然是线程安全的，但是内核在调用信号处理函数时，并不会帮我们恢复 errno 的值
    // 因此我们需要用一个临时变量来保存进入信号处理函数时的 errno，处理结束后再进行恢复
    int prev = errno;

    sigset_t current;
    sigprocmask(SIG_BLOCK, NULL, &current);     // 获取当前进程信号掩码

    // 当我们因为收到某一个信号而进入信号处理函数时，该信号会被默认的添加至当前进程信号掩码中，
    // 防止出现递归打断信号处理函数执行的情况
    if (sigismember(&current, sig)) {

        // printf() 为不可重入函数，但是 read/write 以及 recv/send 等方法为可重入函数
        // printf, fread、fwrite 等带有用户 I/O 缓冲区的标准库函数为不可重入函数，应避免使用
        
        char msg[] = "SIGINT 已被添加至进程信号掩码集中\n";
        write(STDOUT_FILENO, msg, sizeof(msg));
    }

    errno = prev;
}

int main() {
    sigset_t set;
    sigemptyset(&set);
    
    struct sigaction action = {
        signal_handler, 
        set,
        SA_RESTART
    };

    if (sigaction(SIGINT, &action, NULL) == -1) {   // 为 SIGINT 注册信号处理函数
        perror("SIGINT 信号处理函数注册失败");
        return 1;
    }

    while (1) {
        printf("main function \n");
        int remain = sleep(5);
        
        if (remain) {                       // sleep 可以被打断，并且不会重新睡眠，而是返回剩余的睡眠时间
            printf("sleep 被打断，还剩 %d 秒睡眠时间 \n", remain);
            break;
        }
    }
    return 0;
}
