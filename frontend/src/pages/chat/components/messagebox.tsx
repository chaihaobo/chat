import {FC, Ref, useEffect, useImperativeHandle, useRef, useState} from "react";
import {Avatar} from "antd";
import styled from "styled-components";
import useUserInfo from "../../../hooks/useUserInfo.ts";

export interface MessageBoxProps {
    items: MessageBoxItem[]
    ref: Ref<MessageboxRef>
    className?: string
}

export interface MessageBoxItem {
    senderID: number
    senderAvatar: string
    senderName: string
    content: string
}

export interface MessageboxRef {
    scrollToBottom: () => void
}


const Messagebox: FC<MessageBoxProps> = ({items, className, ref}) => {
    let {userInfo} = useUserInfo();
    const messageContainerRef = useRef<HTMLDivElement>(null);
    let [scrollCounter, setScrollCounter] = useState<number>(0);

    const scrollToBottom = () => {
        if (messageContainerRef.current) {
            messageContainerRef.current.scrollTop = messageContainerRef.current.scrollHeight;
        }
    };
    useImperativeHandle(ref, () => {
        return {
            scrollToBottom: () => {
                setScrollCounter(prev => prev + 1);
            }
        }
    })
    useEffect(() => {
        scrollToBottom()
    }, [scrollCounter]);
    return (
        <div className={className} ref={messageContainerRef}>
            {items.map((item, index) => (
                <div
                    key={index}
                    className={`message-item ${item.senderID === userInfo?.id ? 'message-self' : ''}`}
                >
                    <Avatar src={item.senderAvatar} className="avatar"/>
                    <div className="message-content">
                        <div className="sender-name">{item.senderName}</div>
                        <div className="message-bubble">{item.content}</div>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default styled(Messagebox)`
    display: flex;
    flex-direction: column;
    gap: 16px;
    height: 100%;
    overflow-y: auto;
    scroll-behavior: smooth;
    min-height: 0;

    .message-item {
        display: flex;
        gap: 12px;
        align-items: flex-start;

        &.message-self {
            flex-direction: row-reverse;

            .message-content {
                align-items: flex-end;
            }

            .message-bubble {
                background-color: #1890ff;
                color: white;
                border-radius: 8px 2px 8px 8px;
            }

            .sender-name {
                text-align: right;
            }
        }
    }

    .message-content {
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .sender-name {
        font-size: 12px;
        color: #999;
    }

    .message-bubble {
        background-color: #f5f5f5;
        padding: 8px 12px;
        border-radius: 2px 8px 8px 8px;
        max-width: 560px;
        word-wrap: break-word;
    }
`;