#include "blocking_queue.hpp"

#include <thread>
#include <iostream>

using namespace std;


void producer(blocking_queue<int>& queue) {
    for (int i = 0; i < 10; i++) {
        queue.push(i);
        cout << "生产数据完毕: " << i << endl;
    }
}

void consumer(blocking_queue<int>& queue) {
    for (int i = 0; i < 10; i++) {
        int res = queue.pop();
        cout << "开始消费数据: " << res << endl;
        std::this_thread::sleep_for(1s);
    }
}

int main() {
    blocking_queue<int> queue;

    thread t1 {producer, std::ref(queue)};
    thread t2 {consumer, std::ref(queue)};

    t1.join();
    t2.join();

    cout << "Task Done" << endl;
}
