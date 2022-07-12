import axios from 'axios';
import { getItem } from './storage';
import { isCheckTimeout } from './time';

const service = axios.create({
    baseURL: 'http://127.0.0.1:8888',
    timeout: 5000
});

service.interceptors.request.use(
    config => {
        // 在这个位置需要统一的去注入token
        if (getItem('token')) {
            if (isCheckTimeout()) {
                return Promise.reject(new Error('token 失效'))
            }
            // 如果token存在 注入token
            config.headers.Authorization = `Bearer ${getItem('token')}`
        }
        return config;
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

service.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return response.data;
        } else {
            Promise.reject();
        }
    },
    error => {
        console.log(error);
        return Promise.reject();
    }
);

export default service;
