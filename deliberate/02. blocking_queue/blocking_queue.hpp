
// 玩具级别的阻塞队列，condition_variable 的一个简单用例

#ifndef __BLOCKING_QUEUE_H_
#define __BLOCKING_QUEUE_H_

#include <thread>
#include <condition_variable>
#include <queue>

template <typename T>
class blocking_queue {
private:
    size_t m_capacity;
    std::queue<T> m_queue;
    
    std::mutex m_mutex;
    std::condition_variable m_not_empty;
    std::condition_variable m_not_full;
public:
    explicit blocking_queue(int capacity = 2) : m_capacity(capacity) {}

    void push(const T& item) {
        std::unique_lock<std::mutex> lock_guard{m_mutex};

        m_not_full.wait(lock_guard, [this] {
            return this->m_queue.size() != this->m_capacity;
        });

        m_queue.push(item);

        m_not_empty.notify_one();
    }

    T &pop() {
        std::unique_lock<std::mutex> lock_guard{m_mutex};

        m_not_empty.wait(lock_guard, [this] {
            return this->m_queue.size() != 0;
        });

        T &item = m_queue.front();
        m_queue.pop();

        m_not_full.notify_one();

        return item;
    }
};

#endif