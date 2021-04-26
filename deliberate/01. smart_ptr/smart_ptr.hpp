#ifndef __SMART_PTR_H_
#define __SMART_PTR_H_

#include <algorithm>
#include <atomic>


class ref_count {
private:
    // m_count 的自增和自减需要线程安全，故使用 atomic
    std::atomic<long long> m_count;
public:
    ref_count() : m_count(1) {}

    void increase_count() {m_count++;}

    long long decrease_count() {return --m_count;}

    long long get_count() const {return m_count;}
};



template<typename T>
class smart_ptr {
private:
    T *_ptr;
    ref_count *_ref_count;
public:

    template<typename U>
    friend class smart_ptr;

    // 必须显式初始化，不允许进行隐式地类型转换
    explicit smart_ptr(T *ptr = nullptr) : _ptr(ptr) {
        if (ptr) {
            _ref_count = new ref_count();
        }
    }

    ~smart_ptr() {
        if (_ptr && _ref_count->decrease_count() == 0) {
            delete _ptr;
            delete _ref_count;
        }
    }

    // 拷贝构造函数
    smart_ptr(const smart_ptr& other) {
        _ptr = other._ptr;
        if (_ptr) {
            other._ref_count->increase_count();
            _ref_count = other._ref_count;
        }
    }

    /*
     * 用于子类指针向基类指针转换的情况
     */
    template<typename U>
    smart_ptr(const smart_ptr<U>& other) {
        _ptr = other._ptr;
        if (_ptr) {
            other._ref_count->increase_count();
            _ref_count = other._ref_count;
        }
    }

    // 移动构造函数
    template <typename U>
    smart_ptr(smart_ptr<U>&& other) noexcept {
        _ptr = other._ptr;
        if (_ptr) {
            _ref_count = other._ref_count;
            other._ptr = nullptr;
        }
    }

    // 拷贝赋值函数
    smart_ptr &operator= (const smart_ptr& rhs) {
        smart_ptr temp = smart_ptr{rhs};
        swap(temp);
        return *this;
    }


    T *get_ptr() const {return _ptr;}

    long long use_count() const {
        return _ptr ? _ref_count->get_count() : 0;
    }

    T *operator->() const {return _ptr;}
    T &operator*() const {return *_ptr;}
    operator bool() const {return _ptr;}

    void swap(smart_ptr& other) {
        using std::swap;
        swap(_ptr, other._ptr);
        swap(_ref_count, other._ref_count);
    }

};

#endif