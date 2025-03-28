import {alovaInstance} from "./instance.ts";

export interface GetRecentlyMessagesRequest {
    offset: number,
    limit: number,
    friend_user_id?: number
}

export interface GetRecentlyMessagesResponse {
    has_more: number
    messages: RecentlyMessageItem[]
}

export interface RecentlyMessageItem {
    id: number
    from: number
    to: number
    content: string
}

export const getRecentlyMessages = async (request: GetRecentlyMessagesRequest) => {
    return alovaInstance.Get<GetRecentlyMessagesResponse>('/messages/recently', {
        params: request
    });
};