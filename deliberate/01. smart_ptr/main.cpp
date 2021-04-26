#include "smart_ptr.hpp"

#include <iostream>

using namespace std;


class Base {
public:
    Base() = default;
    virtual ~Base() {}

    virtual void foo() {
        cout << "Base foo" << endl;
    }
};



class Dirived : public Base {
public:
    Dirived() = default;
    void foo() {
        cout << "Dirived foo" <<endl;
    }
};



int main() {
    smart_ptr<int> ptr {new int(1024)};

    smart_ptr<int> gtr = ptr;

    cout << "ptr's ref count: " << ptr.use_count() << ", ";
    cout << "gtr's ref count: " << gtr.use_count() << endl;

    smart_ptr<int> etr;
    cout << "etr's ref count: " << etr.use_count() << endl;

    etr = gtr;

    cout << "ptr's ref count: " << ptr.use_count() << ", ";
    cout << "etr's ref count: " << etr.use_count() << endl;

    /* ----------------------------------------------------- */

    /*
     * 若只是单纯的使用类型参数 T 的话，下方的代码将无法执行
     * 因为无法将 smart_ptr<Dirived> 转换成 smart_ptr<Base>。因此，我们需要一个额外的模板参数来帮我们做这件事情
     * 
     * template <typename U>
     * smart_ptr(const smart_ptr<U>& other);
     * 
     * 但是，类模板的不同实例之间是不能直接访问其私有成员的，所以我们得把它变成友元
     * 
     * template <typename U>
     * friend class smart_ptr
     */
    smart_ptr<Dirived> derived {new Dirived()};
    smart_ptr<Base> another = derived;
    another->foo();     // Dirived foo, 多态性保留了下来

}