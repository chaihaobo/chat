import {alovaInstance} from './instance';

export interface LoginResponse {
    id: number;
    avatar: string;
    access_token: string;
    refresh_token: string;
}

export interface PasswordLoginRequest {
    username: string;
    password: string;
}

export const login = async (code: string) => {
    return alovaInstance.Post<LoginResponse>('/user/login', {code});
};

export const loginWithPassword = async (data: PasswordLoginRequest) => {
    return alovaInstance.Post<LoginResponse>('/user/login/password', data);
};