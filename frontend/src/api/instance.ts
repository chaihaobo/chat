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
        const token = localStorage.getItem('access_token');
        method.config.headers.Authorization = token;
    },
    responded: {
        onSuccess: async (response, _) => {
            let responseBody = await response.json() as Response<any>;
            if (response.status >= 400) {
                message.error(responseBody.message);
                throw new Error(responseBody.message)
            }
            return responseBody.data
        },
    }
});
