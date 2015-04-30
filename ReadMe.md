#the weather server service

It get the weather info of a city when providing with the city name
It was built as a thrift server.

## Directory Architecture

+ conf/  the system's configured file it contains the server listen port info.
+ handler/ the realize of each service's function.
+ thrift_interface/  auto generated thrift interface
+ utils/ commonly used tools like system configure
