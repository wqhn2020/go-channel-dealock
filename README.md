# go-channel-dealock

为什么我这个程序会报Deadlock?
按我的理解，从程序运行后的第4秒开始，通道有3个元素，满了，程序就会一直阻塞（在第16行代码）, 不应该报dealock.
