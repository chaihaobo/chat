import {createAlova} from 'alova';
import adapterFetch from 'alova/fetch';
import {message} from "antd";
import ReactHook from 'alova/react';

interface Response<T> {
    code: string
    message: string
    data: T
}


export const alovaInstance = createAlova({
    baseURL: '/api',
    requestAdapter: adapterFetch(),
    statesHook: ReactHook,
    beforeRequest: async (method) => {
        const token = JSON.parse(localStorage.getItem('user_info') || '{}').access_token;
        method.config.headers.Authorization = token;
    },
    responded: {
        onSuccess: async (response, _) => {
            let responseBody = await response.json() as Response<any>;
            // 如果是未登录错误，跳转到/login页面
            if (response.status == 401) {
                window.location.href = `/login?source=${window.location.pathname}`;
            }

            if (response.status >= 400) {
                message.error(responseBody.message);
                throw new Error(responseBody.message)
            }
            return responseBody.data
        },
    }
});
