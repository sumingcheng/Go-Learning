const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = '../proto/todolist.proto';

// 加载.proto文件
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
    });

// 从加载的packageDefinition中创建客户端对象
const todoListPackage = grpc.loadPackageDefinition(packageDefinition);

// 创建客户端实例
const todoListClient = new todoListPackage.TodoList(
    'localhost:8080',
    grpc.credentials.createInsecure()
);

module.exports = todoListClient;
