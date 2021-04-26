

## :card_file_box: 目录

- [C++](#:telescope:C++)
- [Database](:bento:Database)
- [Linux Network Programing](:fire:Linux Network Programing)


## :telescope:C++

#### common

- [const](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/common/1.%20const.pdf)

#### 面向对象程序设计

TODO    

#### 拷贝控制

- [拷贝构造与拷贝赋值](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/copy-control/1.%20%E6%8B%B7%E8%B4%9D%E6%9E%84%E9%80%A0%E4%B8%8E%E6%8B%B7%E8%B4%9D%E8%B5%8B%E5%80%BC.pdf)
- [左值与右值](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/copy-control/2.%20%E5%B7%A6%E5%80%BC%E4%B8%8E%E5%8F%B3%E5%80%BC.pdf)
- [移动构造与移动赋值](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/copy-control/3.%20%E7%A7%BB%E5%8A%A8%E6%9E%84%E9%80%A0%E4%B8%8E%E7%A7%BB%E5%8A%A8%E8%B5%8B%E5%80%BC.pdf)
- [三/五法则](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/copy-control/4.%20%E4%B8%89%E4%BA%94%E6%B3%95%E5%88%99.pdf)

#### 模板与泛型编程

- [Template 的基本使用](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/template/1.%20Template.pdf)
- [万能引用与类型推断问题](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/template/2.%20%E4%B8%87%E8%83%BD%E5%BC%95%E7%94%A8%E4%B8%8E%E7%B1%BB%E5%9E%8B%E6%8E%A8%E6%96%AD%E9%97%AE%E9%A2%98.pdf)
- [完美转发](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/template/3.%20%E5%AE%8C%E7%BE%8E%E8%BD%AC%E5%8F%91.pdf)

#### 类型相关

- [强制类型转换](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/type/1.%20%E5%BC%BA%E5%88%B6%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2.pdf)
- [自动类型推断（auto）](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/type/2.%20%E8%87%AA%E5%8A%A8%E7%B1%BB%E5%9E%8B%E6%8E%A8%E6%96%AD%EF%BC%88auto%EF%BC%89.pdf)
- [decltype](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/type/3.%20decltype.pdf)

#### 函数编程

- [lambda 表达式](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/functional/1.%20lambda%20%E8%A1%A8%E8%BE%BE%E5%BC%8F.pdf)
- [可调用对象——std::function](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/functional/2.%20%E5%8F%AF%E8%B0%83%E7%94%A8%E5%AF%B9%E8%B1%A1%20function.pdf)

#### 智能指针

- [堆、栈与 RAII: C++ 管理资源的方式](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/smart-ptr/1.%20%E5%A0%86%E3%80%81%E6%A0%88%E4%B8%8E%20RAII%20%3A%20C%2B%2B%20%E7%AE%A1%E7%90%86%E8%B5%84%E6%BA%90%E7%9A%84%E6%96%B9%E5%BC%8F.pdf)
- [RAII 与智能指针](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/smart-ptr/2.%20RAII%E4%B8%8E%E6%99%BA%E8%83%BD%E6%8C%87%E9%92%88.pdf)
- [shared_ptr](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/smart-ptr/3.%20shared_ptr.pdf)

#### 并发编程

- [Linux 进程、线程与调度](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/1.%20Linux%20%E8%BF%9B%E7%A8%8B%E3%80%81%E7%BA%BF%E7%A8%8B%E4%B8%8E%E8%B0%83%E5%BA%A6.pdf)
- [线程的创建与执行](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/2.%20%E7%BA%BF%E7%A8%8B%E7%9A%84%E5%88%9B%E5%BB%BA%E4%B8%8E%E6%89%A7%E8%A1%8C.pdf)
- [std::async 与 std::future](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/3.%20async%E4%B8%8Efuture.pdf)
- [互斥量与 std::lock_guard、std::unique_lock](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/4.%20%E4%BA%92%E6%96%A5%E9%87%8F%E3%80%81lock_guard%E4%B8%8Eunique_lock.pdf)
- [通知状态的改变——POSIX 条件变量](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/5.%20%E9%80%9A%E7%9F%A5%E7%8A%B6%E6%80%81%E7%9A%84%E6%94%B9%E5%8F%98%E2%80%94POSIX%E6%9D%A1%E4%BB%B6%E5%8F%98%E9%87%8F.pdf)
- [通知状态的改变——C++ 条件变量](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/6.%20%E9%80%9A%E7%9F%A5%E7%8A%B6%E6%80%81%E7%9A%84%E6%94%B9%E5%8F%98%E2%80%94C%2B%2B%E6%9D%A1%E4%BB%B6%E5%8F%98%E9%87%8F.pdf)
- [POSIX 信号量](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/concurrent/7.%20POSIX%20%E4%BF%A1%E5%8F%B7%E9%87%8F.pdf)

#### 编译与链接

- [编译与链接](https://smartkeyerror.oss-cn-shenzhen.aliyuncs.com/Phyduck/compile/1.%20%E7%BC%96%E8%AF%91%E4%B8%8E%E9%93%BE%E6%8E%A5.pdf)

------

## :bento:Database


------

## :fire:Linux Network Programing