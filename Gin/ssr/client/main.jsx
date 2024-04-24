import React from 'react'; // 确保导入了 React
import {createRoot} from 'react-dom/client'; // 导入 createRoot
import App from './App'; // 导入 App 组件

const container = document.getElementById('app');  // 获取容器元素
const root = createRoot(container);  // 创建 root
root.render(<App/>);  // 使用 root.render 方法渲染 App 组件
