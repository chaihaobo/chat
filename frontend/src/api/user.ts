import {alovaInstance} from './instance';

export interface LoginResponse {
    id: number;
    avatar: string;
    username: string;
    access_token: string;
    refresh_token: string;
}

export interface GetUserInfoResponse {
    id: number;
    username: string;
    avatar: string;
}

export interface PasswordLoginRequest {
    username: string;
    password: string;
}

export interface Friend {
    id: number;
    username: string;
    avatar: string;
}

export const login = async (code: string) => {
    return alovaInstance.Post<LoginResponse>('/user/login', {code});
};

export const loginWithPassword = async (data: PasswordLoginRequest) => {
    return alovaInstance.Post<LoginResponse>('/user/login/password', data);
};

export const getFriends = () => {
    return alovaInstance.Get<Friend[]>('/user/friends');
};

export const getUserInfo = () => {
    return alovaInstance.Get<GetUserInfoResponse>('/user/info');
};