import {useEffect, useRef} from "react";
import useUserInfo from "./useUserInfo.ts";

export interface ChartConnectionHooks {
    onFail: (e: Event) => void
    onOpen: (e: Event) => void
    onMessage: (e: MessageEvent) => void

}

export interface ChatConnection {
    send: (message: string) => void
}

const useChatConnection = (actions: ChartConnectionHooks): ChatConnection => {
    let {userInfo} = useUserInfo()
    const wsRef = useRef<WebSocket | null>(null)
    const onConnectFailed = (e: Event) => {
        actions.onFail(e)
    }
    let ws: WebSocket | null = null
    useEffect(() => {
        ws = new WebSocket("ws://localhost:2222/ws?token=" + userInfo?.access_token);
        wsRef.current = ws
        ws.onmessage = (event) => {
            actions.onMessage(event)
        };
        ws.onerror = (e) => {
            console.log(e)
            onConnectFailed(e)
        };
        ws.onopen = (e) => {
            actions.onOpen(e)
        }
        return () => {
            if (!ws) {
                return
            }
            ws.close();
            wsRef.current = null

        }
    }, [userInfo?.access_token]);
    return {
        send: (message) => {
            if (!wsRef.current) {
                return
            }
            wsRef.current.send(message)
        }
    }

}

export default useChatConnection